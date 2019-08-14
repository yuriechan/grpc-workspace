# Session 04: [Hands-On] gRPC Project Overview & Setup 

---

## 13. Project Setup (Golang + VSCode + Protoc)

### 13.1. Go Setup

* Instructions to install Golang: https://golang.org/doc/install
* Make sure that the go binaries are in your `PATH`
  * (see installation instructions), so that you can run the following command in a terminal:

```bash
go version
go version go1.12.6 darwin/amd64
```

### 13.2. VSCode Setup

* instal VSCode: https://code.visualstudio.com/
* instal VSCode extensions [link](https://code.visualstudio.com/docs/editor/extension-gallery#_browse-and-install-extensions)
  * `vscode-proto3`
  * `Clang-Format`
    * MacOSX: `brew install clang-format`
* instal `protoc` (see below):

### 13.3. Protoc Setup

* in order to perform code generation, you will need to instal `protoc` on your computer
  * open a command line interface and type `brew install protobuf`

---

## 14. Go Dependencies Setup

* visit `gRPC-GO` github: [link](https://github.com/grpc/grpc-go)
* and follow the installation instructions:
  * it takes some time

```bash
go get -u google.golang.org/grpc
```

* visit `golang/protobuf` github: [link](https://github.com/golang/protobuf)
* run this command:

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

---

## 15. Code Generation Test

---

## 16. Server Setup Boilerplate Code

---

## 17. Client Setup Boilerplate Code

---
