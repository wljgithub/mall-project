FROM golang:alpine AS base

WORKDIR /src

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

# build stage
FROM base as build
RUN go build -o /out/main ./cmd/main.go

# bin stage
#FROM scratch AS bin
FROM golang:alpine AS bin
COPY --from=build /out/main /
ADD conf/config.deploy.yml ./conf/config.local.yml
ADD scripts/wait_for.sh ./wait_for.sh

EXPOSE 3000

CMD ["sh","./wait_for.sh","mysql:3306","--","/main"]