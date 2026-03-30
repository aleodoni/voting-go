import { Button } from '@voting/shared';
import { Loader2, Play } from 'lucide-react';
import { useOpenVoting } from '@/hooks/useOpenVoting';
import type { ProjectDTO } from '@/types/meeting';

type OpenVotingButtonProps = {
	project: ProjectDTO;
	reuniaoId?: string;
};

export function OpenVotingButton({
	project,
	reuniaoId,
}: OpenVotingButtonProps) {
	const mutation = useOpenVoting(reuniaoId);

	async function handleOnClick() {
		try {
			await mutation.mutateAsync(project.id);
		} catch (error) {
			console.error(error);
		}
	}

	return (
		<Button onClick={handleOnClick} disabled={mutation.isPending}>
			{mutation.isPending ? (
				<Loader2 className="mr-2 h-4 w-4 animate-spin" />
			) : (
				<Play className="mr-2 h-4 w-4" />
			)}
			Abrir votação
		</Button>
	);
}
