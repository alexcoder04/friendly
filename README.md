
# friendly

A small library with functions I use all the time in my projects.

Also useful if you are starting to learn Go and are struggling with implementing
basic things in your project all the time.

## Functions

```go
// files
func WriteLines(file string, lines []string) error { }

func ReadLines(file string) ([]string, error) { }

func CopyFile(source string, destin string) error { }

func CopyFolder(source string, destination string) error { }

// os
func IsFile(path string) bool { }

func IsDir(path string) bool { }

// zip
func UncompressFolder(source string, destination string) error { }

func CompressFolder(folder string, destination string) error { }

// misc
func ArrayContains(arr []interface{}, value interface{}) bool { }

func SemVersionGreater(v1 string, v2 string) bool { }

func DownloadFile(downloadUrl string, path string) error { }
```

