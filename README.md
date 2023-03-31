# drone-file-secret-extension

A secret provider for [Drone CI](https://www.drone.io/). It simply reads secrets from a given folder, suitable for private use Drone CI instances where running a Vault instance can be undesirable.

This project is based on [drone-vault](https://github.com/drone/drone-vault) secret provider.

**This project is experimental.**

## Installation

Create a shared secret:

```text
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the plugin:

```text
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=DRONE_BASE_PATH=/secrets \
  -v "/path/to/secrets:/secrets" \
  --restart=always \
  --name=drone-file-secret drone/file-secret
```

Update your runner configuration to include the plugin address and the shared secret.

```text
DRONE_SECRET_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_SECRET_PLUGIN_TOKEN=bea26a2221fd8090ea38720fc445eca6
```
