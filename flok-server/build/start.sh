#!/bin/bash
./auth-srv.out &
./chat-srv.out &
./file-srv.out &
./user-srv.out &
./meeting-srv.out &
./project-srv.out &
./portal-srv.out &
# ./api-srv.out &
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
