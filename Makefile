vendor:
	go mod vendor

local.pb.go: vendor
	protoc -I vendor/ -I protos/ protos/local.proto --go_out=plugins=grpc:.
