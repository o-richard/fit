SHELL=/bin/bash

# Generate CSS styles based on the tailwind setup
styles:
	@tools/tailwindcss -c ./assets/dev/tailwind.config.js -i ./assets/dev/input.css -o ./assets/static/css/styles.css --minify
	@echo "Generated styles successfully"

# Generate templ components
generate:
	@templ generate
	@echo "Generated templ components"

.PHONY: styles generate
