PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/orders.proto
OUT_DIR=common/genproto/orders

run-kitchen:
	go run kitchen-service/*.go

run-order:
	go run order-service/*.go
proto:
	mkdir -p $(OUT_DIR)
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
