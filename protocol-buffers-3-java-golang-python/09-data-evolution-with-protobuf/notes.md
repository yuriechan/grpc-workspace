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

* Let's add a field to our schema with new tab number:

Current code is this:

```proto
message MyMessage {
  int32 id = 1;
}
```

and we evolve it as:

```proto
message MyMessage {
  int32 id = 1;
  string first_name = 2;
}
```

* if that field is sent to old code:
  * the old code will NOT know
    * what that tag number corresponds to
    * and the new field will be ignored / dropped
* oppositely, if we read old data with the new code
  * the new field will not be found
    * and the default value will be assumed (empty string)
* Therefore, **default values should always be interpreted with care**

---

## 52. Renaming Fields

* old schema:

```proto
message MyMessage {
  int32 id = 1;
  string first_name = 2;
}
```

* new schema:

```proto
message MyMessage {
  int32 id = 1;
  string person_first_name = 2;
}
```

* in ths case, nothing change except the field name:
  * field name can be changed easily
  * **tag numbers** are ONLY important for protobuf/machines
    * they do not care that much on human-readable field names

---

## 53. Removing Fields

* remove a field in our schema:

old data:

```proto
message MyMessage {
  int32 id = 1;
  string first_name = 2;
}
```

new data:

```proto
message MyMessage {
  int32 = id = 1;
}
```

* if old code does NOT find the deleted field message
  * the **default value** will be used
* oppositely, if we read old data with the new code
  * the deleted field will just be dropped
* Again, **default values should always be interpreted with care**

* Removing Fields - Reserving Tags
  * when removing a field
    * should **ALWAYS** reserve the **tag** and the **name**

before:

```proto
message MyMessage {
  int32 id = 1;
  string first_name = 2;
}
```

after:

```proto
message MyMessage {
  reserved 2;
  reserved "first_name";
  int32 id = 1;
}
```

* to prevent
  * the tag / field name to be reused
* necessary to prevent conflicts in the codebase

* Removing Fields - Make some fields obsolete
  * alternative way:
    * rename the field as: `OBSOLETE_field_name`
  * downside would be:
    * you may have to populate that field while your client get upgraded to use the newer field that replaces it (which has a new tag)
  * **Personally like the `reserved` keyword**

---

## 54. Reserved Keyword

* can reserve TAGS and FIELD NAMES
* can't mi TAGS and FIELD NAMES in the same `reserved` statement
* e.g.

```proto
message Foo {
  reserved 2, 15, 9 to 11;
  reserved "foo", "bar";
}
```

* we reserve TAGS to prevent new fields from reusing tags
  * because that would break old code at runtime
* we reserve FIELD NAMES to prevent code bugs
* **Do NOT EVER remove any reserved tags!!!**

---

## 55. Beware of Defualts

---

## 56. Evolbing Enum Fields

---
