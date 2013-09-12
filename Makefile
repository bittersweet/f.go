all:
	@@go build ff.go
	@@cp ff /usr/local/bin/
	@@echo "built and moved"
