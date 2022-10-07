
# Friendly Go

[![License](https://img.shields.io/github/license/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/blob/main/go.mod)
[![Lines](https://img.shields.io/tokei/lines/github/alexcoder04/friendly?label=lines)](https://github.com/alexcoder04/friendly/pulse)
[![Release](https://img.shields.io/github/v/release/alexcoder04/friendly?display_name=tag&sort=semver)](https://github.com/alexcoder04/friendly/releases/latest)
[![Stars](https://img.shields.io/github/stars/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/stargazers)
[![Contributors](https://img.shields.io/github/contributors-anon/alexcoder04/friendly)](https://github.com/alexcoder04/friendly/graphs/contributors)

A small library with functions I use all the time in my projects.

Also useful if you are starting to learn Go and are annoyed by implementing
basic things in your project all the time.

## Install and Use

In your project directory, type

```sh
go get github.com/alexcoder04/friendly
```

And then, in your code

```go
package ...

import (
    ...
    "github.com/alexcoder04/friendly"
)

...
if friendly.IsDir(path) {
    friendly.CompressFolder(path, destination)
}
...
```

## Functions

### `github.com/alexcoder04/friendly`

```go
// files
func WriteNewFile(file string, content string) error { }
func WriteLines(file string, lines []string) error { }
func ReadLines(file string) ([]string, error) { }
func CopyFile(source string, destin string) error { }
func CopyFolder(source string, destination string) error { }
func GetConfigDir(program string) (string, error) { } // creates the directory if it doesn't exist
func GetLogDir(program string) (string, error) { } // creates the directory if it doesn't exist

// os
func Getpwd() string { }
func Run(command string, arguments []string, workingDir string) error { }
func Exists(path string) bool { }
func IsFile(path string) bool { }
func IsDir(path string) bool { }
func ListFilesRecursively(folder string) ([]string, error) { }

// zip
func UncompressFolder(source string, destination string) error { }
func CompressFolder(folder string, destination string) error { }

// misc
func ArrayContains(arr []interface{}, value interface{}) bool { }
func SemVersionGreater(v1 string, v2 string) bool { }
func DownloadFile(downloadUrl string, path string) error { }
```

### `github.com/alexcoder04/friendly/linux`

```go
// desktop
func GetDisplayServer() string { }
func GuiRunning() bool { }
```

## Contributing

If you use this library and are missing some feature - don't hesiatate to open a
pull request or an issue, I'm always looking forward to improve this project!

