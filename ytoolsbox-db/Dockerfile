FROM postgres:alpine as dumper

COPY demo.sql /docker-entrypoint-initdb.d/

RUN ["sed", "-i", "s/exec \"$@\"/echo \"skipping...\"/", "/usr/local/bin/docker-entrypoint.sh"]

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=test123456
ENV PGDATA=/data

RUN ["/usr/local/bin/docker-entrypoint.sh", "postgres"]

# final build stage
FROM postgres:alpine

COPY --from=dumper /data $PGDATA
