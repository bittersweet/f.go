all:
	@@go build f.go
	@@cp f /usr/local/bin/
	@@echo "built and moved"
