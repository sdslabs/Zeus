#!/bin/bash

# Check if Go is installed
if ! command -v go &>/dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Run go mod tidy
go mod tidy

cd testing
# Compile multiple C files separately with their file names
for file in *.c; do
    if [ -f "$file" ]; then
        filename="${file%.*}"
        gcc -o "$filename" "$file"
    fi
done
cd ..

# Run make build
make build
