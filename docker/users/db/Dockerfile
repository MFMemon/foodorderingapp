FROM postgres
ENV POSTGRES_PASSWORD secret
ENV POSTGRES_USER usersdb
COPY /init/db/userdb.sql /docker-entrypoint-initdb.d/