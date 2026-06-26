MAIN_PATH = internal/http/swagger.go
DOCS_PATH = docs

# ==============================================================================
# SWAGGER COMMANDS
# ==============================================================================

.PHONY: swag
swag:
	@echo "==> Generating Swagger Documentation..."
	swag init -g $(MAIN_PATH) --output $(DOCS_PATH) --parseDependency
	@echo "==> Swagger Documentation generated successfully at /$(DOCS_PATH)"