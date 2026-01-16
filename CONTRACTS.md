# Contratos del Dominio

Este documento define los contratos (interfaces conceptuales) necesarios para implementar los casos de uso del sistema.

---

## Principio general

* Los contratos nacen de los casos de uso.
* Las implementaciones concretas (SQLite, memoria, mocks) deben adaptarse a estos contratos.
* El dominio depende de estos contratos, nunca de sus implementaciones.

---

## Caso de uso: Registrar evento

### Necesidades del dominio

Para registrar un nuevo evento respetando las reglas del dominio, el sistema necesita:

1. Conocer el último evento registrado (si existe)
2. Verificar la inexistencia de un evento con el mismo timestamp
3. Persistir un nuevo evento de forma atómica

---

### Contrato: EventStore

Responsable de almacenar y recuperar eventos de la secuencia.

**Responsabilidades:**

* Obtener el último evento registrado
* Verificar la existencia de un evento en un timestamp dado
* Persistir un nuevo evento

**Garantías esperadas:**

* Los eventos recuperados están ordenados cronológicamente
* Las operaciones de escritura son consistentes

---

## Caso de uso: Consultar períodos derivados

### Necesidades del dominio

Para derivar períodos de sueño y vigilia, el sistema necesita:

1. Obtener una secuencia ordenada de eventos
2. Acceder a sus timestamps y tipos

---

### Contrato: EventReader

Responsable de proveer acceso de solo lectura a la secuencia de eventos.

**Responsabilidades:**

* Listar eventos ordenados por tiempo
* Permitir limitar la cantidad de eventos recuperados (opcional)

**Notas:**

* No permite modificaciones
* Puede compartir implementación con `EventStore`

---

## Consideraciones de diseño

* Un mismo componente puede implementar múltiples contratos.
* Los contratos deben ser pequeños y específicos.
* Si un método no es utilizado por un caso de uso, no debe existir.

---

## Qué NO es un contrato

No forman parte de estos contratos:

* Handlers HTTP
* Templates HTML
* DTOs de transporte
* Esquemas de base de datos
* Detalles de serialización

Estos elementos pertenecen a capas externas al dominio.

