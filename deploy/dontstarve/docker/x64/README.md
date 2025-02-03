# Docker Image For Don't Starve Together Dedicated Server 64-bit

This Docker image provides a ready-to-use Don't Starve Together (DST) dedicated server, built on the [Official SteamCMD Debian Image](https://hub.docker.com/r/steamcmd/steamcmd). Only support running at 64-bit mode.

## Usage

To show the server version

```bash
$ docker run --rm dstgo/dst-server-x64
```

To run the dedicated server
```bash
$ docker run dstgo/dst-server-x64 ./dontstarve_dedicated_server_nullrenderer_x64
```

## Maintenance

This image is maintained by the dstgo team. For issues, feature requests, or contributions, please visit our [GitHub repository](https://github.com/dstgo/wilson).