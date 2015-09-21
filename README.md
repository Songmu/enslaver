enslaver
=======

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[license]: https://github.com/Songmu/enslaver/blob/master/LICENSE

Works forever.

## Description

Run a given command forever.

## Synopsis

	enslaver.Command(os.Args[1], os.Args[2:]...).Run()

## Command Line Tool

You can get the command `enslaver` by following.

    % go get github.com/Songmu/enslaver/cmd/enslaver

### usage

    % enslaver perl -E 'while (1) { 0.1 > rand() ? (die "DEAD") : (say "ALIVE!") }'

## Author

[Songmu](https://github.com/Songmu)
