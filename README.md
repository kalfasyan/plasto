# plasto

A template for small self-hosted apps. It takes you from an idea to a
phone-friendly web app, an API and a database in one container on your own
hardware.

It is built on [PocketBase](https://pocketbase.io) (Go, SQLite, auth, realtime,
admin UI) and [Svelte 5](https://svelte.dev) with
[Tailwind](https://tailwindcss.com). See [docs/decisions.md](docs/decisions.md)
for why, and for what's coming next.

## Quick start

You need [Go](https://go.dev/dl/) 1.27+ and [Node](https://nodejs.org) 22+.

```sh
make setup   # install dependencies
make dev     # backend on :8090, web app with live reload on http://localhost:5173
```

The first start prints a link that creates your admin account. The admin UI is at
http://localhost:8090/_/.

## Layout

```
backend/            Go app on PocketBase
  main.go           entry point
  app/              routes, hooks and jobs (start here)
  migrations/       database schema as code
web/                Svelte 5 + Tailwind web app
  src/App.svelte    the first screen (start here)
  src/lib/          shared code and its tests
Dockerfile          one small multi-arch image
Makefile            run `make` to see every command
```

## Commands

| Command      | What it does                                          |
| ------------ | ----------------------------------------------------- |
| `make dev`   | Run everything with live reload                       |
| `make check` | Format, lint and type-check                           |
| `make test`  | Run the Go and web tests                              |
| `make build` | Build `dist/plasto` with the web app next to it       |
| `make run`   | Build the Docker image and run it on :8090            |

Pushes to `main` publish `ghcr.io/kalfasyan/plasto` for amd64, arm64 and armv7.

## License

MIT
