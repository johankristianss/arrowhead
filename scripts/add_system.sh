curl -X POST "https://10.0.0.200:8443/serviceregistry/mgmt/systems" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"address\": \"localhost\", \"port\": 8080, \"systemName\": \"carconsumerdemo\"}"

curl -X POST "https://10.0.0.200:8445/authorization/mgmt/intracloud" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"consumerId\": 9, \"interfaceIds\": [ 1 ], \"providerIds\": [ 8 ], \"serviceDefinitionIds\": [ 36,37 ]}"
