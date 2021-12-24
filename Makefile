# commands

.PHONY: run build clean run-build ent-gen ent-init ent-describe

PARSE_ARGS = $(filter-out $@,$(MAKECMDGOALS))

run:
	go run main.go

build: clean
	go build -o dist/ent-demo

run-build: build
	./dist/ent-demo

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