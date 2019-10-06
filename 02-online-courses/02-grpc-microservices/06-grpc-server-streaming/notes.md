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

---

## 26. Server Streaming API Server Implementation 

---

## 27. Server Streaming API Client Implementation 

---

## 28. [Exercise] `PrimeNumberDecomposition` API

---

## 29. [Solution] `PrimeNumberDecomposition` API

---
 
