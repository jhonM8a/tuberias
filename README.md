# Tuberías

Este proyecto se encarga de procesar mensajes en una cola RabbitMQ y trabajar con diversas bases de datos. Para ejecutar correctamente el programa, es necesario configurar algunas variables de entorno y asegurarse de que las dependencias estén correctamente instaladas.

## Configuración de Variables de Entorno para Ejecutar el Programa

Antes de ejecutar el programa, es necesario configurar las siguientes variables de entorno. Estas variables permiten establecer la conexión con RabbitMQ, bases de datos SQL y NoSQL, y definir el nombre de la cola a utilizar.

### Variables Requeridas

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

- **`BROKER_NAME`**: Nombre del broker de mensajes.  
  **Ejemplo**: `rabbitmq`

- **`DB_USER`**: Usuario para autenticar en la base de datos SQL (MySQL, por ejemplo).  
  **Ejemplo**: `root`

- **`DB_PASSWORD`**: Contraseña para autenticar en la base de datos SQL.  
  **Ejemplo**: `mysql`

- **`DB_HOST`**: Dirección IP o hostname del servidor de base de datos SQL.  
  **Ejemplo**: `181.79.9.72`

- **`DB_PORT`**: Puerto para conectarse a la base de datos SQL.  
  **Ejemplo**: `6447`

- **`DB_DATABASENAME`**: Nombre de la base de datos SQL.  
  **Ejemplo**: `Archivos`

- **`DB_ENGINE`**: Motor de base de datos (por ejemplo, `mysql`).  
  **Ejemplo**: `mysql`

- **`NOSQL_DB_USER`**: Usuario para autenticar en la base de datos NoSQL (MongoDB, por ejemplo).  
  **Ejemplo**: `root`

- **`NOSQL_DB_PASSWORD`**: Contraseña para autenticar en la base de datos NoSQL.  
  **Ejemplo**: `rootpassword123`

- **`NOSQL_DB_HOST`**: Dirección IP o hostname del servidor de base de datos NoSQL.  
  **Ejemplo**: `181.79.9.72`

- **`NOSQL_DB_PORT`**: Puerto para conectarse a la base de datos NoSQL.  
  **Ejemplo**: `6451`

- **`NOSQL_DB_DATABASE`**: Nombre de la base de datos NoSQL.  
  **Ejemplo**: `admin`

- **`NOSQL_NAME`**: Nombre de la base de datos NoSQL (por ejemplo, `mongo`).  
  **Ejemplo**: `mongo`

### Pasos para Configurar las Variables de Entorno

1. Abre una terminal o consola de comandos.
2. Ejecuta los siguientes comandos para establecer las variables de entorno:

```bash
export RABBITMQ_USER=admin
export RABBITMQ_PASS=admin123
export RABBITMQ_HOST=181.79.9.72
export RABBITMQ_PORT=6457
export RABBITMQ_QUEUE_NAME_FILES=test_queue
export BROKER_NAME=rabbitmq
export DB_USER=root
export DB_PASSWORD=mysql
export DB_HOST=181.79.9.72
export DB_PORT=6447
export DB_DATABASENAME=Archivos
export DB_ENGINE=mysql
export NOSQL_DB_USER=root
export NOSQL_DB_PASSWORD=rootpassword123
export NOSQL_DB_HOST=181.79.9.72
export NOSQL_DB_PORT=6451
export NOSQL_DB_DATABASE=admin
export NOSQL_NAME=mongo
