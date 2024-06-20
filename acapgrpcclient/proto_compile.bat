protoc --go_out=./cctv --go_opt=paths=source_relative --go-grpc_out=./cctv --go-grpc_opt=paths=source_relative cctv.proto
python -m grpc_tools.protoc -I. --python_out=.. --grpc_python_out=.. cctv.proto
