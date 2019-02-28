install: all

all:
	go install ./cmds/...

linux:
	mkdir -p linux_bin
	GOOS=linux GOARCH=amd64 go build -o ./linux_bin/ijson ./cmds/ijson/*.go
	GOOS=linux GOARCH=amd64 go build -o ./linux_bin/lsip ./cmds/lsip/*.go
	GOOS=linux GOARCH=amd64 go build -o ./linux_bin/iproxy ./cmds/iproxy/*.go

	scp linux_bin/iproxy root@m:/root/

clean:
	rm -rf linux_bin
