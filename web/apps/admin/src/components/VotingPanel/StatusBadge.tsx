import { Badge } from '@voting/shared';
import { cn } from '@voting/shared/src/lib';
import { AlertCircle, CheckCircle, Clock } from 'lucide-react';
import { ProjectDTO } from '@/types/meeting';
import { EVotingStatus } from '@/types/voting-status';

type StatusBadgeProps = {
	project: ProjectDTO;
};
export function StatusBadge({ project }: StatusBadgeProps) {
	function getStatusInfo() {
		switch (project.votacao?.status) {
			case EVotingStatus.VOTADA:
				return {
					icon: CheckCircle,
					text: 'Projeto votado',
					color:
						'bg-green-100 text-green-800 dark:bg-green-800 dark:text-green-100',
				};
			case EVotingStatus.ABERTA:
				return {
					icon: Clock,
					text: 'Projeto em votação',
					color: 'bg-yellow-100 text-yellow-800',
				};
			default:
				return {
					icon: AlertCircle,
					text: 'Projeto não votado',
					color:
						'bg-gray-100 text-gray-800 dark:bg-gray-800 dark:text-gray-100',
				};
		}
	}

	const statusInfo = getStatusInfo();
	const StatusIcon = statusInfo.icon;

	return (
		<div className="flex items-center gap-2 mt-2">
			<Badge className={cn('h-6 text-sm', statusInfo.color)}>
				<StatusIcon className="mr-1" />
				{statusInfo.text}
			</Badge>
		</div>
	);
}
