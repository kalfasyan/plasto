# CLAUDE.md

plasto is a template for self-hosted apps: a Go backend on PocketBase (`backend/`)
serving a Svelte 5 + Tailwind web app (`web/`). Decisions and the build plan are in
`docs/decisions.md`.

## Commands

- `make dev`: run the backend (:8090) and the web app (:5173) with live reload
- `make check` and `make test` must pass before every commit

## Conventions

- Database collections are defined in Go migrations under `backend/migrations`
  (`go run . migrate create <name>`), never only in the admin UI.
- Custom API routes live under `/api/app/`.
- Keep it simple: no new dependency or service without a clear need.

## Git and GitHub

- Never add AI co-author or attribution lines to commits, PR descriptions, or GitHub comments
  (no `Co-Authored-By: Claude ...`, `Claude-Session: ...`, "Generated with Claude Code", or similar).
