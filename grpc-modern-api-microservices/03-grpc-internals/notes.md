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

### What's HTTP/2

* gRPC leverages HTTP/2 as a backbone for communications
* demo: [link](https://imagekit.io/demo/http2-vs-http1)
* HTTP/2 is the newer standard for the Internet communication that address common pitfall of HTTP/1.1 on modern web pages
* before we go into HTTP/2
  * let's look at some HTTP/1.1 requests

### How HTTP/1.1 works

* released in 1997, worked great for many years
* HTTP/1.1 opens a new TCP connection to a server at each request
* it does NOT compress headers (which are plaintext)
* it only works with Request/Response mechanism (no server push)
* HTTP was orignally composed of two commands:
  * `GET`: to ask for content
  * `POST`: to send content
* Nowadays, a web page loads 80 assets on average
* Headers are sent at every request and are PLAINTEXT (heavy size)
* each request opens a TCP connection
* these inefficiencies add latency and increase network packet size

```example
Client  ->  request 1   -> HTTP 1.1 Server
Client  <-  response 1  <- HTTP 1.1 Server
Client  ->  request 2   -> HTTP 1.1 Server
Client  <-  response 2  <- HTTP 1.1 Server
```

### How HTTP/2 works

* HTTP 2 was released in 2015
* it has been battled tested for many years
  * and was tested by Google under the name SPDY
* HTTP 2 supports
  * multiplexing
    * the client &  server can push messages in parallel over the same TCP
      * this greatly reduce latency
  * server push
    * server can push streams (multiple messages) for one request from the client
  * header compression
    * headeres (text based) can now be compressed
    * these have much less impact on the packet size
    * (remember the average http request may have over 20 headers, due to cookies, content cache, and application headers)
* HTTP/2 is binary
  * while HTTP/1 text makes it easy for debugging, it's not efficient over the network
  * (Protocol Buffers is a binary protocol and make it a great match for HTTP2)
  * HTTP/2 is secure (SSL is not required but recommended by default)

```example
Client  ->  Single TCP connection -> HTTP 2 Server
Client  ->  Request 1             -> HTTP 2 Server
Client  <-  Response 1            <- HTTP 2 Server
Client  <-  Response 2            <- HTTP 2 Server
Client  <-  Response 3            <- HTTP 2 Server
Client  <-  Response ...          <- HTTP 2 Server
```

### HTTP/2: Bottom Line

* less chatter
* more efficient protocol (less bandwidth)
* reduce latency
* increased security
* **and you get all these improvements out of the box by using the gRPC framework!**

---

## 8. 4 Types of gRPC APIs

* the 4 types
  * Unary
  * Server Streaming
  * Client Streaming
  * Bi-Directional Streaming
* Unary is what a traditional API looks like (HTTP REST)
* HTTP/2 as we've seen, enables APIs to now have streaming capabilities
* the server and the client can push multiple messages as part of one request
* in gRPC, it's very easy to define these APIs as we'll see

```proto
service GreetService {
  // Unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};

  // Streaming Server
  rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesResponse) {};

  // Streaming Client
  rpc LongGreet(stream LongGreetRequest) returns (LongGreetResponse) {};

  // Bi Directional Streaming
  rpc GreetEveryone(stream GreetEveryoneRequest) returns (stream GreetEveryoneResponse) {};
}
```

---

## 9. Scalability in gRPC

* gRPC servers are asynchronous by default
  * this means they do not block threads on request
  * therefore, each gRPC server can serve millions of requests in parallel
* gRPC clients can be asynchronous or synchronous (blocking)
  * the client decides which model works best for the performance needs
  * gRPC clients can perform client side load balancing
* as a proof of scalability:
  * Google has 10 BILLION gRPC requests being made per second internally

---

## 10. Security in gRPC (SSL)

* by default, gRPC strongly advocates for you to use SSL (encryption over the wire) in your API
* this means that gRPC has security as a first class citizen
* each language will provide an API to load gRPC with the required certificates and provide encryption capability out of the box
* additionally, using **internceptors**, we can also provide authentication (will covered by advanced section later)

---

## 11. gRPC vs REST

---

## 12. Section Summary - why use gRPC

---

## Quiz 1: Quiz on gRPC

---

