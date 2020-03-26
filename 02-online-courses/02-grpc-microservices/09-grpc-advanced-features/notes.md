# Session 09: [Hands-On] gRPC Advanced Feature Deep Dive

---

## 43. [Theory] Errors in gRPC

### 43.1. Error Codes

#### 43.1.1. Conventional HTTP Errors

* It is common for your API to something return error codes
* in HTTP, there are **many** error codes
  * 2xx for success
  * 3xx for ...
  * 4xx for ...
  * 5xx for ...
* Whilst HTTP codes are standardisied, they are not usually clear enough

#### 43.1.2. gRPC Errors

* with gRPC, there are a few error codes:
  * https://grpc.io/docs/guides/error.html
* there is also complete reference to implementation of error codes that close a lot of gaps with the documentation:
  * http://avi.im/grpc-errors
* if an application needs to return extra information on the top of an error code, it can use the metadata context

**General Errors:**

| **Case** | **Status Code** |
| --- | --- |
| Client application cancelled the request | `GRPC_STATUS_CANCELLED` |
| Deadline expired before server returned status | `GRPC_STATUS_DEADLINE_EXCEEDED` |
| Method not found on server | `GRPC_STATUS_UNIMPLEMENTED` |
| Server shutting down | `GRPC_STATUS_UNAVAILABLE` |
| Server threw an exception (or did something other than returning a status code to terminate the RPC) | `GRPC_STATUS_UNKNOWN` |

**Network Failures:**

| **Case** | **Status Code** |
| --- | --- |
| No data transmitted before deadline expires. Also applies to cases where some data is transmitted and no other failures are detected before the deadline expires | `GRPC_STATUS_DEADLINE_EXCEEDED`
| Some data transmitted (for example, the request metadata has been written to the TCP connection) before the connection breaks | `GRPC_STATUS_UNAVAILABLE`

**Protocol Errors:**

| **Case** | **Status Code** |
| --- | --- |
| Could not decompress but compression algorithm supported | `GRPC_STATUS_INTERNAL`
| Compression mechanism used by client not supported by the server | `GRPC_STATUS_UNIMPLEMENTED`
| Flow-control resource limits reached | `GRPC_STATUS_RESOURCE_EXHAUSTED`
| Flow-control protocol violation | `GRPC_STATUS_INTERNAL`
| Error parsing returned status | `GRPC_STATUS_UNKNOWN`
| Unauthenticated: credentials failed to get metadata | `GRPC_STATUS_UNAUTHENTICATED`
| Invalid host set in authority metadata | `GRPC_STATUS_UNAUTHENTICATED`
| Error parsing response protocol buffer | `GRPC_STATUS_INTERNAL`
| Error parsing request protocol buffer | `GRPC_STATUS_INTERNAL`

---

## 44. [Hands-On] Errors implementation

### 44.1. Error Codes: Hands-On

* let's implement an error message for a new `SquareRoot` Unary API
* we'll create `SquareRoot` RPC
* we'll implement `Server` with the error handling
* we'll implement `Client` with the error handling

### 44.2. the Implementation

#### 44.2.1. Protobuf

* define the `SquareRoot` RPC first, with unary req and resp):

```proto
...

message SquareRootRequest {
  int32 number = 1;
}

message SquareRootResponse {
  double number_root = 1;
}

service CalculatorService {
  ... // previous RPC definitions

  rpc SquareRoot(SquareRootRequest) returns (SquareRootResponse){};
}
```

* generate `caculator.pb` by using:

```bash
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
```

* One of important things we should do: "documentation for error handling"
  * `// make a comment about error handling for each rpc`

```proto
...
service CalculatorService {
  ... // previous RPC definitions

  // error handling
  // this RPC will throw an exception if the sent number is negative
  // the error being sent is of type `INVALID_ARGUMENT`
  rpc SquareRoot(SquareRootRequest) returns (SquareRootResponse){};
}
```

#### 44.2.2. Server

now, we can go to the `calculator_server/server.go` and have to implement `SquareRoot`

```go
...

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
  fmt.Println("Received SquareRoot RPC")
  number := req.GetNumber()
  if number < 0 {
    return nil, status.Errorf(
      codes.InvalidArgument,
      fmt.Sprintf("Received a negative number: %v", number),
    )
  }
  return &calculatorpb.SquareRootResponse{
    NumberRoot: math.Sqrt(float64(number)),
  }, nil
}

...
```

#### 44.2.3. Client

The server is now supporting the `SquareRoot`, we need to update the client, `calculator_client/client.go`

```go
...

func main() {
  fmt.Println("Calculator Client")
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("could not connect: %v", err)
  }
  defer cc.Close()

  c := calculatorpb.NewCalculatorServiceClient(cc)

  ...

  doErrUnary(c)
}

...

func doErrUnary(c calculatorpb.CalculatorServiceClient) {
  fmt.Println("Starting to do a SquareRoot Unary RPC...")

  // correct call
  doErrorCall(c, 10)

  // error call
  doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
  resp, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
  if err != nil {
    // err converted into the respErr
    // which is gRPC-friendly error that has message and code
    // but only if this `err` is actual gRPC error
    // if NOT, you'll get `ok` being false,
    // you can throw the normal error (the else clause)
    respErr, ok := status.FromError(err)
    if ok {
      // actual error from gRPC (user error)
      fmt.Printf("Error message from server: %v\n", respErr.Message())
      fmt.Println(respErr.Code())
      if respErr.Code() == codes.InvalidArgument {
        fmt.Println("We probably sent a negative number!")
        return
      }
    } else {
      // bigger error, framework type of error
      log.Fatalf("Big Error calling SquareRoot: %v", err)
      return
    }
  }
  fmt.Printf("Result of square root of %v: %v\n", n, resp.GetNumberRoot())
}

...
```

* for the `codes.*` of `codes. we can also use another errors:
  * `codes.NotFound`
  * `codes.OutOfRange`
  * `codes.PermissionDenied`
  * and more ...

#### 44.2.4. Run Server and Client

Now, the server and client are written. Let's run the server:

```bash
go run calculator/calculator_server/server.go
```

and you will see this message:

```bash
Calculator Server
```

then, let's run the client:

```bash
go run calculator/calculator_client/client.go
```

by running the client, the client can send requests to the server so the server-side get this messages:

```bash
Received SquareRoot RPC
Received SquareRoot RPC
```

because client send request 2 times: one is valid request, the other is invalid request

and the client-side get the responses:

```bash
Calculator Client
Starting to do a SquareRoot Unary RPC...
Result of square root of 10: 3.1622776601683795
Error message from server: Received a negative number: -2
InvalidArgument
We probably sent a negative number!
```

---

## 45. [Theory] Deadlines

### 45.1. gRPC Deadlines

* Deadlines allow gRPC clients to specify how long they are willing to wait for an RPC to complete before the RPC is terminated with the error `DEADLINE_EXCEEDED`
* **The gRPC documentation recommends you set a deadline for ALL client RPC calls**
* Setting the deadline is up to you: how long do you feel your API should have to complete?
  * usually small APIs: 100 ms, 500 ms, or 1000 ms (if slower response is okay)
  * for long API call: 5 min?
  * but, it's up to you
* The server should check if the deadline has exceeded and cancel the work it is doing
* This blog describes deadline in depth: `https://grpc.io/blog/deadlines`
* **NOTE**: Deadlines are propagated across if gRPC alls are chained
  * A => B => C (deadline for A is passed to B and then passed to C)
  * thus, C would be "aware" of the deadline of the client A

---

## 46. [Hands-On] Deadlines

### 46.1. gRPC Deadlines: Hands On

* We'll implement the server to return the response after 3000 ms
* The server will check if the client has cancelled the request
* We'll implement the client to set a deadline of 5000 ms
* We'll implement the client to set a deadline of 1000 ms

### 46.2. the Implementation

#### 46.2.1. Protobuf

one more API implementation under `greet.proto`:

```proto
...

message GreetWithDeadlineRequest {
  Greeting greeting = 1;
}

message GreetWithDeadlineResponse {
  string result = 1;
}

service GreetService {
  ...

  // Unary With Deadline
  rpc GreetWithDeadline(GreetWithDeadlineRequest) returns (GreetWithDeadlineResponse) {};
}
```

generate the code:

```bash
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
```

#### 46.2.2. Client

let's go to `greet/greet_client/client.go` and add implementation:

* add `doUnaryWithDeadline()` at the `main` func

```go
func main() {
  fmt.Println("Hello, I am a client.")

  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // WithInsecure() for just now testing
  if err != nil {
  log.Fatalf("Could not connect: %v", err)
  }

  defer cc.Close()

  c := greetpb.NewGreetServiceClient(cc)

  doUnaryWithDeadline(c, 5*time.Second) // should complete
  doUnaryWithDeadline(c, 1*time.Second) // should timeout
}
```

#### 46.2.3. Server

```go
...

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
  fmt.Printf("GreetWithDeadline function was invoked with %v\n", req)
  for i := 0; i < 3; i++ {
    if ctx.Err() == context.Canceled {
      // the client canceled the request
      fmt.Println("The client canceled the request!")
      return nil, status.Error(codes.Canceled, "the client canceled the request!")
    }
    time.Sleep(1 * time.Second)
  }
  firstName := req.GetGreeting().GetFirstName()
  result := "Hello " + firstName
  res := &greetpb.GreetWithDeadlineResponse{
    Result: result,
  }
  return res, nil
}

...
```

#### 46.2.4. Run the Codes

run the server:

```bash
go run greet/greet_server/server.go
```

and your _server terminal_ shows:

```bash
Hello world!
```

and open another terminal and run the client:

```bash
go run greet/greet_client/client.go
```

your _server terminal_ outputs:

```bash
GreetWithDeadline function was invoked with greeting:<first_name:"Mark" last_name:"Hahn" >
GreetWithDeadline function was invoked with greeting:<first_name:"Mark" last_name:"Hahn" >
```

and your _client terminal_ returns:

```bash
Hello, I am a client.
Starting to do a UnaryWithDeadline RPC...
2020/03/19 02:15:19 Response from GreetWithDeadline: Hello Mark
Starting to do a UnaryWithDeadline RPC...
Timeout was hit! Deadline was exceeded
```

---

## 47. [Theory] SSL Security

* Transport Layer Security (TLS), and its now-deprecated predecessor, Secure Sockets Layer (SSL), are cryptographic protocols designed to provide communications security over a computer network.
  * details: [wikipedia](https://en.wikipedia.org/wiki/Transport_Layer_Security)

* SSL Encryption in gRPC
  * In production gRPC call should be running with encryption enabled
  * This is done by generating SSL certificates
  * SSL allows communication to be secure end-to-end and ensuring no man in the middle attack can be performed

How SSL works?

* The need for SSL Encryption (1/3)
  * when you communicate over the internet,
    * your data is visible by all servers that transfer your packet
  * any router in the middle can view the packets you're sending using **PLAIN TEXT**
  * **It is not secure enough when the data is sensitive

```note
       -----                              ----------------                          ---------
HTTP  | You | -------------------------- | U: admin       | ---------------------- | Website |
      |     |                            | P: supersecret |                        |         |
       -----                              ----------------                          ---------

       -----       ----------------       ------------       ----------------       ---------
HTTP  | You | --- | U: admin       | --- |   Router   | --- | U: admin       | --- | Website |
      |     |     | P: supersecret |     | (Internet) |     | P: supersecret |     |         |
       -----       ----------------       ------------       ----------------       ---------
```

* The need for SSL Encryption (2/3)
  * SSL allows clients and servers to encrypt packet

```note
       -------------                      ----------------                          ---------
SSL   | gRPC Client | ------------------ | eBCskdsjEQdkfj | ---------------------- | Website |
      |             |                    | VfEGezKehbB... |                        |         |
       -------------                      ----------------                          ---------

       ----------------                                                      ----------------  
      | U: admin       |                                                    | eBCskdsjEQdkfj |
      | P: supersecret | --- SSL Encryption --        -- SSL Decryption --- | VfEGezKehbB... |
       ----------------                      |        |                      ----------------
                                             |        |
       ----------------                      |        |                      ----------------
      | eBCskdsjEQdkfj | <--------------------        --------------------> | U: admin       |
      | VfEGezKehbB... |                                                    | P: supersecret |
       ----------------                                                      ----------------

```

* The need for SSL Encryption (2/3)
  * SSL enables clients and servers to securely exchanges data
  * Routers cannot view the content of the Internet packets
  * let's do a deep dive into how SSL works

```note
       --------       ----------------       ------------       ----------------       ---------
SSL   | gRPC   | --- | eBCskdsjEQdkfj | --- |   Router   | --- | eBCskdsjEQdkfj | --- | Website |
      | Client |     | VfEGezKehbB... |     | (Internet) |     | VfEGezKehbB... |     |         |
       --------      ----------------        ------------       ----------------       ---------
```

* What is SSL?
  * TLS (Transport Layer Security), successor of SSL, encrypts the connection between 2 endpoints for secure data exchange

    ```note
     ----------                                                                              ------------------------
    | computer | -----> SSL Protocol (https)/Secure and secret exchange of information ---> | https://www.google.com |
     ----------                                                                              ------------------------
    ```

  * based on SSL certificates (e.g. on the web browser, you may see: `Secure | https://www.google.com`)
  * two ways of using SSL (gRPC can use both):
    * 1-way verification, e.g. browser => WebServer (**ENCRYPTION**)
    * 2-way verification, e.g. SSL authentication (**AUTHENTICATION**)

* Detailed Setup of SSL for Encryption

```note
                                                                                                               -------------
                                                                                                              | Server PEM  |
                                                                                                              | Private Key |
                                                                                                               -------------
               ---
              |     -------------                          -------------    Send Signed Certificate             ------------
Certification |    |   CA Root   |   Trust Certificate    | Certificate | -----------------------------------> | Server CRT |
    Setup     |    |   Public    | ---------------------> |  Authority  |                                      | Signed by  |
              |    | Certificate |                        |    (CA)     | <----------------------------------- | CA         |
              |     -------------                          -------------       Request to sign server.crt       ------------
               ---        |                                                  certificate for myapi.example.com        |
                          |                                                                                           |
                          |                                                                                           |
               ---        V                                                                                           V
              |       --------  <-------- 1. Send signed SSL certificate --------------------------------  -------------------
   SSL        |      |  gRPC  | ------*                                                                   |   gRPC Server     |
Handshake     |      | Client | <----/    2. Verify SSL certificate from Server                           |                   |
              |      |        |                                                                           | myapi.example.com |
              |       --------  <======== 3. Secure SSL Encrypted Communication ========================>  -------------------
               ---
```

* i. Certificate Authority (CA): can be public if you want to assign a public URL (e.g. myapi.example.com); or you can have a private CA (e.g. internal URLs), in which case, you have to create or maintain that CA.
* ii. server can create a private key (server PEM Private Key)
  * by using the private key, it will generate a certificate request
  * basically, it'll ask CA for signing the certificate request
  * thus, CA gives signed certificate (`server.crt` is a certificate that's signed by CA)
    * server-side, we need to have
      * a private key
      * signed certificate
* iii. client-side needs "trust certificate" from the CA
  * CA usually issues "CA Root Public Certificate"
* iv. now, it's ready to do SSL Handshake
  * send signed SSL certificate
  * verify SSL certificate from server
  * secure SSL encrypted communication

> Wow, it's a LOT...

---
