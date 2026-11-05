# robmeijerink-go

Personal website platform made in Go with Fiber v3. One Go app serves three sites, chosen by the `Host` header:

| Site | Folder | Production |
| --- | --- | --- |
| robmeijerink.nl | `sites/robmeijerink` | robmeijerink.nl |
| Solvalutions (EN) | `sites/solvalutions_com` | solvalutions.com |
| Solvalutions (NL) | `sites/solvalutions_nl` | solvalutions.nl |

#### Development

After cloning this repository, simply run the Taskfile.yaml:

```sh
task        # Searches through project commands
```

#### Local URLs

Add these to `/etc/hosts` once (they are what the dev server expects):

```
127.0.0.1   robmeijerink.test
127.0.0.1   solvalutions.com.test
127.0.0.1   solvalutions.nl.test
```

| Task | Open in the browser |
| --- | --- |
| `task rob:dev` | http://robmeijerink.test:3000 (or http://localhost:3000) |
| `task solcom:dev` | http://solvalutions.com.test:3000 |
| `task solnl:dev` | http://solvalutions.nl.test:3000 |

- **`http://localhost:3000` always shows robmeijerink.nl**, whichever site you started. For solvalutions.nl and solvalutions.com use the `.test` names above.
- Each task starts the Go app (Air hot reload) on port **3000** and the Vite dev server for that site on port **5173**. Only one site runs at a time, because both ports are shared.
- In dev, HTML templates reload on save. Go code (including the FAQ text in `sites/solvalutions_nl/landings.go`) is rebuilt by Air.

#### Checking a production build locally

Build a site's frontend, then run the app in production mode on another port:

```sh
# frontend (uses Docker, because host npm and the Linux node_modules clash)
docker run --rm -v .:/app -w /app/sites/solvalutions_nl/frontend node:24-alpine sh -c "npm install && npm run build"

# backend
PORT=3600 APP_ENV=production go run .

# pick the site with a Host header
curl -s -H "Host: solvalutions.nl" http://127.0.0.1:3600/website-laten-maken
```

Production mode serves the hashed files from `public/assets/dist`, adds the security headers (CSP, HSTS) and rate limits at 100 requests per minute per IP, so avoid hammering it with many page loads in a row (you get a 429).

#### Tests and checks

```sh
go vet ./... && go test ./...
go run golang.org/x/vuln/cmd/govulncheck@latest ./...    # Go dependencies
docker run --rm -v .:/app -w /app/sites/solvalutions_nl/frontend node:24-alpine npm audit
```

#### Troubleshooting

- **A page shows old content or no FAQ in dev**: Air may still be running an old binary (`failed to listen: address already in use` in the logs). Run `docker restart robmeijerink-go-backend-1`.
- **Dev server not reachable**: check that the containers are running with `docker ps`, and that you use the `.test` URL.
- **`task` build fails on macOS with a native binding error**: build in Docker as shown above, never with host npm.

#### Deploying

Merging to `main` runs the tests and deploys to Google App Engine (`.github/workflows/deploy.yml`, `app.yaml`). Pushing any other branch does not deploy.
