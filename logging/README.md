# Zerolog

## Installation 

`go get -u github.com/rs/zerolog/log`

## Usage



`log.Info().Msg(...)`

### Different log levels

* panic (`zerolog.PanicLevel`, 5)
* fatal (`zerolog.FatalLevel`, 4)
* error (`zerolog.ErrorLevel`, 3)
* warn (`zerolog.WarnLevel`, 2)
* info (`zerolog.InfoLevel`, 1)
* debug (`zerolog.DebugLevel`, 0)
* trace (`zerolog.TraceLevel`, -1)

### Add context/data to log

```
  log.Debug().
    Str("Scale", "833 cents").
    Float64("Interval", 833.09).
    Msg("Fibonacci is everywhere")
```

### Log errors

```
err := errors.New("seems we have an error here")
log.Error().Err(err).Msg("")
```

By default uses JSON logging, to use human readable logging:

`log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})`
