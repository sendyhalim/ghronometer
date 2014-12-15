#Ghronometer
Just a simple tool to count elapsed time for a (script|command line tool) to finish

##Installation
Fire up terminal, run

```
go get github.com/sendyHalim/ghronometer
```

##Example Usage
```
ghronometer composer update
```

```
ghronometer composer self-update
```

```
ghronometer curl google.com -I
```

This package just add elapsed time in ms after each command finished :), feel free to contribute.

##License
Ghronometer is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)