EOS 1V Tagger
=============

Intro
-----

Here's Golang library and CLI binary to parse CSV's from Canon ES-E1 software to allow EXIF tagging film scans according to metadata from such CSV's.

Development status: Just started

Usage
-----

```shell
Usage: tagger [OPTIONS] file.csv

Options:
  -filename-pattern string
        filename pattern for generate exiftool command. %d means frame number on the film (default "FILM_%05d.dng")
  -help
        display help message
  -timezone string
        location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow') (default "UTC")

Version: 0.1-alpha.5, build with go1.13.1 at 2019-10-11T00:01:11+03:00
```

NOTICES
-------

* It's in **deep alpha** state
* It's **NOT** going to make any changes to real data: just prints exiftool commands to STDOUT
* It relies on the data provided by ES-E1 software in CSV format(in EOS 1V Memory just export via `File` -> `Export` -> `CSV`)
* **Always** carefully review exiftool commands *before* applying them.
* It sets ISO for each frame from the film settings. So if you have set ISO for particular frame to the another value it will still use the one from the film properties.
* It allows you to specify timezone set on EOS 1V to properly timestamp scans so please pay attention to `-timezone` flag **which defaults to UTC timezone**
* Since tagger releases are built on the most recent version of Golang it cannot run on Windows prior to Windows 7. Technically, it should compile fine on Golang v1.10 or earlier but it's unstested and not a part of release procedure. More details could be found at [Golang v1.11 changelog](https://golang.org/doc/go1.11).

Licence
-------

This software is licenced under GPLv2 terms.
