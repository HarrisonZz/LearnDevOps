FROM ubuntu:jammy
ARG PY_VERSION=3.10

RUN apt-get update && \
    apt-get install -y \
    curl \
    ca-certificates \
    software-properties-common \
    build-essential && \
    echo "deb [trusted=yes] https://ppa.ipinfo.net/ /" | tee "/etc/apt/sources.list.d/ipinfo.ppa.list" && \
    add-apt-repository -y ppa:deadsnakes/ppa && \
    apt update && \
    apt-get install -y ipinfo python${PY_VERSION}

ENTRYPOINT ["bash","-c"]
CMD ["echo 'Hello World' && /bin/bash"]
