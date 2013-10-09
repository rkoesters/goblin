bindir = bin
mandir = man/man1

all: cmds doc

cmds:
	GOBIN=$(bindir) go install ./cmd/...

doc: $(mandir)
	tools/mkdoc.sh

$(mandir):
	mkdir -p $(mandir)

clean:
	GOBIN=$(bindir) go clean -i ./cmd/...
	rm -f $(mandir)/*
