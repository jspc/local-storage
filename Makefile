.PHONY: test
test: go-test integration-test

.PHONY: go-test
go-test:
	go test -v

local-storage:
	go build -o $@

.PHONY: run
run:
	./local-storage

.PHONY: integration-test
integration-test:
	docker run --network host jspc/storage-service
