install: all

all:
	go install ./cmds/...

linux:
	mkdir -p linux_bin
	GOOS=linux GOARCH=amd64 go build -o ./linux_bin/ijson ./cmds/ijson/*.go
	GOOS=linux GOARCH=amd64 go build -o ./linux_bin/lsip ./cmds/lsip/*.go

clean:
	rm -rf linux_bin
