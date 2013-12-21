package server_test

import (
	. "github.com/mphilpot/gocyberq/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/mphilpot/gocyberq"
	"path/filepath"
)

var _ = Describe("GoCyberQ Server", func() {
	var dirName string

	BeforeEach(func() {
		var err error
		dirName, err = ioutil.TempDir("", "")
		if err != nil {
			Fail("Couldn't create temp dirName")
		}
	})

	AfterEach(func() {
		os.Remove(dirName)
	})

	It("Load configuration", func () {
		expectedConfig := &Server {
			CyberQ: gocyberq.CyberQ {
				URL: "10.0.0.2",
			},
			UpdateMillis: 10000,
			Port: 8181,
		}
		b, err := json.Marshal(expectedConfig);

		if err != nil { Fail("Couldn't marshal json") }

		ioutil.WriteFile(filepath.Join(dirName, ConfigName), b, 0777)

		server, err := LoadConfig(dirName)

		Expect(server).To(Equal(expectedConfig))

		Expect(true).To(Equal(true))
	})
})
