# API de Busca de Temperatura por CEP

Esta é uma aplicação em Go que busca a localização de um CEP e retorna a temperatura atual da cidade em Celsius, Fahrenheit e Kelvin. A aplicação utiliza as APIs ViaCEP e BrasilAPI para buscar o endereço e a WeatherAPI para obter os dados de clima.

## Funcionalidades

- Recebe um CEP válido de 8 dígitos como parâmetro de consulta.
- Retorna a temperatura em Celsius, Fahrenheit e Kelvin.
- Responde com os códigos de status HTTP apropriados para sucesso e erros.

## Pré-requisitos

- Go 1.23.3 ou superior
- Docker
- Docker Compose
- Conta no Google Cloud Platform
- API key da [WeatherAPI](https://www.weatherapi.com/)

## Configuração

1. **Clonar o repositório**:

   ```bash
   git clone https://github.com/seu-usuario/go-temp-service.git
   cd go-temp-service
   ```

2. Configurar o arquivo .env: Crie um arquivo .env na raiz do projeto e defina a chave da API:

   ```.env
   WEATHER_API_KEY=your_api_key_here
   ```

3. Build e execução com Docker:

   ```bash
   docker-compose up --build
   ```

4. Acessar a API: A API estará disponível em [http://localhost:8080/weather?cep=SEU_CEP](http://localhost:8080/weather?cep=SEU_CEP).

## Rotas

### - GET /weather?cep=89500001

**_Exemplo de resposta:_**

```json
{
  "temp_C": 19.2,
  "temp_F": 66.56,
  "temp_K": 292.35
}
```

---

### Tratamento de Erros

**_CEP inválido:_**

```json
{
  "message": "invalid zipcode"
}
```

#### Status: 422

---

**_CEP não encontrado:_**

```json
{
  "message": "can not find zipcode"
}
```

#### Status: 404

---

**_Erro ao buscar dados de clima:_**

```json
{
  "message": "error fetching weather data"
}
```

#### Status: 500

---

## Deploy no Google Cloud Run

**1. Login no Google Cloud:**

```bash
gcloud auth login
gcloud config set project YOUR_PROJECT_ID
```

---

**2. Build da imagem Docker:**

```bash
docker build -t gcr.io/YOUR_PROJECT_ID/temp-service .
```

**3. Push da imagem para o Google Container Registry:**

```bash
docker push gcr.io/YOUR_PROJECT_ID/temp-service
```

**4. Deploy no Google Cloud Run:**

```bash
gcloud run deploy temp-service \
  --image gcr.io/YOUR_PROJECT_ID/temp-service \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=your_api_key_here
```

**5. Acesse a URL fornecida pelo Google Cloud Run após o deploy.**

[Clique aqui para acessar a API de Busca de Temperatura por CEP](https://go-temp-service-871654810459.southamerica-east1.run.app/)

Ex: https://go-temp-service-871654810459.southamerica-east1.run.app/weather?cep=89500001

## Licença

Este projeto está sob a licença MIT. Sinta-se à vontade para modificar e utilizar conforme suas necessidades.
