# Affine stack

Last updated: 2024-04-13 19:40:57

Status: :file_cabinet: Obsolete

## Project dedicated environment variables[^1]

| Name | Default value | Valid example | Description |
| ---- | ------------- | ------------- | ----------- |
| `AFFINE_DB_PASS` |  | `long_and_unguessable_password` | Password for the database |
| `AFFINE_ADM_MAIL` |  | `mail@example.com` | Email address for the admin account |
| `AFFINE_ADM_PASS` |  | `long_and_unguessable_password` | Password for the admin account |
| `AFFINE_MAIL_NAME` |  | `"AFFiNE <affine@example.com>"` | Value in the emails' `from` field |

## Notable services

### <img alt="" src="affine.png" height="32px"> Affine

Category: Knowledge-base

There can be more than Notion and Miro. AFFiNE(pronounced [ə‘fain]) is a next-gen knowledge base that brings planning, sorting and creating all together. Privacy first, open-source, customizable and ready to use.

#### Links for Affine

* [Webpage](https://affine.pro/)
* [Documentation](https://docs.affine.pro/docs/)
* [Repository](https://github.com/toeverything/AFFiNE)
* [GitHub Container Registry](ghcr.io/toeverything/affine-graphql)
* [`docker-compose.yml` example](https://raw.githubusercontent.com/toeverything/AFFiNE/stable/.github/deployment/self-host/compose.yaml)

#### Affine dedicated environment variables[^1]

No dedicated environment variables available.

#### Example [dashy `section.item` entry](https://dashy.to/docs/configuring/#sectionitem) for Affine

```yaml
- title: Affine
  description: Knowledge-base
  icon: >-
    https://cdn.jsdelivr.net/gh/walkxcode/dashboard-icons/png/affine.png
  url: https://affine.example.com/
  tags:
    - electron
    - editor
    - markdown
    - rust
    - app
    - wiki
    - notes
    - table
    - rust-language
    - workspace
    - whiteboard
    - tableview
    - rust-lang
    - crdt
    - knowledge-base
    - notion
    - miro
    - notion-alternative
  color: '#1E96EB'
```

[^1]: Besides and/or instead of those available thanks to used image.
