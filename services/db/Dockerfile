FROM postgres

RUN mkdir -p /db
WORKDIR /db

COPY ./script/ /db/

RUN chmod a+x /db/script.sh /db/destroy.sh
ENV POSTGRES_PASSWORD=pinto
# ENV POSTGRES_HOST_AUTH_METHOD=trust

EXPOSE 5432
