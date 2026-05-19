import { createFileRoute, useRouter } from '@tanstack/react-router';
import { EVoteType } from '@voting/shared';
import { useState } from 'react';
import toast from 'react-hot-toast';
import { VoteModal } from '@/components/VoteModal';
import { VoteSelectionView } from '@/components/VoteSelectionView';
import { useIsProjectVoting } from '@/hooks/useIsProjectVoting';
import { useRegistraVoto } from '@/hooks/useRegistraVoto';
import { VoteType } from '@/types/vote-types';

export const Route = createFileRoute('/vote/$votingId')({
	component: VoteSelection,
});

function VoteSelection() {
	const { data: project } = useIsProjectVoting();
	const router = useRouter();

	// estados consolidados
	const [voteData, setVoteData] = useState<{
		selectedVote: VoteType | null;
		restriction: string;
		contraryReason: string;
		contraryType: string;
		contraryMode: 'text' | 'option';
		isLoading: boolean;
	}>({
		selectedVote: null,
		restriction: '',
		contraryReason: '',
		contraryType: '',
		contraryMode: 'text',
		isLoading: false,
	});

	const [showModal, setShowModal] = useState(false);
	const voteMutation = useRegistraVoto(project?.reuniao_id);

	// abre modal e reseta campos
	function handleVoteSelection(selectedVote: VoteType) {
		setVoteData({
			selectedVote,
			restriction: '',
			contraryReason: '',
			contraryType: '',
			contraryMode: 'text',
			isLoading: false,
		});
		setShowModal(true);
	}

	async function handleConfirmVote() {
		if (!voteData.selectedVote || !project?.votacao) return;

		setVoteData((prev) => ({ ...prev, isLoading: true }));

		try {
			const {
				selectedVote,
				restriction,
				contraryType,
				contraryMode,
				contraryReason,
			} = voteData;

			const voteBody: any = {
				voto: selectedVote.id,
			};

			if (selectedVote.id === EVoteType.FAVORAVEL_RESTRICOES) {
				if (!restriction.trim()) {
					toast.error('Informe a restrição');
					return;
				}

				voteBody.restricao = {
					restricao: restriction.trim(),
				};
			}

			if (selectedVote.id === EVoteType.CONTRARIO) {
				if (contraryMode === 'option') {
					if (!contraryType) {
						toast.error('Selecione o texto do voto contrário');
						return;
					}

					const parecer = project?.pareceres?.find(
						(p) => p.id_texto === contraryType,
					);

					if (!parecer) {
						toast.error('Parecer não encontrado');
						return;
					}

					voteBody.votoContrario = {
						idTexto: Number(contraryType),
						parecerId: parecer.id,
					};
				}

				if (contraryMode === 'text') {
					if (!contraryReason.trim()) {
						toast.error('Informe o motivo do voto contrário');
						return;
					}

					voteBody.restricao = {
						restricao: contraryReason.trim(),
					};
				}
			}

			await voteMutation.mutateAsync({
				votacaoId: project.votacao.id,
				body: voteBody,
			});

			toast.success('Voto realizado com sucesso! 🎉');
			router.navigate({ to: '/' });
		} catch {
			toast.error('Erro ao registrar voto');
		} finally {
			setVoteData((prev) => ({ ...prev, isLoading: false }));
			setShowModal(false);
		}
	}

	return (
		<>
			<VoteSelectionView
				projectVoting={project || null}
				handleVoteSelection={handleVoteSelection}
			/>

			{voteData.selectedVote && project && (
				<VoteModal
					showModal={showModal}
					setShowModal={setShowModal}
					selectedVote={voteData.selectedVote}
					project={project}
					restriction={voteData.restriction}
					setRestriction={(r: string) =>
						setVoteData((prev) => ({ ...prev, restriction: r }))
					}
					contraryMode={voteData.contraryMode}
					setContraryMode={(m: 'text' | 'option') =>
						setVoteData((prev) => ({ ...prev, contraryMode: m }))
					}
					contraryReason={voteData.contraryReason}
					setContraryReason={(r: string) =>
						setVoteData((prev) => ({ ...prev, contraryReason: r }))
					}
					contraryType={voteData.contraryType}
					setContraryType={(t: string) =>
						setVoteData((prev) => ({ ...prev, contraryType: t }))
					}
					handleConfirmVote={handleConfirmVote}
					isLoading={voteData.isLoading}
				/>
			)}
		</>
	);
}
