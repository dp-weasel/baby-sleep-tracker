# Sleep Monitor

Aplicación web simple para registrar y visualizar eventos relacionados con el sueño de un bebé.

El proyecto está pensado para ejecutarse de forma local (por ejemplo, en una Raspberry Pi) y prioriza simplicidad, claridad y bajo mantenimiento.

## Funcionalidad

* Registro de eventos de sueño (inicio y fin).
* Visualización de un historial reciente.
* Cálculo básico de duraciones entre eventos.

El sistema está diseñado de forma genérica para permitir la incorporación futura de otros tipos de eventos.

## Arquitectura

El proyecto se compone de:

* Backend encargado de persistir y exponer eventos.
* Base de datos local SQLite.
* Frontend web liviano para la interacción con el usuario.

Para una descripción detallada de las decisiones de diseño, ver el documento de diseño del proyecto.

## Documentación

* [Diseño del proyecto](DESIGN.md)

## Estado del proyecto

Proyecto personal en desarrollo, de alcance reducido y orientado a uso local.

No está pensado como una solución médica ni como un producto de uso general.

