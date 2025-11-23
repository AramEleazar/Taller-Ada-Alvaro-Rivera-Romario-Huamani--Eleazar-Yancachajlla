package main

import (
	"fmt"
)

// Detectar ciclo en un grafo no dirigido usando DFS con backtracking
func DetectarCiclo(n int, edges [][2]int) (bool, []int) {
	// Crear lista de adyacencia
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	visitado := make([]bool, n)
	padre := make([]int, n)
	for i := 0; i < n; i++ {
		padre[i] = -1
	}

	// Para reconstruir ciclo
	var ciclo []int

	var dfs func(u int) bool
	dfs = func(u int) bool {
		visitado[u] = true

		for _, v := range adj[u] {
			if !visitado[v] {
				padre[v] = u
				if dfs(v) {
					return true
				}
			} else if v != padre[u] {
				// Ciclo encontrado: reconstruir
				ciclo = append(ciclo, v)
				cur := u
				for cur != v {
					ciclo = append(ciclo, cur)
					cur = padre[cur]
				}
				return true
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		if !visitado[i] {
			if dfs(i) {
				return true, ciclo
			}
		}
	}

	return false, nil
}

func main() {
	n := 5
	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}}

	tieneCiclo, nodos := DetectarCiclo(n, edges)
	fmt.Println("¿Tiene ciclo?", tieneCiclo)
	if tieneCiclo {
		fmt.Println("Ciclo encontrado:", nodos)
	}
}
