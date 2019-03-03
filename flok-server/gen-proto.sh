#!/bin/bash
cd auth-srv
make proto
cd ../device-srv
make proto
cd ../hub-connector-srv
make proto
cd ../hub-srv
make proto
cd ../user-srv
make proto