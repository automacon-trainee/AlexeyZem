#gen proto
#first - generate server
#second - generate client
#third - where file.proto
proto_gen:
	protoc \
 	--go_out=internal/API --go_opt=paths=import --go_opt=module=metrics \
 	--go-grpc_out=internal/API --go-grpc_opt=paths=import --go-grpc_opt=module=metrics \
  	proto/*.proto

#start all test in project
tests:
	go test ./... -v