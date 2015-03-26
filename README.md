# etcdjson

## What

This is a pretty simple tool that provides a CLI tool to edit
JSON that's stored in Etcd with a JSONPath-like interface.

## Why

Because I got tired of getting JSON data from Etcd, storing it in a
file, editing it, and then setting the new data.

## Details

This tool has `get` and `set`. It's not concurrency safe for a key, so use
it from 1 node at a time to edit some JSON stored in a key.

### Commands

Assuming this JSON is stored in `/config/json` in Etcd

```json
{
  "a": {
    "b": "c"
  }
}
```

#### get

`etcdjson get /config/json a.b` returns `c`

#### Set
`etcdjson set /config/json a.b d` returns `success` and makes the JSON look like this:

```json
{
  "a": {
    "b": "d"
  }
}
```
