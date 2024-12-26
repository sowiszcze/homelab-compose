# Preparing machine for the deployment

Docker Compose projects included in this repository require some setup before
they can be used. Here are quick first steps to start using this repository.

All setup was done on Linux machines running either Ubuntu 22.04.4 LTS (Jammy
Jellyfish) or Debian GNU/Linux 12 (bookworm), but it should not be that
challenging to translate these for use in other distros or even Windows or Mac
OS.

## Prerequisites

Despite setting up [environment variables](ENVIRONMENT.md), there needs to be
some software installed on destination machine in order to use included
projects.

It is assument that dependencies and requirements for listed software are
already satisfied or will be during installation.

### Required software

- docker-ce
- docker-ce-cli
- docker-ce-rootless-extras
- docker-compose-plugin
- git

```bash
apt install docker-ce docker-ce-cli docker-ce-rootless-extras \
  docker-compose-plugin git
```

### Helpful optional software

- apache2-utils
- lshw
- net-tools
- nnn / mc
- openssl
- pciutils
- powertop
- restic
- systemd-resolved
- tailscale
- tmux
- ufw
- [ufw-docker](https://github.com/chaifeng/ufw-docker)
- usbutils

```bash
apt install apache2-utils lshw net-tools mc openssl pciutils powertop restic \
  systemd-resolved tailscale tmux ufw usbutils
```

### Containerized software

It is assumed you use at least following software in containers:

- [Authentik](authentik) for OIDC, LDAP, and others
- [Traefik](traefik) as reverse proxy and cert manager

## Docker networks

Following set of Docker networks is used:

- ``exposed`` for containers needing a connection to the Internet
- ``mqtt`` for containers using [MQTT service](eclipse-mosquitto)
- ``observability`` for containers exposing prometheus endpoint
- ``proxy`` for containers exposing their services via [reverse proxy](traefik)
- ``s3`` for containers using [S3 compatible object storage](minio)
- ``smarthome`` for IoT and smarthome-related services that need to be connected
- ``smtp`` for containers sending email through [SMTP](docker-mailserver)
- ``wireguard`` for containers making use of [wireguard](wg-easy) network

Setup commands for use in your terminal of choice:

```bash
docker network create --attachable --opt com.docker.network.bridge.name=docker-exposed --opt com.docker.network.bridge.enable_icc=false --gateway=192.168.224.1 --subnet=192.168.224.0/20 --label="tech.sowi.version=beta" exposed;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-mqtt --gateway=192.168.176.1 --subnet=192.168.176.0/20 --label="tech.sowi.version=beta" mqtt;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-observe --gateway=192.168.160.1 --subnet=192.168.160.0/20 --label="tech.sowi.version=beta" observability;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-proxy --gateway=192.168.192.1 --subnet=192.168.192.0/20 --label="tech.sowi.version=beta" proxy;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-s3 --gateway=192.168.240.1 --subnet=192.168.240.0/20 --label="tech.sowi.version=beta" s3;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-smarthome --gateway=192.168.144.1 --subnet=192.168.144.0/20 --label="tech.sowi.version=beta" smarthome;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-smtp --gateway=192.168.208.1 --subnet=192.168.208.0/20 --label="tech.sowi.version=beta" smtp;
docker network create --attachable --internal --opt com.docker.network.bridge.name=docker-wg --gateway=192.168.128.1 --subnet=192.168.128.0/20 --label="tech.sowi.version=beta" wireguard;
```

## `ufw-docker` setup

For installation please refer to the [official instructions](https://github.com/chaifeng/ufw-docker?tab=readme-ov-file#install).

Quick commands list:

```bash
ufw-docker install
ufw-docker allow traefik
ufw-docker allow anonaddy
ufw-docker allow wg-easy
```
