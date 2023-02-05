#!/usr/bin/bash

MONKEY_ROOT="/the_monkeys"

mkdir -p "$MONKEY_ROOT"

function sh_perror()
{
    echo "$@" >/dev/stderr
}

