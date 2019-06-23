# Section 9. Data Evolution with Protobuf

---

## 49. The need for updating the protocol

* The need for updating the protocols
  * When you first declare a message in your protocol
    * you have defined sets of requirements
  * but as time go on, your business will evlove
    * you may have a different set of requirements
  * some fields may
    * change
    * be added
    * be removed
* many apps may read your messages using Protocol Buffers and you may not have time to upgrade them
  * e.g. asking for `FirstName` and `LastName` of customers as `v1` schema
    * but tomorrow (`v2` schema), we may need `PhoneNumber`
  * **Need to be able to evolve** the source data
    * but without breaking other appications reading it
  * **Protocol Buffers** help us tremendously with this situtation.

### Senario 1: Forward compatible change

Write data with New `.proto` -> Read data with Old `.proto`

### Senario 2: Backward compatible change

Write data with Old `.Proto` -> Read data with New `.proto`

---

## 50. Rules for Data Evolution

* Updating Protocol Rules (from the documentation)
  * i. do NOT change the numeric tags for any existing fields
  * ii. can add new fields, and old code will just ignore them
  * iii. likewise, if the new/old code reads unknown data
    * `default` will take place
  * iv. fields can be removed, as long as the tag number is NOT used again in your updated message type
    * You may want to rename the field instead, perhaps adding the prefix:
      * `OBSOLETE_`
    * or make the tag reserved:
      * so that future users of your `.proto` cannot accidentally reuse the specific tag number.
  * v. for data type change (e.g. `int32` to `int64`)
    * please refer to the doc
    * complecated and challenging
    * not recommend to do this
      * we can just create a new field rather

---

## 51. Adding Fields

---

## 52. Renaming Fields

---

## 53. Removing Fields

---

## 54. Reserved Keyword

---

## 55. Beware of Defualts

---

## 56. Evolbing Enum Fields

---
