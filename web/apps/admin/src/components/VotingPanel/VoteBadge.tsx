import { Badge } from '@voting/shared';
import { VotoDTO } from '@/types/meeting';
import { EVoteType } from '@/types/voting-status';

type VoteBadgeProps = {
	vote: VotoDTO;
};

export function VoteBadge({ vote }: VoteBadgeProps) {
	switch (vote.voto) {
		case EVoteType.FAVORAVEL:
			return (
				<Badge className="px-4 bg-green-500 text-white font-bold">
					FAVORÁVEL
				</Badge>
			);
		case EVoteType.FAVORAVEL_RESTRICOES:
			return (
				<div className="flex flex-col w-full items-end justify-center">
					<Badge className="px-4 bg-orange-500 text-white font-bold">
						FAVORÁVEL COM RESTRIÇÕES
					</Badge>
					<span className="text-xs">{vote.restricao?.restricao}</span>
				</div>
			);
		case EVoteType.CONTRARIO:
			return (
				<div className="flex flex-col w-full items-end justify-center">
					<Badge className="px-4 bg-red-500 text-white font-bold">
						CONTRÁRIO
					</Badge>
					<span className="text-xs">
						{vote?.voto_contrario?.id_texto
							? `${vote.voto_contrario.parecer?.vereador} - ${vote.voto_contrario.parecer?.tcp_nome}`
							: `${vote.restricao?.restricao}`}
					</span>
				</div>
			);
		case EVoteType.ABSTENCAO:
			return (
				<Badge className="px-4 bg-gray-500 text-white font-bold">
					ABSTENÇÃO
				</Badge>
			);
		case EVoteType.VISTA:
			return (
				<Badge className="px-4 bg-cyan-500 text-white font-bold">VISTAS</Badge>
			);
		default:
			return null;
	}
}
