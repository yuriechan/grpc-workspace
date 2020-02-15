# Session 09: [Hands-On] gRPC Advanced Feature Deep Dive

---

## 43. [Theory] Errors in gRPC

### 43.1. Error Codes

#### 43.1.1. Conventional HTTP Errors

* It is common for your API to something return error codes
* in HTTP, there are **many** error codes
  * 2xx for success
  * 3xx for ...
  * 4xx for ...
  * 5xx for ...
* Whilst HTTP codes are standardisied, they are not usually clear enough

#### 43.1.2. gRPC Errors

* with gRPC, there are a few error codes:
  * https://grpc.io/docs/guides/error.html
* there is also complete reference to implementation of error codes that close a lot of gaps with the documentation:
  * http://avi.im/grpc-errors
* if an application needs to return extra information on the top of an error code, it can use the metadata context

**General Errors:**

| **Case** | **Status Code** |
| --- | --- |
| Client application cancelled the request | `GRPC_STATUS_CANCELLED` |
| Deadline expired before server returned status | `GRPC_STATUS_DEADLINE_EXCEEDED` |
| Method not found on server | `GRPC_STATUS_UNIMPLEMENTED` |
| Server shutting down | `GRPC_STATUS_UNAVAILABLE` |
| Server threw an exception (or did something other than returning a status code to terminate the RPC) | `GRPC_STATUS_UNKNOWN` |

**Network Failures:**

| **Case** | **Status Code** |
| --- | --- |
| No data transmitted before deadline expires. Also applies to cases where some data is transmitted and no other failures are detected before the deadline expires | `GRPC_STATUS_DEADLINE_EXCEEDED`
| Some data transmitted (for example, the request metadata has been written to the TCP connection) before the connection breaks | `GRPC_STATUS_UNAVAILABLE`

**Protocol Errors:**

| **Case** | **Status Code** |
| --- | --- |
| Could not decompress but compression algorithm supported | `GRPC_STATUS_INTERNAL`
| Compression mechanism used by client not supported by the server | `GRPC_STATUS_UNIMPLEMENTED`
| Flow-control resource limits reached | `GRPC_STATUS_RESOURCE_EXHAUSTED`
| Flow-control protocol violation | `GRPC_STATUS_INTERNAL`
| Error parsing returned status | `GRPC_STATUS_UNKNOWN`
| Unauthenticated: credentials failed to get metadata | `GRPC_STATUS_UNAUTHENTICATED`
| Invalid host set in authority metadata | `GRPC_STATUS_UNAUTHENTICATED`
| Error parsing response protocol buffer | `GRPC_STATUS_INTERNAL`
| Error parsing request protocol buffer | `GRPC_STATUS_INTERNAL`

---

## 44. [Hands-On] Errors implementation

### 44.1. Error Codes: Hands-On

* let's implement an error message for a new `SquareRoot` Unary API
* we'll create `SquareRoot` RPC
* we'll implement `Server` with the error handling
* we'll implement `Client` with the error handling

### 44.2. the Implementation

#### 44.2.1. Protobuf

* define the `SquareRoot` RPC first, with unary req and resp):

```proto
...

message SquareRootRequest {
  int32 number = 1;
}

message SquareRootResponse {
  double number_root = 1;
}

service CalculatorService {
  ... // previous RPC definitions

  rpc SquareRoot(SquareRootRequest) returns (SquareRootResponse){};
}
```

* generate `caculator.pb` by using:

```bash
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```

* One of important things we should do: "documentation for error handling"
  * `// make a comment about error handling for each rpc`

```proto
...
service CalculatorService {
  ... // previous RPC definitions

  // error handling
  // this RPC will throw an exception if the sent number is negative
  // the error being sent is of type `INVALID_ARGUMENT`
  rpc SquareRoot(SquareRootRequest) returns (SquareRootResponse){};
}
```

#### 44.2.2. Server

now, we can go to the `calculator_server/server.go` and have to implement `SquareRoot`

```go
...

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
  fmt.Println("Received SquareRoot RPC")
  number := req.GetNumber()
  if number < 0 {
    return nil, status.Errorf(
      codes.InvalidArgument,
      fmt.Sprintf("Received a negative number: %v", number),
    )
  }
  return &calculatorpb.SquareRootResponse{
    NumberRoot: math.Sqrt(float64(number)),
  }, nil
}

...
```

#### 44.2.3. Client

The server is now supporting the `SquareRoot`, we need to update the client, `calculator_client/client.go`

```go
...

func main() {
  fmt.Println("Calculator Client")
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("could not connect: %v", err)
  }
  defer cc.Close()

  c := calculatorpb.NewCalculatorServiceClient(cc)

  ...

  doErrUnary(c)
}

...

func doErrUnary(c calculatorpb.CalculatorServiceClient) {
  fmt.Println("Starting to do a SquareRoot Unary RPC...")

  // correct call
  doErrorCall(c, 10)

  // error call
  doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
  resp, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
  if err != nil {
    // err converted into the respErr
    // which is gRPC-friendly error that has message and code
    // but only if this `err` is actual gRPC error
    // if NOT, you'll get `ok` being false,
    // you can throw the normal error (the else clause)
    respErr, ok := status.FromError(err)
    if ok {
      // actual error from gRPC (user error)
      fmt.Printf("Error message from server: %v\n", respErr.Message())
      fmt.Println(respErr.Code())
      if respErr.Code() == codes.InvalidArgument {
        fmt.Println("We probably sent a negative number!")
        return
      }
    } else {
      // bigger error, framework type of error
      log.Fatalf("Big Error calling SquareRoot: %v", err)
      return
    }
  }
  fmt.Printf("Result of square root of %v: %v\n", n, resp.GetNumberRoot())
}

...
```

* for the `codes.*` of `codes. we can also use another errors:
  * `codes.NotFound`
  * `codes.OutOfRange`
  * `codes.PermissionDenied`
  * and more ...

#### 44.2.4. Run Server and Client

Now, the server and client are written. Let's run the server:

```bash
go run calculator/calculator_server/server.go
```

and you will see this message:

```bash
Calculator Server
```

then, let's run the client:

```bash
go run calculator/calculator_client/client.go
```

by running the client, the client can send requests to the server so the server-side get this messages:

```bash
Received SquareRoot RPC
Received SquareRoot RPC
```

because client send request 2 times: one is valid request, the other is invalid request

and the client-side get the responses:

```bash
Calculator Client
Starting to do a SquareRoot Unary RPC...
Result of square root of 10: 3.1622776601683795
Error message from server: Received a negative number: -2
InvalidArgument
We probably sent a negative number!
```

---
