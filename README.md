log2 is a little wrapper on top of the original [log package](http://golang.org/pkg/log/) to provide the log level feature, eg:

    log2.Info("Hello world")
    log2.Errorf("Download failed because: %s", err)
    log2.Warning("Something went wrong")
    log2.Fatal(err)

In addition, the `log2.Fatal` function exits after printing the stack trace of the _current_ goroutine.

Finally, level prefixes are formatted as `<LEVEL> -- <msg>` - solely to match the default format of [vcap_logging](http://rubygems.org/gems/vcap_logging). Ideally it should be made configurable using the `text/template` package.