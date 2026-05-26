FROM python:3.12-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install -r requirements.txt

COPY sender.py .
COPY events.json .

CMD ["python", "sender.py"]