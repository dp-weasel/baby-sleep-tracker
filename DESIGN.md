# Sleep Monitor – Diseño del Proyecto

## Propósito

El propósito de este proyecto es proveer una aplicación simple para registrar y visualizar eventos relacionados con el descanso de un bebé, diseñada para correr en hardware de bajos recursos (Raspberry Pi 3B+) y accesible desde cualquier navegador en la red local.

En su versión inicial, la aplicación permite registrar cuándo el bebé se duerme y cuándo se despierta, junto con sus marcas de tiempo, y mostrar un historial corto de la actividad reciente.

El diseño contempla la posibilidad de incorporar otros tipos de eventos en el futuro (por ejemplo, cambios de pañal o baños) sin modificar el modelo conceptual del dominio.

## Alcance

Dentro del alcance:

* Aplicación web para registrar eventos.
* Backend encargado de persistir eventos y renderizar la interfaz.
* Visualización de un historial reciente con duraciones básicas.

Fuera del alcance (versión inicial):

* Usuarios, autenticación o permisos.
* Analíticas avanzadas o reportes históricos extensos.
* Notificaciones o recomendaciones.

## Modelo de datos (conceptual)

El dominio se basa en el concepto de **evento**.

En la versión inicial existen eventos de inicio y fin del sueño.

Cada evento posee:

* Tipo de evento.
* Marca de tiempo.

Los eventos se consideran secuenciales e inmutables. Las duraciones se calculan a partir de la diferencia temporal entre eventos consecutivos.

El modelo está pensado para permitir la incorporación de nuevos tipos de eventos manteniendo una estructura uniforme.

## Arquitectura del sistema

El sistema adopta una arquitectura simple y local y orientada a **hipermedia con renderizado del lado del servidor**.

Componentes:

* Base de datos SQLite para persistencia.
* Backend en Go que concentra la lógica de dominio y el renderizado de HTML.
* Frontend web liviano basado en HTML generado en servidor y HTMX para interacciones dinámicas.

El backend no expone una API REST clásica. En su lugar, actúa como proveedor de hipermedia:

* Los endpoints HTTP devuelven HTML completo o fragmentos de HTML.
* HTMX se utiliza para disparar requests y actualizar partes específicas de la interfaz sin recargar la página completa.

Este enfoque prioriza simplicidad, bajo acoplamiento conceptual y un flujo claro request → respuesta HTML.

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

El backend está implementado en Go y funciona como el **núcleo del sistema**.

Gestiona tanto el dominio de eventos como la generación de la interfaz web.

### Responsabilidades

* Registrar eventos.
* Exponer el historial reciente.
* Aplicar validaciones mínimas.
* Derivar información simple a partir de la secuencia de eventos.
* Renderizar vistas HTML mediante Go Templates.

### Interfaz

* Endpoints HTTP que devuelven HTML (no JSON como contrato principal).
* Operaciones orientadas a acciones del dominio (por ejemplo, registrar un evento).
* Soporte para respuestas parciales destinadas a ser consumidas por HTMX.

No existe una separación estricta entre “API” y “frontend”. El backend y la interfaz forman un sistema cohesivo.

### Reglas del dominio

* Los eventos forman una secuencia temporal.
* No se persisten estados ni duraciones.
* Las correcciones se realizan agregando nuevos eventos.

## Frontend

El frontend provee una interfaz web mínima y funcional, pensada para uso frecuente en un contexto doméstico.

### Principios

* HTML renderizado del lado del servidor.
* Uso de Go Templates para vistas.
* Interactividad acotada mediante HTMX.
* Ausencia de frameworks de frontend complejos o lógica pesada en el cliente.

### Interacción

* Dos acciones principales: registrar que el bebé se durmió o se despertó.
* Selección del momento del evento: ahora, hace unos minutos o hora personalizada.

### Historial

* Visualización de un historial reciente de eventos.
* Inclusión de la duración desde el evento anterior cuando sea posible.

### Relación con el backend

* Comunicación exclusiva mediante HTTP.
* El frontend se limita a disparar acciones y mostrar HTML generado por el backend.
