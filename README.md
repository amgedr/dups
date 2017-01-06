# dups
dups is a program for finding duplicate files in one or more directories. It works by first finding files that have the exact size. Then compares their MD5 hashes to see if their contents are the same too.

### Command-line Options
```bash
-e EXT     Exclude files with extension EXT 
```

### Examples
```bash
# Find duplicates in a directories
dups /tmp

# Find duplicates in multiple directories
dups /tmp ~/tmp

# Ignore jpg files
dups -e jpg /tmp

# Ignore jpg and png files
dups -e jpg -e png /tmp
```

### Installation
If you just want to use the program you can [download the latest release](https://github.com/codehill/dups/releases) from the releases page.

However, if you want the source code too you can use Go as follows:
```
go get github.com/codehill/dups
```

### Contributing
* If you've got any suggestions or questions, [please create an issue here](https://github.com/codehill/dups/issues).
* If you want to contribute a bug or implement a feature, please feel free to do so. Just send me a pull request.

### License
This project is licensed under a MIT license. See the included [LICENSE file](LICENSE) for more details.
