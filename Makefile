.PHONY: generate
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