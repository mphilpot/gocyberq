package server

import (
  "github.com/mphilpot/gocyberq"
  "encoding/json"
  "os"
  // "fmt"
  "io/ioutil"
  "path/filepath"
)

type Server struct {
	CyberQ gocyberq.CyberQ
	UpdateMillis int
	Port int
	HomeDir string
}

const DefaultHome string = "~/.gocyberq"
const ConfigName string = "cyberq.json"

func LoadConfig(homeDir string) (*Server, error) {
	configFile, err := os.Open(filepath.Join(homeDir, ConfigName));
  	if err != nil { return nil, err }

	b, _ := ioutil.ReadAll(configFile)

	server := &Server{}

	err = json.Unmarshal(b, server)
	if err != nil { return nil, err }

	return server, nil
}

func (server *Server) SaveConfig() error {
	b, err := json.Marshal(server);

	if err != nil { return err }

	// TODO(philpott): Change permissions
	ioutil.WriteFile(filepath.Join(server.HomeDir, ConfigName), b, 0777)

	return nil
}

func ConfigHandler() {

}
