import { Card, CardContent } from '@voting/shared';
import { VotoDTO } from '@/types/meeting';
import { VoteBadge } from './VoteBadge';

type VoteListProps = {
	voting: VotoDTO[];
};

export function VoteList({ voting }: VoteListProps) {
	return (
		<Card className="w-full flex-1">
			<CardContent>
				<div className="flex flex-col">
					<p className="pb-4 text-2xl font-bold">Votos dos vereadores</p>

					{voting.length > 0 ? (
						voting.map((vote) => (
							<div
								key={vote.id}
								className="flex w-full items-center gap-4 pb-2"
							>
								<p className="w-[60%] font-bold">
									{vote.usuario.nome_fantasia}
								</p>

								<div className="flex w-[40%] justify-end">
									<VoteBadge vote={vote} />
								</div>
							</div>
						))
					) : (
						<p className="text-muted-foreground">Nenhum voto registrado.</p>
					)}
				</div>
			</CardContent>
		</Card>
	);
}
