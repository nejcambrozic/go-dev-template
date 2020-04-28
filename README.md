# Go developer template

This contains a bare bone template to bootstrap your Golang development project with Docker!
It includes:

* Dev environment in Docker:
  * Debugging breakpoints with Delve and Goland 
  * Live reloading with [CompileDaemon](https://github.com/githubnemo/CompileDaemon)
  * Production build with optimized Docker image
  * docker-compose file to bring dev environment up
  * Makefile to run all the needed commands to get you going
* Golang:
    * Project structure that is recommended for Golang (`cmd` and `pkg` directories)
    * Basic http server with 2 endpoints each demonstrating different implementation:
        * `/ping`: simple "hello world" like function
        * `/health`: uses a service with an interface


# Get Started

Use make for everything, this assumes you have the code checkout and are in the repo root dir:
Makefile contains help and has all commands documented, to see all the commands simply run `make`

## Start the app with live reloading:

`make dev`: this start the server with live realoading and exposes it on port `3000`

You can check the health endpoint: `http://localhost:3000/health`, it should return `{"Status":"ok"}`

Any change in `.go` file will now trigger a rebuild and you should see this in your terminal:
```bash
api_1  | Running build command!
api_1  | Build ok.
api_1  | Hard stopping the current process..
api_1  | Restarting the given command.
api_1  | stderr: Server started
```

## Debugging in Goland

Is described in [a separate document](doc/debugging-in-ide.md)
 

# Licence
[GNU GPLv3](LICENSE)






 



 