FROM debian:bookworm-slim

WORKDIR /app

COPY --link --chmod=0644 foo /foo
COPY --link --chmod=0755 run.sh /run.sh

RUN apt-get update && apt-get install -y build-essential liburing-dev
RUN --mount=type=bind,target=/app \
	gcc -Wall -O2 -D_GNU_SOURCE -o /usr/bin/io_uring-cp ./io_uring-cp.c -luring

RUN adduser --gid 100 vincent
USER vincent
