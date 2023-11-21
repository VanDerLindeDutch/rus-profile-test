PHONY: generate
generate:
		mkdir -p pkg/profile_v1
		protoc --go_out=pkg/profile_v1 --go_opt=paths=import \
				--go-grpc_out=pkg/profile_v1 --go-grpc_opt=paths=import \
				api/profile_v1/service.proto
		mv pkg/profile_v1/rus-profile-test/pkg/profile_v1/* pkg/profile_v1/
		rm -rf pkg/profile_v1/rus-profile-test