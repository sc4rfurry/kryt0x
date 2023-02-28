<h1 align="center">
  Kryt0x
</h1>

<h4 align="center">Utility to Mutate/Encrypt a file using AES-256 Algorithm.</h4>
<div style="text-align:center">
    <div style="align:center">
    <img src="https://img.shields.io/badge/Author-sc4rfurry-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Author">
    <img src="https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square&logo=github&logoColor=white&color=5194f0&bgcolor=110d17" alt="Version">
    <img src="https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17" alt="Go Version">
    <img src="https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17" alt="OS">
    </div>
</div>

#

## Table of Contents

- [Installation](#installation)
- [Features](#features)
- [Running Kryt0x](#running-Kryt0x)
    - [Options](#options)
- [Building Kryt0x](#building-Kryt0x)
- [References](#references)
- [Contributing](#contributing)
- [License](#license)


#

### 🔧 Technologies & Tools

![](https://img.shields.io/badge/Editor-VS_Code-informational?style=flat-square&logo=visual-studio&logoColor=blue&color=5194f0)
![](https://img.shields.io/badge/Language-Go-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)
![](https://img.shields.io/badge/Go_Version-1.18.1-informational?style=flat-square&logo=Go&logoColor=cyan&color=5194f0&bgcolor=110d17)

#

### 📚 Requirements
> - Go 18.1 linux/amd64

#
### Installation

- sudo apt-get update && sudo apt-get golang
- git clone https://github.com/sc4rfurry/kryt0x.git
- cd krty0x
- go get .
- go build main.go
    - or use the `builder.sh` script to build the tool.


### Features

- Encrypts a file using AES-256 Algorithm.
- Static Binary (No Dependencies almost :)
#

## Running Kryt0x
```sh
go run main.go --help
```

### Example

To run the tool on a target, just use the following command.

```console
go run main.go -e /path/to/file
```
OR
```console
go run main.go -d /path/to/file <key>
```
#

## Options
```sh
Usage: ./main -e <file> OR ./main -d <file> <key>
________________________________________________________________________________________________

	Options: 
		File 		File to encrypt/decrypt
		-e 		    Excrypt file
		-d 		    Decrypt file
		-h/--help 	Show this help menu


	!Note:- Decryption need Filename and Encryption Key
```

#

## Building Kryt0x
> To build the tool, you can use the following command.
```sh
env GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w -extldflags "-static"' -o Kryt0x main.go
```

> You can also use the bultin Bash script to build the tool.

- Before running the script, make sure to give it execution permissions.
- The bash script can build both Linux and Windows binaries.
- Binaries will be Stripped and Compressed. (lolcat, strip and upx are required)
```sh
chmod +x builder.sh
./builder.sh main.go
```
#

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

[def]: https://img.shields.io/badge/OS-Linux-informational?style=flat-square&logo=ubuntu&logoColor=green&color=5194f0&bgcolor=110d17