curl -X POST "https://127.0.0.1:8441/orchestrator/orchestration" \
     -H "Content-Type: application/json" \
	 --cert ./cardemoconsumer.p12:123456 \
     --cert-type P12 \
     --cacert ./truststore.pem \
     -d '{
  "requesterSystem": {
    "systemName": "cardemoconsumer",
    "address": "localhost",
    "port": 8080,
    "authenticationInfo": "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAws6zCv0nPnL8lAqSGBB8rMokDSrUEBTrMw1tb0IIsgkH1pbreBhkxyhPlDC7OstPiS6T2DZfo15Fz0ib8QmxJ81s4SJqfeh527DVrEgoDXKBemPzipfSQNBRP1jGDo7RDmNvNMtuQ50rWtMvMJJoo+jOxsS+a0pscr4Te5huaQQnLPDemjFA1YjE38BL3vEsdWt0mL4ueO819wj+BVpUcXQkNspDrRStzh3m8MAsdIzN/KPrPPsBlw8GlhCy+YAxSQ4+uA3M3wYN+FkZd77iNQZP4OgK9P3f/+6/ul2N9QCf0UvQy77JnE/DczXdKz+9y36wU88T8UEaHBaJp8sm6wIDAQAB",
    "metadata": null
  },
  "requesterCloud": null,
  "requestedService": {
    "serviceDefinitionRequirement": "create-car",
    "interfaceRequirements": [ "HTTP-SECURE-JSON" ],
    "securityRequirements": null,
    "metadataRequirements": null,
    "versionRequirement": null,
    "minVersionRequirement": null,
    "maxVersionRequirement": null,
    "pingProviders": false
  },
  "orchestrationFlags": {
    "onlyPreferred": false,
    "overrideStore": true,
    "externalServiceRequest": false,
    "enableInterCloud": false,
    "enableQoS": false,
    "matchmaking": true,
    "metadataSearch": false,
    "triggerInterCloud": false,
    "pingProviders": false
  },
  "preferredProviders": [],
  "commands": {},
  "qosRequirements": {}
}'
