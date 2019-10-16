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

## 25.1. Hands On: `GreetManyTimes` API Definition

* <u>Hands On</u>: Let's define a Streaming Server "`GreetManyTimes`" API
* It will take **ONE** `GreetManyTimesRequest` that contains a `Greeting`
* It will return **MANY** `GreetManyTimesResponse` that contains a result string

## 25.2. `greet.proto`

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

message GreetManyTimesRequest {
    Greeting greeting = 1;
}

message GreetManyTimesResponse {
    string result = 1;
}

service GreetService {
    // Unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

    // Server Streaming
    rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManyTimesResponse) {};
}
```

actually in the `rpc GreetManyTimes()GreetManyTimesRequest) returns (stream GreetManyTimesResponse)`, we can reuse `GreetRequest` and `GreetResponse` rather than `GreetManyTimesRequest` and `GreetManyTimesResponse`, respectively, HOWEVER, usually in rpc, when you define a new rpc, you should create new request and response message types.

Next step: in your `generate.sh` file, get this command and run it:

```sh
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

---

## 26. Server Streaming API Server Implementation 

if we goes to `greet.pb.go`, we can find `interface` of `GreetServiceServer` and `GreetServerClient` we can find `GreetManyTimes()`

```go
// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
}

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	// Server Streaming
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
}
```

at first, we need to implement `GreetManyTimes` under `GreetServiceServer`

and we can add this implementation: 

```go
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}
```

---

## 27. Server Streaming API Client Implementation 

* hands-on:
* We'll implement a client call for our Streaming Server RPC
* We'll test it against our server that is running!

```go
func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mark",
			LastName: "Hahn",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatal("error whilst calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatal("error whilst reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
```

---

## 28. [Exercise] `PrimeNumberDecomposition` API

* goal:
  * to implement a `PrimeNumberDecomposition` PRC Server Streaming API in a `CalculatiorService`
    * The function takes a Request message that has one integer, and returns a stream of Responses that represent the prime number decomposition of that number (see below for the algorithm).
    * Remember to first implement the service definition in a .proto file, alongside the RPC messages
    * Implement the Server code first
    * Test the server code by implementing the Client

* example:
  * The client will send one number (120) and the server will respond with a stream of (2,2,2,3,5), because 120=2*2*2*3*5
  * pseudo code:

```py
k = 2
N = 210
while N > 1:
    if N % k == 0:   // if k evenly divides into N
        print k      // this is a factor
        N = N / k    // divide N by k so that we have the rest of the number left.
    else:
        k = k + 1
```

---

## 29. [Solution] `PrimeNumberDecomposition` API

current `calculator.proto` looks like this:

```proto
syntax = "proto3";

package calculator;
option go_package = "calculatorpb";

message SumRequest {
  int32 first_number = 1;
  int32 second_number = 2;
}

message SumResponse {
  int32 sum_result = 1;
}

service CalculatorService {
  rpc Sum(SumRequest) returns (SumResponse) {};
}
```

and we need to add: `rpc PrimeNumberDecomposition(PrimeNumberDecompositionRequest) returns (stream PrimeNumberDecompositionResponse) {};`, also define each new message like below:

```proto
syntax = "proto3";

package calculator;
option go_package = "calculatorpb";

message SumRequest {
  int32 first_number = 1;
  int32 second_number = 2;
}

message SumResponse {
  int32 sum_result = 1;
}

message PrimeNumberDecompositionRequest{
  int64 number = 1;
}

message PrimeNumberDecompositionResponse{
  int64 prime_factor = 1;
}

service CalculatorService {
  rpc Sum(SumRequest) returns (SumResponse) {};

  rpc PrimeNumberDecomposition(PrimeNumberDecompositionRequest) returns (stream PrimeNumberDecompositionResponse) {};
}
```

then we can generate target code in *.go with this command:

```bash
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```

Now, we can implment the `server.go`: need to implement a new function




---
 
