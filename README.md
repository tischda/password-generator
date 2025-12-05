[![Build Status](https://github.com/tischda/password-generator/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/password-generator/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/password-generator/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/password-generator/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/password-generator/badge.svg)](https://coveralls.io/r/tischda/password-generator)
[![Linter Status](https://github.com/tischda/password-generator/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/password-generator/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/password-generator)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/password-generator.svg)](https://github.com/tischda/password-generator/releases/latest)

# password-generator

Password generator written in Golang + Fyne.
When using the GUI, the password is automaticaly copied to the clipboard.

## Usage

~~~
Usage: password-generator.exe [OPTIONS]

 OPTIONS:
  -gui
        show GUI
  -help
        displays this help message
  -length int
        password length between 1 and 128 (default 20)
  -version
        print version and exit
~~~

## Examples

~~~
$ password-generator
xhGb){dl0lsbk_T6dZDv

$ password-generator --length 15
dtB{#hP_C)C^D4a
~~~

Inspired from:

* https://github.com/gocoder-ai/password-generator
* https://www.socketloop.com/tutorials/golang-how-to-generate-random-string

## Compile

This project has a dependency on [Fyne](https://github.com/fyne-io/fyne).
You need to install these additional packages:

Alpine:
~~~
sudo apk add mesa mesa-dev mesa-dri-nouveau libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers
~~~

Ubuntu:
~~~
sudo apt-get update && sudo apt-get install -y xorg-dev mesa-utils libgl1 libgl1-mesa-dev
~~~

CentOS 8:
~~~
sudo yum install -y mesa-libGL mesa-libGL-devel mesa-dri-drivers libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libXxf86vm-devel make golang
~~~

CentOS 9:
~~~
sudo yum install -y yum-utils
sudo yum-config-manager --set-enabled crb
sudo yum install mesa-dri-drivers mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel libXxf86vm-devel
~~~

libXxf86vm-devel is only available in the CodeReady Linux Builder (CRB) repository.

Windows:
~~~
winget install --id BrechtSanders.WinLibs.POSIX.UCRT

set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64 

make build
~~~

The first build take several minutes, be patient!

I couldn't get it to work with fyne-cross (no console output).

