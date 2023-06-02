
## Contents
* [Specification: component level](part-spec/component.json)
* [Specification: pin level](part-spec/definitions.json)
* [Component examples](support-docs/examples.md)


## Background
### 1 Introduction 
As the demand for hardware design automation tools increases, there is a need for machine readable datasheets. 
Defining a common digital datasheet specification reduces the burden for component vendors to deliver multiple datasheets to support different tools and encourages reuse of tools across multiple designs. 

#### 1.1 Objectives
Part manufacturers create datasheets to document component information including performance, electrical characteristics, size, orientation, packaging, etc. Digital datasheets contain this information in a machine readable format.The objective of this specification is to create a uniform, machine-readable format for representing component datasheets.

#### 1.2 Scope
This document is intended for digital datasheet producers such as components vendors as well as people who will consume these digital datasheets to automate designs such as tool creators.The specification will include common classes of components used in designs. Some examples are included in the appendices to provide context for future additions to this specification.

#### 1.3 Keywords
- Required: The field is required in the component digital datasheet.
- Optional: The field may or may not be included in the component digital datasheet.
- Shall: Indicates a mandatory requirement.
- May: Indicates an optional requirement.

#### 1.4 References
- The JSON data interchange syntax, [ECMA-404, 2nd edition, December 2017](https://www.ecma-international.org/publications-and-standards/standards/ecma-404/)
- JSON Schema: A media Type for Describing JSON Documents, [draft-bhutton-json-schema-00, December 2020](https://datatracker.ietf.org/doc/html/draft-bhutton-json-schema-00).
- Terms, Definitions, and Letter Symbols for Microelectronic Devices, JESD99C – December 2012

### 2. Use Cases 
Digital datasheets can be used for many applications. The list of use cases included here is not meant to be exhaustive and it is expected that new applications will be developed as more people start using these digital datasheets. Sample applications include automated hardware design checks to identify bugs earlier in the design cycle, automated hardware designs to speed up board development, components comparison to identify replacement components on a design. 

### 3.Specifications

#### 3.1 Digital Datasheet Format
A digital datasheet shall be written in the JSON (JavaScript Object Notation) format. The specification is written in JSON schema to facilitate validation of a JSON datasheet. 

#### 3.2 Required Information.
A digital datasheet shall include the following information:
- The manufacturer’s name
- The Manufacturer Part Number (MPN)
- Information to identify the source datasheet
- The datasheet version given by date or GUID 
- The component part type. See appendix for examples.
- List of Component Pins as defined by the specification.

#### 3.3 Date Format
A digital datasheet shall follow the international standard notation, YYYY-MM-DD.

####  3.4.1	 ComponentID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|manufacturer|company that manufactures the part|String|Yes|
|componentName|base part name that describes the form and fit of a component|String| |
|orderableMPN|orderable manufacturer part number|String|Yes|
|sourceDatasheetID|methods for identifying the human-readable source information for a digital datasheet|#/definitions/sourceDatasheetID|Yes|
|digitalDatasheetID|methods for identifying the version of the digital datasheet|#/definitions/digitalDatasheetID|Yes|
|status|production status of a component|String| |
|complianceList|list of standards the part complies with|array of String| |

####  3.4.2	 ComponentProtectionThresholds

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|thermalShutdownThresholdRising|Thermal Shutdown (tsd) Threshold with temperature rising|definitions.json#/unit| |
|thermalShutdownThresholdFalling|Thermal Shutdown (tsd) Threshold with temperature falling|definitions.json#/unit| |
|thermalShutdownHysteresis|Thermal Shutdown (tsd) Hysteresis|definitions.json#/unit| |
|powerSupplyProtection|undervoltage lockout, overvoltage protection thresholds of a supply|array of #/defs/powerSupplyProtection| |

####  3.4.3	 ConditionalProperty

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|value of property|#/$definitions/unit| |
|conditions|conditions under which the property is measured|array of String| |

####  3.4.4	 CurrentConsumption

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|supplyName|Name of the power supply |String| |
|quiescentCurrent|quiescent current (Iq) of a device|definitions.json#/unit| |
|shutdownCurrent|shutdown current (Isd) of a device|definitions.json#/unit| |
|activeCurrent|current consumption when a device is in active mode|definitions.json#/unit| |
|sleepCurrent|current consumption when a device is in sleep mode|definitions.json#/unit| |
|idleCurrent|current consumption when a device is in idle mode|definitions.json#/unit| |

####  3.4.5	 DigitalDatasheetID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|publishedDate|date the digital datasheet was published|String| |
|guid|vendor defined guid (see https://www.guidgenerator.com/) to uniquely identify digital datasheet version|String| |

####  3.4.6	 ExternalComponents

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentType|type of external component required to be connected to a pin|String|Yes|
|configuration|electrical configuration of component connected to pin with respect to the pin|String|Yes|
|minValue|minimum value of component if a range is specified|#/definitions/unit| |
|maxValue|maximum value of component if a range is specified|#/definitions/unit| |
|value|value of component if a range is not specified|#/definitions/unit| |

####  3.4.7	 ExternalFile

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|fileDescription|text description of the contents of an external file (or files)|String| |
|companionSoftware|optional, name of software program used to access file|String| |
|standardReferenced|optional, name of the standard the file is written in|String| |
|fileExtension|type of file referenced by link|array of String| |
|fileName|name of external file within the zipped datasheet package|array of String| |
|fileURI|URI linking to the CAD file|array of String| |

####  3.4.8	 Package

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|length|length of a side of a package|#/$definitions/unit| |
|width|width of a side of a package|#/$definitions/unit| |
|height|height of a package|#/$definitions/unit| |
|nominalFootprints|references to external footprints in standard CAD formats|array of #/$definitions/externalFile| |
|breakoutExamples|references to external board file that contains layout breakout example|array of #/$definitions/externalFile| |
|partModelInformation|reference to an external XML file that contains a part model in IPC/DAC2552 format|#/$definitions/externalFile| |
|standardPackageSize|name of standard package size (imperial)|Must set either Ref or Type| |
|standardPackageType|name of standard package types|Must set either Ref or Type| |

####  3.4.9	 PinFunction

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|perFunctionName|name of the function of a pin|String| |
|interfaceType|type of interface enabled by pin|String| |
|pinUsage|standardized usage of pin|String| |
|direction|direction of a pin's function|String| |
|electricalConfiguration|electrical configuration of a pin|String| |
|polarity|whether the active state of a pin is high or low|String| |

####  3.4.10	 PinSpec

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|terminalIdentifier|pin or ball number as defined by datasheet|String|Yes|
|name|name given to the signal appearing at the terminal of a component|String|Yes|
|standardizedName|standard name of pin|String| |
|description|description of the signal appearing at the terminal of an electric/electronic component|String| |
|numberOfSupportedFunctions|the total number of functions supported by this pin|Number| |
|functionProperties|list of function objects that can apply to an individual pin|array of #/definitions/pinFunction| |
|vihMin|the least positive (most negative) value of high-level input voltage for which operation of the logic element within specification limits is to be expected|#/definitions/unit| |
|vihMax|the most positive (least negative) value of high-level input voltage for which operation of the logic element within specification limits is to be expected|#/definitions/unit| |
|vilMax|the most positive (least negative) value of low-level input voltage for which operation of the logic element within specification limits is to be expected|#/definitions/unit| |
|vilMin|the least positive (most negative) value of low-level input voltage for which operation of the logic element within specification limits is to be expected|#/definitions/unit| |
|vol|the voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a low level at the output|#/definitions/unit| |
|voh|the voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a high level at the output|#/definitions/unit| |
|absVmax|maximum voltage rating beyond which damage to the device may occur|#/definitions/unit| |
|absVmin|absolute minimum voltage that can be applied to a pin|#/definitions/unit| |
|vmax|maximum continuous voltage that can safely be applied to a pin|#/definitions/unit| |
|imax|maximum continuous current that can safely be drawn from a pin|#/definitions/unit| |
|inputLeakage|maximum current draw into a high impedance input pin|#/definitions/unit| |
|outputLeakage|maximum current flow from a pin during the off state|#/definitions/unit| |
|dcResistance|resistance of a pin of a connector|#/definitions/unit| |
|voltageOptions|list of voltage levels supported by a pin|array of #/definitions/unit| |
|floatUnused|description of whether pin can safely be floated if it is not used|Boolean| |
|internalPullUp|indicates the value of an internal pull-up on a pin|#/definitions/unit| |
|internalPullDown|indicates the value of an internal pull-down on a pin|#/definitions/unit| |
|esd|indicates whether ESD protection exists on a pin|Boolean| |
|externalComponents|list of external component structures recommended to be attached to a pin|array of #/definitions/externalComponents| |

####  3.4.11	 Ratio

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numerator|numerator of ratio|Number| |
|denominator|denominator of ratio|Number| |

####  3.4.12	 SourceDatasheetID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|publishedDate|date the source datasheet was published|String| |
|version|version of the source datasheet|String| |
|datasheetURI|uri to the source datasheet pdf or html view|String| |
|productURI|uri to the source datasheet's product page'|String| |

####  3.4.13	 Unit

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|siUnit|name of SI unit of measure|String| |
|typValue|typical unit quantity corresponding to unit text - example 40mV would have a value of 40|Number| |
|minValue|minimum unit quantity corresponding to unit text - example 40mV would have a value of 40|Number| |
|maxValue|maximum unit quantity corresponding to unit text - example 40mV would have a value of 40|Number| |
|unitText|human readable text describing value - example 40mV would have a value of mV|String| |
|unitFactor|multiplier on the value to achieve the SI unit - example for 40mV the unitFactor would be 0.001|Number| |
|relativeValueReference|if unit quantity is based on another reference, value of the reference|String| |
|relativeValueModifier|if a unit quantity is based on another reference, the value that edits that reference|Number| |
|relativeValueOperator|if a unit quantity is based on another reference, the operation that is performed with the modifier|String| |
|valueDefined|a boolean representing whether a value has been defined|Boolean| |

####  3.4.14	 Value

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|typValue|typical value|Number| |
|minValue|minimum value|Number| |
|maxValue|maximum value|Number| |

#### 3.5	 Component

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|methods for identifying the version of the digital datasheet|definitions.json#/componentID|Yes|
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.6.1	 Curve

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|label|description of the data in a curve|String| |
|xData|x value of data being plotted|array of Number| |
|yData|y value of data being plotted|array of Number| |

####  3.6.2	 GraphDefiniton

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|title|title of a graph|String| |
|xUnits|x-axis units|String| |
|xLabel|x-axis title|String| |
|yUnits|y-axis units|String| |
|yLabel|y-axis title|String| |
|numberOfCurves|total number of curves in graph|Number| |
|data|array of curve objects representing actual data being plotted|array of #/graph/curve| |

####  3.7.1	 RegisterDefiniton

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|registerName|Name of a register|String|Yes|
|registerLongName|Full Name of a register|String| |
|registerAddressOffset|Address of a register|String|Yes|
|registerSize|Size of a register|definitions.json#/unit|Yes|
|registerType|Type of a register|String| |
|registerResetValue|Reset value of a register|String| |
|registerValue|Value of a register|String| |
|registerAccessType|Access type of a Register|String| |
|registerBitField|Describes the bit fields in the register|#/registerBitField| |

### 3.8	 Passives

Source: [passives.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/passives.json)

####  3.8.1	 Resistor

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|value|resistor value|definitions.json#/unit|Yes|
|tolerance|nominal tolerance of a resistor|definitions.json#/unit| |
|powerRating|measure of power a resistor can dissipate indefinitely without degrading performance|definitions.json#/unit| |
|temperatureCoefficient|change in resistance when the temperature is changed|Number| |
|maxOverloadVoltage|maximum voltage that can be applied to the resistor for a short period of time|definitions.json#/unit| |
|maxLimitingElementVoltage|maximum voltage value that can be applied continuously to the resistor|definitions.json#/unit| |
|minTemperature|minimum temperature under which a resistor can be expected to reliably operate|definitions.json#/unit| |
|maxTemperature|maximum temperature under which a resistor can be expected to reliably operate|definitions.json#/unit| |
|resistorDerating|graph object to capture resistance changes with temperature|definitions.json#/graphDefiniton| |
|package|package size of resistor|graph.json#/package| |

####  3.8.2	 Capacitor

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|value|capacitor value|definitions.json#/unit|Yes|
|tolerance|nominal tolerance of a capacitor|definitions.json#/unit| |
|ratedVoltage|maximum voltage which may be applied continuously to a capacitance|definitions.json#/unit| |
|dielectric|dielectric material used in the capacitor|String| |
|polarized|describes whether the capacitor is polarized|Boolean| |
|equivalentSerieResistance|equivalent series resistance (ESR) of the capacitor|definitions.json#/unit| |
|temperatureCoefficient|change in capacitance when the temperature is changed|Number| |
|minTemperature|minimum temperature under which a capacitor can be expected to reliably operate|definitions.json#/unit| |
|maxTemperature|maximum temperature under which a capacitor can be expected to reliably operate|definitions.json#/unit| |
|capacitorDerating|graph object to capture capacitance changes with voltage|definitions.json#/graphDefiniton| |
|package|package size of capacitor|definitions.json#/package| |

####  3.8.3	 Inductor

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|value|inductor value|definitions.json#/unit|Yes|
|tolerance|nominal tolerance of an inductor|definitions.json#/unit| |
|ratedCurrent|maximum continuous current the inductor can handle|definitions.json#/unit| |
|saturationCurrent|current where the inductor enters the magnetic state, and the inductance drops a specified amount|definitions.json#/unit| |
|rmsCurrent|DC current that produces an inductor temperature rise of 40 degrees Celsius|definitions.json#/unit| |
|selfResonantFrequency|frequency at which the inductor becomes a capacitor|definitions.json#/unit| |
|dcResistance|DC resistance of the inductor|definitions.json#/unit| |
|temperatureCoefficient|change in inductance when the temperature is changed|Number| |
|minTemperature|minimum temperature under which a inductor can be expected to reliably operate|definitions.json#/unit| |
|maxTemperature|maximum temperature under which a inductor can be expected to reliably operate|definitions.json#/unit| |
|saturationCurve|graph object to capture inductor saturation with current|definitions.json#/graphDefiniton| |
|resonantFrequencyCurve|graph object to capture inductor resonant frequency|definitions.json#/graphDefiniton| |
|package|package size of inductor|definitions.json#/package| |

####  3.8.4	 Common Mode Choke

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|diffModeCutoff|frequency at which the differential mode attenuation equals -3dB|definitions.json#/unit| |
|commonModeAttenuation|graph object to capture common mode attenuation of a common mode choke at various frequencies|definitions.json#/graphDefiniton| |
|dcResistance|dc resistance (DCR) of a common mode choke|definitions.json#/unit| |
|rmsCurrent|applied DC current (IRMS) that produces a common mode choke temperature rise of 40 deg C|definitions.json#/unit| |
|intendedApplication|intended application of a particular common mode choke|String| |
|package|package size of device|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.8.5	 Ferrite Bead

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|dcResistance|dc resistance (DCR) of ferrite bead|definitions.json#/unit| |
|rmsCurrent|applied DC current (IRMS) that produces a ferrite bead temperature rise of 40 deg C|definitions.json#/unit| |
|impedance100MHz|impedance of ferrite bead under standard testing conditions at 100MHz|definitions.json#/unit| |
|impedanceTolerance|variation of ferrite bead impedance expressed as +/- percentage|definitions.json#/unit| |
|package|package size of device|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.9	 Power

Source: [power.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/power.json)

####  3.9.1	 Switching Regulator

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|regulatorTopology|switching voltage regulator topology|String|Yes|
|vin|input voltage under which a device can be expected to operate without output dropping|definitions.json#/unit|Yes|
|vout|output voltage a device can regulate|definitions.json#/unit|Yes|
|feedbackVoltage|voltage comparison point at the feedback node|definitions.json#/unit| |
|loadCurrent|load current supported by a device|definitions.json#/unit|Yes|
|voutAccuracy|output voltage variation at no load|definitions.json#/unit| |
|loadRegulation|output voltage variation from no load to full load |definitions.json#/unit| |
|lineRegulation|output voltage variation from minimum input voltage to maximum input voltage |definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|switchingFrequency|switching frequency (fsw) of voltage regulator|definitions.json#/unit| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|definitions.json#/unit| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|definitions.json#/unit| |
|integratedFets|whether the regulator contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|#/$defs/powerFetProperities| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|definitions.json#/componentProtectionThresholds| |
|efficiency|power efficiency of regulator|graph.json#/graphDefiniton| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.9.2	 Ldo

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|vin|input voltage under which the part can be expected to operate without the output dropping|definitions.json#/unit|Yes|
|vout|output voltage the part can regulate|definitions.json#/unit|Yes|
|feedbackVoltage|voltage comparison point at the feedback node (vref)|definitions.json#/unit| |
|dropoutVoltage|dropout voltage of a device|Number| |
|loadCurrent|load current supported by a device|definitions.json#/unit|Yes|
|currentLimit|sustained output current threshold beyond which the output of a device starts drooping|definitions.json#/unit| |
|voutAccuracy|output voltage variation at no load of a device|definitions.json#/unit| |
|loadRegulation|output voltage variation,from no load to full load, of a device |definitions.json#/unit| |
|lineRegulation|output voltage variation,from minimum input voltage to maximum input voltage, of a device |definitions.json#/unit| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|definitions.json#/unit| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|definitions.json#/unit| |
|powerSupplyRejectionRatio|graph object to capture Power Supply Rejection Ratio (PSRR) of device over various frequencies|definitions.json#/graphDefiniton| |
|rmsOutputNoise|graph object to capture RMS output noise of device over various frequencies|definitions.json#/graphDefiniton| |
|totalOutputNoise|total output noise of a device|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|definitions.json#/componentProtectionThresholds| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.9.3	 Load Switch

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|fetType|type of pass FET in a device|String| |
|loadSwitchCount|number of load switched in the package.|Number| |
|vin|input voltage under which a device can be expected to reliabily operate|definitions.json#/unit|Yes|
|outputCurrent|continuous DC cuurent supported by a device|definitions.json#/unit|Yes|
|onResistance|FET on state resistance|definitions.json#/conditionalProperty|Yes|
|pdResistance|pull-down resistance of a device from the output to the ground|definitions.json#/unit| |
|currentLimitSupport|whether a device supports current limiting|Boolean| |
|adjustableRiseTimeSupport|whether a device supports adjustable rise time|Boolean| |
|quickOutputDischargeSupport|whether a device supports quick output discharge|Boolean| |
|reverseCurrentBlockingSupport|whether a device supports reverse current blocking|Boolean| |
|powerGoodSupport|whether a device supports power good|Boolean| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|definitions.json#/conditionalProperty| |
|offTime|time between enable deasserted and output voltage falling to 90% nominal|definitions.json#/conditionalProperty| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|definitions.json#/conditionalProperty| |
|fallTime|time for output voltage to go from 90% vout nominal to 10% vout nominal|definitions.json#/conditionalProperty| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|definitions.json#/componentProtectionThresholds| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.9.4	 Pmic

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|ldoRegulatorCount|number of ldos in the device|Number| |
|buckRegulatorCount|number of buck regulators in the device|Number| |
|boostRegulatorCount|number of boost regulators in the device|Number| |
|buckBoostRegulatorCount|number of buck-boost regulators in the device|Number| |
|adcCount|number of analog to digital converters in the device|Number| |
|externalClockCount|number of external clocks the device requires|Number| |
|internalClockCount|number of clocks/oscillators in the device|Number| |
|loadSwitchCount|number of load switches in the device|Number| |
|usbSwitchCount|number of USB switches in the device|Number| |
|componentList|List, by title, of components in the device|definitions.json#/component| |
|instances|definition of each instance of a component in the device|array of #/powerComponentDefinitions| |
|vin|input voltage under which a device can be expected to reliabily operate|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|definitions.json#/componentProtectionThresholds| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.9.5	 Display Backlight Driver

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|vin|input voltage under which a device can be expected to operate properly|definitions.json#/unit|Yes|
|vout|output voltage a device can regulate|definitions.json#/unit| |
|ioutPerString|output current per string a device can regulate|definitions.json#/unit| |
|ioutAccuracy|accuracy of per string current regulated by a device|definitions.json#/unit| |
|fsw|switching frequency of a device|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|integratedFets|whether a device contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|#/$defs/powerFetProperities| |
|currentMatchingAccuracy|current matching between LED strings|definitions.json#/unit| |
|dimmingSupport|whether a device supports output current dimming|Boolean| |
|dimmingControl|whether a device is dimmed by PWM or I2C|String| |
|dimmingFrequency|dimming frequency of a device|definitions.json#/unit| |
|dimmingRatio|dimming ratio of a device|definitions.json#/ratio| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|icProtection.json#/componentProtectionThresholds| |
|efficiency|efficiency vs forward current|graph.json#/graphDefiniton| |
|package|package size|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.9.6	 Battery Charger

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|chargerType|battery charger type|String| |
|converterType|switching charger type|String| |
|chargerTopology|type of battery charger topology (Narrow VDC vs Hybrid Power Boost)|String| |
|batteryConfig|battery configuration supported by the device|array of String| |
|batteryCellChemistry|battery cell chemistry supported by the device|array of String| |
|inputPowerSource|input power source supported by the device|array of String| |
|inputCurrentAccuracy|accuracy of input current when set|definitions.json#/unit| |
|batteryChargeCurrent|charging current of a device|definitions.json#/unit|Yes|
|batteryChargeCurrentAccuracy|charging current regulation accuracy of a device|definitions.json#/unit| |
|batteryPreChargeCurrent|charging current of a device in pre-charge phase|definitions.json#/unit| |
|batteryPreChargeCurrentAccuracy|pre-charging current regulation accuracy of a device|definitions.json#/unit| |
|batteryTrickleChargeCurrent|charging current of a device in trickle charge phase|definitions.json#/unit| |
|batteryTrickleChargeCurrentAccuracy|Trickle charging current regulation accuracy of a device|definitions.json#/unit| |
|batteryTerminationChargeCurrent|charging current of a device in charge termination phase|definitions.json#/unit| |
|batteryTerminationChargeCurrentAccuracy|termination charging current regulation accuracy of a device|definitions.json#/unit| |
|batteryChargeVoltage|battery charge voltage regulated by a device|definitions.json#/unit| |
|batteryChargeVoltageAccuracy|accuracy of battery charge voltage regulated by a device|definitions.json#/unit| |
|efficiency|charge efficiency vs charge current of a device|graph.json#/graphDefiniton| |
|vin|input voltage under which a device can be expected to reliabily operate|definitions.json#/unit| |
|fsw|switching frequency of a device|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|definitions.json#/componentProtectionThresholds| |
|batteryChargerProtections|battery charger specific protections supported by device|array of String| |
|integratedLoadSwitch|whether the device contains integrated power path load switch(es)|Boolean| |
|integratedFets|whether the device contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|#/$defs/powerFetProperities| |
|gateCapacitance|describes gate capacitance supported on external fets|definitions.json#/unit| |
|inputSenseResistor|describes intput sense resistor value|definitions.json#/unit| |
|batterySenseResistor|describes battery sense resistor value|definitions.json#/unit| |
|passThroughMode|whether pass through mode is supported|Boolean| |
|bc12Support|whether bc 1.2 detection is built in|Boolean| |
|tcpcSupport|whether type-C port controller support is built in|Boolean| |
|usbTypecRevision|usb type-c spec revision supported by a device|String| |
|pdVersion|version of power delivery spec supported by a device|String| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.9.7	 PowerComponentDefinitions

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentTitle|title used for the component in the digital datasheets specifications|String| |
|instanceName|name of component instances|String| |
### 3.10	 Hardware

Source: [hardware.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/hardware.json)

####  3.10.1	 Switch

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|property describing the way in which the switch is activated|String| |
|contactType|property describing the order in which switch contact is made and broken|String| |
|circuitConfig|property describing the number of poles and throws in a switch|String| |
|cycleRating|number of on/off cycles a mechanical switch can reliably sustain|Number| |
|voltageRating|maximum DC voltage potential that can be applied across an open switch|definitions.json#/unit| |
|currentRating|maximum DC current that can flow through a closed switch without causing excessive heating|definitions.json#/unit| |
|onResistance|nominal resistance of a closed switch|definitions.json#/unit|Yes|
|dielectricRating|maximum AC voltage potential that can be applied across an open switch for one minute|definitions.json#/unit| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.10.2	 Connector

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|function|intended function of a connector|String| |
|contactCount|number of contacts in a connector|Number| |
|type|property describing the method of mating to the connector|String| |
|cycleRating|number of plug/unplug cycles a connector is rated to support|Number| |
|pitch|distance from the center of one contact on the connector to the center of the next contact|definitions.json#/unit| |
|keying|property describing whether a connector has an asymmetry to prevent it from being plugged in the wrong direction|Boolean| |
|numberPins|number of pins on the connector|Number|Yes|
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.11	 IC IO

Source: [ic io.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/ic io.json)

####  3.11.1	 Redriver

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|numberChannels|number of lanes (single ended or differential) supported by redriver|Number| |
|interface|list of interface(s) supported by redriver|array of String| |
|maxDataRate|maximum data rate supported by redriver|definitions.json#/unit| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.11.2	 Retimer

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|numberOfLanes|number of lanes supported by a device|Number| |
|interface|list of interface(s) supported by a device|array of String| |
|dpMaxLinkRate|Max link rate supported by DP interface of a device|String| |
|maxDataRate|maximum data rate supported by a device|definitions.json#/unit| |
|integratedAuxLsSwitch|whether the AUX/LSx switch for SBU is integrated|Boolean| |
|commonClock|whether a device supports common reference clock|Boolean| |
|sris|whether a device supports Seperate Reference clock with Independent Spread spectrum clocking(SRIS)|Boolean| |
|srns|whether a device supports Seperate Reference clock with No Spread spectrum clocking (SRNS)|Boolean| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.11.3	 Bridge Chip

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|muxRatio|ratio of inputs to outputs|String| |
|inputInterfaces|list of interfaces at the input of the bridge|String| |
|outputInterfaces|list of interfaces at the output of the bridge|array of String| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.11.4	 Mux

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|muxRatio|ratio of inputs to outputs|definitions.json#/ratio| |
|inputInterfaces|list of interfaces mux is designed for|array of String| |
|insertionLoss|insertion loss through mux|Number| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.11.5	 Level Shifter

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|inputVoltage|input voltage level of level shifter|definitions.json#/unit| |
|outputVoltage|output voltage level of level shifter|definitions.json#/unit| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.12	 IC Logic

Source: [ic logic.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/ic logic.json)

####  3.12.1	 Logic Gate

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|logical operation performed by logic gate|String|Yes|
|numberGates|number of logical gates encapsulated in logic IC|Number| |
|schmittTrigger|property describing whether logic gate has schmitt trigger inputs|Boolean| |
|propagationDelay|time between input changing to output changing|definitions.json#/unit| |
|rampTime|time for output to go from 10% nominal output voltage to 90% nominal output voltage|definitions.json#/unit| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.12.2	 Clock

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|fixedFrequency|clock frequency value if the clock has a fixed frequency|definitions.json#/unit| |
|numberClockOutputs|number of clock outputs in a clock IC|Number| |
|diffSingleEnded|property describing whether a clock output is single ended or differential|String| |
|jitter|cycle to cycle clock jitter|definitions.json#/unit| |
|frequencyTolerance|amount of frequency variation specced from nominal frequency|definitions.json#/unit| |
|powerSupplyRejectionRatio|power supply rejection ratio (PSRR)or ratio between power supply variation and output variation|definitions.json#/unit| |
|outputFormat|signal format of clock output|String| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.13	 IC Micro

Source: [ic micro.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/ic micro.json)

####  3.13.1	 Microcontroller/ec

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|onChipFlash|quantity of built-in flash in a microprocessor|definitions.json#/unit| |
|onChipRAM|quantity of built-in RAM in a microprocessor|definitions.json#/unit| |
|onChipROM|quantity of built-in ROM in a microprocessor|definitions.json#/unit| |
|coreProcessor|description of core processor|String| |
|coreArchitectureBits|number of bits of data a CPU can transfer per clock cycle|String| |
|clockSpeed|speed of main CPU clock|definitions.json#/unit| |
|firmwareVersion|firmware version of the part|String| |
|activePower|average power of device in active state|definitions.json#/unit| |
|standbyPower|average power of device in standby state|definitions.json#/unit| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

### 3.14	 IC Misc

Source: [ic misc.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/ic misc.json)

####  3.14.1	 Speaker Amplifier

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|dataLength|number of bits in a data word|Number| |
|outputPower|typical output power from speaker amplifier|definitions.json#/conditionalProperty| |
|efficiency|typical speaker amplifier efficiency|definitions.json#/conditionalProperty| |
|thd+n|typical total harmonic distortion plus noise of amplifier|definitions.json#/conditionalProperty| |
|sampleRate|sample rate of data out of amplifier|definitions.json#/unit| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.14.2	 Audio Codec

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|dataLength|number of bits in a data word|Number| |
|hpOutputSNR|headphone amplifier output SNR|Number| |
|hpOutputTHD+N|headphone output total harmonic distortion plus noise|definitions.json#/unit| |
|micInputSNR|microphone input SNR|definitions.json#/unit| |
|micInputTHD+N|microphone input total harmonic distortion plus noise|definitions.json#/unit| |
|jackDetect|describes whether headphone jack detection is supported|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.14.3	 WLAN Module

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|wlanSpec|version of wlan specification supported by module|String| |
|bluetoothVersion|version of bluetooth supported by module|String| |
|txrxChains|number of tx and rx chains in a wifi module|String| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|lteCoexFilter|describes whether module supports lte coexistance filtering|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.14.4	 WWAN Module

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|networkSupport|networks supported by wwan module|String| |
|gpsSupport|whether wwan module has gps support|Boolean| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.14.5	 Tpm

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|interface|describes the communication interface from the chip to the host|array of String| |
|package|component's package size and description|definitions.json#/package| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.15	 Storage/memory

Source: [memory.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/storage/memory.json)

####  3.15.1	 SSD

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|type of ssd storage as defined by interface and technology|String| |
|capacity|capacity of SSD|definitions.json#/unit|Yes|
|dataRate|maximum data rate|definitions.json#/unit| |
|interface|interface of ssd to host|String| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.15.2	 SD Card

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|type of sd card|String| |
|capacity|capacity of SD card|definitions.json#/unit|Yes|
|dataRate|maximum data rate|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.15.3	 Dram

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|type of dram|String| |
|capacity|capacity of dram chip|definitions.json#/unit|Yes|
|diesPerChip|number of dies of dram chip|Number| |
|channelsPerDie|number of channels per die of dram chip|Number| |
|banksPerChannel|number of banks per channel of dram|Number| |
|bitsPerChannel|channel density of dram|definitions.json#/unit| |
|bitsPerDie|total die density of dram|definitions.json#/unit| |
|pageSize|page size of dram|definitions.json#/unit| |
|rows|number of rows per channel of dram|Number| |
|colunms|number of columns per row of dram|Number| |
|speed|dram maximum speed|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.15.4	 Rom

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|capacity|capacity of rom|definitions.json#/unit| |
|interface|interface of rom to host|String| |
|qeStatus|indicates whether the QE bit is set|Boolean| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.15.5	 Eeprom

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|capacity|capacity/density of eeprom|definitions.json#/unit|Yes|
|numberOfWords|number of rows in the eeprom|Number| |
|bitsPerWords|number of columns in the eeprom|Number| |
|bootBlockSize|size of the eeprom boot block|definitions.json#/unit| |
|interface|interface of eeprom to host|String| |
|clockFrequency|eeprom clock frequency|definitions.json#/unit| |
|accessTime|time to access the eeprom|definitions.json#/unit| |
|endurance|time in years a bit in the eeprom can retain its data state |Number| |
|dataRetention|maximum number of read and write cycle the part can support|Number| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.15.6	 Flash Memory

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|capacity|capacity/density of flash memory|definitions.json#/unit|Yes|
|pageSize|page size of flash memory|definitions.json#/unit| |
|blockSize|block size of flash memory|definitions.json#/unit| |
|bootBlockSize|size of the flash memory boot block|definitions.json#/unit| |
|interface|interface of flash memory to host|String| |
|clockFrequency|flash memory clock frequency|definitions.json#/unit| |
|blockEraseTime|time it takes to erase a block (largest erasable unit) of the flash memory|definitions.json#/unit| |
|sectorEraseTime|time it takes to erase a sector (smallest erasable unit) of the flash memory|definitions.json#/conditionalProperty| |
|chipEraseTime|time it takes to erase the flash memory|definitions.json#/unit| |
|pageProgramTime|time it takes to program a page of the flash memory|definitions.json#/unit| |
|endurance|time in years a bit in the eeprom can retain its data state |Number| |
|dataRetention|maximum number of read and write cycle the part can support|Number| |
|hwReset|whether the part supports a hardware reset pin|Boolean| |
|writeProtect|whether the part has a write protect pin|Boolean| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|component's package size and description|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.16	 Usbc

Source: [usbc.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/usbc.json)

####  3.16.1	 Usb-c Pd Controller

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|pdVersion|version of power delivery spec implemented by controller|String|Yes|
|usbTypecRevision|usb type-c spec revision implemented by controller|String| |
|powerRoleSupported|roles supported by pd controller|String| |
|fastRoleSwapSupport|whether the pd controller supports fast role swap (FRS)|Boolean| |
|vconnPowerSupport|whether the pd controller has support for vconn power|Boolean| |
|vconnPowerLimit|power limit supported by internal vconn switch (if supported)|definitions.json#/unit| |
|vconnMaxCurrent|maximum continuous current supported by internal vconn switch (if supported)|definitions.json#/unit| |
|vconnOverCurrentLimit|over current limit supported by internal vconn switch (if supported)|definitions.json#/unit| |
|integratedVbusDischargeSwitch|whether the pd controller has one or more integrated vbus discharge switches |Boolean| |
|integratedLoadSwitch|whether the pd controller has one or more integrated load switches |Boolean| |
|maxSinkCurrent|maximum continuous current supported by pd controller integrated sink load switch|definitions.json#/unit| |
|maxSourceCurrent|maximum continuous current supported by pd controller integrated source load switch|definitions.json#/unit| |
|sinkfetOverCurrentLimit|over current limit supported by pd controller integrated sink load switch|definitions.json#/unit| |
|sourcefetOverCurrentLimit|over current limit supported by pd controller integrated source load switch|definitions.json#/unit| |
|onResistanceSinkFet|on-resistance of the integrated sink load switch|definitions.json#/unit| |
|onResistanceSourceFet|on-resistance of the integrated source load switch|definitions.json#/unit| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|componentProtectionThresholds|Thermal and power supply protection thresholds of a device|definitions.json#/componentProtectionThresholds| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

### 3.17	 Semiconductor

Source: [semiconductor.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/semiconductor.json)

####  3.17.1	 Mosfet

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn.|definitions.json#/componentID|Yes|
|mosfetType|type of MOSFET.|String|Yes|
|mosfetChannelType|doping of a transistor's channel - describes whether a transistor is n-type or p-type.|String| |
|transistorCount|number of transistors in the package.|Number| |
|vgs|gate to source voltage difference of a MOSFET.|definitions.json#/unit| |
|vgsMax|maximum gate to source voltage difference that can be continously applied to a MOSFET. This is a limiting value.|definitions.json#/unit| |
|vds|drain to source voltage difference of a MOSFET.|definitions.json#/unit| |
|vdsMax|maximum drain to source voltage difference that can be continously applied to a MOSFET. This is a limiting value.|definitions.json#/unit| |
|vdsVbr|drain to source breakdown voltage of a MOSFET.|definitions.json#/unit| |
|vgsThMax|maximum gate to source voltage difference required to produce a conducting path between drain and source.|definitions.json#/unit| |
|vgsThTyp|gate to source voltage difference (vgs) required to produce a conducting path between drain and source.|definitions.json#/unit| |
|vgsThMin|minimum gate to source voltage difference required to produce a conducting path between drain and source.|definitions.json#/unit| |
|vsdDiodeVf|reverse diode forward voltage when a MOSFET is in off-state.|definitions.json#/unit| |
|iD|Drain Current of a MOSFET.|definitions.json#/unit| |
|iDrain|maximum continous DC current that can flow through a MOSFET channel.This is a limiting value.|definitions.json#/unit| |
|idPulsed|maximum pulsed DC current that can flow through a MOSFET channel.This is a limiting value.|definitions.json#/unit| |
|iDss|drain-source leakage current of a MOSFET when the gate to source voltage difference is zero|definitions.json#/conditionalProperty| |
|iGss|gate-source leakage current of a MOSFET when the drain to source voltage difference is zero|definitions.json#/conditionalProperty| |
|diodeContinuousCurrent|maximum continuous forward current of the body diode of a MOSFET (IS).This is a limiting value.|definitions.json#/unit| |
|diodePulsedCurrent|maximum pulsed forward current of the body diode of a MOSFET. This is a limiting value.|definitions.json#/unit| |
|forwardTransconductance|signal gain, change in drain current with variation of gate-source voltage of a MOSFET (gFS).|definitions.json#/conditionalProperty| |
|rdson|on-state resistance of a MOSFET.|definitions.json#/conditionalProperty| |
|rg|internal gate resistance of a MOSFET.|definitions.json#/conditionalProperty| |
|ciss|input capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|coss|output capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|crss|reverse transfer capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|qg|total gate charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgd|gate to drain charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgs|gate to source charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qrr|reverse recovery charge of the body diode of a MOSFET.|definitions.json#/conditionalProperty| |
|idVsVds|graph of drain current (iD) vs drain source voltage (vds).|graph.json#/graphDefiniton| |
|idVsVgs|graph of drain current (iD) vs gate source voltage (vgs).|graph.json#/graphDefiniton| |
|tdON|turn-on delay time of a MOSFET|definitions.json#/conditionalProperty| |
|tdOFF|turn-off delay time of a MOSFET|definitions.json#/conditionalProperty| |
|riseTime|rise time of a MOSFET|definitions.json#/conditionalProperty| |
|fallTime|fall time of a MOSFET|definitions.json#/conditionalProperty| |
|trr|reverse recovery time of the body diode of a MOSFET.|definitions.json#/conditionalProperty| |
|pTot|maximum power dissipation of a MOSFET.|definitions.json#/conditionalProperty| |
|pdVsTemp|graph of power dissipation vs temperature.|graph.json#/graphDefiniton| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.17.2	 Diode

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|type of diode|String| |
|diodeCount|number of diodes in the package|Number| |
|diodeConfiguration|configuration of diode|String| |
|numberOfProtectedLines|number of lines a diode can protect|Number| |
|vf|forward voltage of a diode|definitions.json#/conditionalProperty| |
|if|continuous forward current of a diode|definitions.json#/unit| |
|ifm|maximum continuous forward current a diode can support|definitions.json#/unit| |
|ifrm|maximum repetitive peak forward current a diode can support|definitions.json#/unit| |
|ifsm|maximum non-repetitive surge forward current a diode can support|definitions.json#/unit| |
|vbr|breakdown voltage of a diode|definitions.json#/unit| |
|ir|reverse current|definitions.json#/conditionalProperty| |
|vz|breakdown voltage of a zener diode|definitions.json#/conditionalProperty| |
|vrm|maximum reverse standoff voltage a tvs diode can withstand|definitions.json#/unit| |
|vcl|clamping voltage of a tvs diode|definitions.json#/conditionalProperty| |
|vr|maximum continuous reverse biased voltage a diode can support|definitions.json#/unit| |
|vrrm|maximum repetitive reverse voltage pulses a diode can support|definitions.json#/unit| |
|cd|diode junction capacitance - between the anode and cathode- in reverse bias condition|definitions.json#/unit| |
|trr|reverse recovery time it takes the diode to stop conducting when its voltage changes from forward-bias to reverse-bias|definitions.json#/unit| |
|pTot|maximum power dissipation of a forward biased diode|definitions.json#/conditionalProperty| |
|ifVsVf|graph of forward current (If) vs forward voltage (VfTyp)|graph.json#/graphDefiniton| |
|package|package size of resistor|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.17.3	 Bjt

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn.|definitions.json#/componentID|Yes|
|bjtChannelType|doping of a transistor's channel - describes whether a transistor is n-type or p-type.|String| |
|transistorCount|number of transistors in the package.|Number| |
|collectorCurrent|maximum current flow of BJT as measured at the collector (Icc)|definitions.json#/unit| |
|peakCollectorCurrent|maximum pulsed current flow of BJT as measured at the collector (Icm)|definitions.json#/conditionalProperty| |
|baseCurrent|maximum current flow of BJT as measured at the base (Ib)|definitions.json#/unit| |
|peakBaseCurrent|maximum pulsed current flow of BJT as measured at the base (Ibm)|definitions.json#/conditionalProperty| |
|collectorBaseVoltage|maximum voltage between collector and base terminals of BJT with an open emitter terminal (V_CBO)|definitions.json#/unit| |
|collectorEmitterVoltage|maximum voltage between collector and emitter terminals of BJT with an open base terminal (V_CEO)|definitions.json#/unit| |
|emitterBaseVoltage|maximum voltage between emitter and base terminals of BJT with an open collector terminal (V_EBO)|definitions.json#/unit| |
|totalPowerDissipation|maximum power that can be continously dissipated under temperature conditions|definitions.json#/conditionalProperty| |
|collectorBaseCutOffCurrent|current into the collector terminal when the BJT's base and collector are reverse biased and the emitter is open (I_CBO)|definitions.json#/conditionalProperty| |
|emitterBaseCutOffCurrent|current into the base terminal when the BJT's base and emitter are reverse biased and the collector is open (I_EBO)|definitions.json#/conditionalProperty| |
|dcGain|ratio of collector current to base current (hfe)|definitions.json#/conditionalProperty| |
|collectorEmitterSaturationVoltage|collector-emitter voltage below which a change in base current does not impact collector current (VCE_SAT)|definitions.json#/conditionalProperty| |
|baseEmitterSaturationVoltage|base-emitter voltage required to ensure the collector is forward biased for certain current conditions (VBE_SAT)|definitions.json#/conditionalProperty| |
|collectorEmitterBreakdownVoltage|collector-emitter voltage at which a specified current flows with the base open|definitions.json#/conditionalProperty| |
|baseEmitterBreakdownVoltage|base-emitter voltage at which a specified current flows with the collector open|definitions.json#/conditionalProperty| |
|delayTime|time delay between input signal rising and when collector current rises to 10% of Isat (td)|definitions.json#/conditionalProperty| |
|riseTime|time for collector current to rise through active region from 10% to 90% of Isat (tr)|definitions.json#/conditionalProperty| |
|storageTime|time delay between input signal falling and when collector current falls to 90% of Isat (ts)|definitions.json#/conditionalProperty| |
|fallTime|time for collector current to fall from 90% to 10% of Isat (tf)|definitions.json#/conditionalProperty| |
|collectorCapacitance|parasitic capacitance of collector terminal under certain conditions (Cc)|definitions.json#/conditionalProperty| |
|emitterCapacitance|parasitic capacitance of emitter terminal under certain conditions (Ce)|definitions.json#/conditionalProperty| |
|transitionFrequency|frequency of unity current gain with a short circuit output (ft)|definitions.json#/conditionalProperty| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

### 3.18	 LED

Source: [led.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/led.json)

####  3.18.1	 LED

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn.|definitions.json#/componentID|Yes|
|ledColor|LED color|String| |
|vf|forward voltage of an LED |definitions.json#/conditionalProperty| |
|if|continuous forward current of an LED|definitions.json#/unit| |
|ifp|peak forward current of an LED|definitions.json#/unit| |
|vr|maximum continuous reverse biased voltage an LED can support|definitions.json#/unit| |
|ir|reverse current (leakage) of an LED|definitions.json#/conditionalProperty| |
|ledCapacitance|capacitance of an LED|definitions.json#/unit| |
|iv|LED luminous intensity|definitions.json#/conditionalProperty| |
|peakWavelength|light spectrum output value emitted by an LED at highest wavelength|definitions.json#/conditionalProperty| |
|dominantWavelength| dominant wavelength an LED emits the majority of the time|definitions.json#/conditionalProperty| |
|angleHalfIntensity| angle at which LED intensity falls to 50% of its maximum value|definitions.json#/unit| |
|pd|power dissipation of an LED|definitions.json#/unit| |
|package|package size of LED|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

### 3.19	 Adc_dac

Source: [adc_dac.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/adc_dac.json)

####  3.19.1	 Adc

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|digitalResolution|number of bits of resolution in the digital output|Number| |
|conversionTime|time required to convert from an analog signal to digital output|definitions.json#/unit| |
|sampleRate|maximum rate at which the ADC can convert samples|definitions.json#/unit| |
|offsetError|difference (in LSB) of the output at the zero point between an actual and ideal ADC|Number| |
|gainError|difference (in LSB) of how the actual transfer function matches the ideal transfer function, also called full scale error|Number| |
|integralNonlinearity|deviation of an actual transfer function from an ideal transfer function, in LSB|Number| |
|differentialNonlinearity|difference (in LSB) in step width between the actual and ideal transfer functions|Number| |
|rmsNoise|root mean square (RMS) noise of ADC|Number| |
|SNR|signal to noise (SNR) ratio of the converter|Number| |
|interface|digital communication interfaces supported|array of String| |
|inputType|whether the ADC has a single ended or differential input|String| |
|inputChannels|number of input channels|Number| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.19.2	 Dac

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|digitalResolution|number of bits of resolution|Number| |
|offsetError|analog response to an input code of all zeros|definitions.json#/unit| |
|gainError|difference (in percentage of FSR) of how well the slope of the actual transfer function matches the ideal transfer function|definitions.json#/unit| |
|integralNonlinearity|deviation of an actual transfer function from an ideal transfer function, in LSB|Number| |
|differentialNonlinearity|difference between the ideal and the actual output responses for successive DAC codes, in LSB|Number| |
|settlingTime|time from application of input code to valid output response|definitions.json#/unit| |
|sampleRate|maximum rate at which the DAC can convert samples|definitions.json#/unit| |
|interface|digital communication interfaces supported|array of String| |
|otuputType|whether the DAC has a single ended or differential output|String| |
|outputChannels|number of output channels|Number| |
|currentConsumption|current used by device in various power modes|definitions.json#/currentConsumption| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

### 3.20	 Clock_oscillators

Source: [hardware.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/hardware.json)

####  3.20.1	 Oscillator

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn.|definitions.json#/componentID|Yes|
|baseResonator|technology producing resonance|String| |
|frequency|output frequency of oscillator|definitions.json#/unit|Yes|
|frequencyStability|Frequency change over temperature, load, supply voltage change and aging|Must set either Ref or Type| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|outputLoad|maxium capacitive load that can be supported by oscillator|definitions.json#/unit| |
|riseTime|time for output to go from 10% to 90% of output max|definitions.json#/unit| |
|fallTime|time for output to go from 90% to 10% of output max|definitions.json#/unit| |
|startUpTime|time between enable and output reaching 10% of output max|definitions.json#/unit| |
|dutyCycle|time above 50% of output max over entire period|definitions.json#/unit| |
|phaseJitter|variation of waveform period|definitions.json#/conditionalProperty| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

### 3.21	 Sensor

Source: [sensor.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/sensor.json)

####  3.21.1	 Accelerometer

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|accelerometerType|type of accelerometer|String| |
|accelerationRanges|range of force that accelerometer can measure|Number| |
|accelerationSensitivity|smallest change in force the accelerometer is able to capture (typical) at a given acceleration range|definitions.json#/conditionalProperty| |
|accelerationSensitivityOverTemperature|accelerometer sensitivity change over temperature|definitions.json#/value| |
|axis|number of axes of acceleration measurement|Number| |
|zerogOffset|output of the accelerometer when no acceleration is applied|definitions.json#/value| |
|zerogOffsetOverTemperature|accelerometer zero-g offset change over temperature|definitions.json#/value| |
|outputType|measurement output type|String| |
|outputResolution|output resolution of acceleration measurement|definitions.json#/unit| |
|interface|interface(s) for communication to host|array of String| |
|bandwidth|bandwidth of acceleration measurement|definitions.json#/unit| |
|outputDataRate|output Data rate (ODR) of a device|definitions.json#/unit| |
|rmsNoise|broadband rms noise of a device|definitions.json#/value| |
|spectralNoiseDensity|spectral noise density of a device|definitions.json#/value| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|package size of a device|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.21.2	 Gyroscope

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|gyroRanges|range of angular speed that gyro can measure|Number| |
|gyroSensitivity|smallest change in angular speed gyro is able to capture (typical) at a given gyro range|definitions.json#/conditionalProperty| |
|gyroSensitivityOverTemperature|gyroscope sensitivity change over temperature|definitions.json#/value| |
|axis|number of axes of measurement|Number| |
|zeroRateOffset|output of the gyroscope when no angular velocity is applied|definitions.json#/value| |
|zeroRateOffsetOverTemperature|gyro zero rate offset change over temperature|definitions.json#/value| |
|interface|interface(s) for communication to host|array of String| |
|bandwidth|bandwidth of gyroscope|definitions.json#/unit| |
|outputDataRate|output Data rate (ODR) of a device|definitions.json#/unit| |
|outputType|measurement output type|String| |
|rmsNoise|broadband rms noise of a device|definitions.json#/value| |
|spectralNoiseDensity|spectral noise density of a device|definitions.json#/value| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|package size of a device|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.21.3	 Magnetic Sensor

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|outputPolarity|indicates whether the sensor output is active high or active low|String| |
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|magneticSensingTechnology|method by which magnetism is detected|String| |
|outputType|measurement output type|String| |
|quiescentOutput|output of the magnetic sensor when no magnet is present|definitions.json#/unit| |
|outputVoltageLinearRange|output voltage range over which the magnetic flux density response is linear|definitions.json#/unit| |
|linearMagneticSensingRange|magnetic flux density range over which the output voltage is linear |definitions.json#/value| |
|sensitivity|this is the gain - amount of change in the output voltage for a change in the magnetic flux density|definitions.json#/value| |
|operatePoint|magnetic flux density threshold which causes the sensor output to turn on|definitions.json#/value| |
|releasePoint|magnetic flux density threshold which causes the sensor output to turn off|definitions.json#/value| |
|hysteresis|delta between the operate point and the release point threshold|definitions.json#/value| |
|bandwidth|sensing bandwidth|definitions.json#/unit| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|package size of a device|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

####  3.21.4	 Thermal

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|sensingTechnology|method by which temperature is detected|String| |
|outputType|measurement output type|String| |
|interface|interface(s) for communication to host|array of String| |
|accuracy|accuracy of temperature sensor|definitions.json#/unit| |
|temperatureRange|range of temperature sensor|definitions.json#/unit| |
|resolution|maximum resolution of temperature sensor|Number| |
|gain|amount of change in the output voltage for a change in temperature|definitions.json#/value| |
|currentConsumption|current consumption of a device|definitions.json#/currentConsumption| |
|package|package size of a device|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

