import { EVoteType } from '@voting/shared';
import {
	AlertCircle,
	CheckCircle,
	Eye,
	HelpCircle,
	LucideIcon,
	XCircle,
} from 'lucide-react';

export type VoteType = {
	id: EVoteType;
	label: string;
	description: string;
	icon: LucideIcon;
	color: string;
	badge: string;
};

export const voteTypes: VoteType[] = [
	{
		id: EVoteType.FAVORAVEL,
		label: 'Favorável',
		description: 'Voto a favor da proposta',
		icon: CheckCircle,
		color: 'bg-green-500 hover:bg-green-600',
		badge: 'bg-green-100 text-green-800',
	},
	{
		id: EVoteType.FAVORAVEL_RESTRICOES,
		label: 'Favorável com Restrições',
		description: 'Voto a favor, mas com ressalvas',
		icon: AlertCircle,
		color: 'bg-orange-500 hover:bg-orange-600',
		badge: 'bg-orange-100 text-orange-800',
	},
	{
		id: EVoteType.CONTRARIO,
		label: 'Contrário',
		description: 'Voto contrário à proposta',
		icon: XCircle,
		color: 'bg-red-500 hover:bg-red-600',
		badge: 'bg-red-100 text-red-800',
	},
	{
		id: EVoteType.ABSTENCAO,
		label: 'Abstenção',
		description: 'Não voto a favor nem contra',
		icon: HelpCircle,
		color: 'bg-gray-500 hover:bg-gray-600',
		badge: 'bg-gray-100 text-gray-800',
	},
	{
		id: EVoteType.VISTA,
		label: 'Vistas',
		description: 'Solicito mais tempo para análise',
		icon: Eye,
		color: 'bg-blue-500 hover:bg-blue-600',
		badge: 'bg-blue-100 text-blue-800',
	},
];

export type VoteRequest = {
	voto: string;
	restricao?: {
		restricao: string;
	};
	votoContrario?: {
		idTexto: number;
		parecerId: string;
	};
};
