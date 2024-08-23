.PHONY: all run clean

all:
	make clean
	make build

build:
	mkdir build
	cd build && \
	go build -o zeus ../cmd/main.go 
	make arch

arch:
	./build/zeus ./filesystems/arch

fish:
	./build/zeus ./filesystems/fish

clean:
	rm -rf build