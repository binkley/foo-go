# foo-go

Go module experiment

## Building

`go build` - local binary, try with `./foo-go` in project root

### Linting

`./install-golangci-lint-local.sh` - installs linter in `./bin`

### Batect

`./batect build` - checks build success (local binary not created)
`./batect -T` - lists tasks

Note: Troubles when running on MacOS for `lint` task: Running the local
install script pulls down an OSX binary; Batect needs a Linux binary.

## Credits

* https://github.com/golang/go/wiki/Modules
* https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial
