// utils/votacao.ts
import { VotoDTO } from '@/types/meeting';

export type VotingTotals = {
	totalFavorable: number;
	totalRestriction: number;
	totalAgainst: number;
	totalExamination: number;
	totalAbstention: number;
};

export function calcularTotaisVotos(votos: VotoDTO[]): VotingTotals {
	return votos.reduce(
		(acc, voto) => {
			switch (voto.voto) {
				case 'F':
					acc.totalFavorable++;
					break;
				case 'R':
					acc.totalRestriction++;
					break;
				case 'C':
					acc.totalAgainst++;
					break;
				case 'V':
					acc.totalExamination++;
					break;
				case 'A':
					acc.totalAbstention++;
					break;
			}
			return acc;
		},
		{
			totalFavorable: 0,
			totalRestriction: 0,
			totalAgainst: 0,
			totalExamination: 0,
			totalAbstention: 0,
		},
	);
}
