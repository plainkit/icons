.PHONY: update-icons generate

update-icons:
	@echo "Downloading latest Lucide icons..."
	@mkdir -p tmp
	curl -s -L -o tmp/lucide.zip https://github.com/lucide-icons/lucide/archive/refs/heads/main.zip
	@echo "Extracting archive..."
	cd tmp && unzip -q lucide.zip
	@echo "Copying SVG files..."
	mkdir -p lucide/icons
	cp tmp/lucide-main/icons/*.svg lucide/icons/
	@echo "Cleaning up temporary files..."
	rm -rf tmp
	@echo "Icons updated successfully!"

generate:
	@echo "Generating icons..."
	go run lucide/cmd/gen-icons/main.go -icons=lucide/icons -out=lucide -overwrite=true
	@echo "Generating tests..."
	go run lucide/cmd/gen-tests/main.go -dir=lucide -overwrite=true
	@echo "Running goimports..."
	goimports -w .
	@echo "Running gofmt..."
	gofmt -w .
	@echo "Done!"
