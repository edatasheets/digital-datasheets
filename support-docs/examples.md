Single Inverter Example

```json
{
    "componentID": {
        "partType": "inverter",
        "manufacturer": "ACME Components",
        "componentName": "abc123",
        "orderableMPN": "abc123,45",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-23",
            "datasheetURI": "www.acmecomponents.com/abc123datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-03-21",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active",
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "vcc",
            "standardizedName": "vdd",
            "description": "power",
            "absVmax": {
                "siUnit": "volts",
                "absoluteValue": 7,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "absVmin": {
                "siUnit": "volts",
                "absoluteValue": -0.5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "vmax": {
                "siUnit": "volts",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "voltageOptions": [
                {
                    "siUnit": "volts",
                    "absoluteValue": 3.3,
                    "unitText": "V",
                    "unitFactor": 1
                },
                {
                    "siUnit": "volts",
                    "absoluteValue": 1.8,
                    "unitText": "V",
                    "unitFactor": 1
                }
            ],
            "esd": true
        },
        {
            "terminalIdentifier": "2",
            "name": "gnd",
            "standardizedName": "vss",
            "description": "ground"
        },
        {
            "terminalIdentifier": "3",
            "name": "A",
            "description": "input",
            "absVmax": {
                "siUnit": "volts",
                "absoluteValue": 7,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "absVmin": {
                "siUnit": "volts",
                "absoluteValue": -0.5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "vmax": {
                "siUnit": "volts",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "vihMin": {
                "siUnit": "volts",
                "absoluteValue": 1.5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "vilMax": {
                "siUnit": "volts",
                "absoluteValue": 0.9,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "inputLeakage": {
                "siUnit": "amps",
                "absoluteValue": 0.1,
                "unitText": "uA",
                "unitFactor": 1E-6,
                "valueDefined": true
            },
            "floatUnused": true,
            "internalPullUp": {
                "siUnit": "ohms",
                "absoluteValue": 20,
                "unitText": "Kohms",
                "unitFactor": 1E3,
                "valueDefined": true
            },
            "esd": true
        },
        {
            "terminalIdentifier": "4",
            "name": "Y",
            "description": "output",
            "absVmax": {
                "siUnit": "volts",
                "absoluteValue": 7,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "absVmin": {
                "siUnit": "volts",
                "absoluteValue": -0.5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "vmax": {
                "siUnit": "volts",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "vol": {
                "siUnit": "volts",
                "absoluteValue": 0.1,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "voh": {
                "siUnit": "volts",
                "absoluteValue": 3.2,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true
            },
            "imax": {
                "siUnit": "amps",
                "absoluteValue": 8,
                "unitText": "mA",
                "unitFactor": 1E-3,
                "valueDefined": true
            },
            "esd": true
        }
    ]
}
```

Single Complex Pin Example 

```json
{
    "terminalIdentifier": "A6",
    "name": "ALERT#/KSI1/I2C1_SCL/GPIO6",
    "numberOfSupportedFunctions": 4,
    "functionProperties": [
        {
            "perFunctionName": "ALERT#",
            "interfaceType": "interrupt",
            "direction": "out",
            "electricalConfiguration": "open-drain",
            "polarity": "low"
        },
        {
            "perFunctionName": "KSI1",
            "interfaceType": "keyboard",
            "pinUsage": "KSI1",
            "direction": "in",
            "electricalConfiguration": "high-impedance",
            "polarity": "low"
        },
        {
            "perFunctionName": "I2C1_SCL",
            "interfaceType": "i2c"
            "pinUsage": "I2C_SCL",
            "direction": "out",
            "electricalConfiguration": "open-drain",
            "polarity": "low"
        },
        {
            "perFunctionName": "GPIO6",
            "interfaceType": "gpio",
            "direction": "bidir",
            "electricalConfiguration": "push-pull"
        }
    ],
    "vihMin": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "VDD1",
        "relativeValueModifier": 0.7,
        "relativeValueOperator": "multiply",
        "valueDefined": true
    },
    "vihMin": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "VDD1",
        "relativeValueModifier": 0.2,
        "relativeValueOperator": "multiply",
        "valueDefined": true
    },
    "vol": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "GND",
        "relativeValueModifier": 0.1,
        "relativeValueOperator": "add",
        "valueDefined": true
    },
    "voh": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "VDD1",
        "relativeValueModifier": 0.1,
        "relativeValueOperator": "subtract",
        "valueDefined": true
    },
    "externalComponents": {
        "componentType": "resistor",
        "configuration": "pu",
        "minValue": {
            "siUnit": "ohm",
            "absoluteValue": 20,
            "unitText": "Kohms",
            "unitFactor": 1E3,
            "valueDefined": true
        }
    }
}

```

Inductor Example

```json
{
    "componentID": {
        "partType": "inductor",
        "manufacturer": "ACME Components",
        "componentName": "abc123",
        "orderableMPN": "abc123,45",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-23",
            "datasheetURI": "www.acmecomponents.com/abc123datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-03-21",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active",
    },
    "value": {
        "siUnit": "henry",
        "absoluteValue": 10,
        "unitText": "uH",
        "unitFactor": 1E-6,
        "valueDefined": true
    },
    "ratedCurrent": {
        "siUnit": "amp",
        "absoluteValue": 4,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "saturationCurrent": {
        "siUnit": "amp",
        "absoluteValue": 6,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "saturationCurve": {
        "title": "DC Bias Curve for abc123",
        "xUnits": "Amps",
        "xLabel": "Current",
        "yUnits": "uH",
        "yLabel": "Inductance",
        "numberOfCurves": 1,
        "data": {
            "label": "20 degrees C",
            "xData": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9],
            "yData": [10, 9.9, 9.7, 9.5, 9.3, 9, 7, 5, 2.5, 2]
        }
    },
    "package": {
        "standardPackageSize": "0805"  
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "1"
        },
        {
            "terminalIdentifier": "2",
            "name": "2"
        }
    ]
}
```

Buck Example

```json
{
    "componentID": {
        "partType": "inductor",
        "manufacturer": "ACME Components",
        "componentName": "abc123",
        "orderableMPN": "abc123,45",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-23",
            "datasheetURI": "www.acmecomponents.com/abc123datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-03-21",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active",
    },
    "regulatorTopology": "buck",
    "vinMin": {
        "siUnit": "volt",
        "absoluteValue": 12,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vinMax": {
        "siUnit": "volt",
        "absoluteValue": 18,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "voutMax": {
        "siUnit": "volt",
        "absoluteValue": 8,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "voutMin": {
        "siUnit": "volt",
        "absoluteValue": 4,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "maxLoadCurrent": {
        "siUnit": "amp",
        "absoluteValue": 5,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "integratedFets": true,
    "integratedFetProperties": {
        "rdsonHSFET": {
            "siUnit": "ohm",
            "absoluteValue": 5,
            "unitText": "mOhm",
            "unitFactor": 1E-3,
            "valueDefined": true
        },
        "rdsonLSFET": {
            "siUnit": "ohm",
            "absoluteValue": 3.5,
            "unitText": "mOhm",
            "unitFactor": 1E-3,
            "valueDefined": true
        },
    },
    "package": {
        "length": 5,
        "width": 5,
        "height": 1, 
        "dimensionUnit": "millimeter" 
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "Vin"
        },
        {
            "terminalIdentifier": "2",
            "name": "Vout"
        },
        {
            "terminalIdentifier": "3",
            "name": "GND"
        },
         {
            "terminalIdentifier": "4",
            "name": "VFB"
        }
    ]
}
```

Diode Example 

```json
{
    "componentID": {
        "partType": "diode",
        "manufacturer": "ACME Components",
        "componentName": "abc1234",
        "orderableMPN": "abc1234,56",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-30",
            "datasheetURI": "www.acmecomponents.com/abc1234datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-09-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "type": "schottky",
    "diodeCount": 1,
    "diodeConfiguration": "unidirectional",
    "numberOfProtectedLines": 0,
    "vfTyp": {
        "value": {
            "siUnit": "volt",
            "absoluteValue": 600,
            "unitText": "mV",
            "unitFactor": 1e-3,
            "valueDefined": true
        },
        "conditions": {
            "if":{
                "siUnit": "amp",
                "absoluteValue": 800,
                "unitText": "mA",
                "unitFactor": 1e-3,
                "valueDefined": true 
            }
        }            
    },
    "vfMax": {
        "value": {
            "siUnit": "volt",
            "absoluteValue": 650,
            "unitText": "mV",
            "unitFactor": 1e-3,
            "valueDefined": true
        },
        "conditions": {
            "if":{
                "siUnit": "amp",
                "absoluteValue": 800,
                "unitText": "mA",
                "unitFactor": 1e-3,
                "valueDefined": true 
            }
        }       
    },
    "ifm": {
        "siUnit": "amp",
        "absoluteValue": 1,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "ifrm": {
        "siUnit": "amp",
        "absoluteValue": 3,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "ifsm": {
        "siUnit": "amp",
        "absoluteValue": 8,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vr": {
        "siUnit": "volt",
        "absoluteValue": 25,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "irTyp": {
        "value": {
            "siUnit": "amp",
            "absoluteValue": 10,
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "conditions": {
            "vr":{
                "siUnit": "volt",
                "absoluteValue": 10,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }      
    },
    "irMax": {
        "value":{
            "siUnit": "amp",
            "absoluteValue": 20,
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "conditions": {
            "vr":{
                "siUnit": "volt",
                "absoluteValue": 25,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }    
    },
    "ifVsVf":{
        "title": "If vs Vf",
        "xUnits": "mA",
        "xLabel": "forward Current",
        "yUnits": "mV",
        "yLabel": "forward voltage",
        "numberOfCurves": 1,
        "data": {
            "label": "25 degrees C",
            "xData": [1, 10, 50, 100, 150, 200, 250, 500, 800],
            "yData": [200, 250, 350, 380, 400, 450, 500, 575, 600]
        }
    },
    "package": {
        "length": 5,
        "width": 5,
        "height": 1, 
        "dimensionUnit": "millimeter" 
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "cathode"
        },
        {
            "terminalIdentifier": "2",
            "name": "anode"
        }
    ]
}
```

FET Example 

```json
{
    "componentID": {
        "partType": "mosfet",
        "manufacturer": "ACME Components",
        "componentName": "abc123",
        "orderableMPN": "abc123,56",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-30",
            "datasheetURI": "www.acmecomponents.com/abc123datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-09-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "mosfetType": "power",
    "mosfetChannelType": "nType",
    "transistorCount": 1,
    "vgsMax": {
        "siUnit": "volt",
        "absoluteValue": 5,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vdsMax": {
        "siUnit": "volt",
        "absoluteValue": 5,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vdsMax": {
        "siUnit": "volt",
        "absoluteValue": 8,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vgsThMax": {
        "siUnit": "volt",
        "absoluteValue": 2,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vgsThTyp": {
        "siUnit": "volt",
        "absoluteValue": 1.2,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vgsThMin": {
        "siUnit": "volt",
        "absoluteValue": 1,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vsdDiodeVfTyp": {
        "siUnit": "volt",
        "absoluteValue": 0.6,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vsdDiodeVfMax": {
        "siUnit": "volt",
        "absoluteValue": 0.9,
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "iDrain": {
        "siUnit": "amp",
        "absoluteValue": 5,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "idPulsed": {
        "siUnit": "amp",
        "absoluteValue": 10,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "iDss": {
        "value": {
            "siUnit": "amp",
            "absoluteValue": 1,
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "conditions": {
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 10,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }    
    },
    "iGss": {
        "value": {
            "siUnit": "amp",
            "absoluteValue": 5,
            "unitText": "nA",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
        "conditions": {
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }    
    },
    "diodeContinuousCurrent": {
        "siUnit": "amp",
        "absoluteValue": 2,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "diodePulsedCurrent": {
        "siUnit": "amp",
        "absoluteValue": 4,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "rdsonTyp": {
        "value": {
            "siUnit": "ohm",
            "absoluteValue": 20,
            "unitText": "mohm",
            "unitFactor": 1e-3,
            "valueDefined": true
        },
        "conditions": {
            "id":{
                "siUnit": "amp",
                "absoluteValue": 500,
                "unitText": "mA",
                "unitFactor": 1e-3,
                "valueDefined": true 
            },
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 1,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "rgMax": {
        "value": {
            "siUnit": "ohm",
            "absoluteValue": 5,
            "unitText": "ohm",
            "unitFactor": 1,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "siUnit": "volt",
                "absoluteValue": 0,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 0,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "CissMax": {
        "value": {
            "siUnit": "farad",
            "absoluteValue": 500,
            "unitText": "pF",
            "unitFactor": 1e-12,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "siUnit": "volt",
                "absoluteValue": 0,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 0,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "CossMax": {
        "value": {
            "siUnit": "farad",
            "absoluteValue": 100,
            "unitText": "pF",
            "unitFactor": 1e-12,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "siUnit": "volt",
                "absoluteValue": 0,
                "unitText": "V",
                "unitFactor": 5,
                "valueDefined": true 
            },
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 0,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "qgMax": {
        "value": {
            "siUnit": "coulomb",
            "absoluteValue": 20,
            "unitText": "nC",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
                "conditions": {
            "vds":{
                "siUnit": "volt",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "tdONTyp": {
        "value": {
            "siUnit": "second",
            "absoluteValue": 20,
            "unitText": "ns",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "siUnit": "volt",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "siUnit": "volt",
                "absoluteValue": 5,
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "pTot": {
        "value": {
            "siUnit": "watt",
            "absoluteValue": 2,
            "unitText": "w",
            "unitFactor": 1,
            "valueDefined": true
        },
        "conditions": {
            "temperature":{
                "siUnit": "celsius",
                "absoluteValue": 25,
                "unitText": "C",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "package": {
        "length": 5,
        "width": 5,
        "height": 1,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "source"
        },
        {
            "terminalIdentifier": "2",
            "name": "source"
        },
        {
            "terminalIdentifier": "3",
            "name": "gate"
        },
        {
            "terminalIdentifier": "4",
            "name": "drain"
        },
        {
            "terminalIdentifier": "5",
            "name": "drain"
        },
        {
            "terminalIdentifier": "6",
            "name": "drain"
        }
    ]
}
```

Retimer Example 

```json
{
    "componentID": {
        "partType": "retimer",
        "manufacturer": "ACME Components",
        "componentName": "abc1234",
        "orderableMPN": "abc1234,68",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-30",
            "datasheetURI": "www.acmecomponents.com/abc1234datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-09-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "numberOfLanes": 2,
    "interface":"PCIe4.0",
    "maxDataRate":{
        "siUnit": "hertz",
        "absoluteValue": 16,
        "unitText": "Ghz",
        "unitFactor": 1e9,
        "valueDefined": true

    },
    "commonClock": true,
    "sris": false,
    "srns": false,
    "package": {
        "length": 100,
        "width": 100,
        "height": 50,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "data1p"
        },
        {
            "terminalIdentifier": "2",
            "name": "data1n"
        },
        {
            "terminalIdentifier": "3",
            "name": "data2p"
        },
        {
            "terminalIdentifier": "4",
            "name": "data2n"
        },
        {
            "terminalIdentifier": "5",
            "name": "refclkp"
        },
        {
            "terminalIdentifier": "6",
            "name": "refclkn"
        },
        {
            "terminalIdentifier": "7",
            "name": "pwr1"
        },
        {
            "terminalIdentifier": "8",
            "name": "pwr2"
        },
        {
            "terminalIdentifier": "9",
            "name": "clkreq"
        },
        {
            "terminalIdentifier": "10",
            "name": "clkreqn"
        },
        {
            "terminalIdentifier": "11",
            "name": "perstn"
        },
        {
            "terminalIdentifier": "12",
            "name": "gnd"
        }

    ]
}
```

USB-C PD Controller Example 

```json
{
    "componentID": {
        "partType": "usb-c pd controller",
        "manufacturer": "ACME Components",
        "componentName": "abc12345",
        "orderableMPN": "abc12345,68",
        "sourceDatasheetID": {
            "publishedDate": "2021-12-30",
            "datasheetURI": "www.acmecomponents.com/abc12345datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-09-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "pdVersion": "pd3.1",
    "usbTypecRevision": "1.2",
    "powerRoleSupported": "source",
    "fastRoleSwapSupport": false,
    "vconnPowerSupport": true,
    "vconnPowerLimit": {
        "siUnit": "watt",
        "absoluteValue": 2.5,
        "unitText": "w",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vconnMaxCurrent": {
        "siUnit": "amp",
        "absoluteValue": 500,
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "vconnOverCurrentLimit": {
        "siUnit": "amp",
        "absoluteValue": 600,
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "integratedLoadSwitch": true,
    "maxSourceCurrent": {
        "siUnit": "amp",
        "absoluteValue": 3,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "sourcefetOverCurrentLimit": {
        "siUnit": "amp",
        "absoluteValue": 4,
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "onResistanceSourceFet": {
        "siUnit": "ohm",
        "absoluteValue": 5,
        "unitText": "mohm",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "activeCurrent": {
        "siUnit": "amp",
        "absoluteValue": 10,
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "shutDownCurrent": {
        "siUnit": "amp",
        "absoluteValue": 10,
        "unitText": "uA",
        "unitFactor": 1e-6,
        "valueDefined": true
    },
    "idleCurrent": {
        "siUnit": "amp",
        "absoluteValue": 5,
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "vbusOvervoltageProtectionThresholdRising": {
        "siUnit": "volt",
        "absoluteValue": 6.3,
        "unitText": "v",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vbusOvervoltageProtectionThresholdFalling": {
        "siUnit": "volt",
        "absoluteValue": 6.1,
        "unitText": "v",
        "unitFactor": 1,
        "valueDefined": true
    },
    "package": {
        "length": 100,
        "width": 100,
        "height": 30,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "pwr1"
        },
        {
            "terminalIdentifier": "2",
            "name": "pwr2"
        },
        {
            "terminalIdentifier": "3",
            "name": "enable"
        },
        {
            "terminalIdentifier": "4",
            "name": "scl"
        },
        {
            "terminalIdentifier": "5",
            "name": "sda"
        },
        {
            "terminalIdentifier": "6",
            "name": "gnd"
        },
        {
            "terminalIdentifier": "7",
            "name": "vbus"
        },
        {
            "terminalIdentifier": "8",
            "name": "vbus"
        },
        {
            "terminalIdentifier": "9",
            "name": "cc1"
        },
        {
            "terminalIdentifier": "10",
            "name": "cc2"
        },
        {
            "terminalIdentifier": "11",
            "name": "irq"
        },
        {
            "terminalIdentifier": "12",
            "name": "gnd"
        }
    ]
}
```

Microcontroller/EC Example - Limited pins 

```json
{
    "componentID": {
        "partType": "microcontroller/ec",
        "manufacturer": "ACME Components",
        "componentName": "abc12",
        "orderableMPN": "abc12,34",
        "sourceDatasheetID": {
            "publishedDate": "2021-12-30",
            "datasheetURI": "www.acmecomponents.com/abc12datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-09-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "onChipFlash": {
        "siUnit": "byte",
        "absoluteValue": 1,
        "unitText": "MB",
        "unitFactor": 1e6,
        "valueDefined": true
    },
    "onChipRAM": {
        "siUnit": "byte",
        "absoluteValue": 128,
        "unitText": "KB",
        "unitFactor": 1e3,
        "valueDefined": true
    },
    "onChipROM": {
        "siUnit": "byte",
        "absoluteValue": 1,
        "unitText": "KB",
        "unitFactor": 1e3,
        "valueDefined": true
    },
    "coreProcessor": "xyz",
    "coreArchitectureBits": "32-bit",
    "clockSpeed": {
        "siUnit": "hertz",
        "absoluteValue": 120,
        "unitText": "MHz",
        "unitFactor": 1e6,
        "valueDefined": true
    },
    "activePower": {
        "siUnit": "watt",
        "absoluteValue": 500,
        "unitText": "mW",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "standbyPower": {
        "siUnit": "watt",
        "absoluteValue": 10,
        "unitText": "mW",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "package": {
        "length": 20,
        "width": 20,
        "height": 2,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "vdd1"
        },
        {
            "terminalIdentifier": "2",
            "name": "vdd2"
        },
        {
            "terminalIdentifier": "3",
            "name": "ALERT#/KSI1/I2C1_SCL/GPIO6",
            "numberOfSupportedFunctions": 4,
            "functionProperties": [
                {
                    "perFunctionName": "ALERT#",
                    "interfaceType": "interrupt",
                    "direction": "out",
                    "electricalConfiguration": "open-drain",
                    "polarity": "low"
                },
                {
                    "perFunctionName": "KSI1",
                    "interfaceType": "keyboard",
                    "pinUsage": "KSI1",
                    "direction": "in",
                    "electricalConfiguration": "high-impedance",
                    "polarity": "low"
                },
                {
                    "perFunctionName": "I2C1_SCL",
                    "interfaceType": "i2c",
                    "pinUsage": "I2C_SCL",
                    "direction": "out",
                    "electricalConfiguration": "open-drain",
                    "polarity": "low"
                },
                {
                    "perFunctionName": "GPIO6",
                    "interfaceType": "gpio",
                    "direction": "bidir",
                    "electricalConfiguration": "push-pull"
                }
            ],
            "vihMax": {
                "siUnit": "volts",
                "unitText": "V",
                "unitFactor": 1,
                "relativeValueReference": "VDD1",
                "relativeValueModifier": 0.7,
                "relativeValueOperator": "multiply",
                "valueDefined": true
            },
            "vihMin": {
                "siUnit": "volts",
                "unitText": "V",
                "unitFactor": 1,
                "relativeValueReference": "VDD1",
                "relativeValueModifier": 0.2,
                "relativeValueOperator": "multiply",
                "valueDefined": true
            },
            "vol": {
                "siUnit": "volts",
                "unitText": "V",
                "unitFactor": 1,
                "relativeValueReference": "GND",
                "relativeValueModifier": 0.1,
                "relativeValueOperator": "add",
                "valueDefined": true
            },
            "voh": {
                "siUnit": "volts",
                "unitText": "V",
                "unitFactor": 1,
                "relativeValueReference": "VDD1",
                "relativeValueModifier": 0.1,
                "relativeValueOperator": "subtract",
                "valueDefined": true
            },
            "externalComponents": {
                "componentType": "resistor",
                "configuration": "pu",
                "minValue": {
                    "siUnit": "ohm",
                    "absoluteValue": 20,
                    "unitText": "Kohms",
                    "unitFactor": 1E3,
                    "valueDefined": true
                }
            },
        },
        {
            "terminalIdentifier": "4",
            "name": "vrefp"
        },
        {
            "terminalIdentifier": "5",
            "name": "vrefn"
        },
        {
            "terminalIdentifier": "6",
            "name": "gnd"
        }
    ]
}
