# Actual Budget stack

Last updated: 2024-05-14 18:53:59

Status: :heavy_check_mark: Working

## Project dedicated environment variables[^1]

No dedicated environment variables available.

## Notable services

### <img alt="" src="actual.png" height="32px"> Actual Budget

Category: Budget manager

A super fast and privacy-focused app for managing your finances. At its heart is the well proven and much loved Envelope Budgeting methodology. You own your data and can do whatever you want with it. Featuring multi-device sync, optional end-to-end encryption and so much more.

#### Links for Actual Budget

* [Webpage](https://actualbudget.org/)
* [Repository](https://github.com/actualbudget/actual)
* [Docker Hub](https://hub.docker.com/r/actualbudget/actual-server)
* [`docker-compose.yml` example](https://hub.docker.com/r/actualbudget/actual-server)

#### Actual Budget dedicated environment variables[^1]

No dedicated environment variables available.

#### Example [dashy `section.item` entry](https://dashy.to/docs/configuring/#sectionitem) for Actual Budget

```yaml
- title: Actual Budget
  description: Budget manager
  icon: >-
    https://cdn.jsdelivr.net/gh/walkxcode/dashboard-icons/png/actual.png
  url: https://actual.example.com/
  tags:
    - accounting
    - accounts
    - balance
    - banking
    - budgeting
    - cash
    - envelope
    - expenses
    - finances
    - flow
    - management
    - money
    - personal
    - reports
    - saving
    - spending
    - transactions
    - transfers
    - worth
    - zero-based
  color: '#8444e9'
```

[^1]: Besides and/or instead of those available thanks to used image.
