package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jszwec/csvutil"
)

type station struct {
	Name       string
	NameKana   string
	NameRomaji string
	Region     string
}

func (s *station) display() string {
	return fmt.Sprintf("%s, %s (%s)", s.Region, s.Name, s.NameKana)
}

type betweenStation struct {
	Start        string
	End          string
	Stopover     string
	Distance     float64
	RequiredTime int64
}

func (s *betweenStation) display() string {
	return fmt.Sprintf("%s - %s (%s)", s.Start, s.End, s.Stopover)
}

func main() {
	var stations []station
	r, err := ioutil.ReadFile("stations.csv")
	if err != nil {
		log.Fatal(err)
	}
	if err := csvutil.Unmarshal(r, &stations); err != nil {
		log.Fatal(err)
	}

	for _, s := range stations {
		fmt.Printf("%+v\n", s.display())
	}

	var betweens []betweenStation
	r, err = ioutil.ReadFile("betweens.csv")
	if err != nil {
		log.Fatal(err)
	}
	if err := csvutil.Unmarshal(r, &betweens); err != nil {
		log.Fatal(err)
	}

	for _, b := range betweens {
		fmt.Printf("%+v\n", b.display())
	}
}
