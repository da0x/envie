# envie
![Go](https://github.com/da0x/envie/workflows/Go/badge.svg?branch=main)


Welcome to envie! This project helps you read a full struct of env variables
from your system. First, the `Auto` function will attempt to read the environment
variables from a properties file. Default path is `.env`, see `AutoPath`.
Afterwards, it will attempt to read the same values from the system's
environment.
If both methods fail, it will panic if AutoPanic is set to true. (Default).
Otherwise the values from the system's env will override anything loaded
from the properties file. If either method fully succeeds, there will be no
errors.
## Installation
To install the library simply run:
```
$ go get github.com/da0x/envie
```
To import the library, use:
```
import "github.com/da0x/envie"
```
## Usage
### System Environment
```
$ export VARIABLE_ONE = hello
$ export VARIABLE_TWO = world
```
### Properties File
```
// filename: .env
// this is a comment
VARIABLE_ONE = hello
VARIABLE_TWO = world
```
### Code
```
type entity struct {
	V1 string `envie:"SOME_VAR1"`
	V2 string `envie:"SOME_VAR2"`
}

func main() {
	var e entity
    envie.Auto(&e)
    println(e)
}
```
## Maintainer
Daher Alfawares
