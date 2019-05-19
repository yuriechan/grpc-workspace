# Section 3. Protocol Buffers Basics I

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