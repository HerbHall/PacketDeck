# syntax=docker/dockerfile:1

# PacketDeck Docker Desktop Extension
# Multi-stage build: backend + frontend

FROM alpine
LABEL org.opencontainers.image.title="PacketDeck" \
      org.opencontainers.image.description="Visualize and inspect container network traffic inside Docker Desktop" \
      org.opencontainers.image.vendor="Herb Hall" \
      com.docker.desktop.extension.api.version=">= 0.3.3" \
      com.docker.desktop.extension.icon="https://raw.githubusercontent.com/HerbHall/PacketDeck/main/docker.svg" \
      com.docker.extension.screenshots="" \
      com.docker.extension.detailed-description="PacketDeck shows your container network topology, monitors active connections, and captures traffic — all without leaving Docker Desktop." \
      com.docker.extension.publisher-url="https://github.com/HerbHall" \
      com.docker.extension.changelog=""

# TODO: Add backend build stage (with tcpdump/netshoot tools)
# TODO: Add frontend build stage
# TODO: Copy metadata.json and ui assets

COPY metadata.json .
