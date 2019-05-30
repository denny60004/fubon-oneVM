#!/bin/bash
set -u
set -e

basepath=$(cd `dirname $0`; cd ..; pwd)
quorum=$basepath'/quorum/build/bin/geth'
crux=$basepath'/crux/bin/crux'
echo "[*] start crux1"
CMD="$crux --url=http://127.0.0.1:9001/ --port=9001 --networkinterface 0.0.0.0 --workdir=qdata/c1 --socket=tm.ipc --publickeys=tm1.pub --privatekeys=tm1.key --othernodes=http://127.0.0.1:9002 --verbosity=3"
$CMD >> "qdata/logs/crux1.log" 2>&1 &

echo "[*] start crux2"
CMD="$crux --url=http://127.0.0.1:9002/ --port=9002 --networkinterface 0.0.0.0 --workdir=qdata/c2 --socket=tm.ipc --publickeys=tm2.pub --privatekeys=tm2.key --othernodes=http://127.0.0.1:9003 --verbosity=3"
$CMD >> "qdata/logs/crux2.log" 2>&1 &

echo "[*] start crux3"
CMD="$crux --url=http://127.0.0.1:9003/ --port=9003 --networkinterface 0.0.0.0 --workdir=qdata/c3 --socket=tm.ipc --publickeys=tm3.pub --privatekeys=tm3.key --othernodes=http://127.0.0.1:9004 --verbosity=3"
$CMD >> "qdata/logs/crux3.log" 2>&1 &

echo "[*] start crux4"
CMD="$crux --url=http://127.0.0.1:9004/ --port=9004 --networkinterface 0.0.0.0 --workdir=qdata/c4 --socket=tm.ipc --publickeys=tm4.pub --privatekeys=tm4.key --othernodes=http://127.0.0.1:9001 --verbosity=3"
$CMD >> "qdata/logs/crux4.log" 2>&1 &

# DDIR="qdata/c"
# mkdir -p $DDIR
# mkdir -p qdata/logs
# echo $CRUX_PUB >> "$DDIR/tm.pub"
# echo $CRUX_PRIV >> "$DDIR/tm.key"
# rm -f "$DDIR/tm.ipc"
# CMD="crux --url=http://$OWN_URL:$PORT/ --port=$PORT --networkinterface 0.0.0.0 --workdir=$DDIR --socket=tm.ipc --publickeys=tm.pub --privatekeys=tm.key --othernodes=$OTHER_NODES --verbosity=3"
# $CMD >> "qdata/logs/crux.log" 2>&1 &

# DOWN=true
# while $DOWN; do
#     sleep 0.1
#     DOWN=false
# 	if [ ! -S "qdata/c/tm.ipc" ]; then
#             DOWN=true
# 	fi
# done
