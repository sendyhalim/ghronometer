#Ghronometer
Just a simple tool to count elapsed time for a (script|command line tool) to finish
[![Build Status](https://travis-ci.org/sendyHalim/ghronometer.svg)](https://travis-ci.org/sendyHalim/ghronometer)
##Installation
Fire up terminal, run

```sh
go get github.com/sendyHalim/ghronometer
```

##Example Usage
```sh
ghronometer composer update
```

```sh
ghronometer composer self-update
```

```sh
ghronometer curl google.com -I
```

This package adds elapsed time in `ms` after each command has finished :), feel free to contribute.

##License
Ghronometer is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
