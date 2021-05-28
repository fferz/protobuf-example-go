# protobuf-example-go

1. Install protobuf (on Mac)

```brew install protobuf```

2. Create a folder and set up the go dependency manager by writing:

```go mod init github.com/fferz/protobuf-example-go```

This will create a go.mod file with all the dependencies.

3. Create a main.go
4. Create a folder with a proto file: src/simple/simple.proto.  

   Add a go option inside the proto file:  
   ```option go_package = "simple;simplepb" // package path ; package name```

5. Generate the Go code using protoc:  

`protoc -I src/ —go_out=src/ src/simple/simple.proto`

6. Download the new dependencies (from the simple.pb.go file

`go mod tidy`
