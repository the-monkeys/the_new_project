#!/usr/bin/bash

MONKEY_ROOT="/the_monkeys"

MONKEY_ETC="${MONKEY_ROOT}/etc"
MONKEY_ENV_FILE="${MONKEY_ETC}/dev.env"

mkdir -p "$MONKEY_ROOT"
mkdir -p "$MONKEY_ETC"
touch $MONKEY_ENV_FILE
chmod +x $MONKEY_ENV_FILE

function sh_perror()
{
    echo "$@" >/dev/stderr
}

