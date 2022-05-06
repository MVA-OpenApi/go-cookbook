# Installing Go and Cobra-CLI-Generator (system wide):

go: https://go.dev/doc/install
cobra-cli: go install github.com/spf13/cobra-cli@latest

# Initializing project

1. mkdir <project_name>
2. cd <project_name>
3. create go project: go mod init github.com/<link_to_project>
4. add cobra to modules: go get -u github.com/spf13/cobra@latest
5. generate boiler plate code: cobra-cli init

# Run CLI-Tool

- go run main.go

# Build and run CLI-Tool

1. go build
2. .\<tool_name>.exe (for Windows)
