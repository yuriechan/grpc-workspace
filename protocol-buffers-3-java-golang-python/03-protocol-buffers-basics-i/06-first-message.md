# Section 3. Protocol Buffers Basics I

---

## 6. First Message

### First message - Introduction

* Here's our first message:
  * using `proto3` in this course, so we define syntax as:
    * `syntax = "proto3";`
  * In Protocol Buffers we define messages
    * in the example, the message is: `message MyMessage { ... }`
  * and the message has its own fields
    * each field has field type:
      * `int32`, `string`, `bool`
    * each field has field name:
      * `id`, `first_name`, `is_validate`
    * each field has **field tag** (e.g. number)
      * `1`, `2`, `3`

```proto
syntax = "proto3";

message MyMessage {
  int32 id = 1;
  string first_name = 2;
  bool is_validated = 3;
}
```

---