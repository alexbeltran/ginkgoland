# ginkgoland

Ginkgoland is a reporter for ginkgo that creates a `go test` format output. This output can be parsed by `go tool test2json` which is used by ginkgoland.

To use it, add the reporter to the list of custom reporters in your suite file. With this everything should be in place.

```
	RunSpecsWithCustomReporters(t, "Suite Name", []Reporter{&ginkgoland.Reporter{T: t}})
```

This project uses the reporter so look at the test files as an example.
