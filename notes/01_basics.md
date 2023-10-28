# Go Programming Language Guide

## Basics

### Basic Rules

- We use `.go` files
- Code blocks in `{}`
- No styling freedom
- We have semi-colon to separate sentences (optional)
- Case-sensitive
- Strongly typed
- NOT an object-oriented language
- No classes, No exceptions
- We have one file acting as the entry point with main function
- A folder is a package
- Packages can have simple names (services) or URLs (github.com/name/my-library)
- Within one go file we can have: Variables, Functions, Type declarations, Method declarations

## Modules and CLI

- A module is a group of packages
- It's our project
- It contains a go.mod file with configuration and metadata

### CLI Commands

```bash
go mod init
go build
go run
go test
go get
```

#### go mod

Go mod provides access to operations on modules.

```bash
go mod <command> [arguments]
```

Commands

- **download**: download modules to local cache
- **edit**: edit go.mod from tools or scripts
- **graph**: print module requirement graph
- **init**: initialize new module in current directory
- **tidy**: add missing and remove unused modules
- **vendor**: make vendored copy of dependencies
- **verify**: verify dependencies have expected content
- **why**: explain why packages or modules are needed

#### packages

Many commands apply to a set of packages:

```bash
go action [packages]
```

Special Package Names

- "main"
- "all"
- "std"
- "cmd"

#### go run

Usage:

```bash
go run [build flags] [-exec xprog] package [arguments...]
```

#### go build

Usage:

```bash
go build [-o output] [build flags] [packages]
```

**Build Flags**

- **-a**: force rebuilding of packages that are already up-to-date.
- **-n**: print the commands but do not run them.
- **-p n**: the number of programs that can be run in parallel.
- **-race**: enable data race detection.
- **-msan**: enable interoperation with memory sanitizer.
- **-asan**: enable interoperation with address sanitizer.
- **-v**: print the names of packages as they are compiled.
- **-work**: print the name of the temporary work directory.
- **-x**: print the commands.
