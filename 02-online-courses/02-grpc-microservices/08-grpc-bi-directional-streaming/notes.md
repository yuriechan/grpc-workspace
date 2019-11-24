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

* Hands-on: Let's define a Bi Directional Streaming "`GreetEveryone`" API
* It will take **MANY** `GreetEveryoneRequest` that contains a `Greeting`
* It will return **MANY** `GreetEveryoneResponse` that contains a result string

starting with the `greet/greetpb/greet.proto` file, let's add additional code:

```proto
...

message GreetEveryoneRequest {
  string greeting = 1;
}

message GreetEveryoneResponse {
  string result = 1;
}

service GreetService{
  ...

  // BiDi Streaming
  rpc GreetEveryone(stream GreetEveryoneRequest) returns (stream GreetEveryoneResponse) {};
}

then, make sure to generate code again without any issue:

```bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

then we get the newly generated code: `greet/greetpb/greet.pb.go`

---

## 38. Bi-Directional Streaming API Server Implementation

* Hands-on:
* We'll implement a Bi Directional Streaming `GreetEveryone` RPC
* **NOTE**: although we will respond to every message in this example, it is not necessary to do so, and the server is free to choose how many responses to send for each client message

let's have a look: `greet/greet_server/server.go`

```go
func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
  fmt.Printf("GreetEveryone function was invoked with a streaming request\n")
  
  for {
    req, err := stream.Recv()
    if err == io.EOF {
      return nil
    }
    if err != nil {
      log.Fatalf("Error whilst reading client stream: %v", err)
      return err
    }
    firstName := req.GetGreeting().GetFirstName()
    result := "Hello " + firstName + "! "

    sendErr := stream.Send(&greetpb.GreetEveryoneResponse{
      Result: result,
    })
    if sendErr != nil {
      log.Fatalf("Error whilst sending data to client: %v", err)
      return err
    }
  }
}
```

and let's see this server compiles and run well:

```bash
$ go run greet/greet_server/server.go
Hello world!
```

now, it's ready to use.

---

## 39. Bi-Directional Streaming API Client Implementation

---

## 40. [Exercise] `FindMaximum` API

---

## 41. [Solution] `FindMaximum` API

---