# my first go project
## what this is
This is a playground which makes trees of your desired depth.
  - you are able to add nodes to the graph event after it is generated
  - it makes a graph.txt file if you want to hang onto it after you make it
Sample graph (of depth input 3):
```
                         |-- 12
                 |-- 8
                         |-- 11
         |-- 3
                         |-- 10
                 |-- 7
                         |-- 9
 |-ROOT 1
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
