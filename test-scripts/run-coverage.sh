#! /bin/bash

export FULLNODE_API_INFO=/ip4/10.0.1.52/tcp/1234/http

cd ..

gotestsum -- \
  -tags clientqueryask \
  -coverprofile=coverage.txt \
  -coverpkg=github.com/filecoin-project/lotus/...,github.com/filecoin-project/go-fil-markets/... \
  .
go tool cover -html=coverage.txt -o coverage.html
