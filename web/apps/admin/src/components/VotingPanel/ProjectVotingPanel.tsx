import { ProjectDTO } from '@/types/meeting';

type ProjectVotingPanelProps = {
	project: ProjectDTO;
};

type Totals = {
	totalFavorable: number;
	totalRestriction: number;
	totalAgainst: number;
	totalAbstention: number;
	totalExamination: number;
};
export function ProjectVotingPanel({ project }: ProjectVotingPanelProps) {
	// Inicializa os totais
	const totals: Totals = project.votacao?.votos.reduce(
		(acc, voto) => {
			switch (voto.voto) {
				case 'favoravel':
					acc.totalFavorable += 1;
					break;
				case 'restricao':
					acc.totalRestriction += 1;
					break;
				case 'contrario':
					acc.totalAgainst += 1;
					break;
				case 'abstencao':
					acc.totalAbstention += 1;
					break;
				case 'vista':
					acc.totalExamination += 1;
					break;
				default:
					break;
			}
			return acc;
		},
		{
			totalFavorable: 0,
			totalRestriction: 0,
			totalAgainst: 0,
			totalAbstention: 0,
			totalExamination: 0,
		},
	) ?? {
		totalFavorable: 0,
		totalRestriction: 0,
		totalAgainst: 0,
		totalAbstention: 0,
		totalExamination: 0,
	};

	return (
		<div className="grid grid-cols-2 w-full py-4 text-md gap-2">
			<div className="flex items-center gap-2">
				<div className="w-3 h-3 bg-green-500 rounded-full"></div>
				<p className="text-md">Favorável: {totals.totalFavorable}</p>
			</div>
			<div className="flex items-center gap-2">
				<div className="w-3 h-3 bg-orange-500 rounded-full"></div>
				<p className="text-md">
					Favorável com restrições: {totals.totalRestriction}
				</p>
			</div>
			<div className="flex items-center gap-2">
				<div className="w-3 h-3 bg-red-500 rounded-full"></div>
				<p className="text-md">Contrário: {totals.totalAgainst}</p>
			</div>
			<div className="flex items-center gap-2">
				<div className="w-3 h-3 bg-gray-500 rounded-full"></div>
				<p className="text-md">Abstenção: {totals.totalAbstention}</p>
			</div>
			<div className="flex items-center gap-2 col-span-2">
				<div className="w-3 h-3 bg-cyan-500 rounded-full"></div>
				<p className="text-md">Vistas: {totals.totalExamination}</p>
			</div>
		</div>
	);
}
