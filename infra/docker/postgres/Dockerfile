FROM postgres:16

RUN apt-get update \
 && apt-get install -y postgresql-16-cron \
 && rm -rf /var/lib/apt/lists/*

# Passa os parâmetros via ENTRYPOINT para não sobrescrever
# a inicialização padrão da imagem oficial do Postgres
ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["postgres", "-c", "shared_preload_libraries=pg_cron", "-c", "cron.database_name=voting_db"]