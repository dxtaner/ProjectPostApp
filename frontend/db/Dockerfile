FROM postgres:latest

ENV POSTGRES_PASSWORD depixen-pass

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5439
