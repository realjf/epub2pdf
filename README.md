# epub2pdf
convert epub to pdf

## Feature

### v0.0.3(todo)
- -o=/path specify output path
- -f=/path specify input path
- -r recursive directory
- -v output details
- -m output debug details
- -j=N allow N jobs at once
- -d delete source file
- -t specifies the timeout for each conversion.


### v0.0.2
- -o=/path specify output path
- -f=/path specify input path
- -r recursive directory
- -v output details
- -m output debug details
- -j=N allow N jobs at once
- -d delete source file
- -t specifies the timeout for each conversion.


### v0.0.1(discard)


## Supported Convert
- epub to pdf


## Required
- need to install [`calibre`](https://calibre-ebook.com/download)
- add `calibre` and `ebook-convert` to `PATH` environment variable

linux user can run like this to install calibre

```sh
sudo -v && wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | sudo sh /dev/stdin
```
windows and macOS user can download from [Release Assets](https://github.com/realjf/epub2pdf/releases)

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
