# Arquitetura Final — Sistema Híbrido de Trading (v1.0)

**Status:** projeto conceitual fechado. Nenhuma linha de código. Este documento é o contrato de construção.

---

## 0. Princípios inegociáveis

1. **Matemática vive em código nativo.** Todo sinal nasce do motor determinístico. LLM nunca toca o caminho crítico da ordem.
2. **LLM traduz texto em enums, não decide.** Ela classifica contexto não estruturado (notícias, fatos relevantes, agenda); escolher direção, validar contas e enviar ordens é do código.
3. **Fail closed.** Dúvida, silêncio, bug, timeout ou tearsheet de qualquer componente → sistema para ou reduz tamanho. Nunca continua no automático.
4. **Desconfiança estrutural entre componentes.** Posição é a da corretora; a local é cache. Ninguém acredita em ninguém sem reconciliar.
5. **Quem guarda o guardião:** a última trava é um processo independente e um botão humano — fora da pilha inteligente.

---

## 1. crítica do arquiteto — pontas que sobraram (cada uma com a correção)

### 1.1 Inversão de prioridade
Seis rodadas polindo o **filtro** (IA, máquina de estados, shadow mode). Zero rodadas validando o **fundamento**: o motor matemático ter expectativa positiva. Filtro perfeito sobre motor perdedor não salva nada — só perde mais devagar.
**Correção:** ordem de construção invertida (seção 4). A camada de IA é Fase 2, opcional e com ônus da prova.

### 1.2 Piso de engenharia ausente
Tudo até aqui assumiu a plataforma existindo. Faltava a camada mais barata e que evita a quebra com a IA inteira desligada: hard rules no gatekeeper, OCO residente na corretora, reconciliação na reconexão, ordem idempotente, kill switch humano, watchdog em processo separado.
**Correção:** Fase 0 obrigatória. Nada é construído antes dela.

### 1.3 Posição fantasma na reconexão
A forma mais comum de zerar conta real: cai a conexão durante um fill, a corretora executou, o sistema local volta "achando" que está zerado e dispara outra ordem.
**Correção:** regra dura — ao voltar de qualquer queda, o motor fica congelado até a posição local ser reconciliada com a da corretora. Posição local é cache, nunca verdade.

### 1.4 Estado volátil
Um DIRECTIONAL_VETO que existe só em memória do processo morre no restart — e o motor compra o ativo proibido um minuto depois.
**Correção:** estado é persistido (write-through) com TTL por timestamp de relógio, não por processo vivo. No boot, o motor relê os vetos ativos antes de soltar a primeira ordem.

### 1.5 Sensor mudo
Loop de ingestion morto é indistinguível de mercado calmo: nenhuma notícia chega, o semáforo fica NORMAL para sempre e o motor opera no meio de uma bomba sem contexto.
**Correção:** heartbeat. Qualquer loop sem sinal de vida há mais de X minutos → estado degrada para SHOCK_WAIT + alerta. Silêncio de sensor é risco ativo, não normalidade.

### 1.6 Redundância especulativa
Dois agentes LLM dobram custo e latência para resolver um HOLD. A decisão real já está na matriz determinística de conflito (conservadorismo assimétrico), dentro do código.
**Correção:** v1 com UM agente. O segundo existe como flag de configuração e só entra se o ledger mostrar falso-shock sistemático do agente único.

### 1.7 Degradação em hora de dilúvio
Payroll/Copom: fila estoura, rate limit da API de LLM bate, HOLD expira. HOLD **jamais** pode expirar para NORMAL.
**Correção:** fila com prioridade absoluta (fonte oficial > ruído), orçamento diário de LLM com teto, e budget esgotado/timeout/análise inválida → degrada para SHOCK_WAIT. Degradar é sempre para o lado conservador, nunca para o silêncio.

---

## 2. Estrutura definitiva

### 2.1 Componentes

| # | Componente | Onde roda | Responsabilidade | Nunca faz |
|---|---|---|---|---|
| 1 | Motor Matemático | Processo local, loop crítico via WebSocket | Indicadores, gatilho, dimensionamento (entrada, stop, alvo, tamanho) | Chamar LLM no seu caminho |
| 2 | Gatekeeper de Risco | Middleware, mesmo processo do motor | Hard rules: nº máx. de trades, perda diária, R/R mínimo, cooldown, horários, estado; reconciliação | Aprovar sozinho — só sabe vetar |
| 3 | Ingestion de Notícias | Processo separado, contínuo | CVM/BC + calendário + Telegram/RSS; dedup; sanitização; gatilho híbrido (fonte oficial OU anomalia de volume) | Mandar raw text direto pro agente |
| 4 | Agente LLM (1) | Frio, sob demanda | Lê texto sanitizado dentro de tags <untrusted> e responde SOMENTE com schema de enums (event_category, is_unexpected_shock, sentiment_polarity, confidence) | Escolher lado, ver preço, emitir ordem/veto |
| 5 | Tradutor Evento→Estado | Código puro, adjacente ao agente | Matriz determinística + conservatism assimétrico → estado resultante | Interpretar ambiguidade |
| 6 | Máquina de Estados | Memória (leitura < 1 ms) + SQLite (verdade) | NORMAL / HOLD (≤ 10 s) / SHOCK_WAIT (3–5 min) / DIRECTIONAL_VETO (por lado, por ativo, TTL por relógio); publica eventos de transição | Voltar a NORMAL sem TTL explícito |
| 7 | Position Manager | Mesmo processo do motor | OCO residente na corretora; amend para apertar stop (nunca cancel+new); saída protetiva com limite e banda (nunca market nu); espera leilão | Mandar market order no vácuo de liquidez |
| 8 | Position Guardian | Processo separado, independente | Lê SOMENTE saldo/posição direto da corretora; equity stop (perda total do dia = zero e desconecta); expõe o kill switch manual | Compartilhar código, memória ou config com o resto |
| 9 | Shadow Tracker | Worker assíncrono | Segue tick a tick todo sinal vetado (por hard rule ou por semáforo), descontando spread real; grava o ledger de vetos | Enviar qualquer ordem |
| 10 | Ledger | Banco simples (SQLite) | Trades, sinais com motivo, transições de estado, shadow outcomes, heartbeats; alimenta o relatório semanal | Ser dependência operacional de ninguém |

**Observação de terreno:** os primitivos reais da corretora (suporte a OCO/modify, leilão, idempotência) variam. O modelo acima mapeia para o que a API expõe: descoberta de primitivos é tarefa da Fase 0, e componentes se adaptam à API real — não o contrário.

### 2.2 Caminho quente (entrada de ordem)

1. Motor gera o sinal.
2. Gatekeeper aplica hard rules → qualquer violação aborta e o sinal vai ao ledger com motivo.
3. Gatekeeper lê a máquina de estados em memória (< 1 ms) → AMARELO/VERMELHO aborta e loga.
4. VERDE → ordem OCO envolvente dispara para a corretora com chave de idempotência.
5. Position Manager assina as transições de estado enquanto a posição viver.

### 2.3 Caminho frio (evento externo)

1. Ingestion detecta documento/mensagem nova.
2. Fonte oficial e categoria relevante → HOLD em 1 ms. Fonte de ruído → HOLD só com anomalia de volume no WebSocket.
3. Sanitização → fila com prioridade (oficial > ruído) → agente LLM.
4. Agente responde em enums dentro do prazo (8–10 s); estourou → degrada para SHOCK_WAIT.
5. Tradutor aplica a matriz → estado persistido + evento publicado.
6. Position Manager reage. Expirado o TTL, o estado cai sozinho.

### 2.4 Reação de posições abertas por transição

| Transição | Novas entradas | Posição aberta |
|---|---|---|
| → HOLD | Congeladas ≤ 10 s | Stop apertado via amend |
| → SHOCK_WAIT | Bloqueadas 3–5 min | Contra a posição + severidade ≥ limiar → saída protetiva com limite/banda; a favor → trailing apertado |
| → DIRECTIONAL_VETO | Bloqueado só o lado afetado | Incompatível → saída protetiva na próxima janela com liquidez (respeitando leilão) |
| → NORMAL (TTL) | Liberadas | Volta ao OCO original |

### 2.5 Postura de falha — fail fechado

| Falha | Comportamento |
|---|---|
| Ingestion sem heartbeat > X min | → SHOCK_WAIT + alerta |
| LLM: timeout, cota, budget, schema inválido | → SHOCK_WAIT no ativo afetado |
| Conexão com a corretora cai | Motor congela; ao voltar, reconcilia posição antes de qualquer ordem |
| Store de estados inacessível | Motor só opera em modo conservador (tamanho reduzido ou zero) + alerta |
| Crash/restart de qualquer processo | Relê o estado persistido antes de operar; dedup idempotente de ordens |
| Todo o resto falha | Position Guardian sobrevive: equity stop que nenhuma configuração, código ou agente viola |

---

## 3. Ledger shadow e critério de permanência

- Backtest da camada de notícias é impossível sem look-ahead bias (a LLM "conhece" o passado). Aceito como limitação permanente; a medição é prospectiva.
- Todo veto é liquidado tick a tick, com spread real do instante — nunca OHLC de candle fechado.
- Semanal: escore simétrico = prejuízos evitados − lucros perdidos. Barra ajustada pela assimetria do R/R (R/R 1:2 → precisão dos vetos precisa superar ~2/3).
- Duas semanas consecutivas abaixo da barra → a camada de IA é **desligada por configuração**. Não por debate, não por argumento.
- Auditoria extra: % do tempo gasto em HOLD/SHOCK_WAIT sem evento real — cautela demais também é custo, e aparece no ledger.

---

## 4. Ordem de construção com gates

| Fase | Conteúdo | Gate para avançar |
|---|---|---|
| 0 | Piso: gatekeeper, OCO na corretora, reconciliação, watchdog independente, kill switch, descoberta dos primitivos reais da API. Zero IA | Forçar intencionalmente cada limite em sandbox e ver cada trava disparar |
| 1 | Motor em shadow mode (conta demo), sem IA nenhuma | ≥ 200 sinais gravados com expectativa positiva DEPOIS de custos (spread, taxas, slippage). Não passou = voltar ao motor, não adicionar IA |
| 2 | Camada de IA completa, com semáforo real e auditoria shadow | Duas semanas seguidas com escore simétrico acima da barra da seção 3 |
| 3 | Segundo agente, novas fontes, tuning | Somente com evidência do ledger da Fase 2 pedindo isso |

**A ordem é o gate:** nada da Fase 2 existe enquanto a Fase 1 não passar. O filtro só nasce quando existe um motor que mereça ser filtrado.

**Checklist de primitivos da Fase 0 — o que perguntar à corretora:** suporte a OCO/modify (apertar stop sem cancel+new), idempotência de ordem (chave de retransmissão), marcação de leilão/AUCTION no feed, plataforma em que o SDK roda (ex.: API Python do MT5 só em Windows), rate limit da API, reconexão/heartbeat da conexão.

---

## 5. Cortes registrados (por enquanto)

- Segundo agente LLM: flag, só entra com evidência do ledger.
- Auto-tuning de parâmetros por ML: fora de escopo até o motor passar o gate manual.
- Novas fontes (sentimento de redes sociais, fluxo de opções): mesmo critério do ledger, após a Fase 2.

---

## 6. A realidade fora do código

Termos da API da corretora, regime de IR de day trade e responsabilidade são do operador — o sistema assume isso com a fração de capital que se pode perder inteira sem mudança de vida. Com as travas anteriores no lugar, descobrir um erro custa o preço de um sanduíche; sem elas, custa a conta.

---

## 7. Decisão de linguagem — ônus da prova simétrico

A regra que a camada de IA cumpre, a linguagem também cumpre: ninguém entra sem ônus da prova.

1. **Stack é saída da Fase 0, não decisão de design.** Os primitivos reais da corretora arbitram: SDK oficial só-Python, conector FIX ou REST/WebSocket genérico. Escolher linguagem hoje é apostar em terreno que a Fase 0 ainda vai inspecionar.
2. **Partida de um runtime só.** O caminho quente do contrato é trivial (ler estado em memória, comparar números — <1 ms em qualquer linguagem viva); logo, o critério não é velocidade e sim menos peças: um deploy, um log, um runtime no cronograma de um time de um. Python cobre motor, ingestion, LLM, backtest, shadow e guardian. Nota de coerência com §0.1 ("matemática vive em código nativo"): "nativo" ali significa que o sinal nasce de código determinístico — nunca do LLM — e não um requisito de compilador; camada numérica interpretada com C por baixo (NumPy/Pandas) conta como nativa. Nada neste item reabre decisão de stack.
3. **Guardian isolado por deploy, não por linguagem:** processo próprio + ambiente próprio + dependências pinadas compram a independência exigida em §2.1.
4. **Segundo runtime** (Go/Rust/outro) só entra com evidência no ledger: gargalo medido no caminho do dinheiro, ou SDK oficial que não exista em Python. Latência não é evidência neste regime (10 trades/dia, LLM de 8–10 s no frio, OCO residente na quente). A barra, declarada agora e imutável sem nova emenda: toda ordem grava no ledger a latência sinal→envio; o segundo runtime só entra se, numa janela fixa de 4 semanas, os custos (slippage e perdas) das ordens com latência acima de 250 ms somarem 1R — o risco de um stop. Barra móvel é promessa quebrada: o número muda por emenda ao contrato, nunca na hora em que a vontade de escrever outra linguagem bater.
5. **A decisão mortal é operacional, não sintática:** supervisor de processos com restart automático, watchdog de heartbeat, VPS, log e backup são itens da Fase 0. Elixir/OTP resolve isso por design, mas a resposta barata é resolver no runtime escolhido, não migrar.
6. **Os riscos reais do runtime único são operacionais:** dependência que quebra em upgrade (pinar tudo com lockfile) e SDK preso a plataforma (a API Python do MetaTrader5, por exemplo, só roda em Windows) — descobertas da Fase 0, não motivos para pré-escolher stack.

**Máxima registrada:** linguagem rápida sobre motor sem edge só perde mais rápido.
