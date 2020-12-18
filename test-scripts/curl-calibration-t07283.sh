#! /bin/bash

export DATA='{"jsonrpc":"2.0","id":2,"method":"Filecoin.ClientQueryAsk","params":["12D3KooWAU1x4P8XGCWyQBAapXXoGyom4Bx5QHnH4zQSeTqzMyQP","t07283"]}'
echo $DATA | jq .

curl -X POST -H "Content-Type: application/json" \
       	--data "$DATA" \
       	'http://127.0.0.1:9301/rpc/v0' 

