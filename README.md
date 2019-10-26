# EOS 1V Tagger

## Intro

Here's Golang library and CLI binary to parse CSV's from Canon ES-E1 software to allow EXIF tagging film scans according to metadata from such CSV's.

**Development status:** Just started

## What it does

CLI tool called tagger generates [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) commands with metadata right from Canon ES-E1 (via CSV export).
This allows to set exposure, focal length, ISO, timestamp, and much more, even GPS data if you have recorded GPS track log via smartphone, tracker or whatever else allows to write GPS track log.

## Usage

```shell
Usage: tagger [OPTIONS] file.csv

Options:
  -exiftool-binary string
        path to exiftool binary (default "exiftool")
  -filename-pattern string
        filename pattern for generate exiftool command. Available variables: frameNo, cameraID, filmID. More details are available in README. (default "FILM${filmID:d}_${frameNo:05d}.dng")
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
  -timestamp-format string
        the timestamp format in the locale your're using on the system with ES-E1 software. Allowed values: 'US', 'EU' (default "US")
  -timezone string
        location or timezone name used while setting time on EOS 1V, will be used for proper scans timestamping (example: 'Europe/Moscow') (default "UTC")

Version: undefined, build with go1.13.3 at 1970-01-01T03:00:00+03:00
```

So you could try to do the following:

```shell
tagger -timezone='Europe/Berlin' -filename-pattern='FILM_${cameraID:02d}_${filmID:03d}_${frameNo:05d}.tiff' Downloads/138.csv
```

... and it will print you something like this:

```shell
exiftool -overwrite_original -FocusMode="One-Shot AF" -FNumber="1.4" -ApertureValue="1.4" -FocalLength="35mm" -ISO="400" -ISOSpeed="400" -MeteringMode="Evaluative" -ShootingMode="Aperture-priority AE" -DateTimeOriginal="2019-09-21T14:04:21+02:00" -ModifyDate="2019-09-21T14:04:21+02:00" -ExposureTime="1/8000" -ShutterSpeedValue="1/8000" "FILM01_138_00001.tiff"
exiftool -overwrite_original -FocusMode="One-Shot AF" -FNumber="2.5" -ApertureValue="2.5" -FocalLength="35mm" -ISO="400" -ISOSpeed="400" -MeteringMode="Evaluative" -ShootingMode="Aperture-priority AE" -DateTimeOriginal="2019-09-21T14:05:40+02:00" -ModifyDate="2019-09-21T14:05:40+02:00" -ExposureTime="1/3200" -ShutterSpeedValue="1/3200" "FILM01_138_00002.tiff"
..........
```

The tags list generated with exiftool depends on the data present in CSV.
Then you could check carefully the data provided and run these command to apply EXIF metadata.

*Please note:* you need to have [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) installed on your system.

### filename-pattern formatting

Since Go doesn't have Python-style named formatting tools I had to implement my own for purpose of named variable substitution with advance of formatting.
Currently tagger have available three variables:

* `filmID`
* `cameraID`
* `frameNo`

You can use them as follows: `${filmID:d}` where `d` stands for 10-base digit as described in [Go's fmt package docs](https://golang.org/pkg/fmt/#pkg-overview).
Unfortunately specifying type is required to allow formatting features like `.2f`, which means 2-digits float precision or `05d` which means 5 number digit with 10-base integer with leading zeros.

## Some real-life examples

```shell
tagger -geotag ~/Downloads/walk-at-21-09-2019.gpx -make="Ilford Delta" -model="Canon EOS 1V" -serial-number="XXXXX" -timezone='Europe/Moscow' -filename-pattern='FILM_${cameraID:02d}_${filmID:03d}_${frameNo:05d}.tiff' -set-digitized ~/Downloads/139.csv
```

This will generate [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) commands to set:

* GPS tags
* Make to `Ilford Delta`
* Camera model to Canon EOS 1V
* Camera serial number to `XXXXX`
* Timestamps: original photo timetamp to the one set in camera aligned to `Europe/Moscow` timezone, digitized to the one present in `CreateDate` tag
* All the data present in CSV like aperture, exposure, ISO, focal length, etc.

## Build & test

Since tagger is written in Go the compiler is required and could installed from [Go official website](https://golang.org).

Build tagger from source in the most trivial case:

```shell
go build -o ./tagger ./cmd/tagger/...
```

To run tests(if changes made, for instance):

```shell
go test ./...
```

## NOTICES

* It's in **deep alpha** state
* It's **NOT** going to make any changes to real data: just prints exiftool commands to STDOUT
* It relies on the data provided by ES-E1 software in CSV format(in EOS 1V Memory just export via `Data` -> `Export` -> `CSV`)
* Since ES-E1 doesn't mark timezone and uses local regional settings there's no way to determine which date format (dd/mm/yyyy or mm/dd/yyyy) was used on CSV export, please use `-timestamp-format` option to let tagger know if EU timestamp format is set in regional settings.
* **Always** carefully review exiftool commands *before* applying them.
* It allows you to specify timezone set on EOS 1V to properly timestamp scans so please pay attention to `-timezone` flag **which defaults to UTC timezone**
* Since tagger releases are built on the most recent version of Golang it cannot run on Windows prior to Windows 7. Technically, it should compile fine on Golang v1.10 or earlier but it's unstested and not a part of release procedure. More details could be found at [Golang v1.11 changelog](https://golang.org/doc/go1.11).

## Licence

This software is licenced under GPLv2 terms.
