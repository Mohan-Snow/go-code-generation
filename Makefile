.PHONY: generate
generate:
	@TYPE=GenerationService OUTPUT=../generated/service.go go generate gen/gen.go

.PHONY: run
run:
	@go run *.go

.PHONY: clean
clean:
	@rm generated/service.go