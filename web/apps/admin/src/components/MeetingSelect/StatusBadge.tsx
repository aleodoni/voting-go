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
						'bg-emerald-100 text-emerald-700 dark:bg-emerald-900/40 dark:text-emerald-300',
				};

			case EVotingStatus.ABERTA:
				return {
					icon: Clock,
					text: 'Projeto em votação',
					color:
						'bg-amber-100 text-amber-700 dark:bg-amber-900/40 dark:text-amber-300',
				};

			default:
				return {
					icon: AlertCircle,
					text: 'Projeto não votado',
					color:
						'bg-slate-100 text-slate-700 dark:bg-slate-800 dark:text-slate-300',
				};
		}
	}

	const statusInfo = getStatusInfo();
	const StatusIcon = statusInfo.icon;

	return (
		<Badge
			className={cn(
				'h-7 px-3 text-sm font-medium flex items-center gap-1',
				statusInfo.color,
			)}
		>
			<StatusIcon size={14} />
			{statusInfo.text}
		</Badge>
	);
}
