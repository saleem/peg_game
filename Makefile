#Phony targets
.PHONY: go-test

# Go related variables.
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/vendor:$(GOBASE)
GOBIN=$(GOBASE)/bin

go-test:
	@echo "  >  Running unit tests"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go test -timeout 10s -run ^TestUnit.*