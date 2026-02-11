# Use the official RabbitMQ image that includes the management UI
FROM rabbitmq:3.13-management

# Enable the STOMP plugin so we can communicate over port 61613
RUN rabbitmq-plugins enable rabbitmq_stomp