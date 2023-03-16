FROM mongo

RUN mkdir /data/key
COPY ./key /data/key
RUN chown mongodb:mongodb -R /data/key &&\
    chmod -R 600 /data/key &&\
    chmod 700 /data/key

CMD ["mongod"]