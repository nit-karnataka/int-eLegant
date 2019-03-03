#!/bin/bash
cd api-srv
# go build -o ../build/api-srv.out

cd ../auth-srv
go build -o ../build/auth-srv.out

cd ../user-srv
go build -o ../build/user-srv.out

cd ../chat-srv
go build -o ../build/chat-srv.out

cd ../file-srv
go build -o ../build/file-srv.out

cd ../meeting-srv
go build -o ../build/meeting-srv.out

cd ../portal-srv
go build -o ../build/portal-srv.out

cd ../project-srv
go build -o ../build/project-srv.out
