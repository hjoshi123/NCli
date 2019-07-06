<p align="center">
    <img src="https://github.com/hjoshi123/NCli/blob/master/ncli.png">
</p>

# NCli - A simple Command Line Interface to lookup various details of a host like IP, CName

[![NCli](https://img.shields.io/travis/hjoshi123/NCli.svg)](https://travis-ci.org/hjoshi123/NCli) [![Go Report Card](https://goreportcard.com/badge/github.com/hjoshi123/NCli)](https://goreportcard.com/report/github.com/hjoshi123/NCli) ![GitHub](https://img.shields.io/github/license/hjoshi123/NCli.svg) ![GitHub repo size](https://img.shields.io/github/repo-size/hjoshi123/NCli.svg)

```
This is a Simple CLI developed using Golang
It includes lookup for CNAME, IP Address, MX Records and many more features to be added
```

## Installation

* First install golang using homebrew in **mac OS** or **apt** in Ubuntu.
* Once go is set up, run the following command in terminal `go get github.com/hjoshi123/NCli`
* After this, install the module using `go install github.com/hjoshi123/NCli`
* Check if NCli is installed by typing `ncli --help`

### These are following services supported by NCli

- [x]  **NameServers**
- [x]  **IP Lookup**
- [x]  **MX Records**
- [x]  **CNAME**
- [x]  **IP Address of the host**
- [ ] more to come..

## Examples

* **NameServers**

```
ncli ns --host msrit.edu

// output
ns5.he.net.
ns1.he.net.
ns2.he.net.
ns4.he.net.
```
