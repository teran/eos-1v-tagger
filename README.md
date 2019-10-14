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
  -geotag string
      GPS track log file to set location data, supported formats are the ones supported by exiftool. Please refer to exiftool docs for details.
  -help
      display help message
  -timezone string
      location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow') (default "UTC")

Version: 0.1-alpha.6, build with go1.13.1 at 2019-10-14T00:01:11+03:00
```

So you could try to do the following:

```shell
tagger -timezone='Europe/Berlin' -filename-pattern='FILM138_%05d.tiff' Downloads/138.csv
```

... and it will print you something like this:

```shell
exiftool -overwrite_original -FocusMode="One-Shot AF" -FNumber="1.4" -ApertureValue="1.4" -FocalLength="35mm" -ISO="400" -ISOSpeed="400" -MeteringMode="Evaluative" -ShootingMode="Aperture-priority AE" -DateTimeOriginal="2019-09-21T14:04:21+02:00" -ModifyDate="2019-09-21T14:04:21+02:00" -ExposureTime="1/8000" -ShutterSpeedValue="1/8000" "FILM138_00001.tiff"
exiftool -overwrite_original -FocusMode="One-Shot AF" -FNumber="2.5" -ApertureValue="2.5" -FocalLength="35mm" -ISO="400" -ISOSpeed="400" -MeteringMode="Evaluative" -ShootingMode="Aperture-priority AE" -DateTimeOriginal="2019-09-21T14:05:40+02:00" -ModifyDate="2019-09-21T14:05:40+02:00" -ExposureTime="1/3200" -ShutterSpeedValue="1/3200" "FILM138_00002.tiff"
..........
```

The tags list generated with exiftool depends on the data present in CSV.
Then you could check carefully the data provided and run these command to apply EXIF metadata.

*Please note:* you need to have [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) installed on your system.


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
