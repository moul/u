# u

:smile:  Go common utility functions

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/u)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/u/blob/master/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/u.svg)](https://github.com/moul/u/releases)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/u/workflows/Go/badge.svg)](https://github.com/moul/u/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/u/workflows/Release/badge.svg)](https://github.com/moul/u/actions?query=workflow%3ARelease)
[![PR](https://github.com/moul/u/workflows/PR/badge.svg)](https://github.com/moul/u/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/u.svg)](https://golangci.com/r/github.com/moul/u)
[![codecov](https://codecov.io/gh/moul/u/branch/master/graph/badge.svg)](https://codecov.io/gh/moul/u)
[![Go Report Card](https://goreportcard.com/badge/moul.io/u)](https://goreportcard.com/report/moul.io/u)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/u/badge)](https://www.codefactor.io/repository/github/moul/u)

Inspired by https://github.com/kjk/u

## Usage

[embedmd]:# (.tmp/godoc.txt txt /FUNCTIONS/ $)
```txt
FUNCTIONS

func B64Decode(input string) ([]byte, error)
    B64Decode try to decode an input string and returns bytes if success.

func B64Encode(input []byte) string
    B64Encode returns a base64 encoded string of input bytes.

func CaptureStderr() (func() string, error)
    CaptureStderr temporarily pipes os.Stderr into a buffer.

func CaptureStdout() (func() string, error)
    CaptureStdout temporarily pipes os.Stdout into a buffer.

func CaptureStdoutAndStderr() (func() string, error)
    CaptureStdoutAndStderr temporarily pipes os.Stdout and os.Stderr into a
    buffer.

func CheckErr(err error)
    CheckErr panics if the passed error is not nil.

func CombineFuncs(left func(), right ...func()) func()
    CombineFuncs create a chain of functions. This can be particularly useful
    for creating cleanup function progressively. It solves the infinite loop you
    can have when trying to do it manually:
    https://play.golang.org/p/NQem8UJ500t.

func CommandExists(command string) bool
    CommandExists checks whether a command is available in the $PATH.

func CreateEmptyFileWithSize(path string, size uint) error
    CreateEmptyFileWithSize creates a new file of the desired size, filled with
    zeros.

func CurrentUsername(fallback string) string
    CurrentUsename returns the current user's username. If username cannot be
    retrieved, it returns the passed fallback.

func DirExists(path string) bool
    DirExists checks whether a path exists and is a directory.

func ExecStandaloneOutputs(cmd *exec.Cmd) ([]byte, []byte, error)
    ExecStandaloneOutputs runs the command and returns its standard output and
    standard error.

func ExpandUser(path string) (string, error)
func FanIn(chans ...<-chan interface{}) <-chan interface{}
    FanIn merges multiple input chans events into one.

func FileExists(path string) bool
    FileExists checks whether a path exists and is a regular file.

func Future(fn func() (interface{}, error)) <-chan FutureRet
    Future starts running the given function in background and return a chan
    that will return the result of the execution.

func JSON(input interface{}) string
    JSON returns a JSON representation of the passed input.

func MustCaptureStderr() func() string
    MustCaptureStderr wraps CaptureStderr and panics if initialization fails.

func MustCaptureStdout() func() string
    MustCaptureStdout wraps CaptureStdout and panics if initialization fails.

func MustCaptureStdoutAndStderr() func() string
    MustCaptureStdoutAndStderr wraps CaptureStdoutAndStderr and panics if
    initialization fails.

func MustExpandUser(path string) string
    MustExpandUser wraps ExpandUser and panics if initialization fails.

func MustTempFileName(dir, pattern string) string
    MustTempFileName wraps TempFileName and panics if initialization fails.

func MustTempfileWithContent(content []byte) (*os.File, func())
    MustTempfileWithContent wraps TempfileWithContent and panics if
    initialization fails.

func PathExists(path string) bool
    PathExists checks whether a path exists or not.

func PrettyJSON(input interface{}) string
    PrettyJSON returns an indented JSON representation of the passed input.

func RandomLetters(n int) string
    RandomLetters returns a string containing 'n' random letters.

func SafeExec(cmd *exec.Cmd) string
    SafeExec runs a command and return a string containing the combined standard
    output and standard error. If the program fails, the result of `err` is
    appended to the output.

func Sha1(data []byte) []byte
func Sha1Hex(data []byte) string
func ShortDuration(d time.Duration) string
    ShortDuration returns a short human-friendly representation of a duration.
    For duration < 100 days, the output length will be <= 7.

func SilentClose(closer io.Closer)
    SilentClose calls an io.Closer.Close() function and ignore potential errors.

    You can use it as `defer SilenceClose(f)`

func TempFileName(dir, pattern string) (string, error)
    TempFileName returns a valid temporary file name (the file is not created).

func TempfileWithContent(content []byte) (*os.File, func(), error)
    TempfileWithContent creates a tempfile with specified content written in it,
    it also seeks the file pointer so you can read it directly. The second
    returned parameter is a cleanup function that closes and removes the temp
    file.

func UniqueInterfaces(input []interface{}) []interface{}
    UniqueInterfaces removes duplicate values from an interface slice.

func UniqueInts(input []int) []int
    UniqueInts removes duplicate values from an int slice.

func UniqueStrings(input []string) []string
    UniqueStrings removes duplicate values from a string slice.

func Unzip(src string, dest string) ([]string, error)
    Unzip decompresses a zip archive, moving all files and folders within the
    zip file to an output directory. Based on
    https://golangcode.com/unzip-files-in-go/ (MIT).

func UnzipBytes(src []byte, dest string) ([]string, error)
    UnzipBytes is similar to Unzip but takes a zip archive as bytes instead of
    looking for a real file.

func WaitForCtrlC()

TYPES

type FutureRet struct {
	Ret interface{}
	Err error
}
    FutureRet is a generic struct returned by Future.

type MutexMap struct {
	// Has unexported fields.
}
    MutexMap manages a pool of mutexes that can be get by key. MutexMap is
    thread-safe.

func (mm *MutexMap) Lock(key string) func()
    Lock locks a mutex by key, and returns a callback for unlocking unlock. Lock
    will automatically create a new mutex for new keys.

type UniqueChild interface {
	SetChild(childFn func(context.Context))
	CloseChild()
}
    UniqueChild is a goroutine manager (parent) that can only have one child at
    a time. When you call UniqueChild.SetChild(), UniqueChild cancels the
    previous child context (if any), then run a new child. The child needs to
    auto-kill itself when its context is done.

func NewUniqueChild(ctx context.Context) UniqueChild
    NewUniqueChild instantiates and returns a UniqueChild manager.

```

See [![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/u)

## Install

### Using go

```console
$ go get moul.io/u
```

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/master/contribute.gif)

I really welcome contributions. Your input is the most precious material. I'm well aware of that and I thank you in advance. Everyone is encouraged to look at what they can do on their own scale; no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./CONTRIBUTING.md)

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/u/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/u/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/u/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

<!--
### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/u.svg)](https://starchart.cc/moul/u)
-->

## License

© 2020-2021  [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) ([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT) ([`LICENSE-MIT`](LICENSE-MIT)), at your option. See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
