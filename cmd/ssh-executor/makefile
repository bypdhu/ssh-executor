BUILD_NAME         := ssh-executor
BUILD_VERSION    := $(shell git describe --tags)
BUILD_TIME           := $(shell date "+%F-%H-%M-%S")
BUILD_GO_VERSION   := $(shell go version)
TAR_BALL := ${BUILD_NAME}_${BUILD_VERSION}_${BUILD_TIME}.tar.gz

clean:
	rm -f ${BUILD_NAME}

build: clean
	go build -ldflags \
    "-X 'git.eju-inc.com/ops/go-common/version.AppName=${BUILD_NAME}' \
     -X 'git.eju-inc.com/ops/go-common/version.AppVersion=${BUILD_VERSION}' \
     -X 'git.eju-inc.com/ops/go-common/version.BuildTime=${BUILD_TIME}' \
     -X 'git.eju-inc.com/ops/go-common/version.BuildGoVersion=${BUILD_GO_VERSION}' \
     " \
    -o ${BUILD_NAME}

buildlinux: clean
	GOOS=linux GOARCH=amd64 go build -ldflags \
    "-X 'git.eju-inc.com/ops/go-common/version.AppName=${BUILD_NAME}' \
     -X 'git.eju-inc.com/ops/go-common/version.AppVersion=${BUILD_VERSION}' \
     -X 'git.eju-inc.com/ops/go-common/version.BuildTime=${BUILD_TIME}' \
     -X 'git.eju-inc.com/ops/go-common/version.BuildGoVersion=${BUILD_GO_VERSION}' \
     " \
    -o ${BUILD_NAME}

tarball: clean buildlinux
	tar cvf ${TAR_BALL} ${BUILD_NAME}

all: clean buildlinux tarball
	curl -i -X PUT -T ${TAR_BALL} http://demo1.ess.ejucloud.cn/flow-notifier/${TAR_BALL}
