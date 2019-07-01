# FileHash
A simple CLI tool for performing hash functions on files. Supports MD5, SHA1, SHA256 & SHA512 algorithms for hashing given files. Output is in the format 'filename: hashstring'. 

## Installation
1. Download the source for by running "go get -u github.com/MichaelWittgreffe/FileHash"
2. Navigate into "$GOPATH/src/github.com/MichaelWittgreffe/FileHash'
3. Run 'make' - this will perform tests and compile the source into an executable under ./bin
4. Optional: Add the executable to your $PATH to allow use in other directories, this tool is based off of the current working directory

## Usage
Usage is 'filehash [filename] [algorithm]'. Default is MD5.