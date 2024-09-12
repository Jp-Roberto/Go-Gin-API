# Go-Gin-API

O projeto Go-Gin-API tem como finalidade criar uma API RESTful utilizando a linguagem de programação Go e o framework Gin. Ele é configurado para se conectar a um banco de dados PostgreSQL e utiliza o GORM como ORM para gerenciar as operações de banco de dados. A API pode ser usada para gerenciar dados, como informações de alunos, através de operações CRUD (Create, Read, Update, Delete).

## Estrutura do Projeto

- **controllers**: Contém os controladores que lidam com as requisições HTTP.
- **database**: Configuração e inicialização do banco de dados.
- **models**: Definição dos modelos de dados.
- **routes**: Configuração das rotas da API.
- **main.go**: Ponto de entrada da aplicação.

## Dependências

As principais dependências do projeto são:

- [Gin](https://github.com/gin-gonic/gin): Framework web para Go.
- [GORM](https://gorm.io/): ORM para Go.
- [PostgreSQL Driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL): Driver para conexão com PostgreSQL.

## Como Executar

1. Clone o repositório:
   ```sh
   git clone https://github.com/Jp-Roberto/api-go-gin.git

2.Navegue até o diretório do projeto:
cd api-go-gin

3.Execute o Docker Compose para iniciar o banco de dados PostgreSQL:
docker-compose up -d

4.Execute a aplicação:
go run main.go

A API estará disponível em http://localhost:8080.

Rotas Disponíveis
GET /alunos: Retorna a lista de alunos.
GET /alunos/:id: Retorna um aluno específico pelo ID.
POST /alunos: Cria um novo aluno.
PUT /alunos/:id: Atualiza um aluno existente pelo ID.
DELETE /alunos/:id: Deleta um aluno pelo ID.


Contribuição
Sinta-se à vontade para contribuir com o projeto. Para isso, faça um fork do repositório, crie uma branch para suas alterações e envie um pull request.
