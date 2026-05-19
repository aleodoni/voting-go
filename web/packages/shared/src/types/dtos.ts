export type UsuarioVotoDTO = {
	id: string;
	nome: string;
	nome_fantasia: string;
};

export type RestricaoDTO = {
	id: string;
	restricao: string;
};

export type ParecerDTO = {
	id: string;
	codigo_proposicao: string;
	tcp_nome: string;
	vereador: string;
	id_texto: string;
	projeto_id: string;
	created_at: string;
	updated_at: string;
};

export type VotoContrarioDTO = {
	id: string;
	id_texto: string;
	parecer_id: string;
	parecer?: ParecerDTO;
};

export type VotoDTO = {
	id: string;
	voto: string;
	usuario_id: string;
	usuario: UsuarioVotoDTO;
	restricao?: RestricaoDTO;
	voto_contrario?: VotoContrarioDTO;
};

export type VotacaoDTO = {
	id: string;
	projeto_id: string;
	status: string;
	created_at: string;
	updated_at: string;
	votos: VotoDTO[];
};

export type ProjetoDTO = {
	id: string;
	sumula: string;
	relator: string;
	tem_emendas: boolean;
	pac_id: string;
	par_id: string;
	codigo_proposicao: string;
	iniciativa: string;
	conclusao_comissao: string;
	conclusao_relator: string;
	reuniao_id: string;
	created_at: string;
	updated_at: string;
	pareceres?: ParecerDTO[];
	votacao?: VotacaoDTO;
};

export const EVotingStatus = {
	ABERTA: 'A',
	FECHADA: 'F',
	VOTADA: 'V',
	CANCELADA: 'C',
} as const;

export type EVotingStatus = (typeof EVotingStatus)[keyof typeof EVotingStatus];

export const EVoteType = {
	FAVORAVEL: 'F',
	FAVORAVEL_RESTRICOES: 'R',
	CONTRARIO: 'C',
	ABSTENCAO: 'A',
	VISTA: 'V',
} as const;

export type EVoteType = (typeof EVoteType)[keyof typeof EVoteType];

export interface User {
	id: string;
	keycloak_id: string;
	username: string;
	email: string;
	nome: string;
	nome_fantasia?: string;
	credencial: Credencial;
}

export interface Credencial {
	ativo: boolean;
	pode_administrar: boolean;
	pode_votar: boolean;
}
