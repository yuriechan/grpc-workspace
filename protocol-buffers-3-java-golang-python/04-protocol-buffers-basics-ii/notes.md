# Section 4. Protocol Buffer Basics II

---

## 15. Defining Multiple Messages in the Same File

### Defining multiple Messages in the same `.proto` file

* It is possible, in the same `.proto` file, to define multiple types
* It is then super easy to reference them if we want to
* Let's create a message `Date` and add that to our `Person` as a file for a birthday

we can add these code to our previous working code:

```proto
message Person {
  ...
  ...
  
  // Person's birthday
  Date birthday = 9;
}

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

### `1-same-level-message.proto

```proto
// The syntax for this file is proto3
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

  // Person's birthday
  Date birthday = 9;
}

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

---

## 16. Nesting Messages

### Nesting Types

* It is possible to define types within types
* The reasons could be:
  * Avoding naming conflicts
  * Enforcing some level of "locality" for that type
* You can nest types as deeply as you want
* Let's create a field `Address` and use that in our `Person` to have multiple addresses

we can add this code to the previous code:

```proto
message Address {
  string address_line_1 = 1;
  string address_line_2 = 2;
  string zip_code = 3;
  string city = 4;
  string country = 5;
}

// multiple addresses
repeated Address addresses = 10;
```

### `2-nested-messages.proto`

```proto
// The syntax for this file is proto3
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

  // Person's birthday
  Date birthday = 9;

  // we define the type Address within Person (full name is Person.Address)
  message Address {
    string address_line_1 = 1;
    string address_line_2 = 2;
    string zip_code = 3;
    string city = 4;
    string country = 5;
  }

  // multiple addresses
  repeated Address addresses = 10;
}

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

---

## 17. Imports

### Importing Types

* You can also have different types in different `.proto` files
* This is useful if you want to re-use code and import other `.proto` files created by people in your team
* Let's move our `Data` out of our `Person` file and import the date file instead!

* We can create 2 files: `3-date.proto` and `3-person-with-imports.proto`
* The `3-person-with-imports.proto` can import `3-date.proto`

```proto
import "3-basics-part-ii/3-date.proto";
```

---

### `3-date.proto`

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

### `3-person-with-imports.proto`

```proto
// The syntax for this file is proto3
syntax = "proto3";

import "3-basics-part-ii/3-date.proto";

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

  // Person's birthday
  Date birthday = 9;

  // we define the type Address within Person (full name is Person.Address)
  message Address {
    string address_line_1 = 1;
    string address_line_2 = 2;
    string zip_code = 3;
    string city = 4;
    string country = 5;
  }

  // multiple addresses
  repeated Address addresses = 10;
}
```

---

## 18. Packages

### Packages

* It is very important to define the packages in which your protocol buffer messages live
  * when your code gets compiled, it will be placed at the pakage you indicated
  * It also helps to prevent name conflicts between messages (`my.package.Person`)
* Packages will help all the different languages compile correctly from `.proto` files
  * Java, C#, Python, Go, etc...)

So, in this case if we define a package like this in `a.proto`

```proto
package my.date;

message Date {
  // Year of date. Must be from 1 to 9999, or 0 if specifying a date without
  // a year.
  int32 year = 1;
  ...
```

and `b.proto` imports `a.proto`

```proto
import a.proto
```

and want to use the `message Date` in `b.proto`, we have to use the package name like this:

```proto
// Person's birthday
my.date.Date birthday = 9;
```

---

### `4-date-with-package.proto`

```proto
syntax = "proto3";

package my.date;

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

---

### `4-person-with-package.proto`

```proto
// The syntax for this file is proto3
syntax = "proto3";

import "3-basics-part-ii/4-date-with-package.proto";

package person;

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

  // Person's birthday
  my.date.Date birthday = 9;

  // we define the type Address within Person (full name is Person.Address)
  message Address {
    string address_line_1 = 1;
    string address_line_2 = 2;
    string zip_code = 3;
    string city = 4;
    string country = 5;
  }

  // multiple addresses
  repeated Address addresses = 10;
}
```

---

## 19. Practice Exercise II

### Time to Practice!

* **Create the following messages**:
  * City with the following fields:
    * Name
    * Zip Code
    * Country Name
  * Street with the following fields:
    * Street Name
    * City
  * Building with the following fields:
    * Building name
    * Building number
    * Street

* **Organise these messages in 4 ways**:
  * All in one .proto file, as same level messages
  * All in one .proto file, as nested messages
  * Separate files with imports
  * Separate files with imports and packages

**When you're done move on to the next lecture for solutions!**

---

## 20. Solution to Practice Exercise II

### Same File

#### `all-in-one-file.proto`

```proto
syntax = "proto3";

message City {
  string name = 1;
  string zip_code = 2;
  string country_name = 3;
}

message Street {
  string name = 1;
  City city = 2;
}

message Building {
  string building_name = 1;
  int32 building_number = 2;
  Street street = 3;
}
```

### Nested Messages

#### `nested-messages.proto`

```proto
syntax = "proto3";

message Building { // level 0
  string building_name = 1;
  int32 building_number = 2;

  message Street { // level 1
    string name = 1;

    message City { // level 2
      string name = 1;
      string zip_code = 2;
      string country_name = 3;
    }

    City city = 2;
  }

  Street street = 3;
}
```

### Imports

#### `building.proto`

```proto
syntax = "proto3";

// remember to use the full path from the root of the project
import "4-exercises-solutions/3-imports/street.proto";

message Building {
  string building_name = 1;
  int32 building_number = 2;
  Street street = 3;
}
```

#### `city.proto`

```proto
syntax = "proto3";

message City {
  string name = 1;
  string zip_code = 2;
  string country_name = 3;
}
```

#### `street.proto`

```proto
syntax = "proto3";

// remember to use the full path from the root of the project
import "4-exercises-solutions/3-imports/city.proto";

message Street {
  string name = 1;
  City city = 2;
}
```

### Imports with Packages

#### `building.proto`

```proto
syntax = "proto3";

package building;

// remember to use the full path from the root of the project
import "4-exercises-solutions/4-imports-with-packages/street.proto";

message Building {
  string building_name = 1;
  int32 building_number = 2;
  street.Street street = 3;
}
```

#### `city.proto`

```proto
syntax = "proto3";

package city;

message City {
  string name = 1;
  string zip_code = 2;
  string country_name = 3;
}
```

#### `street.proto`

```proto
syntax = "proto3";

// remember to use the full path from the root of the project
import "4-exercises-solutions/4-imports-with-packages/city.proto";

package street;

message Street {
  string name = 1;
  city.City city = 2;
}
```
