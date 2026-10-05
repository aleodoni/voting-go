// components/CloseVotingButton.tsx

import { Button, getApiErrorMessage } from '@voting/shared';
import { Loader2, Square } from 'lucide-react';
import toast from 'react-hot-toast';
import { useCloseVoting } from '@/hooks/useCloseVoting';
import { ProjectDTO } from '@/types/meeting';

type CloseVotingButtonProps = {
	project: ProjectDTO;
};

export function CloseVotingButton({ project }: CloseVotingButtonProps) {
	const mutation = useCloseVoting();

	async function handleOnClick() {
		try {
			await mutation.mutateAsync(project.id);
		} catch (error) {
			console.error(error);
			toast.error(getApiErrorMessage(error, 'Erro ao finalizar a votação'));
		}
	}

	return (
		<Button onClick={handleOnClick} disabled={mutation.isPending}>
			{mutation.isPending ? (
				<Loader2 className="mr-2 h-4 w-4 animate-spin" />
			) : (
				<Square className="mr-2 h-4 w-4" />
			)}
			Finalizar votação
		</Button>
	);
}
