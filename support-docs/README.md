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
Required: The field is required in the component digital datasheet.
Optional: The field may or may not be included in the component digital datasheet.
Shall: Indicates a mandatory requirement.
May: Indicates an optional requirement.

#### 1.4 References
- The JSON data interchange syntax, ECMA-404, 2nd edition, December 2017 
- JSON Schema: A media Type for Describing JSON Documents, draft-bhutton-json-schema-00, December 2020.
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



