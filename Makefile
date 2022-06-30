VERSION = 2.1

ifeq ($(OS),Windows_NT)
    #Windows 
    CURRENT_DIR = $(CURDIR)
else
    #Linux
    CURRENT_DIR = $(shell pwd)
endif

BUILD_VERSION = $(VERSION)
BUILD_UPLOAD_PATH ?= $(CURRENT_DIR)/release

IMAGE= notary-verifyhash-image
CONTAINER = notary-verifyhash-container
BUILD_CMD = bash -e ./build.sh
DOCKER_CMD = docker run --name $(CONTAINER) --rm -i -v "$(CURRENT_DIR):/mnt" -w /mnt $(IMAGE)

build-package: image
	$(DOCKER_CMD) $(BUILD_CMD)

build:
	mkdir -p ${BUILD_UPLOAD_PATH};tar --exclude='./.git' --exclude='./release' -zcvf ${BUILD_UPLOAD_PATH}/archive.tgz .;

image: 
	docker build -t $(IMAGE) -f Dockerfile .
