SRCS := $(shell find . -type f -name '*.go')

# Alpineで動かすのでStatic化
LDFLAGS := -ldflags="-s -w -extldflags \"-static\""
TARGETS := $(shell find ./cmd/ -type d | sed 's!^.*/!!')

all: $(TARGETS)
$(TARGETS):
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$@ ./cmd/$@
# bin/$(NAME): $(SRCS)
# 	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME)