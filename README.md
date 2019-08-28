# fops
A command-line application `fops` written in Golang

[![Build Status](https://travis-ci.com/yufei-peng/fops.svg?branch=master)](https://travis-ci.com/yufei-peng/fops)



## Table of Contents
- Overview (#overview)
  * 3rd-party libraries used in this application
_ Installing (#installing)
- Getting Started (#getting-started)
  * File lines' counting
  * File's checksum
  * Version checking
  * Help
- Features to be implemented (#features-to-be-implemented)
- Known issues (#known-issues)

## Overview
Fops is a command-line application for file operations.

Fops provides:
- Counting line number of the file.
- Calculating checksum of the file, with either md5, sha1 or sha256.

### 3rd-party libraries used in this application
- Cobra
- Packr

## Installing
Use `go get` to install the latest version of the application.
```
go get -u github.com/yufei-peng/fops
```

## Getting Started
### File lines' counting
Using `fops linecount` with flag `-f` or `--file`
```
fops linecount -f filename.txt
fops linecount --file filename.txt
```

### File's checksum
Using `fops checksum` with flag `-f` or `--file` and one of the md5, sha1 and sha256 algorithm.
```
fops checksum -f filename --md5
fops checksum -f filename --sha1
fops checksum -f filename --sha256
```

### Version checking
Using `fops version` to check latest version of fops.
```
fops version
```

### Help
Using `-h` or `--help` in `fops` and its subcommand, and the application will give some help information.
```
fops --help
fops checksum -h
```

## Features to be implemented
- Tests of the application
- Integrate with CI tool for automatic build

## Known issues
