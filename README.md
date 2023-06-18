# Container from scratch

This follows Liz Rice's talk on [Youtube](https://www.youtube.com/watch?v=8fi7uSYlOdc) on implementing a container from scratch. 

## Usage

```bash
$ go run main.go run /bin/sh
```

You need to be on Linux and have root privileges. You can gain root privileges by running:

```bash
sudo -i
```
I have included an Alpine Linux filesystem in the repo. You should copy this and modify this line in the `main.go` file to point to the location of the filesystem on your machine:

```go
const rootfs = "absolute/path/to/rootfs"
```
