# envdiff

[![Actions Status](https://github.com/rsmnarts/envdiff/workflows/build/badge.svg)](https://github.com/rsmnarts/envdiff/actions)
[![codecov](https://codecov.io/gh/rsmnarts/envdiff/branch/master/graph/badge.svg)](https://codecov.io/gh/rsmnarts/envdiff)

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
