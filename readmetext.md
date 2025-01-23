## Background
### 1 Introduction 
As the demand for hardware design automation tools grows, the need for machine-readable datasheets becomes more critical. Establishing a standardized Electronic Datasheet (EDS) specification alleviates the burden on component vendors to produce multiple datasheets for various tools, promoting the reuse of tools across different designs.

#### 1.1 Objectives
Component manufacturers produce datasheets to document various details about their components, such as performance, electrical characteristics, size, orientation, and packaging. EDS provide this information in a machine-readable format. The goal of this specification is to establish a standardized, machine-readable format for representing component datasheets.

#### 1.2 Scope
This document is designed for EDS producers, such as component vendors, and for individuals who will use EDS to automate designs, such as tool creators. The specification will cover common classes of components used in designs. Examples are included in the appendices to provide context for future additions to this specification.

#### 1.3 Keywords
- Required: The field is required in the EDS.
- Optional: The field may or may not be included in the EDS.
- Shall: Indicates a mandatory requirement.
- May: Indicates an optional requirement.

#### 1.4 References
- The JSON data interchange syntax, [ECMA-404, 2nd edition, December 2017](https://www.ecma-international.org/publications-and-standards/standards/ecma-404/)
- JSON Schema: A media Type for Describing JSON Documents, [draft-bhutton-json-schema-00, December 2020](https://datatracker.ietf.org/doc/html/draft-bhutton-json-schema-00).
- Terms, Definitions, and Letter Symbols for Microelectronic Devices, JESD99C – December 2012

### 2. Working with Electronic Datasheets 

#### 2.1 Electronic Datasheet Distribution by Component Vendors
EDS released by a hardware component vendor can include one or multiple files. Multiple files are typically used when a single datasheet covers multiple parts from the same family, when a part is complex or, when a part has additional information, not in JSON format, such as footprints. In this case, the specification has a provision that enables vendors to capture common or core properties across all the parts in one file while using additional files to capture the changes between parts.

There are multiple existing models for hardware component vendors to distribute their parts datasheets. The recommendation is for these vendors to leverage these existing models for EDS. Existing models include private and public distributions directly from the vendors or through an electronic components’ distributor.

In a public release, the hardware component vendor makes their latest EDS available to the public, generally, through a public website. An example of this might be a GitHub repository where each vendor has a repository. The EDS can be directly downloaded through the website without any requirements, and customers can optionally add their contact information to be notified about updates in the EDS. 

Electronic components distributors have agreements with hardware component vendors to distribute their products to customers.  In the existing model, distributors either include a local copy of the component datasheet to their website or add links to the datasheets on the vendor’s website.  In the case of EDS, the recommendation is to use the later model, with an EDS link always pointing to the latest release.  The goal is to avoid a situation where copies of local EDS become out of sync with the latest release. 

In a private distribution model, the EDS is shared by a hardware component vendor to one or more select customers, for Intellectual Property (IP) protection. This may require an NDA between the parties involved. The EDS can be shared through means like a private repository or a private database, or through encrypted emails.  With this model, the vendors have a good visibility of the customers using their EDS and can directly notify them of updates through email.  It’s also recommended to provide a release document identifying EDS changes from one release to the next.

#### 2.2 Importing Electronic Datasheets in a Part Library Management System
Most companies doing hardware design rely on a Part Library Management (PLM) System to store information about electronic components. These systems are often linked to CAD tools to guide component selection during hardware design. Additionally, these systems have an important role in Bill of Materials (BOM) generation and purchasing. Therefore, a natural entry point for EDS information is a PLM system.

Given the importance of component accuracy in a PLM system, it is likely that companies will initially want some manual review of new EDS before they are integrated into a shared PLM system. One suggested method of achieving this is to have an internal datasheet database that is connected to any datasheet repositories (public or private) that are used by that company. Updates can be manually pulled in and compared as new datasheets are available. Once the updates have been checked, new datasheet changes can be submitted to the internal database. This database would be indexed by internal part numbers that correspond to the EDS.

Converting this internal company datasheet database to information in the PLM system can be done automatically because datasheets in the internal database have already been manually verified. To do this conversion, a one-time mapping file between properties in the EDS schema and the specific properties in a company’s PLM is created. Companies can then write a simple generator script using this mapping file to convert EDS entries to PLM properties. Additionally, as EDS will likely contain new information not available in the PLM system, this information can be added or exist in a separate database alongside the PLM system. As tooling advances, these new properties can be leveraged to provide richer schematic validation. 

The advantage of this system is it leverages existing PLM systems, which are often deeply integrated with engineering workflows. However, it makes these PLM systems more efficient to generate and update.

#### 2.3 Tools to Generate Electronic Datasheets 
The Digital Datasheet Creator (DDC) is an open-source tool that is available to make datasheet
creation less time consuming. The tool can be found here: [datasheet creator](https://github.com/intel/digital-datasheet-creator). It consists of a series of template spreadsheets for each 
part type that can be filled in and run through the creator script to generate a specification-compliant JSON datasheet file. 

### 3. Use Cases 
EDS can be used for many applications. The list of use cases included here is not meant to be exhaustive and it is expected that new applications will be developed as more people start using these EDS. Sample applications include automated hardware design checks to identify bugs earlier in the design cycle, automated hardware designs to speed up board development, components comparison to identify replacement components on a design. 

### 4.Specifications

#### 4.1 Electronic Datasheet Format
EDS shall be written in the JSON (JavaScript Object Notation) format. The specification is written in JSON schema to facilitate validation of a JSON datasheet. 

#### 4.2 Required Information.
EDS shall include the following information:
- The manufacturer’s name
- The Manufacturer Part Number (MPN)
- Information to identify the source datasheet
- The datasheet version given by date or GUID 
- The component part type. See appendix for examples.
- List of Component Pins as defined by the specification.

#### 4.3 Date Format
EDS shall follow the international standard notation, YYYY-MM-DD.

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
|componentPropertyExternalFiles|external files that describe key component properties. External files can be used in lieu of defining core properties, pins, and package information in the same file|./common/externalFileMap.json#/externalFileMap|No|
|additionalSpecExternalFiles|external files that contain information outside of the json spec. Examples include layout, simulation, etc.|./common/externalFile.json#/externalFile|No|
|reliability|reliability information about the component|./common/reliability.json#/reliability|No|
|powerSequence|information about component power sequencing|./common/powerSequence.json#/powerSequenceTable|No|

