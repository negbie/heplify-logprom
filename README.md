# heplify-logprom

`heplify-logprom` is a [mtail](https://github.com/google/mtail) fork which is a tool for extracting metrics
from application logs to be exported into a timeseries database
or timeseries calculator for alerting and dashboarding.

`heplify-logprom` provides one additional builtin function 
```
sendhep(msg, cid)
```
which takes one extracted message together with one extracted correlation ID
and sends them protobuf encoded to `heplify-server`. With this addition you 
can ship logs to `heplify-server` and expose metrics for prometheus.

It aims to fill a niche between applications that do not export their own
internal state, and existing monitoring systems, without patching those
applications or rewriting the same framework for custom extraction glue code.

The extraction is controlled by [mtail programs](docs/Programming-Guide.md)
which define patterns and actions:

```
# send sipwise logs, use value of 'ID=' for correlation
/.*ID=(?P<cid>\S+)/  {
  /(?P<msg>.*)/  {
    sendhep($msg,$cid)
  }
}
```

Metrics are exported for scraping by a collector as JSON or Prometheus format
over HTTP, or can be periodically sent to a collectd, StatsD, or Graphite
collector socket.

Read the [programming guide](docs/Programming-Guide.md) if you want to learn how
to write your own programs.


## Installation

Best way of installing `heplify-logprom` is to get a precompiled binary.

### Precompiled binaries

Precompiled binaries for released versions are available in the
[Releases page](https://github.com/negbie/heplify-logprom/releases) on Github. Using the
latest production release binary is the recommended one.