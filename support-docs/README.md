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

### 2. Working with Digital Datasheets 

#### 2.1 Digital Datasheet Distribution by Component Vendors
A component e-datasheet released by a hardware component vendor can include one or multiple files. Multiple files are typically used when a single datasheet covers multiple parts from the same family, when a part is complex or, when a part has additional information, not in json format, such as footprints. In this case, the specification has a provision that enables vendors to capture common or core properties across all the parts in one file while using additional files to capture the changes between parts.

There are multiple existing models for hardware component vendors to distribute their parts datasheets. The recommendation is for these vendors to leverage these existing models for e-datasheets. Existing models include private and public distributions directly from the vendors or through an electronic components’ distributor.

In a public release, the hardware component vendor makes their latest e-datasheets available to the public, generally, through a public website. An example of this might be a Github repository where each vendor has a repository. The e-datasheets can be directly downloaded through the website without any requirements, and customers can optionally add their contact information to be notified about updates in the e-datasheets. 

Electronic components distributors have agreements with hardware component vendors to distribute their products to customers.  In the existing model, distributors either include a local copy of the component datasheet to their website or add links to the datasheets on the vendor’s website.  In the case of e-datasheets, the recommendation is to use the later model, with an e-datasheet link always pointing to the latest release.  The goal is to avoid a situation where copies of local e-datasheet become out of sync with the latest release. 

In a private distribution model, the e-datasheet is shared by a hardware component vendor to one or more select customers, for Intellectual Property (IP) protection. This may require an NDA between the parties involved. The e-datasheet can be shared through means like a private repository or a private database, or through encrypted emails.  With this model, the vendors have a good visibility of the customers using their e-datasheets and can directly notify them of updates through email.  It’s also recommended to provide a release document identifying e-datasheets changes from one release to the next.

#### 2.2 Importing Digital Datasheets in a Part Library Management System
Most companies doing hardware design rely on a Part Library Management (PLM) System to store information about electronic components. These systems are often linked to CAD tools to guide component selection during hardware design. Additionally, these systems have an important role in Bill of Materials (BOM) generation and purchasing. Therefore, a natural entry point for digital datasheet information is a PLM system.

Given the importance of component accuracy in a PLM system, it is likely that companies will initially want some manual review of new digital datasheets before they are integrated into a shared PLM system. One suggested method of achieving this is to have an internal datasheet database that is connected to any datasheet repositories (public or private) that are used by that company. Updates can be manually pulled in and compared as new datasheets are available. Once the updates have been checked, new datasheet changes can be submitted to the internal database. This database would be indexed by internal part numbers that correspond to the digital datasheet.

Converting this internal company datasheet database to information in the PLM system can be done automatically because datasheets in the internal database have already been manually verified. To do this conversion, a one-time mapping file between properties in the digital datasheet schema and the specific properties in a company’s PLM is created. Companies can then write a simple generator script using this mapping file to convert digital datasheet entries to PLM properties. Additionally, as digital datasheets will likely contain new information not available in the PLM system, this information can be added or exist in a separate database alongside the PLM system. As tooling advances, these new properties can be leveraged to provide richer schematic validation. 

The advantage of this system is it leverages existing PLM systems, which are often deeply integrated with engineering workflows. However, it makes these PLM systems more efficient to generate and update.

#### 2.3 Tools to Generate Digital Datasheets 
The digital datasheet creator is an open-source tool that is available to make datasheet
creation less time consuming. The tool can be found here: [datasheet creator](https://github.com/intel/digital-datasheet-creator). It consists of a series of template spreadsheets for each 
part type that can be filled in and run through the creator script to generate a specification-compliant json datasheet file. 

#### 2.3 Tools to Generate Digital Datasheets 
The digital datasheet creator is an open-source tool that is available to make datasheet
creation less time consuming. The tool can be found here: [datasheet creator](https://github.com/intel/digital-datasheet-creator). It consists of a series of template spreadsheets for each 
part type that can be filled in and run through the creator script to generate a specification-compliant json datasheet file. 

### 3. Use Cases 
Digital datasheets can be used for many applications. The list of use cases included here is not meant to be exhaustive and it is expected that new applications will be developed as more people start using these digital datasheets. Sample applications include automated hardware design checks to identify bugs earlier in the design cycle, automated hardware designs to speed up board development, components comparison to identify replacement components on a design. 

### 4.Specifications

#### 4.1 Digital Datasheet Format
A digital datasheet shall be written in the JSON (JavaScript Object Notation) format. The specification is written in JSON schema to facilitate validation of a JSON datasheet. 

#### 4.2 Required Information.
A digital datasheet shall include the following information:
- The manufacturer’s name
- The Manufacturer Part Number (MPN)
- Information to identify the source datasheet
- The datasheet version given by date or GUID 
- The component part type. See appendix for examples.
- List of Component Pins as defined by the specification.

#### 4.3 Date Format
A digital datasheet shall follow the international standard notation, YYYY-MM-DD.

####  4.4	 Top Level Component Specification 

Source: [component.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/component.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentID|methods for identifying the version of the digital datasheet|./common/componentID.json#/componentID|Yes|
|coreProperties|core component properties as defined by the specific component spec file|./common/coreProperties.json#/coreProperties|No|
|pins|array of pin objects with associated properties|./common/pinSpec.json#/pinSpec|No |
|package|component package information|./common/package.json#/package|No|
|register|register information|./common/register.json#/register|No|
|thermal|component temperature and thermal resistance information|./common/thermal.json#/thermal|No|
|componentPropertyExternalFiles|external files that describe key component properties. External files can be used in lieu of defining core properties, pins, and package information in the same file|#/$defs/externalFileMap|No|
|additionalSpecExternalFiles|external files that contain information outside of the json spec. Examples include layout, simulation, etc.|./common/externalFile.json#/externalFile|No|
|reliability|reliability information about the component|./common/reliability.json#/reliability|No|
|powerSequence|information about component power sequencing|./common/powerSequence.json#/powerSequenceTable|No|

### 4.5	 Common

####  4.5.1	 Specification To Capture Information To Identify Components

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

####  4.5.2	 DigitalDatasheetID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|publishedDate|date the digital datasheet was published|String| |
|eDatasheetSpecRevision|revision of the digital datasheet specifications used to create the digital datasheet|String|Yes|
|guid|vendor defined guid (see https://www.guidgenerator.com/) to uniquely identify digital datasheet version|String| |

####  4.5.3	 SourceDatasheetID

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|publishedDate|date the source datasheet was published|String| |
|version|version of the source datasheet|String| |
|datasheetURI|uri to the source datasheet pdf or html view|String| |
|productURI|uri to the source datasheet's product page'|String| |

####  4.5.4	 Specification To Capture Protection Thresholds Data Of a Component

Source: [componentProtectionThresholds.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/componentProtectionThresholds.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|thermalShutdownThresholdRising|thermal Shutdown (tsd) Threshold with temperature rising|../common/values.json#/valueOptions| |
|thermalShutdownThresholdFalling|thermal Shutdown (tsd) Threshold with temperature falling|../common/values.json#/valueOptions| |
|thermalShutdownHysteresis|thermal Shutdown (tsd) Hysteresis|../common/values.json#/valueOptions| |
|powerSupplyProtection|undervoltage lockout, overvoltage protection thresholds of a supply|array of #/$defs/powerSupplyProtection| |

####  4.5.5	 PowerSupplyProtection

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|supplyPinName|name of the power supply pin or any other pin monitored|String| |
|overVoltageProtectionThresholdRising|Overvoltage Protection (OVP) Threshold with power supply rising|../common/values.json#/valueOptions| |
|overVoltageProtectionThresholdFalling|Overvoltage Protection (OVP) Threshold with power supply falling|../common/values.json#/valueOptions| |
|overVoltageProtectionHysteresis|Overvoltage Protection (OVP) Hysteresis|../common/values.json#/valueOptions| |
|underVoltageLockoutThresholdRising|Undervoltage Lockout (UVLO) Threshold with power supply rising|../common/values.json#/valueOptions| |
|underVoltageLockoutThresholdFalling|Undervoltage Lockout (UVLO) Threshold with power supply falling|../common/values.json#/valueOptions| |
|underVoltageLockoutHysteresis|Undervoltage Lockout (UVLO) Hysteresis|../common/values.json#/valueOptions| |

####  4.5.6	 Specification To Capture The Current Consumption Data Of a Component

Source: [currentConsumption.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/currentConsumption.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|supplyName|name of the power supply |String| |
|quiescentCurrent|quiescent current (Iq) of a device|../common/values.json#/valueOptions| |
|shutdownCurrent|shutdown current (Isd) of a device|../common/values.json#/valueOptions| |
|activeCurrent|current consumption when a device is in active mode|../common/values.json#/valueOptions| |
|sleepCurrent|current consumption when a device is in sleep mode|../common/values.json#/valueOptions| |
|idleCurrent|current consumption when a device is in idle mode|../common/values.json#/valueOptions| |

####  4.5.7	 Specification For Referencing An External File

Source: [externalFile.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/externalFile.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|fileDescription|text description of the contents of an external file|String| |
|standardFileType|type of file being linked|String| |
|otherFileType|type of file being linked, if not in the standard list|String| |
|fileExtension|extension of file linked|String| |
|companionSoftware|optional, name of software program used to access file|String| |
|standardReferenced|optional, name of the standard the file is written in|String| |
|fileURI|URI linking to the CAD file|String| |

####  4.5.8	 Specification Of a Graph

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

####  4.5.9	 Curve

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|label|description of the data captured by a curve|String| |
|xData|x-axis values of the curve|array of Number| |
|yData|y-axis values of the curve|array of Number| |

####  4.5.10	 Specification Of An Instance In a Part

Source: [instanceSpec.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/instanceSpec.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partName|name of the part defined in the digital datasheets specifications|String|Yes|
|instanceName|name of an instance of the part|String| |
|instanceProperties|instance properties, as defined in the part specification|../common/coreProperties.json#/coreProperties| |

####  4.5.11	 Specification Of a Package

Source: [package.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/package.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|length|length of a side of a package|../common/values.json#/valueOptions| |
|width|width of a side of a package|../common/values.json#/valueOptions| |
|height|height of a package|../common/values.json#/valueOptions| |
|standardPackageSize|name of standard package size (imperial)|String| |
|standardPackageType|name of standard package types|String| |

####  4.5.12	 Specification Of Pinpaths Through An IC

Source: [pinPaths.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/pinPaths.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numberOfPinPaths|number of pinPaths defined. This number should not be higher than the number of components included in the part|Number| |
|partPinPaths|list of pins associated with each component in a multi-component part|array of #/$defs/partPinPaths|Yes|

####  4.5.13	 PartPinPaths

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentName|name of a component within a part |String| |
|componentPinNames|names of pins associated with each of the component in the part. Pin names must match the name used in part pin definition|array of String| |

####  4.5.14	 Specification Of Pins

Source: [pinSpec.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/pinSpec.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|terminalIdentifier|pin or ball number(s) as defined by datasheet|array of String|Yes|
|name|name given to the signal appearing at the terminal of a component|String|Yes|
|standardizedName|standard name of pin|String| |
|description|description of the signal appearing at the terminal of an electric/electronic component|String| |
|numberOfSupportedFunctions|the total number of functions supported by this pin|Number| |
|pinProperties|list of properties for each pin|#/$defs/pinProperties| |
|functionProperties|list of properties for each pin function configuration|array of #/$defs/functionProperties| |
|pinPaths|information on pin paths - pins associated with each component in a multi-component part (such as A1,Y1 and A2,Y2)|../common/pinPaths.json#/pinPaths| |

####  4.5.15	 ExternalComponents

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|componentType|type of external component required to be connected to a pin|String|Yes|
|configuration|electrical configuration of component connected to pin with respect to the pin|String|Yes|
|value|value of component|../common/values.json#/valueOptions| |

####  4.5.16	 FunctionProperties

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|perFunctionName|name of the function of a pin|String| |
|perFunctionProperties|list of pin properties that change based on the pin function configuration|#/$defs/pinProperties| |

####  4.5.17	 PinProperties

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|vih|the high-level input voltage for which operation of the logic element within specification limits is to be expected|../common/values.json#/valueOptions| |
|vil|the low-level input voltage for which operation of the logic element within specification limits is to be expected|../common/values.json#/valueOptions| |
|vol|the voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a low level at the output|../common/values.json#/valueOptions| |
|voh|the voltage level at an output terminal with input conditions applied that, according to the product specification, will establish a high level at the output|../common/values.json#/valueOptions| |
|absVmax|maximum voltage rating beyond which damage to the device may occur|../common/values.json#/valueOptions| |
|absVmin|absolute minimum voltage that can be applied to a pin|../common/values.json#/valueOptions| |
|voltageOperatingRange|voltage operating range that can safely be applied to a pin|../common/values.json#/valueOptions| |
|currentRange|current range that can safely be drawn/injected from/to a pin|../common/values.json#/valueOptions| |
|inputLeakage|current draw out of a high impedance input pin|../common/values.json#/valueOptions| |
|outputLeakage|current flow from a pin during the off state|../common/values.json#/valueOptions| |
|dcResistance|resistance of a pin of a connector|../common/values.json#/valueOptions| |
|interfaceType|type of interface enabled by pin|String| |
|pinUsage|standardized usage of pin|String| |
|direction|direction of a pin|String| |
|electricalConfiguration|electrical configuration of a pin|String| |
|polarity|whether the active state of a pin is high or low|String| |
|voltageOptions|list of voltage levels supported by a pin|Must set either Ref or Type| |
|floatUnused|description of whether pin can safely be floated if it is not used|Boolean| |
|internalPullUp|indicates the value of an internal pull-up on a pin|../common/values.json#/valueOptions| |
|internalPullDown|indicates the value of an internal pull-down on a pin|../common/values.json#/valueOptions| |
|esd|indicates whether ESD protection exists on a pin|Boolean| |
|externalComponents|list of external component structures recommended to be attached to a pin|array of #/$defs/externalComponents| |

####  4.5.18	 Specification Of Power Fets Properties

Source: [powerFetProperties.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/powerFetProperties.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|singlePowerFetPair|single pair of power fets (typical in buck or boost configuration)|#/$defs/powerFetPair| |
|inputPowerFetPair|input power fet pair (in a buck-boost configuration)|#/$defs/powerFetPair| |
|outputPowerFetPair|output power fet pair (in a buck-boost configuration)|#/$defs/powerFetPair| |

####  4.5.19	 PowerFetPair

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|ilimHSFET|maximum sustained output current under which the high side FET will operate properly|../common/values.json#/valueOptions| |
|ilimLSFET|maximum sustained output current under which the low side FET will operate properly|../common/values.json#/valueOptions| |
|rdsonHSFET|high side FET on-resistance|../common/values.json#/valueOptions| |
|rdsonLSFET|low side FET on-resistance|../common/values.json#/valueOptions| |

####  4.5.20	 Specification Of Power Sequencing Information

Source: [powerSequence.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/powerSequence.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|values|list of power sequence conditions that apply to a component|array of #/$defs/powerSequence| |

####  4.5.21	 PowerSequence

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|powerTransitionName|datasheet specific name given to a particular power transition for tracking|String| |
|powerTransitionDescription|description of power transition|String| |
|signal1|signal that comes up first in the sequence. Also used if there is only one relevant signal for a power transition (such as rise time) |String| |
|signal1TerminalIdentifiers|list of component pins associated with signal 1|array of String| |
|signal2|signal that comes up second in the sequence |String| |
|signal2TerminalIdentifiers|list of component pins associated with signal 2|array of String| |
|timeCondition|time between the signal 1 and signal 2 events|../common/values.json#/valueOptions| |
|powerDirection|whether this is a power up or power down event|Must set either Ref or Type| |
|nextPowerState|power state the device enters after the signal transition|String| |
|currentPowerState|power state the device is currently in before the signal transition|String| |
|riseTime|time for a signal to go from low to high (only applies to one signal)|../common/values.json#/valueOptions| |
|slewRate|maximum rate at which a voltage rail changes per time (only applies to one signal)|../common/values.json#/valueOptions| |
|fallTime|time for a signal to go from high to low (only applies to one signal)|../common/values.json#/valueOptions| |
|transitionStartCondition|the percentage of max voltage where rise or fall time starts being measured for timing (only applies to one signal)|../common/values.json#/valueOptions| |
|transitionEndCondition|the percentage of max voltage where rise or fall time stops being measured for timing (only applies to one signal)|../common/values.json#/valueOptions| |

####  4.5.22	 Specifications Of a Ratio

Source: [ratio.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/ratio.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|numerator|numerator of ratio|Number| |
|denominator|denominator of ratio|Number| |

####  4.5.23	 Specification Of a Register

Source: [register.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/register.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|registerName|name of a register|String|Yes|
|registerLongName|full Name of a register|String| |
|registerAddressOffset|address of a register|String|Yes|
|registerSize|size of a register|../common/values.json#/valueOptions|Yes|
|registerType|type of a register|String| |
|registerEndianness|memory storage order for the bytes|String| |
|registerResetValue|reset value of a register|String| |
|registerValue|value of a register|String| |
|registerIpName|name of the IP or interface controlled by changes in the register|String| |
|registerAccessType|access type of a Register|String| |
|registerBitField|describes the bit fields in the register|#/$defs/registerBitField| |

####  4.5.24	 RegisterBitField

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|bitFieldName|Name of the bit field|String| |
|bitFieldLongName|Long Name of the bit field|String| |
|bitFieldDescription|Describes the bit field|String| |
|bitFieldNumber|Number of a bitfield|Number| |
|bitFieldRange|Range of the bit field|String| |
|bitFieldResetValue|Reset value of a bit field|String| |
|bitFieldAccessType|Access type of a bit field|String| |

####  4.5.25	 Specification Of The Reliability Of a Component

Source: [reliability.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/reliability.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|failuresInTime|the number of expected failures per one billion hours of operation for a device|../common/values.json#/valueOptions| |
|meanTimeToFail|the average time a product or system functions before its first failure under normal conditions|../common/values.json#/valueOptions| |

####  4.5.26	 Specification Of Thermal Properties On Components

Source: [thermal.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/thermal.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|junctionTemperature|recommended operating junction temperature range of a part|../common/values.json#/valueOptions| |
|junctionTemperatureAbsMax|absolute maximum junction temperature of a part|../common/values.json#/valueOptions| |
|ambientTemperature|recommended operating ambient air temperature range of a part|../common/values.json#/valueOptions| |
|ambientTemperatureAbsMax|absolute maximum ambient air temperature of a part|../common/values.json#/valueOptions| |
|caseTemperature|case temperature range of a part|../common/values.json#/valueOptions| |
|caseTemperatureAbsMax|absolute maximum case temperature of a part|../common/values.json#/valueOptions| |
|leadTemperatureAbsMax|absolute maximum lead temperature of a part|../common/values.json#/valueOptions| |
|storageTemperatureAbsMax|absolute maximum storage temperature of a part|../common/values.json#/valueOptions| |
|packageThermalResistance|package thermal resistance of a part|../common/values.json#/valueOptions| |
|thermalResistanceJunctionToAmbient|thermal resistance between the semiconductor junction of a part and the ambient air|../common/values.json#/valueOptions| |
|thermalResistanceJunctionToCase|thermal resistance between the semiconductor junction and the package surface of a part|../common/values.json#/valueOptions| |
|thermalResistanceJunctionToBoard|thermal resistance between the semiconduction junction of a part and the board|../common/values.json#/valueOptions| |
|thermalResistanceJunctionToLead|thermal resistance between the semiconduction junction of a part and the solder pad|../common/values.json#/valueOptions| |
|thermalResistanceCaseToAmbient|thermal resistance between the package surface of a part and the ambient air|../common/values.json#/valueOptions| |
|thermalDesignPower|power at which thermal compliance should be evaluated at steady-state|../common/values.json#/valueOptions| |
|peakPower|maximum transient power the part will dissipate|../common/values.json#/valueOptions| |

####  4.5.27	 Specification Of Value

Source: [values.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/common/values.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|values|list of values that reflect possible values for a property|array of #/$defs/value| |

####  4.5.28	 Value

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|siUnit|name of SI unit of measure|String| |
|unitName|name of unit if not defined in the siUnit list|String| |
|typValue|typical unit quantity corresponding to unit text - example 40mV would have a value of 40|Number| |
|minValue|minimum unit quantity corresponding to unit text - example: 30mV would have a value of 30|Number| |
|maxValue|maximum unit quantity corresponding to unit text - example: 50mV would have a value of 50|Number| |
|unitFactor|multiplier on the value to achieve the SI unit listed - for example if millivolt was selected, 40mV would have a unitFactor value of 1; if volt was selected, the unitFactor value would be 0.001|Number| |
|relativeValueReference|if unit quantity is based on another reference, value of the reference|String| |
|relativeValueModifier|if a unit quantity is based on another reference, the value that edits that reference|Number| |
|relativeValueOperator|if a unit quantity is based on another reference, the operation that is performed with the modifier|String| |
|conditions|conditions under which the property is measured|array of String| |

### 4.6	 Clock

####  4.6.1	 Specification Of Clock

Source: [clock.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/clock/clock.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|fixedFrequency|clock frequency value if the clock has a fixed frequency|../common/values.json#/valueOptions| |
|numberClockOutputs|number of clock outputs in a clock IC|Number| |
|differentialSingleEnded|property describing whether a clock output is single ended or differential|String| |
|jitter|cycle to cycle clock jitter|../common/values.json#/valueOptions| |
|frequencyTolerance|amount of frequency variation from nominal frequency|../common/values.json#/valueOptions| |
|powerSupplyRejectionRatio|power supply rejection ratio (PSRR) or ratio between power supply variation and output variation|Number| |
|outputFormat|signal format of clock output|String| |
|currentConsumption|current used by device in various power modes|../common/currentConsumption.json#/currentConsumption| |

####  4.6.2	 Specification Of Oscillator

Source: [oscillator.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/clock/oscillator.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|baseResonator|technology producing resonance|String| |
|frequency|output frequency of oscillator|../common/values.json#/valueOptions|Yes|
|frequencyStability|frequency change over temperature, load, supply voltage change and aging|Number| |
|currentConsumption|current consumption of a device|../common/currentConsumption.json#/currentConsumption| |
|outputLoad|maximum capacitive load that can be supported by oscillator|../common/values.json#/valueOptions| |
|riseTime|time for output to go from 10% to 90% of output max|../common/values.json#/valueOptions| |
|fallTime|time for output to go from 90% to 10% of output max|../common/values.json#/valueOptions| |
|startUpTime|time between enable and output reaching 10% of output max|../common/values.json#/valueOptions| |
|dutyCycle|time above 50% of output max over entire period|../common/values.json#/valueOptions| |
|phaseJitter|variation of waveform period|../common/values.json#/valueOptions| |

### 4.7	 Data_converter

####  4.7.1	 Specification Of ADC

Source: [adc.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/data_converter/adc.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|digitalResolution|number of bits of resolution in the digital output|Number| |
|conversionTime|time required to convert from an analog signal to digital output|../common/values.json#/valueOptions| |
|sampleRate|maximum rate at which the ADC can convert samples|../common/values.json#/valueOptions| |
|offsetError|difference (in LSB) of the output at the zero point between an actual and ideal ADC|Number| |
|gainError|difference (in LSB) of how the actual transfer function matches the ideal transfer function, also called full scale error|Number| |
|integralNonlinearity|deviation (in LSB) of an actual transfer function from an ideal transfer function|Number| |
|differentialNonlinearity|difference (in LSB) in step width between the actual and ideal transfer functions|Number| |
|rmsNoise|root mean square (RMS) noise of ADC|Number| |
|SNR|signal to noise (SNR) ratio of the converter|Number| |
|interface|digital communication interfaces supported|array of String| |
|inputType|whether the ADC has a single ended or differential input|String| |
|inputChannels|number of input channels|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.7.2	 Specification Of DAC

Source: [dac.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/data_converter/dac.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|digitalResolution|number of bits of resolution|Number| |
|offsetError|analog response to an input code of all zeros|../common/values.json#/valueOptions| |
|gainError|difference (in percentage of FSR) of how well the slope of the actual transfer function matches the ideal transfer function|../common/values.json#/valueOptions| |
|integralNonlinearity|deviation of an actual transfer function from an ideal transfer function, in LSB|Number| |
|differentialNonlinearity|difference between the ideal and the actual output responses for successive DAC codes, in LSB|Number| |
|settlingTime|time from application of input code to valid output response|../common/values.json#/valueOptions| |
|sampleRate|maximum rate at which the DAC can convert samples|../common/values.json#/valueOptions| |
|interface|digital communication interfaces supported|array of String| |
|outputType|whether the DAC has a single ended or differential output|String| |
|outputChannelCount|number of output channels|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.8	 Hardware

####  4.8.1	 Specification Of Connector

Source: [connector.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/hardware/connector.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|function|intended function of a connector|String| |
|contactCount|number of contacts in a connector|Number|Yes|
|type|property describing the method of mating to the connector|String| |
|cycleRating|number of plug/unplug cycles a connector is rated to support|Number| |
|pitch|distance from the center of one contact on the connector to the center of the next contact|../common/values.json#/valueOptions| |
|keying|property describing whether a connector has an asymmetry to prevent it from being plugged in the wrong direction|Boolean| |

####  4.8.2	 Specification Of Switch

Source: [switch.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/hardware/switch.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|type|property describing the way in which the switch is activated|String| |
|contactType|property describing the order in which switch contact is made and broken|String| |
|circuitConfig|property describing the number of poles and throws in a switch|String| |
|cycleRating|number of on/off cycles a mechanical switch can reliably sustain|Number| |
|voltageRating|maximum DC voltage potential that can be applied across an open switch|../common/values.json#/valueOptions| |
|currentRating|maximum DC current that can flow through a closed switch without causing excessive heating|../common/values.json#/valueOptions| |
|onResistance|nominal resistance of a closed switch|../common/values.json#/valueOptions|Yes|
|dielectricRating|maximum AC voltage potential that can be applied across an open switch for one minute|../common/values.json#/valueOptions| |

### 4.9	 Ic_io

####  4.9.1	 Specification Of Bridge Chip

Source: [bridge_chip.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/bridge_chip.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|muxRatio|ratio of inputs to outputs|../common/ratio.json#/ratio| |
|inputInterfaces|list of interfaces at the input of the bridge|array of String| |
|outputInterfaces|list of interfaces at the output of the bridge|array of String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.9.2	 Specification Of High Speed Mux

Source: [highspeed_mux.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/highspeed_mux.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|muxRatio|ratio of inputs to outputs|../common/ratio.json#/ratio| |
|inputInterfaces|list of interfaces high speed mux is designed for|array of String| |
|maxDataRate|maximum data rate supported by high speed mux|../common/values.json#/valueOptions| |
|insertionLoss|insertion loss through high speed mux|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.9.3	 Specification Of Level Shifter

Source: [level_shifter.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/level_shifter.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|inputVoltage|nominal input voltage of level shifter|../common/values.json#/valueOptions| |
|outputVoltage|nominal output voltage of level shifter|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.9.4	 Specification Of Redriver

Source: [redriver.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/redriver.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|numberChannels|number of lanes (single ended or differential) supported by redriver|Number| |
|interface|list of interface(s) supported by redriver|array of String| |
|maxDataRate|maximum data rate supported by redriver|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.9.5	 Specification Of Retimer

Source: [retimer.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/retimer.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|numberOfLanes|number of lanes (single ended or differential) supported by a device|Number| |
|interface|list of interface(s) supported by a device|array of String| |
|dpMaxLinkRate|maximum interface speed supported by DP interface of a device|String| |
|maxDataRate|maximum data rate supported by a device|../common/values.json#/valueOptions| |
|integratedAuxLsSwitch|whether the AUX/LSx switch for SBU is integrated|Boolean| |
|commonClock|whether a device supports common reference clock|Boolean| |
|sris|whether a device supports Separate Reference clock with Independent Spread spectrum clocking(SRIS)|Boolean| |
|srns|whether a device supports Separate Reference clock with No Spread spectrum clocking (SRNS)|Boolean| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.9.6	 Specification Of USB Battery Charging 1.2 (bc12) Detector

Source: [usb_bc12.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/usb_bc12.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|hostMode|whether host mode is supported by bc12 chip|Boolean|Yes|
|deviceMode|whether device mode is supported by bc12 chip|Boolean|Yes|
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.9.7	 Specification Of USB-C PD Controller

Source: [usbc_pdcontroller.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/usbc_pdcontroller.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|pdVersion|version of power delivery spec implemented by controller|String|Yes|
|usbTypecRevision|usb type-c spec revision implemented by controller|String| |
|powerRoleSupported|roles supported by pd controller|String| |
|fastRoleSwapSupport|whether the pd controller supports fast role swap (FRS)|Boolean| |
|vconnPowerSupport|whether the pd controller has support for vconn power|Boolean| |
|vconnPowerLimit|power limit supported by internal vconn switch (if supported)|../common/values.json#/valueOptions| |
|vconnMaxCurrent|maximum continuous current supported by internal vconn switch (if supported)|../common/values.json#/valueOptions| |
|vconnOverCurrentLimit|over current limit supported by internal vconn switch (if supported)|../common/values.json#/valueOptions| |
|integratedVbusDischargeSwitch|whether the pd controller has one or more integrated vbus discharge switches |Boolean| |
|integratedLoadSwitch|whether the pd controller has an integrated load switch for vbus power |Boolean| |
|maxSinkCurrent|maximum continuous current supported by pd controller integrated sink load switch|../common/values.json#/valueOptions| |
|maxSourceCurrent|maximum continuous current supported by pd controller integrated source load switch|../common/values.json#/valueOptions| |
|sinkfetOverCurrentLimit|over current limit supported by pd controller integrated sink load switch|../common/values.json#/valueOptions| |
|sourcefetOverCurrentLimit|over current limit supported by pd controller integrated source load switch|../common/values.json#/valueOptions| |
|onResistanceSinkFet|on-resistance of the integrated sink load switch|../common/values.json#/valueOptions| |
|onResistanceSourceFet|on-resistance of the integrated source load switch|../common/values.json#/valueOptions| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  4.9.8	 Specification Of USB-C Port Controller

Source: [usbc_portcontroller.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_io/usbc_portcontroller.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|pdVersion|version of power delivery spec implemented by controller|String|Yes|
|usbTypecRevision|usb type-c spec revision implemented by controller|String| |
|usbTcpiRevision|usb type-c port controller interface spec revision implemented by controller|String| |
|powerRoleSupported|roles supported by pd controller|String| |
|fastRoleSwapSupport|whether the port controller supports fast role swap (FRS)|Boolean| |
|vconnPowerSupport|whether the port controller has support for vconn power|Boolean| |
|vconnFetOnResistance|on-resistance of the integrated Vconn FET|..common/values.json#/valueOptions| |
|vconnPowerLimit|power limit supported by internal vconn switch (if supported)|../common/values.json#/valueOptions| |
|vconnMaxCurrent|maximum continuous current supported by internal vconn switch (if supported)|../common/values.json#/valueOptions| |
|vconnOverCurrentLimit|over current limit supported by internal vconn switch (if supported)|../common/values.json#/valueOptions| |
|integratedVbusDischargeSwitch|whether the port controller has one or more integrated vbus discharge switches |Boolean| |
|integratedMux|whether the port controller has one or more integrated high speed muxes|Boolean| |
|interface|describes the communication interface from the controller to the policy manager|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

### 4.10	 Ic_microcontroller

####  4.10.1	 Specification Of Microcontroller/ec

Source: [microcontroller.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_microcontroller/microcontroller.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|onChipFlash|capacity of built-in flash in a microprocessor|../common/values.json#/valueOptions| |
|onChipRAM|capacity of built-in RAM in a microprocessor|../common/values.json#/valueOptions| |
|onChipROM|capacity of built-in ROM in a microprocessor|../common/values.json#/valueOptions| |
|coreProcessor|description of core processor|String| |
|coreArchitectureBits|number of bits of data a CPU can transfer per clock cycle|String| |
|clockSpeed|speed of main CPU clock|../common/values.json#/valueOptions| |
|firmwareVersion|firmware version of the part|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.11	 Ic_misc

####  4.11.1	 Specification Of Audio Codec

Source: [audio_codec.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/audio_codec.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|dataLength|number of bits in a data word|Number| |
|hpOutputSNR|headphone amplifier output SNR|Number| |
|hpOutputTHD+N|headphone output total harmonic distortion plus noise|../common/values.json#/valueOptions| |
|micInputSNR|microphone input SNR|Number| |
|micInputTHD+N|microphone input total harmonic distortion plus noise|../common/values.json#/valueOptions| |
|jackDetect|describes whether headphone jack detection is supported|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.11.2	 Specification Of Speaker Amplifier

Source: [speaker_amplifier.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/speaker_amplifier.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|dataLength|number of bits in a data word|Number| |
|outputPower|typical output power from speaker amplifier|../common/values.json#/valueOptions| |
|efficiency|typical speaker amplifier efficiency|../common/values.json#/valueOptions| |
|thd+n|typical total harmonic distortion plus noise of amplifier|../common/values.json#/valueOptions| |
|sampleRate|sample rate of data out of amplifier|../common/values.json#/valueOptions| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.11.3	 Specification Of TPM

Source: [tpm.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/tpm.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|interface|describes the communication interface from the chip to the host|array of String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.11.4	 Specification Of WLAN Module

Source: [wlan_module.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/wlan_module.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|wlanSpec|version of wlan specification supported by module|String| |
|bluetoothVersion|version of bluetooth supported by module|String| |
|txrxChains|number of tx and rx chains in a wifi module|String| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|lteCoexFilter|describes whether module supports lte coexistence filtering|Boolean| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.11.5	 Specification Of WWAN Module

Source: [wwan_module.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/ic_misc/wwan_module.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|networkSupport|networks supported by wwan module|String| |
|gpsSupport|whether wwan module has gps support|Boolean| |
|m2FormFactor|wlan module form factor described by jedec standard m.2 form factors|String| |
|keying|pcie card key|String| |
|interface|describes the communication interface from the chip to the host|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.12	 Logic

####  4.12.1	 Specification Of Logic Gate

Source: [logic_gate.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/logic/logic_gate.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|type|logical operation performed by logic gate|String|Yes|
|numberGates|number of logical gates encapsulated in logic IC|Number| |
|schmittTrigger|property describing whether logic gate has schmitt trigger inputs|Boolean| |
|propagationDelay|time between input changing to output changing|../common/values.json#/valueOptions| |
|rampTime|time for output to go from 10% nominal output voltage to 90% nominal output voltage|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.13	 Memory

####  4.13.1	 Specification Of DRAM

Source: [dram.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/dram.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|type|type of dram|String| |
|capacity|capacity of dram chip|../common/values.json#/valueOptions|Yes|
|ranksperModule|numbers of ranks on dram module|Number| |
|diesPerChip|number of dies on dram chip|Number| |
|channelsPerDie|number of channels per die on dram chip|Number| |
|banksPerChannel|number of banks per channel of dram|Number| |
|bitsPerChannel|channel density of dram|../common/values.json#/valueOptions| |
|bitsPerDie|total die density of dram|../common/values.json#/valueOptions| |
|pageSize|page size of dram|../common/values.json#/valueOptions| |
|rows|number of rows per channel of dram|Number| |
|columns|number of columns per row of dram|Number| |
|dataRate|dram maximum data rate|Number| |
|speed|dram maximum speed|../common/values.json#/valueOptions| |
|latencyCas|cl/tCAS, delay between read command issued and first output data available for read|Number| |
|delayRasCas|tRCD,delay between activation of row and activation of column where data is stored in the dram|../common/values.json#/valueOptions| |
|delayRasPrecharge|tRP, delay between closing access to a row through the precharge command and activating a new row to access data |../common/values.json#/valueOptions| |
|delayActivePrecharge|tRAS, delay between row active command issued and precharge command issued |../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.13.2	 Specification Of EEPROM

Source: [eeprom.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/eeprom.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|accessTime|time to access the eeprom|../common/values.json#/valueOptions| |
|bitsPerWords|number of columns in the eeprom|Number| |
|bootBlockSize|size of the eeprom boot block|../common/values.json#/valueOptions| |
|capacity|capacity/density of eeprom|../common/values.json#/valueOptions|Yes|
|clockFrequency|eeprom clock frequency|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|dataRetention|maximum number of read and write cycle the part can support|Number| |
|endurance|time in years a bit in the eeprom can retain its data state |Number| |
|interface|interface of eeprom to host|String| |
|numberOfWords|number of rows in the eeprom|Number| |

####  4.13.3	 Specification Of Flash Memory

Source: [flash_memory.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/flash_memory.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|capacity|capacity/density of flash memory|../common/values.json#/valueOptions|Yes|
|pageSize|page size of flash memory|../common/values.json#/valueOptions| |
|blockSize|block size of flash memory|../common/values.json#/valueOptions| |
|bootBlockSize|size of the flash memory boot block|../common/values.json#/valueOptions| |
|interface|interface of flash memory to host|String| |
|clockFrequency|flash memory clock frequency|../common/values.json#/valueOptions| |
|blockEraseTime|time it takes to erase a block (largest erasable unit) of the flash memory|../common/values.json#/valueOptions| |
|sectorEraseTime|time it takes to erase a sector (smallest erasable unit) of the flash memory|../common/values.json#/valueOptions| |
|chipEraseTime|time it takes to erase the flash memory|../common/values.json#/valueOptions| |
|pageProgramTime|time it takes to program a page of the flash memory|../common/values.json#/valueOptions| |
|endurance|time in years a bit in the eeprom can retain its data state |Number| |
|dataRetention|maximum number of read and write cycle the part can support|Number| |
|hwReset|whether the part supports a hardware reset pin|Boolean| |
|writeProtect|whether the part has a write protect pin|Boolean| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.13.4	 Specification Of ROM

Source: [rom.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/memory/rom.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|capacity|capacity of rom|../common/values.json#/valueOptions| |
|interface|interface of rom to host|String| |
|qeStatus|indicates whether the Quad Enable(QE) bit is set|Boolean| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.14	 Passives

####  4.14.1	 Specification Of Capacitor

Source: [capacitor.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/capacitor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|value|capacitor value|../common/values.json#/valueOptions|Yes|
|tolerance|nominal tolerance of a capacitor|../common/values.json#/valueOptions| |
|ratedVoltage|maximum voltage which may be applied continuously to a capacitance|../common/values.json#/valueOptions| |
|dielectric|dielectric material used in the capacitor|String| |
|polarized|describes whether the capacitor is polarized|Boolean| |
|equivalentSeriesResistance|equivalent series resistance (ESR) of the capacitor|../common/values.json#/valueOptions| |
|temperatureCoefficient|change in capacitance when the temperature is changed|Number| |
|minTemperature|minimum temperature under which a capacitor can be expected to reliably operate|../common/values.json#/valueOptions| |
|maxTemperature|maximum temperature under which a capacitor can be expected to reliably operate|../common/values.json#/valueOptions| |
|capacitorDerating|graph object to capture capacitance changes with voltage changes|../common/graph.json#/graph| |

####  4.14.2	 Specification Of Common Mode Choke

Source: [common_mode_choke.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/common_mode_choke.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|diffModeCutoff|frequency at which the differential mode attenuation equals -3dB|../common/values.json#/valueOptions| |
|commonModeAttenuation|graph object to capture common mode attenuation of a common mode choke at various frequencies|../common/graph.json#/graph| |
|dcResistance|dc resistance (DCR) of a common mode choke|../common/values.json#/valueOptions| |
|rmsCurrent|applied DC current (IRMS) that produces a common mode choke temperature rise of 40 deg C|../common/values.json#/valueOptions| |
|intendedApplication|intended application of a particular common mode choke|String| |

####  4.14.3	 Specification Of Ferrite Bead

Source: [ferrite_bead.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/ferrite_bead.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|dcResistance|dc resistance (DCR) of ferrite bead|../common/values.json#/valueOptions| |
|rmsCurrent|applied DC current (IRMS) that produces a ferrite bead temperature rise of 40 deg C|../common/values.json#/valueOptions| |
|impedance100MHz|impedance of ferrite bead under standard testing conditions at 100MHz|../common/values.json#/valueOptions| |
|impedanceTolerance|variation of ferrite bead impedance expressed as +/- percentage|../common/values.json#/valueOptions| |

####  4.14.4	 Specification Of Inductor

Source: [inductor.json](https:/github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/inductor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|value|inductor value|../common/values.json#/valueOptions|Yes|
|tolerance|nominal tolerance of an inductor|../common/values.json#/valueOptions| |
|ratedCurrent|maximum continuous current the inductor can handle|../common/values.json#/valueOptions| |
|saturationCurrent|current where the inductor enters the magnetic state, and the inductance drops a specified amount|../common/values.json#/valueOptions| |
|rmsCurrent|DC current that produces an inductor temperature rise of 40 degrees Celsius|../common/values.json#/valueOptions| |
|selfResonantFrequency|frequency at which the inductor starts behaving like a capacitor|../common/values.json#/valueOptions| |
|dcResistance|DC resistance of the inductor|../common/values.json#/valueOptions| |
|temperatureCoefficient|change in inductance when the temperature is changed|Number| |
|minTemperature|minimum temperature under which a inductor can be expected to reliably operate|../common/values.json#/valueOptions| |
|maxTemperature|maximum temperature under which a inductor can be expected to reliably operate|../common/values.json#/valueOptions| |
|saturationCurve|graph object to capture inductance as a function of current|../common/graph.json#/graph| |
|resonantFrequencyCurve|graph object to capture inductance as a function of frequency|../common/graph.json#/graph| |

####  4.14.5	 Specification Of Resistor

Source: [resistor.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/passives/resistor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|value|resistance value|../common/values.json#/valueOptions|Yes|
|tolerance|nominal tolerance of a resistor|../common/values.json#/valueOptions| |
|powerRating|measure of power a resistor can dissipate indefinitely without degrading performance|../common/values.json#/valueOptions| |
|temperatureCoefficient|change in resistance when the temperature is changed|Number| |
|maxOverloadVoltage|maximum voltage that can be applied to the resistor for a short period of time|../common/values.json#/valueOptions| |
|maxLimitingElementVoltage|maximum voltage value that can be applied continuously to the resistor|../common/values.json#/valueOptions| |
|minTemperature|minimum temperature under which a resistor can be expected to reliably operate|../common/values.json#/valueOptions| |
|maxTemperature|maximum temperature under which a resistor can be expected to reliably operate|../common/values.json#/valueOptions| |
|resistorDerating|graph object to capture resistance changes with temperature|../common/graph.json#/graph| |

### 4.15	 Power

####  4.15.1	 Specification Of Battery Charger

Source: [battery_charger.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/battery_charger.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|chargerType|battery charger type|String| |
|converterType|switching charger type|String| |
|chargerTopology|type of battery charger topology (Narrow VDC vs Hybrid Power Boost)|String| |
|batteryConfig|battery configuration supported by the device|array of String| |
|batteryCellChemistry|battery cell chemistry supported by the device|array of String| |
|inputPowerSource|input power source supported by the device|array of String| |
|inputCurrentAccuracy|accuracy of input current when set|../common/values.json#/valueOptions| |
|batteryChargeCurrent|charging (fast charge) current of a device|../common/values.json#/valueOptions|Yes|
|batteryChargeCurrentAccuracy|charging current regulation accuracy of a device|../common/values.json#/valueOptions| |
|batteryPreChargeCurrent|charging current of a device in pre-charge phase|../common/values.json#/valueOptions| |
|batteryPreChargeCurrentAccuracy|pre-charging current regulation accuracy of a device|../common/values.json#/valueOptions| |
|batteryTrickleChargeCurrent|charging current of a device in trickle charge phase|../common/values.json#/valueOptions| |
|batteryTrickleChargeCurrentAccuracy|trickle charging current regulation accuracy of a device|../common/values.json#/valueOptions| |
|batteryTerminationChargeCurrent|charging current of a device in charge termination phase|../common/values.json#/valueOptions| |
|preChargeToChargeThreshold|the per-cell voltage at which a charger transitions from the pre-charge to fast charge current|../common/values.json#/valueOptions| |
|trickleChargeTopreChargeThreshold|the per-cell voltage at which a charger transitions from the trickle charge to pre-charge current|../common/values.json#/valueOptions| |
|minVinRegulation|minimum VIN required to support the max input charger current|../common/values.json#/valueOptions| |
|chargerInputCurrentLimit|the maximum current drawn from the input|../common/values.json#/valueOptions| |
|batteryTerminationChargeCurrentAccuracy|termination charging current regulation accuracy of a device|../common/values.json#/valueOptions| |
|batteryChargeVoltage|battery charge voltage regulated by a device|../common/values.json#/valueOptions| |
|batteryChargeVoltageAccuracy|accuracy of battery charge voltage regulated by a device|../common/values.json#/valueOptions| |
|efficiency|charger efficiency as a function of charge current of a device|../common/graph.json#/graph| |
|vin|input voltage under which a device can be expected to reliably operate|../common/values.json#/valueOptions| |
|fsw|switching frequency of a device|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |
|batteryChargerProtections|battery charger specific protections supported by device|array of String| |
|integratedLoadSwitch|whether the device contains integrated power path load switch(es)|Boolean| |
|integratedFets|whether the device contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|../common/powerFetProperties.json#/powerFetProperties| |
|gateCapacitance|describes gate capacitance supported on external fets|../common/values.json#/valueOptions| |
|inputSenseResistor|describes input sense resistor value|../common/values.json#/valueOptions| |
|batterySenseResistor|describes battery sense resistor value|../common/values.json#/valueOptions| |
|passThroughMode|whether pass through mode is supported|Boolean| |
|bc12Support|whether bc 1.2 detection is built in|Boolean| |
|tcpcSupport|whether type-C port controller support is built in|Boolean| |
|usbTypecRevision|usb type-c spec revision supported by a device|String| |
|pdVersion|version of power delivery spec supported by a device|String| |

####  4.15.2	 Specification Of Display Backlight Driver

Source: [displaybacklight_driver.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/displaybacklight_driver.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|vin|input voltage under which a device can be expected to operate properly|../common/values.json#/valueOptions|Yes|
|vout|output voltage range a device can regulate|../common/values.json#/valueOptions| |
|ioutPerString|output current per string a device can regulate|../common/values.json#/valueOptions|Yes|
|ioutAccuracy|accuracy of per string current regulated by a device|../common/values.json#/valueOptions| |
|fsw|switching frequency of a device|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|integratedFets|whether a device contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|../common/powerFetProperties.json#/powerFetProperties| |
|currentMatchingAccuracy|current matching between LED strings|../common/values.json#/valueOptions| |
|dimmingSupport|whether a device supports output current dimming|Boolean| |
|dimmingControl|whether a device is dimmed by PWM or I2C|String| |
|dimmingFrequency|dimming frequency of a device|../common/values.json#/valueOptions| |
|dimmingRatio|dimming ratio of a device|../common/ratio.json#/ratio| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |
|efficiency|backlight driver efficiency as a function of forward current|../common/graph.json#/graph| |

####  4.15.3	 Specification Of Linear Regulator

Source: [linear_regulator.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/linear_regulator.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|vin|minimum input voltage under which the part can be expected to operate without the output dropping|../common/values.json#/valueOptions|Yes|
|vout|output voltage the part can regulate|../common/values.json#/valueOptions|Yes|
|feedbackVoltage|voltage comparison point at the feedback node (vref)|../common/values.json#/valueOptions| |
|dropoutVoltage|dropout voltage of a device|../common/values.json#/valueOptions| |
|loadCurrent|load current supported by a device|../common/values.json#/valueOptions|Yes|
|currentLimit|sustained output current threshold beyond which the output of a device starts drooping|../common/values.json#/valueOptions| |
|voutAccuracy|output voltage variation at no load of a device|../common/values.json#/valueOptions| |
|loadRegulation|output voltage variation,from no load to full load, of a device |../common/values.json#/valueOptions| |
|lineRegulation|output voltage variation,from minimum input voltage to maximum input voltage, of a device |../common/values.json#/valueOptions| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|../common/values.json#/valueOptions| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|../common/values.json#/valueOptions| |
|powerSupplyRejectionRatio|graph object to capture Power Supply Rejection Ratio (PSRR) of device as a function of frequency|../common/graph.json#/graph| |
|rmsOutputNoise|graph object to capture RMS output noise of device as a function of frequency|../common/graph.json#/graph| |
|totalOutputNoise|total output noise of a device|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  4.15.4	 Specification Of Load Switch

Source: [load_switch.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/load_switch.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|fetType|type of pass FET in a device|String| |
|loadSwitchCount|number of load switches in the package|Number| |
|vin|input voltage under which a device can be expected to reliably operate|../common/values.json#/valueOptions|Yes|
|outputCurrent|continuous DC current supported by a device|../common/values.json#/valueOptions|Yes|
|onResistance|on state resistance of FET|../common/values.json#/valueOptions|Yes|
|pullDownResistance|pull-down resistance of a device from the output to the ground|../common/values.json#/valueOptions| |
|currentLimitSupport|whether a device supports current limiting|Boolean| |
|adjustableRiseTimeSupport|whether a device supports adjustable rise time|Boolean| |
|quickOutputDischargeSupport|whether a device supports quick output discharge|Boolean| |
|reverseCurrentBlockingSupport|whether a device supports reverse current blocking|Boolean| |
|powerGoodSupport|whether a device supports power good|Boolean| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|../common/values.json#/valueOptions| |
|offTime|time between enable deasserted and output voltage falling to 90% nominal|../common/values.json#/valueOptions| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|../common/values.json#/valueOptions| |
|fallTime|time for output voltage to go from 90% vout nominal to 10% vout nominal|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  4.15.5	 Specification Of a PMIC

Source: [pmic.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/pmic.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String| |
|instanceSpec|array of specifications of instances in the pmic |array of ../common/instanceSpec.json#/instanceSpec| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |

####  4.15.6	 Specification Of Switching Regulator

Source: [switching_regulator.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/power/switching_regulator.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|regulatorTopology|switching voltage regulator topology|String|Yes|
|vin|input voltage range under which a device can be expected to reliably operate|../common/values.json#/valueOptions|Yes|
|vout|output voltage a device can regulate|../common/values.json#/valueOptions|Yes|
|feedbackVoltage|voltage comparison point at the feedback node|../common/values.json#/valueOptions| |
|loadCurrent|load current supported by a device|../common/values.json#/valueOptions|Yes|
|voutAccuracy|output voltage variation at no load|../common/values.json#/valueOptions| |
|loadRegulation|output voltage variation from no load to full load |../common/values.json#/valueOptions| |
|lineRegulation|output voltage variation from minimum input voltage to maximum input voltage |../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |
|switchingFrequency|switching frequency (fsw) of voltage regulator|../common/values.json#/valueOptions| |
|enableTime|time between enable asserted and output voltage rising to 10% nominal|../common/values.json#/valueOptions| |
|rampTime|time for output voltage to go from 10% vout nominal to 90% vout nominal|../common/values.json#/valueOptions| |
|integratedFets|whether the regulator contains integrated switching mosfets|Boolean| |
|integratedFetProperties|describes integrated fet current limits and rdson properties|../common/powerFetProperties.json#/powerFetProperties| |
|componentProtectionThresholds|thermal and power supply protection thresholds of a device|../common/componentProtectionThresholds.json#/componentProtectionThresholds| |
|efficiency|power efficiency of a regulator as a function of output current|../common/graph.json#/graph| |

### 4.16	 Semiconductor

####  4.16.1	 Specification Of BJT

Source: [bjt.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/bjt.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|bjtChannelType|doping of a transistor's channel - describes whether a transistor is n-type or p-type|String|Yes|
|transistorCount|number of transistors in the package|Number| |
|collectorCurrent|maximum current flow of BJT as measured at the collector (Icc)|../common/values.json#/valueOptions| |
|peakCollectorCurrent|maximum pulsed current flow of BJT as measured at the collector (Icm)|../common/values.json#/valueOptions| |
|baseCurrent|maximum current flow of BJT as measured at the base (Ib)|../common/values.json#/valueOptions| |
|peakBaseCurrent|maximum pulsed current flow of BJT as measured at the base (Ibm)|../common/values.json#/valueOptions| |
|collectorBaseVoltage|maximum voltage between collector and base terminals of BJT with an open emitter terminal (V_CBO)|../common/values.json#/valueOptions| |
|collectorEmitterVoltage|maximum voltage between collector and emitter terminals of BJT with an open base terminal (V_CEO)|../common/values.json#/valueOptions| |
|emitterBaseVoltage|maximum voltage between emitter and base terminals of BJT with an open collector terminal (V_EBO)|../common/values.json#/valueOptions| |
|pTot|maximum power that can be continuously dissipated under temperature conditions|../common/values.json#/valueOptions| |
|collectorBaseCutOffCurrent|current into the collector terminal when the BJT's base and collector are reverse biased and the emitter is open (I_CBO)|../common/values.json#/valueOptions| |
|emitterBaseCutOffCurrent|current into the base terminal when the BJT's base and emitter are reverse biased and the collector is open (I_EBO)|../common/values.json#/valueOptions| |
|dcCurrentGain|ratio of collector current to base current (hfe)|../common/values.json#/valueOptions| |
|collectorEmitterSaturationVoltage|collector-emitter voltage below which a change in base current does not impact collector current (VCE_SAT)|../common/values.json#/valueOptions| |
|baseEmitterSaturationVoltage|base-emitter voltage required to ensure the collector is forward biased for certain current conditions (VBE_SAT)|../common/values.json#/valueOptions| |
|collectorEmitterBreakdownVoltage|collector-emitter voltage at which a specified current flows with the base open|../common/values.json#/valueOptions| |
|baseEmitterBreakdownVoltage|base-emitter voltage at which a specified current flows with the collector open|../common/values.json#/valueOptions| |
|delayTime|time delay between input signal rising and when collector current rises to 10% of Isat (td)|../common/values.json#/valueOptions| |
|riseTime|time for collector current to rise through active region from 10% to 90% of Isat (tr)|../common/values.json#/valueOptions| |
|storageTime|time delay between input signal falling and when collector current falls to 90% of Isat (ts)|../common/values.json#/valueOptions| |
|fallTime|time for collector current to fall from 90% to 10% of Isat (tf)|../common/values.json#/valueOptions| |
|collectorCapacitance|parasitic capacitance of collector terminal under certain conditions (Cc)|../common/values.json#/valueOptions| |
|emitterCapacitance|parasitic capacitance of emitter terminal under certain conditions (Ce)|../common/values.json#/valueOptions| |
|transitionFrequency|frequency of unity current gain with a short circuit output (ft)|../common/values.json#/valueOptions| |
|icVsHfe|graph of collector current (Ic) as a function of dc current gain (Hfe)|../common/graph.json#/graph| |
|icVsVce|graph of collector current (Ic) as a function of collector-emitter saturation voltage (VCE_SAT)|../common/graph.json#/graph| |
|ibVsVce|graph of base current (Ib) as a function of collector-emitter voltage (VCE)|../common/graph.json#/graph| |
|pdVsTemp|graph of power dissipation as a function of temperature|../common/graph.json#/graph| |

####  4.16.2	 Specification Of Diode

Source: [diode.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/diode.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|type|type of diode|String| |
|diodeCount|number of diodes in the package|Number| |
|diodeConfiguration|configuration of diode|String| |
|numberOfProtectedLines|number of lines a diode can protect|Number| |
|vf|forward voltage of a diode|../common/values.json#/valueOptions| |
|if|continuous forward current of a diode|../common/values.json#/valueOptions| |
|ifm|maximum continuous forward current a diode can support|../common/values.json#/valueOptions| |
|ifrm|maximum repetitive peak forward current a diode can support|../common/values.json#/valueOptions| |
|ifsm|maximum non-repetitive surge forward current a diode can support|../common/values.json#/valueOptions| |
|vbr|breakdown voltage of a diode|../common/values.json#/valueOptions| |
|ir|reverse current|../common/values.json#/valueOptions| |
|vz|breakdown voltage of a zener diode|../common/values.json#/valueOptions| |
|vrm|maximum reverse standoff voltage a tvs diode can withstand|../common/values.json#/valueOptions| |
|vcl|clamping voltage of a tvs diode|../common/values.json#/valueOptions| |
|vr|maximum continuous reverse biased voltage a diode can support|../common/values.json#/valueOptions| |
|vrrm|maximum repetitive reverse voltage pulses a diode can support|../common/values.json#/valueOptions| |
|cd|diode junction capacitance - between the anode and cathode- in reverse bias condition|../common/values.json#/valueOptions| |
|trr|reverse recovery time it takes the diode to stop conducting when its voltage changes from forward-bias to reverse-bias|../common/values.json#/valueOptions| |
|pTot|maximum power dissipation of a forward biased diode|../common/values.json#/valueOptions| |
|ifVsVf|graph of forward current (If) as a function of forward voltage (VfTyp)|../common/graph.json#/graph| |

####  4.16.3	 Specification Of LED

Source: [led.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/led.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|ledColor|LED color|String| |
|vf|forward voltage of an LED |../common/values.json#/valueOptions| |
|if|continuous forward current of an LED|../common/values.json#/valueOptions| |
|ifp|peak forward current of an LED|../common/values.json#/valueOptions| |
|vr|maximum continuous reverse biased voltage an LED can support|../common/values.json#/valueOptions| |
|ir|reverse current (leakage) of an LED|../common/values.json#/valueOptions| |
|ledCapacitance|capacitance of an LED|../common/values.json#/valueOptions| |
|iv|LED luminous intensity|../common/values.json#/valueOptions| |
|peakWavelength|light spectrum output value emitted by an LED at highest wavelength|../common/values.json#/valueOptions| |
|dominantWavelength|dominant wavelength an LED emits the majority of the time|../common/values.json#/valueOptions| |
|angleHalfIntensity|angle at which LED intensity falls to 50% of its maximum value|../common/values.json#/valueOptions| |
|pd|power dissipation of an LED|../common/values.json#/valueOptions| |

####  4.16.4	 Specification Of MOSFET

Source: [mosfet.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/semiconductor/mosfet.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|mosfetType|type of MOSFET|String|Yes|
|mosfetChannelType|doping of a transistor's channel - describes whether a transistor is n-type or p-type|String|Yes|
|transistorCount|number of transistors in the package|Number| |
|vgs|gate to source voltage difference of a MOSFET|../common/values.json#/valueOptions| |
|vgsMax|maximum gate to source voltage difference that can be continuously applied to a MOSFET. This is a limiting value|../common/values.json#/valueOptions| |
|vds|drain to source voltage difference of a MOSFET|../common/values.json#/valueOptions| |
|vdsMax|maximum drain to source voltage difference that can be continuously applied to a MOSFET. This is a limiting value|../common/values.json#/valueOptions| |
|vdsVbr|drain to source breakdown voltage of a MOSFET|../common/values.json#/valueOptions| |
|vgsTh|gate to source voltage difference required to produce a conducting path between drain and source|../common/values.json#/valueOptions| |
|vsdDiodeVf|reverse diode forward voltage when a MOSFET is in off-state|../common/values.json#/valueOptions| |
|iD|drain Current of a MOSFET|../common/values.json#/valueOptions| |
|iDrain|maximum continuous DC current that can flow through a MOSFET channel.This is a limiting value|../common/values.json#/valueOptions| |
|idPulsed|maximum pulsed DC current that can flow through a MOSFET channel.This is a limiting value|../common/values.json#/valueOptions| |
|iDss|drain-source leakage current of a MOSFET when the gate to source voltage difference is zero|../common/values.json#/valueOptions| |
|iGss|gate-source leakage current of a MOSFET when the drain to source voltage difference is zero|../common/values.json#/valueOptions| |
|diodeContinuousCurrent|maximum continuous forward current of the body diode of a MOSFET (IS).This is a limiting value|../common/values.json#/valueOptions| |
|diodePulsedCurrent|maximum pulsed forward current of the body diode of a MOSFET. This is a limiting value|../common/values.json#/valueOptions| |
|forwardTransconductance|signal gain, change in drain current with variation of gate-source voltage of a MOSFET (gFS)|../common/values.json#/valueOptions| |
|rdson|on-state resistance of a MOSFET|../common/values.json#/valueOptions| |
|rg|internal gate resistance of a MOSFET|../common/values.json#/valueOptions| |
|ciss|input capacitance of a MOSFET|../common/values.json#/valueOptions| |
|coss|output capacitance of a MOSFET|../common/values.json#/valueOptions| |
|crss|reverse transfer capacitance of a MOSFET|../common/values.json#/valueOptions| |
|qg|total gate charge of a MOSFET|../common/values.json#/valueOptions| |
|qgd|gate to drain charge of a MOSFET|../common/values.json#/valueOptions| |
|qgs|gate to source charge of a MOSFET|../common/values.json#/valueOptions| |
|qrr|reverse recovery charge of the body diode of a MOSFET|../common/values.json#/valueOptions| |
|idVsVds|graph of drain current (iD) as a function of drain source voltage (vds)|../common/graph.json#/graph| |
|idVsVgs|graph of drain current (iD) as a function of gate source voltage (vgs)|../common/graph.json#/graph| |
|tdON|turn-on delay time of a MOSFET|../common/values.json#/valueOptions| |
|tdOFF|turn-off delay time of a MOSFET|../common/values.json#/valueOptions| |
|riseTime|rise time of a MOSFET|../common/values.json#/valueOptions| |
|fallTime|fall time of a MOSFET|../common/values.json#/valueOptions| |
|trr|reverse recovery time of the body diode of a MOSFET|../common/values.json#/valueOptions| |
|pTot|maximum power dissipation of a MOSFET|../common/values.json#/valueOptions| |
|pdVsTemp|graph of power dissipation vs temperature|../common/graph.json#/graph| |

### 4.17	 Sensor

####  4.17.1	 Specification Of Accelerometer

Source: [accelerometer.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/accelerometer.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|accelerometerType|type of accelerometer|String| |
|accelerationRanges|range of force that accelerometer can measure|array of Number| |
|accelerationSensitivity|smallest change in force the accelerometer is able to capture (typical) at a given acceleration range|Number| |
|accelerationSensitivityOverTemperature|accelerometer sensitivity change over temperature|Number| |
|axis|number of axes of acceleration measurement|Number| |
|zerogOffset|output of the accelerometer when no acceleration is applied|Number| |
|zerogOffsetOverTemperature|accelerometer zero-g offset change over temperature|Number| |
|outputType|measurement output type|String| |
|outputResolution|output resolution of acceleration measurement|../common/values.json#/valueOptions| |
|interface|interface(s) for communication to host|array of String| |
|bandwidth|bandwidth of acceleration measurement|../common/values.json#/valueOptions| |
|outputDataRate|output Data rate (ODR) of a device|../common/values.json#/valueOptions| |
|rmsNoise|broadband rms noise of a device|Number| |
|spectralNoiseDensity|spectral noise density of a device|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.17.2	 Specification Of Gyroscope

Source: [gyroscope.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/gyroscope.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|gyroRanges|range of angular speed that gyro can measure|Number| |
|gyroSensitivity|smallest change in angular speed gyro is able to capture (typical) at a given gyro range|../common/values.json#/valueOptions| |
|gyroSensitivityOverTemperature|gyroscope sensitivity change over temperature|Number| |
|axis|number of axes of measurement|Number| |
|zeroRateOffset|output of the gyroscope when no angular velocity is applied|Number| |
|zeroRateOffsetOverTemperature|gyro zero rate offset change over temperature|Number| |
|interface|interface(s) for communication to host|array of String| |
|bandwidth|bandwidth of gyroscope|../common/values.json#/valueOptions| |
|outputDataRate|output Data rate (ODR) of a device|../common/values.json#/valueOptions| |
|outputType|measurement output type|String| |
|rmsNoise|broadband rms noise of a device|Number| |
|spectralNoiseDensity|spectral noise density of a device|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.17.3	 Specification Of Magnetic Sensor

Source: [magnetic_sensor.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/magnetic_sensor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|magneticSensingTechnology|method by which magnetism is detected|String| |
|outputType|measurement output type|String| |
|quiescentOutput|output of the magnetic sensor when no magnet is present|../common/values.json#/valueOptions| |
|outputVoltageLinearRange|output voltage range over which the magnetic flux density response is linear|../common/values.json#/valueOptions| |
|linearMagneticSensingRange|magnetic flux density range over which the output voltage is linear |Number| |
|sensitivity|this is the gain - amount of change in the output voltage for a change in the magnetic flux density|Number| |
|operatePoint|magnetic flux density threshold which causes the sensor output to turn on|Number| |
|releasePoint|magnetic flux density threshold which causes the sensor output to turn off|Number| |
|outputPolarity|indicates whether the sensor output is active high or active low|String| |
|hysteresis|delta between the operate point and the release point threshold|Number| |
|bandwidth|sensing bandwidth|../common/values.json#/valueOptions| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.17.4	 Specification Of Thermal Sensor

Source: [thermal_sensor.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/sensor/thermal_sensor.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|sensingTechnology|method by which temperature is detected|String| |
|outputType|measurement output type|String| |
|interface|interface(s) for communication to host|array of String| |
|accuracy|accuracy of temperature sensor|../common/values.json#/valueOptions| |
|temperatureRange|range of temperature sensor|../common/values.json#/valueOptions| |
|resolution|maximum resolution of temperature sensor|Number| |
|gain|amount of change in the output voltage for a change in temperature|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.18	 Storage

####  4.18.1	 Specification Of SD Card

Source: [sd_card.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/storage/sd_card.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|type|type of sd card|String| |
|capacity|capacity of sd card|../common/values.json#/valueOptions|Yes|
|dataRate|maximum data rate of the sd card|Number| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

####  4.18.2	 Specification Of SSD

Source: [ssd.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/storage/ssd.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|partType|part type|String|Yes|
|type|type of ssd storage as defined by interface and technology|String| |
|capacity|capacity of the ssd|../common/values.json#/valueOptions|Yes|
|dataRate|maximum data rate of the ssd|Number| |
|interface|interface of ssd to host|String| |
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

### 4.19	 Undefined

####  4.19.1	 Specification Of Undefined IC

Source: [undefined_ic.json](https://github.com/edatasheets/digital-datasheets/blob/main/part-spec/undefined/undefined_ic.json)

|Property|Description|JSON Data Type|Required?|
|:----|:----|:----|:----|
|currentConsumption|current used by device in various power modes|array of ../common/currentConsumption.json#/currentConsumption| |

