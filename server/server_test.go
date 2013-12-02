package server_test

import (
	. "github.com/mphilpot/gocyberq/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/mphilpot/gocyberq"
)

var _ = Describe("GoCyberQ Server", func() {
	var file *os.File

	BeforeEach(func() {
		var err error
		file, err = ioutil.TempFile("", "")
		if err != nil {
			Fail("Couldn't create temp file")
		}
	})

	AfterEach(func() {
		os.Remove(file.Name())
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

		ioutil.WriteFile(file.Name(), b, 0777)

		server, err := LoadConfig(file.Name())

		Expect(server).To(Equal(expectedConfig))

		Expect(true).To(Equal(true))
	})
})
