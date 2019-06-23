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
