# gRPC workspace

This is a personal note for studying gRPC (Golang) + modern API & Microservices  

## Table of Contents

* Protocol Buffer: [link](/protocol-buffers-3-java-golang-python/TOC.md)
* gRPC - modern API & Microservices: [link](/grpc-modern-api-microservices/TOC.md)

## Short Introduction

* [Introduction to gRPC by Microsoft VisualStudio](https://www.youtube.com/watch?v=ru5x_hDZ9Qw)

## How to Instal

### 1. check out your go version:

```bash
go version
```

this is because **gRPC** can be used `go` version `1.6` and above

### 2. instal gRPC

```bash
go get -u -v google.golang.org/grpc
```

### 3. Instal Protobuf version 3

possible to get file [HERE](https://github.com/protocolbuffers/protobuf/releases) and add `PATH`

in case of Mac, use `brew`

```bash
brew install protobuf
```

also, instal protoc plugin for golang

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

and add path

```bash
export PATH=$PATH:$GOPATH/bin
```
