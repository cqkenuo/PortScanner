LinuxOS=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
MacOS=CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
Windows=CGO_ENABLED=0 GOOS=windows GOARCH=amd64
DATE=$(shell date +'%Y-%m-%d %H:%M:%S')

mac:
		@echo "  >  Building mac binary..."
		$(MacOS) go build -o cmd/server_darwin scanServer.go
		$(MacOS) go build -o cmd/agent_darwin scanAgent.go
		@echo "MacOS 版本编译完成"
		@echo $(DATE)

linux:
		@echo "  >  Building linux binary..."
		$(LinuxOS) go build -o cmd/server_linux scanServer.go
		$(LinuxOS) go build -o cmd/agent_linux scanAgent.go
		@echo "Linux 版本编译完成"
		@echo $(DATE)

windows:
		@echo "  >  Building windows binary..."
		$(Windows) go build -o bin/server.exe scanServer.go
		$(Windows) go build -o bin/agent.exe scanAgent.go
		@echo "Windows 版本编译完成"
