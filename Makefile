# commands

.PHONY: run run-migrate-down build clean run-build ent-gen ent-init ent-describe

PARSE_ARGS = $(filter-out $@,$(MAKECMDGOALS))

run:
	go run cmd/college/main.go

run-migrate-down:
	go run cmd/college/main.go -migrate-down

build: clean
	go build -o dist/college ./cmd/college

run-build: build
	./dist/college

clean:
	rm -rf dist

ent-init:
	go run entgo.io/ent/cmd/ent init $(PARSE_ARGS)

ent-gen:
	go generate ./ent

ent-describe:
	go run entgo.io/ent/cmd/ent describe ./ent/schema

log:
	@echo $(PARSE_ARGS)

%:
	@: