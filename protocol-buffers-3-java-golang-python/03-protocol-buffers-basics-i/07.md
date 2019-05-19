# Section 3. Protocol Buffers Basics I

---

## 7. Scalar Types

### Scala Types - Number

* Numbers can take various forms based on what values you expect them to have:
  * `double`, `float`, `int32`, `int64`, `unit32`, `uint64`, `sint32`, `sint64`, `fixed64`, `sfixed32`, `sfixed64`
* **Integer**: for now, let's use `int32`
  * _There's a discussion in the advanced section of advantages of each specific type_
* **Floating point numbers**: `float`, `double`
  * `float`: 32 bits
  * `double`: 64 bits - for more precision _(if you really need it)_

### Scala Types - Boolean

* Boolean can hold hte value:
  * `true`
  * `false`
* It's represented as `bool` in protobuf

### Scala Types - String

* String represents an arbitrary length of text
* It's represented as `string` in protobuf
* A string **MUST** always contain:
  * `UTF-8` encoded OR 7-bit `ASCII` text

### Scala Types - Bytes

* Bytes represent any sequence of byte array.
* It's represented as `bytes` in protobuf
* It'll be up to you to interpret what these bytes mean.
  * e.g. You could use these bytes to include a small image
  * or whatever you want!

### Scala Types - Summary

* Let's create a message `Person` that has:
  * `int32` (age)
  * `string` (first name)
  * `string` (last name)
  * `bytes` (small picture)
  * `bool` (profile verified)
  * `float` (height)

#### `1-scalar-types.proto`

```proto
syntax = "proto3";

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes small_picture = 4;
  bool is_profile_verified = 5;
  float height = 6;
}
```

This coude above is also available at [HERE](../codes/1-basics-part-i/1-scalar-types.proto)  

easy, eh? **Protocol Buffers, meant to be easy and easy to read!**

---