# Session 03: [Theory] gRPC Internals Deep Dive

---

## 5. Why this section?

* learning framework is fun, but being able to understand theory is integral
* if you don't care about theory, you can skip this section

---

## 6. Protocol Buffers & Language Interoperability

* gRPC website: [link](https://grpc.io/)

### Protocol Buffers role in gRPC

* Protocol Buffers is used to define the:
  * messages (data, request and response)
  * services (service name and RPC endpoints)
* We then generate code from it

#### `example.proto`

```proto
syntax = "proto3";

message Greeting {
  string first_name = 1;
}

message GreetRequest {
  Greeting greeting = 1;
}

message GreetResponse {
  string result = 1;
}

service GreetService {
  rpc Greet(GreetRequest) returns (GreetResponse) ();
}
```

### Efficiency of Protocol Buffers over JSON

* gRPC uses Protocol Buffers for communications
* the payload sizes: Protobuf vs JSON

* JSON: 55 bytes

```json
{
  "age": 35,
  "first_name": "Mark",
  "last_name": "Hahn",
}
```

* Same in Protocol Buffers: 20 bytes

```proto
message Person {
  int32   age         = 1;
  string  first_name  = 2;
  string  last_name   = 3;
}
```

* Therefore, JSON (55 bytes) vs Protobuf (20 bytes)
  * **we save in Network Bandwidth**

* parsing JSON is actually CPU intensive
  * since the format is human-readable
* parsing Protocol Buffers (binary format) is less CPU intensive
  * since it's closer to how a machine represents data
* by using gRPC, the use of Protocol Buffers means
  * faster
  * more efficient
* in terms of communication
  * which friendly with mobile devices
    * because that relatively have slower CPU and limited hardware support

### Quick tour on `grpc.io`

* officially supported platforms
  * C/C++
  * C#
  * Dart
  * Go
  * Java
  * Node.js
  * PHP
  * Python
  * Ruby
* works across languages and platforms

### gRPC Languages

* gRPC will have these main implementations:
  * GRPC-JAVA: Pure implementation of gRPC in Java
  * GRPC-GO: Pure implementation of gRPC in Go
  * GRPC-C: Pure implementation of gRPC in C
    * gRPC C++
    * gRPC Python
    * gRPC Ruby
    * gRPC Objective-C
    * gRPC PHP
    * gRPC C#
  * other languages implement gRPC natively or rely on C implementation

### gRPC can be used by any language

* because the code can be generated for any lanugage,
  * it makes it super simple to create microservices in any language that interact each other
    * e.g.
      * an mobile app written (in Java) has a stub
      * purchase service (in Go) and it has a stub
      * pricing service (in Python)
        * the mobile app's stub (in Java) can do
          * proto req/res to
            * the purchase service (in Go)
            * the pricing service (in Python)
        * also, the purchase service (in Go), the stub can do
          * proto req/res to
            * shipment service (in C#)

### Summary: Why Protocol Buffer?

* easy to write message definition
* the definition of the API is independent from the implementation
* A huge amount of code can be generated, in any language, from a simple `.proto` file
* the payload is binary
  * therefore, very efficient to send/receive on a network and
  * serialise/de-serialiser on a CPU

---

## 7. HTTP/2

---

## 8. 4 Types of gRPC APIs

---

## 9. Scalability in gRPC

---

## 10. Security in gRPC (SSL)

---

## 11. gRPC vs REST

---

## 12. Section Summary - why use gRPC

---

## Quiz 1: Quiz on gRPC

---

