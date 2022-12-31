# SimpleDB

A simple key-value database (on-disk) written in Go.

It uses hashmap data structure as index.

## Benchmark results

```
[index] Inserted 10000 entries in 23.430793ms
[index] Fetched 10000 entries in 49.957959ms
[no index] Inserted 10000 entries in 27.739258ms
[no index] Fetched 10000 entries in 9.645895597s
```

> Run `go run benchmark/main.go -n <num_of_entries>` to benchmark yourself.
