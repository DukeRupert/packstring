.PHONY: dev build css-watch css-build setup images deploy lighthouse

TAILWIND := ./tailwindcss
CSS_INPUT := static/css/input.css
CSS_OUTPUT := static/css/output.css

dev:
	@echo "Starting dev server..."
	@$(MAKE) -j2 css-watch air

air:
	PACKSTRING_DEV=1 air

css-watch:
	$(TAILWIND) -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch

css-build:
	$(TAILWIND) -i $(CSS_INPUT) -o $(CSS_OUTPUT) --minify

build: css-build
	go build -o bin/packstring ./cmd/server

setup:
	@if [ ! -f $(TAILWIND) ]; then \
		echo "Downloading Tailwind CSS standalone CLI..."; \
		curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64; \
		mv tailwindcss-linux-x64 $(TAILWIND); \
		chmod +x $(TAILWIND); \
		echo "Tailwind CLI installed."; \
	else \
		echo "Tailwind CLI already present."; \
	fi

images:
	bash scripts/optimize-images.sh

deploy:
	docker build -t packstring:latest .
	@echo "Image built: packstring:latest"
	@echo "Run with: docker run -p 8080:8080 packstring:latest"

lighthouse:
	npx lighthouse http://localhost:8080 --output html --output-path ./lighthouse-report.html --chrome-flags="--headless --no-sandbox"
