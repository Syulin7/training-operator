IMG ?= registry.cn-beijing.aliyuncs.com/acs/tf_operator
VERSION ?= v1.0-aliyun
GIT_SHORT_COMMIT=$(shell git rev-parse --short HEAD)
TAG=${VERSION}-${GIT_SHORT_COMMIT}

.PHONY: build
build: ## build exec file gpushare-device-plugin on host
	@docker build . --rm --no-cache -t ${IMG}:${TAG}
