curl -X POST "https://localhost:8445/authorization/mgmt/intracloud" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"consumerId\": 8, \"interfaceIds\": [ 1 ], \"providerIds\": [ 7 ], \"serviceDefinitionIds\": [ 36,37 ]}" \
     --cert ./sysop.p12:123456 \
     --cert-type P12 \
     --cacert ./truststore.pem
