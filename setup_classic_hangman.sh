#!/bin/bash

echo "cd src && ./../bin/classic_hangman words" > start_classique_hangman.sh
mkdir bin
cd src
go build -o ../bin/classic_hangman main.go