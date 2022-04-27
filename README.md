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
