#!/usr/bin/env bash

GENDOC_DIRPATH="${GENDOC_DIRPATH:-"/output"}"

if [ "$1" = "kafkactl" ]; then
    shift
    exec kafkactl "${@}"
elif [ "$1" = "gendoc" ]; then
    exec gendoc "${GENDOC_DIRPATH}"
else
    exec "$@"
fi
