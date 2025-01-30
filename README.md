# go-api-template

REST API / [Fiber](https://github.com/gofiber/fiber) starter repository 🐢

## Tech Stack

| Environments   | Languages  | Linters                                   | Frameworks  | Testing    | CI/CD                 |
| :------------- | :--------- | :---------------------------------------- | :---------- | :--------- | :-------------------- |
| ![docker-logo] | ![go-logo] | ![golangci-lint-logo]<br>![lefthook-logo] | ![fiber-logo] | ![go-logo] | ![githubactions-logo] |

[docker-logo]: https://img.shields.io/badge/-Docker-2496ED.svg?logo=docker&style=flat&logoColor=ffffff
[go-logo]: https://img.shields.io/badge/-Go-00ADD8.svg?logo=go&style=flat&logoColor=ffffff
[golangci-lint-logo]: https://img.shields.io/badge/-golangci--lint-00ADD8.svg?logo=golangcilint&style=flat&logoColor=ffffff
[lefthook-logo]: https://img.shields.io/badge/-Lefthook-FF1E1E.svg?logo=lefthook&style=flat&logoColor=ffffff
<!-- [gin-logo]: https://img.shields.io/badge/-Gin-646CFF.svg?logo=gin&style=flat&logoColor=ffffff -->
[fiber-logo]: https://img.shields.io/badge/-Fiber-00ACD7.svg?logo=gofiber&style=flat&logoColor=ffffff
[githubactions-logo]: https://img.shields.io/badge/-GitHub%20Actions-2088FF.svg?logo=githubactions&style=flat&logoColor=ffffff

### Other

- [air](https://github.com/air-verse/air)

### Gin vs Fiber

These articles are very helpful/informative.

- [Go Gin vs Fiber: Hello World performance](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-hello-world-performance-6863e597b654)
- [Go servers benchmark: Echo, Fiber, and Gin](https://blog.stackademic.com/go-servers-benchmark-echo-fiber-and-gin-caadd9a78319)
- [Go — Gin vs Fiber vs Echo: How much performance difference is really there for a real-world use case?](https://medium.com/deno-the-complete-reference/go-gin-vs-fiber-vs-echo-how-much-performance-difference-is-really-there-for-a-real-world-use-1ed29d6a3e4d)

## Development

1. [Linter](#linter)
2. [Running the app](#running-the-app)

### VS Code

Enable `gopls` to use [IntelliSense features](https://code.visualstudio.com/docs/languages/go).

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
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v1.63.4
```

> [!IMPORTANT]
>
> For _Windows_, use [Git Bash](https://git-scm.com).

### Running the app

Run with Docker to develop with live-reloading (hot-reload).

```bash
docker compose up
```

The API route's root is `localhost:8000`.

You can execute the sample API (`ping`) at `localhost:8000/api/v1/ping`.
