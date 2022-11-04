# epub2pdf
convert epub to pdf

## Feature

- specify output path -o=/path
- specify input path by -f=/path
- recursive directory by -r
- output details by -v
- output debug details by -m
- allow N jobs at once by -j=N
- delete source file by -d
- optimize code

## Supported Convert
- epub to pdf


## Required
- need to install [`calibre`](https://calibre-ebook.com/download)
- add `calibre` and `ebook-convert` to `PATH` environment variable

linux user can run like this to install calibre

```sh
sudo -v && wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sudo sh /dev/stdin
```


## Quick start
### Option 1: Download binary

[see release](https://github.com/realjf/epub2pdf/releases)


### Option 2: Build from source code
#### Required
- [go](https://go.dev/dl/) development environment
- [make](https://gnuwin32.sourceforge.net/packages/make.htm)

**`Linux`**
```sh
make build
cd bin
./epub2pdf convert /path/to/epub_directory
```
**`Windows`**
```powershell
make.exe build_win
cd bin
epub2pdf.exe convert /path/to/epub_directory
```
**`MacOS`**
```sh
make build_darwin
cd bin
epub2pdf convert /path/to/epub_directory
```

**Output**
The output default is in source directory

## License
epub2pdf is released under the Apache 2.0 license. See [LICENSE](https://github.com/realjf/epub2pdf/blob/master/LICENSE)
