# envdiff

How to use ?

- First, you need install envdiff in your local with go
```bash
go install github.com/rsmnarts/envdiff
```

- then, you run command like this
```bash
envdiff -source={source_path} -target={target_path}
```
example :
```bash
envdiff -source .env -target .env.bak
````
