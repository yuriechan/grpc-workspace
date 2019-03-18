# Section 3. Protocol Buffers Basics I

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