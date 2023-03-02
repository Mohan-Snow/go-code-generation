.PHONY: generate
generate:
	@TYPE=EncodingService OUTPUT=../generated/service.go go generate gen/gen.go

.PHONY: run
run:
	@go run *.go

.PHONY: clean
clean:
	@rm generated/service.go