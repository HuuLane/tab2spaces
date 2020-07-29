package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/atotto/clipboard"
)

func filesConcat(filenames []string) string {
	var b strings.Builder
	t := template.Must(template.New("fileHeader").Parse("\n/*\n * Filename: {{.}}\n */\n\n"))

	for _, filename := range filenames {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		// write file header
		err = t.Execute(&b, filename)
		if err != nil {
			panic(err)
		}

		// eliminate Tab
		_, err = b.Write(
			bytes.ReplaceAll(content, []byte("	"), []byte("    ")))
		if err != nil {
			panic(err)
		}
	}
	return b.String()
}

func main() {
	// TODO: ignore dot file
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Which filenames or dir?")
		return
	}
	fmt.Println("will read", args)

	filenames := make([]string, 0, len(args))
	for _, arg := range args {
		fi, err := os.Stat(arg)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			// TODO
			fmt.Println("directory")
		case mode.IsRegular():
			filenames = append(filenames, arg)
		default:
			fmt.Printf("unknown file: %s\n", arg)
		}
	}

	r := filesConcat(filenames)
	err := clipboard.WriteAll(r)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully write to your clipboard!")
	}
}
