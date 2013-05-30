bindir = bin

all: cmds-go cmds-sh

cmds-go:
	GOBIN=$(bindir) go install ./cmd/...

cmds-sh:
	cp sh/* $(bindir)

clean:
	rm -f $(bindir)/*
