TOOLS_BIN=tools/bin
MOCKGEN ?= $(TOOLS_BIN)/mockgen

$(MOCKGEN): tools/go.mod
	cd tools && go build -o bin/mockgen github.com/golang/mock/mockgen
