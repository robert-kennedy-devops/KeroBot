# KeroBot

![Go](https://img.shields.io/badge/Go-1.25-blue)
![Docker](https://img.shields.io/badge/Docker-Compose-blue)
![License](https://img.shields.io/badge/License-MIT-green)

KeroBot é um **cliente de usuário do Telegram (MTProto)** que automatiza fluxos de jogo em estilo RPG. O foco do projeto é **arquitetura robusta, segurança operacional e automação confiável**, com logs estruturados, persistência em PostgreSQL e engine orientada a eventos.

## Destaques para recrutamento
- **Arquitetura modular** com separação clara entre parsing, engine e execução.
- **Automação resiliente** com fila priorizada, retry e rate limit.
- **Observabilidade** via logs estruturados e métricas.
- **Onboarding seguro**: login por QR com reenvio automático quando o token expira.
- **Customização por usuário**: bot de configuração + regras aprendidas.

## Caso de uso (resumo)
1. O cliente MTProto conecta ao Telegram e escuta o bot alvo.
2. Mensagens são parseadas e convertidas em estados do jogo.
3. A engine decide ações com base em regras e estado atual.
4. Workers atuam como fallback para estabilidade e só enfileiram ações quando o botão está visível.
5. Ações vão para uma fila com delay, retry, anti-flood e prioridade.

## Principais recursos
- Cliente MTProto (sem bot token)
- Detecção de botões inline com matching inteligente
- Engine baseada em eventos + state manager
- Fila de ações com anti-flood, retry e prioridade
- Workers concorrentes context-aware
- Persistência em PostgreSQL + migrations
- Bot de configuração (onboarding e ajustes por usuário)
- Login via QR por sessão (reenvio automático se expirar)
- Modo de captura + regras aprendidas
- Endpoint de métricas
- Docker + Docker Compose

## Arquitetura (mapa rápido)
- `internal/telegram`: cliente MTProto, listener, botões, auth
- `internal/parser`: parser de mensagens e detecção de estado
- `internal/engine`: state manager, action executor, automation engine
- `internal/automation`: workers e loops
- `internal/database`: PostgreSQL e migrations
- `pkg/logger`: logger JSON com `slog`
- `pkg/retry`: helper de retry

## Estrutura do projeto
```
kerobot/
  cmd/kerobot/main.go
  internal/
    telegram/
    parser/
    engine/
    automation/
    database/
    models/
    config/
  pkg/
    logger/
    retry/
  migrations/
  docker/
  docker-compose.yml
  .env.example
  README.md
```

## Setup rápido
1. Crie o `.env`:
```
cp .env.example .env
```
2. Suba o stack:
```
docker compose up --build
```
3. Configure API_ID/API_HASH via bot de configuração.

## .env (exemplo)
```
TG_PASSWORD=
TG_SESSION=./data/telegram.session
TARGET_BOT=bot_alvo

DB_HOST=postgres
DB_PORT=5432
DB_USER=kerobot
DB_PASS=kerobot
DB_NAME=kerobot
DB_SSLMODE=disable

HUNT_INTERVAL=15s
COMBAT_INTERVAL=5s
HEAL_INTERVAL=10s
POTION_INTERVAL=30s
DUNGEON_INTERVAL=1m
HEAL_PERCENT=40
MIN_POTIONS=5

CLICK_DELAY=900ms
RATE_PER_SECOND=2
RETRY_ATTEMPTS=3
RETRY_DELAY=1500ms
METRICS_ADDR=:9090
BOT_ENABLED=true
BOT_TOKEN=seu_bot_token
ADMIN_CHAT_ID=0
```

## Bot de configuração
Comandos principais:
- `/config` mostra o menu
- `/set_api <id>` e `/set_hash <hash>` definem as credenciais do app
- `/qr` gera o QR para login de usuário (novo QR é enviado automaticamente se expirar)
- `/capture_on` e `/capture_off` ativam o modo captura
- `/last` mostra a última captura
- `/learn_last_click <label>` cria regra com base no último botão
- `/learn_last_text <texto>` cria regra com base no último texto

## Métricas
Endpoint JSON em `METRICS_ADDR`:
- `/metrics` retorna contadores de ações e mensagens.

## Decisões técnicas (trade-offs)
- **MTProto (cliente de usuário) vs Bot API**: permite interagir com bots como se fosse um usuário real, porém exige cuidado com autenticação e políticas do Telegram.
- **Engine por eventos + workers de fallback**: melhora a resiliência, mas aumenta concorrência e exige estado consistente.
- **Fila com prioridade e rate limit**: evita flood e banimento, ao custo de latência controlada.
- **Parser heurístico**: simples e rápido, porém exige manutenção quando o texto do bot muda.
- **Persistência em Postgres**: garante histórico e configurações por usuário, com custo operacional maior do que storage local.

## Roadmap (próximos passos)
- Dashboard simples para visualizar métricas e estado das contas.
- Testes de integração para fluxo completo (parser + engine + executor).
- Feature flags por usuário para ativar/desativar automações com segurança.

## Testes
Execute todos os testes:
```
go test ./...
```

## Aviso
Use com responsabilidade e respeite as regras do Telegram e do jogo.
