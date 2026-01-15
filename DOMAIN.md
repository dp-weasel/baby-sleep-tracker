# Dominio: Baby Sleep Tracking

## Objetivo del dominio

Modelar y registrar eventos temporales relacionados con el descanso de un bebé, manteniendo una secuencia consistente e inmutable de hechos a partir de la cual se puede inferir información como períodos de sueño y vigilia.

El dominio no persiste estados ni duraciones. Toda información derivada se calcula a partir de la secuencia de eventos registrados.

---

## Conceptos principales

### Evento

Representa un hecho ocurrido en un instante específico del tiempo.

Un evento:

* Ocurre en un único timestamp
* Es inmutable una vez registrado
* Pertenece a un tipo de evento

### Tipo de evento

Clasificación que define el significado del evento dentro del dominio.

En la versión inicial del sistema existen eventos relacionados exclusivamente con el sueño, pero el dominio está preparado para admitir nuevos tipos en el futuro.

### Secuencia de eventos

Conjunto ordenado de eventos que representa la historia registrada.

La secuencia es la única fuente de verdad del sistema. No existe un estado actual persistido.

---

## Tipos de eventos

### SUEÑO_INICIO

Indica que el bebé se durmió.

### SUEÑO_FIN

Indica que el bebé se despertó.

---

## Reglas e invariantes del dominio

Las siguientes reglas deben cumplirse en todo momento:

* Los eventos son inmutables una vez registrados.
* No pueden existir dos eventos con el mismo timestamp.
* Los eventos deben estar estrictamente ordenados en el tiempo.
* No pueden existir dos eventos consecutivos del mismo tipo.
* El primer evento del sistema debe ser siempre `SUEÑO_FIN`.
* No se persiste ningún estado derivado (por ejemplo, "está durmiendo").
* Las duraciones entre eventos no se almacenan; se calculan dinámicamente.

Si una operación viola cualquiera de estas reglas, el dominio debe rechazarla.

---

## Inferencia de períodos y duraciones

A partir de la secuencia de eventos, el dominio permite inferir períodos continuos de tiempo entre dos eventos consecutivos.

Cada período se caracteriza por:

* Hora de inicio (timestamp del evento anterior)
* Hora de fin (timestamp del evento siguiente)
* Tipo de período (despierto o durmiendo)
* Duración calculada como la diferencia entre ambos timestamps

Estas inferencias existen solo en tiempo de consulta y nunca se persisten.

---

## Secuencias válidas e inválidas

### Secuencia válida

```
SUEÑO_FIN    07:10
SUEÑO_INICIO 08:50
SUEÑO_FIN    11:10
```

Inferencias posibles:

* 07:10 – 08:50: despierto (1h 40m)
* 08:50 – 11:10: durmiendo (2h 20m)

---

### Secuencias inválidas

Dos eventos consecutivos del mismo tipo:

```
SUEÑO_INICIO 08:50
SUEÑO_INICIO 09:30
```

Eventos fuera de orden temporal:

```
SUEÑO_FIN    09:00
SUEÑO_INICIO 08:30
```

Eventos con timestamps duplicados:

```
SUEÑO_FIN    08:00
SUEÑO_INICIO 08:00
```

---

## Casos límite definidos

### Primer evento del sistema

Antes del primer `SUEÑO_FIN` no existe historial ni se pueden calcular duraciones.

### Registro de eventos pasados

El dominio permite registrar eventos con timestamps en el pasado siempre que:

* Respeten el orden temporal
* Cumplan todas las reglas e invariantes definidas

El origen del timestamp (evento actual o carga manual) es irrelevante para el dominio.

