# Section 3. Protocol Buffers Basics I

---

## 11. Default Values for Fields

### Default Values for fields

* All fields, if NOT specified or unknown,
  * will take a default value
* examples:
  * `bool`: `false`
  * number (`int32`, etc...): `0`
  * `string`: empty string
  * `bytes`: empty bytes
  * `enum`: first value
  * `repeated`: empty list

---