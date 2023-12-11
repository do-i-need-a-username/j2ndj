# j2ndj #

[![goreleaser](https://github.com/do-i-need-a-username/j2ndj/actions/workflows/release.yml/badge.svg)](https://github.com/do-i-need-a-username/j2ndj/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/do-i-need-a-username/j2ndj)](https://goreportcard.com/report/github.com/do-i-need-a-username/j2ndj)
[![codecov](https://codecov.io/gh/do-i-need-a-username/j2ndj/branch/main/graph/badge.svg?token=ZQZQZQZQZQ)](https://codecov.io/gh/do-i-need-a-username/j2ndj)
[![Go Reference](https://pkg.go.dev/badge/github.com/do-i-need-a-username/j2ndj.svg)](https://pkg.go.dev/github.com/do-i-need-a-username/j2ndj)

Converts json to [ndjson](https://ndjson.org/), New Line Delimited JSON.

## Installation ##

```bash
# Install with go
go install github.com/do-i-need-a-username/j2ndj@latest
```

## Usage ##

```bash
# find install path
which j2ndj
/Users/myuser/go/bin/j2ndj

# stdin t ostdout
cat input.json | ./bin/j2ndj
{"age":30,"city":"New York","name":"John Doe"}
{"age":28,"city":"Los Angeles","name":"Jane Doe"}
{"age":35,"city":"Chicago","name":"Bob Smith"}

# stdin to output file
cat input.json | ./bin/j2ndj -output=./output.ndjson

# input file to stdout
./bin/j2ndj -input=./input.json
{"age":30,"city":"New York","name":"John Doe"}
{"age":28,"city":"Los Angeles","name":"Jane Doe"}
{"age":35,"city":"Chicago","name":"Bob Smith"}

# input file to output file
./bin/j2ndj -input=./input.json -output=./output.ndjson
```

## Alternative ##

Use jq to convert json to ndjson

```bash
jq -c '.[]' input.json > output.ndjson
```
