vendor:
	go mod vendor

local:
	mkdir -p local

local/local.pb.go: vendor local
	protoc -I vendor/ -I protos/ protos/local.proto --go_out=plugins=grpc:local/
