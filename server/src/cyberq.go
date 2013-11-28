package gocyberq

import (
        // "encoding/json"
        "encoding/xml"
        // "net/http"
        "fmt"
        "io/ioutil"
        "os"
)

type Nutcstatus struct {
     OutputPercent string `xml:"OUTPUT_PERCENT"`
     TimerCurrent string `xml:"TIMER_CURR"`
     CookTemp string `xml:"COOK_TEMP"`
     Food1Temp string `xml:"FOOD1_TEMP"`
     Food2Temp string `xml:"FOOD2_TEMP"`
     Food3Temp string `xml:"FOOD3_TEMP"`
     CookStatus string `xml:"COOK_STATUS"`
     Food1Status string `xml:"FOOD1_STATUS"`
     Food2Status string `xml:"FOOD2_STATUS"`
     Food3Status string `xml:"FOOD3_STATUS"`
     TimerStatus string `xml:"TIMER_STATUS"`
     DegUnits string `xml:"DEG_UNITS"`
     CookCycleTime string `xml:"COOK_CYCTIME"`
     CookPropBand string `xml:"COOK_PROPBAND"`
     CookRamp string `xml:"COOK_RAMP"`
}

type Nutcallstatus struct {
    Cook Cook `xml:"COOK"`
    Food1 Food1 `xml:"FOOD1"`
    Food2 Food2 `xml:"FOOD2"`
    Food3 Food3 `xml:"FOOD3"`
    OutputPercent string `xml:"OUTPUT_PERCENT"`
    TimerCurrent string `xml:"TIMER_CURR"`
    TimerStatus string `xml:"TIMER_STATUS"`
    DegUnits string `xml:"DEG_UNITS"`
    CookCycleTime string `xml:"COOK_CYCTIME"`
    CookPropBand string `xml:"COOK_PROPBAND"`
    CookRamp string `xml:"COOK_RAMP"`

    // Used only in the config call
    System System `xml:"SYSTEM"`
    Control Control `xml:"CONTROL"`
}

type Cook struct {
    Name string `xml:"COOK_NAME"`
    Temp string `xml:"COOK_TEMP"`
    Set string `xml:"COOK_SET"`
    Status string `xml:"COOK_STATUS"`
}

type Food1 struct {
    Name string `xml:"FOOD1_NAME"`
    Temp string `xml:"FOOD1_TEMP"`
    Set string `xml:"FOOD1_SET"`
    Status string `xml:"FOOD1_STATUS"`
}

type Food2 struct {
    Name string `xml:"FOOD2_NAME"`
    Temp string `xml:"FOOD2_TEMP"`
    Set string `xml:"FOOD2_SET"`
    Status string `xml:"FOOD2_STATUS"`
}

type Food3 struct {
    Name string `xml:"FOOD3_NAME"`
    Temp string `xml:"FOOD3_TEMP"`
    Set string `xml:"FOOD3_SET"`
    Status string `xml:"FOOD3_STATUS"`
}

type System struct {
    MenuScrolling string `xml:"MENU_SCROLLING"`
    LcdBacklight string `xml:"LCD_BACKLIGHT"`
    LcdContrast string `xml:"LCD_CONTRAST"`
    DegUnits string `xml:"DEG_UNITS"`
    AlarmBeeps string `xml:"ALARM_BEEPS"`
    KeyBeeps string `xml:"KEY_BEEPS"`
}

type Control struct {
    TimeoutAction string `xml:"TIMEOUT_ACTION"`
    CookHold string `xml:"COOKHOLD"`
    AlarmDev string `xml:"ALARMDEV"`
    CookRamp string `xml:"COOK_RAMP"`
    OpenDetect string `xml:"OPENDETECT"`
    CycleTime string `xml:"CYCTIME"`
    PropBand string `xml:"PROPBAND"`
}

func Status() (status Nutcstatus) {
    xmlFile, err := os.Open("../docs/cyberq_status.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)

    xml.Unmarshal(b, &status)

    // fmt.Println(status)
    return status
}

func All() (status Nutcallstatus) {
    xmlFile, err := os.Open("../docs/cyberq_all.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)

    xml.Unmarshal(b, &status)

    // fmt.Println(status)
    return status
}

func Config() (status Nutcallstatus) {
    xmlFile, err := os.Open("../docs/cyberq_config.xml")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer xmlFile.Close()

    b, _ := ioutil.ReadAll(xmlFile)

    xml.Unmarshal(b, &status)

    // fmt.Println(status)
    return status
}
