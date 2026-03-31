import {
	Button,
	Dialog,
	DialogClose,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	EVoteType,
	Label,
	ProjetoDTO,
	Select,
	SelectContent,
	SelectItem,
	SelectTrigger,
	SelectValue,
	Textarea,
} from '@voting/shared';
import * as React from 'react';
import { VoteType } from '@/types/vote-types';

interface VoteModalProps {
	showModal: boolean;
	setShowModal: (value: boolean) => void;
	selectedVote: VoteType;
	project: ProjetoDTO | null;
	restriction: string;
	setRestriction: (value: string) => void;
	contraryMode: 'text' | 'option';
	setContraryMode: (value: 'text' | 'option') => void;
	contraryReason: string;
	setContraryReason: (value: string) => void;
	contraryType: string;
	setContraryType: (value: string) => void;
	handleConfirmVote: () => Promise<void>;
	isLoading: boolean;
}

export const VoteModal: React.FC<VoteModalProps> = ({
	showModal,
	setShowModal,
	selectedVote,
	project,
	restriction,
	setRestriction,
	contraryMode,
	setContraryMode,
	contraryReason,
	setContraryReason,
	contraryType,
	setContraryType,
	handleConfirmVote,
	isLoading,
}) => {
	const restrictionId = React.useId();
	const contraryTextId = React.useId();
	const contrarySelectId = React.useId();

	const hasPareceres = (project?.pareceres?.length ?? 0) > 0;

	if (!hasPareceres && contraryMode === 'option') {
		setContraryMode('text');
	}

	return (
		<Dialog open={showModal} onOpenChange={setShowModal}>
			<DialogContent className="max-w-lg">
				<DialogHeader>
					<DialogTitle>Confirme seu voto</DialogTitle>
					<DialogDescription>
						Você está prestes a votar <strong>{selectedVote.label}</strong> no
						projeto {project?.codigo_proposicao}.
					</DialogDescription>
				</DialogHeader>

				<div className="mt-4 flex flex-col gap-4">
					{selectedVote.id === EVoteType.FAVORAVEL_RESTRICOES && (
						<div className="flex flex-col gap-2">
							<Label htmlFor={restrictionId}>Detalhes da restrição</Label>
							<Textarea
								id={restrictionId}
								value={restriction}
								onChange={(e) => setRestriction(e.target.value)}
								placeholder="Descreva a restrição..."
								rows={3}
							/>
						</div>
					)}

					{selectedVote.id === EVoteType.CONTRARIO && (
						<div className="flex flex-col gap-3">
							<Label
								htmlFor={
									contraryMode === 'text' ? contraryTextId : contrarySelectId
								}
							>
								Motivo do voto contrário
							</Label>

							<div className="flex gap-2">
								<Button
									type="button"
									variant={contraryMode === 'text' ? 'default' : 'outline'}
									size="sm"
									onClick={() => setContraryMode('text')}
								>
									Texto
								</Button>

								<Button
									type="button"
									variant={contraryMode === 'option' ? 'default' : 'outline'}
									size="sm"
									onClick={() => setContraryMode('option')}
									disabled={!hasPareceres}
								>
									Opção
								</Button>
								{!hasPareceres && (
									<p className="text-sm text-muted-foreground">
										Este projeto não possui parecer disponível para seleção.
									</p>
								)}
							</div>

							{contraryMode === 'text' && (
								<Textarea
									id={contraryTextId}
									value={contraryReason}
									onChange={(e) => setContraryReason(e.target.value)}
									placeholder="Descreva o motivo..."
									rows={3}
								/>
							)}

							{contraryMode === 'option' && (
								<Select value={contraryType} onValueChange={setContraryType}>
									<SelectTrigger id={contrarySelectId}>
										<SelectValue placeholder="Selecione um parecer" />
									</SelectTrigger>

									<SelectContent>
										{project?.pareceres?.map((parecer) => (
											<SelectItem key={parecer.id} value={parecer.id_texto}>
												{parecer.tcp_nome} - {parecer.vereador}
											</SelectItem>
										))}
									</SelectContent>
								</Select>
							)}
						</div>
					)}
				</div>

				<DialogFooter className="mt-4 flex justify-end gap-2">
					<DialogClose asChild>
						<Button type="button" variant="outline" disabled={isLoading}>
							Cancelar
						</Button>
					</DialogClose>

					<Button
						type="button"
						onClick={handleConfirmVote}
						disabled={isLoading}
					>
						{isLoading ? 'Enviando...' : 'Confirmar voto'}
					</Button>
				</DialogFooter>
			</DialogContent>
		</Dialog>
	);
};
