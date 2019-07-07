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

* First install golang using **homebrew** in **mac OS** or **apt** in **Ubuntu**.
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

* **IP Lookup**

```
ncli ip --host msrit.edu

// Output
54.230.216.176
```

* Check out `ncli --help` for more details on various commands. All commands unless mentioned require **--host** option.

## TODO
* Port Scanning


### If this library helps you in anyway, show your love :heart: by putting a :star: on this project :v:

## Contributors
[Hemant Joshi](https://github.com/hjoshi123)


## License

```
Copyright (c) 2019 Hemant Joshi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
