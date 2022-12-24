
# Friendly Go

[![License](https://img.shields.io/github/license/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/blob/main/go.mod)
[![Lines](https://img.shields.io/tokei/lines/github/alexcoder04/friendly?label=lines)](https://github.com/alexcoder04/friendly/pulse)
[![Release](https://img.shields.io/github/v/release/alexcoder04/friendly?display_name=tag&sort=semver)](https://github.com/alexcoder04/friendly/releases/latest)
[![Stars](https://img.shields.io/github/stars/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/stargazers)
[![Contributors](https://img.shields.io/github/contributors-anon/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/graphs/contributors)

A library with functions I use all the time in my projects.

Also useful if you are starting to learn Go and are annoyed by implementing basic things in your project all the time.

## Install and Use

**Warning:** the `v1` version is deprecated, please use `friendly/v2` instead.

In your project directory, type

```sh
go get github.com/alexcoder04/friendly/v2
# or `go get github.com/alexcoder04/friendly/v2/...` depending on the sub-package you need
```

And then, in your code

```go
package ...

import (
    ...
    "github.com/alexcoder04/friendly/v2"
    "github.com/alexcoder04/friendly/v2/ffiles"
)

...
folder, _ := friendly.Input()
if ffiles.IsDir(folder) {
    friendly.CompressFolder(folder, destination)
}
...
```

## Documentation

See also the documentation at `pkg.go.dev`:

 - [`alexcoder04/friendly/v2`](https://pkg.go.dev/github.com/alexcoder04/friendly/v2)
 - [`alexcoder04/friendly/v2/ffiles`](https://pkg.go.dev/github.com/alexcoder04/friendly/v2/ffiles)
 - [`alexcoder04/friendly/v2/flinux`](https://pkg.go.dev/github.com/alexcoder04/friendly/v2/flinux)

### `github.com/alexcoder04/friendly/v2`

Please refer to [pkg.go.dev](https://pkg.go.dev/github.com/alexcoder04/friendly/v2) for this section.

### `github.com/alexcoder04/friendly/v2/ffiles`

```go
// exist
func Exists(path string) (bool, error) { } // returns true error when cannot stat file (and error is not os.ErrNotExists)
func IsDir(path string) bool { } // true only when file exists and IS a directory
func IsFile(path string) bool { } // true only when file exists and is NOT a directory

// read-write
func ListFilesRecursively(folder string) ([]string, error) { }
func ReadLines(file string) ([]string, error) { }
func WriteLines(file string, lines []string) error { }
func WriteNewFile(file string, content string) error { }

// copy
func CopyFile(source string, destin string) error { }
func CopyFolder(source string, destination string) error { }

// locations
func GetConfigDirFor(program string) (string, error) { } // creates the directory if it doesn't exist
func GetCacheDirFor(program string) (string, error) { } // creates the directory if it doesn't exist
func GetLogDirFor(program string) (string, error) { } // creates the directory if it doesn't exist
func GetRuntimeDir() string { } // creates the directory if it doesn't exist
```

### `github.com/alexcoder04/friendly/v2/flinux`

```go
// desktop
func GetDisplayServer() string { }
func GuiRunning() bool { }
```

## Contributing

If you use this library and are missing some feature - don't hesitate to open a
pull request or an issue, I'm always looking forward to improve this project!
