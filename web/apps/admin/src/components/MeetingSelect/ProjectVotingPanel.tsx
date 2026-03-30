import { ProjectDTO } from '@/types/meeting';
import { calcularTotaisVotos, VotingTotals } from '@/utils/votacao';
import { OpenVotingPanelButton } from './OpenVotingPanelButton';

const EMPTY_TOTALS: VotingTotals = {
	totalFavorable: 0,
	totalRestriction: 0,
	totalAgainst: 0,
	totalExamination: 0,
	totalAbstention: 0,
};

type VoteItemProps = {
	label: string;
	value: number;
	color: string;
};

function VoteItem({ label, value, color }: VoteItemProps) {
	return (
		<div className="flex items-center justify-between rounded-md border px-3 py-2">
			<div className="flex items-center gap-2">
				<div className={`w-2.5 h-2.5 rounded-full ${color}`} />
				<p className="text-sm">{label}</p>
			</div>
			<p className="font-semibold text-sm">{value}</p>
		</div>
	);
}

type ProjectVotingPanelProps = {
	project: ProjectDTO;
};

export function ProjectVotingPanel({ project }: ProjectVotingPanelProps) {
	const totals = project.votacao?.votos
		? calcularTotaisVotos(project.votacao.votos)
		: EMPTY_TOTALS;

	return (
		<div className="grid grid-cols-1 md:grid-cols-2 gap-2 py-4">
			<VoteItem
				label="Favorável"
				value={totals.totalFavorable}
				color="bg-green-500"
			/>

			<VoteItem
				label="Favorável c/ restrições"
				value={totals.totalRestriction}
				color="bg-orange-500"
			/>

			<VoteItem
				label="Contrário"
				value={totals.totalAgainst}
				color="bg-red-500"
			/>

			<VoteItem
				label="Abstenção"
				value={totals.totalAbstention}
				color="bg-gray-500"
			/>

			<VoteItem
				label="Vistas"
				value={totals.totalExamination}
				color="bg-cyan-500"
			/>

			<div className="pt-2">
				<OpenVotingPanelButton project={project} />
			</div>
		</div>
	);
}
