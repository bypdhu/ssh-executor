set one=%1

if "%one%"=="linux" (
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
) else (
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
)


SET BUILD_NAME=ssh-executor
SET BUILD_VERSION=V1.0.0
SET BUILD_TIME=%date:~,4%%date:~5,2%%date:~8,2%_%time:~,2%%time:~3,2%%time:~6,2%
SET BUILD_GO_VERSION=go version go1.8 windows/amd64

go build -ldflags "-X 'git.eju-inc.com/ops/go-common/version.AppName=%BUILD_NAME%' -X 'git.eju-inc.com/ops/go-common/version.AppVersion=%BUILD_VERSION%' -X 'git.eju-inc.com/ops/go-common/version.BuildTime=%BUILD_TIME%' -X 'git.eju-inc.com/ops/go-common/version.BuildGoVersion=%BUILD_GO_VERSION%' "