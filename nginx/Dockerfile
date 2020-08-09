FROM nginx:alpine

WORKDIR /opt/src

RUN apk add --no-cache bash openssl \
    && rm /etc/nginx/conf.d/default.conf \
    && echo "daemon off;" >> /etc/nginx/nginx.conf

COPY ./conf.d/code-education.conf /opt/src/conf.d/code-education.conf
COPY ./domains/codeeducation/index.html /var/www/codeeducation.com.br/index.html
COPY ./boot.sh .

RUN ln -s /opt/src/conf.d/code-education.conf /etc/nginx/conf.d/code-education.conf

EXPOSE 80

ENTRYPOINT [ "./boot.sh" ]