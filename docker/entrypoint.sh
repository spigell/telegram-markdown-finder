#!/bin/bash

if [[ -d /tmp/entrypoint.d ]]; then
  for f in /tmp/entrypoint.d/*.sh ;do
    source $f
  done
fi

/root/tg-markdown-finder -config /etc/tg-markdown-finder.json