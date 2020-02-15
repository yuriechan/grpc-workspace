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

* define the `SquareRoot` RPC first, with unary req and resp):

```proto
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

---
