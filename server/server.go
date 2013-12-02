package server

import (
  "github.com/mphilpot/gocyberq"
  "encoding/json"
  "os"
  // "fmt"
  "io/ioutil"
)

type Server struct {
	CyberQ gocyberq.CyberQ
	UpdateMillis int
	Port int
}


func LoadConfig(filePath string) (*Server, error) {
	configFile, err := os.Open(filePath);
  	if err != nil { return nil, err }

	b, _ := ioutil.ReadAll(configFile)

	server := &Server{}

	err = json.Unmarshal(b, server)
	if err != nil { return nil, err }

	return server, nil
}

func ConfigHandler() {

}
