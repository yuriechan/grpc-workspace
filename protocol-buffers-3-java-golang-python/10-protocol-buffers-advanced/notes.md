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

---

## 59. Protocol Buffers Options

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
