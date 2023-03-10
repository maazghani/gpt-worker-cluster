#!/bin/bash

# Define the task
task='{"text": "Write Terraform code"}'

# Publish the task to the NATS queue
curl -X POST -H "Content-Type: application/json" \
     --data "$task" \
     nats://localhost:4222/publish/tasks

