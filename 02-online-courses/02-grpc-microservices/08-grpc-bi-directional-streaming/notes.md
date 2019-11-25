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

* Hands-on:
* We'll implement a client call for our Bi Directional Streaming RPC
* We'll test it against our server that is running!

let's work with the exisitng client code: `greet/greet_client/client.go`

```go
func doBiDiStreaming(c greetpb.GreetServiceClient) {
  fmt.Println("Starting to do a BiDi Streaming RPC...")

  // we create a stream by invoking the client
  stream, err := c.GreetEveryone(context.Background())
  if err != nil {
    log.Fatalf("Error whilst creating stream: %v", err)
  }

  requests := []*greetpb.GreetEveryoneRequest{
    &greetpb.GreetEveryoneRequest{
        Greeting: &greetpb.Greeting{
        FirstName: "Mark",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Chris",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "JD",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Stephan",
      },
    },
    &greetpb.GreetEveryoneRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Deepak",
      },
    },
  }

  waitc := make(chan struct{})
  // we send a bunch of messages to the client (go routine)
  go func() {
    // function to send a bunch of messages
    for _, req := range requests {
      fmt.Printf("Sending message: %v\n", req)
      stream.Send(req)
      time.Sleep(1000 * time.Millisecond)
    }
    stream.CloseSend()
  }()

  // we receive a bunch of messages from the client (go routine)
  go func() {
    // function to receive a bunch of messages
    for {
      res, err := stream.Recv()
      if err == io.EOF {
        break
      }
      if err != nil {
        log.Fatalf("Error whilst receiving: %v", err)
        break
      }
      fmt.Printf("Received: %v", res.GetResult())
    }
    close(waitc)
  }()

  // block until everyone is done
  <-waitc
}
```

and let's run the server:

```bash
$ go run greet/greet_server/server.go
Hello world!
```

and the client

```bash
$ go run greet/greet_client/client.go
Hello, I am a client.
Starting to do a BiDi Streaming RPC...
Sending message: greeting:<first_name:"Mark" >
Received: Hello Mark!
Sending message: greeting:<first_name:"Chris" >
Received: Hello Chris!
Sending message: greeting:<first_name:"JD" >
Received: Hello JD!
Sending message: greeting:<first_name:"Stephan" >
Received: Hello Stephan!
Sending message: greeting:<first_name:"Deepak" >
Received: Hello Deepak!
```

and the server-side terminal get these messages:

```bash
$ go run greet/greet_server/server.go
Hello world!
GreetEveryone function was invoked with a streaming request
```

---

## 40. [Exercise] `FindMaximum` API

* **Goal**: to implement a `FindMaximum` RPC Bi-Directional Streaming API in a `CalculatorService`:
  * the function takes a stream of request message that has one integer, and returns a stream of responses that represent the current maximum between all these integers
  * remember to first implement the service definition in a `.proto` file, alongside the RPC message
  * implement the Server code first
  * Test the server code by implementing the Client

* Example:
  * The client will send a stream of number `(1,5,3,6,2,20)` and the server will respond with a stream of `(1,5,6,20)`.

---

## 41. [Solution] `FindMaximum` API

---