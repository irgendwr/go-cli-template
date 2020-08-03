# go-cli-template

<!-- FIXME: replace 'irgendwr/go-cli-template' with your repository name -->
[![Build status](https://github.com/irgendwr/go-cli-template/workflows/build/badge.svg)](https://github.com/irgendwr/go-cli-template/actions?query=workflow%3Abuild)
[![Release status](https://github.com/irgendwr/go-cli-template/workflows/release/badge.svg)](https://github.com/irgendwr/go-cli-template/actions?query=workflow%3Arelease)
[![Go Report Card](https://goreportcard.com/badge/github.com/irgendwr/go-cli-template)](https://goreportcard.com/report/github.com/irgendwr/go-cli-template)
[![GitHub Release](https://img.shields.io/github/release/irgendwr/go-cli-template.svg)](https://github.com/irgendwr/go-cli-template/releases)

<!-- FIXME: add short description here -->
This is a template for building CLI tools in go.

To use it, simply clone this repository and replace all occurrences of `irgendwr/go-cli-template` with your repository and `APPNAME` with your application name. You need to adapt this in all files, as indicated by comments containing `FIXME`. Furthermore, change the LICENSE to fit your needs.

Features:

- based on [cobra](https://github.com/spf13/cobra) and its template
- config parsing using [viper](https://github.com/spf13/viper)
- Makefile (build, test, clean, update)
- GitHub actions (build, release)
- build/release using [goreleaser](https://github.com/goreleaser/goreleaser/)

Read the [cobra readme](https://github.com/spf13/cobra#readme) for more details and guides to expand functionality.

## Installation

### Linux

<!-- FIXME: replace 'APPNAME' with your application name -->
Download the latest release into `/opt/APPNAME/` (or any other folder):

<!-- FIXME: replace 'APPNAME' with your application name and 'irgendwr/go-cli-template' with your repository name -->
```bash
mkdir -p /opt/APPNAME/
cd /opt/APPNAME/
curl -O -L https://github.com/irgendwr/go-cli-template/releases/latest/download/APPNAME_Linux_x86_64.tar.gz
tar -xvzf APPNAME_Linux_x86_64.tar.gz
```

<!-- FIXME: replace 'APPNAME' with your application name -->
Create a file called `.APPNAME.yaml` inside this folder (e.g. using `nano .APPNAME.yaml`) and edit it to fit your needs.
See [config](#config) section for examples.

### Config

**Note: Do not use Tabs! Indent config with spaces instead.**

<!-- FIXME: document options -->
<!-- FIXME: add examples -->
Example:

```yaml
Username: foobar
Password: your-password-here
```

## Usage

<!-- FIXME: replace 'APPNAME' with your application name -->
Run `/opt/APPNAME/APPNAME`.
<!-- FIXME: add syntax, explanations, ... -->

## Build

Run `make`.
