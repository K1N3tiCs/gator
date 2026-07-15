# Gator — RSS Feed Aggregator CLI

Gator is a command-line RSS feed aggregator. Register accounts, subscribe to RSS feeds, and browse posts from your followed feeds — all from the terminal.

## Prerequisites

- Go 1.26+
- PostgreSQL
- [goose](https://github.com/pressly/goose) (for database migrations)

## Setup

### 1. Configuration

Create `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable"
}
```

### 2. Database

```bash
createdb gator
goose postgres "postgres://user:password@localhost:5432/gator?sslmode=disable" up
```

### 3. Build

```bash
go build -o gator
```

## Usage

### User management

| Command | Description |
|---|---|
| `register <name>` | Create a new user and log in |
| `login <name>` | Switch to an existing user |
| `users` | List all users |

### Feed subscriptions

| Command | Description |
|---|---|
| `addfeed <name> <url>` | Add an RSS feed (auto-follows) |
| `feeds` | List all feeds in the system |
| `follow <url>` | Follow an existing feed |
| `following` | List feeds you follow |
| `unfollow <url>` | Unfollow a feed |

### Aggregation & browsing

| Command | Description |
|---|---|
| `agg <duration>` | Continuously fetch feeds every `<duration>` (e.g. `30s`, `1m`) |
| `browse [limit]` | Show recent posts (default 2) |

### Maintenance

| Command | Description |
|---|---|
| `reset` | Delete all users and data |

## Architecture

```
main.go              — CLI entry point, command dispatch
handler_*.go         — Command handlers (users, feeds, posts)
rss_feed.go          — RSS fetcher and XML parser
commands.go          — Generic command registry
internal/config/     — ~/.gatorconfig.json reader/writer
internal/database/   — sqlc-generated Go data access layer
sql/schema/          — PostgreSQL migrations (goose)
sql/queries/         — Named SQL queries (sqlc input)
```

## Regenerating code

After modifying `sql/queries/` or `sql/schema/`:

```bash
sqlc generate
goose postgres "$DATABASE_URL" up
```
