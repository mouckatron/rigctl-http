# rigctl-http

Aims to provide an HTTP API interface between a user/software and `rigctld`.

It connects to `rigctld`s telnet interface to run its commands.

It will make the `rigctl` long form commands available as endpoints, for example `dump_caps` at `/dump_caps` but will use `GET` and `POST` requests for long form commands prefixed with `get_` and `set_`, for example `get_frequency` will be at endpoint `/frequency`.

## Example

Some quick examples using `curl` and assuming `rigctl-http` is running at `localhost:8080`

```sh
$ curl http://localhost:8080/frequency
{
    "success": true,
    "raw": "get_freq:\nFrequency: 145500000\nRPRT 0\n",
    "data": {
        "frequency": "145500000"
    }
}

$ curl -XPOST http://localhost:8080/frequency -d '{"frequency":"144800000"}'
{
    "success": true,
    "raw": "set_freq: 144800000\nRPRT 0\n",
    "data": {
        "frequency": "144800000"
    }
}
```

## Testing

The suite of integration tests can be run against `rigctl`s dummy rig. `make integration` will start `rigctld` and `rigctl-http` and run the tests in `/test/integration`.

```sh
make integration
```
