#!/bin/sh
mango=${mango:-mango-doc}
manual=${manual:-GOBLIN}

mkman ()
{
	section="$1"
	path="$2"
	name="`basename $path`"
	manpage="man/man$section/$name.$section"
	$mango -name "$name" -manual "$manual" "$path" >"$manpage"
}

for i in cmd/*
do
	mkman 1 "$i"
done
