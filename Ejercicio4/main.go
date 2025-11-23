package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Estructura para la cola de prioridad
type Item struct {
	node     int
	priority float64
	index    int
}

// Implementación del heap (cola de prioridad)
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra: devuelve distancia mínima y la ruta
func Dijkstra(grafo map[int]map[int]float64, origen, destino int) (float64, []int) {

	dist := make(map[int]float64)
	prev := make(map[int]int)

	for nodo := range grafo {
		dist[nodo] = math.Inf(1)
		prev[nodo] = -1
	}

	dist[origen] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: origen, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.node

		if u == destino {
			break
		}

		for v, peso := range grafo[u] {
			alt := dist[u] + peso
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				heap.Push(pq, &Item{node: v, priority: alt})
			}
		}
	}

	// reconstruir ruta
	ruta := []int{}
	cur := destino
	for cur != -1 {
		ruta = append([]int{cur}, ruta...)
		cur = prev[cur]
	}

	if len(ruta) == 1 && ruta[0] != origen {
		return math.Inf(1), nil
	}

	return dist[destino], ruta
}

func main() {

	grafo := map[int]map[int]float64{
		0: {1: 4, 2: 1},
		1: {3: 1},
		2: {1: 2, 3: 5},
		3: {},
	}

	distancia, ruta := Dijkstra(grafo, 0, 3)

	fmt.Println("Distancia mínima:", distancia)
	fmt.Println("Ruta:", ruta)
}
