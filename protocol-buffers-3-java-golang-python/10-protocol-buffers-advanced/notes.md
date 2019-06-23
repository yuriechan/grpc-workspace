# Section 10. Protocol Buffers Advanced

---

## 57. Integer Types Deep Dive

* Integer Types
  * many ways to represent an integer in protocol buffers:
    * `int32`
    * `int64`
    * `uint32`
    * `uint64`
    * `sint32`
    * `sint64`
    * `fixed32`
    * `fixed64`
    * `sfixed32`
    * `sfixed64`
  * each tpye is constructed to handle:
    * i. range of allowed values: 64 bits has wider value range than 32 bits
    * ii. whether negative values are allowed: `uint32` vs `sint32`
    * iii. size efficiency on serialisation
  * This is advanced topics for:
    * better performance
    * space optimisation

* Integer Types - Range of allowed values
  * 64 bits allow for a greater range
    * 32 bit:
      * `int32` / `sint32`: -2,147,483,648 to 2,147,483,647
      * `uint32`: 0 to 4,294,967,295
    * 64 bit:
      * `int64` / `sint64`: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
      * `uint64`: 0 to 18,446,744,073,709,551,615

* Integer Types - Negative Values
  * i. `uint32`, `uint64` do NOT allow negative values
  * ii. `int32`, `int64` do NOT encode negative values efficiently
    * negative values constantly use 10 bytes in spaec
  * iii. `sint32`, `sint64` encode negative values well (through the use of a technique called ZigZag)
* choose accordingly based on if your field can have negataive values!

* Integer Types - Size efficiency
  * `uint32`, `uint64`, `int32`, `int64` `sint32`, `sint64` are variable encoding meaning that if they can use less space, they will (for small values)
  * `fixed32` use 4 bytes constantly
    * more efficient than `uint32` if values are often greater than 2^28
  * `fixed64` use 8 bytes constantly
    * more efficient than `uint64` if vaules are often greater than 2^56

---

## 58. Advanced Data Types (`oneof`, `map`, `Timestamp`, and `Duration)

* Advacned Types - `oneof`
  * you can use `oneof` to tell protocol buffers that only one field can have a value:

  ```proto
  message MyMessage {
    int32 id = 1;
    one of example_oneof {
      string my_string = 2;
      bool my_bool = 3;
    }
  }
  ```

  * `oneof` fields **CANNOT** repeated
  * evolving schemas using `oneof` is complicated
    * see documentation if you really need
  * on read, all fields will be `null` except the last one that was set at write

* Advanced Types - `map`s
  * maps can be used to map scalars (except `float` / `double`) to values of any type

  ```proto
  message MyMessage {
    int32 id = 1;
    map<string, Result> results = 2;
  }
  ```

  * map fields cannot be repeated
  * there's no ordering for `map`
    * it's key => value store

* Advanced Types - `Timestamps` (Well Known Types)
  * protocol buffers contain a set of Well Known Types
    * e.g. advanced types known to all programming languages
  * the list is here: [link](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/timestamp)
  * one of the types is `Timestamp`
    * fields are:
      * `seconds`
      * `nanoseconds` (UTC)
  * do not forget to use `import` statement

  ```proto
  syntax = "proto3";

  import "google/protobuf/timestamp.proto";

  message MyMessage {
    google.protobuf.Timestamp my_field = 1;
  }
  ```

* Advanced Types - `Duration`
  * `Duration` is another Well Known Type
  * It represents the time span between two timestamps
  * It contains, just like `Timestamp`, `seconds` and `nanoseconds`

  ```proto
  syntax = "proto3";

  import "google/protobuf/timestamp.proto";
  import "google/protobuf/duration.proto";

  message MyMessage {
    google.protobuf.Timestamp msg_date = 1;
    google.protobuf.Duration validaty = 2;
  }
  ```

---

## 59. Protocol Buffers Options

* `option`s allow to alter the behaviour of the `protoc` compiler
  * when generating code for specific languages
* there are few bundled options
  * check the docs for details
  * examples:
  
  ```proto
  option csharp_namespace = "Google.Protobuf.WellKnownTypes";
  option cc_enable_arenas = true;
  option go_package = "github.com/golang/protobuf/ptypes/duration";
  option java_package = "com.google.protobuf";
  option java_outer_classname = "DurationProto";
  option java_multiple_files = true;
  option objc_class_prefix = "GPB";
  ```

---

## 60. Naming Conventions

---

## 61. Uber style guiding

---

## 62. Services

---

## 63. Introduction to gRPC (from gRPC Course)

---

## 64. Protocol Buffers Internals

---
