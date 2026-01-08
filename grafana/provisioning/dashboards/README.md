# Grafana Dashboard Provisioning Guide

## Dashboard Structure

```json
{
  "title": "Dashboard Name",
  "uid": "unique-id",
  "tags": ["prometheus", "monitoring"],
  "templating": { ... },
  "panels": [ ... ]
}
```

## Template Variables (`templating.list`)

### Query Variable (Most Common)

```json
{
  "type": "query",
  "name": "environment", // Use as $environment in queries
  "label": "Environment", // Dropdown label
  "datasource": "Prometheus",
  "query": "label_values(environment)", // PromQL to fetch values
  "multi": false, // Single selection only
  "includeAll": true, // Show "All" option
  "allValue": ".*", // Regex for "All"
  "refresh": 1 // 1=on load, 2=on time change
}
```

### Common Queries

```promql
label_values(environment)                                    # All env labels
label_values(up{environment="prod"}, instance)              # Cascading filter
label_values(http_requests_total, job)                      # Job names
```

### Variable Types

- `query` - Values from datasource
- `custom` - Manual list: `"query": "dev,staging,prod"`
- `interval` - Time intervals: `"query": "1m,5m,15m,1h"`
- `constant` - Fixed value
- `textbox` - Free text input

### Multi-Select

```json
{
  "multi": true,
  "includeAll": true,
  "current": {
    "text": ["instance1", "instance2"],
    "value": ["instance1", "instance2"]
  }
}
```

## Using Variables in Panels

```json
{
  "targets": [
    {
      "expr": "rate(http_requests_total{environment=~\"$environment\", instance=~\"$instance\"}[5m])"
    }
  ]
}
```

- Single value: `environment=\"$environment\"`
- Multi-select: `instance=~\"$instance\"` (regex match)

## Cascading Variables

```json
"list": [
  {"name": "env", "query": "label_values(environment)"},
  {"name": "cluster", "query": "label_values(up{environment=\"$env\"}, cluster)"},
  {"name": "instance", "query": "label_values(up{environment=\"$env\", cluster=\"$cluster\"}, instance)"}
]
```

## Tips

- **No comments in JSON** - Use this README instead
- Dashboard auto-reloads every 10s (see `dashboard.yml`)
- Use `allowUiUpdates: true` to edit in UI, then export JSON
- Test queries in Prometheus UI first: `http://localhost:9090`
