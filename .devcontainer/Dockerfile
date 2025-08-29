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

# Project Setup
WORKDIR /app

# Step 1: Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Step 2: Download all dependencies with a cache mount
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
