#!/usr/bin/env bash

./scripts/merge.sh && \
sam build -t modules/api/template.yaml -b .aws-sam/build/api && \
sam build -t modules/services/product/template.yaml -b .aws-sam/build/services/product && \
sam build -t template.yaml -b .aws-sam/build/root
