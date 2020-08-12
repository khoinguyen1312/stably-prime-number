all: bin/example
test: lint unit-test

PLATFORM=local
TARGET_IMAGE_NAME=main:latest

.PHONY: bin/example
bin/example:
	@docker build -f Dockerfile.build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}
	@docker build -f Dockerfile.serve . -t ${TARGET_IMAGE_NAME} --platform ${PLATFORM}

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: lint
lint:
	@docker build . --target lint
