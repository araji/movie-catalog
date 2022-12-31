IMAGE ?= "araji/movie-catalog"
TAG ?= "0.0.2"

.PHONY: docker-push docker-build deploy
docker-push:
	docker push $(IMAGE):$(TAG)

docker-build:
	docker build -t $(IMAGE):$(TAG) .

deploy:
	cd helm && helm upgrade --install movie-catalog . --set image.tag=$(TAG) --set postgres.password=$$(kubectl get secret movie-db-cluster-app -o jsonpath="{.data.password}" | base64 --decode)