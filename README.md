## How to Run the Project

1. **Download the code:** Clone or download the repository to your local machine.

2. **Navigate to the directory:** Open your terminal and go to the directory where the `main.go` file is located.

3. **Execute the code:** Run the project by executing the following command:
    ```bash
    go run main.go
    ```
   This will start the API on port `8080`.

4. **Send a GET request:** With the API running, use a tool like Postman, Insomnia, or your browser to send a GET request to the following URL:
    ```
    http://localhost:8080/?zip=12334
    ```
   Replace `12334` with any Brazilian ZIP code.

5. **Get the response:** You should receive a response similar to the following:
    ```
`{
	"cep": "14850000",
	"state": "SP",
	"city": "Pradópolis",
	"neighborhood": "",
	"street": "",
	"errors": null
}`

## Como Executar o Projeto

1. **Baixe o código:** Clone ou baixe o repositório para sua máquina local.

2. **Navegue até o diretório:** Abra seu terminal e vá para o diretório onde está localizado o arquivo `main.go`.

3. **Execute o código:** Execute o projeto executando o seguinte comando:
    ```bash
    go run main.go
    ```
   Isso iniciará a API na porta `8080`.

4. **Envie uma solicitação GET:** Com a API em execução, use uma ferramenta como Postman, Insomnia ou seu navegador para enviar uma solicitação GET para a seguinte URL:
    ```
    http://localhost:8080/?zip=12334
    ```
   Substitua `12334` por qualquer CEP brasileiro.

5. **Obtenha a resposta:** Você deve receber uma resposta semelhante a seguinte:
    ```
`{
	"cep": "14850000",
	"state": "SP",
	"city": "Pradópolis",
	"neighborhood": "",
	"street": "",
	"errors": null
 }`
