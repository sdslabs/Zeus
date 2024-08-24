.PHONY: all run clean

all:
	make clean
	make build

build:
	mkdir build
	cd build && \
	go build -o zeus ../cmd/main.go 
	sudo make arch

arch:
	sudo ./build/zeus ./filesystems/arch

arch_test:
	sudo ./build/zeus ./filesystems/arch 100M 1000

fish:
	sudo ./build/zeus ./filesystems/fish

ubuntu:
	sudo ./build/zeus ./filesystems/ubuntu

clean:
	sudo rm -rf build