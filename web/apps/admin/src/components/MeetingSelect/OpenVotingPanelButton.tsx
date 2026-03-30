import { Button } from '@voting/shared';
import { ExternalLink } from 'lucide-react';
import { ProjectDTO } from '@/types/meeting';

type OpenVotingPanelButtonProps = {
	project: ProjectDTO;
};

export function OpenVotingPanelButton({ project }: OpenVotingPanelButtonProps) {
	function handleOpenVotingPanel() {
		const win = window.open(`/voting-panel/${project.id}`, 'painel_votacao');

		win?.focus();
	}

	return (
		<div className="col-span-2 pt-4">
			<Button
				className="w-fit"
				variant="outline"
				onClick={handleOpenVotingPanel}
			>
				<ExternalLink className="mr-2 h-4 w-4" />
				Abrir painel de votação
			</Button>
		</div>
	);
}
