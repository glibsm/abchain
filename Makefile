VERSION ?= 0.0.3
TAG ?= glibsm/abchain:$(VERSION)

.PHONY: docker
docker:
	docker buildx build --platform linux/amd64,linux/arm64 --push -t $(TAG) .

.PHONY: push
push: docker
	docker push $(TAG)

