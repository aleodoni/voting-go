# Prompt de início de sessão

Cole no início de uma conversa nova. O repositório é público, então não precisa de token; se um dia ficar privado, forneça um token **somente leitura** na conversa e **nunca** o grave neste arquivo nem em nada versionado.

```
Acesse o repositório https://github.com/aleodoni/voting-go
A branch mais atualizada é a develop.

Leia, na raiz, CLAUDE.md — tem as diretivas de como trabalhar neste projeto
(convenções de commit, quando pode/não pode commitar e pushar, formato de
entrega de arquivos, etc). Leia também BACKLOG.md e docs/roadmap.md pra pegar
o contexto do que já foi feito e o que falta. Se for mexer em arquitetura ou
regra de negócio, leia também docs/architecture.md e docs/domain.md.

Depois de ler, me diga um resumo rápido do estado atual e aguarde eu indicar
por onde continuar.
```

## Variante: sessão com objetivo definido

```
<mesmo bloco acima, até "Depois de ler">

Objetivo da sessão: <descreva>
Prompt de apoio: docs/prompts/<feature-backend|feature-frontend|migration|review|bugfix>.md
```