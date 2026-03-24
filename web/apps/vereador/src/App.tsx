import { useAuth } from '@voting/shared'

export default function App() {
  const { user } = useAuth()

  return (
    <div>
      <h1>Voting Vereador</h1>
      <p>Logado como: {user?.nome_fantasia}</p>
    </div>
  )
}