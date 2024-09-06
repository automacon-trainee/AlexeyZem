#gen proto
#first - generate server
#second - generate client
#third - where file.proto
proto_gen:
	protoc \
 	--go_out=cmd/gRPCGeo --go_opt=paths=import --go_opt=module=metrics/cmd/gRPCGeo \
 	--go-grpc_out=cmd/gRPCGeo --go-grpc_opt=paths=import --go-grpc_opt=module=metrics/cmd/gRPCGeo \
  	cmd/gRPCGeo/*.proto

#start all test in project
tests:
	go test ./... -v