## Contents
* [Specification: component level](part-spec/component.json)
* [Specification: pin level](part-spec/definitions.json)
* [Component examples](support-docs/examples.md)


## Background
1	Introduction 
As the demand for hardware design automation tools increases, there is a need for machine readable datasheets. 
Defining a common digital datasheet specification reduces the burden for component vendors to deliver multiple datasheets to support different tools and encourages reuse of tools across multiple designs. 

1.1	Objectives
Part manufacturers create datasheets to document component information including performance, electrical characteristics, size, orientation, packaging, etc. Digital datasheets contain this information in a machine readable format.The objective of this specification is to create a uniform, machine-readable format for representing component datasheets.

1.2	Scope
This document is intended for digital datasheet producers such as components vendors as well as people who will consume these digital datasheets to automate designs such as tool creators.The specification will include common classes of components used in designs. Some examples are included in the appendices to provide context for future additions to this specification.

1.3	Keywords
Required: The field is required in the component digital datasheet.
Optional: The field may or may not be included in the component digital datasheet.
Shall: Indicates a mandatory requirement.
May: Indicates an optional requirement.

1.4	References
1)	The JSON data interchange syntax, ECMA-404, 2nd edition, December 2017 
2)	JSON Schema: A media Type for Describing JSON Documents, draft-bhutton-json-schema-00, December 2020.
3)	Terms, Definitions, and Letter Symbols for Microelectronic Devices, JESD99C – December 2012

2	Use Cases 
Digital datasheets can be used for many applications. The list of use cases included here is not meant to be exhaustive and it is expected that new applications will be developed as more people start using these digital datasheets. Sample applications include automated hardware design checks to identify bugs earlier in the design cycle, automated hardware designs to speed up board development, components comparison to identify replacement components on a design. 

3	Specifications
3.1	Digital Datasheet Format
A digital datasheet shall be written in the JSON (JavaScript Object Notation) format. The specification is written in JSON schema to facilitate validation of a JSON datasheet. 
3.2	Required Information.
A digital datasheet shall include the following information:
The manufacturer’s name
The Manufacturer Part Number (MPN)
Information to identify the source datasheet
The datasheet version given by date or GUID 
The component part type. See appendix for examples.
List of Component Pins as defined by the specification.

3.3	Date Format
A digital datasheet shall follow the international standard notation, YYYY-MM-DD.

3.4	Units
The digital datasheet shall follow the International System of Units, SI except for specific cases, like package dimensions, where the imperial system of units may be used. These specific cases will be addressed in the specifications.
Properties used the represent units in the datasheets are listed in Github: 
They include:
|Unit Property|Description|Data Type|Required?|
|:----|:----|:----|:----|
|siUnit|Name of the SI unit of measure|String|Yes|
|absoluteValue|Unit quantity corresponding to unit text|Number|Yes|
|unitText|Human readable text describing value|String|Yes|
|unitFactor|Multiplier on the value to achieve the SI unit|Number|Yes|
|relativeValueReference|if unit quantity is based on another reference, value of the reference|String| |
|relativeValueModifier|if a unit quantity is based on another reference, the value that edits that reference|Number| |
|relativeValueOperator|if a unit quantity is based on another reference, the operation that is performed with the modifier|String| |
|valueDefined|a boolean representing whether a value has been defined|Boolean| |

3.5	Missing Values
Missing fields shall be left blank.

3.6	Pins Specifications
The pins specification is included in Github at:
Table below shows the list of properties used to describe pins

|Pin Property|Description|Data Type|Required?|
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






