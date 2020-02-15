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
              Client Streaming API
************                               ************
*          * --> req ..., 3, 2, 1, 0 ----> *          *
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

### 32.1. Streaming Client API: Server Implementation

* Hands-on:
* We'll implement a Streaming Client `LongGreet` RPC
* As we'll see the API implementation will be a bit more difficult
* **NOTE**: the server will only respond to the client once the client is done sending request.
  * (but in theory, the server can respond whenever it wants.)

let's start with `greet/greet_server/server.go` and add these code:

```go
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
  fmt.Printf("LongGreet function was invoked with a streaming request\n")
  result := ""

  for {
    req, err := stream.Recv()
    if err == io.EOF {
      // we have finished reading the client stream
      return stream.SendAndClose(&greetpb.LongGreetResponse{
        Result: result,
      })
    }
    if err != nil {
      log.Fatalf("Error whilst reading client stream: %v", err)
    }

    firstName := req.GetGreeting().GetFirstName()
    result += "Hello " + firstName + "! "
  }
}
```

then make sure the server is runnable:

```bash
go run greet/greet_server/server.go
```

You'll see this message (printed by `main()`):

```bash
Hello World
```

---

## 33. Client Streaming API Client Implementation

### 33.1. Streaming Client API: Client Implementation

* Hands-on:
* We'll implement a client call for our Streaming Client RPC
* We'll test against our server that is running!

let's start with `greet/greet_client/client.go` and add these code:

```go
func doClientStreaming(c greetpb.GreetServiceClient) {
  fmt.Println("Starting to do a Client Streaming RPC...")

  requests := []*greetpb.LongGreetRequest{
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Mark",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Chris",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "JD",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Stephan",
      },
    },
    &greetpb.LongGreetRequest{
      Greeting: &greetpb.Greeting{
        FirstName: "Deepak",
      },
    },
  }

  stream, err := c.LongGreet(context.Background())  // since it's streaming, req it not required
  if err != nil {
    log.Fatalf("error whlist calling LongGreet: %v", err)
  }

  for _, req := range requests {
    fmt.Printf("Sending req: %v\n", req)
    stream.Send(req)
    time.Sleep(1000 * time.Millisecond)
  }

  res, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("error whilst receiving response from LongGreet: %v", err)
  }
  fmt.Printf("LongGreet Response: %v\n", res)
}
```

and let's revise the `main()` function:

```go
func main() {
  fmt.Println("Hello, I am a client.")

  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // WithInsecure() for just now testing
  if err != nil {
    log.Fatalf("Could not connect: %v", err)
  }

  defer cc.Close()

  c := greetpb.NewGreetServiceClient(cc)
  // fmt.Printf("Created client: %f", c)

  // doUnary(c)
  // doServerStreaming(c)
  doClientStreaming(c)
}
```

now we need to do:

1. run the server: `go run greet/greet_server/server.go`
2. run the client: `go run greet/greet_client/client.go`

when you run the server, you can only get this message from `main()` function:

```bash
Hello world!
```

then, in another terminal, run the client, you get this message at a terminal that runs server:

```bash
LongGreet function was invoked with a streaming request
```

in the terminal that runs the client you can see:

```bash
Hello, I am a client.
Starting to do a Client Streaming RPC...
Sending req: %v
 greeting:<first_name:"Mark" >
Sending req: %v
 greeting:<first_name:"Chris" >
Sending req: %v
 greeting:<first_name:"JD" >
Sending req: %v
 greeting:<first_name:"Stephan" >
Sending req: %v
 greeting:<first_name:"Deepak" >
LongGreet Response: result:"Hello Mark! Hello Chris! Hello JD! Hello Stephan! Hello Deepak! "
```

---

## 34. [Exercise] `ComputeAverage` API

* Goal: **to implement a `ComputeAverage` RPC Client Streaming API in a `CalculatorService`**:
  * the function takes a stream of Request message that has one integer, and returns a Response with a double that represents the computed average
  * Remember to first implement the service definition in a `.proto` file, alongside the RPC messages
  * Implement the Server code first
  * Test the server code by implementing the Client

* Example:
  * The client will send a stream of number `(1,2,3,4)` and the server will response with `(2.5`), because `(1,2,3,4)`/4=`2.5`

---

## 35. [Solution] `ComputeAverage` API

First, let's improve `calculator/calculatorpb/calculator.proto` file:

```proto
message ComputeAverageRequest{
  int32 number = 1;
}

message ComputeAverageResponse{
  double average = 1;
}

service CalculatorService {
  
  ...  
  
  rpc ComputeAverage(stream ComputeAverageRequest) returns (ComputeAverageResponse) {};
}
```

We add the new messages as Request and Response for the `ComputeAverage`.

then we can generate code:

```bash
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```

then, we need to implement the server: `calculator/calculator_server/server.go`

```go
func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
  fmt.Printf("Received ComputeAverage RPC: %v\n")

  sum := int32(0)
  count := 0

  for {
    req, err := stream.Recv()
    if err == io.EOF {
      average = float64(sum) / count
      return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
        Average: average, 
      })
    }
    if err != nil {
      log.Fatalf("Error whilst reading client stream: %v", err)
    }
    sum += req.GetNumber()
    count++
  }
}
```

The next is implementing the client: `calculator/calculator_client/client.go`

```go
func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
  fmt.Println("Starting to do a ComputeAverage Client Streaming RPC...")

  stream, err := c.ComputeAverage(context.Background())
  if err != nil {
    log.Fatalf("Error whlist opening stream: %v", err)
  }

  numbers := []int32{3, 5, 9, 54, 23}

  for _, number := range numbers {
    fmt.Printf("Sending number: %v\n", number  )
    stream.Send(&calculatorpb.ComputeAverageRequest{
      Number: number,
    })
  }

  res, err := stream.CloseAndRecv()
  if err != nil {
    log.Fatalf("Error whlist receiving response: %v", err)
  }

  fmt.Printf("The Average is: %v\n", res.GetAverage())
}
```

and update `main()` function:

```go
func main() {
  fmt.Println("Calculator Client")
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("could not connect: %v", err)
  }
  defer cc.Close()

  c := calculatorpb.NewCalculatorServiceClient(cc)
  // fmt.Printf("Created client: %f", c)

  // doUnary(c)

  // doServerStreaming(c)
  doClientStreaming(c)
}
```

Now, since the server and client is ready, we can run the server and client respectively.

```bash
go run calculator/calculator_server/server.go
```

```bash
go run calculator/calculator_client/client.go
```

as soon as you run the server, you'll see:

```bash
Calculator Server
```

and you can run the client:

```bash
Starting to do a ComputeAverage Client Streaming RPC...
Sending number: 3
Sending number: 5
Sending number: 9
Sending number: 54
Sending number: 23
The Average is: 18.8
```

the server side also shows:

```bash
Received ComputeAverage RPC
```

---
