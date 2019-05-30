
all: compile

help:
	@echo "Make targets are: compile, test, clean"

compile:
	go build

test:
	go test ./...

clean:
	rm -rf frwd
