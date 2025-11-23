package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yaricom/goGraphML/graphml"
)

// Estructura del grafo como diccionarios, clave: ID de una persoma, valor: los amigos que tiene
type Graph map[string][]string

// busca los nodos que estén a n grados del usuario inicial en el grafo
func BFS(g Graph, nodoInicial string, n int) ([]string, error) {
	//validaciones iniciales
	//si no existe en el grafo el usuario inicial, retorna un mensaje indicando que no existe
	if _, exists := g[nodoInicial]; !exists {
		return nil, fmt.Errorf("el nodo inicial '%s' no existe en el grafo", nodoInicial)
	}
	//Si el grado es negativo, mensaje de error, no puede ser negativo
	if n < 0 {
		return nil, fmt.Errorf("los grados de separación deben ser no negativos")
	}
	//si el grado es 0, retorna  el grafo del usuario inicial, no recorre nada
	if n == 0 {
		return []string{nodoInicial}, nil
	}

	//estructura de cada nodo, id y su grado
	type nodeState struct {
		id   string
		dist int
	}
	//cola para el BFS, almacena los nodos, con su id y grado, inicializa con el usuario inicial
	cola := []nodeState{{id: nodoInicial, dist: 0}}

	//diccionario que almacena los nodos ya visitados, incializa con el usuario inicial, true
	visited := make(map[string]bool)
	visited[nodoInicial] = true

	//almacenar los nodos que estén en el grado 2
	var result []string

	//bucle para el analisis del grafo, hasta vaciar la cola
	for len(cola) > 0 {
		//se saca al primer elemento de la cola
		actual := cola[0]
		cola = cola[1:]

		//si el nodo actual se encuentra en el grado objetivo(n)
		//se agrega a result ese nodo, se ignoran los otros nodos que están a n+1 grados
		if actual.dist == n {
			result = append(result, actual.id)
			continue
		}

		//si el grado es mayor a n, se ignora
		if actual.dist > n {
			continue
		}

		//bucle para explorar los vecinos del nodo actual
		for _, vecino := range g[actual.id] {
			//si los nodos vecinos ya fueron visitados, se ignoran
			//si no se agregan como nodos ya visitados
			if !visited[vecino] {
				visited[vecino] = true
				//se agregan a la cola la cola anterior y a los vecinos que aún no fueron visitidos
				//y se les asigna +1 grado, porque son vecinos del nodo actual
				cola = append(cola, nodeState{id: vecino, dist: actual.dist + 1})
			}
		}
	}
	//retorna el resultado
	return result, nil
}

// cargarGrafo lee el archivo y construye el mapa de adyacencia.
// Lee el archivo karate.gml y construye un mapa de adyacencia
func cargarGrafo(filePath string) (Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// CORRECCIÓN:
	// 1. Instanciamos la estructura GraphML (el constructor suele pedir una descripción)
	gmlDoc := graphml.NewGraphML("Social Network")

	// 2. Usamos el método .Decode() pasando el archivo (io.Reader)
	if err := gmlDoc.Decode(file); err != nil {
		return nil, fmt.Errorf("error al decodificar GML: %v", err)
	}

	// Asumimos que hay al menos un grafo en el archivo
	if len(gmlDoc.Graphs) == 0 {
		return nil, fmt.Errorf("no se encontraron grafos en el archivo")
	}

	g := gmlDoc.Graphs[0]
	adjList := make(Graph)

	// 1. Registrar todos los nodos
	for _, node := range g.Nodes {
		// Usamos node.ID como identificador
		adjList[node.ID] = []string{}
	}

	// 2. Procesar las aristas (edges)
	for _, edge := range g.Edges {
		source := edge.Source
		target := edge.Target

		// Verificar que los nodos existan en el mapa (por seguridad)
		if _, ok := adjList[source]; ok {
			adjList[source] = append(adjList[source], target)
		}
		if _, ok := adjList[target]; ok {
			adjList[target] = append(adjList[target], source)
		}
	}

	return adjList, nil
}

func main() {
	//nombre del archivo a leer
	nombreArchivo := "karate.gml"
	//asignar un ID inicial
	nodoInicial := "1"
	//definir el grado a buscar
	grado := 2
	//construir el mapa de adyacencia
	graph, err := cargarGrafo(nombreArchivo)
	//si se genera un error en la construcción
	if err != nil {
		log.Fatalf("Error al cargar el grafo: %v", err)
	}

	fmt.Printf("Grafo cargado exitosamente. Nodos: %d\n", len(graph))
	fmt.Printf("Buscando usuarios a exactamente %d grados de '%s'...\n", grado, nodoInicial)

	//buscar los amigos que estén en el grado 2
	amigos, err := BFS(graph, nodoInicial, grado)
	//si se produce un error en la busqueda
	if err != nil {
		log.Fatalf("Error en BFS: %v", err)
	}

	//imprimir los resultados de la busqueda
	fmt.Printf("Resultado: %v\n", amigos)
	fmt.Printf("Total encontrados: %d\n", len(amigos))
}
