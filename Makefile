PHONY: generate
generate:
		mkdir -p pkg/profile_v1
		protoc -I . --go_out=pkg/profile_v1 --go_opt=paths=import \
				--go-grpc_out=pkg/profile_v1 --go-grpc_opt=paths=import \
				 --grpc-gateway_out=pkg/profile_v1 \
                 --grpc-gateway_opt generate_unbound_methods=true \
                  --openapiv2_out . \
				api/profile_v1/service.proto
		mv pkg/profile_v1/rus-profile-test/pkg/profile_v1/* pkg/profile_v1/
		rm -rf pkg/profile_v1/rus-profile-test