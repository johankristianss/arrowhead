curl -X POST "https://localhost:8443/serviceregistry/mgmt/systems" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"address\": \"localhost\", \"port\": 8080, \"systemName\": \"cardemoconsumer\"}" \
     --cert ./sysop.p12:123456 \
     --cert-type P12 \
     --cacert ./truststore.pem


