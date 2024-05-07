[![GitHub version](https://img.shields.io/github/v/release/jeffalyanak/check_sl_swtgw218as)](https://github.com/jeffalyanak/check_sl_swtgw218as/releases/latest)
[![License](https://img.shields.io/github/license/jeffalyanak/check_sl_swtgw218as)](https://github.com/jeffalyanak/check_sl_swtgw218as/blob/master/LICENSE)
[![Donate](https://img.shields.io/badge/donate--green)](https://jeff.alyanak.ca/donate)

# Golang sl-swtgw218as Switch Port Checker

Icinga/Nagios plugin, checks for bad packets on any interfaces.

## Installation and requirements

The pre-compiled binaries available on the [releases page](https://github.com/jeffalyanak/check_sl_swtgw218as/releases) are self-contained and have no dependancies to run.

If you wish to compile it yourself, you'll need to install `go`. It's been tested on:

* Golang 1.21.6

It'll probably build just fine on many other versions.

## Usage

```bash
usage:
  required
    -h string
        Fully-qualified domain name to check.
    -u string
        Username.
    -p string
        Password.
```

## Version history

-0.5—Initial release.

## License

Golang Icinga/Nagios sl-swtgw218as Switch Port Checker is licensed under the terms of the GNU General Public License Version 3.
