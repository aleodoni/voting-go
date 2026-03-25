import { useQuery } from '@tanstack/react-query'
import { getApi } from '@voting/shared'

export type VotingStats = {
  total_projects: number
  total_voted_projects: number
}

async function fetchVotingStats(): Promise<VotingStats> {
  const { data } = await getApi().get<VotingStats>('/votacao/stats')
  return data
}

export function useVotingStats() {
  return useQuery({
    queryKey: ['voting-stats'],
    queryFn: fetchVotingStats,
    staleTime: 0,
  })
}