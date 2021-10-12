# envdiff

[![Build Status](https://github.com/rsmnarts/envdiff/workflows/build/badge.svg)](https://github.com/rsmnarts/envdiff/actions?query=workflow%3Abuild)
[![Go Report Card](https://goreportcard.com/badge/github.com/rsmnarts/envdiff)](https://goreportcard.com/report/github.com/rsmnarts/envdiff)
[![GoDoc](https://godoc.org/github.com/rsmnarts/envdiff?status.svg)](https://godoc.org/github.com/rsmnarts/envdiff)
[![codecov](https://codecov.io/gh/rsmnarts/envdiff/branch/main/graph/badge.svg?token=slMBQ6cxp0)](https://codecov.io/gh/rsmnarts/envdiff)

How to use ?

- First, you need install envdiff in your local with go

```bash
go install github.com/rsmnarts/envdiff@latest
```

- then, you run command like this

```bash
envdiff -source={source_path} -target={target_path}
```

example :

```bash
envdiff -source=./testdata/example.env -target=./testdata/example-another.env
```
