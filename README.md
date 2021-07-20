# ncmbox

Netease Cloud Music Box

## Build

### Install Golang

Please refer to the documentation: [Download and install](https://golang.org/doc/install).

### Get Source Code

```bash
$ git clone https://github.com/zexi/ncmbox
$ cd ncmbox
```

### Build Binary

```bash
$ make build
$ ls ./_output/bin
ncmbox
```

## Run

```bash
# create cellphone login config
$ mkdir ~/.ncmbox
$ echo "username: YOUR_PHONE_NUMBER\npassword: YOUR_PASSWORD" > ~/.ncmbox/config.yml

# start app
$ ./_output/bin/ncmbox
```
