package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ==============================================================================
// 1. ESTRUCTURAS DE DATOS Y UNION-FIND
// ==============================================================================

// KruskalEdge representa una arista con costo (distancia)
type KruskalEdge struct {
	u    int
	v    int
	cost float64 // Usamos float64 para la distancia euclidiana
}

// NodeCoord almacena las coordenadas (x, y) simuladas de un edificio
type NodeCoord struct {
	x float64
	y float64
}

// UnionFind y sus métodos (Find, Union) se mantienen exactamente igual
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
	}
	for i := 1; i <= n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i])
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) bool {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)

	if rootI != rootJ {
		if uf.rank[rootI] < uf.rank[rootJ] {
			uf.parent[rootI] = rootJ
		} else if uf.rank[rootI] > uf.rank[rootJ] {
			uf.parent[rootJ] = rootI
		} else {
			uf.parent[rootJ] = rootI
			uf.rank[rootI]++
		}
		return true
	}
	return false
}

// ==============================================================================
// 2. FUNCIÓN DE COSTO POR DISTANCIA
// ==============================================================================

// calculateDistance usa la fórmula Euclidiana para simular la longitud/costo
func calculateDistance(c1, c2 NodeCoord) float64 {
	dx := c1.x - c2.x
	dy := c1.y - c2.y
	// Distancia = sqrt((x2-x1)^2 + (y2-y1)^2)
	return math.Sqrt(dx*dx + dy*dy)
}

// ==============================================================================
// 3. ALGORITMO DE KRUSKAL (MST)
// ==============================================================================

// KruskalMST calcula el Árbol de Expansión Mínima (MST)
func KruskalMST(edges []KruskalEdge, n int) (float64, []KruskalEdge) {
	// Paso 1: Ordenar todas las aristas por costo (distancia) ascendente
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].cost < edges[j].cost
	})

	uf := NewUnionFind(n)

	mstConnections := []KruskalEdge{}
	totalCost := 0.0
	edgesCount := 0

	// Paso 2: Iterar y usar Union-Find para evitar ciclos
	for _, edge := range edges {
		if uf.Union(edge.u, edge.v) {
			totalCost += edge.cost
			mstConnections = append(mstConnections, edge)
			edgesCount++

			if edgesCount == n-1 {
				break
			}
		}
	}

	return totalCost, mstConnections
}

// ==============================================================================
// 4. FUNCIÓN PRINCIPAL (LECTURA DE DATOS, ASIGNACIÓN DE COORDENADAS Y EJECUCIÓN)
// ==============================================================================

func main() {
	const filename = "power-US-Grid.mtx"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var N int
	var totalEdges int
	var totalSimulatedCost float64
	var allEdges []KruskalEdge

	// Nuevo: Mapa para almacenar las coordenadas simuladas de cada edificio
	coordinates := make(map[int]NodeCoord)

	rand.Seed(time.Now().UnixNano())

	// --- Parsing y Lectura ---

	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "%") {
			continue
		}

		if lineCount == 0 {
			parts := strings.Fields(line)
			if len(parts) < 3 {
				return
			}

			N, _ = strconv.Atoi(parts[0])
			totalEdges, _ = strconv.Atoi(parts[2])

			// SIMULACIÓN DE COORDENADAS: Asignar coordenadas aleatorias a todos los N edificios
			// Simulamos que el mapa geográfico es de 0 a 1000
			for i := 1; i <= N; i++ {
				coordinates[i] = NodeCoord{
					// Coordenadas aleatorias para simular la posición geográfica
					x: float64(rand.Intn(1000)),
					y: float64(rand.Intn(1000)),
				}
			}
			lineCount++
			continue
		}

		// Líneas de datos: u v
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			u, _ := strconv.Atoi(parts[0])
			v, _ := strconv.Atoi(parts[1])

			// ASIGNACIÓN CLAVE: Costo = Distancia Euclidiana Simulada
			cost := calculateDistance(coordinates[u], coordinates[v])

			totalSimulatedCost += cost

			// Almacenar la arista
			allEdges = append(allEdges, KruskalEdge{u: u, v: v, cost: cost})
		}

		lineCount++
	}

	// --- Ejecución del Algoritmo Kruskal ---

	minCost, mstEdges := KruskalMST(allEdges, N)

	// ==========================================================================
	// 5. OUTPUT
	// ==========================================================================

	fmt.Println("--- Red Eléctrica Óptima (Costo por Distancia Simulada) ---")
	fmt.Printf("Total de Edificios (N): %d\n", N)
	fmt.Printf("Total de Conexiones Posibles (E): %d\n", totalEdges)

	// Salida 3: Costo total si se conectaran todos contra todos (para comparación)
	// Usamos %.2f para redondear a dos decimales
	fmt.Printf("\nCosto Total (Todos contra Todos, basado en Distancia): %.2f\n", totalSimulatedCost)

	// Salida 1: Costo total mínimo
	fmt.Printf("Costo Total Mínimo (Algoritmo de Kruskal - MST): %.2f\n", minCost)

	// Salida 2: Lista de conexiones a instalar
	fmt.Println("\nLista de las primeras 10 Conexiones a Instalar:")

	for i, edge := range mstEdges {
		if i >= 10 {
			break
		}
		fmt.Printf("  - Edificios %d <-> %d (Costo/Longitud: %.2f)\n", edge.u, edge.v, edge.cost)
	}
	if len(mstEdges) > 10 {
		fmt.Printf("  ... y %d conexiones más de un total de %d.\n", len(mstEdges)-10, len(mstEdges))
	}
}
