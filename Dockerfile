FROM golang:latest as builder

MAINTAINER Andrey Vladislavov

RUN mkdir /go/src/BugTrackerBackend_DB

COPY . /go/src/BugTrackerBackend_DB

WORKDIR /go/src/BugTrackerBackend_DB

RUN go build -o bugTrackerBackend .

FROM ubuntu:20.04 AS release

MAINTAINER Andrey Vladislavov

RUN apt-get update -y && apt-get install -y locales gnupg2
RUN locale-gen en_US.UTF-8
RUN update-locale LANG=en_US.UTF-8

ENV PGVER 12
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update -y && apt-get install -y postgresql postgresql-contrib

USER postgres

COPY db.sql /home

WORKDIR /home

RUN /etc/init.d/postgresql start &&\
    psql --command "CREATE USER bugtrack WITH SUPERUSER PASSWORD 'bugtrack_pass';" &&\
    createdb -E UTF8 bugtrackcourse &&\
    psql --command "\i '/home/db.sql'" &&\
    /etc/init.d/postgresql stop

RUN echo "listen_addresses='*'\n" >> /etc/postgresql/$PGVER/main/postgresql.conf
RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/$PGVER/main/pg_hba.conf

VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root

COPY --from=builder /go/src/BugTrackerBackend_DB/bugTrackerBackend /usr/bin/bugTrackerBackend

#EXPOSE 5432
#EXPOSE 5000

CMD service postgresql start && bugTrackerBackend

#---------------------

#COPY go.mod go.sum /go/src/github.com/AndVl1/bugTrackerBackend/
#WORKDIR /go/src/github.com/AndVl1/bugTrackerBackend
#RUN go mod download
#COPY . /go/src/github.com/AndVl1/bugTrackerBackend
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bugTrackerBackend github.com/AndVl1/bugTrackerBackend
#
#FROM alpine
#RUN apk add --no-cache ca-certificates && update-ca-certificates
#COPY --from=builder /go/src/github.com/AndVl1/bugTrackerBackend/build/bugTrackerBackend /usr/bin/bugTrackerBackend
#EXPOSE 8080 8080
#ENTRYPOINT ["/usr/bin/bugTrackerBackend"]