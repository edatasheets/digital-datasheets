
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

#### 3.4 Units
The digital datasheet shall follow the International System of Units, SI except for specific cases, like package dimensions, where the imperial system of units may be used. These specific cases will be addressed in the specifications.
Properties used the represent units in the datasheets are listed in Github: 
They include:
|Unit Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|siUnit|Name of the SI unit of measure|String|Yes|
|absoluteValue|Unit quantity corresponding to unit text|Number|Yes|
|unitText|Human readable text describing value|String|Yes|
|unitFactor|Multiplier on the value to achieve the SI unit|Number|Yes|
|relativeValueReference|if unit quantity is based on another reference, value of the reference|String| |
|relativeValueModifier|if a unit quantity is based on another reference, the value that edits that reference|Number| |
|relativeValueOperator|if a unit quantity is based on another reference, the operation that is performed with the modifier|String| |
|valueDefined|a boolean representing whether a value has been defined|Boolean| |

#### 3.5	Missing Values
Missing fields shall be left blank.

#### 3.6	Pins Specifications

The pins specification is included in Github at:https://github.com/edatasheets/edatasheets.github.io/blob/main/part-spec/definitions.json

Table below shows the list of properties used to describe pins

|Pin Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|terminalIdentifier|Pin or ball number as defined by datasheet|String|Yes|
|name|Name given to the signal appearing at the terminal of a component"|String|Yes|
|standardizedName  |Standard Name of pin|String| |
|numberOfSupportedFunctions|The total number of functions supported by this pin|Number | |
|functionProperties|List of function objects that can apply to an individual pin|Array| |
|vihMin|The least positive (most negative) value of high-level input voltage for which operation of the logic element within specification limits is to be expected.|Number| |
|vihMax|The most positive (least negative) value of high-level input voltage for which operation of the logic element within specification limits is to be expected. |Number| |
|vilMin|The least positive (most negative) value of low-level input voltage for which operation of the logic element within specification limits is to be expected.|Number| |
|vilMax|The most positive (least negative) value of low-level input voltage for which operation of the logic element within specification limits is to be expected|Number| |
|vol|The voltage level at an output terminal with input conditions applied that,according to the product specification, will establish a low level at the output|Number| |
|voh|The voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a high level at the output.|Number| |
|absVmax |Maximum voltage rating beyond which damage to the device may occur.|Number| |
|absVmin |Absolute minimum voltage that can be applied to a pin|Number| |
|imax|Maximum continuous current that can safely be drawn from a pin|Number| |
|inputLeakage|Maximum current draw into a high impedance input pin|Number| |
|outputLeakage|Maximum current flow from a pin during the off state|Number| |
|dcResistance|Resistance of a pin of a connector|Number| |
|voltageOptions|List of voltage levels supported by a pin|Array| |
|floatUnused|Describes  whether pin can safely be floated if it is not used|Boolean| |
|internalPullUp|Indicates the value of an internal pull-up on a pin|Number| |
|internalPullDown|Indicates the value of an internal pull-down on a pin|Number| |
|esd|Indicates whether ESD protection exists on a pin|Boolean| |
|externalComponents|List of external component structures recommended to be attached to a pin|Object| |
|configuration|electrical configuration of component connected to pin with respect to the pin.|String| |
|minValue|Minimum value of component if a range is specified|Number| |
|maxValue|Maximum value of component if a range is specified|Number| |
|value|Value of component if a range is not specified|Number| |
|componentUnit|Unit of min/max value|String| |
|connectionPin|Name of pin to which an external component should be pulled up|String| |
|perFunctionName|Name of the function of a pin|String| |
|interfaceType|Type of interface enabled by pin|String| |
|pinUsage|Standardized usage of pin|String| |
|direction|Direction of a pin's function|String| |
|electricalConfiguration|Electrical configuration of a pin|String| |
|polarity|Whether the active state of a pin is high or low|String| |


#### 3.7	Graph Specifications
Datasheets include additional information through figures and graphs. Examples include but are not limited to efficiency vs load curves and power derating curves. A graph object and a curve object are specified to capture that information in the digital datasheet. The graph object focuses on defining the type of data captured, including title and axis information. The curve object focuses on capturing the data. Multiple curve objects may be included in a graph object to capture relationships between variables under different conditions. Examples include load current vs efficiency curves for different input voltages.Properties used to specify a curve object are included in the table below:
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|label|Description of the data in a curve|string| |
|xData|x value of data being plotted|array of numbers|Yes|
|yData|y value of data being plotted|array of numbers|Yes|

Properties used to specify a graph object are included below:
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|title|Title of a graph|string|Yes|
|xUnits|x-axis units|string|Yes|
|xLabel|x-axis title|String|Yes|
|yUnits|y-axis units|string|Yes|
|yLabel|y-axis title|String|Yes|
|numberOfCurves|total number of curves in graph|Number|Yes|

#### 3.8	Physical Package Specifications
The digital datasheet has specifications to capture the first level physical package specifications. Properties used to specify a package are included below:
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|length|Length of the package|Number|Yes|
|width|Width of the package|Number|Yes|
|height|Height of the package|Number|Yes|
|dimensionUnit|Unit (mm or mil) used to describe the package dimensions|String|Yes|
|standardPackageSize|Standard name used to designate the package size|String| |
|standardPackageType|Standard name used to designate the package type|String| |

#### 3.9	Passive Components Specifications
The passive components digital datasheet specifications are included in GitHub: https://github.com/edatasheets/edatasheets.github.io/tree/main/part-spec

#### 3.9.1	Resistor Specifications
The table below give a description of the properties used to specify a resistor in a digital datasheet.
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|Resistance value|Number|Yes|
|tolerance|Nominal tolerance of a resistor|Number| |
|powerRating|Maximum power a resistor can dissipate without degrading performance|Number| |
|temperatureCoefficient|Change in resistance when the temperature is changed|Number| |
|maxOverloadVoltage|Maximum voltage that can be applied to the resistor for a short period of time|Number| |
|maxLimitingElementVoltage|Maximum voltage value that can be applied continuously to the resistor|Number| |
|pinSpec|Pins definition of the resistor|object| |
|package|Package definition of the resistor|Object| |
|minTemperature|Minimum temperature under which a resistor can be expected to reliably operate|Number| |
|maxTemperature|Maximum temperature under which a resistor can be expected to reliably operate|Number| |
|complianceList|List of compliances met by the resistor|Array| |
|resistorDerating|Graph to capture resistor derating with temperature|Object| |


#### 3.9.2	Capacitor Specifications
The table below gives a description of the properties used to specify a capacitor in a digital datasheet.
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|Capacitance value|Number|Yes|
|tolerance|Nominal tolerance of a capacitor|Number| |
|ratedVoltage|Maximum voltage which may be applied continuously to a capacitance |Number| |
|dielectric|Dielectric material used in the capacitor|String| |
|polarized|Describes whether the capacitor is polarized|Boolean| |
|esr|Equivalent Series resistance of the capacitor|Number| |
|minTemperature|Minimum temperature under which a capacitor can be expected to reliably operate|Number| |
|maxTemperature|Maximum temperature under which a capacitor can be expected to reliably operate|Number| |
|temperatureCoefficient|Change in capacitance when the temperature is changed|Number| |
|pinSpec|Pins definition of the capacitor|object| |
|package|Package definition of the capacitor|Object| |
|complianceList|List of compliances met by the capacitor component|Array| |
|capacitorDerating|Graph to capture capacitor derating with voltage|Object| |

#### 3.9.3	Inductor Specifications
The table below gives a description of the properties used to specify an inductor in a digital datasheet.
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|value|Inductance value|Number|Yes|
|tolerance|Nominal tolerance of an Inductor|Number| |
|ratedCurrent|Maximum continuous current the inductor can handle|Number| |
|saturationCurrent|Current where the inductor enters the magnetic state, and the inductance drops a specified amount |Number| |
|rmsCurrent|DC current that produces an inductor temperature rise of 40 degrees Celsius|Number| |
|selfResonantFrequency|Frequency at which the inductor becomes a capacitor|Number| |
|dcResistance|DC resistance of the inductor|Number| |
|minTemperature|Minimum temperature under which an inductor can be expected to reliably operate|Number| |
|maxTemperature|Maximum temperature under which an inductor can be expected to reliably operate|Number| |
|temperatureCoefficient|Change in inductance when the temperature is changed|Number| |
|pinSpec|Pins definition of the inductor|object| |
|package|Package definition of the inductor|Object| |
|complianceList|List of compliances met by the inductor component|Array| |
|saturationCurve|Graph to capture current saturation curve|Object| |
|resonantFrequencyCurve|Graph to capture resonant frequency curve|Object| |

#### 3.10	Power Components Specifications
The power component specification is included in Github at: https://github.com/edatasheets/edatasheets.github.io/tree/main/part-spec

#### 3.10.1	 Power FET Specifications
Power FETs are important components of swithing regulators. They can be integrated in the regulator or external to the regulator. The table below gives a description of the properties used to specify FETs in a digital datasheet. 
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|ilimHSFET|High Side FET maximum current above which the output voltage starts dropping |Number| |
|ilimLSFET|Low Side FET maximum current above which the output voltage starts dropping|Number| |
|rdsonHSFET|High side FET on-resistance|Number| |
|rdsonLSFET|Low side FET on-resistance|Number| |

#### 3.10.2	 Switching Regulator Specifications
The table below gives a description of the properties used to specify a switching regulator in a digital datasheet.
|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|regulatorTopology|Switching regulator topology|String|Yes|
|vinMin|Minimum input voltage under which the part can be expected to operate properly |Number|Yes|
|vinMax|Maximum input voltage under which the part can be expected to operate properly|Number|Yes|
|voutMin|Minimum output voltage the part can regulate|Number|Yes|
|voutmax|Maximum output voltage the part can regulate|Number|Yes|
|feedbackVoltage|Voltage at the feedback node|Number| |
|voutAccuracy|Output voltage variation|Number| |
|minLoadCurrent|Minimum output current the part can support without going out of regulation|Number| |
|maxLoadCurrent|Maximum output current the part can support without going out of regulation|Number|Yes|
|loadRegulation|Output voltage variation from no load to full load|Number| |
|lineRegulation|Output voltage variation from minimum input voltage to maximum input voltage|Number| |
|quiescentCurrent|Quiescent current of voltage regulator at no load|Number| |
|shutdownCurrent |Shutdown current of voltage regulator|Number| |
|switchingFrequency|Switching frequency of voltage regulator|Number| |
|integratedFets|Describes whether the FETs are integrated in the regulator package|Boolean| |
|powerFetProperties| Current limits and rdson values for the integrated FETs|Object| |
|enableTime|Time between enable asserted and output voltage rising to 10% nominal|Number| |
|rampTime|Time for output to go from 10% nominal output voltage to 90%nominal output voltage|Number| |
|underVoltageLockoutThresholdRising|Undervoltage Lockout threshold with input voltage rising|Number| |
|underVoltageLockoutThresholdFalling|Undervoltage Lockout threshold with input voltage falling|Number| |
|overVoltageProtectionThresholdRising|Overvoltage Protection threshold with input voltage rising|Number| |
|overVoltageProtectionThresholdFalling|Overvoltage Protection threshold with input voltage Falling|Number| |
|thermalShutdownThresholdRising|Thermal Shutdown Threshold with temperature rising|Number| |
|thermalShutdownThresholdFalling|Thermal Shutdown Threshold with temperature Falling|Number| |
|thermalShutdownHysteresis|Thermal Shutdown Hysteresis|Number| |
|efficiency|Power efficiency Graph of the switching regulator|object| |
|pinSpec|Pins definition of the switching regulator|object| |
|package|Package definition of the switching regulator|Object| |
|complianceList|List of compliances met by the switching regulator component|Array| |

### 3.10	 Hardware

####  3.10.1	 Switch

The table below gives a description of the properties used to specify a switch in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|type|property describing the way in which the switch is activated|String| |
|contactType|property describing the order in which switch contact is made and broken|String| |
|circuitConfig|property describing the number of poles and throws in a switch|String| |
|cycleRating|number of on/off cycles a mechanical switch can reliably sustain|Number| |
|voltageRating|maximum DC voltage potential that can be applied across an open switch|Object| |
|currentRating|maximum DC current that can flow through a closed switch without causing excessive heating|Object| |
|onResistance|nominal resistance of a closed switch|Object|Yes|
|dielectricRating|maximum AC voltage potential that can be applied across an open switch for one minute|Object| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.10.2	 Connector

The table below gives a description of the properties used to specify a connector in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|function|intended function of a connector|String| |
|contactCount|number of contacts in a connector|Number| |
|type|property describing the method of mating to the connector|String| |
|cycleRating|number of plug/unplug cycles a connector is rated to support|Number| |
|pitch|distance from the center of one contact on the connector to the center of the next contact|Object| |
|keying|property describing whether a connector has an asymmetry to prevent it from being plugged in the wrong direction|Boolean| |
|numberPins|number of pins on the connector|Number|Yes|
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

### 3.11	 IC IO (Integrated Circuit Input/Output)

ICs often require various external components to correctly use them in a design.
This section contains components commonly used to fit that need.

####  3.11.1	 Redriver

The table below gives a description of the properties used to specify a redriver in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|numberChannels|number of lanes (single ended or differential) supported by redriver|Number| |
|interface|list of interface(s) supported by redriver|String| |
|maxDataRate|maximum data rate supported by redriver|Object| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.11.2	 Bridge Chip

The table below gives a description of the properties used to specify a bridge chip in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|muxRatio|ratio of inputs to outputs|String| |
|inputInterfaces|list of interfaces at the input of the bridge|String| |
|outputInterfaces|list of interfaces at the output of the bridge|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.11.3	 Mux

The table below gives a description of the properties used to specify a mux in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|muxRatio|ratio of inputs to outputs|String| |
|inputInterfaces|list of interfaces mux is designed for|String| |
|insertionLoss|insertion loss through mux|Number| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.11.4	 Level Shifter

The table below gives a description of the properties used to specify level shifters in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|numberChannels|number of lanes (single ended or differential) supported by redriver|Number| |
|interface|interface supported by redriver|String| |
|inputVoltage|input voltage level of level shifter|Object| |
|outputVoltage|output voltage level of level shifter|Object| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

### 3.12	 Logic Integrated Circuits

This section contains ICs which can be used to implement digital logic in a design.

####  3.12.1	 Logic Gate

The table below gives a description of the properties used to specify logic gates in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|type|logical operation performed by logic gate|String|Yes|
|numberGates|number of logical gates encapsulated in logic IC|Number| |
|schmittTrigger|property describing whether logic gate has schmitt trigger inputs|Boolean| |
|propagationDelay|time between input changing to output changing|Object| |
|rampTime|time for output to go from 10% nominal output voltage to 90% nominal output voltage|Object| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.12.2	 Clock

The table below gives a description of the properties used to specify a clock in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|fixedFrequency|clock frequency value if the clock has a fixed frequency|Object| |
|minFrequency|minimum clock frequency value if the clock has a configurable frequency|Object| |
|maxFrequency|maximum clock frequency value if the clock has a configurable frequency|Object| |
|numberClockOutputs|number of clock outputs in a clock IC|Number| |
|diffSingleEnded|property describing whether a clock output is single ended or differential|String| |
|jitter|cycle to cycle clock jitter|Object| |
|frequencyTolerance|amount of frequency variation specced from nominal frequency|Object| |
|powerSupplyRejectionRatio|power supply rejection ratio (PSRR)or ratio between power supply variation and output variation|Object| |
|outputFormat|signal format of clock output|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

### 3.13	 Microcontrollers

This section contains non-trivial microcontroller components.

####  3.13.1	 Microcontroller/EC (Electronic Controller)

The table below gives a description of the properties used to specify an embedded controller ("EC") in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|onChipFlash|quantity of built-in flash in a microprocessor|Object| |
|onChipRAM|quantity of built-in RAM in a microprocessor|Object| |
|onChipROM|quantity of built-in ROM in a microprocessor|Object| |
|coreProcessor|description of core processor|String| |
|coreArchitectureBits|number of bits of data a CPU can transfer per clock cycle|String| |
|clockSpeed|speed of main CPU clock|Object| |
|firmwareVersion|firmware version of the part|String| |
|activePower|average power of device in active state|Object| |
|standbyPower|average power of device in standby state|Object| |
|pins|array of pin objects with associated properties|array of Object| |
|package|component's package size and description|Object| |

### 3.14	 IC Misc

This section contains miscellaneous integrated circuits.

####  3.14.1	 Speaker Amplifier

The table below gives a description of the properties used to specify a speacker amplifier in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|dataLength|number of bits in a data word|Number| |
|outputPower|typical output power from speaker amplifier|Object| |
|efficiency|typical speaker amplifier efficiency|Object| |
|thd+n|typical total harmonic distortion plus noise of amplifier|Object| |
|sampleRate|sample rate of data out of amplifier|Object| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.14.2	 Audio Codec

The table below gives a description of the properties used to specify an audio codec in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|dataLength|number of bits in a data word|Number| |
|hpOutputSNR|headphone amplifier output SNR|Number| |
|hpOutputTHD+N|headphone output total harmonic distortion plus noise|Object| |
|micInputSNR|microphone input SNR|Object| |
|micInputTHD+N|microphone input total harmonic distortion plus noise|Object| |
|jackDetect|describes whether headphone jack detection is supported|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.14.3	 WLAN Module

The table below gives a description of the properties used to specify a WLAN module in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|wlanSpec|version of wlan specification supported by module|String| |
|bluetoothVersion|version of bluetooth supported by module|String| |
|txrxChains|number of tx and rx chains in a wifi module|String| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|lteCoexFilter|describes whether module supports lte coexistance filtering|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.14.4	 WWAN Module

The table below gives a description of the properties used to specify WWAN module in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|networkSupport|networks supported by wwan module|String| |
|gpsSupport|whether wwan module has gps support|Boolean| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.14.5	 TPM (Trusted Platform Module)

The table below gives a description of the properties used to specify a TPM in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|interface|describes the communication interface from the chip to the host|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

### 3.15	 Storage/memory

This section contains data storage components, for both volatile and
non-volatile memory, as well as read-only and read-write memory.

####  3.15.1	 SSD

The table below gives a description of the properties used to specify an SSD in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|type|type of ssd storage as defined by interface and technology|String| |
|capacity|capacity of SSD|Object|Yes|
|dataRate|maximum data rate|Object| |
|interface|interface of ssd to host|String| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.15.2	 SD Card

The table below gives a description of the properties used to specify an SD card in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|type|type of sd card|String| |
|capacity|capacity of SD card|Object|Yes|
|dataRate|maximum data rate|Object| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.15.3	 DRAM

The table below gives a description of the properties used to specify DRAM in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|type|type of dram|String| |
|capacity|capacity of dram chip|Object|Yes|
|speed|dram maximum speed|Object| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

####  3.15.4	 ROM

The table below gives a description of the properties used to specify ROM in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|capacity|capacity of rom|Object| |
|interface|interface of rom to host|String| |
|qeStatus|indicates whether the QE bit is set|Boolean| |
|package|component's package size and description|Object| |
|pins|array of pin objects with associated properties|array of Object| |

### 3.16	 USB-C

This section contains components related to implementing USB-C.

####  3.16.1	 USB-C Power Delivery Controller

The table below gives a description of the properties used to specify a USB-C PD controller in a digital datasheet.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|Object|Yes|
|pdVersion|version of power delivery spec implemented by controller|String|Yes|
|usbTypecRevision|usb type-c spec revision implemented by controller|String| |
|powerRoleSupported|roles supported by pd controller|String| |
|fastRoleSwapSupport|whether the pd controller supports fast role swap (FRS)|Boolean| |
|vconnPowerSupport|whether the pd controller has support for vconn power|Boolean| |
|vconnPowerLimit|power limit supported by internal vconn switch (if supported)|Object| |
|vconnMaxCurrent|maximum continuous current supported by internal vconn switch (if supported)|Object| |
|vconnOverCurrentLimit|over current limit supported by internal vconn switch (if supported)|Object| |
|integratedVbusDischargeSwitch|whether the pd controller has one or more integrated vbus discharge switches |Boolean| |
|integratedLoadSwitch|whether the pd controller has one or more integrated load switches |Boolean| |
|maxSinkCurrent|maximum continuous current supported by pd controller integrated sink load switch|Object| |
|maxSourceCurrent|maximum continuous current supported by pd controller integrated source load switch|Object| |
|sinkfetOverCurrentLimit|over current limit supported by pd controller integrated sink load switch|Object| |
|sourcefetOverCurrentLimit|over current limit supported by pd controller integrated source load switch|Object| |
|onResistanceSinkFet|on-resistance of the integrated sink load switch|Object| |
|onResistanceSourceFet|on-resistance of the integrated source load switch|Object| |
|activeCurrent|active current of pd controller (during PD communication)|Object| |
|shutDownCurrent|shutdown current of pd controller|Object| |
|idleCurrent|idle current of pd controller (cable connected but no PD communiation)|Object| |
|thermalShudownThresholdRising|Thermal Shudown (tsd) Threshold with temperature rising|Object| |
|thermalShudownThresholdFalling|Thermal Shudown (tsd) Threshold with temperature falling|Object| |
|vbusOvervoltageProtectionThresholdRising|Overvoltage Protection (OVP) Threshold with vbus voltage rising|Object| |
|vbusOvervoltageProtectionThresholdFalling|Overvoltage Protection (OVP) Threshold with vbus voltage falling|Object| |
|vbusUndervoltageLockoutThresholdRising|Undervoltage Lockout Out (UVLO) Threshold with vbus voltage rising|Object| |
|vbusUndervoltageLockoutThresholdFalling|Undervoltage Lockout Out (UVLO) Threshold with vbus voltage rising|Object| |
|vconnOvervoltageProtectionThresholdRising|Overvoltage Protection (OVP) Threshold with vconn voltage rising|Object| |
|vconnOvervoltageProtectionThresholdFalling|Overvoltage Protection (OVP) Threshold with vconn voltage falling|Object| |
|vconnUndervoltageLockoutThresholdRising|Undervoltage Lockout Out (UVLO) Threshold with vconn voltage rising|Object| |
|vconnUndervoltageLockoutThresholdFalling|Undervoltage Lockout Out (UVLO) Threshold with vconn voltage rising|Object| |
|pins|array of pin objects with associated properties|array of Object| |
|package|component's package size and description|Object| |

### 3.17	 Semiconductor

Source: [semiconductor.json](https://github.com/chromeos/digital-datasheets/blob/main/part-spec/semiconductor.json)

####  3.17.1	 Mosfet

The table below gives a description of the properties used to specify a MOSFET.

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
|vsdDiodeVfTyp|typical reverse diode forward voltage when a MOSFET is in off-state.|definitions.json#/unit| |
|vsdDiodeVfMax|maximum reverse diode forward voltage when a MOSFET is in off-state.|definitions.json#/unit| |
|iD|Drain Current of a MOSFET.|definitions.json#/unit| |
|iDrain|maximum continous DC current that can flow through a MOSFET channel.This is a limiting value.|definitions.json#/unit| |
|idPulsed|maximum pulsed DC current that can flow through a MOSFET channel.This is a limiting value.|definitions.json#/unit| |
|iDss|drain-source leakage current of a MOSFET when the gate to source voltage difference is zero|definitions.json#/conditionalProperty| |
|iGss|gate-source leakage current of a MOSFET when the drain to source voltage difference is zero|definitions.json#/conditionalProperty| |
|diodeContinuousCurrent|maximum continuous forward current of the body diode of a MOSFET (IS).This is a limiting value.|definitions.json#/unit| |
|diodePulsedCurrent|maximum pulsed forward current of the body diode of a MOSFET. This is a limiting value.|definitions.json#/unit| |
|forwardTransconductance|signal gain, change in drain current with variation of gate-source voltage of a MOSFET (gFS).|definitions.json#/conditionalProperty| |
|rdsonTyp|typical on-state resistance of a MOSFET.|definitions.json#/conditionalProperty| |
|rdsonMax|maximum on-state resistance of a MOSFET.|definitions.json#/conditionalProperty| |
|rgTyp|typical internal gate resistance of a MOSFET.|definitions.json#/conditionalProperty| |
|rgMax|maximum internal gate resistance of a MOSFET.|definitions.json#/conditionalProperty| |
|cissTyp|typical input capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|cissMax|maximum input capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|cossTyp|typical output capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|cossMax|maximum output capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|crssTyp|typical reverse transfer capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|crssMax|maximum reverse transfer capacitance of a MOSFET.|definitions.json#/conditionalProperty| |
|qgTyp|typical total gate charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgMax|maximum total gate charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgdTyp|typical gate to drain charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgdMax|maximum gate to drain charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgsTyp|typical gate to source charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qgsMax|maximum gate to source charge of a MOSFET.|definitions.json#/conditionalProperty| |
|qrrTyp|typical reverse recovery charge of the body diode of a MOSFET.|definitions.json#/conditionalProperty| |
|qrrMax|maximum reverse recovery charge of the body diode of a MOSFET.|definitions.json#/conditionalProperty| |
|idVsVds|graph of drain current (iD) vs drain source voltage (vds).|graph.json#/graphDefiniton| |
|idVsVgs|graph of drain current (iD) vs gate source voltage (vgs).|graph.json#/graphDefiniton| |
|tdONTyp|typical turn-on delay time of a MOSFET|definitions.json#/conditionalProperty| |
|tdONMax|maximum turn-on delay time of a MOSFET|definitions.json#/conditionalProperty| |
|tdOFFTyp|typical turn-off delay time of a MOSFET|definitions.json#/conditionalProperty| |
|tdOFFMax|maximum turn-off delay time of a MOSFET|definitions.json#/conditionalProperty| |
|riseTimeTyp|typical rise time of a MOSFET|definitions.json#/conditionalProperty| |
|riseTimeMax|maximum rise time of a MOSFET|definitions.json#/conditionalProperty| |
|fallTimeTyp|typical fall time of a MOSFET|definitions.json#/conditionalProperty| |
|fallTimeMax|maximum fall time of a MOSFET|definitions.json#/conditionalProperty| |
|trrTyp|typical reverse recovery time of the body diode of a MOSFET.|definitions.json#/conditionalProperty| |
|trrMax|maximum reverse recovery time of the body diode of a MOSFET.|definitions.json#/conditionalProperty| |
|pTot|maximum power dissipation of a MOSFET.|definitions.json#/conditionalProperty| |
|pdVsTemp|graph of power dissipation vs temperature.|graph.json#/graphDefiniton| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |
|package|component's package size and description|definitions.json#/package| |

####  3.17.2	 Diode

The table below gives a description of the properties used to specify a diode.

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|common component identifying information, such as mpn|definitions.json#/componentID|Yes|
|type|type of diode|String| |
|diodeCount|number of diodes in the package|Number| |
|diodeConfiguration|configuration of diode|String| |
|numberOfProtectedLines|number of lines a diode can protect|Number| |
|vfTyp|typical forward voltage of a diode|definitions.json#/conditionalProperty| |
|vfMax|maximum forward voltage of a diode|definitions.json#/conditionalProperty| |
|vfMin|minimum forward voltage of a diode|definitions.json#/conditionalProperty| |
|if|continuous forward current of a diode|definitions.json#/unit| |
|ifm|maximum continuous forward current a diode can support|definitions.json#/unit| |
|ifrm|maximum repetitive peak forward current a diode can support|definitions.json#/unit| |
|ifsm|maximum non-repetitive surge forward current a diode can support|definitions.json#/unit| |
|vbr|breakdown voltage of a diode|definitions.json#/unit| |
|irTyp|typical reverse current|definitions.json#/conditionalProperty| |
|irMax|maximum reverse current|definitions.json#/conditionalProperty| |
|vzTyp|typical breakdown voltage of a zener diode|definitions.json#/conditionalProperty| |
|vzMax|maximum breakdown voltage of a zener diode|definitions.json#/conditionalProperty| |
|vzMin|minimum breakdown voltage of a zener diode|definitions.json#/conditionalProperty| |
|vrm|maximum reverse standoff voltage a tvs diode can withstand|definitions.json#/unit| |
|vclTyp|typical clamping voltage of a tvs diode|definitions.json#/conditionalProperty| |
|vclMax|maximum clamping voltage of a tvs diode|definitions.json#/conditionalProperty| |
|vclMin|minimum clamping voltage of a tvs diode|definitions.json#/conditionalProperty| |
|vr|maximum continuous reverse biased voltage a diode can support|definitions.json#/unit| |
|vrrm|maximum repetitive reverse voltage pulses a diode can support|definitions.json#/unit| |
|cd|diode junction capacitance - between the anode and cathode- in reverse bias condition|definitions.json#/unit| |
|trr|reverse recovery time it takes the diode to stop conducting when its voltage changes from forward-bias to reverse-bias|definitions.json#/unit| |
|pTot|maximum power dissipation of a forward biased diode|definitions.json#/conditionalProperty| |
|ifVsVf|graph of forward current (If) vs forward voltage (VfTyp)|graph.json#/graphDefiniton| |
|package|package size of resistor|definitions.json#/package| |
|pins|array of pin objects with associated properties|array of definitions.json#/pinSpec| |

