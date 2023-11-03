# ccwc

A simplified version of one of the most famous, valuabe Unix command line tool `wc` written in Go/Golang with executables for windows, osx and linux.

`ccwc` just like the original `wc` is a command line tool that allows you to count the number of lines, words, bytes, and characters in the file specified by the `File` parameter.

## Install

Find the corresponding executable according to your operating system platform here:

https://github.com/emanuelquerty/ccwc/releases/tag/v0.0.1

## Usage

`ccwc [OPTION] [FILE]`

If `OPTION` is not specified, it prints the line, word, and byte count in this same order of the file specified by the `FILE` parameter.

### Example usage
`ccwc test.txt` 

`ccwc` supports the following `Flags` as of this time:

|    Flag     |        Description         |      
| ----------- |:--------------------------:|
|    -c       | print the byte count       |
|    -m       | print the character count  |
|    -w       | print the word count       |
|    -l       | print the newline count    |

### Example usage

`ccwc -c test.txt`

The above command prints the number of bytes in `test.txt`

## License

Copyright (c) 2023 [Emanuel Inacio](https://github.com/emanuelquerty)

Licensed under [MIT License](./LICENSE)
