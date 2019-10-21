from receiver.receiver import RabbitMqWorker
from time import sleep


if __name__ == "__main__":
    sleep(30)
    RabbitMqWorker().start_receive()