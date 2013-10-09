bindir = bin
mandir = man

all: cmds doc
clean: clean-cmds clean-doc

cmds:
	GOBIN=$(bindir) go install ./cmd/...

clean-cmds:
	GOBIN=$(bindir) go clean -i ./cmd/...

doc:
	mandir=$(mandir) tools/mkdoc.sh build

clean-doc:
	mandir=$(mandir) tools/mkdoc.sh clean
