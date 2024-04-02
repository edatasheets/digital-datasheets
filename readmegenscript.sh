#!/bin/bash

CUR=5
for A in common/ clock/  data_converter/  hardware/  ic_io/  ic_microcontroller/  ic_misc/  logic/  memory/  passives/  power/  semiconductor/  sensor/  storage/; do
    go run readmescript.go  part-spec/$A 4.$CUR
    ((CUR++))
done  > tempREADME.md


