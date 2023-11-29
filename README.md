# j2ndj #

[![goreleaser](https://github.com/do-i-need-a-username/j2ndj/actions/workflows/release.yml/badge.svg)](https://github.com/do-i-need-a-username/j2ndj/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/do-i-need-a-username/j2ndj)](https://goreportcard.com/report/github.com/do-i-need-a-username/j2ndj)
[![codecov](https://codecov.io/gh/do-i-need-a-username/j2ndj/branch/main/graph/badge.svg?token=ZQZQZQZQZQ)](https://codecov.io/gh/do-i-need-a-username/j2ndj)
[![Go Reference](https://pkg.go.dev/badge/github.com/do-i-need-a-username/j2ndj.svg)](https://pkg.go.dev/github.com/do-i-need-a-username/j2ndj)

Converts a json file to [ndjson](https://ndjson.org/) file. (New Line Delimited JSON)

## Installation ##

```bash
# Install go
go install github.com/do-i-need-a-username/j2ndj@latest
# fins install path
which j2ndj
/Users/myuser/go/bin/j2ndj
# Run the command
/Users/myuser/go/bin/j2ndj -input /tmp/logs.json -output /tmp/logs.ndjson
```
