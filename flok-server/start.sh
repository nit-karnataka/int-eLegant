#!/bin/bash
./auth-srv.out &
./device-srv.out &
./user-srv.out &
./hub-srv.out &
./hub-connector-srv.out &
./api-srv.out &
wait
# cd ../protocol-adapter
# go build -o ../protocol-adapter.out

# cd ../room-handler
# go build -o ../room-handler.out

# cd ../routine-engine
# go build -o ../routine-engine.out

# cd ../rule-engine
# go build -o ../rule-engine.out

# cd ../scene-handler
# go build -o ../scene-handler.out
