CREATE OR REPLACE VIEW spl.v_votacao_reunioes2 AS
SELECT 
  a.rec_data,
  a.versao,
  a.rec_id,
  a.rec_tipo_reuniao,
  a.rec_numero,
  b.ini_nome AS con_desc,
  b.con_id,
  c.pac_id,
  b.con_sigla
FROM spl.reuniao_comissao a
JOIN spl.reuniao_comissao_conjunto rcc using (rec_id)
JOIN spl.conjunto_vereadores b using (con_id)
JOIN spl.pauta_comissao c using (rec_id);

CREATE OR REPLACE VIEW spl.v_projetos_pareceres2
AS 
SELECT 
  texto_conclusao_autor.pro_codigo,
  texto_conclusao_autor.txt_id,
  texto_conclusao_autor.vereador,
  texto_conclusao_autor.tcp_nome
FROM spl.texto_conclusao_autor;

CREATE OR REPLACE VIEW spl.v_projetos_comissao2
AS 
SELECT 
  proposicoes_pauta_reuniao.sumula,
  proposicoes_pauta_reuniao.comissao,
  proposicoes_pauta_reuniao.paralelo,
  proposicoes_pauta_reuniao.relator,
  proposicoes_pauta_reuniao.tem_emendas,
  proposicoes_pauta_reuniao.pac_id,
  proposicoes_pauta_reuniao.par_id,
  proposicoes_pauta_reuniao.codigo_proposicao,
  proposicoes_pauta_reuniao.iniciativa,
  proposicoes_pauta_reuniao.conclusao_comissao,
  proposicoes_pauta_reuniao.conclusao_relator
FROM spl.proposicoes_pauta_reuniao;

GRANT SELECT ON spl.v_votacao_reunioes2 TO PUBLIC;
GRANT SELECT ON spl.v_projetos_pareceres2 TO PUBLIC;
GRANT SELECT ON spl.v_projetos_comissao2 TO PUBLIC;