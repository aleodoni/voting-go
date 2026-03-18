DELETE FROM votacao;
DELETE FROM reuniao;
-- DELETE FROM usuario;

INSERT INTO reuniao(
  id, 
  con_id, 
  con_desc, 
  rec_id, 
  con_sigla, 
  rec_tipo_reuniao, 
  rec_numero, 
  pac_id, 
  rec_data, 
  created_at, 
  updated_at
) VALUES (
  'gdmo94u77a080zyi71nlelwv',
  6136,
  'Comissão de Urbanismo, Obras Públicas e TI',
  3877,
  'C.UrbanismoTI',
  'Ordinária',
  '14ª, às 08:00, Videoconferência',
  1592,
  NOW()::date,
  NOW(),
  NOW()
);

INSERT INTO reuniao(
  id, 
  con_id, 
  con_desc, 
  rec_id, 
  con_sigla, 
  rec_tipo_reuniao, 
  rec_numero, 
  pac_id, 
  rec_data, 
  created_at, 
  updated_at
) VALUES (
  'd3xhufegbwe1mdldxr9pofng',
  69,
  'Comissão de Economia, Finanças e Fiscalização',
  3878,
  'C.Economia',
  'Ordinária',
  '29ª, 14h, no Plenário',
  1593,
  NOW()::date,
  NOW(),
  NOW()
);

INSERT INTO reuniao(
  id, 
  con_id, 
  con_desc, 
  rec_id, 
  con_sigla, 
  rec_tipo_reuniao, 
  rec_numero, 
  pac_id, 
  rec_data, 
  created_at, 
  updated_at
) VALUES (
  'xfdisna9nwtkxe5fjveasn9b',
  9066,
  'Comissão de Meio Ambiente, Desenvolvimento Sustentável e Assuntos Metropolitanos',
  3879,
  'C.Meio Ambiente',
  'Ordinária',
  '11ª, APÓS A SESSÃO PLENÁRIA',
  1594,
  NOW()::date,
  NOW(),
  NOW()
);

-- Projetos
insert into projeto (
    id,
    reuniao_id,
    sumula,
    relator,
    tem_emendas,
    pac_id,
    par_id,
    codigo_proposicao,
    iniciativa,
    conclusao_comissao,
    conclusao_relator,
    created_at,
    updated_at
)
values (
  'bgxjyguyqcuuy0b6hupw01o4',
  'gdmo94u77a080zyi71nlelwv',
  'Autoriza o Poder Executivo a transferir à Companhia de Habitação Popular de Curitiba - COHAB-CT, a título de alienação gratuita, imóvel que especifica.',
  'Mauro Bobato',
  false,
  1592,
  31868,
  '005.00147.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'a84v5s7ulu21fxjy0qsudfwy',
  'gdmo94u77a080zyi71nlelwv',
  'Autoriza o Poder Executivo a alienar a favor do espólio de Osni Prates Pacheco, representado por Giselle Brunor Pacheco Ebrahim, a área que especifica.',
  'Sidnei Toaldo',
  false,
  1592,
  31901,
  '005.00052.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'tcmrr7a89cyvaztczvfl1xuv',
  'gdmo94u77a080zyi71nlelwv',
  'Acrescenta-se o parágrafo 2º e incisos, e renumera-se o parágrafo único do artigo 9º da Lei Municipal nº 8.670, de 29 de junho de 1995',
  'Toninho da Farmácia',
  false,
  1592,
  31902,
  '005.00151.2022',
  'Dalton Borba',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'mhmgl60nogltubu1t5vuobxh',
  'gdmo94u77a080zyi71nlelwv',
  'Suprime, desafeta e incorpora área de terreno aos bens dominicais e autoriza o Poder Executivo a alienar em favor de Marco Aurélio Ferla, Luciana Ferla, Claudia Ferla e Fernanda Ferla, a área que especifica.',
  'Hernani',
  false,
  1592,
  31924,
  '005.00059.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'b0olar2jw8ldwll23lkx96lt',
  'gdmo94u77a080zyi71nlelwv',
  'Institui o controle da poluição sonora veicular no âmbito do Município de Curitiba e dá outras providências.\r\n ',
  'Mauro Bobato',
  false,
  1592,
  31925,
  '005.00070.2022',
  'Jornalista Márcio Barros, Sargento Tania Guerreiro',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'e53cdtftq9s4a4n8f74jr2sh',
  'gdmo94u77a080zyi71nlelwv',
  'Altera o art. 6º da Lei nº 11.095, de 21 de julho de 2004, que dispõe sobre as normas que regulam a aprovação de projetos, o licenciamento de obras e atividades, a execução, manutenção e conservação de obras no Município, e dá outras providências.',
  'Sidnei Toaldo',
  false,
  1592,
  31933,
  '005.00075.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'eumc83oykn63om1m3l7silmd',
  'd3xhufegbwe1mdldxr9pofng',
  'Altera dispositivos da Lei nº 13.900, de 9 de dezembro de 2011, que cria o Conselho Municipal da Juventude.',
  'Tito Zeglin',
  false,
  1593,
  31896,
  '005.00156.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'rk1vz8ocx3mohzo7zdih2hu6',
  'd3xhufegbwe1mdldxr9pofng',
  'Dispõe sobre a criação do programa "Passe Livre à Internet", para garantir acesso e navegação à internet de caráter gratuito aos estudantes do município de Curitiba.',
  'João da 5 Irmãos',
  false,
  1593,
  31899,
  '005.00054.2021',
  'Renato Freitas',
  '',
  'Por mais informações',
  now(),
  now()
),
(
  'wj69pik5a3dua5vi6bdk3q9f',
  'd3xhufegbwe1mdldxr9pofng',
  'Dispõe sobre a Planta Genérica de Valores - PGV, altera dispositivos das Leis Complementares nºs 7/1993, 40/2001 e 44/2002 e revoga dispositivos das Leis Complementares nºs 53/2004 e 91/2014.',
  'Sergio R. B. Balaguer (Serginho do Posto)',
  false,
  1593,
  31940,
  '002.00008.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
),
(
  'eztzzn54sauixarng0v5fix3',
  'xfdisna9nwtkxe5fjveasn9b',
  'Propõe mudar nome da avenida Manoel Ribas para Manoel Ribeiro.',
  'Sergio R. B. Balaguer (Serginho do Posto)',
  false,
  1594,
  31941,
  '002.00009.2022',
  'Prefeito',
  'Pela tramitação',
  'Pela tramitação',
  now(),
  now()
);

-- Pareceres
insert into parecer (
    id,
    codigo_proposicao,
    tcp_nome,
    vereador,
    id_texto,
    projeto_id,
    created_at,
    updated_at
)
values
(
    'hpprelurepah4h1p05bprdck',
    '005.00147.2022',
    'Por mais informações',
    'Pier Petruzziello',
    38195,
    'bgxjyguyqcuuy0b6hupw01o4',
    now(),
    now()
),
(
    'ga6ei6yq556bs7oqeotnqv0e',
    '005.00147.2022',
    'Pela tramitação',
    'Dalton Borba',
    37933,
    'bgxjyguyqcuuy0b6hupw01o4',
    now(),
    now()
),
(
    'pkg6qwrg08s2oua0tjt34oce',
    '005.00052.2022',
    'Pela anexação',
    'Beto Moraes',
    38192,
    'a84v5s7ulu21fxjy0qsudfwy',
    now(),
    now()
);
