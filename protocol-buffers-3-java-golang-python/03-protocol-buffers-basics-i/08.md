# Section 3. Protocol Buffers Basics I

---

## 8. Tags

### Tags

* In Protocol Buffers, field names are not important!
  * but, when programming, the fields are important
* for protobuf, the important element is the **tag**
  * smallest tag: `1`
  * largest tag: `(2^29)-1` or 536,870,911
  * You also cannot use the numbers: `19000` through `19999`
    * **because**: reserved by Google for special usage
  
* Tags numbered from `1` to `15` use **1 byte** in space
  * so, use them for frequently populated fields
* Tags numbered from `16` to `2047` use **2 bytes**
  * less populated fields
* There's a concept of **reserved** tags
  * will cover this in advanced lectures

---