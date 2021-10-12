# envdiff

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
