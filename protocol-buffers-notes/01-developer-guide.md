# Developer Guide

* protocol buffers is a
  * language-neutral,
  * platform-netural,
  * extensible way
    * to serialising structured data
    * for use in communications protocols, data storage, and more

* filename extension: `*.proto`

## What are protocol buffers?

* protocol buffers are a:
  * flexible,
  * efficient,
  * automated
    * mechanism for serialising **structured data**
  * like XML, **BUT**
    * smaller,
    * faster,
    * simpler!

* define how you want your data to be structured once,
  * you can use **special generated source code**,
  * to easily write and read your structured data
    * to and from a variety of data streams
    * using a variety of languages

* Be able to update your data structure
  * WITHOUT breaking deployed program that are compiled againtest the **"older"** format

## How do they work?

* the information we want to serialising to be structed
  * by defining protocol buffers types in `.proto` files
* each protocol buffer message:
  * a small logical record of information
    * containing a series of **name-value** pairs

### An exmaple to define a person

```proto
message Person {
  required string name = 1;
  required int32 id = 2;
  optional string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    required string number = 1;
    optional PhoneType type = 2 [default = HOME];
  }

  repeated PhoneNumber phone = 4;
}
```

* each message type has one or more uniquely numbered fields
* each field has a name and a value type
  * where value types can be
    * numbers (integer or floating-point),
      * e.g. `required int32 id = 2;`
    * booleans,
    * strings,
      * e.g. `required string name = 1;`
    * raw bytes,
    * or even (as in the example above) other protocol buffer message types
      * e.g. `repeated PhoneNumber phone = 4;`
* allowing you to structure your data hierarchically

* can specify
  * optional fields
  * required fields
  * repeated fields
    * e.g. `repeated PhoneNumber phone = 4;`

* after defining a message (`.proto` file)
  * you can run **protocol buffer compiler**
  * for your application's language on your `.proto` file
    * to generate data access classes
  * details about how to write `.proto` files: [LINK](https://developers.google.com/protocol-buffers/docs/proto)
* These provide
  * simple accessors for
    * each field (e.g. `name()` and `set_name()`)
    * methods to serialise/parse the whole structure to/from raw bytes
      * e.g. with C++, running the compiler on the above example will generate a `class` called `Person`
        * you can use this `class` in your application to
          * populate,
          * serialise,
          * retrieve,
            * `Person` protocol buffer messaages
        * then we can write a code like below:

          ```cpp
          Person person;
          person.set_name("John Doe");
          person.set_id(1234);
          person.set_email("jdoe@example.com");
          fstream output("myfile", ios::out | ios::binary);
          person.SerializeToOstream(&output);
          ```

        * then you can read the message back in

        ```cpp
        fstream input("myfile", ios::in | ios::binary);
        Person person;
        person.ParseFromIstream(&input);
        cout << "Name: " << person.name() << endl;
        cout << "E-mail: " << person.email() << endl;
        ```

* you can add new fields to your message formats **w/o breaking backwards compatibility**
  * old binaries simplify ignore the new field when parsing
* thus, if you have a communication protocol that uses protocol buffers as its data format,
  * you can extend your protocol without having worry about breaking existing code

* using generated protocol buffer code: [LINK](https://developers.google.com/protocol-buffers/docs/reference/overview)
* how protocol buffer messages are encoded: [LINK](https://developers.google.com/protocol-buffers/docs/encoding)

<!-- ## Why not just use XML? -->