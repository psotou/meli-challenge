# MELI challange

## Cómo correr la aplicación

La aplicación realizada consiste en una API que puede cargar información tando de empleados, roles, accesos a BD y aplicaciones. También se puede actualizar la data de los empleados.

Para levantar la aplicación primero que todo se debe tener Docker instalado. Luego, debemos editar el archivo `.env.example` y renombrarlo por `.env`. En este último ingresaremos el nombre de BD y la contraseña de usuario root. Idealmente, estas contraseñas deberían compartirse a través de una ambiente seguro tipo Vault, pero los fines de este proyecto vamos a escribir `root` para contraseña y `pasidb` como BD. 

Con este cambio y asegurándonos de estar parados en la raíz del proyecto corremos:

```bash
make up
```
Una vez levantada la aplicación podemos revisar que las tablas y vistas necesarias están disponibles en la BD, para ellos, primero abrimos otro instancia de la terminal y escribimos:

```bash
make db
```

E ingresamos la contraseña `root`. Una vez dentro de la BD, corremos lo siguiente:

```mysql
use pasidb;
show tables;
```

Esto debería devolver el siguiente listado:

```bash 
+----------------------+
| Tables_in_pasidb     |
+----------------------+
| application          |
| db_access            |
| department_risk_view |
| employee             |
| employee_risk_view   |
| risk_view            |
| role                 |
+----------------------+
```

Donde `application`, `db_access`, `employee` y `role` son las tablas que guardan información relativa a esos datos; `risk_view` es una vista que consolida las relaciones entre las tablas y el puntaje de riesgo asociado tanto para empleados como para departamentos; `employee_risk_view` y `department_risk_view` son tablas consolidades finales con la información de riesgo para empleados y departmentos respectivamente.

## Carga de data

Para la carga de cualquier tipo de información se debe ingresar una lista de objetos en formato `.json` en el endpoint que corresponda a la información que se desee cargar. Por ejemplo, para cargar data de empleados, se debe usar el endpoint `/employees`. El formato de la data es el siguiente:

```json
[
    {
        "id": 1,
        "name": "nombre ejemplo",
        "department": "departamento ejemplo"
    },
    {
        "id": 2,
        "name": "nombre ejemplo 2",
        "department": "departamento ejemplo 3"
    }
]
```

Para ver ejemplos con la data para cargar entregada en el desafío ir a [scripts/sh-script/json-data/](scripts/sh-script/json-data/).

### Endpoints de carga y actualización de datos

Los endpoint para cargar la data son los siguientes:

| HTTP request method | Endpoint | Comentario |
| :-- | :-- | :-- |
| `POST` | `/employees` | para carga de datos de empleados |
| `PUT` | `/employees` | para actualización de datos de empleados (mismo body que el usado para la carga) |
| `POST` | `/roles` | para carga de roles |
| `POST` | `/applications` | para carga de aplicaciones |
| `POST` | `/db_accesses` | para carga de accesos a la BD |

Para cargar de manera "masiva" la data relacionada a estos objetos, debemos correr:

```bash
make seed-data
```

Esto ejecutará POST requests a los endpoints de carga y de esta manera poblaremos la BD. Para asegurarnos de que la data esté cargada podemos ingresar a la BD y hacer un select sobre alguna de las tablas mencionadas arriba.

Para probar las actualizaciones a empleados podemos correr el script que actualiza dos empleados (ver la data a actualziar en [esta ruta](scripts/sh-script/json-data/update-employee.json). El script es el siguiente: 

```bash
cd scripts/sh-script/ && ./update-data.sh
```

### Endpoints de consulta de riesgo

Los endpoint para la consulta del riesgo asociado a un empleado o un departamento son los siguientes:

| HTTP request method | Endpoint | Comentario |
| :-- | :-- | :-- |
| `GET` | `/employeerisk/:username` | donde `:username` es el nombre de usario del empleado a consultar |
| `GET` | `/departmentrisk/:code` | donde `:code` es el código del departamento a consultar |

Se ha generado un script para facilitar la consulta de esta data. Para conocer el riesgo asociado a un empleado, corremos:

```bash
./scripts/sh-script/get-employee-risk.sh {username}
```
 Y para conocer el riesgo asociado a un departamento, corremos:

 ```bash
 ./scripts/sh-script/get-department-risk.sh {department code}
 ```

## Generación el índice de riesgo

Primero debemos cruzar las tablas según sus relaciones para ver cómo se comporta la data. Según la data compartida, vemos que lo cruces son los siguientes:

- `employee.username = role.username`
- `employee.username = db_access.username`
- `role.role_id = application.role_id`

Con estos cruces generamos una tabla virtual (ver **`cte`** [scripts/sql-script/init.sql](scripts/sql-script/init.sql)). Como vamos a cruzar todo con la tabla employee (left join sobre employee), vemos que este cruce genera 3 valores para las columnas `is_pii` (tal vez el dato más relevante) y `is_critical`, y esto valores son `NULL`, `0`, o `1`.

Con esta información en mente agregamos los siguientes valores de riesgo asociados a dichas columnas, lo valores son:

| Columna | Valor de cruce | Valor riesgo asignado | Columna generada |
| :-- | :-- | :-- | :-- |
| `is_pii` | NULL | 0 | `table_risk` |
| `is_pii` | 0 | 1 | `table_risk` |
| `is_pii` | 1 | 2 | `table_risk` |
| `is_critical` | NULL | 0 | `app_risk` |
| `is_critical` | 0 | 0 | `app_risk` |
| `is_critical` | 1 | 1 | `app_risk` |

Como la columan `is_pii` es relevante, asignamos un punto a extra a cada valor, dado que tener accesos a una tabla con información personal es más riesgoso que la criticidad de una aplicación (o al menos esto lo tomé como supuesto).

Luego, la segunda parte más relevantea a la hora de asignar el riesgo es el acceso de usarios a esquemas o tablas de la BD. En la información que cargamos, vemos que hay varios roles, los cuales homologamos y asignamos un puntaje de riesgo según el nivel de privilegios de cada rol. Los puntajes son:

| Rol | Rol homologado | puntaje riesgo rol | Columna generada |
| :-- | :-- | :-- | :-- |
| `ADMINS_SU` | `ADMIN` | 4 | `role_risk` |
| `WRITER_CS` | `WRITER` | 3 | `role_risk` |
| `READER` | `READER` | 2 | `role_risk` |
| `READER_SU` | `READER` | 2 | `role_risk` |
| `USER_READ` | `READER` | 2 | `role_risk` |
| `OPERATOR` | `OTHER` | 1 | `role_risk` |
| `CONSULTANT` | `OTHER` | 1 | `role_risk` |

Y 0 a aquellos con valores nulos.

De esta manera, el puntaje de riesgo asociado a un empleado será el máximo de la suma de las tablas 

Con esto, tenemos que los puntajes de riesgo finales se calculan de la siguiente manera:

- **Riesgo empleado**: `max(sum(table_risk, app_risk, role_risk))`
- **Riesgo departamento**: `max(sum(table_risk, app_risk))`

Además, los empleados con estado `Inactive` tienen un riesgo de 0.

Esto dado que el rol es intrínseco a un empleado, pero no a un departamento. 

El calculo genera distintos rangos numéricos asociados al riesgo de empleados y departamentos. Donde el rango de los empleados es `[0, 7]` y el de los departamentos es `[0, 3]`. Con esto asignamos el riesgo asociado un rango de la siguiente manera:

**Empleados**

| Rango riesgo | Riesgo |
| :-- | :-- | 
| 0 | `no risk` |
| [1, 2] | `low` |
| [3, 4] | `mid` |
| [5, 6] | `high` |
| 7 | `very high` |

**Departamentos**

| Rango riesgo | Riesgo |
| :-- | :-- | 
| 0 | `no risk` |
| 1 | `low` |
| 2 | `mid` |
| 3 | `high` |

## Tests

Para correr los test unitarios asociados a los endpoints generados, basta con correr:

```bash
make tests
```

