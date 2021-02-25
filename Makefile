protoc-std:
	@docker run --rm -v $$(pwd):/work -w /work thethingsindustries/protoc:3.1.32 --go_out=./proto_out --proto_path=. spec.proto

protoc-faster:
	@docker run --rm -v $$(pwd):/work -w /work thethingsindustries/protoc:3.1.32 --gogofaster_out=plugins=grpc:./proto_out --gogofaster_opt=paths=source_relative --proto_path=. spec.proto