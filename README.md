EOS 1V Tagger
=============

Intro
-----

Here's Golang library and CLI binary to parse CSV's from Canon ES-E1 software to allow EXIF tagging film scans according to metadata from such CSV's.

Development status: Just started

Usage
-----

```shell
tagger CSV_FILE
```

At the momment it suppose you have the following format film scan files: `FILM_%5f.dng`, for example: `FILM_00001.dng` so exiftool commands are generated based on that file naming scheme.

NOTICES
-------

* It's in **deep alpha** state
* It's **NOT** going to make any changes to real data: just prints exiftool commands to STDOUT
* It relies on the data provided by ES-E1 software in CSV format(in EOS 1V Memory just export via `File` -> `Export` -> `CSV`)
* It does **NOT** perform timezone detection since ES-E1 exports timestamps without timezone mark.
* **Always** carefully review exiftool commands *before* applying them.
* It sets ISO for each frame from the film settings. So if you have set ISO for particular frame to the another value it will still use the one from the film properties.

Licence
-------

This software is licenced under GPLv2 terms.
