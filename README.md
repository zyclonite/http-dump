[![Docker Pulls](https://badgen.net/docker/pulls/zyclonite/http-dump)](https://hub.docker.com/r/zyclonite/http-dump)
[![Quay.io Enabled](https://badgen.net/badge/quay%20pulls/enabled/green)](https://quay.io/repository/zyclonite/http-dump)
[![build](https://github.com/zyclonite/http-dump/actions/workflows/build.yml/badge.svg)](https://github.com/zyclonite/http-dump/actions/workflows/build.yml)

## http-dump
Dump raw http requests to the console

### build

`docker build -t zyclonite/http-dump .`

### run

`docker run --name http-dump -d -p 8080:8080 zyclonite/http-dump`

### use

open `http://127.0.0.1:8080/` in your browser and watch the container log
