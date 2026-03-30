import { Badge, Card, CardContent } from '@voting/shared';
import { User } from 'lucide-react';
import { ProjectDTO } from '@/types/meeting';
import { EVotingStatus } from '@/types/voting-status';
import { ProjectVotingPanel } from './ProjectVotingPanel';
import { StatusBadge } from './StatusBadge';

type ProjectCardProps = {
	project: ProjectDTO;
};

export function ProjectCard({ project }: ProjectCardProps) {
	return (
		<Card className="w-full py-4">
			<CardContent className="space-y-4">
				<div className="flex items-center">
					<Badge className="px-4 py-1 text-md font-bold" variant="secondary">
						{project.codigo_proposicao}
					</Badge>
				</div>

				<p className="text-md">{project.sumula}</p>

				<div className="flex flex-wrap gap-4 text-md">
					<div className="flex items-center gap-2">
						<User size={16} />
						<p className="font-bold">Iniciativa:</p>
						<p>{project.iniciativa}</p>
					</div>

					<div className="flex items-center gap-2">
						<User size={16} />
						<p className="font-bold">Relator:</p>
						<p>{project.relator}</p>
					</div>
				</div>

				<div>
					<StatusBadge project={project} />
				</div>

				{project.votacao?.status === EVotingStatus.VOTADA ||
				project.votacao?.status === EVotingStatus.ABERTA ? (
					<ProjectVotingPanel project={project} />
				) : null}
			</CardContent>
		</Card>
	);
}
