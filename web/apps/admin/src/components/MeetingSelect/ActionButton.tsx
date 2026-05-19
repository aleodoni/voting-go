import { Badge } from '@voting/shared';
import { Lock } from 'lucide-react';
import { ProjectDTO } from '@/types/meeting';
import { EVotingStatus } from '@/types/voting-status';
import { CancelVotingButton } from './CancelVotingButton';
import { CloseVotingButton } from './CloseVotingButton';
import { OpenVotingButton } from './OpenVotingButton';

type ActionButtonProps = {
	project: ProjectDTO;
	isVotingProject: boolean;
};

export function ActionButton({ project, isVotingProject }: ActionButtonProps) {
	function returnButtonType() {
		switch (project.votacao?.status) {
			case EVotingStatus.VOTADA:
				return <CancelVotingButton project={project} />;

			case EVotingStatus.ABERTA:
				return <CloseVotingButton project={project} />;

			default:
				if (isVotingProject) {
					return (
						<Badge className="h-8 px-3 flex items-center gap-1 bg-slate-100 text-slate-700 dark:bg-slate-800 dark:text-slate-300">
							<Lock size={14} />
							Outra votação aberta
						</Badge>
					);
				}

				return <OpenVotingButton project={project} />;
		}
	}

	return (
		<div className="flex items-center gap-2 pt-3">{returnButtonType()}</div>
	);
}
