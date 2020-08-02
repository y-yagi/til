package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const libtpl = `package main

func hello(name string) string {
	return "Hello, " + name
}
`

const libtesttpl = `package main
import (
	"reflect"
	"testing"
)

var tests = []struct {
	in1   string
	want string
}{
	{"Bob", "Hello, Bob"},
}

func TestSuccess(t *testing.T) {
	for _, tt := range tests {
	  got := hello(tt.in1)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in1 %v, got %v, want %v", tt.in1, got, tt.want)
	  }
	}
}
`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "specify name")
		return
	}

	dirname := os.Args[1]
	dirname = strings.Replace(dirname, " ", "-", -1)
	dirname = strings.ToLower(dirname)

	if err := os.Mkdir(dirname, 0755); err != nil {
		log.Fatal(err)
	}
	file := filepath.Join(dirname, "main.go")
	if err := ioutil.WriteFile(file, []byte(libtpl), 0644); err != nil {
		log.Fatal(err)
	}
	testfile := filepath.Join(dirname, "main_test.go")
	if err := ioutil.WriteFile(testfile, []byte(libtesttpl), 0644); err != nil {
		log.Fatal(err)
	}

	return
}
