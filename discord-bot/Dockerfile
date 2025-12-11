FROM python:3.12.12-alpine3.23


WORKDIR /app

COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt

COPY . .

CMD ["python", "src/main.py"]
