ENV_FILE=.env
Nginx_Config_Folder=webapp/nginx

all: help

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

## build: 打包前端，后端docker镜像
.PHONY: build
build:
	@rm -rf webapp/nginx
	@mkdir -p webapp/nginx
	@cp nginx/nginx.conf webapp/nginx/
	@docker-compose build
	@echo building docker image...

## run: 启动项目
.PHONY: run
run:
	docker-compose --env-file ${ENV_FILE} up -d


## log: 查看docker-compose日志
.PHONY: log
log:
	docker-compose logs -f

## run-server-only: 只运行后端应用
.PHONY: run-server-only
run-server-only:
	@echo run server..
	@cd server && go run cmd/main.go


.PHONY: help
help:
	@echo "直接 make start 即可启动项目，make stop停止,更多用法如下: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
