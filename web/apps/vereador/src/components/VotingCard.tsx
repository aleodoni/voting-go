import { useRouter } from '@tanstack/react-router';
import {
	Badge,
	Button,
	Card,
	CardContent,
	CardFooter,
	CardHeader,
	CardTitle,
	calcularTotaisVotos,
	EVoteType,
	EVotingStatus,
	ProjetoDTO,
	useAuth,
	VotingTotals,
} from '@voting/shared';
import {
	AlertTriangle,
	Check,
	Clock,
	Eye,
	Info,
	MinusCircle,
	Vote,
	X,
} from 'lucide-react';
import { useRegistraVoto } from '@/hooks/useRegistraVoto';

const EMPTY_TOTALS: VotingTotals = {
	totalFavorable: 0,
	totalRestriction: 0,
	totalAgainst: 0,
	totalExamination: 0,
	totalAbstention: 0,
};

type VotingCardProps = {
	projectVoting: ProjetoDTO | null;
};

// ----- Components externos para performance -----
type TotalsBadgeProps = {
	totals: VotingTotals;
};

export function VotingTotalsBadge({ totals }: TotalsBadgeProps) {
	return (
		<div className="flex flex-wrap gap-2">
			<Badge className="bg-green-600 text-white flex items-center gap-1">
				<Check className="w-3 h-3" />
				{totals.totalFavorable} favorável
			</Badge>
			<Badge className="bg-orange-600 text-white flex items-center gap-1">
				<AlertTriangle className="w-3 h-3" />
				{totals.totalRestriction} com restrições
			</Badge>
			<Badge className="bg-red-600 text-white flex items-center gap-1">
				<X className="w-3 h-3" />
				{totals.totalAgainst} contra
			</Badge>
			<Badge className="bg-gray-600 text-white flex items-center gap-1">
				<MinusCircle className="w-3 h-3" />
				{totals.totalAbstention} abstenção
			</Badge>
			<Badge className="bg-cyan-600 text-white flex items-center gap-1">
				<Eye className="w-3 h-3" />
				{totals.totalExamination} vistas
			</Badge>
		</div>
	);
}

type UserVoteBadgeProps = {
	userVote: { voto: string; usuario: { nome_fantasia: string } } | null;
};

export function UserVoteBadge({ userVote }: UserVoteBadgeProps) {
	function voteBadge(vote: string) {
		switch (vote) {
			case EVoteType.FAVORAVEL:
				return (
					<Badge className="bg-green-600 text-white flex items-center gap-1">
						<Check className="w-3 h-3" />
						Favorável
					</Badge>
				);
			case EVoteType.FAVORAVEL_RESTRICOES:
				return (
					<Badge className="bg-orange-600 text-white flex items-center gap-1">
						<AlertTriangle className="w-3 h-3" />
						Favorável com restrições
					</Badge>
				);
			case EVoteType.CONTRARIO:
				return (
					<Badge className="bg-red-600 text-white flex items-center gap-1">
						<X className="w-3 h-3" />
						Contrário
					</Badge>
				);
			case EVoteType.ABSTENCAO:
				return (
					<Badge className="bg-gray-600 text-white flex items-center gap-1">
						<MinusCircle className="w-3 h-3" />
						Abstenção
					</Badge>
				);
			case EVoteType.VISTA:
				return (
					<Badge className="bg-cyan-600 text-white flex items-center gap-1">
						<Eye className="w-3 h-3" />
						Vistas
					</Badge>
				);
			default:
				return <Badge className="bg-gray-600 text-white">Voto nulo</Badge>;
		}
	}

	if (!userVote) return null;

	return (
		<div className="flex items-center justify-center gap-2 text-sm w-full text-foreground/70">
			<span>Seu voto ({userVote.usuario.nome_fantasia}):</span>
			{voteBadge(userVote.voto)}
		</div>
	);
}

// ----- VotingCard principal -----
export function VotingCard({ projectVoting }: VotingCardProps) {
	const router = useRouter();
	const { user } = useAuth();
	const mutation = useRegistraVoto();

	const totals = projectVoting?.votacao?.votos
		? calcularTotaisVotos(projectVoting.votacao.votos)
		: EMPTY_TOTALS;

	const userVote =
		projectVoting?.votacao?.votos.find(
			(vote) => vote.usuario.id === user?.id,
		) ?? null;

	const isVotingOpen = projectVoting?.votacao?.status === EVotingStatus.ABERTA;

	async function _handleVote(voto: EVoteType) {
		if (!projectVoting?.votacao) return;
		try {
			await mutation.mutateAsync({
				votacaoId: projectVoting.votacao.id,
				body: { voto },
			});
		} catch {
			// erro tratado pelo mutation.isError
		}
	}

	return (
		<Card className="h-fit w-full">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<Info className="h-6 w-6 text-primary" />
					{projectVoting
						? `Projeto ${projectVoting.codigo_proposicao}`
						: 'Nenhuma votação em andamento'}
				</CardTitle>
			</CardHeader>

			<CardContent className="flex flex-col gap-4">
				{projectVoting ? (
					<div className="flex flex-col gap-3">
						{/* Súmula */}
						<h4 className="text-sm font-semibold text-foreground mb-2 flex items-center gap-2">
							<div className="w-1 h-4 bg-primary rounded-full"></div>
							Súmula
						</h4>
						<div className="text-sm text-foreground leading-relaxed pl-3">
							{projectVoting.sumula}
						</div>

						{/* Totais de votos */}
						<VotingTotalsBadge totals={totals} />
					</div>
				) : (
					<div className="flex flex-col justify-center items-center py-6 gap-4">
						<div className="relative flex items-center justify-center">
							<div className="flex items-center justify-center w-14 h-14 bg-primary/50 rounded-full">
								<Clock className="w-6 h-6 text-white/80" />
							</div>
							<div className="absolute inset-0 w-14 h-14 bg-primary rounded-full animate-ping opacity-20"></div>
							<div className="absolute inset-0 w-14 h-14 bg-primary-foreground rounded-full animate-ping opacity-10 delay-75"></div>
						</div>
						<div className="text-muted-foreground text-sm font-bold">
							Aguardando votação
						</div>
						<div className="text-muted-foreground text-xs flex items-center gap-2">
							<span>Em breve será liberada</span>
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
				<CardFooter className="flex flex-col gap-3 w-full">
					{/* User vote or vote button */}
					{userVote ? (
						<UserVoteBadge userVote={userVote} />
					) : (
						<div className="flex w-full justify-end">
							<Button
								variant="outline"
								disabled={!isVotingOpen} // desabilita se votação não estiver aberta
								onClick={() =>
									router.navigate({
										to: '/vote/:votingId',
										params: { votingId: projectVoting.votacao?.id },
									})
								}
							>
								<Vote className="mr-2 h-4 w-4" />
								Votar agora
							</Button>
						</div>
					)}

					{/* Error feedback */}
					{mutation.isError && (
						<p className="text-sm text-destructive w-full text-center">
							{mutation.error.message}
						</p>
					)}
				</CardFooter>
			)}
		</Card>
	);
}
