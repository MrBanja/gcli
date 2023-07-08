#!/bin/bash

# Define program name
PROG='gcli'

# Compile for Darwin
env GOOS=darwin go build -o ${PROG}-darwin-amd64

# Compile for Linux
env GOOS=linux go build -o ${PROG}-linux-amd64

# Compile for Windows
env GOOS=windows go build -o ${PROG}-windows-amd64.exe

echo "Compilation finished. The executables are named ${PROG}-darwin-amd64, ${PROG}-linux-amd64, and ${PROG}-windows-amd64.exe"
