# Section 5. Setting up Protoc Compiler

---

## 21. Setup Protoc Compiler

### Setup Protoc Compiler

In order to perform code generation, you will need to install `protoc` on your computer.

============ MacOSX =============

It is actually very easy, open a command line interface and type `brew install protobuf`

---

## 22. Use `protoc` to generate code in any language

Want to see all documents for `protoc` then just do this command in your bash:

```bash
protoc
```

and it returns

```bash
Usage: protoc [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
                              If not found in any of the these directories,
                              the --descriptor_set_in descriptors will be
                              checked for required proto file.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
                              standard input and write the raw tag/value
                              pairs in text format to standard output.  No
                              PROTO_FILES should be given when using this
                              flag.
  --descriptor_set_in=FILES   Specifies a delimited list of FILES
                              each containing a FileDescriptorSet (a
                              protocol buffer defined in descriptor.proto).
                              The FileDescriptor for each of the PROTO_FILES
                              provided will be loaded from these
                              FileDescriptorSets. If a FileDescriptor
                              appears multiple times, the first occurrence
                              will be used.
...
```

make directory for Java and Python

```bash
mkdir java python
```

and

```bash
protoc -I=proto --python_out=python proto/*.proto
```

this means we want to generate files for Python from `*.proto` files that located under a directory located in `(pwd)/proto`. So if we run the command, `protoc` generate `*.py` file from `*.proto` files.

if you want to do same thing for Java:

```bash
protoc -I=proto --java_out=java proto/*.proto
```

we can do same things for these languages (check this out by typing `protoc` in your bash)

```bash
--cpp_out=OUT_DIR           Generate C++ header and source.
--csharp_out=OUT_DIR        Generate C# source file.
--java_out=OUT_DIR          Generate Java source file.
--js_out=OUT_DIR            Generate JavaScript source.
--objc_out=OUT_DIR          Generate Objective C header and source.
--php_out=OUT_DIR           Generate PHP source file.
--python_out=OUT_DIR        Generate Python source file.
--ruby_out=OUT_DIR          Generate Ruby source file.
```

---

## 23. Practice Using `protoc`

### Practice Using `protoc`

This was a short video, and the best way to get a feel for `protoc` is to practice!

So go ahead! As an exercise, use `protoc` to create javascript code (`--js_out`  option) on the files we have created in the code folder. 

Happy learning :)

### Actual Trial

We have pre-written files with this structure

`/proto/complex.proto`  
`/proto/enum_example.proto`  
`/proto/simple.proto`  

and run this command:

```bash
protoc -I=proto --js_out=outputs/javascript proto/*.proto
```

therefore, we can get `*.javascript` file from these `*.proto` file under the directory `outputs/javascript`. If you want to get another codes for different languages, for example, python and java, you can also do this:

```bash
protoc -I=proto --java_out=outputs/java proto/*.proto
protoc -I=proto --python_out=outputs/python proto/*.proto
```
