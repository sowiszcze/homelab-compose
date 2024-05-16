# AdGuard Home Sync stack

Last updated: 2024-05-03 10:38:19

Status: :heavy_check_mark: Working

## Project dedicated environment variables[^1]

| Name | Default value | Valid example | Description |
| ---- | ------------- | ------------- | ----------- |
| `FALLBACK_DNS` |  | `9.9.9.9` | DNS IP used as a fallback in case the main one fails to start |
| `PRIMARY_DNS` |  | `10.0.0.10` | Primary DNS IP, equals to IP address of AdGuard Home service in `br_adguard` network |
| `PRIMARY_DNS_SUBNET` |  | `10.0.0.0/24` | CIDR notation for subnet mask of `br_adguard` network |

## Notable services

### ISC DHCP Relay Agent

Category: Networking

The Internet Systems Consortium DHCP Relay Agent, dhcrelay, provides a means for relaying DHCP and BOOTP requests from a subnet to which no DHCP server is directly connected to one or more DHCP servers on other subnets. It supports both DHCPv4/BOOTP and DHCPv6 protocols.

#### Links for ISC DHCP Relay Agent

* [Webpage](https://linux.die.net/man/8/dhcrelay)
* [Repository](https://github.com/modem7/DHCP-Relay)
* [Docker Hub](https://hub.docker.com/r/modem7/dhcprelay)
* [`docker-compose.yml` example](https://github.com/modem7/DHCP-Relay#configuration)

#### ISC DHCP Relay Agent dedicated environment variables[^1]

No dedicated environment variables available.

### <img alt="" src="adguardhome.png" height="32px"> AdGuard Home

Category: Networking

AdGuard Home is a network-wide software for blocking ads & tracking. After you set it up, it’ll cover ALL your home devices, and you don’t need any client-side software for that. With the rise of Internet-Of-Things and connected devices, it becomes more and more important to be able to control your whole network.

#### Links for AdGuard Home

* [Webpage](https://adguard.com/en/adguard-home/overview.html)
* [Repository](https://github.com/AdguardTeam/AdGuardHome)
* [Docker Hub](https://hub.docker.com/r/adguard/adguardhome)
* [`docker-compose.yml` example](https://github.com/actualbudget/actual-server/blob/master/docker-compose.yml)

#### AdGuard Home dedicated environment variables[^1]

No dedicated environment variables available.

#### Example [dashy `section.item` entry](https://dashy.to/docs/configuring/#sectionitem) for AdGuard Home

```yaml
- title: AdGuard Home
  description: Networking
  icon: >-
    https://cdn.jsdelivr.net/gh/walkxcode/dashboard-icons/png/adguard-home.png
  url: https://adguard-home.example.com/
  color: '#68BC71'
```

### <img alt="" src="adguardhome-sync.png" height="32px"> AdGuardHome Sync

Category: Networking

Synchronize AdGuardHome config to replica instances.

#### Links for AdGuardHome Sync

* [Repository](https://github.com/bakito/adguardhome-sync)
* [GitHub Container Registry](https://ghcr.io/bakito/adguardhome-sync)
* [`docker-compose.yml` example](https://github.com/bakito/adguardhome-sync#docker-compose)

#### AdGuardHome Sync dedicated environment variables[^1]

No dedicated environment variables available.

#### Example [dashy `section.item` entry](https://dashy.to/docs/configuring/#sectionitem) for AdGuardHome Sync

```yaml
- title: AdGuardHome Sync
  description: Networking
  icon: >-
    https://cdn.jsdelivr.net/gh/walkxcode/dashboard-icons/png/adguardhome-sync.png
  url: https://adguardhome-sync.example.com/
  color: '#68BC71'
```

[^1]: Besides and/or instead of those available thanks to used image.
