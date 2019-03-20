# Section 3. Protocol Buffers Basics I

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

#K## `date.proto`

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