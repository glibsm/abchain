VERSION ?= 0.0.2
TAG ?= glibsm/abchain:$(VERSION)

.PHONY: docker
docker:
	docker build -t $(TAG) .

.PHONY: push
push: docker
	docker push $(TAG)

