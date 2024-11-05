FROM ubuntu:jammy
ARG PY_VERSION=3.10
ARG ACTIVATE_PYVENV=py_venv/bin/activate
COPY app.py /src/
WORKDIR /src/
RUN apt-get update && \
    apt-get install -y \
    curl \
    ca-certificates \
    software-properties-common \
    build-essential && \
    echo "deb [trusted=yes] https://ppa.ipinfo.net/ /" | tee "/etc/apt/sources.list.d/ipinfo.ppa.list" && \
    add-apt-repository -y ppa:deadsnakes/ppa && \
    apt update && \
    apt-get install -y ipinfo python${PY_VERSION} python3-pip \
    python3.10-venv && \
    mkdir /flask_project && cd /flask_project && \
    python3 -m venv py_venv && \
    chmod +x ${ACTIVATE_PYVENV} && \
    bash -c "${ACTIVATE_PYVENV}" && \
    pip3 install Flask
ENTRYPOINT ["flask","run"]
