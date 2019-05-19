# Section 3. Protocol Buffers Basics I

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