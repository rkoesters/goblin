bindir = bin
docdir = man

mango   = mango-doc
manual  = GOBLIN
version = 1

all: cmds-go cmds-sh cmds-doc

cmds-go: $(bindir)
	GOBIN=$(bindir) go install ./cmd/...

cmds-sh: $(bindir)
	cp sh/* $(bindir)

cmds-doc: $(docdir)
	tools/mkdoc.sh $(mango) $(docdir) $(manual) $(version)

$(bindir):
	mkdir $(bindir)

$(docdir):
	mkdir $(docdir)

clean:
	rm -f $(bindir)/*
	rm -f $(docdir)/*
