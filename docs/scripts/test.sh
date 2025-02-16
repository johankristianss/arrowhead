curl -X GET "https://127.0.0.1:8443/serviceregistry/mgmt?direction=ASC&sort_field=id" \
     -H "accept: */*" \
     --cert ./sysop.p12:123456 \
     --cert-type P12 \
     --cacert ./truststore.pem
