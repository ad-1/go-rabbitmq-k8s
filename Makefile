# IMPORTANT: If you are using Minikube, you must set the Docker environment variables in your shell before running any build commands.
# Run the following command in your terminal before using this Makefile:
# eval $(minikube docker-env)
# This ensures Docker builds images inside the Minikube VM, not on your host.

KUBECTL = minikube kubectl --
NAMESPACE = rabbitmq-demo

CONSUMER_IMAGE = go-rabbitmq-k8s-consumer:latest
PRODUCER_IMAGE = go-rabbitmq-k8s-producer:latest

CONSUMER_CHART = ./charts/consumer
PRODUCER_CHART = ./charts/producer

.PHONY: all build build-consumer build-producer deploy deploy-consumer deploy-producer deploy-rabbitmq wait-for-rabbitmq status clean uninstall

build: build-consumer build-producer

build-consumer:
	docker build -t $(CONSUMER_IMAGE) -f Dockerfile.consumer .

build-producer:
	docker build -t $(PRODUCER_IMAGE) -f Dockerfile.producer .

deploy: deploy-rabbitmq wait-for-rabbitmq deploy-consumer deploy-producer

deploy-rabbitmq:
	@if helm status rabbitmq --namespace $(NAMESPACE) >/dev/null 2>&1; then \
		echo "RabbitMQ already exists. Skipping install."; \
	else \
		helm repo add bitnami https://charts.bitnami.com/bitnami; \
		helm repo update; \
		helm upgrade --install rabbitmq bitnami/rabbitmq --namespace $(NAMESPACE) --create-namespace; \
	fi

wait-for-rabbitmq:
	@echo "Waiting for RabbitMQ pod to be ready..."
	@$(KUBECTL) -n $(NAMESPACE) wait --for=condition=ready pod -l app.kubernetes.io/name=rabbitmq --timeout=360s

deploy-consumer:
	pwd
	helm upgrade --install rabbitmq-consumer $(CONSUMER_CHART) \
	  --namespace $(NAMESPACE) --create-namespace \
	  -f ./values-common.yaml -f $(CONSUMER_CHART)/values.yaml

deploy-producer:
	pwd
	helm upgrade --install rabbitmq-producer $(PRODUCER_CHART) \
	  --namespace $(NAMESPACE) --create-namespace \
	  -f ./values-common.yaml -f $(PRODUCER_CHART)/values.yaml

status:
	$(KUBECTL) -n $(NAMESPACE) get jobs,pods

clean:
	$(KUBECTL) -n $(NAMESPACE) delete job consumer producer || true
	-helm uninstall rabbitmq-consumer --namespace $(NAMESPACE)
	-helm uninstall rabbitmq-producer --namespace $(NAMESPACE)

uninstall: clean

all: clean build deploy status
