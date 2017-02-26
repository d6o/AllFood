# ![AllFood](http://image.prntscr.com/image/09975b3d3a664237a08c2c7811692415.png)

# AllFood ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/AllFood) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Beta-brightgreen.svg)

The AllFood's goal is to be a perfect tool providing a stupidly easy-to-use and fast program to search all Nearby restaurants based on a latitude and longitude.

## Project Status

AllFood is on beta. Pull Requests [are welcome](https://github.com/DiSiqueira/AllFood#social-coding)

## Features

- Integrated with Facebook Place Search
- Integrated with FourSquare Place Search
- Integrated with Google Place Search
- It's perfect to find your restaurant
- CUSTOMIZE to your needs
- EASY to more providers
- Easy to use
- STUPIDLY [EASY TO USE](https://github.com/DiSiqueira/AllFood#usage)
- Very fast start up and response time
- Uses native libs

## Installation

### Option 1: Go Get

```bash
$ go get github.com/DiSiqueira/AllFood
$ AllFood -h
```

### Option 2: From source

```bash
$ go get gopkg.in/ini.v1
$ git clone https://github.com/DiSiqueira/AllFood.git
$ cd AllFood/
$ go build *.go
```

## Usage

### Basic usage Google Places

```bash
# Find restaurants using Only Google Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -google-key=AIzaSyAOEARYwKidXTiNkM982OP21A8LOAAAZ
```

### Basic usage Facebook Places

```bash
# Find restaurants using Only Facebook Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -fb-app-id=1054064656577111 -fb-app-secret=5abbabac4f45764d534704f24e4a5aaa
```

### Basic usage FourSquare Places

```bash
# Find restaurants using Only FourSquare Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -fs-app-id=KBAGZ3RDIEGOI42A241ADIEGOSDWULOLIQCCHK20NFSCU -fs-app-secret=2G4VSOHLOLNHBW3CBCSKWIGSLA5XM5NG11WHLOLR52FW4GCYAAA
```

### Combine Google Places And Facebook Places Results

```bash
# Find restaurants using Google Places and Facebook Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -google-key=AIzaSyAOEARYwKidXTiNkM982OP21A8LOLAAA -fb-app-id=1054064656577111 -fb-app-secret=5abbabac4f45764d534704f24e4a5aaa
```

### Combine Google Places And FourSquare Places Results

```bash
# Find restaurants using Google Places and FourSquare Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -google-key=AIzaSyAOEARYwKidXTiNkM982OP21A8LOLAAA -fs-app-id=KBAGZ3RDIEGOI42A241ADIEGOSDWULOLIQCCHK20NFSCU -fs-app-secret=2G4VSOHLOLNHBW3CBCSKWIGSLA5XM5NG11WHLOLR52FW4GCYAAA
```

### Combine Facebook Places And FourSquare Places Results

```bash
# Find restaurants using Facebook Places and FourSquare Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -fb-app-id=1054064656577111 -fb-app-secret=5abbabac4f45764d534704f24e145aaa -fs-app-id=KBAGZ3RDIEGOI42A241ADIEGOSDWULOLIQCCHK20NFSCU -fs-app-secret=2G4VSOHLOLNHBW3CBCSKWIGSLA5XM5NG11WHLOLR52FW4GCYAAA
```

### Combine All Results

```bash
# Find restaurants using Google Places, Facebook Places and FourSquare Places
$ AllFood -lat=-22.314459 -lng=-49.058695 -radius=1 -google-key=AIzaSyAOEARYwKidXTiNkM982OP21A8LOLAAA -fb-app-id=1054064656577111 -fb-app-secret=5abbabac4f45764d534704f24e145aaa -fs-app-id=KBAGZ3RDIEGOI42A241ADIEGOSDWULOLIQCCHK20NFSCU -fs-app-secret=2G4VSOHLOLNHBW3CBCSKWIGSLA5XM5NG11WHLOLR52FW4GCYAAA
```

### Show help

```bash
$ AllFood -h
```

## Program Output

![](http://image.prntscr.com/image/bfd678a833a949f3b3caa8331b376f37.png)

## Program Help

![](http://image.prntscr.com/image/7ad9e94e11d049a992b202cbc764eadf.png)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/AllFood/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ go get gopkg.in/ini.v1
$ git clone --recursive git@github.com:DiSiqueira/AllFood.git
$ cd AllFood/
$ go run *.go
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/AllFood/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.