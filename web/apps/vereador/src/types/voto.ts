export type RegistraRestricaoRequest = {
	restricao: string;
};

export type RegistraVotoContrarioRequest = {
	idTexto: number;
	parecerId: string;
};

export type RegistraVotoRequest = {
	voto: string;
	restricao?: RegistraRestricaoRequest;
	votoContrario?: RegistraVotoContrarioRequest;
};
