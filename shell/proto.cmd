cd ../proto_file
protoc --go_out=plugins=grpc:../src/eassy/proto/  *.proto
