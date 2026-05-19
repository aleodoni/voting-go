-- spl.tipo_iniciativa
INSERT INTO spl.tipo_iniciativa (tin_id, tin_nome, tin_multipla, versao, tin_adm_leg)
VALUES (1, 'Mesa Diretora', false, 1, 'A');

-- spl.conjunto_vereadores
INSERT INTO spl.conjunto_vereadores (con_id, con_sigla, con_maximo, ut_id, ini_nome, ini_ativa, tin_id, versao, ini_codigo_prefeitura, con_descricao, con_site_ordem, con_email)
VALUES (6136, 'C.UrbanismoTI', 5, 100, 'Comissão de Urbanismo, Obras Públicas e TI', true, 1, 1, 999, 'C.Meio Ambiente', 1, 'conjunto@email.com');

INSERT INTO spl.conjunto_vereadores (con_id, con_sigla, con_maximo, ut_id, ini_nome, ini_ativa, tin_id, versao, ini_codigo_prefeitura, con_descricao, con_site_ordem, con_email)
VALUES (69, 'C.Economia', 5, 100, 'Comissão de Economia, Finanças e Fiscalização', true, 1, 1, 999, 'C.Economia', 1, 'conjunto@email.com');

INSERT INTO spl.conjunto_vereadores (con_id, con_sigla, con_maximo, ut_id, ini_nome, ini_ativa, tin_id, versao, ini_codigo_prefeitura, con_descricao, con_site_ordem, con_email)
VALUES (9066, 'C.Meio Ambiente', 5, 100, 'Comissão de Meio Ambiente, Desenvolvimento Sustentável e Assuntos Metropolitanos', true, 1, 1, 999, 'C.Meio Ambiente', 1, 'conjunto@email.com');

INSERT INTO spl.conjunto_vereadores (con_id, con_sigla, con_maximo, ut_id, ini_nome, ini_ativa, tin_id, versao, ini_codigo_prefeitura, con_descricao, con_site_ordem, con_email)
VALUES (1234, 'C.Teste ConjuntaA', 5, 100, 'Comissão teste A para reunião conjunta', true, 1, 1, 999, 'C.Teste ConjuntaA', 1, 'conjuntoA@email.com');

INSERT INTO spl.conjunto_vereadores (con_id, con_sigla, con_maximo, ut_id, ini_nome, ini_ativa, tin_id, versao, ini_codigo_prefeitura, con_descricao, con_site_ordem, con_email)
VALUES (1235, 'C.Teste ConjuntaB', 5, 100, 'Comissão teste B para reunião conjunta', true, 1, 1, 999, 'C.Teste ConjuntaB', 1, 'conjuntoB@email.com');


-- spl.composicao
INSERT INTO spl.composicao (com_id, con_id, ver_id, crg_id, versao, com_ordem)
VALUES (1, 1, 1, 1, 1, 1);
INSERT INTO spl.composicao (com_id, con_id, ver_id, crg_id, versao, com_ordem)
VALUES (2, 6136, 2, 1, 1, 1);

-- spl.pauta_comissao
INSERT INTO spl.pauta_comissao (pac_id, rec_id, pac_liberada, pac_notificada, versao, pac_texto)
VALUES (1592, 3877, true, true, 1, 'Texto da pauta para a reunião 14ª 3877');

INSERT INTO spl.pauta_comissao (pac_id, rec_id, pac_liberada, pac_notificada, versao, pac_texto)
VALUES (1593, 3878, true, true, 1, 'Texto da pauta para a reunião 29ª, 14h, no Plenário 3878');

INSERT INTO spl.pauta_comissao (pac_id, rec_id, pac_liberada, pac_notificada, versao, pac_texto)
VALUES (1594, 3879, true, true, 1, 'Texto da pauta para a reunião 11ª, APÓS A SESSÃO PLENÁRIA 3879');

INSERT INTO spl.pauta_comissao (pac_id, rec_id, pac_liberada, pac_notificada, versao, pac_texto)
VALUES (1595, 3880, true, true, 1, 'Texto da pauta para a reunião Conjunta 3880');

-- spl.reuniao_comissao
INSERT INTO spl.reuniao_comissao (rec_id, rec_tipo_reuniao, rec_numero, versao, rec_data, rec_notificada)
VALUES (3877, 'Ordinária', '14ª, às 08:00, Videoconferência', 1, CURRENT_DATE, true);

INSERT INTO spl.reuniao_comissao (rec_id, rec_tipo_reuniao, rec_numero, versao, rec_data, rec_notificada)
VALUES (3878, 'Ordinária', '29ª, 14h, no Plenário', 1, CURRENT_DATE, true);

INSERT INTO spl.reuniao_comissao (rec_id, rec_tipo_reuniao, rec_numero, versao, rec_data, rec_notificada)
VALUES (3879, 'Ordinária', '11ª, APÓS A SESSÃO PLENÁRIA', 1, CURRENT_DATE, true);

INSERT INTO spl.reuniao_comissao (rec_id, rec_tipo_reuniao, rec_numero, versao, rec_data, rec_notificada)
VALUES (3880, 'Ordinária', '13ª, Teste Conjunta', 1, CURRENT_DATE, true);

-- spl.reuniao_comissao_conjunto
INSERT INTO spl.reuniao_comissao_conjunto (rec_id, con_id, versao)
VALUES (3877, 6136, 1);

INSERT INTO spl.reuniao_comissao_conjunto (rec_id, con_id, versao)
VALUES (3878, 69, 1);

INSERT INTO spl.reuniao_comissao_conjunto (rec_id, con_id, versao)
VALUES (3879, 9066, 1);

INSERT INTO spl.reuniao_comissao_conjunto (rec_id, con_id, versao)
VALUES (3880, 1234, 1);

INSERT INTO spl.reuniao_comissao_conjunto (rec_id, con_id, versao)
VALUES (3880, 1235, 1);

-- spl.texto_conclusao_autor
INSERT INTO spl.texto_conclusao_autor (
    pro_id, pro_codigo, par_id, txt_data, txt_finalizado,
    txt_relator, txt_id, ver_id, vereador, tcp_id, tcp_nome,
    par_finalizado, con_id
) VALUES 
(
    0, '005.00147.2022', 31868, '2024-05-01', true,
    false, 38195, 1, 'Pier Petruzziello', 0, 'Por mais informações',
    true, 1
),
(
    0, '005.00147.2022', 31868, '2024-05-01', true,
    false, 37933, 1, 'Dalton Borba', 0, 'Pela tramitação',
    true, 1
),
(
    0, '005.00052.2022', 31901, '2024-05-01', true,
    false, 38192, 1, 'Beto Moraes', 0, 'Pela anexação',
    true, 1
);

-- spl.proposicoes_pauta_reuniao
INSERT INTO spl.proposicoes_pauta_reuniao (
    pac_id, par_id, link_prop, codigo_proposicao, iniciativa, sumula,
    comissao, paralelo, relator, conclusao_relator, conclusao_comissao,
    link_texto, tem_emendas, prazo_comissao_fim
) 
VALUES (
    1592, 
    31868, 
    'https://linkprop.com', 
    '005.00147.2022', 
    'Prefeito', 
    'Autoriza o Poder Executivo a transferir à Companhia de Habitação Popular de Curitiba - COHAB-CT, a título de alienação gratuita, imóvel que especifica.',
    'Comissão A - não usado', 
    false, 
    'Mauro Bobato', 
    'Pela tramitação', 
    'Pela tramitação',
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1592, 
    31901, 
    'https://linkprop.com', 
    '005.00052.2022', 
    'Prefeito', 
    'Autoriza o Poder Executivo a alienar a favor do espólio de Osni Prates Pacheco, representado por Giselle Brunor Pacheco Ebrahim, a área que especifica.',
    'Comissão A - não usado', 
    false, 
    'Sidnei Toaldo', 
    'Pela tramitação', 
    'Pela tramitação',
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1592, 
    31902, 
    'https://linkprop.com', 
    '005.00151.2022', 
    'Dalton Borba', 
    'Acrescenta-se o parágrafo 2º e incisos, e renumera-se o parágrafo único do artigo 9º da Lei Municipal nº 8.670, de 29 de junho de 1995', 
    'Comissão A - não usado', 
    false, 
    'Toninho da Farmácia', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1592, 
    31924, 
    'https://linkprop.com', 
    '005.00059.2022', 
    'Prefeito', 
    'Suprime, desafeta e incorpora área de terreno aos bens dominicais e autoriza o Poder Executivo a alienar em favor de Marco Aurélio Ferla, Luciana Ferla, Claudia Ferla e Fernanda Ferla, a área que especifica.', 
    'Comissão A - não usado', 
    false, 
    'Hernani', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1592, 
    31925, 
    'https://linkprop.com', 
    '005.00070.2022', 
    'Jornalista Márcio Barros, Sargento Tania Guerreiro', 
    'Institui o controle da poluição sonora veicular no âmbito do Município de Curitiba e dá outras providências.', 
    'Comissão A - não usado', 
    false, 
    'Mauro Bobato', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1592, 
    31933, 
    'https://linkprop.com', 
    '005.00075.2022', 
    'Prefeito', 
    'Altera o art. 6º da Lei nº 11.095, de 21 de julho de 2004, que dispõe sobre as normas que regulam a aprovação de projetos, o licenciamento de obras e atividades, a execução, manutenção e conservação de obras no Município, e dá outras providências.', 
    'Comissão A - não usado', 
    false, 
    'Sidnei Toaldo', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1593, 
    31896, 
    'https://linkprop.com', 
    '005.00156.2022', 
    'Prefeito', 
    'Altera dispositivos da Lei nº 13.900, de 9 de dezembro de 2011, que cria o Conselho Municipal da Juventude.', 
    'Comissão A - não usado', 
    false, 
    'Tito Zeglin', 
    'Pela tramitação', 
    'Pela tramitação', 
    'RF-2103', 
    false, 
    '2024-06-01'
),
(
    1593, 
    31899, 
    'https://linkprop.com', 
    '005.00054.2021', 
    'Renato Freitas', 
    'Dispõe sobre a criação do programa "Passe Livre à Internet", para garantir acesso e navegação à internet de caráter gratuito aos estudantes do município de Curitiba.', 
    'Comissão A - não usado', 
    false, 
    'João da 5 Irmãos', 
    'Por mais informações', 
    '', 
    'https Obi-1', 
    false, 
    '2024-06-01'
),
(
    1593, 
    31940, 
    'https://linkprop.com', 
    '002.00008.2022', 
    'Prefeito', 
    'Dispõe sobre a Planta Genérica de Valores - PGV, altera dispositivos das Leis Complementares nºs 7/1993, 40/2001 e 44/2002 e revoga dispositivos das Leis Complementares nºs 53/2004 e 91/2014.', 
    'Comissão A - não usado', 
    false, 
    'Sergio R. B. Balaguer (Serginho do Posto)', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1594, 
    31941, 
    'https://linkprop.com', 
    '002.00009.2022', 
    'Prefeito', 
    'Propõe mudar nome da avenida Manoel Ribas para Manoel Ribeiro.', 
    'Comissão A - não usado', 
    false, 
    'Sergio R. B. Balaguer (Serginho do Posto)', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
),
(
    1595, 
    31500, 
    'https://linkprop.com', 
    '001.00001.2025', 
    'Prefeito', 
    'Propõe instituir o dia do analista de sistemas.', 
    'Comissão A - não usado', 
    true, 
    'Sergio R. B. Balaguer (Serginho do Posto)', 
    'Pela tramitação', 
    'Pela tramitação', 
    'https://linktexto.com', 
    false, 
    '2024-06-01'
);
