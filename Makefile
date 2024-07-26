# Build the application
all: build

build:
	@echo "Building..."
	@templ generate
	@tailwindcss -i styles/styles.css -o public/output.css
	@go build -o main main.go

# Live Reload
watch:
	air

