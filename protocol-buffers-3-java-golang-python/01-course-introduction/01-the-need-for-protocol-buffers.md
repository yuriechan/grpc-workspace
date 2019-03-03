# Section 1. Course Introduction

---

## 1. The need for Protocol Buffers

### An Evolution of data: **Comma Separated Values (CSV)**

* First CSV:

| rownum | column1 | column2 | column3 | column4 | column5 | column6 |
| --- | --- | --- | --- | --- | --- | --- |
| row1 | John | Doe | 25 | John.Doe | true | OK |
| row2 | Mary | Poppins | sixty | Mary.pop | yes | OK |
| row3 | Tom | Cruise | 45 | Tom.Cru |  |  |

* Advantages:
  * Easy to parse
  * Easy to read
  * Easy to make sense of

* Disadvantages:
  * The data types of elements has to be inferred and it not a guarantee
  * Parsing becomes tricky when data contains commas
  * Column names may or may not be there

### An Evolution of data: **Relational tables definitions**

* Relational table definitions add types:

```sql
CREATE TABLE distributors (
  did   integer PRIMARY KEY,
  name  varchar(40)
)
```

* Advantages:
  * Data is fully typed
  * Data fits in a table

* Disadvantages:
  * Data has to be flat
  * Data is stored in a database, and data definition will be different for each database
    * *tough to share that data across different programming languages or databases*, **additioanl customisation needed!!!**

### An Evolution of data: **JSON (JavaScript Object Notation)**

* JSON format can be shared across the network!

```json
{
  "id": "0001",
  "type": "donut",
  "name": "Cake",
  "image":
    {
      "url": "images/0001.jpg",
      "width": 200,
      "height": 200,
    },
  "thumbnail":
    {
      "url": "images/thumbnails/0001.jpg",
      "width": 32,
      "height": 32
    }
}
```

* Advantages:
  * Data can take any form (arrays, nested elements)
  * JSON is a widely accepted format on the web
  * JSON can be read by pretty much any language
  * JSON can be easily shared over a network
    * JSON is text!

* Disadvantages:
  * Data has no schema enforcing
    * **You can put ANYTHING, JSON won't complain!**
  * JSON Objects can be quite big in size because of repeated keys
  * No commnets, meetadata, documentation

### An Evolution of data: **Protocol Buffers**

* `example.proto`

```proto
syntax = "proto3";

message MyMessage {
  int32 id = 1;
  string first_name = 2;
  bool is_validated = 3;
}
```

* Protocol Buffers are defined by a `*.proto` text file
* You can easily read it and understand it as a human

* Advantages:
  * Data is fully typed
  * Data is compressed automatically
    * lesss CPU usage, when you read it
  * Schema (defined using `.proto` file) is needed to generate code and read the data
  * Documentation can be embedded in the schema
  * Data can be read across any languages
    * C#, Java, Go, Python, JavaScript, etc. ...
  * Schema can evolve over time, in a safe manner (schema evolution)
  * overall, **3-10x** smaller, **20-100x** faster than XML
  * Code is generated for you automatically!

* Disadvantages:
  * Protobuf support for some languages might be lacking 
    * but, the main ones are fine
  * Can't "open" the serialised data with a text editor
    * because it's compressed and serialised

* Today Protocol Buffers are used as Google for almost all their internal applicaitons. 
* They have over 48000 Protobuf messages types in 12000 `.proto` files
* If it's working for Google, there's a great chance it'll be working for you!

---