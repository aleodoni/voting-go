import { createFileRoute } from '@tanstack/react-router';
import { ContainerPage, H2 } from '@voting/shared';
import { VotingPanel } from '@/components/VotingPanel';
import { useProject } from '@/hooks/useProject';

export const Route = createFileRoute('/voting-panel/$projectId')({
	component: VotingPanelPage,
});

function VotingPanelPage() {
	const { projectId } = Route.useParams();
	const {
		data: project,
		isLoading,
		isFetching,
		isRefetching,
		isError,
	} = useProject(projectId);

	if (isLoading) return <p>Carregando...</p>;
	if (isError || !project) return <p>Projeto não encontrado</p>;

	return (
		<ContainerPage>
			<H2 className="mb-8">Painel de votação</H2>
			<VotingPanel
				project={project}
				isFetching={isFetching}
				isRefetching={isRefetching}
			/>
		</ContainerPage>
	);
}
