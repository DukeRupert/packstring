#!/usr/bin/env bash
set -euo pipefail

# Image optimization script for Packstring
# Converts source JPG/PNG images to responsive WebP variants.
#
# Source directories:  static/img/source/{hero,trips,gallery}/
# Output directories:  static/img/{hero,trips,gallery}/
#
# Each image produces 3 variants: 400w, 800w, 1600w
# Gallery images also produce a 400x300 crop thumbnail (-thumb.webp)
# Hero images use quality 85; all others use quality 80.
# If a file exceeds 200KB, it is re-encoded at progressively lower quality.

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
SOURCE_DIR="$PROJECT_ROOT/static/img/source"
OUTPUT_BASE="$PROJECT_ROOT/static/img"

WIDTHS=(400 800 1600)
QUALITY=80
HERO_QUALITY=85
MAX_SIZE=204800  # 200KB in bytes

# Use magick (IMv7) if available, fall back to convert
if command -v magick &>/dev/null; then
    CONVERT="magick"
else
    CONVERT="convert"
fi

WARN_COUNT=0
PROCESSED=0

encode_webp() {
    local tmp_png="$1"
    local out="$2"
    local quality="$3"

    cwebp -q "$quality" -quiet "$tmp_png" -o "$out"

    # Re-encode at lower quality if file exceeds 200KB
    local size
    size=$(stat -c%s "$out" 2>/dev/null || stat -f%z "$out" 2>/dev/null)
    local q="$quality"
    while [ "$size" -gt "$MAX_SIZE" ] && [ "$q" -gt 40 ]; do
        q=$(( q - 10 ))
        cwebp -q "$q" -quiet "$tmp_png" -o "$out"
        size=$(stat -c%s "$out" 2>/dev/null || stat -f%z "$out" 2>/dev/null)
    done

    if [ "$size" -gt "$MAX_SIZE" ]; then
        echo "  WARNING: $out is $(( size / 1024 ))KB (still exceeds 200KB at q=$q)"
        WARN_COUNT=$(( WARN_COUNT + 1 ))
    fi
}

process_image() {
    local src="$1"
    local out_dir="$2"
    local quality="$3"
    local is_gallery="$4"

    local basename
    basename="$(basename "$src")"
    local name="${basename%.*}"

    for w in "${WIDTHS[@]}"; do
        local out="$out_dir/${name}-${w}w.webp"
        local tmp="$out_dir/${name}-${w}w.tmp.png"
        $CONVERT "$src" -resize "${w}x>" -strip "$tmp"
        encode_webp "$tmp" "$out" "$quality"
        rm -f "$tmp"
    done

    # Gallery thumbnail: 400x300 center crop
    if [ "$is_gallery" = "yes" ]; then
        local thumb="$out_dir/${name}-thumb.webp"
        local tmp="$out_dir/${name}-thumb.tmp.png"
        $CONVERT "$src" -resize "400x300^" -gravity center -extent 400x300 -strip "$tmp"
        encode_webp "$tmp" "$thumb" "$quality"
        rm -f "$tmp"
    fi

    PROCESSED=$(( PROCESSED + 1 ))
}

echo "=== Packstring Image Optimization ==="
echo "Using: $CONVERT"
echo ""

# Hero images
HERO_SRC="$SOURCE_DIR/hero"
HERO_OUT="$OUTPUT_BASE/hero"
mkdir -p "$HERO_OUT"
if compgen -G "$HERO_SRC"/*.{jpg,jpeg,png,JPG,JPEG,PNG} > /dev/null 2>&1; then
    echo "Processing hero images..."
    for f in "$HERO_SRC"/*.{jpg,jpeg,png,JPG,JPEG,PNG}; do
        [ -f "$f" ] || continue
        echo "  $(basename "$f")"
        process_image "$f" "$HERO_OUT" "$HERO_QUALITY" "no"
    done
fi

# Trip images
TRIPS_SRC="$SOURCE_DIR/trips"
TRIPS_OUT="$OUTPUT_BASE/trips"
mkdir -p "$TRIPS_OUT"
if compgen -G "$TRIPS_SRC"/*.{jpg,jpeg,png,JPG,JPEG,PNG} > /dev/null 2>&1; then
    echo "Processing trip images..."
    for f in "$TRIPS_SRC"/*.{jpg,jpeg,png,JPG,JPEG,PNG}; do
        [ -f "$f" ] || continue
        echo "  $(basename "$f")"
        process_image "$f" "$TRIPS_OUT" "$QUALITY" "no"
    done
fi

# Gallery images
GALLERY_SRC="$SOURCE_DIR/gallery"
GALLERY_OUT="$OUTPUT_BASE/gallery"
mkdir -p "$GALLERY_OUT"
if compgen -G "$GALLERY_SRC"/*.{jpg,jpeg,png,JPG,JPEG,PNG} > /dev/null 2>&1; then
    echo "Processing gallery images..."
    for f in "$GALLERY_SRC"/*.{jpg,jpeg,png,JPG,JPEG,PNG}; do
        [ -f "$f" ] || continue
        echo "  $(basename "$f")"
        process_image "$f" "$GALLERY_OUT" "$QUALITY" "yes"
    done
fi

echo ""
echo "Done. Processed $PROCESSED images."
if [ "$WARN_COUNT" -gt 0 ]; then
    echo "WARNING: $WARN_COUNT files still exceed the 200KB target."
else
    echo "All files are within the 200KB target."
fi
