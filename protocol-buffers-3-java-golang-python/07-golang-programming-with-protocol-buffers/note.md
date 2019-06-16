# Section 7. Golang Programming with Protocol Buffers

---

## 33. Setup & Code Download in Golang

### TODO: Setup your Mac

- Golang installed: https://golang.org/doc/install
- VSCode installed: https://code.visualstudio.com/
- The VSCode Golang extension: https://code.visualstudio.com/docs/languages/go
- The source code for this project (download / star the project): https://github.com/simplesteph/protobuf-example-go

- Golang packages:

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
```

- Protoc Compiler (see Setting up Protoc Compiler Section)

---

## 34. Code generation in Golang

### Hello World in golang

Check out your Mac has golang compliler

#### `main.go`

```go
package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello world!")
}
```

then run the code on your terminal:

```bash
go run main.go
```

If you can see and output: `Hello world!`, you are ready to go!

###
