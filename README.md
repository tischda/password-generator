# password-generator [![Test](https://github.com/tischda/password-generator/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/password-generator/actions/workflows/test.yml)

Password generator written in [Go](https://www.golang.org).

### Usage

~~~
Usage: password-generator.exe [option]

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

Examples:

~~~
$ password-generator
xhGb){dl0lsbk_T6dZDv

$ password-generator --length 15
dtB{#hP_C)C^D4a
~~~

Inspired from:

* https://github.com/gocoder-ai/password-generator
* https://www.socketloop.com/tutorials/golang-how-to-generate-random-string

### Compile:

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
