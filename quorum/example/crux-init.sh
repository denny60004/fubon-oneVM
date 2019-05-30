#!/bin/bash
set -u
set -e

echo "[*] initialize crux1"
mkdir -p qdata/c1
cp keys/tm1.pub qdata/c1
cp keys/tm1.key qdata/c1

echo "[*] initialize crux2"
mkdir -p qdata/c2
cp keys/tm2.pub qdata/c2
cp keys/tm2.key qdata/c2

echo "[*] initialize crux3"
mkdir -p qdata/c3
cp keys/tm3.pub qdata/c3
cp keys/tm3.key qdata/c3

echo "[*] initialize crux4"
mkdir -p qdata/c4
cp keys/tm4.pub qdata/c4
cp keys/tm4.key qdata/c4