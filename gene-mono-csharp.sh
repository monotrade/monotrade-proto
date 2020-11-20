#! /bin/bash
cd src
mkdir -p ../../../MMDashboard/src/Monotrade.API/Proto/
protoc --csharp_out=../../../MMDashboard/src/Monotrade.API/Proto *.proto
cd ..
