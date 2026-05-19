import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
	ProjetoDTO,
} from '@voting/shared';
import { FileText } from 'lucide-react';
import { VoteType, voteTypes } from '@/types/vote-types';

type VoteSelectionViewProps = {
	projectVoting: ProjetoDTO | null;
	handleVoteSelection: (voteType: VoteType) => void;
};
export function VoteSelectionView({
	projectVoting,
	handleVoteSelection,
}: VoteSelectionViewProps) {
	if (!projectVoting) {
		return null;
	}

	return (
		<div className="flex flex-col gap-8 w-full items-center h-full pt-8">
			<Card className="w-full">
				<CardContent>
					<div className="flex items-start gap-3">
						<FileText className="w-5 h-5 text-primary mt-1" />
						<div className="flex-1">
							<p className="text-md text-muted-foreground">
								Projeto {projectVoting.codigo_proposicao}
							</p>
							<p className="text-sm text-muted-foreground mb-1">
								{projectVoting.sumula}
							</p>
							<p className="text-sm text-muted-foreground">
								Iniciativa: {projectVoting.iniciativa}
							</p>
							<p className="text-sm text-muted-foreground">
								Relator: {projectVoting.relator}
							</p>
						</div>
					</div>
				</CardContent>
			</Card>
			<Card className="w-full">
				<CardHeader>
					<CardTitle>Selecione seu voto</CardTitle>
					<CardDescription>
						Escolha uma das opções de voto disponíveis para este projeto.
					</CardDescription>
				</CardHeader>
				<CardContent>
					<div className="grid gap-4 sm:grid-cols-1md:grid-cols-2 lg:grid-cols-3">
						{voteTypes.map((vote) => {
							const Icon = vote.icon;
							return (
								<Card
									key={vote.id}
									className="cursor-pointer transition-all hover:shadow-md border-2 hover:border-primary/50"
									onClick={() => handleVoteSelection(vote)}
								>
									<CardContent className="p-2 text-center">
										<div
											className={`w-12 h-12 rounded-full ${vote.color} flex items-center justify-center mx-auto mb-3`}
										>
											<Icon className="w-6 h-6 text-white" />
										</div>
										<h3 className="font-semibold text-foreground mb-2">
											{vote.label}
										</h3>
										<p className="text-sm text-muted-foreground">
											{vote.description}
										</p>
									</CardContent>
								</Card>
							);
						})}
					</div>
				</CardContent>
			</Card>
		</div>
	);
}
