.PHONY: proto run

proto:
	for f in internal/proto/*/*.proto; do \
		protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative $$f; \
		echo compiled: $$f; \
	done

run:
	docker-compose build
	docker-compose up --remove-orphans
