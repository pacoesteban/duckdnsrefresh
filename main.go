//
// main.go
// Copyright (C) 2017 Paco Esteban <paco@onna.be>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var (
	token   string
	domains []string
)

func main() {
	var v = flag.Bool("v", false, "Enable verbose output")
	flag.Parse()
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath("/etc/duckdns/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.duckdns") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	token = viper.GetString("token")
	domains = viper.GetStringSlice("domains")

	for _, d := range domains {
		if *v {
			log.Printf("Trying to update domain: %s", d)
		}
		req := fmt.Sprintf("https://www.duckdns.org/update?domains=%s&token=%s", d, token)
		r, err := http.Get(req)
		if err != nil {
			log.Fatalf("cannot get url: %s : %q", req, err)
		}
		if r.StatusCode != http.StatusOK {
			log.Fatalf("incorrect response from server: %s ", r.Status)
		}
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("could not read the body: %q", err)
		}
		if string(body) != "OK" {
			log.Fatalf("error while refreshing domain: %s", string(body))
		}

		if *v {
			log.Printf("All good. got response '%s' from DuckDNS.", string(body))
		}
	}
}
