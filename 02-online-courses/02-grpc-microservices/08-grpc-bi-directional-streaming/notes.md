# Session 08: [Hands-On] gRPC Bi-Directional Streaming

---

## 36. What's a Bi-Directional Streaming API?

### 36.1. What is a Bi Directional Streaming API?

* Bi Directional Streaming RPC API are a **NEW** kind API enabled thanks to **HTTP/2**
* The client will send **many** messages to the server and will receive **many** responses from the server
* The number of requests and responses **does not have to match**
* Bi Directional Streaming RPC are well suited for
  * When the client and the server needs to send a lot of data asynchronously
  * "Chat" protocol
  * Long running connections

```note
              Bi-Directional Streaming API
************                               ************
*          * --> req ..., 3, 2, 1, 0 ----> *          *
*  client  *                               *  server  *
*          * <-- resp 0, 1, 2, 3, ... <--- *          *
************                               ************
```

* In gRPC, Bi Directional Streaming API are defined using the keyword "`stream`", **twice**
* As for each RPC call, we have to define a "Request" message and a "Response" message

```proto
message GreetEveryoneRequest {
  Greeting greeting = 1;
}

message GreetEveryoneResponse {
  string result = 1;
}

service GreetService {
  // Bi Directional Streaming
  rpc GreetEveryone(stream GreetEveryoneRequest) returns (stream GreetEveryoneResponse) {};
}
```

---

## 37. `GreetEveryone` API Definition

---

## 38. Bi-Directional Streaming API Server Implementation

---

## 39. Bi-Directional Streaming API Client Implementation

---

## 40. [Exercise] `FindMaximum` API

---

## 41. [Solution] `FindMaximum` API

---