
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

In your project directory, type

```sh
go get github.com/alexcoder04/friendly
# or `go get github.com/alexcoder04/friendly/...` depending on the sub-package you need
```

And then, in your code

```go
package ...

import (
    ...
    "github.com/alexcoder04/friendly"
    "github.com/alexcoder04/friendly/ffiles"
)

...
folder, _ := friendly.Input()
if ffiles.IsDir(folder) {
    friendly.CompressFolder(folder, destination)
}
...
```

## Functions in Sub-Ppackages

### `github.com/alexcoder04/friendly`

```go
// io
func Input(prompt string) (string, error) { }

// net
func DownloadFile(downloadUrl string, path string) error { }

// os
func Getpwd() string { }
func Run(commandline []string, workingDir string) error { } // pass "" for workingDir to use current working dir
func GetOutput(commandLine []string, workingDir string) (string, error) { } // pass "" for workingDir to use current working dir

// zip
func UncompressFolder(source string, destination string) error { }
func CompressFolder(folder string, destination string) error { }

// misc
func ArrayContains[T comparable](arr []T, value T) bool { }
func IsInt(num string) bool { }
func SemVersionGreater(v1 string, v2 string) bool { }
```

### `github.com/alexcoder04/friendly/ffiles`

```go
// exist
func Exists(path string) (bool, error) { } // returns true error when cannot stat file (and error is not os.ErrNotExists)
func IsFile(path string) bool { } // true only when file exists and is NOT a directory
func IsDir(path string) bool { } // true only when file exists and IS a directory

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
func GetLogDirFor(program string) (string, error) { } // creates the directory if it doesn't exist
func GetRuntimeDir() string { } // creates the directory if it doesn't exist
```

### `github.com/alexcoder04/friendly/flinux`

```go
// desktop
func GetDisplayServer() string { }
func GuiRunning() bool { }
```

## Contributing

If you use this library and are missing some feature - don't hesitate to open a
pull request or an issue, I'm always looking forward to improve this project!
