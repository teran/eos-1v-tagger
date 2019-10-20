EOS 1V Tagger
=============

Intro
-----

Here's Golang library and CLI binary to parse CSV's from Canon ES-E1 software to allow EXIF tagging film scans according to metadata from such CSV's.

**Development status:** Just started

What it does
------------

CLI tool called tagger generates [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) commands with metadata right from Canon ES-E1 (via CSV export).
This allows to set exposure, focal length, ISO, timestamp, and much more, even GPS data if you have recorded GPS track log via smartphone, tracker or whatever else allows to write GPS track log.

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
  -make string
        Make tag value. NOTE: it will overwrite the value set by your film scanner software
  -model string
        Model tag value. NOTE: it will overwrite the value set by your film scanner software
  -serial-number string
        SerialNumber tag value. NOTE: it will overwrite the value set by your film scanner software
  -set-digitized
        set DateTimeDigitized from CreateDate field
  -timezone string
        location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow') (default "UTC")

Version: 0.1-alpha.8, build with go1.13.1 at 2019-10-17T02:14:02+03:00
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

Some real-life examples
-----------------------

```shell
tagger -geotag ~/Downloads/walk-at-21-09-2019.gpx -make="Ilford Delta" -model="Canon EOS 1V" -serial-number="XXXXX" -timezone='Europe/Moscow' -filename-pattern="FILM139_%05d.dng" -set-digitized ~/Downloads/139.csv
```

This will generate [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) commands to set:

* GPS tags
* Make to `Ilford Delta`
* Camera model to Canon EOS 1V
* Camera serial number to `XXXXX`
* Timestamps: original photo timetamp to the one set in camera aligned to `Europe/Moscow` timezone, digitized to the one present in `CreateDate` tag
* All the data present in CSV like aperture, exposure, ISO, focal length, etc.

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
