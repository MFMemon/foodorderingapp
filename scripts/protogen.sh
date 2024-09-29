export GO111MODULE='on'
protoc --go_out=internal --go_opt=paths=import --go-grpc_out=internal --go-grpc_opt=paths=import --proto_path=api/protos usersvc/usersvc.proto