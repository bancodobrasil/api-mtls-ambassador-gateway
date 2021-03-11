FROM nginx:1.18-alpine

EXPOSE 80

VOLUME [ "/etc/nginx/conf.d/certs" ]

ENV PROXY_PASS ''

COPY nginx.conf /nginx.template

CMD ["/bin/sh","-c", "envsubst '${PROXY_PASS} ${PROXY_SSL_VERIFY}' < /nginx.template > /etc/nginx/conf.d/default.conf; nginx-debug -g 'daemon off;'"]