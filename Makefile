ENV_FILE=.env
Nginx_Config_Folder=webapp/nginx

all: help

## serve:
.PHONY: serve
serve: build-local run

## re-serve:
.PHONY: re-serve
re-serve: stop serve

.PHONY: build-local
build-local:

	@rm -rf webapp/nginx
	@mkdir -p webapp/nginx
	@cp nginx/nginx.local.conf webapp/nginx/nginx.conf
	@docker-compose build
	@echo building docker image...

## start: 打包docker镜像并启动项目
.PHONY: start
start: build run

## stop: 停止项目
.PHONY: stop
stop:
	docker-compose down

## restart: 重新启动
.PHONY: restart
restart: stop start

.PHONY: build
build:
	@rm -rf webapp/nginx
	@mkdir -p webapp/nginx
	@cp nginx/nginx.conf webapp/nginx/
	@docker-compose build
	@echo building docker image...

.PHONY: run
run:
	docker-compose --env-file ${ENV_FILE} up -d


## log: 查看docker-compose日志
.PHONY: log
log:
	docker-compose logs -f

.PHONY: run-server-only
run-server-only:
	@echo run server..
	@cd server && go run cmd/main.go


.PHONY: help
help:
	@echo "直接 make start 即可启动项目，make stop停止,更多用法如下: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
