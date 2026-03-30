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
			<CardContent className="p-5 space-y-5">
				<div className="flex items-center justify-between">
					<Badge variant="secondary" className="text-sm font-semibold">
						{project.codigo_proposicao}
					</Badge>

					<StatusBadge project={project} />
				</div>

				<div>
					<p className="text-sm leading-relaxed text-foreground">
						{project.sumula}
					</p>
				</div>

				<div className="grid grid-cols-2 gap-6 text-sm">
					<div className="flex items-start gap-2">
						<User size={16} strokeWidth={1} className="mt-1" />
						<div>
							<p className="font-medium text-muted-foreground">Iniciativa</p>
							<p>{project.iniciativa}</p>
						</div>
					</div>

					<div className="flex items-start gap-2">
						<User size={16} strokeWidth={1} className="mt-1" />
						<div>
							<p className="font-medium text-muted-foreground">Relator</p>
							<p>{project.relator}</p>
						</div>
					</div>
				</div>

				{project.votacao?.status === EVotingStatus.VOTADA ||
				project.votacao?.status === EVotingStatus.ABERTA ? (
					<div className="border-t pt-4">
						<ProjectVotingPanel project={project} />
					</div>
				) : null}

				<div className="border-t pt-4">
					<ActionButton project={project} isVotingProject={isVotingProject} />
				</div>
			</CardContent>
		</Card>
	);
}
