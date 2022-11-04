# MELI challange

WIP del challange para MercadoLibre. El challange consiste de 3 objetivos:

1. Inventario.
2. Generación de índice de riesgo.
3. Disponibilización de la información.

El trabajo se va a realizar de forma dinámica y se irán disponilizando de los objetivos a medida que estén.

Por ahora, este WIP contempla la parte (1), es decir, todo lo que comprende el inventario. Para ellos se realizaron las siguientes tareas:

- Dockerización tanto de la BD como el servicio que disponibiliza los endpoints para el llenado y actualización de datos al inventario.
- Desarrollo de API REST en Golang.

Para el desarrollo de la primera parte (Inventario), se requería, antes que todo, poblar la BD con los datos entregados. Por esta razón es que en primerísimo lugar se requiere inicializar la BD creando las tablas necesarias.

La data entregada para insertar en la BD se encuentra en la ruta [cmd/seed/json-data/](cmd/seed/json-data). Para cargar esta data a la BD de manera más sencilla es que decidí transformar a `.csv` y así poder cargarla más fácilmente a la BD.

La transformación de la data de `.json` a `.csv` se logró a través de un golang script, el que vive en la ruta [cmd/seed/main.go](cmd/seed/main.go). Y para correrlo debemos ejecutar lo siguiente en la raíz del proyecto:

```bash
cd cmd/seed && go run main.go
```

Este script genera archivos `.csv` para cada una de las fuentes de datos en [cmd/seed/csv-data](cmd/seed/csv-data).

La población de la BD con esta data ocurre al integrar esta ruta a un volumen en `docker-compose`, que luego va a ser usado para insertar la data en las tablas creadas cuando se levanta el orquestador de contenedores a través de un `init.sql` script que vive en [scripts/init.sql](scripts/init.sql).

Dado que la parte (1) requiere que:

- Se puedan ingresar nuevos empleados.
- Se pueda actualizar el estado de un empleado.
- Se puedan ingresar el resto de datos del anexo, esto es, los datos del tipo Aplicación, Roles, Acceso a la BD.

Tomando esto en cuenta se disponibilizan los siguientes endponts:

| Request | Endpoint |
| :-- | :-- |
| POST | `/employee` |
| PUT | `/employee/{path_parameter}` |
| POST | `/role` |
| POST | `/application` |
| POST | `/dbaccess` |

**WIP: los payloads y explicación se explicarán prontamente.**

## Cómo levantar el servicio

Estando en la raíz del proyecto, corremos primero el script que genera los `.csv`. Luego, editamos el archivo `.env.example` y ponemos la contraseña para el root user (root para hacernos la vida más simple), y el nombre de la BD, que debe ser `pasidb` dado que el script de generación y carga de las tablas está configurado para usar esta BD. Luego cambiamos el nombre del archivo de `.env.example` a `.env`.

A continuación corremos lo siguiente:

```bash
make up
```

Esto va a levantar el proyecto y a disponibilizar tanto la API como la BD. Si queremos mirar que las tablas estén correctamente pobladas en la BD, corremos:

```bash
make bd
```

Nos va a pedir la contraseña de `root` user, que pare efectos de este proyecto será `root`. Una vez dentro, podemos, por ejemplo, mirar la data de la tabla `employee`. Para esto corremos lo siguiente en la consola de `mysql`:

```sql
SELECT * FROM pasidb.employee;
```

Si todo está correcto deberíamos poder ver info.

El proyecto además expones el puerto `:8080` para las consultas tipo REST. Probemos con la creación de un empleado. Para ellos, en otra terminal corremos:

```bash
curl -X POST -H 'Content-Type: application/json' \
-d '{"username": "pasifrola","department_code": 111,"date_in": "2022-10-12"}' \
localhost:8080/employee
```

Este request debería devolver un status code 201. Para ver si el dato se creó, una alternativa es ir a la instancia de la BD que abrimos previamente y volver a ejecutar el `SELECT` sobre la tabla de empleados.


### Disclaimer

En adelante voy  a trabajar sobre dos ramas; `main` que es donde pretendo tener el proyecto final, y `wip`, que es la rama con la data más actualizada dado que es ahí donde iré agregando y probando cosas.

