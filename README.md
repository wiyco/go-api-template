# go-api-template

[Gin](https://github.com/gin-gonic/gin) starter repository 🐢

## Tech Stack

| Environments   | Languages  | Linters                                   | Frameworks  | Testing    | CI/CD                 |
| :------------- | :--------- | :---------------------------------------- | :---------- | :--------- | :-------------------- |
| ![docker-logo] | ![go-logo] | ![golangci-lint-logo]<br>![lefthook-logo] | ![gin-logo] | ![go-logo] | ![githubactions-logo] |

[docker-logo]: https://img.shields.io/badge/-Docker-2496ED.svg?logo=nodedotjs&style=flat&logoColor=ffffff
[golangci-lint-logo]: https://img.shields.io/badge/-golangci--lint-00ADD8.svg?logo=nodedotjs&style=flat&logoColor=ffffff
[lefthook-logo]: https://img.shields.io/badge/-Lefthook-FF1E1E.svg?logo=lefthook&style=flat&logoColor=000000
[go-logo]: https://img.shields.io/badge/-Go-00ADD8.svg?logo=go&style=flat&logoColor=ffffff
[gin-logo]: https://img.shields.io/badge/-Gin-646CFF.svg?logo=gin&style=flat&logoColor=ffffff
[githubactions-logo]: https://img.shields.io/badge/-GitHub%20Actions-2088FF.svg?logo=githubactions&style=flat&logoColor=ffffff

### Other

- [air](https://github.com/air-verse/air)
- [godotenv](https://github.com/joho/godotenv)
- [gorm](https://gorm.io/docs)
- [uuid](https://github.com/google/uuid)

## Development

1. [Linter](#linter)
2. [Running the app](#running-the-app)

### Linter

First, you need to install [`golangci-lint`](https://golangci-lint.run).

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
> For _Windows_, use [Git Bash](https://git-scm.com).

### Running the app

Run with Docker to develop with live-reloading (hot-reload).

```bash
docker compose up
```

API route is `localhost:8080` as the root.
