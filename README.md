## Put your project in someplace safe and warm
### For Windows

```bash
GOOS=windows GOARCH=amd64 go build -o tree.exe tree.go
```
This command compiles your Go application for the Windows operating system and the amd64 architecture. It generates an executable named tree.exe.

### For macOS
```
GOOS=darwin GOARCH=amd64 go build -o tree tree.go
```
### For Linux
```
GOOS=linux GOARCH=amd64 go build -o tree tree.go

```
