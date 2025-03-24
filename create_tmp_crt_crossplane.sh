#!/usr/bin/env bash
cat <<EOF > req.cnf
[req]
distinguished_name = req_distinguished_name
x509_extensions = v3_req
prompt = no
[req_distinguished_name]
C = US
ST = MD
O = home
localityName = home
commonName = *.crossplane-system.svc
organizationalUnitName = home
emailAddress = some@email.com
[v3_req]
keyUsage = keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names
[alt_names]
DNS.1   = *.crossplane-system.svc5
EOF

# Generate our Private Key, and Certificate directly
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 \
  -keyout "tls.key" -config req.cnf \
  -out "tls.crt" -sha256
rm req.cnf

# If you know the exact folder name needed
SPECIFIC_NAME="k8s-webhook-server/serving-certs/"
CERT_DIR="$(dirname $(mktemp -u))/$SPECIFIC_NAME" && \
mkdir -p "$CERT_DIR" && \
#mv tls.key tls.crt "$CERT_DIR/" && \
echo "Certificates moved to $CERT_DIR"
