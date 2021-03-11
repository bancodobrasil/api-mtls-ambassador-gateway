FROM local/api-mtls-ambassador-client

COPY client-certs/client.cert.pem /etc/nginx/conf.d/certs/client.cert.pem
COPY client-certs/client.key.pem /etc/nginx/conf.d/certs/client.key.pem
COPY client-certs/trusted_ca_cert.pem /etc/nginx/conf.d/certs/trusted_ca_cert.pem