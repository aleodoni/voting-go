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
				<p className="text-sm text-muted-foreground">Carregando...</p>
			</div>
		);
	}

	return (
		<div className="flex flex-col w-full rounded-xl border bg-background shadow-sm p-4 sm:p-6">
			<div className="pb-3 sm:pb-4 border-b">
				<p className="text-base sm:text-xl font-semibold">
					Projetos para votação
				</p>
				{projects.length > 0 && (
					<p className="text-xs text-muted-foreground mt-0.5">
						{projects.length} projeto{projects.length > 1 ? 's' : ''}
					</p>
				)}
			</div>

			{projects.length === 0 ? (
				<div className="flex items-center justify-center py-10">
					<p className="text-sm text-muted-foreground">
						Nenhum projeto encontrado para esta reunião.
					</p>
				</div>
			) : (
				<div className="pt-4 flex flex-col gap-3 sm:gap-4">
					{projects.map((project) => (
						<ProjectCard
							key={project.id}
							project={project}
							isVotingProject={isVotingProject}
						/>
					))}
				</div>
			)}
		</div>
	);
}
