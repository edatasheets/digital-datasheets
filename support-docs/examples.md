Single Inverter Example

```json
{
    "partType": "inverter",
    "manufacturer": "ACME Components",
    "mpn": "abc123",
    "datasheetVersion": "v0.3",
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
            "siUnit": "ohms",
            "absoluteValue": 20,
            "unitText": "Kohms",
            "unitFactor": 1E3,
            "valueDefined": true
        },
    }
}

```