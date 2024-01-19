# msfetch

msfetch is a CLI webscraper for www.musicstore.de
It allows for a product search without leaving the terminal

## Build with Docker
To build the image with docker use the following command.
```
docker build . -t msfetch
```

## Run with Docker
To run the docker image use the following command.
```
docker run -it msfetch <command_name> <flag> <args>
```
## Build and use locally
For local usage you can compile the source code. 

```
go build -o msfetch
```

## Usage
The help prints all the possible commands and flags
```
$ msfetch --help
```

