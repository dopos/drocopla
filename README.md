# drocopla

A drone.io conversion extension to set host platform as drone pipeline platform. 

_Please note this project requires Drone server version 1.4 or higher._

Default drone behaviour:

If `os/arch` is not set in `.drone.yml` - fill them with `linux/amd64`.

As result, such pipelines becomes not available for drone-runners on non-default plaforms.
Meanwhile, if you set platform values in `.drone.yml`, such pipelines becomes not available for drone-runners on other plaforms

**drocopla** adds to `.drone.yml` platform tags with values from its runtime.

If `.drone.yml` contains platform tag, it does nothing.

There is a [fix for drone.io sources](https://github.com/LeKovr/drone/commit/886f18f8ad368e4b7c8882d70f709f9535bd277f) which solves the same problem, but it is not merged yet.

## Installation

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --restart=always \
  --name=converter ghcr.io/dopos/drocopla
```

Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6

## Acknowledgements

This code based on [boilr-convert](https://github.com/drone/boilr-convert) template.
