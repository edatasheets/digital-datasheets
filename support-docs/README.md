
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

### 3.4	 Common

####  3.4.1	 Specification To Capture Information To Identify Components

Source: [componentID.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/componentID.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|manufacturer|company that manufactures the part|String|Yes|
|componentName|base part name that describes the form and fit of a component|String| |
|orderableMPN|orderable manufacturer part numbers, including packing and software information|array of String|Yes|
|sourceDatasheetID|methods for identifying the human-readable source information for a digital datasheet|#/$defs/sourceDatasheetID|Yes|
|digitalDatasheetID|methods for identifying the version of the digital datasheet|#/$defs/digitalDatasheetID|Yes|
|status|production status of a component|String| |
|complianceList|list of standards the part complies with|array of String| |

####  3.4.2	 DigitalDatasheetID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|publishedDate|date the digital datasheet was published|String| |
|guid|vendor defined guid (see https://www.guidgenerator.com/) to uniquely identify digital datasheet version|String| |

####  3.4.3	 SourceDatasheetID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|publishedDate|date the source datasheet was published|String| |
|version|version of the source datasheet|String| |
|datasheetURI|uri to the source datasheet pdf or html view|String| |
|productURI|uri to the source datasheet's product page'|String| |

####  3.4.3	 specification to capture protection thresholds data of a component

Source: [componentProtectionThresholds.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/componentProtectionThresholds.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|thermalShutdownThresholdRising|thermal Shutdown (tsd) Threshold with temperature rising|../common/unit.json#/unit| |
|thermalShutdownThresholdFalling|thermal Shutdown (tsd) Threshold with temperature falling|../common/unit.json#/unit| |
|thermalShutdownHysteresis|thermal Shutdown (tsd) Hysteresis|../common/unit.json#/unit| |
|powerSupplyProtection|undervoltage lockout, overvoltage protection thresholds of a supply|array of #/$defs/powerSupplyProtection| |

####  3.4.4	 PowerSupplyProtection

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|supplyName|name of the power supply|String| |
|overVoltageProtectionThresholdRising|Overvoltage Protection (OVP) Threshold with power supply rising|../common/unit.json#/unit| |
|overVoltageProtectionThresholdFalling|Overvoltage Protection (OVP) Threshold with power supply falling|../common/unit.json#/unit| |
|overVoltageProtectionHysteresis|Overvoltage Protection (OVP) Hysteresis|../common/unit.json#/unit| |
|underVoltageLockoutThresholdRising|Undervoltage Lockout (UVLO) Threshold with power supply rising|../common/unit.json#/unit| |
|underVoltageLockoutThresholdFalling|Undervoltage Lockout (UVLO) Threshold with power supply falling|../common/unit.json#/unit| |
|underVoltageLockoutHysteresis|Undervoltage Lockout (UVLO) Hysteresis|../common/unit.json#/unit| |

####  3.4.5	 Specification Of The Conditional Property Object

Source: [conditionalProperty.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/conditionalProperty.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|value of property|../common/unit.json#/unit| |
|conditions|conditions under which the property is measured|array of String| |

####  3.4.6	 specification to capture the current consumption data of a component

Source: [currentConsumption.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/currentConsumption.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|supplyName|name of the power supply |String| |
|quiescentCurrent|quiescent current (Iq) of a device|../common/unit.json#/unit| |
|shutdownCurrent|shutdown current (Isd) of a device|../common/unit.json#/unit| |
|activeCurrent|current consumption when a device is in active mode|../common/unit.json#/unit| |
|sleepCurrent|current consumption when a device is in sleep mode|../common/unit.json#/unit| |
|idleCurrent|current consumption when a device is in idle mode|../common/unit.json#/unit| |

####  3.4.7	 Specification For Referencing An External File

Source: [externalFile.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/externalFile.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|fileDescription|text description of the contents of an external file|String| |
|fileType|type of file being linked|String| |
|fileExtension|extension of file linked|String| |
|companionSoftware|optional, name of software program used to access file|String| |
|standardReferenced|optional, name of the standard the file is written in|String| |
|fileURI|URI linking to the CAD file|String| |

####  3.4.8	 specification of a graph

Source: [graph.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/graph.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|title|title of a graph|String| |
|xUnits|x-axis units|String| |
|xLabel|x-axis title|String| |
|yUnits|y-axis units|String| |
|yLabel|y-axis title|String| |
|numberOfCurves|total number of curves in graph|Number| |
|curve|data represented by one or more curves in a graph|array of #/$defs/curve| |

####  3.4.9	 Curve

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|label|description of the data captured by a curve|String| |
|xData|x-axis values of the curve|array of Number| |
|yData|y-axis values of the curve|array of Number| |

####  3.4.10	 specification of a package

Source: [package.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/package.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|length|length of a side of a package|../common/unit.json#/unit| |
|width|width of a side of a package|../common/unit.json#/unit| |
|height|height of a package|../common/unit.json#/unit| |
|nominalFootprints|references to external footprints in standard CAD formats|array of ../common/externalFile.json#/externalFile| |
|breakoutExamples|references to external board file that contains layout breakout example|array of ../common/externalFile.json#/externalFile| |
|partModelInformation|reference to an external XML file that contains a part model in IPC/DAC2552 format|../common/externalFile.json#/externalFile| |
|standardPackageSize|name of standard package size (imperial)|String| |
|standardPackageType|name of standard package types|String| |

####  3.4.11	 Specification Of Pinpaths Through An IC

Source: [pinPaths.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/pinPaths.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numberOfPinPaths|number of pinPaths defined. This number should not be higher than the number of components included in the part|Number| |
|partPinPaths|list of pins associated with each component in a multi-component part|array of #/$defs/partPinPaths|Yes|

####  3.4.12	 PartPinPaths

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentName|name of a component within a part |String| |
|componentPinNames|names of pins associated with each of the component in the part. Pin names must match the name used in part pin definition|array of String| |

####  3.4.13	 Specification Of Pins

Source: [pinSpec.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/pinSpec.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|terminalIdentifier|pin or ball number as defined by datasheet|String|Yes|
|name|name given to the signal appearing at the terminal of a component|String|Yes|
|standardizedName|standard name of pin|String| |
|description|description of the signal appearing at the terminal of an electric/electronic component|String| |
|numberOfSupportedFunctions|the total number of functions supported by this pin|Number| |
|functionProperties|list of function objects that can apply to an individual pin|array of #/$defs/functionProperties| |
|vihMin|the lowest value of high-level input voltage for which operation of the logic element within specification limits is to be expected|../common/unit.json#/unit| |
|vihMax|the highest value of high-level input voltage for which operation of the logic element within specification limits is to be expected|../common/unit.json#/unit| |
|vilMax|the highest value of low-level input voltage for which operation of the logic element within specification limits is to be expected|../common/unit.json#/unit| |
|vilMin|the lowest value of low-level input voltage for which operation of the logic element within specification limits is to be expected|../common/unit.json#/unit| |
|vol|the voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a low level at the output|../common/unit.json#/unit| |
|voh|the voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a high level at the output|../common/unit.json#/unit| |
|absVmax|maximum voltage rating beyond which damage to the device may occur|../common/unit.json#/unit| |
|absVmin|absolute minimum voltage that can be applied to a pin|../common/unit.json#/unit| |
|vmax|maximum continuous voltage that can safely be applied to a pin|../common/unit.json#/unit| |
|imax|maximum continuous current that can safely be drawn from a pin|../common/unit.json#/unit| |
|inputLeakage|maximum current draw out of a high impedance input pin|../common/unit.json#/unit| |
|outputLeakage|maximum current flow from a pin during the off state|../common/unit.json#/unit| |
|dcResistance|resistance of a pin of a connector|../common/unit.json#/unit| |
|voltageOptions|list of voltage levels supported by a pin|array of ../common/unit.json#/unit| |
|floatUnused|description of whether pin can safely be floated if it is not used|Boolean| |
|internalPullUp|indicates the value of an internal pull-up on a pin|../common/unit.json#/unit| |
|internalPullDown|indicates the value of an internal pull-down on a pin|../common/unit.json#/unit| |
|esd|indicates whether ESD protection exists on a pin|Boolean| |
|externalComponents|list of external component structures recommended to be attached to a pin|array of #/$defs/externalComponents| |

####  3.4.14	 ExternalComponents

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentType|type of external component required to be connected to a pin|String|Yes|
|configuration|electrical configuration of component connected to pin with respect to the pin|String|Yes|
|value|value of component|../common/unit.json#/unit| |
|connectionPin|name of pin to which an external component should be pulled up|String| |

####  3.4.15	 FunctionProperties

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|perFunctionName|name of the function of a pin|String| |
|interfaceType|type of interface enabled by pin|String| |
|pinUsage|standardized usage of pin|String| |
|direction|direction of a pin's function|String| |
|electricalConfiguration|electrical configuration of a pin|String| |
|polarity|whether the active state of a pin is high or low|String| |

####  3.4.15	 Specification Of Power Fets Properties

Source: [powerFetProperties.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/powerFetProperties.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|ilimHSFET|maximum sustained output current under which the high side FET will operate properly|../common/unit.json#/unit| |
|ilimLSFET|maximum sustained output current under which the low side FET will operate properly|../common/unit.json#/unit| |
|rdsonHSFET|high side FET on-resistance|../common/unit.json#/unit| |
|rdsonLSFET|low side FET on-resistance|../common/unit.json#/unit| |

####  3.4.16	 specifications of a ratio

Source: [ratio.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/ratio.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numerator|numerator of ratio|Number| |
|denominator|denominator of ratio|Number| |

####  3.4.17	 specification of a register

Source: [register.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/register.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|registerName|name of a register|String|Yes|
|registerLongName|full Name of a register|String| |
|registerAddressOffset|address of a register|String|Yes|
|registerSize|size of a register|../common/unit.json#/unit|Yes|
|registerType|type of a register|String| |
|registerResetValue|reset value of a register|String| |
|registerValue|value of a register|String| |
|registerAccessType|access type of a Register|String| |
|registerBitField|describes the bit fields in the register|#/$defs/registerBitField| |

####  3.4.18	 RegisterBitField

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|bitFieldName|Name of the bit field|String| |
|bitFieldLongName|Long Name of the bit field|String| |
|bitFieldDescription|Describes the bit field|String| |
|bitFieldNumber|Number of a bitfield|Number| |
|bitFieldRange|Range of the bit field|String| |
|bitFieldResetValue|Reset value of a bit field|String| |
|bitFieldAccessType|Access type of a bit field|String| |

####  3.4.19	 Specification Of Unit

Source: [unit.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/unit.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|siUnit|name of SI unit of measure|String| |
|unitName|name of unit if not defined in the siUnit list|String| |
|typValue|typical unit quantity corresponding to unit text - example 40mV would have a value of 40|Number| |
|minValue|minimum unit quantity corresponding to unit text - example: 30mV would have a value of 30|Number| |
|maxValue|maximum unit quantity corresponding to unit text - example: 50mV would have a value of 50|Number| |
|unitFactor|multiplier on the value to achieve the SI unit listed - for example if millivolt was selected, 40mV would have a unitFactor value of 1; if volt was selected, the unitFatctor value would be 0.001|Number| |
|relativeValueReference|if unit quantity is based on another reference, value of the reference|String| |
|relativeValueModifier|if a unit quantity is based on another reference, the value that edits that reference|Number| |
|relativeValueOperator|if a unit quantity is based on another reference, the operation that is performed with the modifier|String| |

### 3.5	 Clock

####  3.5.1	 Specification Of Clock

Source: [clock.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/clock/clock.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|fixedFrequency|clock frequency value if the clock has a fixed frequency|../common/unit.json#/unit| |
|numberClockOutputs|number of clock outputs in a clock IC|Number| |
|differentialSingleEnded|property describing whether a clock output is single ended or differential|String| |
|jitter|cycle to cycle clock jitter|../common/unit.json#/unit| |
|frequencyTolerance|amount of frequency variation from nominal frequency|../common/unit.json#/unit| |
|powerSupplyRejectionRatio|power supply rejection ratio (PSRR) or ratio between power supply variation and output variation|Number| |
|outputFormat|signal format of clock output|String| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.5.2	 Specification Of Oscillator

Source: [oscillator.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/clock/oscillator.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|baseResonator|technology producing resonance|String| |
|frequency|output frequency of oscillator|../common/unit.json#/unit|Yes|
|frequencyStability|frequency change over temperature, load, supply voltage change and aging|Number| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|outputLoad|maxium capacitive load that can be supported by oscillator|../common/unit.json#/unit| |
|riseTime|time for output to go from 10% to 90% of output max|../common/unit.json#/unit| |
|fallTime|time for output to go from 90% to 10% of output max|../common/unit.json#/unit| |
|startUpTime|time between enable and output reaching 10% of output max|../common/unit.json#/unit| |
|dutyCycle|time above 50% of output max over entire period|../common/unit.json#/unit| |
|phaseJitter|variation of waveform period|../common/conditionalProperty.json#/conditionalProperty| |

### 3.7	 Data_converter

####  3.7.1	 Specification Of Adc

Source: [adc.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/data_converter/adc.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|digitalResolution|number of bits of resolution in the digital output|Number| |
|conversionTime|time required to convert from an analog signal to digital output|../common/unit.json#/unit| |
|sampleRate|maximum rate at which the ADC can convert samples|../common/unit.json#/unit| |
|offsetError|difference (in LSB) of the output at the zero point between an actual and ideal ADC|Number| |
|gainError|difference (in LSB) of how the actual transfer function matches the ideal transfer function, also called full scale error|Number| |
|integralNonlinearity|deviation (in LSB) of an actual transfer function from an ideal transfer function|Number| |
|differentialNonlinearity|difference (in LSB) in step width between the actual and ideal transfer functions|Number| |
|rmsNoise|root mean square (RMS) noise of ADC|Number| |
|SNR|signal to noise (SNR) ratio of the converter|Number| |
|interface|digital communication interfaces supported|array of String| |
|inputType|whether the ADC has a single ended or differential input|String| |
|inputChannels|number of input channels|Number| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.7.2	 Specification Of Dac

Source: [dac.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/data_converter/dac.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|digitalResolution|number of bits of resolution|Number| |
|offsetError|analog response to an input code of all zeros|../common/unit.json#/unit| |
|gainError|difference (in percentage of FSR) of how well the slope of the actual transfer function matches the ideal transfer function|../common/unit.json#/unit| |
|integralNonlinearity|deviation of an actual transfer function from an ideal transfer function, in LSB|Number| |
|differentialNonlinearity|difference between the ideal and the actual output responses for successive DAC codes, in LSB|Number| |
|settlingTime|time from application of input code to valid output response|../common/unit.json#/unit| |
|sampleRate|maximum rate at which the DAC can convert samples|../common/unit.json#/unit| |
|interface|digital communication interfaces supported|array of String| |
|otuputType|whether the DAC has a single ended or differential output|String| |
|outputChannelCount|number of output channels|Number| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

### 3.8	 Hardware

####  3.8.1	 Specification Of Connector

Source: [connector.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/hardware/connector.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|function|intended function of a connector|String| |
|contactCount|number of contacts in a connector|Number|Yes|
|type|property describing the method of mating to the connector|String| |
|cycleRating|number of plug/unplug cycles a connector is rated to support|Number| |
|pitch|distance from the center of one contact on the connector to the center of the next contact|../common/unit.json#/unit| |
|keying|property describing whether a connector has an asymmetry to prevent it from being plugged in the wrong direction|Boolean| |

####  3.8.2	 Specification Of Switch

Source: [switch.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/hardware/switch.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|type|property describing the way in which the switch is activated|String| |
|contactType|property describing the order in which switch contact is made and broken|String| |
|circuitConfig|property describing the number of poles and throws in a switch|String| |
|cycleRating|number of on/off cycles a mechanical switch can reliably sustain|Number| |
|voltageRating|maximum DC voltage potential that can be applied across an open switch|../common/unit.json#/unit| |
|currentRating|maximum DC current that can flow through a closed switch without causing excessive heating|../common/unit.json#/unit| |
|onResistance|nominal resistance of a closed switch|../comon/unit.json#/unit|Yes|
|dielectricRating|maximum AC voltage potential that can be applied across an open switch for one minute|../common/unit.json#/unit| |

### 3.9	 Ic_io

####  3.9.1	 Specification Of Bridge Chip

Source: [bridge_chip.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/bridge_chip.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|muxRatio|ratio of inputs to outputs|../common/ratio.json#/ratio| |
|inputInterfaces|list of interfaces at the input of the bridge|array of String| |
|outputInterfaces|list of interfaces at the output of the bridge|array of String| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.9.2	 Specification Of High Speed Mux

Source: [highspeed_mux.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/highspeed_mux.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|muxRatio|ratio of inputs to outputs|../common/ratio.json#/ratio| |
|inputInterfaces|list of interfaces high speed mux is designed for|array of String| |
|maxDataRate|maximum data rate supported by high speed mux|../common/unit.json#/unit| |
|insertionLoss|insertion loss through high speed mux|Number| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.9.3	 Specification Of Level Shifter

Source: [level_shifter.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/level_shifter.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|inputVoltage|nominal input voltage of level shifter|../common/unit.json#/unit| |
|outputVoltage|nominal output voltage of level shifter|../common/unit.json#/unit| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.9.4	 Specification Of Redriver

Source: [redriver.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/redriver.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numberChannels|number of lanes (single ended or differential) supported by redriver|Number| |
|interface|list of interface(s) supported by redriver|array of String| |
|maxDataRate|maximum data rate supported by redriver|../common/unit.json#/unit| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.9.5	 Specification Of Retimer

Source: [retimer.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/retimer.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numberOfLanes|number of lanes (single ended or differential) supported by a device|Number| |
|interface|list of interface(s) supported by a device|array of String| |
|dpMaxLinkRate|maximum interface speed supported by DP interface of a device|String| |
|maxDataRate|maximum data rate supported by a device|../common/unit.json#/unit| |
|integratedAuxLsSwitch|whether the AUX/LSx switch for SBU is integrated|Boolean| |
|commonClock|whether a device supports common reference clock|Boolean| |
|sris|whether a device supports Seperate Reference clock with Independent Spread spectrum clocking(SRIS)|Boolean| |
|srns|whether a device supports Seperate Reference clock with No Spread spectrum clocking (SRNS)|Boolean| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  3.9.6	 Specification Of USB Battery Charging 1.2 (bc12) Detector

Source: [usb_bc12.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/usb_bc12.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|hostMode|whether host mode is supported by bc12 chip|Boolean|Yes|
|deviceMode|whether device mode is supported by bc12 chip|Boolean|Yes|

####  3.9.7	 Specification Of Usb-c Pd Controller

Source: [usbc_pdcontroller.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/usbc_pdcontroller.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|pdVersion|version of power delivery spec implemented by controller|String|Yes|
|usbTypecRevision|usb type-c spec revision implemented by controller|String| |
|powerRoleSupported|roles supported by pd controller|String| |
|fastRoleSwapSupport|whether the pd controller supports fast role swap (FRS)|Boolean| |
|vconnPowerSupport|whether the pd controller has support for vconn power|Boolean| |
|vconnPowerLimit|power limit supported by internal vconn switch (if supported)|../common/unit.json#/unit| |
|vconnMaxCurrent|maximum continuous current supported by internal vconn switch (if supported)|../common/unit.json#/unit| |
|vconnOverCurrentLimit|over current limit supported by internal vconn switch (if supported)|../common/unit.json#/unit| |
|integratedVbusDischargeSwitch|whether the pd controller has one or more integrated vbus discharge switches |Boolean| |
|integratedLoadSwitch|whether the pd controller has an integrated load switch for vbus power |Boolean| |
|maxSinkCurrent|maximum continuous current supported by pd controller integrated sink load switch|../common/unit.json#/unit| |
|maxSourceCurrent|maximum continuous current supported by pd controller integrated source load switch|../common/unit.json#/unit| |
|sinkfetOverCurrentLimit|over current limit supported by pd controller integrated sink load switch|../common/unit.json#/unit| |
|sourcefetOverCurrentLimit|over current limit supported by pd controller integrated source load switch|../common/unit.json#/unit| |
|onResistanceSinkFet|on-resistance of the integrated sink load switch|../common/unit.json#/unit| |
|onResistanceSourceFet|on-resistance of the integrated source load switch|../common/unit.json#/unit| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  3.9.8	 Specification Of Usb-c Port Controller

Source: [usbc_portcontroller.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/usbc_portcontroller.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|pdVersion|version of power delivery spec implemented by controller|String|Yes|
|usbTypecRevision|usb type-c spec revision implemented by controller|String| |
|usbTcpiRevision|usb type-c port controller interface spec revision implemented by controller|String| |
|powerRoleSupported|roles supported by pd controller|String| |
|fastRoleSwapSupport|whether the port controller supports fast role swap (FRS)|Boolean| |
|vconnPowerSupport|whether the port controller has support for vconn power|Boolean| |
|vconnFetOnResistance|on-resistance of the integrated Vconn FET|..common/unit.json#/unit| |
|vconnPowerLimit|power limit supported by internal vconn switch (if supported)|../common/unit.json#/unit| |
|vconnMaxCurrent|maximum continuous current supported by internal vconn switch (if supported)|../common/unit.json#/unit| |
|vconnOverCurrentLimit|over current limit supported by internal vconn switch (if supported)|../common/unit.json#/unit| |
|integratedVbusDischargeSwitch|whether the port controller has one or more integrated vbus discharge switches |Boolean| |
|integratedMux|whether the port controller has one or more integrated high speed muxes|Boolean| |
|interface|describes the communication interface from the controller to the policy manager|String| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

### 3.10	 Ic_microcontroller

####  3.10.1	 Specification Of Microcontroller/ec

Source: [microcontroller.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_microcontroller/microcontroller.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|onChipFlash|capacity of built-in flash in a microprocessor|../common/unit.json#/unit| |
|onChipRAM|capacity of built-in RAM in a microprocessor|../common/unit.json#/unit| |
|onChipROM|capacity of built-in ROM in a microprocessor|../common/unit.json#/unit| |
|coreProcessor|description of core processor|String| |
|coreArchitectureBits|number of bits of data a CPU can transfer per clock cycle|String| |
|clockSpeed|speed of main CPU clock|../common/unit.json#/unit| |
|firmwareVersion|firmware version of the part|String| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

### 3.11	 Ic_misc

####  3.11.1	 Specification Of Audio Codec

Source: [audio_codec.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/audio_codec.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|dataLength|number of bits in a data word|Number| |
|hpOutputSNR|headphone amplifier output SNR|Number| |
|hpOutputTHD+N|headphone output total harmonic distortion plus noise|../common/unit.json#/unit| |
|micInputSNR|microphone input SNR|number| |
|micInputTHD+N|microphone input total harmonic distortion plus noise|../common/unit.json#/unit| |
|jackDetect|describes whether headphone jack detection is supported|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|..common/currentConsumption.json#/currentConsumption| |

####  3.11.2	 Specification Of Speaker Amplifier

Source: [speaker_amplifier.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/speaker_amplifier.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|dataLength|number of bits in a data word|Number| |
|outputPower|typical output power from speaker amplifier|../common/conditionalProperty.json#/conditionalProperty| |
|efficiency|typical speaker amplifier efficiency|../common/conditionalProperty.json#/conditionalProperty| |
|thd+n|typical total harmonic distortion plus noise of amplifier|../common/conditionalProperty.json#/conditionalProperty| |
|sampleRate|sample rate of data out of amplifier|..common/unit.json#/unit| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|..common/currentConsumption.json#/currentConsumption| |

####  3.11.3	 Specification Of Tpm

Source: [tpm.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/tpm.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|interface|describes the communication interface from the chip to the host|array of String| |
|currentConsumption|current used by device in various power modes|..common/currentConsumption.json#/currentConsumption| |

####  3.11.4	 Specification Of WLAN Module

Source: [wlan_module.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/wlan_module.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|wlanSpec|version of wlan specification supported by module|String| |
|bluetoothVersion|version of bluetooth supported by module|String| |
|txrxChains|number of tx and rx chains in a wifi module|String| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|lteCoexFilter|describes whether module supports lte coexistance filtering|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|..common/currentConsumption.json#/currentConsumption| |

####  3.11.5	 Specification Of WWAN Module

Source: [wwan_module.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/wwan_module.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|networkSupport|networks supported by wwan module|String| |
|gpsSupport|whether wwan module has gps support|Boolean| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|..common/currentConsumption.json#/currentConsumption| |

### 3.12	 Logic

####  3.12.1	 Specification Of Logic Gate

Source: [logic_gate.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/logic/logic_gate.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|type|logical operation performed by logic gate|String|Yes|
|numberGates|number of logical gates encapsulated in logic IC|Number| |
|schmittTrigger|property describing whether logic gate has schmitt trigger inputs|Boolean| |
|propagationDelay|time between input changing to output changing|../common/unit.json#/unit| |
|rampTime|time for output to go from 10% nominal output voltage to 90% nominal output voltage|units.json#/unit| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

### 3.13	 Memory

####  3.13.1	 Specification Of Dram

Source: [dram.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/dram.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|type|type of dram|String| |
|capacity|capacity of dram chip|../common/unit.json#/unit|Yes|
|ranksperModule|numbers of ranks on dram module|Number| |
|diesPerChip|number of dies on dram chip|Number| |
|channelsPerDie|number of channels per die on dram chip|Number| |
|banksPerChannel|number of banks per channel of dram|Number| |
|bitsPerChannel|channel density of dram|../common/unit.json#/unit| |
|bitsPerDie|total die density of dram|../common/unit.json#/unit| |
|pageSize|page size of dram|../common/unit.json#/unit| |
|rows|number of rows per channel of dram|Number| |
|colunms|number of columns per row of dram|Number| |
|dataRate|dram maximum data rate|Number| |
|speed|dram maximum speed|../common/unit.json#/unit| |
|latencyCas|cl/tCAS, delay between read command issued and first output data available for read|Number| |
|delayRasCas|tRCD,delay between activation of row and activation of column where data is stored in the dram|../common/unit.json#/unit| |
|delayRasPrecharge|tRP, delay between closing access to a row through the precharge command and activating a new row to access data |../common/unit.json#/unit| |
|delayActivePrecharge|tRAS, delay between row active command issued and precharge command issued |../common/unit.json#/unit| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

####  3.13.2	 Specification Of Eeprom

Source: [eeprom.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/eeprom.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|accessTime|time to access the eeprom|../common/unit.json#/unit| |
|bitsPerWords|number of columns in the eeprom|Number| |
|bootBlockSize|size of the eeprom boot block|../common/unit.json#/unit| |
|capacity|capacity/density of eeprom|../common/unit.json#/unit|Yes|
|clockFrequency|eeprom clock frequency|../common/unit.json#/unit| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|dataRetention|maximum number of read and write cycle the part can support|Number| |
|endurance|time in years a bit in the eeprom can retain its data state |Number| |
|interface|interface of eeprom to host|String| |
|numberOfWords|number of rows in the eeprom|Number| |

####  3.13.3	 Specification Of Flash Memory

Source: [flash_memory.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/flash_memory.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|capacity|capacity/density of flash memory|../common/unit.json#/unit|Yes|
|pageSize|page size of flash memory|../common/unit.json#/unit| |
|blockSize|block size of flash memory|../common/unit.json#/unit| |
|bootBlockSize|size of the flash memory boot block|../common/unit.json#/unit| |
|interface|interface of flash memory to host|String| |
|clockFrequency|flash memory clock frequency|../common/unit.json#/unit| |
|blockEraseTime|time it takes to erase a block (largest erasable unit) of the flash memory|../common/unit.json#/unit| |
|sectorEraseTime|time it takes to erase a sector (smallest erasable unit) of the flash memory|../common/unit.json#/unit| |
|chipEraseTime|time it takes to erase the flash memory|../common/unit.json#/unit| |
|pageProgramTime|time it takes to program a page of the flash memory|../common/unit.json#/unit| |
|endurance|time in years a bit in the eeprom can retain its data state |Number| |
|dataRetention|maximum number of read and write cycle the part can support|Number| |
|hwReset|whether the part supports a hardware reset pin|Boolean| |
|writeProtect|whether the part has a write protect pin|Boolean| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

####  3.13.4	 Specification Of Rom

Source: [rom.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/rom.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|capacity|capacity of rom|../common/unit.json#/unit| |
|interface|interface of rom to host|String| |
|qeStatus|indicates whether the Quad Enable(QE) bit is set|Boolean| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

### 3.14	 Passives

####  3.14.1	 Specification Of Capacitor

Source: [capacitor.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/capacitor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|capacitor value|../common/unit.json#/unit|Yes|
|tolerance|nominal tolerance of a capacitor|../common/unit.json#/unit| |
|ratedVoltage|maximum voltage which may be applied continuously to a capacitance|../common/unit.json#/unit| |
|dielectric|dielectric material used in the capacitor|String| |
|polarized|describes whether the capacitor is polarized|Boolean| |
|equivalentSeriesResistance|equivalent series resistance (ESR) of the capacitor|../common/unit.json#/unit| |
|temperatureCoefficient|change in capacitance when the temperature is changed|Number| |
|minTemperature|minimum temperature under which a capacitor can be expected to reliably operate|../common/unit.json#/unit| |
|maxTemperature|maximum temperature under which a capacitor can be expected to reliably operate|../common/unit.json#/unit| |
|capacitorDerating|graph object to capture capacitance changes with voltage changes|../common/graph.json#/graph| |

####  3.14.2	 Specification Of Common Mode Choke

Source: [common_mode_choke.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/common_mode_choke.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|diffModeCutoff|frequency at which the differential mode attenuation equals -3dB|../common/unit.json#/unit| |
|commonModeAttenuation|graph object to capture common mode attenuation of a common mode choke at various frequencies|../common/graph.json#/graph| |
|dcResistance|dc resistance (DCR) of a common mode choke|../common/unit.json#/unit| |
|rmsCurrent|applied DC current (IRMS) that produces a common mode choke temperature rise of 40 deg C|../common/unit.json#/unit| |
|intendedApplication|intended application of a particular common mode choke|String| |

####  3.14.3	 Specification Of Ferrite Bead

Source: [ferrite_bead.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/ferrite_bead.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|dcResistance|dc resistance (DCR) of ferrite bead|../common/unit.json#/unit| |
|rmsCurrent|applied DC current (IRMS) that produces a ferrite bead temperature rise of 40 deg C|../common/unit.json#/unit| |
|impedance100MHz|impedance of ferrite bead under standard testing conditions at 100MHz|../common/unit.json#/unit| |
|impedanceTolerance|variation of ferrite bead impedance expressed as +/- percentage|../common/unit.json#/unit| |

####  3.14.4	 Specification Of Inductor

Source: [inductor.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/inductor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|inductor value|../common/unit.json#/unit|Yes|
|tolerance|nominal tolerance of an inductor|../common/unit.json#/unit| |
|ratedCurrent|maximum continuous current the inductor can handle|../common/unit.json#/unit| |
|saturationCurrent|current where the inductor enters the magnetic state, and the inductance drops a specified amount|../common/unit.json#/unit| |
|rmsCurrent|DC current that produces an inductor temperature rise of 40 degrees Celsius|../common/unit.json#/unit| |
|selfResonantFrequency|frequency at which the inductor starts behaving like a capacitor|../common/unit.json#/unit| |
|dcResistance|DC resistance of the inductor|../common/unit.json#/unit| |
|temperatureCoefficient|change in inductance when the temperature is changed|Number| |
|minTemperature|minimum temperature under which a inductor can be expected to reliably operate|../common/unit.json#/unit| |
|maxTemperature|maximum temperature under which a inductor can be expected to reliably operate|../common/unit.json#/unit| |
|saturationCurve|graph object to capture inductance as a function of current|../common/graph.json#/graph| |
|resonantFrequencyCurve|graph object to capture inductance as a function of frequency|../common/graph.json#/graph| |

####  3.14.5	 Specification Of Resistor

Source: [resistor.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/resistor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|resistance value|../common/unit.json#/unit|Yes|
|tolerance|nominal tolerance of a resistor|../common/unit.json#/unit| |
|powerRating|measure of power a resistor can dissipate indefinitely without degrading performance|../common/unit.json#/unit| |
|temperatureCoefficient|change in resistance when the temperature is changed|Number| |
|maxOverloadVoltage|maximum voltage that can be applied to the resistor for a short period of time|../common/unit.json#/unit| |
|maxLimitingElementVoltage|maximum voltage value that can be applied continuously to the resistor|../common/unit.json#/unit| |
|minTemperature|minimum temperature under which a resistor can be expected to reliably operate|../common/unit.json#/unit| |
|maxTemperature|maximum temperature under which a resistor can be expected to reliably operate|../common/unit.json#/unit| |
|resistorDerating|graph object to capture resistance changes with temperature|../common/graph.json#/graph| |

### 3.15	 Power

####  3.15.1	 Specification Of Battery Charger

Source: [battery_charger.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/battery_charger.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|chargerType|battery charger type|String| |
|converterType|switching charger type|String| |
|chargerTopology|type of battery charger topology (Narrow VDC vs Hybrid Power Boost)|String| |
|batteryConfig|battery configuration supported by the device|array of String| |
|batteryCellChemistry|battery cell chemistry supported by the device|array of String| |
|inputPowerSource|input power source supported by the device|array of String| |
|inputCurrentAccuracy|accuracy of input current when set|../common/unit.json#/unit| |
|batteryChargeCurrent|charging current of a device|../common/unit.json#/unit|Yes|
|batteryChargeCurrentAccuracy|charging current regulation accuracy of a device|../common/unit.json#/unit| |
|batteryPreChargeCurrent|charging current of a device in pre-charge phase|../common/unit.json#/unit| |
|batteryPreChargeCurrentAccuracy|pre-charging current regulation accuracy of a device|../common/unit.json#/unit| |
|batteryTrickleChargeCurrent|charging current of a device in trickle charge phase|../common/unit.json#/unit| |
|batteryTrickleChargeCurrentAccuracy|trickle charging current regulation accuracy of a device|../common/unit.json#/unit| |
|batteryTerminationChargeCurrent|charging current of a device in charge termination phase|../common/unit.json#/unit| |
|batteryTerminationChargeCurrentAccuracy|termination charging current regulation accuracy of a device|../common/unit.json#/unit| |
|batteryChargeVoltage|battery charge voltage regulated by a device|../common/unit.json#/unit| |
|batteryChargeVoltageAccuracy|accuracy of battery charge voltage regulated by a device|../common/unit.json#/unit| |
|efficiency|charger efficiency as a function of charge current of a device|../common/graph.json#/graph| |
|vin|input voltage under which a device can be expected to reliabily operate|../common/unit.json#/unit| |
|fsw|switching frequency of a device|../common/unit.json#/unit| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |
|batteryChargerProtections|battery charger specific protections supported by device|array of String| |
|integratedLoadSwitch|whether the device contains integrated power path load switch(es)|Boolean| |
|integratedFets|whether the device contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|../common/powerFetProperties.json#/powerFetProperties| |
|gateCapacitance|describes gate capacitance supported on external fets|../common/unit.json#/unit| |
|inputSenseResistor|describes intput sense resistor value|../common/unit.json#/unit| |
|batterySenseResistor|describes battery sense resistor value|../common/unit.json#/unit| |
|passThroughMode|whether pass through mode is supported|Boolean| |
|bc12Support|whether bc 1.2 detection is built in|Boolean| |
|tcpcSupport|whether type-C port controller support is built in|Boolean| |
|usbTypecRevision|usb type-c spec revision supported by a device|String| |
|pdVersion|version of power delivery spec supported by a device|String| |

####  3.15.2	 Specification Of Display Backlight Driver

Source: [displaybacklight_driver.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/displaybacklight_driver.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|vin|input voltage under which a device can be expected to operate properly|../common/unit.json#/unit|Yes|
|vout|output voltage range a device can regulate|../common/unit.json#/unit| |
|ioutPerString|output current per string a device can regulate|../common/unit.json#/unit| |
|ioutAccuracy|accuracy of per string current regulated by a device|../common/unit.json#/unit| |
|fsw|switching frequency of a device|../common/unit.json#/unit| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|integratedFets|whether a device contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|../common/powerFetProperties.json#/powerFetProperties| |
|currentMatchingAccuracy|current matching between LED strings|../common/unit.json#/unit| |
|dimmingSupport|whether a device supports output current dimming|Boolean| |
|dimmingControl|whether a device is dimmed by PWM or I2C|String| |
|dimmingFrequency|dimming frequency of a device|../common/unit.json#/unit| |
|dimmingRatio|dimming ratio of a device|../common/ratio.json#/ratio| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |
|efficiency|backlight driver efficiency as a function of forward current|../common/graph.json#/graph| |

####  3.15.3	 Specification Of Linear Regulator

Source: [linear_regulator.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/linear_regulator.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|vin|minimum input voltage under which the part can be expected to operate without the output dropping|../common/unit.json#/unit|Yes|
|vout|output voltage the part can regulate|../common/unit.json#/unit|Yes|
|feedbackVoltage|voltage comparison point at the feedback node (vref)|../common/unit.json#/unit| |
|dropoutVoltage|dropout voltage of a device|../common/unit.json#/unit| |
|loadCurrent|load current supported by a device|../common/unit.json#/unit|Yes|
|currentLimit|sustained output current threshold beyond which the output of a device starts drooping|../common/unit.json#/unit| |
|voutAccuracy|output voltage variation at no load of a device|../common/unit.json#/unit| |
|loadRegulation|output voltage variation,from no load to full load, of a device |../common/unit.json#/unit| |
|lineRegulation|output voltage variation,from minimum input voltage to maximum input voltage, of a device |../common/unit.json#/unit| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|../common/unit.json#/unit| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|../common/unit.json#/unit| |
|powerSupplyRejectionRatio|graph object to capture Power Supply Rejection Ratio (PSRR) of device as a function of frequency|../common/graph.json#/graph| |
|rmsOutputNoise|graph object to capture RMS output noise of device as a function of frequency|../common/graph.json#/graph| |
|totalOutputNoise|total output noise of a device|../common/unit.json#/unit| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  3.15.4	 Specification Of Load Switch

Source: [load_switch.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/load_switch.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|fetType|type of pass FET in a device|String| |
|loadSwitchCount|number of load switches in the package|Number| |
|vin|input voltage under which a device can be expected to reliabily operate|../common/unit.json#/unit|Yes|
|outputCurrent|continuous DC cuurent supported by a device|../common/unit.json#/unit|Yes|
|onResistance|on state resistance of FET|../common/conditionalProperty.json#/conditionalProperty|Yes|
|pullDownResistance|pull-down resistance of a device from the output to the ground|../common/unit.json#/unit| |
|currentLimitSupport|whether a device supports current limiting|Boolean| |
|adjustableRiseTimeSupport|whether a device supports adjustable rise time|Boolean| |
|quickOutputDischargeSupport|whether a device supports quick output discharge|Boolean| |
|reverseCurrentBlockingSupport|whether a device supports reverse current blocking|Boolean| |
|powerGoodSupport|whether a device supports power good|Boolean| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|../common/conditionalProperty.json#/conditionalProperty| |
|offTime|time between enable deasserted and output voltage falling to 90% nominal|../common/conditionalProperty.json#/conditionalProperty| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|../common/conditionalProperty.json#/conditionalProperty| |
|fallTime|time for output voltage to go from 90% vout nominal to 10% vout nominal|../common/conditionalProperty.json#/conditionalProperty| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  3.15.5	 Specification Of Pmic

Source: [pmic.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/pmic.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|linearRegulatorCount|number of linear regulators in the device|Number| |
|buckRegulatorCount|number of buck regulators in the device|Number| |
|boostRegulatorCount|number of boost regulators in the device|Number| |
|buckBoostRegulatorCount|number of buck-boost regulators in the device|Number| |
|adcCount|number of analog to digital converters in the device|Number| |
|externalClockCount|number of external clocks the device requires|Number| |
|internalClockCount|number of clocks/oscillators in the device|Number| |
|loadSwitchCount|number of load switches in the device|Number| |
|usbSwitchCount|number of USB switches in the device|Number| |
|componentList|list of components integrated in the device|array of String| |
|instances|definition (specification) of each instance of a component in the device|array of #/$defs/powerComponentDefinitions| |
|vin|input voltage under which a device can be expected to reliabily operate|../common/unit.json#/unit| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  3.15.6	 PowerComponentDefinitions

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentName|name of the component defined in the digital datasheets specifications|String| |
|instanceName|name of component instance|String| |
### 3.16	 Semiconductor

####  3.16.1	 Specification Of Bjt

Source: [bjt.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/bjt.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|bjtChannelType|doping of a transistor's channel - describes whether a transistor is n-type or p-type|String|Yes|
|transistorCount|number of transistors in the package|Number| |
|collectorCurrent|maximum current flow of BJT as measured at the collector (Icc)|../common/unit.json#/unit| |
|peakCollectorCurrent|maximum pulsed current flow of BJT as measured at the collector (Icm)|../common/unit.json#/unit| |
|baseCurrent|maximum current flow of BJT as measured at the base (Ib)|../common/unit.json#/unit| |
|peakBaseCurrent|maximum pulsed current flow of BJT as measured at the base (Ibm)|../common/unit.json#/unit| |
|collectorBaseVoltage|maximum voltage between collector and base terminals of BJT with an open emitter terminal (V_CBO)|../common/unit.json#/unit| |
|collectorEmitterVoltage|maximum voltage between collector and emitter terminals of BJT with an open base terminal (V_CEO)|../common/unit.json#/unit| |
|emitterBaseVoltage|maximum voltage between emitter and base terminals of BJT with an open collector terminal (V_EBO)|../common/unit.json#/unit| |
|pTot|maximum power that can be continously dissipated under temperature conditions|../common/conditionalProperty.json#/conditionalProperty| |
|collectorBaseCutOffCurrent|current into the collector terminal when the BJT's base and collector are reverse biased and the emitter is open (I_CBO)|../common/conditionalProperty.json#/conditionalProperty| |
|emitterBaseCutOffCurrent|current into the base terminal when the BJT's base and emitter are reverse biased and the collector is open (I_EBO)|../common/conditionalProperty.json#/conditionalProperty| |
|dcCurrentGain|ratio of collector current to base current (hfe)|../common/conditionalProperty.json#/conditionalProperty| |
|collectorEmitterSaturationVoltage|collector-emitter voltage below which a change in base current does not impact collector current (VCE_SAT)|../common/conditionalProperty.json#/conditionalProperty| |
|baseEmitterSaturationVoltage|base-emitter voltage required to ensure the collector is forward biased for certain current conditions (VBE_SAT)|../common/conditionalProperty.json#/conditionalProperty| |
|collectorEmitterBreakdownVoltage|collector-emitter voltage at which a specified current flows with the base open|../common/unit.json#/unit| |
|baseEmitterBreakdownVoltage|base-emitter voltage at which a specified current flows with the collector open|../common/unit.json#/unit| |
|delayTime|time delay between input signal rising and when collector current rises to 10% of Isat (td)|../common/conditionalProperty.json#/conditionalProperty| |
|riseTime|time for collector current to rise through active region from 10% to 90% of Isat (tr)|../common/conditionalProperty.json#/conditionalProperty| |
|storageTime|time delay between input signal falling and when collector current falls to 90% of Isat (ts)|../common/conditionalProperty.json#/conditionalProperty| |
|fallTime|time for collector current to fall from 90% to 10% of Isat (tf)|../common/conditionalProperty.json#/conditionalProperty| |
|collectorCapacitance|parasitic capacitance of collector terminal under certain conditions (Cc)|../common/conditionalProperty.json#/conditionalProperty| |
|emitterCapacitance|parasitic capacitance of emitter terminal under certain conditions (Ce)|../common/conditionalProperty.json#/conditionalProperty| |
|transitionFrequency|frequency of unity current gain with a short circuit output (ft)|../common/conditionalProperty.json#/conditionalProperty| |
|icVsHfe|graph of collector current (Ic) as a function of dc current gain (Hfe)|../common/graph.json#/graph| |
|icVsVce|graph of collector current (Ic) as a function of collector-emitter saturation voltage (VCE_SAT)|../common/graph.json#/graph| |
|ibVsVce|graph of base current (Ib) as a function of collector-emitter voltage (VCE)|../common/graph.json#/graph| |
|pdVsTemp|graph of power dissipation as a function of temperature|../common/graph.json#/graph| |

####  3.16.2	 Specification Of Diode

Source: [diode.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/diode.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|type|type of diode|String| |
|diodeCount|number of diodes in the package|Number| |
|diodeConfiguration|configuration of diode|String| |
|numberOfProtectedLines|number of lines a diode can protect|Number| |
|vf|forward voltage of a diode|../common/conditionalProperty.json#/conditionalProperty| |
|if|continuous forward current of a diode|../common/unit.json#/unit| |
|ifm|maximum continuous forward current a diode can support|../common/unit.json#/unit| |
|ifrm|maximum repetitive peak forward current a diode can support|../common/unit.json#/unit| |
|ifsm|maximum non-repetitive surge forward current a diode can support|../common/unit.json#/unit| |
|vbr|breakdown voltage of a diode|../common/unit.json#/unit| |
|ir|reverse current|../common/conditionalProperty.json#/conditionalProperty| |
|vz|breakdown voltage of a zener diode|../common/conditionalProperty.json#/conditionalProperty| |
|vrm|maximum reverse standoff voltage a tvs diode can withstand|../common/unit.json#/unit| |
|vcl|clamping voltage of a tvs diode|../common/conditionalProperty.json#/conditionalProperty| |
|vr|maximum continuous reverse biased voltage a diode can support|../common/unit.json#/unit| |
|vrrm|maximum repetitive reverse voltage pulses a diode can support|../common/unit.json#/unit| |
|cd|diode junction capacitance - between the anode and cathode- in reverse bias condition|../common/unit.json#/unit| |
|trr|reverse recovery time it takes the diode to stop conducting when its voltage changes from forward-bias to reverse-bias|../common/unit.json#/unit| |
|pTot|maximum power dissipation of a forward biased diode|../common/conditionalProperty.json#/conditionalProperty| |
|ifVsVf|graph of forward current (If) as a function of forward voltage (VfTyp)|../common/graph.json#/graph| |

####  3.16.3	 Specification Of LED

Source: [led.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/led.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|ledColor|LED color|String| |
|vf|forward voltage of an LED |../common/conditionalProperty.json#/conditionalProperty| |
|if|continuous forward current of an LED|../common/unit.json#/unit| |
|ifp|peak forward current of an LED|../common/unit.json#/unit| |
|vr|maximum continuous reverse biased voltage an LED can support|../common/unit.json#/unit| |
|ir|reverse current (leakage) of an LED|../common/conditionalProperty.json#/conditionalProperty| |
|ledCapacitance|capacitance of an LED|../common/unit.json#/unit| |
|iv|LED luminous intensity|../common/conditionalProperty.json#/conditionalProperty| |
|peakWavelength|light spectrum output value emitted by an LED at highest wavelength|../common/conditionalProperty.json#/conditionalProperty| |
|dominantWavelength|dominant wavelength an LED emits the majority of the time|../common/conditionalProperty.json#/conditionalProperty| |
|angleHalfIntensity|angle at which LED intensity falls to 50% of its maximum value|../common/unit.json#/unit| |
|pd|power dissipation of an LED|../common/unit.json#/unit| |

####  3.16.4	 Specification Of Mosfet

Source: [mosfet.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/mosfet.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|mosfetType|type of MOSFET|String|Yes|
|mosfetChannelType|doping of a transistor's channel - describes whether a transistor is n-type or p-type|String| |
|transistorCount|number of transistors in the package|Number| |
|vgs|gate to source voltage difference of a MOSFET|../common/unit.json#/unit| |
|vgsMax|maximum gate to source voltage difference that can be continously applied to a MOSFET. This is a limiting value|../common/unit.json#/unit| |
|vds|drain to source voltage difference of a MOSFET|../common/unit.json#/unit| |
|vdsMax|maximum drain to source voltage difference that can be continously applied to a MOSFET. This is a limiting value|../common/unit.json#/unit| |
|vdsVbr|drain to source breakdown voltage of a MOSFET|../common/unit.json#/unit| |
|vgsTh|gate to source voltage difference required to produce a conducting path between drain and source|../common/unit.json#/unit| |
|vsdDiodeVf|reverse diode forward voltage when a MOSFET is in off-state|../common/unit.json#/unit| |
|iD|drain Current of a MOSFET|../common/unit.json#/unit| |
|iDrain|maximum continous DC current that can flow through a MOSFET channel.This is a limiting value|../common/unit.json#/unit| |
|idPulsed|maximum pulsed DC current that can flow through a MOSFET channel.This is a limiting value|../common/unit.json#/unit| |
|iDss|drain-source leakage current of a MOSFET when the gate to source voltage difference is zero|../common/conditionalProperty.json#/conditionalProperty| |
|iGss|gate-source leakage current of a MOSFET when the drain to source voltage difference is zero|../common/conditionalProperty.json#/conditionalProperty| |
|diodeContinuousCurrent|maximum continuous forward current of the body diode of a MOSFET (IS).This is a limiting value|../common/unit.json#/unit| |
|diodePulsedCurrent|maximum pulsed forward current of the body diode of a MOSFET. This is a limiting value|../common/unit.json#/unit| |
|forwardTransconductance|signal gain, change in drain current with variation of gate-source voltage of a MOSFET (gFS)|../common/conditionalProperty.json#/conditionalProperty| |
|rdson|on-state resistance of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|rg|internal gate resistance of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|ciss|input capacitance of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|coss|output capacitance of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|crss|reverse transfer capacitance of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|qg|total gate charge of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|qgd|gate to drain charge of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|qgs|gate to source charge of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|qrr|reverse recovery charge of the body diode of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|idVsVds|graph of drain current (iD) as a function of drain source voltage (vds)|../common/graph.json#/graph| |
|idVsVgs|graph of drain current (iD) as a function of gate source voltage (vgs)|../common/graph.json#/graph| |
|tdON|turn-on delay time of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|tdOFF|turn-off delay time of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|riseTime|rise time of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|fallTime|fall time of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|trr|reverse recovery time of the body diode of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|pTot|maximum power dissipation of a MOSFET|../common/conditionalProperty.json#/conditionalProperty| |
|pdVsTemp|graph of power dissipation vs temperature|../common/graph.json#/graph| |

### 3.17	 Sensor

####  3.17.1	 Specification Of Accelerometer

Source: [accelerometer.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/accelerometer.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|accelerometerType|type of accelerometer|String| |
|accelerationRanges|range of force that accelerometer can measure|array of Number| |
|accelerationSensitivity|smallest change in force the accelerometer is able to capture (typical) at a given acceleration range|Number| |
|accelerationSensitivityOverTemperature|accelerometer sensitivity change over temperature|Number| |
|axis|number of axes of acceleration measurement|Number| |
|zerogOffset|output of the accelerometer when no acceleration is applied|Number| |
|zerogOffsetOverTemperature|accelerometer zero-g offset change over temperature|Number| |
|outputType|measurement output type|String| |
|outputResolution|output resolution of acceleration measurement|../common/unit.json#/unit| |
|interface|interface(s) for communication to host|array of String| |
|bandwidth|bandwidth of acceleration measurement|../common/unit.json#/unit| |
|outputDataRate|output Data rate (ODR) of a device|../common/unit.json#/unit| |
|rmsNoise|broadband rms noise of a device|Number| |
|spectralNoiseDensity|spectral noise density of a device|Number| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

####  3.17.2	 Specification Of Gyroscope

Source: [gyroscope.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/gyroscope.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|gyroRanges|range of angular speed that gyro can measure|Number| |
|gyroSensitivity|smallest change in angular speed gyro is able to capture (typical) at a given gyro range|../common/conditionalProperty.json#/conditionalProperty| |
|gyroSensitivityOverTemperature|gyroscope sensitivity change over temperature|Number| |
|axis|number of axes of measurement|Number| |
|zeroRateOffset|output of the gyroscope when no angular velocity is applied|Number| |
|zeroRateOffsetOverTemperature|gyro zero rate offset change over temperature|Number| |
|interface|interface(s) for communication to host|array of String| |
|bandwidth|bandwidth of gyroscope|../common/unit.json#/unit| |
|outputDataRate|output Data rate (ODR) of a device|../common/unit.json#/unit| |
|outputType|measurement output type|String| |
|rmsNoise|broadband rms noise of a device|Number| |
|spectralNoiseDensity|spectral noise density of a device|Number| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

####  3.17.3	 Specification Of Magnetic Sensor

Source: [magnetic_sensor.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/magnetic_sensor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|magneticSensingTechnology|method by which magnetism is detected|String| |
|outputType|measurement output type|String| |
|quiescentOutput|output of the magnetic sensor when no magnet is present|../common/unit.json#/unit| |
|outputVoltageLinearRange|output voltage range over which the magnetic flux density response is linear|../common/unit.json#/unit| |
|linearMagneticSensingRange|magnetic flux density range over which the output voltage is linear |Number| |
|sensitivity|this is the gain - amount of change in the output voltage for a change in the magnetic flux density|Number| |
|operatePoint|magnetic flux dens####  3.4.1	 ComponentID

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
n Of Thermal Sensor

Source: [thermal_sensor.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/thermal_sensor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|sensingTechnology|method by which temperature is detected|String| |
|outputType|measurement output type|String| |
|interface|interface(s) for communication to host|array of String| |
|accuracy|accuracy of temperature sensor|../common/unit.json#/unit| |
|temperatureRange|range of temperature sensor|../common/unit.json#/unit| |
|resolution|maximum resolution of temperature sensor|Number| |
|gain|amount of change in the output voltage for a change in temperature|Number| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

### 3.18	 Storage

####  3.18.1	 Specification Of SD Card

Source: [sd_card.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/storage/sd_card.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|type|type of sd card|String| |
|capacity|capacity of sd card|../common/unit.json#/unit|Yes|
|dataRate|maximum data rate of the sd card|Number| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

####  3.18.2	 Specification Of SSD

Source: [ssd.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/storage/ssd.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|type|type of ssd storage as defined by interface and technology|String| |
|capacity|capacity of the ssd|../common/unit.json#/unit|Yes|
|dataRate|maximum data rate of the ssd|Number| |
|interface|interface of ssd to host|String| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |

