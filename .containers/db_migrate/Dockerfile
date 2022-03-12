FROM migrate/migrate:v4.15.1

RUN mkdir /db

COPY ./.containers/db_migrate/scripts/ /opt/bin
RUN chmod -R +x /opt/bin

ENTRYPOINT [ "/opt/bin/entrypoint.sh" ]
