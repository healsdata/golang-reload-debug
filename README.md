# go-reload-debug

An example of using Air (for live reload), Delve (for debugging), and Cobra (for CLI commands) together in Docker for local development.

## Highlights

This is intended as a simple example of getting the two tools working together. 

For this example, the entry point in the Dockerfile is air 

```CMD ["air", "-c", ".air.toml"]```

On first run and reload, Air builds the new executable.

```go build -o ./.dev/main .```

Then Air uses Delve to run the Cobra Cmd

```dlv exec ... ./.dev/main serve```

## Usage

To see it in action, you can run `docker compose up -d` from the root directory.

### Reloading

After that, any edit you make to main.go will trigger a reload. You can test it by changing the message in `indexHandler` to a different string.

### Debugging

To configure debugging in Goland, you'll want to follow the steps outlined in [Attach to a process on a remote machine](https://www.jetbrains.com/help/go/attach-to-running-go-processes-with-debugger.html#attach-to-a-process-on-a-remote-machine)

## Tools
* [cosmtrek/air](https://github.com/cosmtrek/air): ☁️ Live reload for Go apps
* [go-delve/delve](https://github.com/go-delve/delve): Debugger for the Go programming language.
* [spf13/cobra](https://github.com/spf13/cobra): Library for creating powerful modern CLI applications.

## References
* [Today I Learned: Golang Live-Reload for Development Using Docker Compose + Air](https://medium.easyread.co/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076)
* [Debugging & Live Reloading Go app within Docker Container](https://medium.com/@hananrok/debugging-hot-reloading-go-app-within-docker-container-b44d2929e8bd)
* [Go development with Docker Containers](https://blog.jetbrains.com/go/2020/05/04/go-development-with-docker-containers/)
* [Creating a web server with Golang](https://blog.logrocket.com/creating-a-web-server-with-golang/)