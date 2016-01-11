go-steno
========

<a href="https://raw.githubusercontent.com/vjkoskela/gosteno/master/LICENSE">
    <img src="https://img.shields.io/hexpm/l/plug.svg"
         alt="License: Apache 2">
</a>
<a href="https://travis-ci.org/vjkoskela/gosteno/">
    <img src="https://travis-ci.org/vjkoskela/gosteno.png"
         alt="Travis Build">
</a>

Implementation of [ArpNetworking's LogbackSteno](https://github.com/ArpNetworking/logback-steno) for [Go](https://golang.org). Extends [Sirupsen's logrus](https://github.com/Sirupsen/logrus)
logging implementation with a [Steno compatible](testdata/steno.schema.json) Formatter as well as providing named Logger
instances and a fluent log builder.

Dependency
----------

First, retrieve the library into your workspace:

    go> go get github.com/vjkoskela/gosteno

To use the library in your project(s) simply import it:

```go
import "github.com/vjkoskela/gosteno"
```

Formatter
---------

Create the formatter as follows:

```go
var formatter *gosteno.Formatter = gosteno.NewFormatter()
```

The gosteno.Formatter supports a subset of the options available in [LogbackSteno](https://github.com/ArpNetworking/logback-steno):

* LogEventName - Set the default event name. The default is "log".
* InjectContextProcess - Add the process identifier to the context block. The default is true.
* InjectContextHost - Add the host name to the context block. The default is true.
* InjectContextLogger - Add the logger name to the context block. The default is false. (1)

_Note 1_: Injecting additional key-value pairs into context is not strictly compliant with the current definition of Steno.<br>

These may be configured after instantiating the Formatter. For example:

```go
formatter.SetInjectContextLogger(true)
```

Logrus
------

The underlying logging implementation is [logrus](https://github.com/Sirupsen/logrus). At minimum you must set the formatter; however, you may also want to configure the writer and minimum output level. For example to configure the global logger instance:

```go
logrus.SetOutput(os.Stdout)
logrus.SetFormatter(formatter)
logrus.SetLevel(logrus.DebugLevel)
```

Alternatively, you may create a new [logrus](https://github.com/Sirupsen/logrus) logger and configure that. For example:

```
var logrusLogger *logrus.Logger = &logrus.Logger{
    Out: os.Stdout,
    Formatter: formatter,
    Level: logrus.ErrorLevel,
}
```

Logger
------

Instantiate a named Logger instance for each particular logging context. By default the Logger is bound to the global default [logrus](https://github.com/Sirupsen/logrus) logger instance. For example:

```go
var logger *gosteno.Logger = gosteno.GetLogger("http.server")
```

Alternatively, you may bind a Logger instance to particular [logrus](https://github.com/Sirupsen/logrus) logger instance:

```go

var logger *gosteno.Logger = gosteno.GetLoggerForLogger("http.server", logrusLogger)
```

Log Builder
-----------

Aside from providing a full set of logging methods compatible with Go's standard logging and with the [logrus](https://github.com/Sirupsen/logrus) logger, gosteno.Logger also provides access to buildable logs:

```go
logger.DebugBuilder().SetMessage("This is a log builder debug message").Log()

logger.InfoBuilder().SetEvent("my_event").SetMessage("This is a log builder info message with event").Log()

logger.WarnBuilder().
        SetEvent("my_event").
        SetMessage("This is a warn builder info message with event and error").
        SetError(errors.New("This is also another error")).
        Log()

logger.ErrorBuilder().
        SetEvent("my_event").
        SetMessage("This is a log builder info message with event, error, data and context").
        SetError(errors.New("This is also another error")).
        AddContext("requestId", uuid.New()).
        AddData("userId", uuid.New()).
        Log()
```

This produces output like this:

```json
{"time":"2016-01-08T17:45:35.895560313-08:00","name":"log","level":"debug","data":{"message":"This is a log builder debug message"},"context":{"host":"Mac-Pro.local","processId":"16358","logger":"examples.main"},"id":"e4c0f58d-74c1-425e-8f8c-017f03bc0171","version":"0"}
{"time":"2016-01-08T17:45:35.895584789-08:00","name":"my_event","level":"info","data":{"message":"This is a log builder info message with event"},"context":{"host":"Mac-Pro.local","processId":"16358","logger":"examples.main"},"id":"5f7acb89-c498-4b30-b0d4-248de0d8e060","version":"0"}
{"time":"2016-01-08T17:45:35.895611498-08:00","name":"my_event","level":"warn","data":{"message":"This is a warn builder info message with event and error"},"context":{"host":"Mac-Pro.local","processId":"16358","logger":"examples.main"},"exception":{"type":"error","message":"This is also another error","backtrace":[]},"id":"b75fd67c-2831-4aff-8dff-277393da6eed","version":"0"}
{"time":"2016-01-08T17:45:35.895643617-08:00","name":"my_event","level":"crit","data":{"message":"This is a log builder info message with event, error, data and context","userId":"bb486dfd-d7c5-4e3f-8391-c39d9fee6cac"},"context":{"requestId":"3186ea94-bca3-4a75-8ba2-b01151e9935c","host":"Mac-Pro.local","processId":"16358","logger":"examples.main"},"exception":{"type":"error","message":"This is also another error","backtrace":[]},"id":"67c13e4d-12de-4ae4-8606-271d6e4ae13f","version":"0"}
```

For more examples please see [performance.go](performance/performance.go).

Performance
-----------

Some very non-scientific relative benchmarking was performed against the Arpnetworking LogbackSteno Java implementation.

Go with error, mind you it's effectively just a string with no stack trace:

```
Elapsed 2.236446 seconds
Elapsed 2.198458 seconds
Elapsed 2.356763 seconds
Elapsed 2.196019 seconds
Elapsed 2.206734 seconds
```

Average = 2.238884 seconds

Java with exception for error, which includes the message and a stack trace:

```
Elapsed 10.815678 seconds
Elapsed 11.896151 seconds
Elapsed 10.839127 seconds
Elapsed 11.803035 seconds
Elapsed 10.903178 seconds
```

Average = 11.2514338 seconds

Java without exception for error:

```
Elapsed 2.790097 seconds
Elapsed 1.639426 seconds
Elapsed 1.470144 seconds
Elapsed 1.612253 seconds
Elapsed 1.575005 seconds
```

Average = 1.817385 seconds

The gosteno implementation is 5.025 times _faster_ when exceptions are logged in Java and 1.232 times _slower_ when they are not.
The cost of stack trace generation and encoding in Java is not surprising; however, the scale of the improvement with it disabled is.
Finally, although this is neither a complete nor rigorous performance test it does provide an interesting baseline.
The next step would be to compare serialization times including complex data structures as that will simulate real world usage more closely.

Details:
* JDK/JRE version 1.8.0_66
* GO version 1.5.2
* 3.5 Ghz 6-Core Intel Xeon E5
* 32 GB 1866 Mhz DDR ECC
* Mac OS X with El Capitan
* Loggers configured to stdout redirected in bash to /dev/null

Development
-----------

To build the library locally you must satisfy these prerequisites:
* [Go](https://golang.org/)

Next, fork the repository, get and build:

Getting and Building:

```bash
go> go get github.com/$USER/gosteno
go> go install github.com/$USER/gosteno
```

Testing:

```bash
go> go test -coverprofile=coverage.out github.com/$USER/gosteno
go> go tool cover -html=coverage.out
```

To use the local forked version in your project simply import it:

```go
import "github.com/$USER/gosteno"
```

_Note:_ The above assumes $USER is the name of your Github organization containing the fork.

License
-------

Published under Apache Software License 2.0, see LICENSE


&copy; Ville Koskela, 2016
