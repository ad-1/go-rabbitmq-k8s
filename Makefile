KUBECTL = minikube kubectl --
NAMESPACE = rabbitmq-demo

# Docker images
CONSUMER_IMAGE = go-rabbitmq-k8s-consumer:latest
PRODUCER_IMAGE = go-rabbitmq-k8s-producer:latest

# Chart paths
CONSUMER_CHART = ./charts/consumer
PRODUCER_CHART = ./charts/producer

.PHONY: all build build-consumer build-producer deploy deploy-consumer deploy-producer restart restart-consumer restart-producer status status-consumer status-producer logs logs-consumer logs-producer clean clean-consumer clean-producer uninstall uninstall-consumer uninstall-producer

# Build images
build: build-consumer build-producer

build-consumer:
	eval $$(minikube docker-env) && docker build -t $(CONSUMER_IMAGE) -f Dockerfile.consumer .

build-producer:
	eval $$(minikube docker-env) && docker build -t $(PRODUCER_IMAGE) -f Dockerfile.producer .

# Deploy charts
deploy: deploy-consumer deploy-producer

deploy-consumer:
	helm upgrade --install rabbitmq-consumer $(CONSUMER_CHART) --namespace $(NAMESPACE) --create-namespace

deploy-producer:
	helm upgrade --install rabbitmq-producer $(PRODUCER_CHART) --namespace $(NAMESPACE) --create-namespace

# Restart deployments and jobs
restart: restart-consumer restart-producer

restart-consumer:
	# For jobs, delete the existing job to restart it
	$(KUBECTL) -n $(NAMESPACE) delete job consumer || true
	helm upgrade --install rabbitmq-consumer $(CONSUMER_CHART) --namespace $(NAMESPACE) --create-namespace

restart-producer:
	# For jobs, delete the existing job to restart it
	$(KUBECTL) -n $(NAMESPACE) delete job producer || true
	helm upgrade --install rabbitmq-producer $(PRODUCER_CHART) --namespace $(NAMESPACE) --create-namespace

# Get rollout/job status
status: status-consumer status-producer

status-consumer:
	$(KUBECTL) -n $(NAMESPACE) get jobs consumer || echo "Consumer job not found"
	$(KUBECTL) -n $(NAMESPACE) get pods -l job-name=consumer

status-producer:
	# Show job status and pods
	$(KUBECTL) -n $(NAMESPACE) get jobs producer || echo "Producer job not found"
	$(KUBECTL) -n $(NAMESPACE) get pods -l job-name=producer

# Tail logs
logs: logs-consumer logs-producer

logs-consumer:
	$(KUBECTL) -n $(NAMESPACE) logs -l app=consumer --tail=50

logs-producer:
	$(KUBECTL) -n $(NAMESPACE) logs -l job-name=producer --tail=50

# Cleanup resources
clean: clean-consumer clean-producer

clean-consumer:
	$(KUBECTL) -n $(NAMESPACE) delete deploy consumer || true
	helm uninstall rabbitmq-consumer --namespace $(NAMESPACE) || true

clean-producer:
	$(KUBECTL) -n $(NAMESPACE) delete job producer || true
	helm uninstall rabbitmq-producer --namespace $(NAMESPACE) || true

# Uninstall all Helm releases
uninstall: uninstall-consumer uninstall-producer

uninstall-consumer:
	helm uninstall rabbitmq-consumer --namespace $(NAMESPACE) || true

uninstall-producer:
	helm uninstall rabbitmq-producer --namespace $(NAMESPACE) || true

# Combined: build, deploy, restart for both consumer and producer
all: clean build deploy restart