package gocyberq_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/mphilpot/gocyberq"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("Cyberq", func() {
	var (
		server *httptest.Server
		cyberq *CyberQ
	)

	BeforeEach(func() {
		handler := func(w http.ResponseWriter, r *http.Request) {
			var xmlFile *os.File
			var err error

			if r.URL.Path == "/status.xml" {
				xmlFile, err = os.Open("docs/cyberq_status.xml")
			} else if r.URL.Path == "/all.xml" {
				xmlFile, err = os.Open("docs/cyberq_all.xml")
			} else if r.URL.Path == "/config.xml" {
				xmlFile, err = os.Open("docs/cyberq_config.xml")
			}
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer xmlFile.Close()

			b, _ := ioutil.ReadAll(xmlFile)
			fmt.Fprintf(w, string(b))
		}

		server = httptest.NewServer(http.HandlerFunc(handler))
		cyberq = &CyberQ{URL: server.URL}
	})

	AfterEach(func() {
		server.Close()
	})

	It("should have a test", func() {
		expectedValue := Nutcstatus{
			OutputPercent: "100",
			TimerCurrent:  "00:00:00",
			CookTemp:      "3343",
			Food1Temp:     "823",
			Food2Temp:     "OPEN",
			Food3Temp:     "OPEN",
			CookStatus:    "0",
			Food1Status:   "0",
			Food2Status:   "4",
			Food3Status:   "4",
			TimerStatus:   "0",
			DegUnits:      "1",
			CookCycleTime: "6",
			CookPropBand:  "500",
			CookRamp:      "0",
		}
		Expect(cyberq.Status()).To(Equal(expectedValue))
	})

	It("should return a all object", func() {
		expectedValue := Nutcallstatus{
			Cook: Cook{
				Name:   "Big Green Egg",
				Temp:   "3216",
				Set:    "4000",
				Status: "0",
			},
			Food1: Food1{
				Name:   "Chicken Quarters",
				Temp:   "1482",
				Set:    "1750",
				Status: "0",
			},
			Food2: Food2{
				Name:   "Food2",
				Temp:   "OPEN",
				Set:    "1000",
				Status: "4",
			},
			Food3: Food3{
				Name:   "Food3",
				Temp:   "OPEN",
				Set:    "1000",
				Status: "4",
			},
			OutputPercent: "100",
			TimerCurrent:  "00:00:00",
			TimerStatus:   "0",
			DegUnits:      "1",
			CookCycleTime: "6",
			CookPropBand:  "500",
			CookRamp:      "0",
		}
		Expect(cyberq.All()).To(Equal(expectedValue))
	})

	It("should return a config object", func() {
		// status := All()
		// fmt.Println(status)
		expectedValue := Nutcallstatus{
			Cook: Cook{
				Name:   "Big Green Egg",
				Temp:   "3220",
				Set:    "4000",
				Status: "0",
			},
			Food1: Food1{
				Name:   "Chicken Quarters",
				Temp:   "1493",
				Set:    "1750",
				Status: "0",
			},
			Food2: Food2{
				Name:   "Food2",
				Temp:   "OPEN",
				Set:    "1000",
				Status: "4",
			},
			Food3: Food3{
				Name:   "Food3",
				Temp:   "OPEN",
				Set:    "1000",
				Status: "4",
			},
			OutputPercent: "100",
			TimerCurrent:  "00:00:00",
			TimerStatus:   "0",
			DegUnits:      "",
			CookCycleTime: "",
			CookPropBand:  "",
			CookRamp:      "",
			System: System{
				MenuScrolling: "1",
				LcdBacklight:  "47",
				LcdContrast:   "10",
				DegUnits:      "1",
				AlarmBeeps:    "0",
				KeyBeeps:      "0",
			},
			Control: Control{
				TimeoutAction: "0",
				CookHold:      "2000",
				AlarmDev:      "500",
				CookRamp:      "0",
				OpenDetect:    "1",
				CycleTime:     "6",
				PropBand:      "500",
			},
		}
		actualValue := cyberq.Config()
		Expect(actualValue.Cook).To(Equal(expectedValue.Cook))
		Expect(actualValue.Food1).To(Equal(expectedValue.Food1))
		Expect(actualValue.Food2).To(Equal(expectedValue.Food2))
		Expect(actualValue.Food3).To(Equal(expectedValue.Food3))
		Expect(actualValue.System).To(Equal(expectedValue.System))
		Expect(actualValue.Control).To(Equal(expectedValue.Control))
		Expect(actualValue).To(Equal(expectedValue))
	})
})
