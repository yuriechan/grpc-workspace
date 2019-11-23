# Session 07: [Hands-On] gRPC Client Streaming

---

## 30. What's a Client Streaming API?

### 30.1. What's a Client Streaming API?

* Client Streaming RPC API are a **NEW** kind API enabled thanks to **HTTP/2**
* The client will send **many** message to the server and will receive **one** response from the server (at any time)
* Streaming Client are well suited for
  * when the client needs to send a lot of data (big data)
  * when the server processing is expensive and should happen as the client sends data
  * when the client needs to "PUSH" data to the server without really expecting a response

```note
************                               ************
*          * --> req 0, 1, 2, 3, ... ----> *          *
*  client  *                               *  server  *
*          * <-- resp 0 <----------------- *          *
************                               ************
```

* In gRPC Client Streaming Calls are defined using the keyword "`stream`"
* As for each RPC call we have to define a "Request" message and a "Response" message

```proto
message LongGreetRequest {
  Greeting greeting = 1;
}

message LongGreetResponse {
  string result = 1;
}

service GreetService {
  // Streaming Client
  rpc LongGreet(stream LongGreetResponse) returns (LongGreetResponse {}:
}
```

---

## 31. `LongGreet` API Definition

## 31.1. Hands On: `LongGreet` API Definition

* Hands On: Let's define a Streaming Client "`LongGreet`" API
* It will take **MANY** `LongGreetRequest` that contains a `Greeting`
* It will return **ONE** `LongGreetResponse` that contains a result string

```proto
...

message LongGreetRequest {
  Greeting greeting = 1;
}

message LongGreetResponse {
  string retult = 1;
}

service GreetService {
  
  ...

  // Client Streaming
  rpc LongGreet(stream LongGreetRequest) returns LongGreetResponse) {};
}
```

then run the command to generate new code:

```bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

and you can see the generated code with new functions for `LongGreet`.

---

## 32. Client Streaming API Server Implementation

---

## 33. Client Streaming API Client Implementation

---

## 34. [Exercise] `ComputeAverage` API

---

## 35. [Solution] `ComputeAverage` API

---
