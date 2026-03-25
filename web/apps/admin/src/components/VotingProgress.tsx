import { VotingStats } from '@/hooks/useVotingStats';
import { Card, CardContent, CircularProgress } from '@voting/shared';
import React from 'react';


interface VotingProgressProps {
	stats?: VotingStats
	isLoading: boolean
}
export function VotingProgress({isLoading, stats}: VotingProgressProps) {

	if (isLoading) return <p>Carregando...</p>

	const percent = calcularPercentual(stats ? stats.total_voted_projects: 0, stats ? stats.total_projects: 0);

	const [progress] = React.useState(percent);

	return (
		// <div className="flex items-center justify-center my-8">
		<Card className="flex w-1/3 h-fit">
			<CardContent>
				<div className="flex items-center justify-center gap-2">
					<CircularProgress
						value={progress}
						size={150}
						strokeWidth={15}
						shape="square"
						showLabel
						labelClassName="text-xl font-bold"
						renderLabel={(percentToNumber) => `${percentToNumber}%`}
						className="stroke-indigo-500/25 bg-amber-300"
						progressClassName="stroke-indigo-600"
					/>
					Projetos já votados
				</div>
			</CardContent>
		</Card>
	);
}

function calcularPercentual(valor: number, total: number): number {
  if (total === 0) return 0;
  return Math.round((valor / total) * 100);
}
