#!/bin/sh
mango=${mango:-mango-doc}
manual=${manual:-GOBLIN}

mkman ()
{
	section="$1"
	path="$2"
	base="`basename $path`"
	manpage="man/man$section/$base.$section"
	$mango -name "$base" -manual "$manual" "$path" >"$manpage"
}

for i in cmd/*
do
	mkman 1 "$i"
done
