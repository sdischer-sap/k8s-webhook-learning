#!/usr/bin/env bash
# Generate a private key
openssl genrsa -out webhook-server-tls.key 2048

# Generate a CSR (Certificate Signing Request)
openssl req -new -key webhook-server-tls.key -out webhook-server-tls.csr -subj "/CN=host.docker.internal"

# Add Subject Alternative Name (SAN)
cat > webhook-server-tls.ext << EOF
subjectAltName = DNS:host.docker.internal
EOF

# Generate the certificate
openssl x509 -req -in webhook-server-tls.csr -signkey webhook-server-tls.key -out webhook-server-tls.crt -days 365 -extfile webhook-server-tls.ext

# Get the CA bundle for the webhook configuration
caBundle=$(cat webhook-server-tls.crt | base64 | tr -d '\n')

# If you know the exact folder name needed
SPECIFIC_NAME="k8s-webhook-server/serving-certs/"
CERT_DIR="$(dirname $(mktemp -u))/$SPECIFIC_NAME" && \
mkdir -p "$CERT_DIR" && \
#mv tls.key tls.crt "$CERT_DIR/" && \
echo "Certificates moved to $CERT_DIR"
