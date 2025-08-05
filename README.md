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

Minikube

alias mkubectl='minikube kubectl --'

ERROR: error during connect: Head "https://127.0.0.1:53716/_ping": tls: failed to verify certificate: x509: certificate has expired or is not yet valid: current time 2025-08-01T12:43:11+01:00 is after 2025-08-01T11:05:00Z

minikube delete
minikube start

🛑 Issue 1: PersistentVolumeClaim "consumer-pvc" is invalid: spec.accessModes: Required value
Cause: your PVC template didn’t specify any accessModes.

Fix: added .Values.volume.accessMode (ReadWriteOnce) into the PVC spec.

🛑 Issue 2: no persistent volumes available for this claim and no storage class is set
Cause: you explicitly set storageClassName: "" in the PVC, which disables dynamic provisioning.

Kubernetes then looked for a static PV to bind but none existed.

Fix: either remove storageClassName to use the cluster’s default dynamic StorageClass, or create a matching static PV.

You chose to remove it → letting Minikube’s default StorageClass provision a PV dynamically.

🛑 Issue 3: spec is immutable after creation except resources.requests...
Cause: Helm upgrade tried to patch the existing PVC, changing storageClassName: "" to nil.

PVC spec.storageClassName (and most of the spec) is immutable once created.

Fix: had to delete the old PVC (mkubectl delete pvc consumer-pvc) before re-deploying so Helm could create it fresh.

✅ Final State
PVC now has accessModes: [ReadWriteOnce] and no explicit storageClassName, so it successfully bound to a dynamically provisioned PV.

The consumer job can mount /names (or whatever you’ve set) and persist cache data across pod restarts.

⚡ Lesson learned:
With PVCs:

Always specify accessModes.

Use the default StorageClass unless you have a reason not to.

PVC specs are mostly immutable — if you need to change the class or binding, you must delete and recreate the PVC.