#! /bin/bash
cd src
protoc --csharp_out=../../../MMDashboard/src/Monotrade.API/Message *
cd ..
