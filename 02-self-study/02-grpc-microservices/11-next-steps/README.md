# Session 11. Next Steps

---

## 68. gRPC Services in the Real World

### 68.1. In the real world, what does a service look like?

* Google uses gRPC for some of its most important cloud services:
  * Google Pub/Sub: [actual protobuf definition](https://github.com/googleapis/googleapis/blob/master/google/pubsub/v1/pubsub.proto)
  * Google Spanner: [actual protobuf definition](https://github.com/googleapis/googleapis/blob/master/google/spanner/v1/spanner.proto)

### 68.2. Importance of Comments

It's very important to keep comments like this:

```proto
// The service that an application uses to manipulate subscriptions and to
// consume messages from a subscription via the `Pull` method or by
// establishing a bi-directional stream using the `StreamingPull` method.
service Subscriber {
  option (google.api.default_host) = "pubsub.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/cloud-platform,"
      "https://www.googleapis.com/auth/pubsub";

  // Creates a subscription to a given topic. See the
  // <a href="https://cloud.google.com/pubsub/docs/admin#resource_names">
  // resource name rules</a>.
  // If the subscription already exists, returns `ALREADY_EXISTS`.
  // If the corresponding topic doesn't exist, returns `NOT_FOUND`.
  //
```

* the comments defines well about the service
* it's also telling us how should it works for 2 cases especially comments like this:
  
  ```proto
  // If the subscription already exists, returns `ALREADY_EXISTS`.
  // If the corresponding topic doesn't exist, returns `NOT_FOUND`.
  ```

### 68.3. Not only for gRPC but REST

When you see the code like this:

```proto
  rpc CreateTopic(Topic) returns (Topic) {
    option (google.api.http) = {
      put: "/v1/{name=projects/*/topics/*}"
      body: "*"
    };
    option (google.api.method_signature) = "name";
  }
```

* this is a way to support [REST](https://en.wikipedia.org/wiki/Representational_state_transfer). If you are interested in, please check below:
  * [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)

Try to have a look how `.proto` files are defined in the real world and see the use-cases and implementation.

---

## 69. What about Gogo?

* In the open source world, it is not unusual to see active community members fork a popular project and add features they find the community needs.

* One notable project is "`gogo`", created by Walter Schulze, which aims to provide an alternative at how Protocol Buffers and gRPC are implemented in Go.

* This different implementation from the one provided by Google has many benefits, mainly around performance and usability, as well as providing more code generation options, which makes this project one of the most important in the gRPC ecosystem today.

* You can find the project alongside its 1800+ stars on GitHub here: https://github.com/gogo/protobuf

* If you'd like to read about it, I strongly recommend this blog: https://jbrandhorst.com/post/gogoproto/

* Regardless of the implementation you will use for your projects, the gRPC concepts are the exact same. Happy learning!

---

## 70. Next Steps

* read the documentation: [https://grpc.io/](https://grpc.io/)
* trying migrate your REST API to gRPC or use a new project to experiment with gRPC
* the gRPC ecosystem is booming,
  * stay up to date with the gRPC & Protocol Buffers release and on the lookout for new projects
* play with the technology and have some fun!

---
