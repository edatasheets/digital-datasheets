#!/bin/bash

CUR=1
for A in common/ clock/   component.json/  data_converter/  hardware/  ic_io/  ic_microcontroller/  ic_misc/  logic/  memory/  passives/  power/  semiconductor/  sensor/  storage/; do
    go run readmescript.go  part-spec/$A 3.$CUR
    ((CUR++))
done  > tempREADME.md


