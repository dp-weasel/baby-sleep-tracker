# Sleep Monitor – Diseño del Proyecto

## Propósito

El propósito de este proyecto es proveer una aplicación simple para registrar y visualizar eventos relacionados con el descanso de un bebé.

En su versión inicial, la aplicación permite registrar cuándo el bebé se duerme y cuándo se despierta, junto con sus marcas de tiempo, y mostrar un historial corto de la actividad reciente.

El diseño contempla la posibilidad de incorporar otros tipos de eventos en el futuro (por ejemplo, cambios de pañal o baños) sin modificar el modelo conceptual del dominio.

## Alcance

Dentro del alcance:

* Aplicación web para registrar eventos.
* Backend encargado de persistir y exponer los eventos.
* Visualización de un historial reciente con duraciones básicas.

Fuera del alcance (versión inicial):

* Usuarios, autenticación o permisos.
* Analíticas avanzadas o reportes históricos extensos.
* Notificaciones o recomendaciones.

## Modelo de datos (conceptual)

El dominio se basa en el concepto de **evento**.

En la versión inicial existe un único tipo funcional de evento: inicio y fin del sueño.

Cada evento posee:

* Tipo de evento.
* Marca de tiempo.

Los eventos se consideran secuenciales e inmutables. Las duraciones se calculan a partir de la diferencia temporal entre eventos consecutivos.

El modelo está pensado para permitir la incorporación de nuevos tipos de eventos manteniendo una estructura uniforme.

## Arquitectura del sistema

El sistema adopta una arquitectura simple y local.

Componentes:

* Base de datos SQLite para persistencia.
* Backend en Go para exponer la API y aplicar reglas del dominio.
* Frontend web para la interacción con el usuario.

El frontend se comunica con el backend mediante HTTP. Se evalúa el uso de HTMX para mantener una implementación liviana.

## Modelo de datos (SQL)

La persistencia se implementa sobre SQLite mediante un esquema normalizado.

**event_types**
Tabla maestra que define los tipos de eventos disponibles.

**activity_logs**
Tabla principal que almacena la secuencia temporal de eventos.

* Los eventos se almacenan de forma inmutable.
* Las duraciones y estados no se persisten.
* Se define un índice temporal para optimizar consultas de historial.

El sistema se inicializa con tipos de eventos mínimos relacionados con el sueño.

## Backend

El backend gestiona el dominio de eventos y expone una API orientada a eventos.

### Responsabilidades

* Registrar eventos.
* Exponer eventos recientes.
* Aplicar validaciones mínimas.
* Derivar información simple a partir de la secuencia de eventos.

### Interfaz

* API HTTP con intercambio de datos en JSON.
* Una única operación conceptual para el registro de eventos.
* Consulta de eventos ordenados cronológicamente, con información derivada opcional.

### Reglas del dominio

* Los eventos forman una secuencia temporal.
* No se persisten estados ni duraciones.
* Las correcciones se realizan agregando nuevos eventos.

## Frontend

El frontend provee una interfaz web mínima para registrar eventos y consultar el historial.

### Interacción

* Dos acciones principales: registrar que el bebé se durmió o se despertó.
* Selección del momento del evento: ahora, hace unos minutos o hora personalizada.

### Historial

* Visualización de un historial reciente de eventos.
* Inclusión de la duración desde el evento anterior cuando sea posible.

### Relación con el backend

* Comunicación exclusiva mediante HTTP.
* El frontend no contiene lógica compleja; se limita a registrar eventos y mostrar resultados.
