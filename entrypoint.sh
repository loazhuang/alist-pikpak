#!/bin/bash

chown -R ${PUID}:${PGID} /opt/alist/

umask ${UMASK}

if [ "$1" = "version" ]; then
  ./alist version
else

 if [ "$RUN_ARIA2" = "true" ]; then
    chown -R ${PUID}:${PGID} /opt/aria2/
    exec su-exec ${PUID}:${PGID} nohup aria2c \
      --enable-rpc \
      --rpc-allow-origin-all \
      --conf-path=/opt/aria2/.aria2/aria2.conf \
      >/dev/null 2>&1 &
  fi
  exec su-exec ${PUID}:${PGID} ./alist server --no-prefix
fi
