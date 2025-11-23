package main

import "testing"

func TestDetectarCiclo_ConCiclo(t *testing.T) {
	n := 5
	edges := [][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 1}, // crea ciclo
	}

	tieneCiclo, ciclo := DetectarCiclo(n, edges)

	if !tieneCiclo {
		t.Fatalf("Se esperaba ciclo, pero DetectarCiclo retornó false")
	}

	if len(ciclo) == 0 {
		t.Fatalf("Se esperaba una lista de nodos del ciclo, pero retornó vacía")
	}
}

func TestDetectarCiclo_SinCiclo(t *testing.T) {
	n := 4
	edges := [][2]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}

	tieneCiclo, ciclo := DetectarCiclo(n, edges)

	if tieneCiclo {
		t.Fatalf("No debe detectar ciclo, pero retornó true")
	}

	if ciclo != nil {
		t.Fatalf("No debe retornar nodos si no hay ciclo, pero retornó %v", ciclo)
	}
}
