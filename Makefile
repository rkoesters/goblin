all: cmds

cmds:
	GOBIN=bin go install ./cmd/...

clean:
	rm -f bin/*
