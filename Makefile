
all: compile

help:
	@echo "Make targets are: compile, test, clean"

compile:
	go build

test:
	go test ./...

clean: snap-clean
	rm -rf frwd

#

SNAPTARGETS = amd64 arm64 armhf # ppc64 i686
.PHONY: snap
snap:
	make clean
	make compile
	make test
	for t in $(SNAPTARGETS) ; do set -e ; snapcraft clean ; snapcraft  --debug --target-arch $$t ; done
	snapcraft login
	for s in $$(ls *.snap) ; do set -e ; snapcraft push  --release edge,beta  $$s ; done
	snapcraft logout

snap-install:
	sudo snap install --devmode --dangerous frwd*_amd64.snap

snap-uninstall:
	sudo snap remove frwd

snap-test:
	make clean
	snapcraft clean
	snapcraft --debug

snap-clean:
	rm -rf *.snap snap stage parts prime

