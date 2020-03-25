#!/bin/sh

reflex -G 'temp-images/upload-*.png' -G '*.db' -G '*.db-journal' -s -- sh -c 'clear; sleep 1s; go run ./cmd/web'
