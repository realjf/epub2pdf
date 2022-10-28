# epub2pdf
convert epub to pdf

## Features

- support linux/windows/macOS

## Required
- need to install [`calibre`](https://calibre-ebook.com/download)
- add `calibre` to PATH Environment Variable


## Quick start
### Option 1: Download binary

see release


### Option 2: Build from source code
#### Required
- [go](https://go.dev/dl/) development environment
- [make](https://gnuwin32.sourceforge.net/packages/make.htm)

**`Linux`**
```sh
make build
cd bin
./epub2pdf /path/to/epub_directory
```
**`Windows`**
```powershell
make.exe build_win
cd bin
epub2pdf.exe /path/to/epub_directory
```
**`MacOS`**
```sh
make build_darwin
cd bin
epub2pdf.app /path/to/epub_directory
```

