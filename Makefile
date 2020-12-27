SOURCES := $(wildcard *.go */*.go)
VERSION := $(shell cat .version)
BINARY := shit

# Development build

TARGET := target

all: $(TARGET)/$(BINARY)

$(TARGET)/$(BINARY): $(SOURCES)
	go build -ldflags "-X main.version=$(VERSION) -X main.mode=dev" -o $@

# Production build

DIST := dist

build: $(DIST)/$(BINARY)

$(DIST)/$(BINARY): $(SOURCES)
	go build -ldflags "-X main.version=$(VERSION) -X main.mode=prod" -a -o $@

# Follow logs

LOG := ~/.shit.log

log:
	@touch $(LOG)
	@tail -f $(LOG)

# Clean

clean:
	@$(RM) -r $(TARGET) $(DIST) $(LOG)

.PHONY: all build log clean
