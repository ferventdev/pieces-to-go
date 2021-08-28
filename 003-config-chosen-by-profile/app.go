package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

const PROFILES_FOLDER = "profiles"

type config struct {
	Prop1    string
	Prop2    int
	PropSet1 struct {
		Prop3 string
		Prop4 float64
	}
	PropSet2 struct {
		Prop5 []float64
		Prop6 map[string]float32
		Prop7 struct {
			Subprop1 bool
			Subprop2 bool
		}
		Subset struct {
			Prop8 string
			Prop9 time.Time
		}
	}
}

func main() {
	profile := readAppProfile()

	var cfg config
	if _, err := toml.DecodeFile(PROFILES_FOLDER+"/base.toml", &cfg); err != nil {
		log.Fatalln("config file parsing error:", err)
	}

	if profile == "" || profile == "base" {
		log.Println("App profile is not explicitly stated, so only the default (base) profile will be applied")
	} else {
		checkThatProfileExists(profile)
		log.Printf("%s profile is explicitly stated, so it will be applied (it's config will override the default one)\n", profile)
		path := PROFILES_FOLDER + "/" + profile + ".toml"
		if _, err := toml.DecodeFile(path, &cfg); err != nil {
			log.Fatalln("config file parsing error:", err)
		}
	}

	log.Printf("%+v", cfg)
}

func readAppProfile() string {
	var profile string
	flag.StringVar(&profile, "profile", "", "profile to start the app with")
	flag.Parse()
	if profile == "" {
		profile = os.Getenv("PROFILE")
	}
	return strings.ToLower(profile)
}

func checkThatProfileExists(profile string) {
	files, err := ioutil.ReadDir(PROFILES_FOLDER)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files {
		name := f.Name()
		if f.IsDir() || !strings.HasSuffix(name, ".toml") {
			continue
		}
		if profile == strings.TrimSuffix(name, ".toml") {
			return
		}
	}
	log.Fatalf("%s profile does not exist, as its config file not found in ./%s/", profile, PROFILES_FOLDER)
}
