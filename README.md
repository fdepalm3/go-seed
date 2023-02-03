# <div align="center"><img src="https://raw.githubusercontent.com/gist/GeekTree0101/05d338bb59109fc71871711c6fa49377/raw/3ff868ffcf2f84d419c392667335fe7e9f1bf155/dancing-gopher.gif" width="40" height="40"> Semilla Go <img src="https://raw.githubusercontent.com/gist/GeekTree0101/05d338bb59109fc71871711c6fa49377/raw/3ff868ffcf2f84d419c392667335fe7e9f1bf155/dancing-gopher.gif" width="40" height="40"></div>
# <p align="center"><img src="https://gophers-latam.github.io/img/community.png" width="500" height="175" /></p>

La idea de este repositorio es proveer una posible implementación de un web server en Go con los minimos features necesarios 
para atacar el MVP de un proyecto típico en Redbee. Pero también buscamos que sirva como base para construir un espacio de aprendizaje
de manera tal que el futuro cercano GO se integre como lenguaje de desarrollo en los proyectos de Redbee.

Nos podes encontrar en:

<img src="https://upload.wikimedia.org/wikipedia/commons/d/d5/Slack_icon_2019.svg" width="15" height="15"> [semilla-go](https://redbee.slack.com/archives/C041MR2A5GB)

<img src="https://upload.wikimedia.org/wikipedia/commons/d/d5/Slack_icon_2019.svg" width="15" height="15"> [redbeego](https://redbee.slack.com/archives/C041MR2A5GB)

<img src="https://img2.freepng.es/20181201/ib/kisspng-portable-network-graphics-trello-scalable-vector-g-5c025368ae6bb9.4395252315436562967144.jpg" width="15" height="15"> [Trello](https://trello.com/b/UdjUuG5G/semilla-go)



## Features
* Implementación de hexagonal
* Configuración de rutas para recibir requests HTTP
* Tests unitarios
* Tests endpoints HTTP
* Integración a servicio REST externo
* Integración a Redis
* Integración a base SQL (Postgres)
* (TODO) Integración a Kafka
* (TODO) Seguridad


## Prerrequisitos
* Tener instalada una versión de Go >=1.19
* Asegurarse que `$(go env GOPATH)/bin`esté en el `PATH`
* Instalar [mockgen](go install github.com/golang/mock/mockgen@v1.6.0)



## Cómo levantar la semilla:
1. Clonar repositorio
2. Ejecutar `make run .`



## Makefile:
Se incluye un makefile con las siguientes utilidades:

* **build**:    compila los packages y genera el binario de la aplicación
* **test**:     ejecuta todos los tests y genera un reporte de coverage
* **run**:      levanta la aplicación 
* **lint**:     corre un linter y genera un reporte en /reports/linter-report.out
* **generate**: corre todas las reglas de ogeneración de código 

Para correr un comando del Makefile basta ejecutar: `make nombreDeComando` (ej: `make test`)



## Testing

### Testdata

En el directorio `testdata` agregamos datos de prueba comunes a varios casos de prueba, como convención
generamos un archivo por cada entidad del modelo, por ejemplo `pokemon.go`tiene datos de prueba relacionados
a la generación de un `Pokemon`. 

### Mocking

Usamos [gomock](https://github.com/golang/mock) para mockear 

### HTTP Test




## Estructura del proyecto:

TODO



## Dependencias

### Routing:

Usamos [Chi](https://github.com/go-chi/chi) para routing.

