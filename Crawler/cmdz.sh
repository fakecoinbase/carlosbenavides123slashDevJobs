cd protos
protoc -I=. --python_out=. ./job.proto
