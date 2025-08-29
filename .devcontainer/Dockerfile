FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Basic tools
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
  && rm -rf /var/lib/apt/lists/*

# Install Go
RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-go \
  && rm -rf /var/lib/apt/lists/*

# Dependencies for GUI applications
RUN apt-get update && apt-get install -y --no-install-recommends \
    libx11-dev \
    libxcursor-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxi-dev \
    libgl1-mesa-dev \
    libxxf86vm-dev \
    xvfb xauth \
  && rm -rf /var/lib/apt/lists/*

# Use a shared, writable module cache path and pre-create it
ARG USERNAME=vscode
RUN mkdir -p /go/pkg/mod/cache/download && chown -R ${USERNAME}:${USERNAME} /go

# Project Setup
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
