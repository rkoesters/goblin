#!/bin/sh
if [ $# != 4 ]
then
	echo "usage: $0 mango docdir manual version" >/dev/stderr
	exit 1
fi

mango="$1"
docdir="$2"
manual="$3"
version="$4"

for i in cmd/*
do
	cmdname="`basename $i`"
	$mango -name "$cmdname" -manual "$manual" -version "$version" "$i" >"$docdir/$cmdname.1"
done
