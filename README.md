# DevQuest🚀
Um aplicativo de aprendizado de programação com gamificação

Projeto em desenvolvimento


### Link do projeto e tarefas:

https://github.com/users/Bulletdev/projects/1/views/4

<div>

```bash

codequest/
│
├── backend/
│   ├── cmd/
│   │   └── main.go                  # Arquivo principal do servidor Go
│   │
│   ├── internal/
│   │   ├── models/
│   │   │   └── lesson.go            # Definição do modelo de Lição
│   │   │
│   │   └── handlers/
│   │       └── lesson_handler.go    # Handlers para rotas de lições
│   │
│   ├── pkg/
│   │   └── utils/
│   │       └── cors.go              # Utilitários como configuração de CORS
│   │
│   └── go.mod                       # Dependências do Go
│
├── frontend/
│   ├── public/
│   │   ├── index.html
│   │   └── favicon.ico
│   │
│   ├── src/
│   │   ├── components/
│   │   │   ├── LessonList.js        # Lista de lições
│   │   │   └── LessonDetail.js      # Detalhes de uma lição
│   │   │
│   │   ├── services/
│   │   │   └── api.js               # Configuração de chamadas à API
│   │   │
│   │   ├── App.js                   # Componente principal
│   │   └── index.js                 # Ponto de entrada do React
│   │
│   ├── package.json
│   └── README.md
│
└── README.md                        # Documentação principal do projeto
```

</div>

--------------------------------------------------

Detalhes do Projeto "CodeQuest"

<div>
Nome Sugerido: CodeQuest 🚀
</div>

Diferenciais do Aplicativo:

Gamificação Inteligente


Sistema de pontos baseado em conclusão diária de lições
Perda de pontos/energia se não realizar lições
Badges e conquistas para motivação
Ranking global e entre amigos

Instruções para iniciar o projeto:

No diretório backend:

```bash
go mod tidy
go run cmd/main.go
```

No diretório frontend:

```bash
npm install
npm start
```

O projeto estará rodando:

Backend: http://localhost:8080
Frontend: http://localhost:3000


<details>
Estrutura de Aprendizado


Lições curtas e práticas (5-15 minutos)
Foco em código prático, não apenas teoria
Progressão baseada em desempenho
Múltiplas linguagens (Python, JavaScript, Go, etc.)


Tecnologias Propostas


Frontend: React com Tailwind CSS
Backend: Go (Golang)
Banco de Dados: PostgreSQL
Autenticação: JWT
Deploy: Kubernetes


Fluxo de Aprendizado


Tutorial inicial para definir nível
Lições adaptativas baseadas no progresso
Exercícios variados: completar código, debugar, criar funções
Mini projetos ao final de cada módulo


Recursos de Monetização


Plano básico gratuito
Plano premium com:

Mais linguagens
Projetos completos
Mentoria online
Certificados


Anúncios não intrusivos no plano gratuito

Desafios Técnicos:

Criar um sistema de avaliação de código seguro
Garantir experiência de usuário fluida
Sistema de pontuação justo e motivador
Segurança na execução de códigos enviados
</details>


Próximos Passos:

<div>
  
Definir escopo detalhado

Criar protótipo de baixa fidelidade

Desenvolver MVP (Produto Mínimo Viável)

Teste com usuários inicial

</div>
