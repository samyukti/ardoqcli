# Processing ardoqcli Output

ardoqcli outputs JSON to stdout, making it easy to pipe into other tools for filtering, querying, and exploring data.

## jq

[jq](https://jqlang.github.io/jq/) is the standard command-line JSON processor.

### Count all components

The API wraps results in a `values` array. Access it and count with `length`.

```sh
ardoqcli component list | jq '.values | length'
```

### Filter by type

`select()` filters objects matching a condition. Wrap in `[...]` to collect results into an array.

```sh
ardoqcli component list | jq '[.values[] | select(.type == "Application")]'
```

### Count filtered results

Chain `length` after filtering to get the count.

```sh
ardoqcli component list | jq '[.values[] | select(.type == "Application")] | length'
```

### Extract specific fields

After filtering, pick the fields you need. Here, just the name.

```sh
ardoqcli component list | jq '[.values[] | select(.type == "Application") | .name]'
```

### Flatten custom fields

`{name} + .customFields` merges the name with all custom fields into a single flat object.

```sh
ardoqcli component list | jq '[.values[] | select(.type == "Application") | {name} + .customFields]'
```

### Select specific custom fields

Pick individual custom fields and rename them at the top level.

```sh
ardoqcli component list | jq '[.values[] | select(.type == "Application") | {name, review_date: .customFields.review_date, hosting_type: .customFields.hosting_type}]'
```

## dsq

[dsq](https://github.com/multiprocessio/dsq) lets you run SQL queries against JSON. It uses SQLite under the hood. Use `{"values"}` to point dsq at the nested array.

Nested fields like `customFields` are flattened with dot notation (e.g., `customFields.review_date`). Quote the dotted names in the SQL query.

### Query with SQL

```sh
ardoqcli component list | dsq -s json 'SELECT name FROM {"values"} WHERE type = "Application"'
```

### Select nested fields

```sh
ardoqcli component list | dsq -s json 'SELECT name, "customFields.review_date" as review_date, "customFields.hosting_type" as hosting_type FROM {"values"} WHERE type = "Application"'
```

## fx

[fx](https://fx.wtf/) is an interactive JSON viewer. Pipe output and browse with keyboard and mouse.

```sh
ardoqcli component list | fx
```
