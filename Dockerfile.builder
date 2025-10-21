FROM goreleaser/goreleaser-cross:latest

RUN apt-get update && apt-get install -y \
    xorg-dev \
    mesa-utils \
    libgl1 \
    libgl1-mesa-dev && \
    rm -rf /var/lib/apt/lists/*
