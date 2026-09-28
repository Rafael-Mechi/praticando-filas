FROM python:3.12-slim

WORKDIR /app

# Copia o requirements.txt e instala as dependências
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copia o código fonte
COPY main.py .

# Executa o script
CMD ["python", "main.py"]
