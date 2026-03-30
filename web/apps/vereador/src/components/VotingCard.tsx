import { useRouter } from '@tanstack/react-router';
import {
	Badge,
	Button,
	Card,
	CardContent,
	CardFooter,
	CardHeader,
	CardTitle,
	EVoteType,
	ProjetoDTO,
	useAuth,
} from '@voting/shared';
import { Clock, FileText, Info, Vote } from 'lucide-react';

type VotingCardProps = {
	projectVoting: ProjetoDTO | null;
};

export function VotingCard({ projectVoting }: VotingCardProps) {
	const router = useRouter();

	const { user } = useAuth();

	function handleVoteButton() {
		if (projectVoting) {
			router.navigate({
				to: '/vote',
				search: { votingId: projectVoting.votacao?.id },
			});
		}
	}

	function voteBadge(vote: string) {
		switch (vote) {
			case EVoteType.FAVORAVEL:
				return (
					<Badge className="font-medium bg-green-600 text-white">
						Favorável
					</Badge>
				);
			case EVoteType.FAVORAVEL_RESTRICOES:
				return (
					<Badge className="font-medium bg-orange-600 text-white">
						Favorável com restrições
					</Badge>
				);
			case EVoteType.CONTRARIO:
				return (
					<Badge className="font-medium bg-red-600 text-white">Contrário</Badge>
				);
			case EVoteType.ABSTENCAO:
				return (
					<Badge className="font-medium bg-gray-600 text-white">
						Abstenção
					</Badge>
				);
			case EVoteType.VISTA:
				return (
					<Badge className="font-medium bg-cyan-600 text-white">Vistas</Badge>
				);
			default:
				return (
					<Badge className="font-medium bg-gray-600 text-white">
						Voto nulo
					</Badge>
				);
		}
	}

	function getUserVote() {
		if (!projectVoting || !user) return null;
		const alreadyVoted = projectVoting?.votacao?.votos.find(
			(vote) => vote.usuario.id === user.id,
		);

		return alreadyVoted ? alreadyVoted : null;
	}

	const userVote = getUserVote();

	return (
		<Card className="h-fit w-3xl">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<Info className="h-8 w-8 rounded-full border-primary text-primary " />
					Votações em andamento
				</CardTitle>
			</CardHeader>
			<CardContent>
				{projectVoting ? (
					<div className="flex flex-col justify-center text-sm gap-2">
						<div className="flex items-start gap-3 mb-4">
							<div className="p-1.5 bg-foreground/10 rounded-lg">
								<FileText className="w-4 h-4 text-foreground/50" />
							</div>
							<div>
								<span className="text-sm font-medium text-foreground">
									Projeto
								</span>
								<div className="text-lg font-mono font-semibold text-primary mt-1">
									{projectVoting.codigo_proposicao}
								</div>
							</div>
						</div>
						<div>
							<h4 className="text-sm font-semibold text-foreground mb-3 flex items-center gap-2">
								<div className="w-1 h-4 bg-primary rounded-full"></div>
								Súmula
							</h4>
							<div className="text-sm text-foreground leading-relaxed pl-3">
								{projectVoting.sumula}
							</div>
						</div>
					</div>
				) : (
					<div className="flex flex-col justify-center items-center py-4">
						<div className="relative flex items-center justify-center">
							<div className="flex items-center justify-center w-12 h-12 bg-primary/50 rounded-full">
								<Clock className="w-6 h-6 text-white/80" />
							</div>
							<div className="absolute inset-0 w-12 h-12 bg-primary rounded-full animate-ping opacity-20"></div>
							<div className="absolute inset-0 w-12 h-12 bg-primary-foreground rounded-full animate-ping opacity-10 delay-75"></div>
						</div>
						<div className="text-muted-foreground text-sm pt-8 font-bold">
							Nenhuma votação em andamento
						</div>
						<div className="text-muted-foreground text-xs flex items-center justify-center gap-2 pt-4">
							<span>Aguardando votação ser liberada</span>
							<div className="flex gap-1">
								<div className="w-1 h-1 bg-gray-500 rounded-full animate-pulse"></div>
								<div className="w-1 h-1 bg-gray-500 rounded-full animate-pulse delay-100"></div>
								<div className="w-1 h-1 bg-gray-500 rounded-full animate-pulse delay-200"></div>
							</div>
						</div>
					</div>
				)}
			</CardContent>
			{projectVoting && (
				<CardFooter className="flex flex-col gap-2 justify-between items-center w-full h-fit">
					<div className="flex items-center justify-between">
						<div className="flex items-center gap-4 text-xs">
							<Badge className="font-medium bg-green-600 text-white">
								{projectVoting?.votacao?.totals.totalFavorable || 0} favorável
							</Badge>
							<Badge className="font-medium bg-orange-600 text-white">
								{projectVoting?.votacao?.totals.totalRestriction || 0} favorável
								com restrições
							</Badge>
							<Badge className="font-medium bg-red-600 text-white">
								{projectVoting?.votacao?.totals.totalAgainst || 0} contra
							</Badge>
							<Badge className="font-medium bg-gray-600 text-white">
								{projectVoting?.votacao?.totals.totalAbstention || 0} abstenção
							</Badge>
							<Badge className="font-medium bg-cyan-600 text-white">
								{projectVoting?.votacao?.totals.totalExamination || 0} vistas
							</Badge>
						</div>
					</div>
					{userVote ? (
						<div className="flex w-full text-md items-center text-foreground/50 pt-4">
							<div className="flex items-center justify-center gap-2 text-sm w-full">
								<span>Voto do vereador {userVote.usuario.nome_fantasia}:</span>
								{voteBadge(userVote.voto)}
							</div>
						</div>
					) : (
						<div className="flex w-full justify-end pt-4">
							<Button
								variant={'outline'}
								className="justify-end"
								disabled={!projectVoting || !!userVote}
								onClick={() => handleVoteButton()}
							>
								<Vote className="mr-2 h-4 w-4" />
								Votar agora
							</Button>
						</div>
					)}
				</CardFooter>
			)}
		</Card>
	);
}
