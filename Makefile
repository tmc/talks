.PHONY: run
run: deps
	present -use_playground

.PHONY: deps
deps:
	command -v present || go intall golang.org/x/tools/cmd/present@latest
