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
