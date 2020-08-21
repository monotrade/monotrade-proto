# go get -v  github.com/golang/protobuf/protoc-gen-go

mkdir gen-go
protoc --go_out=./gen-go src/demo.proto
