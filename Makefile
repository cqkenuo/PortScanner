LinuxOS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
MacOS=CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
Windows=CGO_ENABLED=0 GOOS=windows GOARCH=amd64
DATE=$(shell date +'%Y-%m-%d %H:%M:%S')

mac:
		@echo "  >  Building mac binary..."
		$(MacOS) go build -o cmd/server server.go
		$(MacOS) go build -o cmd/agent agent.go
		@echo "MacOS 版本编译完成"
		@echo $(DATE)

linux:
		@echo "  >  Building linux binary..."
		$(LinuxOS) go build -o cmd/server server.go
		$(LinuxOS) go build -o cmd/agent agent.go
		@echo "Linux 版本编译完成"
		@echo $(DATE)

all:
		$(MacOS) go build -o cmd/server_darwin server.go
		$(MacOS) go build -o cmd/agent_darwin agent.go
		$(LinuxOS) go build -o cmd/server_linux server.go
		$(LinuxOS) go build -o cmd/agent_linux agent.go