# Section 1. Course Introduction

---

## 2. How are Protocol Buffers used?

### How is Protocol Buffer Used? **To share data across languages!**

`.proto` files (human readable) -> *(automated generation of code)* -> Create Objects (in Java, Python, Go, etc.) -> *(encode/decode)* -> Serialised Data (can be interpreted by any language)

* Serialised Data
  * Universal:
    * can be read by any languages

With the `.proto` file, human-readable data, we can get the **sericalised data**, can be used in any programming languages.

### How is Protocol Buffer Used?

* Some databases may have support for Protocol Buffers data format
* lots of RPC frameworks, use Protocol Buffers to exchange data
  * gRPC (by Google)
* Google uses it for all their internal API
* Big projects
  * `etcd`
    * use Protocol Buffers for transporting data

### Proto2 vs Proto3

* Mid 2016, Google released the 3rd iteration of Protocol Buffers
  * `proto3`
* In this course:
  * only cover `proto3`
    * the most common format used forward
    * the best compatibility across a wide array of programming languages
* the easiest to learn

---