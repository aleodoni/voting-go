# Domínio

## Glossário

| Termo | Significado |
| ----- | ----------- |
| Reunião | Sessão legislativa; vem do SPL via sincronização |
| Projeto | Matéria em pauta numa reunião; pode ter pareceres |
| Parecer | Manifestação sobre um projeto (também sincronizado) |
| Votação | Ato de votar um projeto; agregado raiz com status |
| Voto | Escolha de um usuário numa votação |
| Restrição | Ressalva textual associada a um voto |
| Voto contrário | Voto contrário vinculado a um parecer |
| Credencial | Permissões do usuário: `Ativo`, `PodeVotar`, `PodeAdministrar` |
| Nome fantasia | Nome de exibição do usuário |
| Sincronia | Execução de sincronização SPL → banco local, com contadores e status |
| Vereador / Admin | Perfis dos dois apps (votar / operar a sessão) |

## Status da votação

`A` aberta · `F` fechada · `V` votada/apurada · `C` cancelada — *confirmar semântica exata de `F` vs `V` no código antes de documentar como regra*.

Opções de voto (`OpcaoVoto`): `F`, `R`, `C`, `V`, `A` — *documentar o significado de cada letra aqui (favorável/contrário/etc.)*.

## Regras de negócio

- Só administrador ativo abre, fecha e cancela votações, consulta projetos completos, pesquisa usuários e executa sincronização.
- Só usuário ativo com `PodeVotar` registra voto.
- Um usuário vota uma única vez por votação.
- Um projeto tem no máximo uma votação.
- `IsAdmin` e `CanVote` exigem credencial ativa.
- Eventos publicados: `votacao_aberta`, `votacao_fechada`, `votacao_cancelada`, `voto_registrado`.
- Um job interno fecha votações que permaneceram abertas.

## Pontos a confirmar

Itens marcados em itálico acima são inferências a partir do código; valide com quem opera a sessão antes de tratá-los como regra oficial.