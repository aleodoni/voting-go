import { Badge, Card, CardContent } from '@voting/shared';
import { User } from 'lucide-react';
import { ProjectDTO } from '@/types/meeting';
import { EVotingStatus } from '@/types/voting-status';
import { ActionButton } from './ActionButton';
import { ProjectVotingPanel } from './ProjectVotingPanel';
import { StatusBadge } from './StatusBadge';

type ProjectCardProps = {
	project: ProjectDTO;
	isVotingProject: boolean;
};

export function ProjectCard({ project, isVotingProject }: ProjectCardProps) {
	return (
		<Card className="w-full transition-all hover:shadow-md border-muted">
			<CardContent className="p-4 sm:p-5 space-y-4">
				{/* Cabeçalho: código + status */}
				<div className="flex items-center justify-between gap-2 flex-wrap">
					<Badge
						variant="secondary"
						className="text-xs sm:text-sm font-semibold"
					>
						{project.codigo_proposicao}
					</Badge>
					<StatusBadge project={project} />
				</div>

				{/* Súmula */}
				<p className="text-sm leading-relaxed text-foreground">
					{project.sumula}
				</p>

				{/* Iniciativa + Relator */}
				<div className="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm">
					<div className="flex items-start gap-2">
						<User
							size={15}
							strokeWidth={1}
							className="mt-0.5 shrink-0 text-muted-foreground"
						/>
						<div>
							<p className="font-medium text-muted-foreground text-xs">
								Iniciativa
							</p>
							<p className="text-sm">{project.iniciativa}</p>
						</div>
					</div>

					<div className="flex items-start gap-2">
						<User
							size={15}
							strokeWidth={1}
							className="mt-0.5 shrink-0 text-muted-foreground"
						/>
						<div>
							<p className="font-medium text-muted-foreground text-xs">
								Relator
							</p>
							<p className="text-sm">{project.relator}</p>
						</div>
					</div>
				</div>

				{/* Painel de votação */}
				{project.votacao?.status === EVotingStatus.VOTADA ||
				project.votacao?.status === EVotingStatus.ABERTA ? (
					<div className="border-t pt-4">
						<ProjectVotingPanel project={project} />
					</div>
				) : null}

				{/* Ação */}
				<div className="border-t pt-4">
					<ActionButton project={project} isVotingProject={isVotingProject} />
				</div>
			</CardContent>
		</Card>
	);
}
