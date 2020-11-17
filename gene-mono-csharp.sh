#! /bin/bash
cd src
mkdir -p ../../../MMDashboard/src/Monotrade.API/Message/
protoc --csharp_out=../../../MMDashboard/src/Monotrade.API/Message *.proto
cd ..
