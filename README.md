# Entrega del Proyecto Final de Taller de Go

Este proyecto incluye la configuración de un contenedor Docker y su conexión con DBeaver para gestionar una base de datos MySQL. A continuación se detallan los pasos para descargar, configurar y ejecutar el proyecto.

## Integrantes
- María Fernanda Martínez May
- Valeria Ixchel Ramírez Martínez
- Juan Fernando Ruíz Jiménez  
- Fernanda Guadalupe Reséndiz Montes



## Instrucciones

### Requisitos Previos
Asegúrate de tener instalados los siguientes programas:
- Go (versión 1.16 o superior)
- Git
- Docker
- DBeaver (para gestionar la base de datos MySQL)

### 1. Descargar el Proyecto
Clona el repositorio en tu máquina local ejecutando el siguiente comando:
git clone https://github.com/fergiee7/ProyectoFrames.git

### 2. Ingresar al Directorio del Proyecto
Cambia al directorio del proyecto clonado:
cd ProyectoFinal

### 3. Crear Contenedor en Docker
Crea un contenedor en Docker para MySQL con las siguientes configuraciones:
- Nombre del contenedor: my-sql
- Usuario: root
- Contraseña: test

Ejecuta el siguiente comando para crear el contenedor:
docker run --name mysql -e MYSQL_ROOT_PASSWORD=test -d mysql:latest
¡Recuerda que debes iniciar el contenedor cada vez que lo uses!

### 4. Configuración de DBeaver
Abre DBeaver y crea una nueva conexión a la base de datos MySQL. Configura la conexión con los siguientes parámetros:
- Nombre base de datos: proyecto_1
- Host: localhost
- Port: 8080
- Username: root
- Password: test

Una vez conectado, podrás ver y gestionar las tablas del proyecto. En el archivo main encontrarás el backend y en el archivo index el frontend.

Cualquier cambio en la base de datos modifícalo en la línea 38:
dsn := "root:test@tcp(127.0.0.1:3306)/proyecto_1?charset=utf8mb4&parseTime=True&loc=Local"

### 5. Corriendo el Programa
Para utilizar y/o actualizar este proyecto:
1. Haz un fork del proyecto.
2. Crea una nueva rama:
git checkout -b feature/nueva-funcionalidad
3. Realiza tus cambios y haz un commit:
git commit -am 'Añadir nueva funcionalidad'
4. Sube los cambios:
git push origin feature/nueva-funcionalidad
5. Crea un Pull Request.
