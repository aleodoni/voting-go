// components/CancelVotingButton.tsx

import { Button } from '@voting/shared';
import { Loader2, X } from 'lucide-react';
import { useCancelVoting } from '@/hooks/useCancelVoting';
import { ProjectDTO } from '@/types/meeting';

type CancelVotingButtonProps = {
	project: ProjectDTO;
};

export function CancelVotingButton({ project }: CancelVotingButtonProps) {
	const mutation = useCancelVoting();

	async function handleOnClick() {
		try {
			await mutation.mutateAsync(project.id);
		} catch (error) {
			console.error(error);
		}
	}

	return (
		<Button
			variant="destructive"
			onClick={handleOnClick}
			disabled={mutation.isPending}
		>
			{mutation.isPending ? (
				<Loader2 className="mr-2 h-4 w-4 animate-spin" />
			) : (
				<X className="mr-2 h-4 w-4" />
			)}
			Cancelar votação
		</Button>
	);
}
