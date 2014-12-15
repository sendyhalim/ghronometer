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

Composer could not find a composer.json file in /home/sendy
To initialize a project, please create a composer.json file as described in the http://getcomposer.org/ "Getting Started" section

elapsed time: 58 ms
```

```sh
ghronometer composer self-update

You are already using composer version 523aef76d0fa29c18bf42264e88fc9b58acf825c.

elapsed time: 2326 ms
```

```sh
ghronometer curl google.com -I

<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="http://www.google.com/">here</A>.
</BODY></HTML>

elapsed time: 173 ms
```

This package adds elapsed time in `ms` after each command has finished :), feel free to contribute.

##Root command
If you need to run command with root access(e.g in ubuntu `sudo`):
```
sudo cp ~/go/bin/ghronometer /usr/bin/ghronometer
```
just remember that you need to re-run above command everytime you update `ghronometer` package

##License
Ghronometer is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
