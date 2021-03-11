# API mTLS Ambassador (Client)

Ambassador Docker container used to authenticate using mTLS for the Open Banking and PIX API communication as client (request)

## Quick Start

First, bring up the [mTLS API server sidecar example](https://github.com/bancodobrasil/api-mtls-sidecar-server).

Now, to use this Ambassador to connect to this service using mTLS, open your terminal and:

```bash

git clone https://github.com/bancodobrasil/api-mtls-ambassador-client.git

```

```bash

docker-compose up

```

Then run `curl`:

```bash
$ curl http://localhost:9090
{"data":"I'm secured by an mTLS!"}
```
