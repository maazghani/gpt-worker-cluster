#!/bin/bash

# Create the Amazon EKS cluster
eksctl create cluster --name gpt3-workers --nodes 3

# Install the necessary tools
kubectl apply -f https://raw.githubusercontent.com/openfaas/faas-netes/master/namespaces.yml

# Install OpenFaaS
helm repo add openfaas https://openfaas.github.io/faas-netes/
helm upgrade openfaas --install openfaas/openfaas --namespace openfaas --set functionNamespace=openfaas-fn

# Install the NATS chart
helm repo add bitnami https://charts.bitnami.com/bitnami
helm upgrade nats bitnami/nats --namespace nats

# Deploy the worker function
kubectl apply -f worker-function.yml

# Create the task queue
kubectl apply -f task-queue.yml

