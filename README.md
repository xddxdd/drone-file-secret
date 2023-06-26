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
$ go build
$ env \
  DRONE_BIND=127.0.0.1:3000 \
  DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  DRONE_BASE_PATH=/path/to/secrets \
  ./drone-file-secret
```

Update your runner configuration to include the plugin address and the shared secret.

```text
DRONE_SECRET_PLUGIN_ENDPOINT=http://127.0.0.1:3000
DRONE_SECRET_PLUGIN_TOKEN=bea26a2221fd8090ea38720fc445eca6
```

Define secrets in `.drone.yml`:

```yaml
kind: secret
name: secret_name       # Reference this secret by this name in your pipeline
get:
  path: secret_filename # Relative path of secret DRONE_BASE_PATH, slashes are allowed
  name: anything        # This value has no effect
```
