# idem

[![License](https://img.shields.io/github/license/gonvenience/idem.svg)](https://github.com/gonvenience/idem/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gonvenience/idem)](https://goreportcard.com/report/github.com/gonvenience/idem)
[![Tests](https://github.com/gonvenience/idem/workflows/Tests/badge.svg)](https://github.com/gonvenience/idem/actions?query=workflow%3A%22Tests%22)
[![Codecov](https://img.shields.io/codecov/c/github/gonvenience/idem/main.svg)](https://codecov.io/gh/gonvenience/idem)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gonvenience/idem)](https://pkg.go.dev/github.com/gonvenience/idem)
[![Release](https://img.shields.io/github/release/gonvenience/idem.svg)](https://github.com/gonvenience/idem/releases/latest) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gonvenience/idem)

Go library to detect a entry renaming given list with deleted and list with added items.

The code for this library is a [spork](https://twitter.com/jonjonsonjr/status/1372415386516262915) of code from [go-git/go-git](https://github.com/go-git/go-git), which was modified to match the needs of [`dyff`](https://github.com/homeport/dyff). Originally envisioned to be in `dyff` with [PR #409](https://github.com/homeport/dyff/pull/409), it was moved into its separate repository, due to its complexity and setup.
