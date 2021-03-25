FROM labbsr0x/api-mtls-ambassador-client:0.0.2

COPY client-certs/client.pem /etc/nginx/conf.d/certs/client.pem
COPY client-certs/client-key.pem /etc/nginx/conf.d/certs/client-key.pem
COPY client-certs/server-ca.pem /etc/nginx/conf.d/certs/server-ca.pem