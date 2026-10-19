# plasto

A template for small self-hosted apps. It takes you from an idea to a
phone-friendly web app, an API and a database in one container on your own
hardware.

It is built on [PocketBase](https://pocketbase.io) (Go, SQLite, auth, realtime,
admin UI) and [Svelte 5](https://svelte.dev) with
[Tailwind](https://tailwindcss.com). See [docs/decisions.md](docs/decisions.md)
for why, and for what's coming next.

## Quick start

You need [Go](https://go.dev/dl/) 1.27+, [Node](https://nodejs.org) 22+ and
[uv](https://docs.astral.sh/uv/) (for the Python workers).

```sh
make setup   # install dependencies
make dev     # backend on :8090, web app with live reload on http://localhost:5173
```

Open the app and create your account: the first person becomes the **owner**
and can also sign in to the admin UI at http://localhost:8090/_/ with the same
email and password. Add more people there, in the `users` collection.

## What you get

- **First-run setup and sign-in.** No terminal needed to create the first account.
- **Users and roles.** The owner manages everyone; members can see each other and
  edit themselves. Nobody can sign themselves up.
- **Push notifications** straight from your server, with no account or third-party
  service. Each person turns them on per device; send one from Go with
  `push.ToUser(app, userID, push.Message{Title: "…", URL: "/…"})`. Phones need the
  app on HTTPS.
- **Scheduled jobs that don't get skipped.** `jobs.MustAdd(app, "digest", "0 7 * * *", fn)`
  runs on a cron schedule in your time zone (`TZ`). If the server was off when a job
  was due, it runs once when the server starts again. Every run and its error are
  listed in the admin UI under the `jobs` collection.
- **Nightly backups** of the database and uploads at 03:00, keeping the last 7, in
  `pb_data/backups`. Change the schedule or add S3 storage in the admin UI under
  Settings > Backups.
- **Python workers** for the heavy lifting (ML, scraping, LLMs), on this machine or
  another. They sign in with a worker account and read and write records through
  the API; see [workers/README.md](workers/README.md).
- **A demo** that uses every piece: a live items list on the home screen, a worker
  that adds to it, and a push notification for each new item.
- **An installable app.** It installs to a phone's home screen or a desktop, opens
  offline from its cached shell, and follows the system's light or dark mode.

## Deploy

On any Linux machine with Docker (a NAS, a Raspberry Pi, a server):

```sh
curl -fsSL https://raw.githubusercontent.com/kalfasyan/plasto/main/deploy/install.sh | sh
```

It asks for a Tailscale auth key, so your devices reach the app at
`https://plasto.<your-tailnet>.ts.net`, and starts it with Docker Compose. See
[docs/deploy.md](docs/deploy.md) for updates, backups, workers, Synology and the Pi.

## Layout

```
backend/            Go app on PocketBase
  main.go           entry point
  app/              routes, hooks and jobs (start here)
  jobs/             scheduled jobs with catch-up
  push/             Web Push: keys, subscriptions, sending
  migrations/       database schema as code
web/                Svelte 5 + Tailwind web app
  src/App.svelte    picks setup, sign-in or the app
  src/components/   screens; Home.svelte is where your app starts
  src/lib/          shared code (API client, session, install) and its tests
  public/           manifest, service worker and icons
workers/            Python worker kit, the demo worker, their tests and image
deploy/             compose files, Tailscale sidecar, install script, smoke test
docs/               decisions and the deploy guide
Dockerfile          the app's multi-arch image
Makefile            run `make` to see every command
```

## Commands

| Command      | What it does                                          |
| ------------ | ----------------------------------------------------- |
| `make dev`   | Run everything with live reload                       |
| `make check` | Format, lint and type-check                           |
| `make test`  | Run the Go, web and worker tests                      |
| `make build` | Build `dist/plasto` with the web app next to it       |
| `make run`   | Build the Docker image and run it on :8090            |
| `make smoke` | Deploy both images with Compose and test end to end   |

Pushes to `main` publish `ghcr.io/kalfasyan/plasto` and
`ghcr.io/kalfasyan/plasto-worker` for amd64, arm64 and armv7; a `v*` tag
publishes that version.

## License

MIT
