all: cmds

cmds:
	GOBIN=bin go install ./cmd/...
	cp sh/* bin

clean:
	rm -f bin/*
