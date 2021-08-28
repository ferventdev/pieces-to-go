package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var profile string
	flag.StringVar(&profile, "profile", "", "profile to start the app with")
	flag.Parse()
	log.Println("profile (from CLI argument, i.e. flag) =", profile)
	if profile == "" {
		profile = os.Getenv("PROFILE")
		log.Println("profile (from env. variable) =", profile)
	}
	if profile == "" {
		log.Fatalln("App profile must be defined either as an env. variable (PROFILE) or a CLI flag (-profile)")
	}
}
