#!/bin/sh

HOST=${RAB_HOST}
PORT=${RAB_PORT}
if [ -z "$HOST" ] || [ -z "$PORT" ]; then
  echo "Uso: $0 Rabbitmq host end port not defined"
  exit 1
fi

while ! nc -z -w3 "$HOST" "$PORT"; do


  echo "Aguardando a porta $PORT responder em $HOST..."
  sleep 2

done
  echo "Porta $PORT está aberta em $HOST"
exec /app/server