FROM golang:1.25.5 AS build
WORKDIR /usr/src/app

SHELL [ "/bin/bash", "-o", "pipefail", "-c" ]

RUN apt-get -y update && apt-get -y --no-install-recommends install unzip=6.0-29
RUN wget -qO- https://bun.com/install | ENV="$HOME/.bashrc" SHELL="$(which bash)" bash -
ENV PATH="/root/.bun/bin:$PATH"

COPY go.mod go.sum package.json bun.lock ./
RUN go mod download && bun i --frozen-lockfile

COPY . .
RUN make

FROM alpine:latest AS run
WORKDIR /usr/src/app
COPY --from=build /usr/src/app/bin/app ./
EXPOSE 8080
CMD [ "./app" ]
