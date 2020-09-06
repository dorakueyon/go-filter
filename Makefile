linux:
	GOOS=linux GOARCH=amd64 go build -o linux-amd64/mdfilter ./main.go
win:
	GOOS=windows GOARCH=amd64 go build -o windows-amd64/md_filter.exe ./main.go
