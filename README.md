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
  -file-source string
        adds file source EXIF tag. Available options: 'Film Scanner', 'Reflection Print Scanner', 'Digital Camera'
  -filename-pattern string
        filename pattern for generate exiftool command. Available variables: frameNo, cameraID, filmID. More details are available in README. (default "FILM_${cameraID:02d}${filmID:03d}${frameNo:05d}.dng")
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

Version: undefined, build with go1.13.4 at 1970-01-01T03:00:00+03:00
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
Then you could check carefuly the data provided and run these command to apply EXIF metadata.

*Please note:* you need to have [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) installed on your system.

### The workflow tagger is designed for

#### Making shots

Please always ensure you have proper time set on your EOS 1V. If you have left your camera without battery for some time, please connect it to ES-E1 and set date & time. Otherwise it won't store any date & time data for your shots making impossible to set accurate timestamps and geotaging.

When you making shots to your film camera it's not such a bad idea to run some location tracking software on your smartphone for furter geotagging your shots.
So tagger supports `-geotag` flag to align the photos against GPS track log and set proper coordinates to each one.

This feature relies on [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) geotagging functionality.

#### Scanning

I'm not going to talk about technical part of the process like film washing to avoid dust, dirt and strawberry jam on the film itself, just pointing what is important for tagger instead: scanning your shots in the order they were captured makes no problems on tagging them. Otherwise it will require to rename the scan files or change order in ES-E1 CSV exported data to match the scans properly.

At the same time EOS 1V provides a set of unique data to identify particular frame in global namespace like camera ID, film ID and frame ID. If you have more than one EOS 1V the best option is to use different camera IDs for them and use respective filenaming scheme for scans.

#### Completing the data in ES-E1 & CSV export

ES-E1 software allows to manually complete the data about your shots if it wasn't done automatically by EOS 1V, normally all the data from CSV should properly be handled by tagger, including timestamps, exposure and much-much more. So here should be no issues.

**Note**: most of timestamps in EXIF have no way to use only date part of the timestamp. So timestamps with date field only are ignored. Please keep it in mind when setting frame timestamp manually and leaving time field unchecked/empty in ES-E1.

To export data in CSV format for tagger just use `Data` menu in the main EOS-1V Memory window, choose `Export` -> `CSV`.

#### Tagging the files

After downloading the latest version of tagger from [releases page](https://github.com/teran/eos-1v-tagger/releases) please make sure you installed it in the place it's accessible via terminal application. Then run your terminal (`Terminal.app` in macOS; `Win`+`R`, then type `cmd` in Windows) and run tagger application. It should print help message like in Usage section on run without any option. Please read it carefuly. Choose the ones you need and point CSV export file. If you have any issues please refer to examples in README.

**Hint**: to avoid copy-pasting huge amount of [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) commands feel free to redirect tagger's output to the file: `> exiftool-commands.bat` (`bat` extension in Windows means batch command file which could be run like a shell script). Please check paths carefuly to avoid updating wrong files. And just that file like any other command.
**Note**: tagger uses CSV file alone to generate [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) commands, it could run on any system it's compiled for(frankly speaking any system [Go](https://golang.org/) could build for), even without ES-E1 installed.
**Note**: tagger doesn't use [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) directly it just generates and prints commands. Technically the systems you're running tagger and exiftool could be two different ones.
**Note**: tagger uses `-geotag` flag to pass it to [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) only as a path. It doesn't parse it or read it. So you need your geotag file to be avaiable on the system you're going to run [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/) on.

#### Reviewing results

To ensure the EXIF data was updated you could use [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/):

```shell
exiftool imagefile.tiff
```

It will print all the EXIF data from the file.

**Note**: If your scans were added to Adobe Lightroom it's required to read metadata via Lightroom interface to see the changes. Otherwise writing metadata to files from Lightroom will overwrite the changes made by [exiftool](https://www.sno.phy.queensu.ca/~phil/exiftool/).

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

## Installation

### Binary release

Just refer to [releases page](https://github.com/teran/eos-1v-tagger/releases), choose the latest one, pick the binary package under `Assets` list and unpack it to appropriate place.

### Compile from source

```shell
go get github.com/teran/eos-1v-tagger/cmd/tagger
```

**Note**: Golang compiler is required to compile and install tagger from source code

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
* **Always** carefuly review exiftool commands *before* applying them.
* It allows you to specify timezone set on EOS 1V to properly timestamp scans so please pay attention to `-timezone` flag **which defaults to UTC timezone**
* Since tagger releases are built on the most recent version of Golang it cannot run on Windows prior to Windows 7. Technically, it should compile fine on Golang v1.10 or earlier but it's unstested and not a part of release procedure. More details could be found at [Golang v1.11 changelog](https://golang.org/doc/go1.11).

## Licence

This software is licenced under GPLv2 terms.
