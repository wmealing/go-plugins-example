package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"plugin"
	"strings"

	"github.com/sourcegraph/go-diff/diff"
)

type Greeter interface {
	Greet(diff.FileDiff)
}

func patch_parsed_hook(d diff.FileDiff) {

	/* There will be an observer hook here here to loop through
	   and msg all observers */

	log.Println("👌 - Running Patch parsed hooks.")
	log.Println("🛈 D = ", d)

	/* This is a long dance, that reads the compiled so files
	   in plugins */

	files, err := ioutil.ReadDir("./plugins/")

	if err != nil {
		log.Fatal(err)
	}

	/* For each plugin file */
	for _, f := range files {
		var name = f.Name()

		if !strings.HasSuffix(name, ".so") {
			continue
		}

		log.Println("Found plugin", name)

		plug, err := plugin.Open("./plugins/" + name)
		symGreeter, err := plug.Lookup("Greeter")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var greeter Greeter
		greeter, ok := symGreeter.(Greeter)

		if !ok {
			fmt.Println("unexpected type from module symbol")
			os.Exit(1)
		}

		greeter.Greet(d)
	}

}

func parse_patch(p string) {
	log.Println("🛈 - Parsing patch file:", p)

	diffData, err := ioutil.ReadFile("test.diff")

	if err != nil {
		log.Println("ERROR")
	}

	diff, err := diff.ParseFileDiff(diffData)

	if err != nil {
		log.Printf("🛈 - %s: parseHunks err %v, want %v", p, err, nil)
	}

	patch_parsed_hook(*diff)
}

func main() {
	log.Println("🛈 - Starting")
	parse_patch("test.diff")
	log.Println("🛈 - Ending")
}
