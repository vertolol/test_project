FROM python:3.7

ENV PYTHONUNBUFFERED 1

RUN mkdir /rabbitmq
WORKDIR /rabbitmq


COPY /rabbitmq /rabbitmq
RUN pip install -r requirements.txt


CMD ["python", "launch_receive.py"]