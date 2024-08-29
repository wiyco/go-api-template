# go-api-template

[Gin](https://github.com/gin-gonic/gin) starter repository.

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

Run with Docker to develop with live-reloading (hot-reload).

```bash
docker compose up
```

API route is `localhost:8080` as the root.

## Modules

1. [air](#1-air)
2. [godotenv](#2-godotenv)
3. [gorm](#3-gorm)
4. [uuid](#4-uuid)

### 1. [air](https://github.com/air-verse/air)

Used for live-reloading (hot-reload) development.

### 2. [godotenv](https://github.com/joho/godotenv)

Environment variables library such as dotenv (e.g. `.env`).

### 3. [gorm](https://gorm.io/docs)

ORM library.

### 4. [uuid](https://github.com/google/uuid)

Used to generate UUIDs.
