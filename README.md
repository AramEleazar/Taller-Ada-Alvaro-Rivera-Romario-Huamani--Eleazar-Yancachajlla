# Taller de Grafos - Análisis y Diseño de Algoritmos

**Estudiantes:** Romario Huamni Paccaya  
**Estudiantes:** Aram Eleazar Yancachajlla Sucso
---

## Ejercicio 1: Red Social Básica (BFS)

### 1. Explicación del Enfoque
El objetivo es encontrar amigos que están a **N** grados de separación de un usuario inicial en una red social. 

Para resolverlo, se modeló la red como un grafo no dirigido y se implementó el algoritmo **BFS** o Búsqueda en Anchura.
* **Grafo:** Se usa un `map[string][]string`.
* **Algoritmo:** Utilizamos una cola (FIFO) para explorar el grafo por capas  . Cada nivel representa un grado de separación.
* **Datos:** Se implementó código encargado de leer archivos **.gml** y construir el grafo.

### 2. Análisis de Complejidad

* **Temporal:** En el peor de los casos, el BFS visita cada vértice una vez **(V)** y, por consecuencia, revisa cada arista **(E)** una sola vez. Lo que resulta en una complejidad de **O(V + E)**

* **Espacial:** Se necesita espacio para almacenar el grafo completo y vaariables auxiliares como el mapa de **visitado** y la cola. Todo ello resulta en una complejidad de **O(V + E)**.

### 3. Instrucciones de Ejecución

**Requisitos:** Tener instalado Go y el archivo **karate.gml** en la carpeta **ejercicio01**.

**Comandos:**

Desde la raíz del proyecto (`taller-grafos/`):

```bash
# Ejecutar el programa principal
go run ./ejercicio01/main.go

# Ejecutar los tests unitarios
go test -v ./ejercicio01
```

### 4. Casos de Prueba Incluidos
Los tests en **main_test.go** verifican la lógica del algoritmo BFS usando un grafos de ejemplos ya incluidos en el código:

* **Grado 0:** El usuario debe encontrarse a sí mismo.

* **Grado 1:** Debe encontrar solo a los amigos directos.

* **Grado 2:** Debe encontrar amigos de amigos sin repeticiones y excluyendo al usuario incial.

* **Grado Inalcanzable:** Si N es mayor al tamaño de la red, debe retornar lista vacía.

* **Nodo Inexistente:** Manejo de error si el usuario inicial no existe, no retorna nada.

----
## Ejercicio 2: Detección de Ciclos en un Grafo No Dirigido (DFS)

### 1. Explicación del Enfoque
Este ejercicio consiste en determinar si un grafo no dirigido contiene al menos un ciclo y, en caso afirmativo, retornar los nodos que forman dicho ciclo.

Para resolverlo, se implementó el algoritmo DFS (Depth-First Search) con backtracking, siguiendo esta lógica:

El grafo se representa mediante una lista de adyacencia.

Se recorre el grafo usando DFS.

Para evitar falsos positivos, se lleva un arreglo padre[] que mantiene la referencia del nodo desde el que se llegó.

Cuando DFS encuentra un nodo ya visitado que no es el padre, significa que existe un ciclo.

Se reconstruye el ciclo recorriendo la cadena de padres hasta volver al nodo repetido.

Este enfoque permite detectar ciclos correctamente incluso en grafos grandes y dispersos.

### 2. Análisis de Complejidad

* **Temporal:** Complejidad temporal=O(V+E)

* **Espacial:** Complejidad espacial=O(V)

### 3. Instrucciones de Ejecución

**Requisitos:** Tener instalado Go y estar en la carpeta **ejercicio02**.

**Comandos:**

Desde la raíz del proyecto (`taller-grafos/`):

```bash
# Ejecutar el programa principal
go run ./Ejercicio2/main.go

# Ejecutar los tests unitarios
go test -v ./Ejercicio2
```

### 4. Casos de Prueba Incluidos
Los tests en main_test.go validan dos escenarios fundamentales:

* **Grafo con ciclo:** Se usa un grafo donde existe un ciclo 1 → 2 → 3 → 1. 
El test verifica:

Que la función retorne true.

Que se genere correctamente una lista de nodos que forman parte del ciclo.

* **Grafo sin ciclo:** Se prueba con un grafo lineal (sin ciclos).

La función debe retornar:

false indicando que no hay ciclo.

nil como lista de nodos.


