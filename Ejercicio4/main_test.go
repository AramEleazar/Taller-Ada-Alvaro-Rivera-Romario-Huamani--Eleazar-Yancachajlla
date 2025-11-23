package main

import (
	"math"
	"testing"
)

func TestDijkstra_RutaCorrecta(t *testing.T) {
	grafo := map[int]map[int]float64{
		0: {1: 4, 2: 1},
		1: {3: 1},
		2: {1: 2, 3: 5},
		3: {},
	}

	dist, ruta := Dijkstra(grafo, 0, 3)

	if dist != 4 {
		t.Fatalf("Distancia incorrecta. Esperado 4, obtenido %v", dist)
	}

	esperado := []int{0, 2, 1, 3}
	if len(ruta) != len(esperado) {
		t.Fatalf("Ruta incorrecta. Esperado %v, obtenido %v", esperado, ruta)
	}
}

func TestDijkstra_SinRuta(t *testing.T) {
	grafo := map[int]map[int]float64{
		0: {1: 2},
		1: {},
		2: {3: 1},
		3: {},
	}

	dist, ruta := Dijkstra(grafo, 0, 3)

	if !math.IsInf(dist, 1) {
		t.Fatalf("Se esperaba distancia infinita, obtenido %v", dist)
	}

	if ruta != nil {
		t.Fatalf("Se esperaba ruta nil, obtenido %v", ruta)
	}
}
