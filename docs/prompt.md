# Prompt de início de sessão

Cole o bloco abaixo no início de **toda** conversa nova. O repositório é público, então não precisa de token; se um dia ficar privado, forneça um token **somente leitura** na conversa e **nunca** o grave neste arquivo nem em nada versionado.

```
Acesse o repositório https://github.com/aleodoni/voting-go
A branch mais atualizada é a develop.

Leia, nesta ordem:
1. CLAUDE.md (raiz) — diretivas de como trabalhar neste projeto (convenções
   de commit, quando pode/não pode commitar e pushar, formato de entrega de
   arquivos, etc).
2. docs/roadmap.md — em especial a seção "Próximo passo" e o que está em andamento.
3. BACKLOG.md — dívida técnica por prioridade.
4. docs/architecture.md e docs/domain.md — só se o próximo passo tocar
   arquitetura ou regra de negócio; leia sob demanda.

Depois de ler:
- Faça um resumo curto do estado atual (o que está feito e o que está em andamento).
- Diga qual é o próximo passo: o indicado em "Próximo passo" do roadmap. Se a
  seção estiver vazia ou desatualizada em relação ao BACKLOG/estado do repo,
  proponha o de maior prioridade e explique por quê.
- Comece por ele, a menos que eu indique outro caminho. Se depender de algo
  meu (rodar um comando, colar uma saída, tomar uma decisão), peça.

Ao final de cada iteração, ANTES de eu commitar:
1. Mostre o checklist de documentação: para cada arquivo (docs/roadmap.md,
   BACKLOG.md, CLAUDE.md "Estado atual", docs/architecture.md, docs/domain.md,
   ADRs, runbooks, README) diga "atualizar" ou "sem mudança" e o motivo, com
   base nos passos executados na iteração.
2. Entregue os arquivos de docs atualizados junto com o resto da mudança
   (inclusive a seção "Próximo passo" do roadmap).
3. Só então dê a mensagem de commit, em inglês, num único bloco de código.
```

## Variante: sessão com objetivo definido

```
<mesmo bloco acima>

Objetivo da sessão: <descreva>
Prompt de apoio: docs/prompts/<feature-backend|feature-frontend|migration|review|bugfix>.md
```