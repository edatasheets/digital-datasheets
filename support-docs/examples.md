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
        "partType": "switching regulator",
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
    "vin": {
        "minValue": 12,
        "maxValue": 18,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vout": {
        "minValue": 4,
        "maxValue": 8,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "loadCurrent": {
        "maxValue": 5,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "integratedFets": true,
    "integratedFetProperties": {
        "rdsonHSFET": {
            "typValue": 5,
            "siUnit": "ohm",
            "unitText": "mOhm",
            "unitFactor": 1E-3,
            "valueDefined": true
        },
        "rdsonLSFET": {
            "typValue": 3.5,
            "siUnit": "ohm",
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

LDO Example

```json
{
    "componentID": {
        "partType": "LDO",
        "manufacturer": "ACME Components",
        "componentName": "abc123456",
        "orderableMPN": "abc123456,45",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-23",
            "datasheetURI": "www.acmecomponents.com/abc123456datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-12-21",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active",
    },
    "vin": {
        "maxValue": 3.3,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vout": {
        "minValue": 1.35,
        "maxValue": 1.65,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "loadCurrent": {
        "maxValue": 1,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
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
            "name": "VIN"
        },
        {
            "terminalIdentifier": "2",
            "name": "VOUT"
        },
        {
            "terminalIdentifier": "3",
            "name": "GND"
        },
        {
            "terminalIdentifier": "4",
            "name": "VFB"
        },
        {
            "terminalIdentifier": "5",
            "name": "EN"
        }
    ]
}
```

Load Switch Example 

```json
{
    "componentID": {
        "partType": "load switch",
        "manufacturer": "ACME Components",
        "componentName": "abc678",
        "orderableMPN": "abc678,56",
        "sourceDatasheetID": {
            "publishedDate": "2021-05-30",
            "datasheetURI": "www.acmecomponents.com/abc678datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-10-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "fetType": "PFET",
    "loadSwitchCount": 1,
     "vin":{
        "typValue": 1.8,
        "minValue": 1.65,
        "maxValue": 3.3,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
     "outputCurrent":{
        "maxValue": 2,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "onResistance":{
        "typValue": 10,
        "siUnit": "ohm",
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "currentLimitSupport": true,
    "powerGoodSupport": true,
    "quickOutputDischargeSupport": true,
    "reverseCurrentBlockingSupport": false,
    "enableTime": {
        "typValue": 1,
        "siUnit": "second",
        "unitText": "s",
        "unitFactor": 1,
        "valueDefined": true
    },
     "offTime": {
        "typValue": 1,
        "siUnit": "second",
        "unitText": "ms",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "package": {
        "length": 3,
        "width": 3,
        "height": 1,
        "dimensionUnit": "millimeter"
    },
     "pins": [
        {
            "terminalIdentifier": "1",
            "name": "VIN"
        },
        {
            "terminalIdentifier": "2",
            "name": "VOUT"
        },
        {
            "terminalIdentifier": "3",
            "name": "GND"
        },
        {
            "terminalIdentifier": "4",
            "name": "EN"
        },
          {
            "terminalIdentifier": "5",
            "name": "PWRGD"
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
    "vf": {
        "value": {
            "typValue": 600,
            "maxValue": 650,
            "siUnit": "volt",
            "unitText": "mV",
            "unitFactor": 1e-3,
            "valueDefined": true
        },
        "conditions": {
            "if":{
                "typValue": 800,
                "siUnit": "amp",
                "unitText": "mA",
                "unitFactor": 1e-3,
                "valueDefined": true 
            }
        }            
    },
    "ifm": {
        "maxValue": 1,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "ifrm": {
       "maxValue": 3,
       "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "ifsm": {
        "maxValue": 8,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vr": {
        "maxValue": 25,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "ir": {
        "value": {
            "typValue": 10,
            "maxValue": 20,
            "siUnit": "amp",
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "conditions": {
            "vr":{
                "maxValue": 10,
                "siUnit": "volt",
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
        "maxValue": 5,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vdsMax": {
        "maxValue": 5,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vgsThMax": {
        "maxValue": 2,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vgsThTyp": {
        "typValue": 1.2,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vgsThMin": {
        "minValue": 1,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vsdDiodeVf": {
        "typValue": 0.6,
        "maxValue": 0.9,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "iDrain": {
        "maxValue": 5,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "idPulsed": {
       "maxValue": 10,
       "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "iDss": {
        "value": {
            "typValue": 1
            "siUnit": "amp",
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "conditions": {
            "vgs":{
                "typValue": 10,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }    
    },
    "iGss": {
        "value": {
            "typValue": 5,
            "siUnit": "amp",
            "unitText": "nA",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
        "conditions": {
            "vgs":{
                "typValue": 5,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }    
    },
    "diodeContinuousCurrent": {
        "maxValue": 2,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "diodePulsedCurrent": {
        "maxValue": 4,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "rdson": {
        "value": {
            "typValue": 20,
            "siUnit": "ohm",
            "unitText": "mohm",
            "unitFactor": 1e-3,
            "valueDefined": true
        },
        "conditions": {
            "id":{
                "typValue": 500,
                "siUnit": "amp",
                "unitText": "mA",
                "unitFactor": 1e-3,
                "valueDefined": true 
            },
            "vgs":{
                "typValue": 1,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "rg": {
        "value": {
            "maxValue": 5,
            "siUnit": "ohm",
            "unitText": "ohm",
            "unitFactor": 1,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "typValue": 0,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "typValue": 0,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "Ciss": {
        "value": {
            "maxValue": 500,
            "siUnit": "farad",
            "unitText": "pF",
            "unitFactor": 1e-12,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "typValue": 0,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "typValue": 0,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "Coss": {
        "value": {
            "maxValue": 100,
            "siUnit": "farad",
            "unitText": "pF",
            "unitFactor": 1e-12,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "typValue": 0.
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 5,
                "valueDefined": true 
            },
            "vgs":{
                "typValue": 0,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "qg": {
        "value": {
            "maxValue": 20,
            "siUnit": "coulomb",
            "unitText": "nC",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
                "conditions": {
            "vds":{
                "typValue": 5,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "typValue": 5,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "tdON": {
        "value": {
            "typValue": 20,
            "siUnit": "second",
            "unitText": "ns",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
        "conditions": {
            "vds":{
                "typValue": 5,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            },
            "vgs":{
                "typValue": 5,
                "siUnit": "volt",
                "unitText": "V",
                "unitFactor": 1,
                "valueDefined": true 
            }
        }
    },
    "pTot": {
        "value": {
            "typValue": 2,
            "siUnit": "watt",
            "unitText": "w",
            "unitFactor": 1,
            "valueDefined": true
        },
        "conditions": {
            "temperature":{
                "typValue": 25,
                "siUnit": "celsius",
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
        "maxValue": 16,
        "siUnit": "hertz",
        "unitText": "GHz",
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
        "maxValue": 2.5,
        "siUnit": "watt",
        "unitText": "w",
        "unitFactor": 1,
        "valueDefined": true
    },
    "vconnMaxCurrent": {
        "maxValue": 500,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "vconnOverCurrentLimit": {
        "typValue": 600,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "integratedLoadSwitch": true,
    "maxSourceCurrent": {
        "typValue": 3,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "sourcefetOverCurrentLimit": {
        "typValue": 4,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
    },
    "onResistanceSourceFet": {
        "typValue": 5,
        "siUnit": "ohm",
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
    "currentConsumption": {
        "supplyName": "pwr1",
        "shutDownCurrent": {
            "typValue": 10,
            "siUnit": "amp",
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "idleCurrent": {
            "typValue": 5,
            "siUnit": "amp",
            "unitText": "mA",
            "unitFactor": 1e-3,
            "valueDefined": true
         }
    },
    "componentProtectionThresholds": [
        {
            "powerSupplyProtection":{
                "supplyName": "vbus",
                "OvervoltageProtectionThresholdRising": {
                    "typValue": 6.3,
                    "siUnit": "volt",
                    "unitText": "v",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "OvervoltageProtectionThresholdFalling": {
                    "typValue": 6.1,
                    "siUnit": "volt",
                    "unitText": "v",
                    "unitFactor": 1,
                    "valueDefined": true
                }
            }
        },
        {
            "powerSupplyProtection":{
                "supplyName": "vconn",
                "OvervoltageProtectionThresholdRising": {
                    "typValue": 5.5,
                    "siUnit": "volt",
                    "unitText": "v",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "OvervoltageProtectionThresholdFalling": {
                    "typValue": 5.3,
                    "siUnit": "volt",
                    "unitText": "v",
                    "unitFactor": 1,
                    "valueDefined": true
                }
            }
        }
    ],
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
            }
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
```

Display Backlight Driver Example - Limited pins 

```json
{
    "componentID": {
        "partType": "display backlight driver",
        "manufacturer": "ACME Components",
        "componentName": "abc100",
        "orderableMPN": "abc100,45",
        "sourceDatasheetID": {
            "publishedDate": "2021-12-23",
            "datasheetURI": "www.acmecomponents.com/abc100datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-12-20",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "vin":{
        "typValue": 10,
        "minValue": 5,
        "maxValue": 24,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
     "ioutPerString":{
        "typValue": 100,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 1e-3,
        "valueDefined": true
    },
    "currentMatchingAccuracy":{
        "typValue": 1,
        "siUnit": "percentage",
         "unitText": "%",
        "unitFactor": 1,
        "valueDefined": true
    },
     "dimmingRatio":{
        {"numerator": 1000},
        {"denominator": 1}
     },
    "fsw":{
        "minValue": 500,
        "maxValue": 2000
        "siUnit": "hertz",
        "unitText": "KHz",
        "unitFactor": 1e3,
        "valueDefined": true
    },
      "package": {
        "length": 5,
        "width": 5,
        "height": 1,
        "dimensionUnit": "millimeter"
    },
    "standardPackageType": "qfn",
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "VIN"
        },
        {
            "terminalIdentifier": "2",
            "name": "OUT0"
        },
        {
            "terminalIdentifier": "3",
            "name": "OUT1"
        },
        {
            "terminalIdentifier": "4",
            "name": "OUT2"
        },
        {
            "terminalIdentifier": "5",
            "name": "OUT3"
        },
        {
            "terminalIdentifier": "6",
            "name": "GND"
        },
        {
            "terminalIdentifier": "7",
            "name": "IREF"
        },
        {
            "terminalIdentifier": "8",
            "name": "VFB"
        },
        {
            "terminalIdentifier": "9",
            "name": "EN"
        }
    ]
}
       
```

PMIC Example - Limited pins 

```json
{
    "componentID": {
        "partType": "pmic",
        "manufacturer": "ACME Components",
        "componentName": "abc10",
        "orderableMPN": "abc10,45",
        "sourceDatasheetID": {
            "publishedDate": "2021-11-23",
            "datasheetURI": "www.acmecomponents.com/abc10datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-03-21",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "ldoRegulatorCount": 2,
    "buckRegulatorCount": 2,
    "componentList": [
        "ldo",
        "switching regulator"
    ],
    "instances": [
        {
            "componentTitle": "ldo",
            "instanceName": "ldo1",
            "instanceProperties": {
                "vin": {
                    "minValue": 3,
                    "maxvalue": 5,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "vout": {
                    "minValue": 1.62,
                    "maxValue": 1.98
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "loadCurrent": {
                   "maxValue": 1,
                   "siUnit": "amp",
                    "unitText": "A",
                    "unitFactor": 1,
                    "valueDefined": true
                }
            }
        },
        {
            "componentTitle": "ldo",
            "instanceName": "ldo2",
            "instanceProperties": {
                "vin": {
                    "minValue": 2.8,
                    "maxValue": 3.3,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "vout": {
                    "minValue": 1.35,
                    "maxValue": 1.65,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "loadCurrent": {
                    "maxValue": 1,
                    "siUnit": "amp",
                    "unitText": "A",
                    "unitFactor": 1,
                    "valueDefined": true
                }
            }
        },
        {
            "componentTitle": "switching regulator",
            "instanceName": "buck1",
            "instanceProperty": {
                "regulatorTopology": "buck",
                "vin": {
                    "minValue": 10,
                    "maxValue": 15,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "vout": {
                    "minValue": 2.7,
                    "maxValue": 3.3,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "loadCurrent": {
                    "maxValue": 10,
                    "siUnit": "amp",
                    "unitText": "A",
                    "unitFactor": 1,
                    "valueDefined": true
                }
            }
        },
        {
            "componentTitle": "switching regulator",
            "instanceName": "buck2",
            "instanceProperty": {
                "regulatorTopology": "buck",
                "vin": {
                    "minValue": 10,
                    "maxvalue": 15,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "voutMax": {
                    "minValue": 4.5,
                    "maxValue": 5.5,
                    "siUnit": "volt",
                    "unitText": "V",
                    "unitFactor": 1,
                    "valueDefined": true
                },
                "loadCurrent": {
                    "maxValue": 5,
                    "siUnit": "amp",
                    "unitText": "A",
                    "unitFactor": 1,
                    "valueDefined": true
                }
            }
        }
    ],
    "package": {
        "length": 10,
        "width": 10,
        "height": 1,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "Vin_ldo1"
        },
        {
            "terminalIdentifier": "2",
            "name": "Vout_ldo1"
        },
        {
            "terminalIdentifier": "3",
            "name": "GND"
        },
        {
            "terminalIdentifier": "4",
            "name": "VFB_ldo1"
        },
        {
            "terminalIdentifier": "5",
            "name": "Vin_ldo2"
        },
        {
            "terminalIdentifier": "6",
            "name": "Vout_ldo2"
        },
        {
            "terminalIdentifier": "7",
            "name": "VFB_ldo2"
        },
        {
            "terminalIdentifier": "8",
            "name": "Vsw_buck1"
        },
        {
            "terminalIdentifier": "9",
            "name": "Vsw_buck2"
        },
        {
            "terminalIdentifier": "10",
            "name": "Vin_buck"
        },
        {
            "terminalIdentifier": "11",
            "name": "GND"
        },
        {
            "terminalIdentifier": "12",
            "name": "sda"
        },
        {
            "terminalIdentifier": "13",
            "name": "scl"
        },
        {
            "terminalIdentifier": "14",
            "name": "gpio1"
        },
        {
            "terminalIdentifier": "15",
            "name": "gpio2"
        },
        {
            "terminalIdentifier": "16",
            "name": "GND"
        }
    ]
}
```

Sensor Example - Limited pins 

```json     
{
    "componentID": {
        "partType": "accelerometer",
        "manufacturer": "ACME Components",
        "componentName": "abc100",
        "orderableMPN": "abc100,456",
        "sourceDatasheetID": {
            "publishedDate": "2022-03-23",
            "datasheetURI": "www.acmecomponents.com/abc100datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2023-03-23",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "accelerationRanges": 8,
    "accelerationSensitivity": {
        "value": {
            "typValue": 256,
            "maxvalue": 257
        },
        "conditions": "max acceleration range is 8g"
    },
    "accelerationSensitivityOverTemperature": {
        "typValue": 0.02,
        "minValue": 0.01,
        "maxValue": 0.03
    },
    "axis": 3,
    "zerogOffset": {
        "minValue": -50,
        "maxValue": 50
    },
    "outputType": "digital",
    "interface": [
        "i2c",
        "spi"
    ],
    "bandwidth": {
        "typValue": 20,
        "minValue": 10,
        "maxValue": 30,
        "siUnit": "hertz",
        "unitText": "MHz",
        "unitFactor": 1000000,
        "valueDefined": true
    },
    "currentConsumption": {
        "supplyName": "vdd",
        "quiescentCurrent": {
            "typValue": 100,
            "minValue": 50,
            "maxValue": 150,
            "siUnit": "amp",
            "unitText": "uA",
            "unitFactor": 1e-6,
            "valueDefined": true
        },
        "shutdownCurrent": {
            "typValue": 250,
            "minValue": 100,
            "maxValue": 300,
            "siUnit": "amp",
            "unitText": "nA",
            "unitFactor": 1e-9,
            "valueDefined": true
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
            "name": "vdd"
        },
        {
            "terminalIdentifier": "2",
            "name": "dout"
        },
        {
            "terminalIdentifier": "3",
            "name": "GND"
        },
        {
            "terminalIdentifier": "4",
            "name": "scl"
        },
        {
            "terminalIdentifier": "5",
            "name": "sda"
        }
    ]
}

```

Converter Example - Limited pins 

```json
{
  "componentID": {
    "partType": "adc",
    "manufacturer": "ACME Components",
    "componentName": "abc1000",
    "orderableMPN": "abc1000,456",
    "sourceDatasheetID": {
      "publishedDate": "2022-03-23",
      "datasheetURI": "www.acmecomponents.com/abc1000datasheet"
    },
    "digitalDatasheetID": {
      "publishedDate": "2023-03-23",
      "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
    },
    "status": "active"
  },
  "digitalResolution": 8,
  "conversionTime": {
    "typValue": 8,
    "siUnit": "second",
    "unitText": "uS",
    "unitFactor": 0.000001,
    "valueDefined": true
  },
  "sampleRate": {
    "typValue": 10,
    "minValue": 1,
    "maxValue": 20,
    "siUnit": "hertz",
    "unitText": "MHz",
    "unitFactor": 1000000,
    "valueDefined": true
  },
  "offsetError": 0.5,
  "gainError": 0.1,
  "integralNonlinearity": 0.1,
  "differentialNonlinearity": 0.1,
  "SNR": 60,
  "inputType": "singleEnded",
  "inputChannels": 8,
  "currentConsumption": {
    "supplyName": "vdd",
    "quiescentCurrent": {
      "typValue": 200,
      "minValue": 50,
      "maxValue": 500,
      "siUnit": "amp",
      "unitText": "uA",
      "unitFactor": 0.000001,
      "valueDefined": true
    },
    "shutdownCurrent": {
      "typValue": 50,
      "minValue": 10,
      "maxValue": 80,
      "siUnit": "amp",
      "unitText": "nA",
      "unitFactor": 1e-9,
      "valueDefined": true
    }
  },
  "package": {
    "length": 10,
    "width": 5,
    "height": 1,
    "dimensionUnit": "millimeter"
  },
  "pins": [
    {
      "terminalIdentifier": "1",
      "name": "vdd"
    },
    {
      "terminalIdentifier": "2",
      "name": "scl"
    },
    {
      "terminalIdentifier": "3",
      "name": "sda"
    },
    {
      "terminalIdentifier": "4",
      "name": "GND"
    },
    {
      "terminalIdentifier": "5",
      "name": "channel1"
    },
    {
      "terminalIdentifier": "6",
      "name": "channel2"
    },
    {
      "terminalIdentifier": "7",
      "name": "vref"
    }
  ]
}
```

BJT Example - 

```json
{
  "componentID": {
    "partType": "bjt",
    "manufacturer": "ACME Components",
    "componentName": "abc123456",
    "orderableMPN": "abc123456,56",
    "sourceDatasheetID": {
      "publishedDate": "2022-05-30",
      "datasheetURI": "www.acmecomponents.com/abc123456datasheet"
    },
    "digitalDatasheetID": {
      "publishedDate": "2023-05-26",
      "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
    },
    "status": "active"
  },
  "bjtChannelType": "nType",
  "transistorCount": 1,
  "collectorCurrent": {
    "maxValue": 150,
    "siUnit": "amp",
    "unitText": "mA",
    "unitFactor": 0.001,
    "valueDefined": true
  },
  "peakCollectorCurrent": {
    "maxValue": 300,
    "siUnit": "amp",
    "unitText": "mA",
    "unitFactor": 0.001,
    "valueDefined": true
  },
  "collectorBaseVoltage": {
    "maxValue": 30,
    "siUnit": "volt",
    "unitText": "V",
    "unitFactor": 1,
    "valueDefined": true
  },
  "emitterBaseVoltage": {
    "maxValue": 5,
    "siUnit": "volt",
    "unitText": "V",
    "unitFactor": 1,
    "valueDefined": true
  },
  "emitterBaseCutOffCurrent": {
    "value": {
      "typValue": 40,
      "siUnit": "amp",
      "unitText": "nA",
      "unitFactor": 1e-9,
      "valueDefined": true
    },
    "conditions": {
      "collectorEmitterVoltage": {
        "typValue": 10,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
      },
      "emitterBaseVoltage": {
        "typValue": 2,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
      }
    }
  },
  "dcGain": {
    "value": {
      "typValue": 10
    },
    "conditions": {
      "collectorEmitterVoltage": {
        "typValue": 5,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
      },
      "collectorCurrent": {
        "typValue": 5,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 0.001,
        "valueDefined": true
      }
    }
  },
  "collectorEmitterBreakdownVoltage": {
    "value": {
      "typValue": 40,
      "siUnit": "volt",
      "unitText": "V",
      "unitFactor": 1,
      "valueDefined": true
    },
    "conditions": {
      "collectorCurrent": {
        "typValue": 5,
        "siUnit": "amps",
        "unitText": "mA",
        "unitFactor": 0.001,
        "valueDefined": true
      },
      "baseCurrent": {
        "typValue": 0,
        "siUnit": "amp",
        "unitText": "A",
        "unitFactor": 1,
        "valueDefined": true
      }
    }
  },
  "delayTime": {
    "value": {
      "typValue": 40,
      "siUnit": "second",
      "unitText": "nS",
      "unitFactor": 1e-9,
      "valueDefined": true
    },
    "conditions": {
      "baseEmitterVoltage": {
        "typValue": 0.7,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
      },
      "baseCurrent": {
        "typValue": 2,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 0.001,
        "valueDefined": true
      },
      "collectorCurrent": {
        "typValue": 20,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 0.001,
        "valueDefined": true
      }
    }
  },
  "package": {
    "length": 1,
    "width": 1,
    "height": 0.5,
    "dimensionUnit": "millimeter"
  },
  "pins": [
    {
      "terminalIdentifier": "1",
      "name": "emitter"
    },
    {
      "terminalIdentifier": "2",
      "name": "base"
    },
    {
      "terminalIdentifier": "3",
      "name": "collector"
    }
  ]
}
```

EEPROM Example - Limited pins 

```json
{
    "componentID": {
        "partType": "eeprom",
        "manufacturer": "ACME Components",
        "componentName": "abc3456",
        "orderableMPN": "abc3456,56",
        "sourceDatasheetID": {
            "publishedDate": "2022-05-30",
            "datasheetURI": "www.acmecomponents.com/abc3456datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2023-05-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "capacity": {
        "typValue": 16,
        "siUnit": "bits",
        "unitText": "Kb",
        "unitFactor": 1000,
        "valueDefined": true
    },
    "numberOfWords": 2048,
    "bitsPerWords": 8,
    "interface": "i2c",
    "clockFrequency": {
        "typValue": 100,
        "siUnit": "hertz",
        "unitText": "KHz",
        "unitFactor": 1000,
        "valueDefined": true
    },
    "endurance": 10,
    "package": {
        "length": 1.5,
        "width": 2,
        "height": 0.5,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "SDA"
        },
        {
            "terminalIdentifier": "2",
            "name": "GND"
        },
        {
            "terminalIdentifier": "3",
            "name": "SCL"
        },
        {
            "terminalIdentifier": "4",
            "name": "WP"
        },
        {
            "terminalIdentifier": "5",
            "name": "VDD"
        }
    ]
}
```

Flash Memory Example - Limited pins 

```json
{
    "componentID": {
        "partType": "flash memory",
        "manufacturer": "ACME Components",
        "componentName": "abc456",
        "orderableMPN": "abc456,56",
        "sourceDatasheetID": {
            "publishedDate": "2022-05-30",
            "datasheetURI": "www.acmecomponents.com/abc456datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2023-05-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "capacity": {
        "typValue": 64,
        "siUnit": "bit",
        "unitText": "Mb",
        "unitFactor": 1000000,
        "valueDefined": true
    },
    "pageSize": {
        "typValue": 256,
        "siUnit": "bytes",
        "unitText": "B",
        "unitFactor": 1,
        "valueDefined": true
    },
    "blockSize": {
        "typValue": 64,
        "siUnit": "bytes",
        "unitText": "KB",
        "unitFactor": 1000,
        "valueDefined": true
    },
    "interface": "spi",
    "clockFrequency": {
        "typValue": 1,
        "siUnit": "hertz",
        "unitText": "MHz",
        "unitFactor": 1000000,
        "valueDefined": true
    },
    "blockEraseTime": {
        "typValue": 100,
        "siUnit": "second",
        "unitText": "ms",
        "unitFactor": 0.001,
        "valueDefined": true
    },
    "pageProgramTime": {
        "typValue": 250,
        "siUnit": "second",
        "unitText": "us",
        "unitFactor": 0.000001,
        "valueDefined": true
    },
    "writeProtect": true,
    "package": {
        "length": 5,
        "width": 5,
        "height": 1,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "VDD"
        },
        {
            "terminalIdentifier": "2",
            "name": "VSS"
        },
        {
            "terminalIdentifier": "3",
            "name": "SDI"
        },
        {
            "terminalIdentifier": "4",
            "name": "SDO"
        },
        {
            "terminalIdentifier": "5",
            "name": "CS"
        },
        {
            "terminalIdentifier": "6",
            "name": "SCLK"
        },
        {
            "terminalIdentifier": "7",
            "name": "RST"
        },
        {
            "terminalIdentifier": "8",
            "name": "WP#"
        }
    ]
}
```

LED Example - Limited pins 

```json
{
    "componentID": {
        "partType": "LED",
        "manufacturer": "ACME Components",
        "componentName": "abc4567",
        "orderableMPN": "abc4567,567",
        "sourceDatasheetID": {
            "publishedDate": "2022-05-30",
            "datasheetURI": "www.acmecomponents.com/abc4567datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2023-05-26",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "ledColor": "red",
    "vf": {
        "value": {
            "typValue": 2,
            "maxValue": 3,
            "siUnit": "volt",
            "unitText": "V",
            "unitFactor": 1,
            "valueDefined": true
        },
        "conditions": {
            "if": {
                "typValue": 15,
                "siUnit": "amp",
                "unitText": "mA",
                "unitFactor": 0.001,
                "valueDefined": true
            }
        }
    },
    "ifp": {
        "maxValue": 200,
        "siUnit": "amp",
        "unitText": "mA",
        "unitFactor": 0.001,
        "valueDefined": true
    },
    "vr": {
        "maxValue": 5,
        "siUnit": "volt",
        "unitText": "V",
        "unitFactor": 1,
        "valueDefined": true
    },
    "peakWavelength": {
        "value": {
            "typValue": 620,
            "siUnit": "m",
            "unitText": "nm",
            "unitFactor": 1e-9,
            "valueDefined": true
        },
        "conditions": {
            "if": {
                "typValue": 15,
                "siUnit": "amp",
                "unitText": "mA",
                "unitFactor": 0.001,
                "valueDefined": true
            }
        }
    },
    "angleHalfIntensity": {
        "typValue": 25,
        "siUnit": "degrees",
        "unitText": "degrees",
        "unitFactor": 1,
        "valueDefined": true
    },
    "package": {
        "length": 2,
        "width": 2,
        "height": 1,
        "dimensionUnit": "millimeter"
    },
    "pins": [
        {
            "terminalIdentifier": "1",
            "name": "anode"
        },
        {
            "terminalIdentifier": "2",
            "name": "cathode"
        }
    ]
}
```

MCU Example - Limited pins, multiple files

stm32f205RBT6.json (main component file)
```json
{
    "componentID": {
        "partType": "ic micro",
        "manufacturer": "ExampleManufacturer",
        "componentName": "stm32f205RBT6",
        "orderableMPN": [
            "stm32f205RBT6xxxTR",
            "stm32f205RBT6xxx"
        ],
        "sourceDatasheetID": {
            "publishedDate": "2021-05-23",
            "datasheetURI": "www.Examplemanufacturer.com/stm32datasheet"
        },
        "digitalDatasheetID": {
            "publishedDate": "2022-03-21",
            "guid": "3e4cd9de-657a-41ae-902e-beca95aff51d"
        },
        "status": "active"
    },
    "datasheetFiles": {
        "coreProperties": {
            "filedescription": "core specifications for stm32f205",
            "fileType": "json",
            "fileExtension": "json",
            "fileURI": "stm32f205_core.json"
        },
        "additionalCoreProperties": {
            "filedescription": "additional property list for parts following the part number: stm32f205xBx6",
            "fileType": "json",
            "fileExtension": "json",
            "fileURI": "stm32f205xBx6.json"
        },
        "pins": {
            "filedescription": "pin specifications for 64 pins",
            "fileType": "json",
            "fileExtension": "json",
            "fileURI": "stm32f205R_pin.json"
        },
        "package": {
            "filedescription": "package specifications for 64 pin LQFP",
            "fileType": "json",
            "fileExtension": "json",
            "fileURI": "stm32f205xxT_package.json"
        }
    }
}
```
stm32f205_core.json (core properties file)
```json
{
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
    }
}

```

stm32f205xBx6.json (additional properties file)
```json
{
    "onChipFlash": {
        "siUnit": "byte",
        "absoluteValue": 128,
        "unitText": "KB",
        "unitFactor": 1e6,
        "valueDefined": true
    }
}

```

stm32f205R_pin.json (pin file)
```json
{
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
            }
        }
    ]
}

```

stm32f205xxT_package (package file)
```json
{
    "length": {
        "siUnit": "millimeter",
        "absoluteValue": 10,
        "unitText": "mm",
    },
    "width": {
        "siUnit": "millimeter",
        "absoluteValue": 10,
        "unitText": "mm",
    },
    "height": {
        "siUnit": "millimeter",
        "absoluteValue": 1.6,
        "unitText": "mm",
    },
    "partModelInformation":  {
        "filedescription": "package specifications for 64 pin LQFP",
        "fileType": "xml",
        "fileExtension": "xml",
        "standardReferenced": "IPC2552",
        "fileURI": "lqfp64.xml"
    }
}

```
    
        
    
        
    
    
             
    
    
    
     
    
        
        
    
    
