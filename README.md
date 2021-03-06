**wego** is a weather client for the terminal.
## A fork of wego with Dark Sky support

![Screenshots](https://raw.githubusercontent.com/tylerwolf35/wego/gh-pages/wego.png)

## Features

* show forecast for 1 to 7 days
* nice ASCII art icons
* displayed info (metric or imperial units):
  * temperature range ([felt](https://en.wikipedia.org/wiki/Wind_chill) and measured)
  * windspeed and direction
  * viewing distance
  * precipitation amount and probability
* ssl, so the NSA has a harder time learning where you live or plan to go
* multi language support
* config file for default location which can be overridden by commandline
* Automatic config management with [ingo](https://github.com/schachmat/ingo)

## Dependencies

* A [working](https://golang.org/doc/install#testing) [Go](https://golang.org/) environment >= 1.5 (You can use
  [goenv](https://github.com/syndbg/goenv) if your distribution does not
  support Go >= 1.5 yet)
* utf-8 terminal with 256 colors
* A sane monospaced font containing all the required runes
* An API key for the backend (see Setup below)

## Installation

To install or update the wego binary into your `$GOPATH` as usual, run:
```shell
go get -u github.com/tylerwolf35/wego
```
If you use Arch Linux or an Arch based distribution you can install this fork of wego from the [AUR](https://aur.archlinux.org/packages/wego-darksky) using the following command (edit for use with your AUR helper):
```shell
pikaur -S wego-darksky
```

## Setup

0. Run `wego` once. You will get an error message, but the `.wegorc` config file
   will be generated in your `$HOME` directory (it will be hidden in some file
   managers due to the filename starting with a dot). If the command is not found put `alias wego=/home/$USER/go/bin/wego` (or wherever your go path is) in your shell's rc file.
0. __With a [Dark Sky](https://darksky.net) account__ (new default)
    * Create your account at https://darksky.net/dev
    * Update the following `.wegorc` config variables to fit your needs:
    ```
      backend=darksky
      location=37.4275,-122.1697
      darksky-api-key=YOUR_DARKSKY_API_KEY_HERE
    ```
0. __With an [Openweathermap](https://home.openweathermap.org/) account__
    * You can create an account and get a free API key by [signing up](https://home.openweathermap.org/users/sign_up)
    * Update the following `.wegorc` config variables to fit your needs:
    ```
      backend=openweathermap
      location=Palo Alto
      owm-api-key=YOUR_OPENWEATHERMAP_API_KEY_HERE
    ```
0. __With a [Worldweatheronline](http://www.worldweatheronline.com/) account__
    * Worldweatheronline no longer gives out free API keys. [#83](https://github.com/schachmat/wego/issues/83)
    * Update the following `.wegorc` config variables to fit your needs:
    ```
      backend=worldweatheronline
      location=Palo Alto
      wwo-api-key=YOUR_WORLDWEATHERONLINE_API_KEY_HERE
    ```
0. You may want to adjust other preferences like `days`, `units` and `…-lang` as
   well. Save the file.
0. Run `wego` once again and you should get the weather forecast for the current
   and next few days for your chosen location.
0. If you're visiting someone in e.g. London over the weekend, just run `wego 4
   London` or `wego London 4` (the ordering of arguments makes no difference) to
   get the forecast for the current and the next 3 days. Unfortunately that does
   not currently work with the forecast.io backend, as it only supports
   latitude,longitude location specification.

You can set the `$WEGORC` environment variable to override the default config
file location.

## License - ISC

Copyright (c) 2020, Tyler Wolf

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
