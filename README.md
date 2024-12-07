# tuberias
# Configuración de Variables de Entorno para Ejecutar el Programa

Antes de ejecutar el programa, es necesario configurar las siguientes variables de entorno. Estas variables permiten establecer la conexión con RabbitMQ y definir el nombre de la cola a utilizar.

## Variables Requeridas

A continuación, se describen las variables de entorno que deben ser configuradas:

- **`RABBITMQ_USER`**: Usuario para autenticar en RabbitMQ.  
  **Ejemplo**: `admin`

- **`RABBITMQ_PASS`**: Contraseña para autenticar en RabbitMQ.  
  **Ejemplo**: `admin123`

- **`RABBITMQ_HOST`**: Dirección IP o hostname del servidor RabbitMQ.  
  **Ejemplo**: `181.79.9.72`

- **`RABBITMQ_PORT`**: Puerto para conectarse a RabbitMQ.  
  **Ejemplo**: `6457`

- **`RABBITMQ_QUEUE_NAME_FILES`**: Nombre de la cola donde se enviarán o procesarán los mensajes.  
  **Ejemplo**: `fileToProcess`

## Pasos para Configurar las Variables de Entorno

1. Abre una terminal o consola de comandos.
2. Ejecuta los siguientes comandos para establecer las variables de entorno:

   ```bash
   export RABBITMQ_USER=admin
   export RABBITMQ_PASS=admin123
   export RABBITMQ_HOST=181.79.9.72
   export RABBITMQ_PORT=6457
   export RABBITMQ_QUEUE_NAME_FILES=fileToProcess
