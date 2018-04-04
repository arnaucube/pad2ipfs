# pad2ipfs [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/pad2ipfs)](https://goreportcard.com/report/github.com/arnaucode/pad2ipfs)

Simply Go lang library to get the content from a pad (etherpad) and put into IPFS.


Needs to have installed IPFS (https://ipfs.io), and the daemon running ('> ipfs daemon').

## Install
```
go get github.com/arnaucode/pad2ipfs
```

## Usage

The added pads are stored in 'addedPads' directory.
The getted pads are stored in 'gettedPads' directory.


#### - Add
Adds the content from a pad to IPFS

```go
hash, err := pad2ipfs.Add(link, format)
```
```go
hash, err := pad2ipfs.Add("https://board.net/p/selectedpad", "md")
if err!=nil{
  fmt.Println(err)
}
```
Supported formats:
  - md
  - txt
  - html
  - pdf
  - odt



#### - Get
Gets the content from IPFS and stores it into a file

```go
err := pad2ipfs.Get(hash, filename)
```
```go
err := pad2ipfs.Get("QmVyp4JSREK5syLmNRCafkZkhzC7CfvS9qYWKfvfffqK2B", "selectedpad.md")
if err!=nil {
  fmt.Println(err)
}
```



### CLI
In the directory /pad2ipfs-cli is placed the cli to interact directly with the library from the command line. Here is a screenshot:

![pad2ipfs-cli-screenshot](https://raw.githubusercontent.com/arnaucode/pad2ipfs/master/pad2ipfs-cli-screenshot.png "pad2ipfs-cli-screenshot")
