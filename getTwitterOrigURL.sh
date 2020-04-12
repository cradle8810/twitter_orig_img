#!/bin/sh

if [ 0 -eq $# ]; then
	echo "Usage: $0 [file]" >&2
	exit 1
fi

# <meta  property="og:image" content="(HERE)">

cat "$1" | grep 'pbs.twimg.com/media/' | grep 'meta' | \
    perl -nle '$F[0]=/content=\"(.+?)\"/; print $1'|\
    sed -e 's/:large//g'| awk '{print $1 ":orig"}'
