# Session 06: [Hands-On] gRPC Server Streaming

---

## 24. What's a Server Streaming API?

### 24.1 What's a Server Streaming API?

* Server Streaming RPC API are a **NEW** kind API enabled thanks to **HTTP/2**
* the client will send **one** message to the server and will receive **many** responses from the server, possibly an infinite number
* streaming server are well suited for:
  * when the server needs to send a LOT of data (big data)
  * when the server needs to "PUSH" data to the client without having the client request for more (e.g. live feed, chat, etc.)

```note
************                               ************
*          * -------> req 0 -------------> *          *
*  client  *                               *  server  *
*          * <-- resp 0, 1, 2, 3, ... <--- *          *
************                               ************
```

* in gRPC Server Streaming Calls are defined using the keyword "`stream`"
* as for each RPC call, we have to define a "Request" message and a "Response" messsage.

```proto
message GreetManyTimesRequest {
    Greeting greeting = 1;
}

message GreetManyTimesResponse {
    string result = 1;
}

service GreetService {
    // Streaming Server
    rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesRequest) {};
}
```

---

## 25. `GreetManyTimes` API Definition

## 25.1. Hands On: `GreetManyTimes` API Definition

* <u>Hands On</u>: Let's define a Streaming Server "`GreetManyTimes`" API
* It will take **ONE** `GreetManyTimesRequest` that contains a `Greeting`
* It will return **MANY** `GreetManyTimesResponse` that contains a result string

## 25.2. `greet.proto`

```proto
syntax = "proto3";

package greet;
option go_package="greetpb";

message Greeting {
    string first_name = 1;
    string last_name = 2;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse {
    string result = 1;
}

message GreetManyTimesRequest {
    Greeting greeting = 1;
}

message GreetManyTimesResponse {
    string result = 1;
}

service GreetService {
    // Unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

    // Server Streaming
    rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesResponse) {};
}
```

actually in the `rpc GreetManyTimes()GreetManyTimesRequest) returns (stream GreetManyTimesResponse)`, we can reuse `GreetRequest` and `GreetResponse` rather than `GreetManyTimesRequest` and `GreetManyTimesResponse`, respectively, HOWEVER, usually in rpc, when you define a new rpc, you should create new request and response message types.

Next step: in your `generate.sh` file, get this command and run it:

```sh
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

---

## 26. Server Streaming API Server Implementation 

---

## 27. Server Streaming API Client Implementation 

---

## 28. [Exercise] `PrimeNumberDecomposition` API

---

## 29. [Solution] `PrimeNumberDecomposition` API

---
 
