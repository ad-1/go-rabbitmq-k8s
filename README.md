📦 go-rabbitmq-k8s – Message Queue App with Go, RabbitMQ, and Kubernetes
This project is a hands-on DevOps exercise demonstrating how to design, containerize, and deploy a simple message queue-based system built with Go and RabbitMQ, orchestrated entirely within a local Minikube Kubernetes cluster.

🧩 Components
Producer: A Go job that connects to RabbitMQ and publishes 10 messages.

Consumer: A Go job (formerly a Deployment) that consumes a fixed number of messages with concurrent workers using goroutines, channels, and WaitGroups.

RabbitMQ: Deployed via the Bitnami Helm chart into the same namespace.

Kubernetes Jobs: Used for both producer and consumer for clean, finite processing.

Helm Charts: One for each component, configured with templated environment variables.

Makefile: A local DevOps command center — build Docker images, deploy charts, tail logs, and clean up the cluster.

💡 Project Purpose
This isn't intended to be a polished tutorial. It’s a working technical spike designed to demonstrate deep integration across technologies and capture the real-world workflow of building and debugging an end-to-end message queue system.

The eventual blog series will focus not just on the how-to, but the “why isn't this working?” moments — covering key pain points and lessons learned from debugging issues like:

Local Docker image not being found inside Minikube

Job vs. Deployment semantics for finite workers

RabbitMQ queues showing 0 messages even after publishing

Image pull failures due to missing imagePullPolicy

Secrets management and password sync across services

⚙️ Tech Stack Highlights
Go: Lightweight message publishing/consuming with goroutines

RabbitMQ: AMQP broker running in the cluster

Docker: Multi-container setup with separate Dockerfiles

Helm: Parameterized deployment of Jobs with Secrets and ConfigMaps

Kubernetes: All components run as Jobs inside Minikube

Makefile: Everything orchestrated with a single command: make all

🚧 Next Steps
Convert environment variables from ConfigMaps to Kubernetes Secrets

Store RabbitMQ credentials securely and share across services

Write up a developer-focused blog series based on the questions, errors, and debugging process rather than a standard tutorial