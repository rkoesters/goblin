#!/bin/sh
if [ "x$1" != "xbuild" ] && [ "x$1" != "xclean" ]
then
	echo "usage: $0 [ build | clean ]"
	exit 1
fi

mandir=${mandir:-man}
mango=${mango:-mango-doc}
manual=${manual:-Goblin Commands}

manpage () {
	section="$1"
	path="$2"
	name="`basename $path`"
	echo "$mandir/man$section/$name.$section"
}

build () {
	$mango -manual "$manual" "$2" >"`manpage "$@"`"
}

clean () {
	rm "`manpage "$@"`"
}

mkdir -p $mandir/man1
for i in cmd/*
do
	$1 1 "$i"
done
