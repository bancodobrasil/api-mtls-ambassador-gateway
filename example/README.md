# Example: Invoking a mTLS secured API Endpoint from your component

To see this example running, do the following:

First, bring up the [mTLS API server sidecar example](https://github.com/bancodobrasil/api-mtls-sidecar-proxy/tree/main/example).

Then:

```bash
docker-compose up
```

You should see in the output

```bash
invoker_1          | ******
invoker_1          |
invoker_1          |
invoker_1          | Invocation return: `I'm a POST secured by an mTLS!. You said: I'm behind the Ambassador`
invoker_1          |
invoker_1          |
invoker_1          | ******
```
