# my first go project
## what this is
This is a playground which makes trees of your desired depth.
Sample graph:
```
                         |-- 12
                 |-- 8
                         |-- 11
         |-- 3
                         |-- 10
                 |-- 7
                         |-- 9
 |-- 1
                         |-- 6
                 |-- 2
                         |-- 5
         |-- 2
                         |-- 4
                 |-- 1
                         |-- 3

```



## How to install (after you put your project in someplace safe and warm)
### For Windows

```bash
GOOS=windows GOARCH=amd64 go build -o tree.exe tree.go
```
### For macOS
```
GOOS=darwin GOARCH=amd64 go build -o tree tree.go
```
### For Linux
```
GOOS=linux GOARCH=amd64 go build -o tree tree.go

```
