bindir = bin
mandir = man

all: cmds-go cmds-sh cmds-doc

cmds-go: $(bindir)
	GOBIN=$(bindir) go install ./cmd/...

cmds-sh: $(bindir)
	cp sh/* $(bindir)

cmds-doc: $(mandir)
	tools/mkdoc.sh

$(bindir):
	mkdir -p $(bindir)

$(mandir):
	mkdir -p $(mandir)

clean:
	rm -f $(bindir)/*
	rm -f $(mandir)/*/*
