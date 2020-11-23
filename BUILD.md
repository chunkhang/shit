# Building

Before building, make sure the repository is cloned into the right path:
`$GOPATH/src/github.com/chunkhang/shit`.

### Versioning

The current version is always reflected in [.version](./.version).
Building with `make` ensures that `shit` is built with the correct
version.

### Build Modes

#### Development

```
make
```

The built binary can be found in `target/shit`.

#### Production

```
make build
```

The built binary can be found in `dist/shit`.

### Logging

```
make log
```

The development build will log to `~/.shit.log` when running. This
command allows us to follow the logs.

### Cleaning up

```
make clean
```

This command cleans up the `target` and `dist` folders.
