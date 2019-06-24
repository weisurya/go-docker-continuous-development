# Continuous Development Environment
## using Docker and VSCode - Go version

## Directory Structure
```bash
|--- .vscode
|       |--- launch.json
|
|--- Dockerfile
|--- Dockerfile.base
|--- Dockerfile.debug
|--- main.go
|--- README.md
|--- request.http
```

### Establish Docker container and integrate with project file in host directory
1. type `docker run --rm -it -v ${pwd}:/go/src/app -w /go/src/app -p 3000:3000 golang:1.12`
2. to check whether you have bind the volume correctly or not, type `ls`
3. you should see your local project files inside there
4. run `go run main.go`

**Explanation regarding the Docker command**
Command | Explanation 
--- | ---
docker run | to create a new Docker container
--rm | delete container after exit, for sanity purpose
-it | to enable interactive mode
-v ${pwd}:/go/src/app | to mount bind between local directory and container directory
-w | to change into a specific directory inside container
-p 3000:3000 | to set inbound-outbound port
golang:1.12 | base image

___
### Use Debug module from VSCode
#### Pre-requisite
1. build the base image first, `docker image build -f Dockerfile.base -t go-base .`

#### Main
1. On your VSCode, click `ctrl + L shift + p`
2. Select/write `Go: Install/Update Tools`
3. Search `dlv` 
4. Click `OK` button
5. Set up `launch.json` file inside `.vscode` folder
6. For debugging purpose, open up a new port to let dlv listen for changes. i.e. `:1234`
7. Create a Dockerfile specifically for debugging purpose. i.e. `Dockerfile.debug`
8. Build it, `docker image build -f Dockerfile.debug -t go-debug .`
9. Run it, `docker run --rm -d -it -v ${pwd}:/go/src/app -w /go/src/app -p 3000:3000 -p 1234:1234 --security-opt=seccomp:unconfined go-debug:latest`
10. Run your debugger and start to place your breakpoint

**Explanation regarding the Docker command**
Command | Explanation 
--- | ---
-d | detach, so you will not be directed into the container's CLI
-p 1234:1234 | open up another port for debugging purpose
--security-opt=seccomp:unconfined | in order to let dlv run into your container, you need to allow it by without using default secure computing mode (seccomp) profile

**Additional explanation about Debug step**
1. To know more about dlv CLI command, check [here](https://github.com/go-delve/delve/blob/master/Documentation/usage/dlv.md)
2. To know more about VSCode debugging, especially in Go, check [here](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code)

#### Separate Docker image between base, debug, and production purpose
1. Create the base image on `Dockerfile.base`. The purpose of this Dockerfile is to build the standard environment that could be applicable for debug and production
2. Create the debugging-purpose image on `Dockerfile.debug`. We do not need to have unrelated packages for production, such as `go-delve` (for debugging purpose) and exposing a new port, which may harm your container.
3. Create the production-purpose image on `Dockerfile`. We set the standard naming style for production.