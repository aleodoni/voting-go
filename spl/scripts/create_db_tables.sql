CREATE DATABASE spl;

\c spl

CREATE SCHEMA IF NOT EXISTS spl;

CREATE TABLE IF NOT EXISTS spl.proposicoes_pauta_reuniao (
    pac_id integer,
    par_id integer,
    link_prop text,
    codigo_proposicao character varying(23),
    iniciativa text,
    sumula text,
    comissao character varying,
    paralelo boolean,
    relator character varying(100),
    conclusao_relator character varying(40),
    conclusao_comissao character varying,
    link_texto text,
    tem_emendas boolean,
    prazo_comissao_fim date
);

CREATE TABLE IF NOT EXISTS spl.texto_conclusao_autor (
    pro_id integer,
    pro_codigo character varying(23),
    par_id integer,
    txt_data date,
    txt_finalizado boolean,
    txt_relator boolean,
    txt_id integer,
    ver_id integer,
    vereador character varying(100),
    tcp_id integer,
    tcp_nome character varying(40),
    par_finalizado boolean,
    con_id integer
);

CREATE TABLE IF NOT EXISTS spl.reuniao_comissao (
    rec_id integer NOT NULL,
    rec_tipo_reuniao character varying(30) NOT NULL,
    rec_numero character varying(40) NOT NULL,
    versao integer NOT NULL,
    rec_data date NOT NULL,
    rec_notificada boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS spl.reuniao_comissao_conjunto (
    rec_id integer NOT NULL,
    con_id integer NOT NULL,
    versao integer NOT NULL
);

CREATE TABLE IF NOT EXISTS spl.conjunto_vereadores (
    con_id integer NOT NULL,
    con_sigla character varying(20),
    con_maximo integer,
    ut_id integer,
    ini_nome text,
    ini_ativa boolean NOT NULL,
    tin_id integer NOT NULL,
    versao integer NOT NULL,
    ini_codigo_prefeitura integer,
    con_descricao text,
    con_site_ordem integer,
    con_email character varying(120)
);

CREATE TABLE IF NOT EXISTS spl.pauta_comissao (
    pac_id integer NOT NULL,
    rec_id integer,
    pac_liberada boolean NOT NULL,
    pac_notificada boolean NOT NULL,
    versao integer NOT NULL,
    pac_texto text
);

CREATE TABLE IF NOT EXISTS spl.vereador (
    ver_id integer NOT NULL primary key,
    ver_sexo character(1) NOT NULL,
    ut_id integer,
    ini_nome character varying(100) NOT NULL,
    ini_ativa boolean NOT NULL,
    ver_nome_completo character varying(100) NOT NULL,
    versao integer NOT NULL,
    ini_codigo_prefeitura integer,
    ver_site character varying(250),
    ver_biografia text,
    ver_redes_sociais character varying(250),
    ver_fone_principal character varying(20),
    ver_fones character varying(100),
    ver_legislaturas character varying(100),
    ver_localizacao character varying(50),
    ver_partido character varying(100),
    arq_id integer,
    arq_id_biografia integer
);

CREATE TABLE IF NOT EXISTS spl.usuario (
    usu_id integer,
    usu_nome character varying(100),
    usu_login character varying(30),
    usu_matricula character varying(10),
    usu_email character varying(120),
    usu_compositor_universal boolean,
    usu_recebe_diario boolean,
    usu_recebe_od boolean,
    usu_tipo character varying(10),
    usu_ativo boolean,
    ver_id integer,
    usu_recebe_aviso_aprovacao boolean,
    versao integer
);

CREATE TABLE IF NOT EXISTS spl.composicao (
    com_id integer NOT NULL,
    con_id integer NOT NULL,
    ver_id integer NOT NULL,
    crg_id integer NOT NULL,
    versao integer NOT NULL,
    com_ordem integer
);

CREATE TABLE IF NOT EXISTS spl.cargo_vereador (
    crg_id integer NOT NULL,
    crg_nome character varying(50) NOT NULL,
    versao integer NOT NULL,
    crg_ordem integer
);

CREATE TABLE IF NOT EXISTS spl.tipo_iniciativa (
    tin_id integer NOT NULL,
    tin_nome character varying(100) NOT NULL,
    tin_multipla boolean NOT NULL,
    versao integer NOT NULL,
    tin_adm_leg character(1) NOT NULL
);

CREATE DOMAIN spl.tipo_voto_parecer AS character varying(25)
    COLLATE "default"
    CONSTRAINT voto_parecer_constraint
    CHECK (
        VALUE::text = 'Favorável'::text OR VALUE::text = 'Favorável com restrições'::text OR VALUE::text = 'Contrário'::text OR VALUE::text = 'Abstenção'::text
    );

CREATE TABLE
    IF NOT EXISTS spl.voto_parecer (
        vot_id integer NOT NULL primary key,
        par_id integer NOT NULL,
        ver_id integer REFERENCES spl.vereador (ver_id),
        vot_voto spl.tipo_voto_parecer NULL,
        versao integer DEFAULT 0 NOT NULL
    );

ALTER TABLE spl.reuniao_comissao OWNER TO CURRENT_USER;
ALTER TABLE spl.reuniao_comissao_conjunto OWNER TO CURRENT_USER;
ALTER TABLE spl.conjunto_vereadores OWNER TO CURRENT_USER;
ALTER TABLE spl.pauta_comissao OWNER TO CURRENT_USER;
ALTER VIEW spl.v_votacao_reunioes2 OWNER TO CURRENT_USER;

GRANT SELECT ON ALL TABLES IN SCHEMA spl TO PUBLIC;