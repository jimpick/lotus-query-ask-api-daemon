#! /bin/bash

export FULLNODE_API_INFO=/ip4/10.0.1.52/tcp/1234/http

cd ..

go run -tags clientqueryask . daemon
