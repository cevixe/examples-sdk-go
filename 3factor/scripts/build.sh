#!/usr/bin/env bash

sam build -t components/store/template.yaml -b .aws-sam/build/store && \
sam build -t components/api/template.yaml -b .aws-sam/build/api && \
sam build -t components/core/template.yaml -b .aws-sam/build/core && \
sam build -t components/resolvers/template.yaml -b .aws-sam/build/resolvers && \
sam build -t components/functions/template.yaml -b .aws-sam/build/functions && \
sam build -t template.yaml -b .aws-sam/build/root
