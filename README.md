# go-api-template

REST API / [Gin](https://github.com/gin-gonic/gin) starter repository.

## Docs

- [Gin](https://gin-gonic.com/docs)

## Development

1. [Linter](#linter)
2. [Running the app](#running-the-app)

### Linter

First, you need to install [`golangci-lint`](https://golangci-lint.run/).

Please run the following script according to your environment:

> [!TIP]
>
> Make sure that `~/go/bin` is exist in `PATH`.

#### macOS

```zsh
brew install golangci-lint
```

#### Windows & Linux

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v1.60.3
```

> [!IMPORTANT]
>
> For Windows, use [Git Bash](https://git-scm.com/).

### Running the app

Run with Docker. Development with live-reloading (hot-reload).

```bash
docker compose up
```

## Modules

1. [air](#1-air)

### 1. [air](https://github.com/air-verse/air)

Used for live-reloading (hot-reload) development.
