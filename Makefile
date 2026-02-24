.PHONY: dev build css-watch css-build setup images deploy lighthouse

TAILWIND := ./tailwindcss
CSS_INPUT := static/css/input.css
CSS_OUTPUT := static/css/output.css

dev:
	@echo "Starting dev server..."
	@$(MAKE) -j2 css-watch air

air:
	air

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
	@echo "Image optimization not yet configured."

deploy:
	@echo "Deploy not yet configured."

lighthouse:
	@echo "Lighthouse CI not yet configured."
