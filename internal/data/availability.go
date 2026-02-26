package data

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

// DateSlot represents a single date range with availability status.
type DateSlot struct {
	Dates  string `yaml:"dates"`
	Status string `yaml:"status"`
	Note   string `yaml:"note,omitempty"`
}

// AvailabilityFile is the top-level YAML structure.
type AvailabilityFile struct {
	Trips map[string][]DateSlot `yaml:"trips"`
}

// AvailabilityStore loads and caches trip availability from a YAML file.
type AvailabilityStore struct {
	path    string
	devMode bool

	mu      sync.RWMutex
	trips   map[string][]DateSlot
	modTime time.Time
}

// NewAvailabilityStore creates a store that reads availability from the given YAML path.
// In dev mode, the file is re-checked on every Get call. A missing or malformed file
// logs a warning but does not crash the server.
func NewAvailabilityStore(path string, devMode bool) *AvailabilityStore {
	s := &AvailabilityStore{
		path:    path,
		devMode: devMode,
		trips:   make(map[string][]DateSlot),
	}
	s.load()
	return s
}

// Get returns availability slots for a trip slug. In dev mode it stat-checks
// the file and reloads if modified.
func (s *AvailabilityStore) Get(slug string) []DateSlot {
	if s.devMode {
		s.reloadIfChanged()
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.trips[slug]
}

func (s *AvailabilityStore) load() {
	raw, err := os.ReadFile(s.path)
	if err != nil {
		log.Printf("[availability] warning: cannot read %s: %v", s.path, err)
		return
	}

	var af AvailabilityFile
	if err := yaml.Unmarshal(raw, &af); err != nil {
		log.Printf("[availability] warning: cannot parse %s: %v", s.path, err)
		return
	}

	// Validate statuses
	for slug, slots := range af.Trips {
		for i, slot := range slots {
			switch slot.Status {
			case "open", "limited", "booked":
				// valid
			default:
				log.Printf("[availability] warning: unknown status %q for %s[%d], treating as open", slot.Status, slug, i)
				af.Trips[slug][i].Status = "open"
			}
		}
	}

	info, err := os.Stat(s.path)
	if err != nil {
		log.Printf("[availability] warning: cannot stat %s: %v", s.path, err)
		return
	}

	s.mu.Lock()
	s.trips = af.Trips
	s.modTime = info.ModTime()
	s.mu.Unlock()

	log.Printf("[availability] loaded %d trips from %s", len(af.Trips), s.path)
}

// GetAll returns a copy of the full trips map for the admin editor.
func (s *AvailabilityStore) GetAll() map[string][]DateSlot {
	if s.devMode {
		s.reloadIfChanged()
	}
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make(map[string][]DateSlot, len(s.trips))
	for k, v := range s.trips {
		slots := make([]DateSlot, len(v))
		copy(slots, v)
		out[k] = slots
	}
	return out
}

// Save validates the given trips, writes them atomically to the YAML file, and updates the in-memory cache.
func (s *AvailabilityStore) Save(trips map[string][]DateSlot) error {
	// Validate statuses
	for slug, slots := range trips {
		for i, slot := range slots {
			switch slot.Status {
			case "open", "limited", "booked":
				// valid
			default:
				return fmt.Errorf("invalid status %q for %s[%d]", slot.Status, slug, i)
			}
		}
	}

	af := AvailabilityFile{Trips: trips}
	out, err := yaml.Marshal(&af)
	if err != nil {
		return fmt.Errorf("marshal yaml: %w", err)
	}

	header := "# Trip Availability — MT Hunt & Fish Outfitters\n" +
		"# Edit this file to update availability on the website.\n" +
		"# Status options: open, limited, booked\n\n"

	// Atomic write: temp file in same dir + rename
	dir := filepath.Dir(s.path)
	tmp, err := os.CreateTemp(dir, "availability-*.yaml")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpName := tmp.Name()

	if _, err := tmp.WriteString(header); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("write header: %w", err)
	}
	if _, err := tmp.Write(out); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("write yaml: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("close temp file: %w", err)
	}
	if err := os.Rename(tmpName, s.path); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("rename temp file: %w", err)
	}

	// Update in-memory cache
	info, _ := os.Stat(s.path)
	s.mu.Lock()
	s.trips = trips
	if info != nil {
		s.modTime = info.ModTime()
	}
	s.mu.Unlock()

	log.Printf("[availability] saved %d trips to %s", len(trips), s.path)
	return nil
}

func (s *AvailabilityStore) reloadIfChanged() {
	info, err := os.Stat(s.path)
	if err != nil {
		return
	}
	s.mu.RLock()
	changed := info.ModTime().After(s.modTime)
	s.mu.RUnlock()
	if changed {
		s.load()
	}
}
