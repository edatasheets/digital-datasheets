package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Component struct {
	ComponentID    ComponentID `json:"componentID,omitempty"`
	CoreProperties interface{} `json:"coreProperties,omitempty"`
	Pins           []pinSpec   `json:"pins,omitempty"`
	Package        Package     `json:"package,omitempty"`
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
	TypValue               float64  `json:"typValue,omitempty"`
	MinValue               float64  `json:"minValue,omitempty"`
	MaxValue               float64  `json:"maxValue,omitempty"`
	RelativeValueReference string   `json:"relativeValueReference,omitempty"`
	RelativeValueModifier  float64  `json:"relativeValueModifier,omitempty"`
	RelativeValueOperator  string   `json:"relativeValueOperator,omitempty"`
	UnitName               string   `json:"unitName,omitempty"`
	UnitFactor             float64  `json:"unitFactor,omitempty"`
	Conditions             []string `json:"conditions,omitempty"`
}

type FunctionProperties struct {
	PerFunctionName         string    `json:"perFunctionName,omitempty"`
	InterfaceType           string    `json:"interfaceType,omitempty"`
	PinUsage                string    `json:"pinUsage,omitempty"`
	Direction               string    `json:"direction,omitempty"`
	ElectricalConfiguration string    `json:"electricalConfiguration,omitempty"`
	Polarity                string    `json:"polarity,omitempty"`
	PerFunctionProperties   *PinProps `json:"perFunctionProperties,omitempty"`
}

type pinSpec struct {
	TerminalIdentifier         []string             `json:"terminalIdentifier,omitempty"`
	Name                       string               `json:"name,omitempty"`
	StandardizedName           string               `json:"standardizedName,omitempty"`
	NumberOfSupportedFunctions int                  `json:"numberOfSupportedFunctions,omitempty"`
	Description                string               `json:"description,omitempty"`
	FunctionProperties         []FunctionProperties `json:"functionProperties,omitempty"`
	PinProperties              *PinProps            `json:"pinProperties,omitempty"`
}

type PinProps struct {
	Vih                *ValueOptions `json:"vih,omitempty"`
	Vil                *ValueOptions `json:"vil,omitempty"`
	Vol                *ValueOptions `json:"vol,omitempty"`
	Voh                *ValueOptions `json:"voh,omitempty"`
	AbsVmax            *ValueOptions `json:"absVmax,omitempty"`
	AbsVmin            *ValueOptions `json:"absVmin,omitempty"`
	Vmax               *ValueOptions `json:"vmax,omitempty"`
	Imax               *ValueOptions `json:"imax,omitempty"`
	InputLeakage       *ValueOptions `json:"inputLeakage,omitempty"`
	VoltageOptions     *ValueOptions `json:"voltageOptions,omitempty"`
	Esd                bool          `json:"esd,omitempty"`
	InternalPullUp     *ValueOptions `json:"internalPullUp,omitempty"`
	InternalPullDown   *ValueOptions `json:"internalPullDown,omitempty"`
	ExternalComponents *interface{}  `json:"externalComponents,omitempty"`
}

type OldComponent struct {
	ComponentID    ComponentID  `json:"componentID,omitempty"`
	CoreProperties interface{}  `json:"coreProperties,omitempty"`
	Pins           []oldPinSpec `json:"pins,omitempty"`
	Package        Package      `json:"package,omitempty"`
}

type oldPinSpec struct {
	TerminalIdentifier         string               `json:"terminalIdentifier,omitempty"`
	Name                       string               `json:"name,omitempty"`
	StandardizedName           string               `json:"standardizedName,omitempty"`
	NumberOfSupportedFunctions int                  `json:"numberOfSupportedFunctions,omitempty"`
	FunctionProperties         []FunctionProperties `json:"functionProperties,omitempty"`
	Description                string               `json:"description,omitempty"`
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
	InternalPullDown           *ValueOptions        `json:"internalPullDown,omitempty"`
	ExternalComponents         *interface{}         `json:"externalComponents,omitempty"`
}

func main() {
	path := "/home/rachelsn/Documents/digital-datasheets/examples"
	subDirs, err := os.ReadDir(path)
	if err != nil {
		panic("Error while reading files from directory")
	}
	for _, d := range subDirs {
		if d.IsDir() {
			subDirPath := filepath.Join(path, d.Name())
			fls, err := os.ReadDir(subDirPath)
			if err != nil {
				panic(fmt.Sprintf("Error while reading sub directory: %s: %w", subDirPath, err))
			}
			for _, f := range fls {
				if f.Name() == "STM32F302R6T6TR.json" || f.Name() == "diode_BAS16_output.json" {
					continue
				}
				fPath := filepath.Join(subDirPath, f.Name())
				fmt.Println(fPath)
				b, err := os.ReadFile(fPath)
				if err != nil {
					panic(fmt.Sprintf("Error while reading file %s: %w", f.Name(), err))
				}
				comp := OldComponent{}
				err = json.Unmarshal(b, &comp)
				if err != nil {
					// skip failures
					fmt.Printf("could not unmarshall: %w", err)
					continue
				}
				newComp := Component{}
				newComp.ComponentID = comp.ComponentID
				newComp.CoreProperties = comp.CoreProperties
				newComp.Package = comp.Package
				for _, pin := range comp.Pins {
					newPinSpec := pinSpec{
						TerminalIdentifier:         []string{pin.TerminalIdentifier},
						Name:                       pin.Name,
						StandardizedName:           pin.StandardizedName,
						Description:                pin.Description,
						NumberOfSupportedFunctions: pin.NumberOfSupportedFunctions,
						FunctionProperties:         pin.FunctionProperties,
						PinProperties: &PinProps{
							Vih:                pin.Vih,
							Vil:                pin.Vil,
							Vol:                pin.Vol,
							Voh:                pin.Voh,
							AbsVmax:            pin.AbsVmax,
							AbsVmin:            pin.AbsVmin,
							Vmax:               pin.Vmax,
							Imax:               pin.Imax,
							InputLeakage:       pin.InputLeakage,
							VoltageOptions:     pin.VoltageOptions,
							Esd:                pin.Esd,
							InternalPullUp:     pin.InternalPullUp,
							InternalPullDown:   pin.InternalPullDown,
							ExternalComponents: pin.ExternalComponents,
						},
					}
					if fmt.Sprintf("%v", newPinSpec.PinProperties) == fmt.Sprintf("%v", &PinProps{}) {
						newPinSpec.PinProperties = nil
					}

					newComp.Pins = append(newComp.Pins, newPinSpec)
				}

				out, err := json.MarshalIndent(newComp, "", "  ")
				if err != nil {
					panic(err)
				}

				err = os.WriteFile(f.Name(), out, 0644)
				if err != nil {
					panic(err)
				}
			}
		}

	}

}
