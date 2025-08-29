FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Basic tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
  && rm -rf /var/lib/apt/lists/*

# Install Go
RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-go \
  && rm -rf /var/lib/apt/lists/*
