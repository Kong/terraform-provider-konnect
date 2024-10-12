FILES := generated.yaml generated-with-plugins.yaml
FORMAT := npx openapi-format --sortFile .openapi-format-sort.json

.PHONY: lint-fix
lint-fix:
	@for f in $(FILES); do \
		$(FORMAT) -o ./src/konnect/definitions/control-planes-config/v2/$$f ./src/konnect/definitions/control-planes-config/v2/$$f; \
	done

.PHONY: lint-check
lint-check:
	@for f in $(FILES); do \
		$(FORMAT) -o /tmp/formatted_$$f ./src/konnect/definitions/control-planes-config/v2/$$f; \
		if ! diff /tmp/formatted_$$f ./src/konnect/definitions/control-planes-config/v2/$$f >/dev/null; then \
			echo "The file $$f is not formatted correctly. Try running make lint-fix"; \
			exit 1; \
		fi; \
	done
