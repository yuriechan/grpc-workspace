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

## 9. Repeated Fields

### Repeated Fields

* To make a "list" or an "array"
  * you can use the concept of **`repeated` fields**
* the list can take any number (0 or more) of elements you want
* the opposite of **`repeated`** is "singular" (we don't write it)
* let's add a list of phone numbers to our `Person` example!

#### `2-repeated-fields.proto`

```proto
syntax = "proto3"

message Person {
  int32 age = 1;
  string first_name = 2;
  string last_name = 3;
  bytes small_picture = 4;
  bool is_profile_verified = 5;
  float height = 6;

  repeated string phone_number = 7;
}
```

---

## 10. Comments

### Comments

* It's possible to embed comments in your `.proto` file.
* It's actually recommended to use comments as a form of documentation for oyur schemas.
* Comments can be of these two forms:

  ```proto
  // this is a comment
  ```

  ```proto
  /* this is a
   * multiline comment */
  ```

* Adding comment to the example: `Person`

### `3-comments.proto`

```proto
// The syntac for this file is proto3
syntax = "proto3";

/* Person is used to identify users
 * across our system */
message Person {
  // the age as of the person's creation
  int32 age = 1;
  // the first name as documented in the signup form
  string first_name = 2;
  string last_name = 3; // last name as documented in the signup form
  // small_picture represents a small .jpg file
  bytes small_picture = 4;
  bool is_profile_verified = 5;
  // height of the person in cms
  float height = 6;

  // a list of phone numbers that is optional to provide at signup
  repeated string phone_numbers = 7;
}
```

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

## 12. Enumerations (Enums)

### Enums

* If you know all the values a field can take in advance,
  * you can leverage the **Enum** type
* **<u>The first value of an Enum is the default value</u>**
* **Enum** must start by the tag `0`
  * which is the default value
* Example of an Enum with the `Person` example for:
  * field Eye Colour: (blue, brown, green)

### `4-enums.proto`

```proto
// The syntac for this file is proto3
syntax = "proto3";

/* Person is used to identify users
 * across our system */
message Person {
  // the age as of the person's creation
  int32 age = 1;
  // the first name as documented in the signup form
  string first_name = 2;
  string last_name = 3; // last name as documented in the signup form
  // small_picture represents a small .jpg file
  bytes small_picture = 4;
  bool is_profile_verified = 5;
  // height of the person in cms
  float height = 6;

  // a list of phone numbers that is optional to provide at signup
  repeated string phone_numbers = 7;

  // we currently consider only 3 eye colours
  enum EyeColour {
    UNKNOWN_EYE_COLOUR = 0;
    EYE_GREEN = 1;
    EYE_BROWN = 2;
    EYE_BLUE = 3;
  }

  // it's an enum as defined above
  EyeColour eye_colour = 8;
}
```

so in the example, first we need to **define** enum

```proto
 enum EyeColour {
    UNKNOWN_EYE_COLOUR = 0;
    EYE_GREEN = 1;
    EYE_BROWN = 2;
    EYE_BLUE = 3;
  }

  EyeColour eye_colour = 8;
```

therefore, in the `enum EyeColour {}` block, we define the Enum type first, then we also need to declare the type to use as `EyeColour eye_colour = 8;`

The last line of the code means we use the enum type `EyeColour`, which we will call as `eye_colour` and its tag is `8`.

**AGAIN**, do NOT forget to give the default value first (in this case `UNKNOWN_EYE_COLOUR` was defined first), then create some elements (like `EYE_GREEN`). Usually, these elements are all capitalised.

---

## 13. Practice Exercises I

### Create the following messages in 5 different files:

* Date (example: 2018/03/15). With fields:
  * Year (number)
  * Month (number)
  * Day (number)
* Latitude Longitude (example: -33.865382, 151.192861). With fields:
  * Latitude (number)
  * Longitude (number)
* Money (example: USD 35.42  (35 = integral, 42 = decimal)):
  * Currency Code (string)
  * Integral amount (number)
  * Decimal amount (number)
* DayOfWeek (example Monday) using Enum, and keeping the 0 as “undefined” day
* Person (example: Michael Jordan):
  * First Name
  * List of middle names
  * Last Name

When you're done move on to the next lecture for solutions!

### Personal Code

### `date.proto`

```proto
syntax = "proto3";

message Date {
  int32 Year = 1;
  int32 Month = 2;
  int32 DAy = 3;
}
```

#### `latitude-longitude.proto`

```proto
syntax = "proto3";

message LatitudeLongitude {
  double Latitude = 1;
  double Longitude = 2;
}
```

#### `money.proto`

```proto
syntax = "proto3";

message Money {
  string CurrentyCode = 1;
  double IntegralAmount = 2;
  double DecimalAmount = 3;
}
```

#### `day-of-week.proto`

```proto
syntax = "proto3";

enum DayOfWeek {
  DAY_UNSPECIFIED = 0;
  MONDAY = 1;
  TUESDAY = 2;
  WEDNESDAY = 3;
  THURSDAY = 4;
  FRIDAY = 5;
  SATURDAY = 6;
  SUNDAY = 7;
}
```

#### `person.proto`

```proto
syntax = "proto3";

message Person {
  string first_name = 1;
  repeated middle_names = 2;
  string last_name = 3;
}
```

---

## 14. Solution for Practice Exercises I

### `date.proto`

```proto
syntax = "proto3";

message Date {
  // Year of date. Must be from 1 to 9999, or 0 if specifying a date without
  // a year.
  int32 year = 1;

  // Month of year. Must be from 1 to 12.
  int32 month = 2;

  // Day of month. Must be from 1 to 31 and valid for the year and month, or 0
  // if specifying a year/month where the day is not significant.
  int32 day = 3;
}
```

### `dayofweek.proto`

```proto
syntax = "proto3";

// Represents a day of week.
enum DayOfWeek {
  // The unspecified day-of-week.
  DAY_OF_WEEK_UNSPECIFIED = 0;

  // The day-of-week of Monday.
  MONDAY = 1;

  // The day-of-week of Tuesday.
  TUESDAY = 2;

  // The day-of-week of Wednesday.
  WEDNESDAY = 3;

  // The day-of-week of Thursday.
  THURSDAY = 4;

  // The day-of-week of Friday.
  FRIDAY = 5;

  // The day-of-week of Saturday.
  SATURDAY = 6;

  // The day-of-week of Sunday.
  SUNDAY = 7;
}
```

### `latlng.proto`

```proto
syntax = "proto3";

message LatLng {
  // The latitude in degrees. It must be in the range [-90.0, +90.0].
  double latitude = 1;

  // The longitude in degrees. It must be in the range [-180.0, +180.0].
  double longitude = 2;
}
```

### `money.proto`

```proto
syntax = "proto3";

// Represents an amount of money with its currency type.
message Money {
  // The 3-letter currency code defined in ISO 4217.
  string currency_code = 1;

  // The whole units of the amount.
  // For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
  int64 units = 2;

  // Number of nano (10^-9) units of the amount.
  // The value must be between -999,999,999 and +999,999,999 inclusive.
  // If `units` is positive, `nanos` must be positive or zero.
  // If `units` is zero, `nanos` can be positive, zero, or negative.
  // If `units` is negative, `nanos` must be negative or zero.
  // For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
  int32 nanos = 3;
}
```

### `person.proto`

```proto
syntax = "proto3";

message Person {
  string first_name = 1;
  repeated string middle_names = 2;
  string last_name = 3;
}
```

---
