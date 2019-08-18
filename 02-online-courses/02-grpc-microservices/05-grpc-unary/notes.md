# Session 05: [Hands-On] gRPC Project Overview & Setup

---

## 18. What's an Unary API

* Unary RPC calls are basic Request / Response that everyone is familar with
* The client will send **one** message to the server and will receive **one** response from the server
* Unary RPC calls will be the most common for your APIs
  * Unary calls are very well suited when your data is small
  * Start with Unary when writing APIs and use streaming API if performance is an issue

```note
              Unary API
------------             ------------
|          | ----------> |          |
|  Client  |             |  Server  |
|          | <---------- |          |
------------             ------------
```

* In gRPC, Unary Calls are defined using Protocol Buffers
* For each RPC call, we have to define a "Request" message and a "Response" message

```proto
message GreetRequest {
  Greeting greeting = 1;
}

message GreetResponse {
  string result = 1;
}

service GreetService {
  // Unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

---

## 19. Greet API Definition

### 19.1 Hands On - Greet API Definition

* **Hands On**: define a Unary "Greet" API
* our message is: `Greeting`
  * contains: `first_name`, `last_name` string field
  * will take: a `GreetRequest` that contains a `Greeting`
  * will return: a `GreetResponse` that contains a `result` string

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

service GreetService {
  // Unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

---

## 20. Unary API Server Implementation

---

## 21. Unary API Client Implementation

---

## 22. [Exercise] Sum API

---

## 23. [Solution] Sum API

---

