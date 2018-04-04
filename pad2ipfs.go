package pad2ipfs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	sh "github.com/ipfs/go-ipfs-api"
)

//AddedPads is the directory where are stored the pads that are added to IPFS
const AddedPads = "addedPads"

//GettedPads is the directory where are stored the pads that are getted from IPFS
const GettedPads = "gettedPads"

//Add gets the content from the etherpad specified in the link, and downloads it in the format of the specified extension, and then, puts it into IPFS
func Add(link string, extension string) (string, error) {
	if extension != "md" && extension != "txt" && extension != "html" && extension != "pdf" && extension != "odt" {
		return "", errors.New("No valid extension")
	}
	format := extension
	if extension == "md" {
		format = "markdown"
		extension = "md"
	}

	//create the pads directory
	_ = os.Mkdir(AddedPads, os.ModePerm)

	//get the pad name
	linkSplitted := strings.Split(link, "/")
	padName := linkSplitted[len(linkSplitted)-1]
	completeLink := link + "/export/" + format

	//get the content from the url
	r, err := http.Get(completeLink)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}

	//save the content into a file
	err = ioutil.WriteFile(AddedPads+"/"+padName+"."+extension, content, 0644)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//connect to ipfs shell
	s := sh.NewShell("localhost:5001")
	//save the file into IPFS
	ipfsHash, err := s.AddDir(AddedPads + "/" + padName + "." + extension)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return ipfsHash, nil
}

//Get gets the content from IPFS for a given hash, and saves it into a file
func Get(hash string, filename string) error {
	//create the pads directory
	_ = os.Mkdir(GettedPads, os.ModePerm)

	//connect to ipfs shell
	s := sh.NewShell("localhost:5001")
	err := s.Get(hash, GettedPads+"/"+filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
