import { ProjectDTO } from '@/types/meeting';
import { ProjectCard } from './ProjectCard';
import { VoteList } from './VoteList';

type VotingPanelProps = {
	project: ProjectDTO;
	isFetching: boolean;
	isRefetching: boolean;
};

export function VotingPanel({
	project,
	isFetching,
	isRefetching,
}: VotingPanelProps) {
	if (isFetching && !isRefetching) {
		return (
			<div className="w-full h-full flex items-center justify-center">
				<p>Carregando...</p>
			</div>
		);
	}
	return (
		<div className="flex flex-col w-full h-full gap-4">
			<ProjectCard project={project} />
			<VoteList voting={project.votacao?.votos || []} />
		</div>
	);
}
