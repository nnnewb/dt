//go:generate protoc -I. -I../../3rd/protoc/include -I../../3rd/googleapis --go_opt paths=source_relative --go_out . --go-grpc_opt paths=source_relative --go-grpc_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true --grpc-gateway_out . wallet.proto shop.proto
package pb
