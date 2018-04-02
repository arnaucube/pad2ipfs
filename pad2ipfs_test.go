package pad2ipfs

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
)

const checkIcon = "\xE2\x9C\x94 "

func TestGenerateEmptyMT(t *testing.T) {
	hash1, err := Add("http://board.net/p/pad1", "md")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		color.Green(checkIcon + "added file1")
		fmt.Print("	IPFS hash:")
		color.Blue(hash1)
	}
	hash2, err := Add("http://board.net/p/pad2", "txt")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		color.Green(checkIcon + "added file2")
		fmt.Print("	IPFS hash:")
		color.Blue(hash2)
	}
	hash3, err := Add("http://board.net/p/pad3", "html")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		color.Green(checkIcon + "added file3")
		fmt.Print("	IPFS hash:")
		color.Blue(hash3)
	}

	err = Get(hash1, "pad1.md")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		color.Green(checkIcon + "getted file1")
	}
	err = Get(hash2, "pad2.txt")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		color.Green(checkIcon + "getted file2")
	}
	err = Get(hash3, "pad3.html")
	if err != nil {
		t.Errorf(err.Error())
	} else {
		color.Green(checkIcon + "getted file3")
	}
}
