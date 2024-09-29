FROM ubuntu:jammy
RUN apt-get update && apt-get install -y curl iputils-ping
WORKDIR /app
COPY ./cmd/usersvc/usvc .
RUN chmod +x usvc
# ENTRYPOINT [ "tail", "-f", "/dev/null" ]
ENTRYPOINT [ "./usvc" ]

# RUN curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
# RUN echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | \
# tee /etc/apt/sources.list.d/mongodb-org-7.0.list

# RUN apt-get install -y postgresql
# RUN apt-get update && apt-get install -y mongodb-org
# WORKDIR /app
# COPY /init/runmaster.sh ./
# COPY /init/master ./
# # COPY /init/tester ./
# RUN chmod +x runmaster.sh
# EXPOSE 80
# EXPOSE 50051

# ENTRYPOINT [ "./runmaster.sh" ]
# ENTRYPOINT [ "tail", "-f", "/dev/null" ]