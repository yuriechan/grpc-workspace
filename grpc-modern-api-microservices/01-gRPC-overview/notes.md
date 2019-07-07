# Session 01 gRPC Course Overview

---

## 1. gRPC Introduction

* Today's trend is to build microservices (MS)
  * MS are
    * built in different languages
    * encompass a function of your business
  * These MS must exchange information and need to agree on:
    * the API to exchange data
    * the data format
    * the error patterns
    * load balancing
    * many other
  * one of popular choice for building API
    * REST (HTTP-JSON)

* Building an API is hard
  * need to think about data model
    * JSON
    * XML
    * binary?
  * need to think about the endpoints
    * `GET` `/api/v1/user/123/post/456`
    * `POST` `/api/v1/user/123/post`
  * need to think about how to invoke it and handle errors
    * API
    * errors
  * need to think about efficiency of the API
    * how much data do I get out of one call?
    * too much data?
    * too little data? -> many API calls?
  * how about:
    * latency?
    * scalability to 1000s of clients?
    * load balancing?
    * inter operability with many different languages?
    * authentication
    * monitoring
    * logging

> "Don't you wish you could leave the boring and hard stuff to the framework?"

* What's an API?
  * at its core, an API is a contract, saying:
    * send me this `REQUEST` (client)
    * will send you this `RESPONSE` (server)
    * **It's all about data!**
    * the rest, we'll leave to the gRPC framework

* What's gRPC?
  * free and open-source framework developed by Google
  * part of the CNCF(Cloud Native Computation Foundation) like Docker & Kubernentes
  * at a high level:
    * it allows you to define `REQUEST` and `RESPONSE` for `RPC` (Remote Procedure Calls) and handles all the rest for you
  * on top of it:
    * it's modern, fast, and efficient.
    * build on top of `HTTP/2`
    * low latency
    * supports streaming
    * language independent
    * makes it super easy to plub in:
      * authentication
      * load balancing
      * logging
      * monitoring

* What's an RPC?
  * Remote Procedure Call
  * in your `CLIENT` code, it looks like you're just calling a function directly on the `SERVER`

Client code (any other language):

```py
(code)
...
server.CreateUser(user)
...
(code)
```

Server code (any other language):

```py
// function creating users
def CreateUser(User user) {
  ...
}
```

In the examples above, the `CLIENT` do `RPC call` to the `SERVER` over the network.

* What's an RPC?
  * it's not a new concept
    * (CORBA had this before)
  * with gRPC:
    * implemented very cleanly
    * solves a lot of problems
  * e.g.
    * A: C++ service that has gRPC server inside
    * B: Ruby client that has gRPC stub
    * C: Android-Java client that has gRPC stub
    * In this situation, gRPC stub at Ruby client can do **proto request** to the gRPC server inside of C++ service and the gRPC server can return **proto response** to B
    * In addition, C can also do the same thing to A, just like what B does to A (vice versa)

* How to get started?
  * core of gRPC:
    * need to define:
      * `message`s
      * `service`s
      * using **Protocol Buffers** (`*.proto` files)
  * the rest of gRPC code:
    * auto generated
    * we just need to provide implementation for it
  * one `*.proto` file works for over 12 programming languages (server & client)
  * allows you to use a framework that scales to millions of RPC per seconds

* an Example: **The GreetService**

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
  rpc Greet(GreetRequest) returns (GreetResponse) {};
}
```

* Why Protocol Buffers?
  * protocol buffers are language agnostic
  * code can be generated for many languages
  * data is binary
    * efficiently serialised (small payload)
  * convenient for transporting a lot of data
  * protobufs allow for easy API evolution using their rules

* Why should we learn?
  * many companies have embraced it fully in production
    * Google (internally for Google Cloud services)
      * e.g. Pub/Sub
    * Netflix
    * Square (first contributor, replacement of all their APIs)
    * CoreOS (etcd 3 is built on gRPC for server-server communication)
    * CoackroachDB
    * Mercari
  * gRPC is the future of micro-service API and mobile-server API (and maybe Web APIs)

## 2. Course Objective

* Part 1: theory
  * gRPC concept
* Part 2: Hands-On Programming
  * gRPC implementation of all kind of APIs
  * exercises
* Part 3: Hands-On Advanced
  * gRPC advanced concepts
  * implementations

* Course Objectives
  * learn the gRPC theory to understand how gRPC works
  * compare gRPC and REST API paradigm
  * write your gRPC service definition in `.proto` files
  * generate server & client codes
  * implement
    * unary
    * server streaming
    * client streaming
    * bi-directional streaming API
  * practice your learning with exercises & solutions
  * implement advanced concepts such as
    * error handling
    * deadlines
    * SSL security
  * get pointeres to expand your learning journey and get inspired by real world gRPC services

* Pre-requisties
  * good understanding of the programming language
  * asynchronous programming is a plus
  * good understanding of protocol buffers
  * lots of willingness to learn something new
  * this course will be challenging

* Who is this course for?
  * **developers**: who want to understand how to write gRPC services and clients
  * **architects**: who want ot understand how gRPC works and the concepts behind the different types of API

## 3. Important Message

* programming parts will be after the theory

## 4. Code Download

* github repo: [link](https://github.com/simplesteph/grpc-go-course)

---
