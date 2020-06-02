[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## GTFO

Reimplementation of my tool [gtfo](https://github.com/mzfr/gtfo) in Go.

I'm mostly doing this as a means of learning Go as it seemed like a nice first project to start with. Also, it's much easier to make binaries from Go scripts.

## Gallery

* `gtfo -b nmap`

![](images/gtfo.png)

* `gtfo -b randomnamehere`

![](images/err.png)

## Usage

I decided to keep this plain and simple so there is only 1 option now, to search the gtfobins. I haven't implemented lolbas searching functionality because I never used that in the python version	of gtfo and I thought that might just be a waste of time. So if you would like to have that supported please open an issue.

```
Search gtfobin from terminal

Options:
  -b, --bin <binary>       Search Linux binaries on gtfobins
```

## Installation

You can download the pre-compiled binary from [here](https://github.com/mzfr/go-gtfo/releases)

If you want to make changes to the code and then compile the binary you can clone this repo and then run:

```
go build
```

Also, you can run the following command to install it directly:

```
go get github.com/mzfr/go-gtfo
```

If you want to run this locally then do the following:

1) Clone this repo: `git clone https://github.com/mzfr/go-gtfo`
2) run: `go run main.go -b <binary_name>`

__Note__: Make sure you have go installed.

## Support

If you'd like you can buy me some coffee:

<a href="https://www.buymeacoffee.com/mzfr" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" ></a>
