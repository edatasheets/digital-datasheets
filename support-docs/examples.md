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
        "relativeValueOperator": "multiply"
        "valueDefined": true
    },
    "vihMin": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "VDD1",
        "relativeValueModifier": 0.2,
        "relativeValueOperator": "multiply"
        "valueDefined": true
    },
    "vol": "vihMin": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "GND",
        "relativeValueModifier": 0.1,
        "relativeValueOperator": "add"
        "valueDefined": true
    },
    "voh": "vihMin": {
        "siUnit": "volts",
        "unitText": "V",
        "unitFactor": 1,
        "relativeValueReference": "VDD1",
        "relativeValueModifier": 0.1,
        "relativeValueOperator": "subtract"
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
        },
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
        "xLabelUnits": "Amps",
        "xLabelTitle": "Current",
        "yLabelUnits": "uH",
        "yLabelTitle": "Inductance",
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
        "dimensionUnit": "milimeter" 
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

