package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func formatJsonData(data []byte) (res []byte, err error) {
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "   ")
	if err != nil {
		return
	}
	res = out.Bytes()
	return
}

func formatStdinData() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := formatJsonData(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", res)
}

func formatFileData(fname string) {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := formatJsonData(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(fname, res, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func main() {

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe {
		formatStdinData()
		return
	}

	if len(os.Args) < 2 {
		return
	}
	formatFileData(os.Args[1])
}
