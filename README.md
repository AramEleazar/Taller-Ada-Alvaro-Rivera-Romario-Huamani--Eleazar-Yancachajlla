# Taller de Grafos – Análisis y Diseño de Algoritmos

**Estudiantes:**  
- Romario Huamni Paccaya  
- Aram Eleazar Yancachajlla Sucso  
- Alvaro Nicolas Rivera Tarque

---

## Ejercicio 1: Red Social Básica (BFS)

### 1. Explicación del Enfoque
El objetivo es encontrar amigos que están a *N* grados de separación de un usuario inicial en una red social.

Para resolverlo, se modeló la red como un grafo no dirigido e implementamos el algoritmo **BFS** (Búsqueda en Anchura):

* **Grafo:** Se usa un `map[string][]string`.
* **Algoritmo:** Utiliza una cola (FIFO) para explorar por capas. Cada nivel representa un grado de separación.
* **Datos:** Se implementó código para leer archivos `.gml` y construir el grafo.

### 2. Análisis de Complejidad

* **Temporal:** En el peor caso, BFS visita cada vértice *(V)* y cada arista *(E)* una vez.  
  Complejidad: **O(V + E)**.
* **Espacial:** Se almacena el grafo, la cola y el mapa de visitados.  
  Complejidad: **O(V + E)**.

### 3. Instrucciones de Ejecución

**Requisitos:**  
Tener instalado Go y el archivo `karate.gml` en la carpeta `Ejercicio1`.

**Comandos (desde la raíz del proyecto `taller-grafos/`):**

```bash
# Ejecutar el programa principal
go run ./Ejercicio1/main.go

# Ejecutar los tests unitarios
go test -v ./Ejercicio1
```

### 4. Casos de Prueba Incluidos
Los tests en `main_test.go` verifican:

* **Grado 0:** El usuario se encuentra a sí mismo.  
* **Grado 1:** Encuentra solo a sus amigos directos.  
* **Grado 2:** Encuentra amigos de amigos sin repeticiones.  
* **Grado inalcanzable:** Si N excede la red, retorna lista vacía.  
* **Nodo inexistente:** Si el usuario inicial no existe, se retorna vacío o error.

---

## Ejercicio 2: Detección de Ciclos en un Grafo No Dirigido (DFS)

### 1. Explicación del Enfoque
Este ejercicio busca determinar si un grafo no dirigido contiene ciclos y, de existir, devolver los nodos que lo componen.

La solución implementa **DFS** con backtracking:

* El grafo se representa mediante una lista de adyacencia.  
* DFS recorre los nodos llevando un arreglo `padre[]`.  
* Si DFS encuentra un nodo visitado que **no es el padre**, existe un ciclo.  
* El ciclo se reconstruye siguiendo la cadena de padres hasta regresar al nodo repetido.

Este método detecta ciclos de forma eficiente incluso en grafos grandes o dispersos.

### 2. Análisis de Complejidad

* **Temporal:** **O(V + E)**  
* **Espacial:** **O(V)**

### 3. Instrucciones de Ejecución

**Requisitos:**  
Tener instalado Go y ubicarse en la carpeta `Ejercicio2`.

**Comandos (desde la raíz `taller-grafos/`):**

```bash
# Ejecutar el programa principal
go run ./Ejercicio2/main.go

# Ejecutar los tests unitarios
go test -v ./Ejercicio2
```

### 4. Casos de Prueba Incluidos
Los tests en `main_test.go` validan dos escenarios principales:

####  Grafo con ciclo
Ejemplo: 1 → 2 → 3 → 1  
Se valida:

* Que se retorne `true`.  
* Que se genere correctamente la lista de nodos del ciclo.

####  Grafo sin ciclo
Ejemplo: grafo lineal.

Se espera:

* Retornar `false`.  
* La lista de nodos debe ser `nil`.

---
---


---

## Ejercicio 4: Navegación GPS (Dijkstra Básico)

### 1. Explicación del Enfoque
El objetivo es encontrar la **ruta más corta** entre dos puntos de una ciudad, considerando que las calles tienen distintas longitudes (pesos).

Para resolverlo, se modela la ciudad como un **grafo ponderado** e implementamos el algoritmo **Dijkstra** optimizado usando un **heap (priority queue)**:

* **Grafo:** Se representa como `map[int]map[int]float64`, donde cada nodo es una intersección.
* **Algoritmo:** Dijkstra selecciona siempre el nodo con la menor distancia acumulada gracias al heap.
* **Ruta:** Se utiliza un mapa `prev[]` para reconstruir la ruta completa desde el destino hacia el origen.
* **Datos:** El grafo puede ser construido con datos reales de un distrito de Lima (según el tutorial indicado).

### 2. Análisis de Complejidad

* **Temporal:**  
  Usando priority queue, la complejidad es:  
  **O((V + E) log V)**  
  Por las operaciones de insertar y extraer elementos del heap.

* **Espacial:**  
  Se almacena el grafo, el arreglo de distancias, predecesores y el heap.  
  Complejidad: **O(V + E)**.

### 3. Instrucciones de Ejecución

**Requisitos:**  
Tener instalado Go y el grafo de prueba dentro de la carpeta `Ejercicio4`.

**Comandos (desde la raíz del proyecto `taller-grafos/`):**

```bash
# Ejecutar el programa principal
go run ./Ejercicio4/main.go

# Ejecutar los tests unitarios
go test -v ./Ejercicio4

### 5. Pasos para ejecutar los casos de prueba

Para validar el correcto funcionamiento del algoritmo, sigue estos pasos:

1. Abre una terminal en la raíz del proyecto `taller-grafos/`.
2. Ejecuta el siguiente comando para correr los tests del Ejercicio 4:

```bash
go test -v ./Ejercicio4
