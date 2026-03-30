import { ProjectDTO } from '@/types/meeting';
import { ProjectCard } from './ProjectCard';

type ProjectsMeetingProps = {
	projects: ProjectDTO[];
	isFetching: boolean;
	isRefetching: boolean;
	isVotingProject: boolean;
};

export function ProjectsMeeting({
	projects,
	isFetching,
	isRefetching,
	isVotingProject,
}: ProjectsMeetingProps) {
	if (isFetching && !isRefetching) {
		return (
			<div className="flex items-center justify-center py-10">
				<p>Carregando...</p>
			</div>
		);
	}

	return (
		<div className="flex flex-col w-full rounded-xl border bg-background shadow-sm p-4">
			<div className="pb-4 border-b">
				<p className="text-xl font-semibold">Projetos para votação</p>
			</div>

			<div className="pt-4 flex flex-col gap-4">
				{projects.map((project) => (
					<ProjectCard
						key={project.id}
						project={project}
						isVotingProject={isVotingProject}
					/>
				))}
			</div>
		</div>
	);
}
