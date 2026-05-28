import { useQueryClient } from '@tanstack/react-query';
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
import { useEffect, useRef, useState } from 'react';
import toast from 'react-hot-toast';
import { useExecuteSynch } from '@/hooks/useExecuteSynch';
import { useLastSynchs } from '@/hooks/useLastSyncs';

export function LastSyncsCard() {
	const queryClient = useQueryClient();

	// controla polling
	const [pollingEnabled, setPollingEnabled] = useState(false);

	const { data: synchronizations = [], isLoading } = useLastSynchs(
		pollingEnabled ? 5000 : false,
	);

	const { mutate: executeSynchronization, isPending } = useExecuteSynch();

	const latestSync = synchronizations[0];

	const isRunning = !!latestSync?.iniciado_em && !latestSync?.finalizado_em;

	// guarda estado anterior
	const previousRunningRef = useRef(isRunning);

	// detecta finalização
	useEffect(() => {
		if (!latestSync) return;

		// terminou agora
		if (previousRunningRef.current && !isRunning) {
			// para polling
			setPollingEnabled(false);

			// invalida query
			queryClient.invalidateQueries({
				queryKey: ['today-meetings'],
			});

			if (latestSync.sucesso) {
				toast.success('Sincronia finalizada com sucesso!');
			} else {
				toast.error('Sincronia finalizada com erro');
			}
		}

		previousRunningRef.current = isRunning;
	}, [
		isRunning,
		latestSync, // invalida query
		queryClient.invalidateQueries,
	]);

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
					synchronizations.map((sync) => {
						const isExecuting = !!sync.iniciado_em && !sync.finalizado_em;

						return (
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
										isExecuting
											? 'border-yellow-500/30 bg-yellow-500/10 text-yellow-400'
											: sync.sucesso
												? 'border-green-500/30 bg-green-500/10 text-green-400'
												: 'border-red-500/30 bg-red-500/10 text-red-400'
									}`}
								>
									{isExecuting
										? 'Executando'
										: sync.sucesso
											? 'Sucesso'
											: 'Erro'}
								</div>
							</div>
						);
					})
				)}
			</CardContent>

			<CardFooter className="flex justify-end">
				<Button
					variant="outline"
					onClick={() => {
						// inicia polling
						setPollingEnabled(true);

						// marca como executando
						previousRunningRef.current = true;

						executeSynchronization();
					}}
					disabled={isPending || isRunning}
				>
					{isPending || isRunning ? 'Sincronizando...' : 'Executar sincronia'}
				</Button>
			</CardFooter>
		</Card>
	);
}
