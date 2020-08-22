# go get -v  github.com/golang/protobuf/protoc-gen-go

cd src
protoc --go_out=../../monotrade-platform-go/api/message market.proto
