# my first go project
## what this is
This is a playground which makes trees of your desired depth.
Sample graph:

Your tree:
                                 |-- 28
                         |-- 24
                                 |-- 27
                 |-- 16
                                 |-- 26
                         |-- 23
                                 |-- 25
         |-- 3
                                 |-- 22
                         |-- 18
                                 |-- 21
                 |-- 15
                                 |-- 20
                         |-- 17
                                 |-- 19
 |-- 1
                                 |-- 14
                         |-- 10
                                 |-- 13
                 |-- 2
                                 |-- 12
                         |-- 9
                                 |-- 11
         |-- 2
                                 |-- 8
                         |-- 4
                                 |-- 7
                 |-- 1
                                 |-- 6
                         |-- 3
                                 |-- 5




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
