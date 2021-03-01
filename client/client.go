package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "localhost:8080"
const endpoint = "/endpoint"

func main() {
	fmt.Printf("pid: %d\n", os.Getpid())

	t := flag.String("type", "readandclose", `specify the type of request to send:
close				- send and close
read				- send and read
readandclose		- send, read and close
nothing 			- send
readbyte 			- send and read one byte
readbyteandclose	- send, read  one byte and close
`)
	n := flag.Int("number", 5, `specify the number of requests to send`)
	flag.Parse()

	var fn func() error

	switch *t {
	case "close":
		fn = makeRequestAndCloseBody
	case "read":
		fn = makeRequestAndReadBody
	case "readandclose":
		fn = makeRequestAndReadAndCloseBody
	case "nothing":
		fn = makeRequest
	case "readonebyte":
		fn = makeRequestAndReadByte
	case "readonebyteandclose":
		fn = makeRequestAndReadByteAndCloseBody
	default:
		log.Fatalln("unknown request type")
	}

	for i := 0; i < *n; i++ {
		if err := fn(); err != nil {
			log.Fatalln(err)
		}
	}
}

func makeRequestAndReadAndCloseBody() error {
	res, err := http.Get("http://" + url + endpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = io.Copy(ioutil.Discard, res.Body)
	return err
}

func makeRequestAndReadBody() error {
	res, err := http.Get("http://" + url + endpoint)
	if err != nil {
		return err
	}

	_, err = io.Copy(ioutil.Discard, res.Body)

	return err
}

func makeRequestAndCloseBody() error {
	res, err := http.Get("http://" + url + endpoint)
	if err != nil {
		return err
	}

	return res.Body.Close()
}

func makeRequest() error {
	_, err := http.Get("http://" + url + endpoint)
	return err
}

func makeRequestAndReadByte() error {
	res, err := http.Get("http://" + url + endpoint)
	if err != nil {
		return err
	}

	var oneByteBuff [1]byte
	_, err = io.ReadFull(res.Body, oneByteBuff[:])

	return err
}

func makeRequestAndReadByteAndCloseBody() error {
	res, err := http.Get("http://" + url + endpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var oneByteBuff [1]byte
	_, err = io.ReadFull(res.Body, oneByteBuff[:])

	return err
}
