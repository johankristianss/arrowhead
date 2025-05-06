#!/bin/bash

get_my_public_key() {
    SSL_ENABLED=true  # Set dynamically if needed

    if [ "$SSL_ENABLED" = true ]; then
        if [ -f "./cardemoprovider.p12" ]; then
            # Extract public key from certificate and clean output
            PUBLIC_KEY=$(openssl pkcs12 -in cardemoprovider.p12 -clcerts -nokeys -passin pass:123456 2>/dev/null | \
                         openssl x509 -pubkey -noout 2>/dev/null | \
                         sed -e '1d' -e '$d' | tr -d '\n')

            echo "$PUBLIC_KEY"  # Return clean string
        else
            echo "Error: SSL is enabled, but cardemoprovider.p12 not found." >&2
            return 1
        fi
    else
        echo ""  # Return empty string if SSL is disabled
    fi
}

# Get the public key
PUB_KEY=$(get_my_public_key)

# Check if extraction was successful
if [ -z "$PUB_KEY" ]; then
    echo "Failed to extract public key. Exiting..."
    exit 1
fi

# Create JSON payload
JSON_PAYLOAD=$(cat <<EOF
{
  "endOfValidity": null,
  "interfaces": ["HTTP-SECURE-JSON"],
  "metadata": {
    "http-method": "POST"
  },
  "providerSystem": {
    "address": "127.0.0.1",
    "authenticationInfo": "$PUB_KEY",
    "metadata": null,
    "port": 8888,
    "systemName": "cardemoprovider"
  },
  "secure": "TOKEN",
  "serviceDefinition": "create-car",
  "serviceUri": "/car",
  "version": null
}
EOF
)

# Send POST request to Service Registry
curl -X POST "https://localhost:8443/serviceregistry/register" \
    -H "Content-Type: application/json" \
    --cert ./cardemoprovider.p12:123456 \
    --cert-type P12 \
    --cacert ./truststore.pem \
    -d "$JSON_PAYLOAD"
