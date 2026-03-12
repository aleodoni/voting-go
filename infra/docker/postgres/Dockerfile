FROM postgres:16

RUN apt-get update \
 && apt-get install -y postgresql-16-cron

CMD ["postgres", "-c", "shared_preload_libraries=pg_cron", "-c", "cron.database_name=voting_db"]