import {
	Button,
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from '@voting/shared';
import { RefreshCw } from 'lucide-react';

import { useExecuteSynch } from '@/hooks/useExecuteSynch';
import { useLastSynchs } from '@/hooks/useLastSyncs';

export function LastSyncsCard() {
	const { data: synchronizations = [], isLoading } = useLastSynchs();

	const { mutate: executeSynchronization, isPending } = useExecuteSynch();

	return (
		<Card className="w-full h-fit">
			<CardHeader>
				<CardTitle className="flex items-center gap-2">
					<RefreshCw className="h-8 w-8 rounded-full border-primary text-primary" />
					Últimas sincronias
				</CardTitle>

				<CardDescription>
					Histórico das últimas sincronizações executadas
				</CardDescription>
			</CardHeader>

			<CardContent className="space-y-4">
				{isLoading ? (
					<p className="text-sm text-muted-foreground">Carregando...</p>
				) : synchronizations.length === 0 ? (
					<p className="text-sm text-muted-foreground">
						Nenhuma sincronização encontrada
					</p>
				) : (
					synchronizations.slice(0, 3).map((sync) => (
						<div key={sync.id} className="flex items-center justify-between">
							<div className="flex flex-col">
								<p className="text-sm font-medium">
									{new Date(sync.iniciado_em).toLocaleString('pt-BR')}
								</p>

								<p className="text-xs text-muted-foreground">
									{sync.projetos_sincronizados} projetos ·{' '}
									{sync.reunioes_sincronizadas} reuniões ·{' '}
									{sync.pareceres_sincronizados} pareceres
								</p>
							</div>

							<div
								className={`px-2 py-1 rounded-full text-xs font-medium border ${
									sync.sucesso
										? 'border-green-500/30 bg-green-500/10 text-green-400'
										: 'border-red-500/30 bg-red-500/10 text-red-400'
								}`}
							>
								{sync.sucesso ? 'Sucesso' : 'Erro'}
							</div>
						</div>
					))
				)}
			</CardContent>

			<CardFooter className="flex justify-end">
				<Button
					variant="outline"
					onClick={() => executeSynchronization()}
					disabled={isPending}
				>
					{isPending ? 'Sincronizando...' : 'Executar sincronia'}
				</Button>
			</CardFooter>
		</Card>
	);
}
