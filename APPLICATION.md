# Application Layer

Este documento define la capa de aplicación del sistema.

La capa de aplicación orquesta los casos de uso del dominio. No contiene lógica de transporte (HTTP, UI) ni detalles de persistencia (SQLite, SQL). Su responsabilidad es coordinar contratos del dominio y aplicar reglas de negocio.

---

## Responsabilidades

La capa de aplicación:

* Expone casos de uso claros y explícitos
* Valida reglas del dominio
* Coordina acceso a contratos (lectura y escritura de eventos)
* Devuelve resultados o errores de dominio

No renderiza HTML ni conoce protocolos de transporte.

---

## Caso de uso: Registrar evento

### Descripción

Registra un nuevo evento en la secuencia respetando todas las reglas del dominio.

### Flujo

1. Obtener el último evento registrado (si existe)
2. Validar que el tipo de evento sea consistente con la secuencia
3. Validar que el timestamp sea posterior al último evento
4. Validar que no exista otro evento con el mismo timestamp
5. Persistir el nuevo evento

### Entradas

* Tipo de evento
* Timestamp del evento

### Salida

* Evento registrado

### Errores de dominio posibles

* SecuenciaVacíaInvalida (primer evento distinto de SUEÑO_FIN)
* EventoFueraDeOrden
* EventoDuplicadoEnTiempo
* TipoDeEventoInvalido

---

## Caso de uso: Consultar períodos

### Descripción

Obtiene la secuencia de eventos y deriva períodos continuos de sueño o vigilia.

### Flujo

1. Obtener la secuencia ordenada de eventos
2. Iterar eventos consecutivos
3. Derivar períodos a partir de pares de eventos

### Entradas

* Cantidad máxima de eventos (opcional)

### Salida

* Lista de períodos derivados

Cada período incluye:

* Timestamp de inicio
* Timestamp de fin
* Tipo de período (durmiendo / despierto)
* Duración calculada

### Errores posibles

* SecuenciaInsuficiente (menos de dos eventos)

---

## Consideraciones de diseño

* La capa de aplicación no decide cómo se muestran los datos
* No persiste resultados derivados
* No reordena eventos recibidos de los contratos
* Los errores son explícitos y semánticos

---

## Qué NO pertenece a esta capa

* Handlers HTTP
* Templates HTML
* Serialización
* SQL o consultas
* Validaciones de UI

