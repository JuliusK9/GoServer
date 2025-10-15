# Process Manager
---
<br>


## Features

- **HTTP REST API** for process management
- **Command Line Interface** for easy interaction  
- **Real-time process monitoring** 

<br>

<div id="toc">

### 🚩 Table of Contents

- [Install](#install)
- [Using The CLI](#using-the-cli)
- [CLI commands](#cli-commands)
    -[server](#server)
    -[start](#start)
    -[list](#list)
    -[help](#help)
    

</div>

***
## Install
```bash
# Clone the repository
git clone https://github.com/JuliusK9/GoServer
cd GoServer

# Build the CLI tool
go build -o processmanager ./cmd/cli
```


**[🔝 back to top](#toc)**

***
## Using the CLI

The CLI provides an intuitive way to manage processes:

```bash
# Start the server (in one terminal)
./processmanager server

# In another terminal, manage processes
./processmanager start "process_1"
./processmanager start "process_2"
./processmanager list
```

**[🔝 back to top](#toc)**

***
## CLI commands

```server```

Starts the HTTP server on port 8080.

```bash
./processmanager server
```

```start [name]```

Starts a new process with the specified name.

```bash
./processmanager start "NewProcess"
```

```list```

Lists all currently running processes.

```bash
./processmanager list
```

```help```

Shows help information for any command.

```bash
./processmanager help
./processmanager help start
```

**[🔝 back to top](#toc)**