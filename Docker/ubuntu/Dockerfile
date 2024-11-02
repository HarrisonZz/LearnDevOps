FROM ubuntu:jammy

RUN apt-get update && \
    apt-get install -y \
    curl \
    ca-certificates && \
    echo "deb [trusted=yes] https://ppa.ipinfo.net/ /" | tee "/etc/apt/sources.list.d/ipinfo.ppa.list" && \
    apt-get update && \
    apt-get install -y ipinfo \


