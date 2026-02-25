---
name: ardoqcli
description: Operates Ardoq enterprise architecture platform via the ardoqcli command-line tool. Use when user asks to list, get, create, update, or delete components, references, workspaces, or reports in Ardoq. Also use when user says "query Ardoq", "architecture data", "component list", or "batch update".
license: MIT
metadata:
  author: Samyukti
  version: 0.1.0
---

# ardoqcli

A CLI tool wrapping the Ardoq REST API v2. All output is JSON to stdout, status messages go to stderr.

## Prerequisites

- `ardoqcli` installed and configured (`ardoqcli configure` or env vars `ARDOQ_BASE_URL`, `ARDOQ_API_KEY`)
- Verify connection with `ardoqcli me`

## Commands

### Read Operations

```bash
ardoqcli me                                       # Test connection
ardoqcli workspace list                           # List workspaces
ardoqcli workspace get <id>                       # Get workspace details
ardoqcli workspace context <id>                   # Get workspace context
ardoqcli component list                           # List all components
ardoqcli component list -q rootWorkspace=<id>     # Filter by workspace
ardoqcli component get <id>                       # Get single component
ardoqcli reference list                           # List references
ardoqcli reference get <id>                       # Get single reference
ardoqcli report list                              # List reports
ardoqcli report get <id>                          # Get report details
ardoqcli report run <id>                          # Run a report
ardoqcli report run <id> -t tabular               # Run report as tabular
```

### Write Operations

```bash
ardoqcli component create -d '{"name":"X","rootWorkspace":"id","typeId":"id"}'
ardoqcli component create -f data.csv -t csv      # Create from CSV
ardoqcli component update <id> -d '{"name":"Y"}'
ardoqcli component delete <id>
ardoqcli reference create -d '{"source":"id1","target":"id2","type":1}'
ardoqcli reference update <id> -f changes.json
ardoqcli reference delete <id>
ardoqcli batch -f batch.json                      # Bulk operations
```

### Global Flags

- `-o <file>` — Write JSON output to file instead of stdout
- `-q key=val,key=val` — Query parameters for filtering list endpoints
- `--config <path>` — Custom config file path

## Input Formats

- `-d` — Inline JSON string
- `-f` — File path (JSON by default)
- `-t csv` — Interpret file as CSV (headers become JSON keys, single row produces object, multiple rows produce array)

## Processing Output

Output is JSON, designed for piping. Use these tools:

### jq (recommended)

```bash
# Count components
ardoqcli component list | jq '.values | length'

# Filter by type
ardoqcli component list | jq '[.values[] | select(.type == "Application")]'

# Extract names
ardoqcli component list | jq '[.values[] | select(.type == "Application") | .name]'

# Flatten custom fields
ardoqcli component list | jq '[.values[] | {name} + .customFields]'

# Pick specific custom fields
ardoqcli component list | jq '[.values[] | {name, review_date: .customFields.review_date, hosting_type: .customFields.hosting_type}]'
```

### dsq (SQL queries on JSON)

```bash
# Query with SQL (use {"values"} to access nested array)
ardoqcli component list | dsq -s json 'SELECT name FROM {"values"} WHERE type = "Application"'

# Access nested fields with quoted dot notation
ardoqcli component list | dsq -s json 'SELECT name, "customFields.hosting_type" as hosting_type FROM {"values"} WHERE type = "Application"'
```

## API Reference

For full request/response schemas, consult `references/openapi.json` or the [Ardoq API docs](https://developer.ardoq.com/public-api/).

## Important Notes

- API responses wrap arrays in a `values` key
- PATCH requests automatically use `ifVersionMatch=latest`
- The `-q` flag maps to API query parameters (supported: name, rootWorkspace, parent, componentKey, typeId for components)
- The `type` field in responses (e.g., "Application") is not a filterable API query param; filter client-side with jq or dsq
- If config is missing, the CLI shows which values are needed and how to set them
