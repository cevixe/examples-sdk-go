# 3 Factor App Example

## Getting Started

Primero debemos crear una aplicación nueva desplegando el stack común de **cevixe**:

```bash
$ git clone github.com/cevixe/aws-stack-go.git
$ cd aws-stack-go
$ ./scripts/build.sh
$ ./scripts/deploy.sh --stack poc3factor --bucket my-bucket
```

Luego descargamos el proyecto ejemplo de implementación del patrón 3 factor app:
```bash
$ git clone github.com/cevixe/examples-sdk-go.git
$ cd examples-sdk-go/3factor
```

Dentro del proyecto ejemplo encontraremos la definición del api graphql de la aplicación.
Este proyecto contiene los tipos y operaciones comunes de **cevixe** y los módulos
de nuestra aplicación, uno por cada dominio de negocio.
```
📦 3factor
┗ 📂 modules
  ┗ 📂 api
    ┣ 📂 schemas
    ┃ ┣ 📂 cevixe (tipos y operaciones comunes de cevixe)
    ┃ ┗ 📂 product (tipos y operaciones del módulo de productos)
    ┗ 📜 template.yaml (definición graphql schema resource y resolvers)
```

Además tenemos la definición de los servicios, uno por cada dominio de negocio o 
transacción multi dominio de negocio, en este último escenario nuestro agregado 
será el modelado de la máquina de estados de la transacción.
```
📦 3factor
┗ 📂 modules
  ┗ 📂 services
    ┗ 📂 product
      ┣ 📂 cmd (definiciones de puntos de entrada al servicio)
      ┃ ┣ 📂 create (punto de entrada para creación de productos)
      ┃ ┃ ┗ 📜 main.go (iniciar runtime de cevixe y exponer event handler)
      ┃ ┣ 📂 update
      ┃ ┃ ┗ 📜 main.go
      ┃ ┗ 📂 delete
      ┃   ┗ 📜 main.go
      ┣ 📂 pkg (componentes internos del servicio)
      ┣ 📜 go.mod (declaración de dependecias al SDK de cevixe)
      ┣ 📜 makefile (archivo de configuración para la compilación)
      ┗ 📜 template.yaml (declaración de funciones)
```

Para desplegar el proyecto ejemplo debemos ejecutar lo siguiente:
```bash
$ ./scripts/build.sh
$ ./scripts/deploy.sh --application poc3factor --stack poc3factor-impl --bucket my-bucket
```



