package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Component struct {
	ComponentID    ComponentID     `json:"componentID,omitempty"`
	CoreProperties Microcontroller `json:"coreProperties,omitempty"`
	Pins           []pinSpec       `json:"pins,omitempty"`
	Package        Package         `json:"package,omitempty"`
}

type Package struct {
	Length               *ValueOptions `json:"length,omitempty"`
	Width                *ValueOptions `json:"width,omitempty"`
	Height               *ValueOptions `json:"height,omitempty"`
	NominalFootprints    *ExternalFile `json:"nominalFootprints,omitempty"`
	BreakoutExamples     *ExternalFile `json:"breakoutExamples,omitempty"`
	PartModelInformation *ExternalFile `json:"partModelInformation,omitempty"`
	StandardPackageSize  string        `json:"standardPackageSize,omitempty"`
	StandardPackageType  string        `json:"standardPackageType,omitempty"`
}

type ExternalFile struct {
	FileDescription    string `json:"fileDescription,omitempty"`
	FileType           string `json:"fileType,omitempty"`
	FileExtension      string `json:"fileExtension,omitempty"`
	CompanionSoftware  string `json:"companionSoftware,omitempty"`
	StandardReferenced string `json:"standardReferenced,omitempty"`
	FileURI            string `json:"fileURI,omitempty"`
}

type ComponentID struct {
	PartType           string              `json:"partType,omitempty"`
	Manufacturer       string              `json:"manufacturer,omitempty"`
	ComponentName      string              `json:"componentName,omitempty"`
	OrderableMPN       []string            `json:"orderableMPN,omitempty"`
	SourceDatasheetID  *SourceDatasheetID  `json:"sourceDatasheetID,omitempty"`
	DigitalDatasheetID *DigitalDatasheetID `json:"digitalDatasheetID,omitempty"`
	Status             string              `json:"status,omitempty"`
	ComplianceList     []string            `json:"complianceList,omitempty"`
}

type SourceDatasheetID struct {
	PublishedDate string `json:"publishedDate,omitempty"`
	Version       string `json:"version,omitempty"`
	DatasheetURI  string `json:"datasheetURI,omitempty"`
	ProductURI    string `json:"productURI,omitempty"`
}

type DigitalDatasheetID struct {
	PublishedDate          string `json:"publishedDate,omitempty"`
	EDatasheetSpecRevision string `json:"eDatasheetSpecRevision,omitempty"`
	Guid                   string `json:"guid,omitempty"`
}

type Microcontroller struct {
	PartType             string               `json:"partType,omitempty"`
	OnChipFlash          *ValueOptions        `json:"onChipFlash,omitempty"`
	OnChipRAM            *ValueOptions        `json:"onChipRAM,omitempty"`
	OnChipROM            *ValueOptions        `json:"onChipROM,omitempty"`
	CoreProcessor        string               `json:"coreProcessor,omitempty"`
	CoreArchitectureBits string               `json:"coreArchitectureBits,omitempty"`
	ClockSpeed           *ValueOptions        `json:"clockSpeed,omitempty"`
	FirmwareVersion      string               `json:"firmwareVersion,omitempty"`
	CurrentConsumption   []CurrentConsumption `json:"currentConsumption,omitempty"`
}

type CurrentConsumption struct {
	SupplyName       string        `json:"supplyName,omitempty"`
	QuiescentCurrent *ValueOptions `json:"quiescentCurrent,omitempty"`
	ShutdownCurrent  *ValueOptions `json:"shutdownCurrent,omitempty"`
	ActiveCurrent    *ValueOptions `json:"activeCurrent,omitempty"`
	SleepCurrent     *ValueOptions `json:"sleepCurrent,omitempty"`
	IdleCurrent      *ValueOptions `json:"idleCurrent,omitempty"`
}

type ValueOptions struct {
	Values []Value `json:"values,omitempty"`
}

type Value struct {
	SiUnit                 string   `json:"siUnit,omitempty"`
	UnitName               string   `json:"unitName,omitempty"`
	UnitFactor             float64  `json:"unitFactor,omitempty"`
	TypValue               float64  `json:"typValue,omitempty"`
	MinValue               float64  `json:"minValue,omitempty"`
	MaxValue               float64  `json:"maxValue,omitempty"`
	RelativeValueReference string   `json:"relativeValueReference,omitempty"`
	RelativeValueModifier  float64  `json:"relativeValueModifier,omitempty"`
	RelativeValueOperator  string   `json:"relativeValueOperator,omitempty"`
	Conditions             []string `json:"conditions,omitempty"`
}

type FunctionProperties struct {
	PerFunctionName         string `json:"perFunctionName,omitempty"`
	InterfaceType           string `json:"interfaceType,omitempty"`
	PinUsage                string `json:"pinUsage,omitempty"`
	Direction               string `json:"direction,omitempty"`
	ElectricalConfiguration string `json:"electricalConfiguration,omitempty"`
}

type pinSpec struct {
	TerminalIdentifier         string               `json:"terminalIdentifier,omitempty"`
	Name                       string               `json:"name,omitempty"`
	NumberOfSupportedFunctions int                  `json:"numberOfSupportedFunctions,omitempty"`
	FunctionProperties         []FunctionProperties `json:"functionProperties,omitempty"`
	Vih                        *ValueOptions        `json:"vih,omitempty"`
	Vil                        *ValueOptions        `json:"vil,omitempty"`
	Vol                        *ValueOptions        `json:"vol,omitempty"`
	Voh                        *ValueOptions        `json:"voh,omitempty"`
	AbsVmax                    *ValueOptions        `json:"absVmax,omitempty"`
	AbsVmin                    *ValueOptions        `json:"absVmin,omitempty"`
	Vmax                       *ValueOptions        `json:"vmax,omitempty"`
	Imax                       *ValueOptions        `json:"imax,omitempty"`
	InputLeakage               *ValueOptions        `json:"inputLeakage,omitempty"`
	VoltageOptions             *ValueOptions        `json:"voltageOptions,omitempty"`
	Esd                        bool                 `json:"esd,omitempty"`
	InternalPullUp             *ValueOptions        `json:"internalPullUp,omitempty"`
}

func addVOLProps(props pinSpec) pinSpec {
	props.Vol = &ValueOptions{
		Values: []Value{
			{
				SiUnit:     "volt",
				MaxValue:   .4,
				Conditions: []string{"CMOS port", "IIO = +8 mA", "2.7 V < VDD < 3.6 V"},
			},
			{
				SiUnit:     "volt",
				MaxValue:   .4,
				Conditions: []string{"TTL port", "IIO = +8 mA", "2.7 V < VDD < 3.6 V"},
			},
			{
				SiUnit:     "volt",
				MaxValue:   1.3,
				Conditions: []string{"IIO = +20 mA", "2.7 V < VDD < 3.6 V"},
			},
			{
				SiUnit:     "volt",
				MaxValue:   .4,
				Conditions: []string{"IO = +6 mA", "2 V < VDD < 2.7 V"},
			},
		},
	}
	props.Voh = &ValueOptions{
		Values: []Value{
			{
				SiUnit:                 "volt",
				RelativeValueReference: "VDD",
				RelativeValueModifier:  -.4,
				RelativeValueOperator:  "add",
				Conditions:             []string{"CMOS port", "IIO = +8 mA", "2.7 V < VDD < 3.6 V"},
			},
			{
				SiUnit:     "volt",
				MinValue:   2.4,
				Conditions: []string{"TTL port", "IIO = +8 mA", "2.7 V < VDD < 3.6 V"},
			},
			{
				SiUnit:                 "volt",
				RelativeValueReference: "VDD",
				RelativeValueModifier:  -1.3,
				RelativeValueOperator:  "add",
				Conditions:             []string{"IIO = +20 mA", "2.7 V < VDD < 3.6 V"},
			},
			{
				SiUnit:                 "volt",
				RelativeValueReference: "VDD",
				RelativeValueModifier:  -.4,
				RelativeValueOperator:  "add",
				Conditions:             []string{"IO = +6 mA", "2 V < VDD < 2.7 V"},
			},
		},
	}
	props.Imax = &ValueOptions{
		Values: []Value{
			{
				SiUnit:   "milliamp",
				MaxValue: 25,
			},
		},
	}
	props.Esd = true
	return props
}

func getElectricalProps(ioType string) pinSpec {
	props := pinSpec{}
	switch ioType {
	case "TTa", "TT":
		props = pinSpec{
			Vih: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "0.445VDD",
						RelativeValueModifier:  0.398,
						RelativeValueOperator:  "add",
					},
				},
			},
			Vil: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "0.3VDD",
						RelativeValueModifier:  0.07,
						RelativeValueOperator:  "add",
					},
				},
			},
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MaxValue: 4,
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "VSS",
						RelativeValueModifier:  -0.3,
						RelativeValueOperator:  "add",
					},
				},
			},
		}
		props = addVOLProps(props)
	case "FT", "FTf":
		props = pinSpec{
			Vih: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "0.5VDD",
						RelativeValueModifier:  0.2,
						RelativeValueOperator:  "add",
					},
				},
			},
			Vil: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: ".475VDD",
						RelativeValueModifier:  -0.2,
						RelativeValueOperator:  "add",
					},
				},
			},
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "VDD",
						RelativeValueModifier:  0.4,
						RelativeValueOperator:  "add",
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "VSS",
						RelativeValueModifier:  -0.3,
						RelativeValueOperator:  "add",
					},
				},
			},
		}
		props = addVOLProps(props)
	case "B":
		props = pinSpec{
			Vih: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "0.2VDD",
						RelativeValueModifier:  0.95,
						RelativeValueOperator:  "add",
					},
				},
			},
			Vil: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "0.3VDD",
						RelativeValueModifier:  -0.3,
						RelativeValueOperator:  "add",
					},
				},
			},
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MaxValue: 9,
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MinValue: 0,
					},
				},
			},
		}
		props = addVOLProps(props)
	case "RST":
		props = pinSpec{
			Vih: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "0.445VVDD",
						RelativeValueModifier:  0.398,
						RelativeValueOperator:  "add",
					},
				},
			},
			Vil: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: ".3VDD",
						RelativeValueModifier:  0.7,
						RelativeValueOperator:  "add",
					},
				},
			},
			InternalPullUp: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "ohm",
						MaxValue: 55,
						TypValue: 40,
						MinValue: 25,
					},
				},
			},
			Esd: true,
		}
	case "VDD":
		props = pinSpec{
			Vmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MinValue: 2,
						MaxValue: 3.6,
					},
				},
			},
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MaxValue: 4,
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MinValue: -.3,
					},
				},
			},
			Esd: true,
		}
	case "VBAT":
		props = pinSpec{
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MaxValue: 4,
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MinValue: -.3,
					},
				},
			},
			Vmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MinValue: 1.65,
						MaxValue: 3.6,
					},
				},
			},
			Esd: true,
		}
	case "VDDA":
		props = pinSpec{
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MaxValue: 4,
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MinValue: -.3,
					},
				},
			},
			Vmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:     "volt",
						MaxValue:   3.6,
						MinValue:   2,
						Conditions: []string{"OPAMP and DAC not used"},
					},
					{
						SiUnit:     "volt",
						MaxValue:   3.6,
						MinValue:   2.4,
						Conditions: []string{"OPAMP and DAC used"},
					},
				},
			},
			Esd: true,
		}
	default:
		props = pinSpec{
			Vih: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "VDD",
						RelativeValueModifier:  0.7,
						RelativeValueOperator:  "multiply",
					},
				},
			},
			Vil: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "VDD",
						RelativeValueModifier:  0.3,
						RelativeValueOperator:  "multiply",
					},
				},
			},
			AbsVmax: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "volt",
						MaxValue: 4,
					},
				},
			},
			AbsVmin: &ValueOptions{
				Values: []Value{
					{
						SiUnit:                 "volt",
						RelativeValueReference: "VSS",
						RelativeValueModifier:  -0.3,
						RelativeValueOperator:  "add",
					},
				},
			},
		}
		props = addVOLProps(props)

	}

	switch ioType {
	case "TC":
		props.Vmax = &ValueOptions{
			Values: []Value{
				{
					SiUnit:                 "volt",
					RelativeValueReference: "VDD",
					RelativeValueModifier:  0.3,
					RelativeValueOperator:  "add",
				},
			},
		}
	case "TTa":
		props.Vmax = &ValueOptions{
			Values: []Value{
				{
					SiUnit:                 "volt",
					RelativeValueReference: "VDDA",
					RelativeValueModifier:  0.3,
					RelativeValueOperator:  "add",
				},
			},
		}
	case "FT", "FTf", "B":
		props.Vmax = &ValueOptions{
			Values: []Value{
				{
					SiUnit:   "volt",
					MaxValue: 5.5,
				},
			},
		}
	case "TT":
		props.Vmax = &ValueOptions{
			Values: []Value{
				{
					SiUnit:   "volt",
					MaxValue: 3.6,
				},
			},
		}
	}

	switch ioType {
	case "TC":
		props.InputLeakage = &ValueOptions{
			Values: []Value{
				{
					SiUnit:   "microamp",
					MaxValue: 0.1,
				},
			},
		}
	case "TTa":
		props.InputLeakage = &ValueOptions{
			Values: []Value{
				{
					SiUnit:     "microamp",
					MaxValue:   0.1,
					Conditions: []string{"digital mode", "VSS ≤ VIN ≤ VDD"},
				},
				{
					SiUnit:     "microamp",
					MaxValue:   1,
					Conditions: []string{"digital mode", "VDD ≤ VIN ≤ VDDA"},
				},
				{
					SiUnit:     "microamp",
					MaxValue:   0.2,
					Conditions: []string{"analog mode", "VSS ≤ VIN ≤ VDDA"},
				},
			},
		}
	case "FT", "FTf":
		props.InputLeakage = &ValueOptions{
			Values: []Value{
				{
					SiUnit:     "microamp",
					MaxValue:   0.1,
					Conditions: []string{"digital mode", "VSS ≤ VIN ≤ VDD"},
				},
				{
					SiUnit:     "microamp",
					MaxValue:   10,
					Conditions: []string{"VDD ≤ VIN ≤ 5 V"},
				},
			},
		}
	}

	return props
}

func generateFunctionProps(name, extraFns1, extraFns2 string) []FunctionProperties {
	fnNameSlice := []string{name}
	fnNameSlice = append(fnNameSlice, strings.Split(extraFns1, ",")...)
	fnNameSlice = append(fnNameSlice, strings.Split(extraFns2, ",")...)
	var out []FunctionProperties
	for _, fn := range fnNameSlice {
		if fn == "" {
			continue
		}
		fnNoSpace := strings.TrimSpace(fn)
		fnProp := translateFnName(fnNoSpace)
		fnProp.PerFunctionName = fnNoSpace
		out = append(out, fnProp)
	}
	return out
}

var (
	vddRE  = regexp.MustCompile(`VDD_*`)
	vddaRE = regexp.MustCompile(`VDDA*`)
	vbatRE = regexp.MustCompile(`VBAT`)
	vssRE  = regexp.MustCompile(`VSS*`)

	gpioRE = regexp.MustCompile(`^P[A-Z]?\d\d?$`)

	compinRE   = regexp.MustCompile(`COMP\d_IN*`)
	compoutRE  = regexp.MustCompile(`COMP\d_OUT*`)
	opampinRE  = regexp.MustCompile(`OPAMP*VIN*`)
	opampoutRE = regexp.MustCompile(`OPAMP*VOUT*`)
	adcRE      = regexp.MustCompile(`ADC*`)

	uarttxRE  = regexp.MustCompile(`USART*_TX`)
	uartrxRE  = regexp.MustCompile(`USART*_RX`)
	uartRE    = regexp.MustCompile(`USART*`)
	i2csdaRE  = regexp.MustCompile(`I2C*_SDA`)
	i2csclRE  = regexp.MustCompile(`I2C*_SCL`)
	i2cRE     = regexp.MustCompile(`I2C*`)
	jtagRE    = regexp.MustCompile(`JT*`)
	spicopiRE = regexp.MustCompile(`SPI*_MOSI*`)
	spicipoRE = regexp.MustCompile(`SPI*_MISO*`)
	spiclkRE  = regexp.MustCompile(`SPI*_SCK*`)
	spicsRE   = regexp.MustCompile(`SPI*_NSS*`)
	usbRE     = regexp.MustCompile(`USB_*`)
	canRE     = regexp.MustCompile(`CAN_*`)
)

func translateFnName(fnName string) FunctionProperties {
	if uarttxRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "uart", PinUsage: "UART_TX", Direction: "out", ElectricalConfiguration: "push-pull"}
	} else if uartrxRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "uart", PinUsage: "UART_RX", Direction: "in"}
	} else if uartRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "uart"}
	} else if i2csdaRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "i2c", PinUsage: "I2C_SDA", Direction: "bidir", ElectricalConfiguration: "open-drain"}
	} else if i2csclRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "i2c", PinUsage: "I2C_SCL", Direction: "out", ElectricalConfiguration: "open-drain"}
	} else if i2cRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "i2c", ElectricalConfiguration: "open-drain"}
	} else if jtagRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "jtag"}
	} else if spicopiRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "spi", PinUsage: "SPI_COPI", Direction: "out", ElectricalConfiguration: "push-pull"}
	} else if spicipoRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "spi", PinUsage: "SPI_CIPO", Direction: "in"}
	} else if spiclkRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "spi", PinUsage: "SPI_CLK", Direction: "out", ElectricalConfiguration: "push-pull"}
	} else if spicsRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "spi", PinUsage: "SPI_CS", Direction: "in"}
	} else if vddRE.MatchString(fnName) || vddaRE.MatchString(fnName) || vbatRE.MatchString(fnName) {
		return FunctionProperties{ElectricalConfiguration: "power"}
	} else if vssRE.MatchString(fnName) {
		return FunctionProperties{ElectricalConfiguration: "ground"}
	} else if gpioRE.MatchString(fnName) {
		return FunctionProperties{Direction: "bidir"}
	} else if compinRE.MatchString(fnName) || adcRE.MatchString(fnName) || opampinRE.MatchString(fnName) {
		return FunctionProperties{Direction: "in", ElectricalConfiguration: "analog"}
	} else if compoutRE.MatchString(fnName) || opampoutRE.MatchString(fnName) {
		return FunctionProperties{Direction: "out", ElectricalConfiguration: "analog"}
	} else if usbRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "USB", Direction: "bidir", ElectricalConfiguration: "push-pull"}
	} else if canRE.MatchString(fnName) {
		return FunctionProperties{InterfaceType: "CAN", ElectricalConfiguration: "open-drain"}
	}
	return FunctionProperties{}
}

func sanitizeIOType(ioType, pinName string) string {
	knownIOTypes := map[string]bool{
		"TTa": true,
		"TT":  true,
		"TC":  true,
		"FTf": true,
		"FT":  true,
		"B":   true,
		"RST": true,
	}
	ioType = strings.TrimSpace(ioType)
	if ok := knownIOTypes[ioType]; ok {
		return ioType
	}

	if ioType == "-" {
		return strings.TrimSpace(pinName)
	}
	return ""
}

func generateCurrentConsumption() []CurrentConsumption {
	out := []CurrentConsumption{
		{
			SupplyName: "VDD",
			ActiveCurrent: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "milliamp",
						TypValue: 5.2,
						MaxValue: 7,
						Conditions: []string{
							"Run mode, executing from flash",
							"Internal clock (HSI)",
							"clock freq = 8 MHz",
							"max temp at 105 deg C",
							"All peripherals enabled",
						},
					},
					{
						SiUnit:   "milliamp",
						TypValue: 4.8,
						MaxValue: 6.5,
						Conditions: []string{
							"Run mode, executing from RAM",
							"Internal clock (HSI)",
							"clock freq = 8 MHz",
							"max temp at 105 deg C",
							"All peripherals enabled",
						},
					},
				},
			},
			SleepCurrent: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "milliamp",
						TypValue: 3.4,
						MaxValue: 5.1,
						Conditions: []string{
							"Run mode, executing from flash or RAM",
							"Internal clock (HSI)",
							"clock freq = 8 MHz",
							"max temp at 105 deg C",
							"All peripherals enabled",
						},
					},
				},
			},
			ShutdownCurrent: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "microamp",
						TypValue: 1.10,
						MaxValue: 12.60,
						Conditions: []string{
							"LSI OFF and IWDG OFF",
							"Vdd=3.6V",
							"max temp at 105 deg C",
						},
					},
				},
			},
		},
		{
			SupplyName: "VDDA",
			ActiveCurrent: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "microamp",
						TypValue: 69,
						MaxValue: 96,
						Conditions: []string{
							"Run mode, executing from flash or RAM",
							"Internal clock (HSI)",
							"clock freq = 8 MHz",
							"max temp at 105 deg C",
							"VDDA = 3.6V",
						},
					},
				},
			},
			ShutdownCurrent: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "microamp",
						TypValue: 1.09,
						Conditions: []string{
							"LSI OFF and IWDG OFF",
							"Vdd=3.6V",
							"max temp at 105 deg C",
						},
					},
				},
			},
		},
		{
			SupplyName: "VBAT",
			ActiveCurrent: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "microamp",
						TypValue: 0.82,
						Conditions: []string{
							"LSE & RTCON; Xtal mode lower driving capability; LSEDRV[1: 0] = 00",
							"VBAT = 3.6V",
						},
					},
				},
			},
		},
	}
	return out
}

func generatePackage() Package {
	out := Package{
		Length: &ValueOptions{
			Values: []Value{
				{
					SiUnit:   "millimeters",
					TypValue: 12,
				},
			},
		},
		Width: &ValueOptions{
			Values: []Value{
				{
					SiUnit:   "millimeters",
					TypValue: 12,
				},
			},
		},
		Height: &ValueOptions{
			Values: []Value{
				{
					SiUnit:   "millimeters",
					TypValue: 1.6,
				},
			},
		},
	}
	return out
}

// This Marshal function doesn't escape HTML so you can use < and >
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

func main() {
	stm32 := Component{
		ComponentID: ComponentID{
			PartType:      "microcontroller",
			Manufacturer:  "stmicroelectronics",
			ComponentName: "STM32F302R6",
			OrderableMPN:  []string{"STM32F302R6T6TR", "STM32F302R6T6xxx"},
			SourceDatasheetID: &SourceDatasheetID{
				PublishedDate: "March2020",
				Version:       "Rev 8",
				DatasheetURI:  "https://www.st.com/resource/en/datasheet/stm32f302r6.pdf",
				ProductURI:    "https://www.st.com/en/microcontrollers-microprocessors/stm32f302r6.html",
			},
			DigitalDatasheetID: &DigitalDatasheetID{
				PublishedDate:          "April2024",
				EDatasheetSpecRevision: "1.0",
			},
			Status:         "production",
			ComplianceList: []string{"Ecopack2"},
		},
		CoreProperties: Microcontroller{
			PartType: "microcontroller",
			OnChipFlash: &ValueOptions{
				Values: []Value{
					{
						SiUnit:     "byte",
						UnitName:   "Kilobytes",
						UnitFactor: 1000,
						TypValue:   32,
					},
				},
			},
			OnChipRAM: &ValueOptions{
				Values: []Value{
					{
						SiUnit:     "byte",
						UnitName:   "Kilobytes",
						UnitFactor: 1000,
						TypValue:   16,
					},
				},
			},
			CoreProcessor:        "Cortex-M4",
			CoreArchitectureBits: "32",
			ClockSpeed: &ValueOptions{
				Values: []Value{
					{
						SiUnit:   "megahertz",
						TypValue: 8,
					},
				},
			},
			CurrentConsumption: generateCurrentConsumption(),
		},
		Package: generatePackage(),
	}

	// Generate Pin Information
	// datasheetPins.csv was generated from:
	//	1) opening the datasheet pdf as a google doc
	//	2) copy-paste the pin information tables into a spreadsheet
	//	3) hand modify a few of the values for processing
	//	4) download the spreadsheet as: datasheetPins.csv
	// https://docs.google.com/spreadsheets/d/1Dr2eBk2hWs5tvedTAUk-lWu2t_hDv5RiGWgafALyFdA/edit?resourcekey=0-gkLZA0rFdx_pqpIIBQUYRw#gid=0
	file, err := os.Open("datasheetPins.csv")
	if err != nil {
		panic("Error while reading the file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Error reading records")
	}
	// Column number each of these pins is in
	pinNum := 3
	pinName := 4
	ioType := 6
	extraFns1 := 8
	extraFns2 := 9

	for _, eachrecord := range records[1:] {
		fmt.Println(eachrecord)
		ioTypeStr := sanitizeIOType(eachrecord[ioType], eachrecord[pinName])
		pin := getElectricalProps(ioTypeStr)
		pin.TerminalIdentifier = eachrecord[pinNum]
		pin.Name = eachrecord[pinName]
		pin.FunctionProperties = generateFunctionProps(eachrecord[pinName], eachrecord[extraFns1], eachrecord[extraFns2])
		pin.NumberOfSupportedFunctions = len(pin.FunctionProperties)
		stm32.Pins = append(stm32.Pins, pin)
	}

	b, err := JSONMarshal(stm32)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("STM32F302R6T6TR.json", b, 0644); err != nil {
		panic(err)
	}
}
