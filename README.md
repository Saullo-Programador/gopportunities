# GoOpportunities 🚀

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![Framework](https://img.shields.io/badge/Framework-Gin-00ADD8?style=flat-square)](https://gin-gonic.com/)
[![Database](https://img.shields.io/badge/Database-SQLite-003B57?style=flat-square&logo=sqlite)](https://www.sqlite.org/)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)](LICENSE)

GoOpportunities é uma API REST de alta performance desenvolvida em Go para o gerenciamento de oportunidades de emprego. Fornece um backend escalável para plataformas de recrutamento, permitindo operações completas de CRUD em vagas, com foco em código limpo e experiência do desenvolvedor.

## ✨ Principais Funcionalidades

- **Ciclo de Vida CRUD Completo**: Gestão total de vagas (Criação, Leitura, Atualização e Exclusão).
- **Documentação Interativa**: Documentação automática da API utilizando Swagger/OpenAPI para integração simplificada.
- **Persistência Eficiente**: Utilização do GORM para interações otimizadas com o banco de dados e migrações de schema.
- **Logging Estruturado**: Sistema de logs centralizado para melhor observabilidade e depuração.
- **Schemas Tipados**: Modelos de dados estritamente definidos, garantindo a consistência da API.

## 🛠️ Stack Tecnológica

| Tecnologia | Propósito | Justificativa |
| :--- | :--- | :--- |
| **Go** | Linguagem | Escolhida por seu modelo de concorrência superior e eficiência de runtime. |
| **Gin Gonic** | Framework Web | Selecionado por sua abordagem minimalista e roteamento de alta velocidade. |
| **GORM** | ORM | Fornece uma camada de abstração poderosa para operações de banco de dados. |
| **SQLite** | Banco de Dados | Leve e portátil, ideal para desenvolvimento rápido e testes. |
| **Swagger** | Documentação | Padroniza o contrato da API para consumidores de front-end e terceiros. |

## 🏗️ Arquitetura

O projeto segue uma arquitetura modular para garantir a separação de preocupações e a manutenibilidade:

- `config/`: Inicialização do sistema, variáveis de ambiente e conexão com o banco de dados.
- `handler/`: Camada de controle responsável pelo parsing de requisições e orquestração de respostas.
- `router/`: Definições de rotas e integração de middlewares.
- `schemas/`: Entidades de domínio e modelos de banco de dados (GORM).
- `docs/`: Especificações OpenAPI geradas automaticamente.

## 🚀 Como Começar

### Pré-requisitos
- Go 1.20+
- Git

### Instalação e Configuração
1. **Clone o repositório**:
   ```bash
   git clone https://github.com/Saullo-Programador/gopportunities.git
   cd goopportunities
   ```

2. **Instale as dependências**:
   ```bash
   go mod tidy
   ```

3. **Execute a aplicação**:
   ```bash
   go run main.go
   ```

A API estará disponível em `http://localhost:8080`.

## 📖 Referência da API

Acesse a interface interativa do Swagger para explorar e testar os endpoints:
👉 `http://localhost:8080/swagger/index.html`

### Principais Endpoints

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `GET` | `/api/v1/openings` | Lista todas as oportunidades disponíveis |
| `GET` | `/api/v1/opening` | Detalhes de uma vaga específica |
| `POST` | `/api/v1/opening` | Cria uma nova oportunidade de emprego |
| `PUT` | `/api/v1/opening` | Atualiza dados de uma vaga existente |
| `DELETE` | `/api/v1/opening` | Remove uma vaga do sistema |

## 🗺️ Roadmap

- [ ] **Autenticação**: Implementar JWT (JSON Web Tokens) para acesso seguro.
- [ ] **Migração de Banco**: Suporte para PostgreSQL para ambientes de produção.
- [ ] **Testes Unitários**: Implementar suítes de testes para handlers e serviços.
- [ ] **Dockerização**: Criar imagem Docker para facilitar o deploy.
- [ ] **CI/CD**: Configurar GitHub Actions para linting e testes automatizados.

## 🤝 Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para enviar um Pull Request.

1. Faça um Fork do projeto
2. Crie sua branch de feature (`git checkout -b feature/MinhaFuncionalidade`)
3. Faça o commit de suas alterações (`git commit -m 'Adiciona MinhaFuncionalidade'`)
4. Push para a branch (`git push origin feature/MinhaFuncionalidade`)
5. Abra um Pull Request

---
Desenvolvido por [Saullo Programador](https://github.com/Saullo-Programador)
