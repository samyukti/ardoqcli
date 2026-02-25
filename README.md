# ardoqcli

A command-line tool for the [Ardoq](https://www.ardoq.com/) REST API v2.

Thin JSON-passthrough wrapper — pipe output to `jq`, script with shell, automate your enterprise architecture workflows.

## Install

```sh
brew install samyukti/tap/ardoqcli
```

Or with Go:

```sh
go install com.samyukti.ardoqcli@latest
```

Or build from source:

```sh
go build -o ardoqcli .
```

## Configuration

Interactive setup:

```sh
ardoqcli configure
```

This saves credentials to `~/.config/ardoqcli/hosts.yml`.

Alternatively, use environment variables:

```sh
export ARDOQ_BASE_URL=https://myorg.ardoq.com
export ARDOQ_API_KEY=your-api-key
```

Environment variables take precedence over the config file. Use `--config` to specify an alternate config file path.

## Usage

### Test connection

```sh
ardoqcli me
```

### Components

```sh
ardoqcli component list
ardoqcli component list -q rootWorkspace=abc123
ardoqcli component get <id>
ardoqcli component create -d '{"name":"My Component","rootWorkspace":"workspace","typeId":"xyz"}'
ardoqcli component create -f component.csv -t csv
ardoqcli component update <id> -d '{"name":"Updated Name"}'
ardoqcli component delete <id>
```

### References

```sh
ardoqcli reference list
ardoqcli reference get <id>
ardoqcli reference create -d '{"source":"id1","target":"id2","type":1}'
ardoqcli reference update <id> -f changes.json
ardoqcli reference delete <id>
```

### Workspaces

```sh
ardoqcli workspace list
ardoqcli workspace get <id>
ardoqcli workspace context <id>
```

### Reports

```sh
ardoqcli report list
ardoqcli report get <id>
ardoqcli report run <id>
ardoqcli report run <id> -t tabular
```

### Batch operations

```sh
ardoqcli batch -f batch.json
```

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Write JSON output to a file |
| `--query` | `-q` | Query parameters as `key=val,key=val` |
| `--config` | | Path to config file |

## Input formats

Write commands (`create`, `update`, `batch`) accept input via:

- `-d` — inline JSON string
- `-f` — path to a file (JSON by default)
- `-t csv` — interpret the file as CSV (headers become JSON keys)

A single-row CSV produces a JSON object; multiple rows produce an array.

## Design

- **No Go structs for API entities** — raw JSON passthrough keeps the CLI resilient to API changes
- **PATCH uses `ifVersionMatch=latest`** — no need to pre-fetch the current version
- **Output to stdout, messages to stderr** — JSON is always on stdout for clean piping; status and errors go to stderr with colored output
