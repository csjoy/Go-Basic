# Go-Basic

## Go Debug Environment Setup for Visual Studio Code

### Step 1
put the below code in settings.json file
```json
"launch": {

        "configurations": [
            {
                "name": "Remote",
                "type": "go",
                "request": "attach",
                "mode": "remote",
                "remotePath": "${fileDirname}",
                "port": 2345,
                "host": "127.0.0.1",
                "apiVersion": 2
            }
        ],
        "compounds": []
    },
```
 
### Step 2
Install delve
```go
go install github.com/go-delve/delve/cmd/dlv@latest
```
 
### Step 3
add GOBIN environment variable and a alias to start debug session
```bash
export PATH="$HOME/go/bin:$PATH"
alias godebug="dlv debug --headless --listen=:2345 --log --api-version=2 -- $@"
```
### Step 4
To use VSCode debug window make sure to include go.mod file and a breakpoint then type the alias godebug in the terminal from the project directory. Finally click the green button next to Remote or press F5 to go into the debug mode.
