package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestBFS(t *testing.T) {
	//grafo pequeño de prueba
	mockGraph := Graph{
		"A": {"B", "D"},
		"B": {"A", "C", "E"},
		"C": {"B"},
		"D": {"A", "E"},
		"E": {"B", "D"},
	}

	tests := []struct {
		name        string
		nodoInicial string
		n           int
		want        []string
		wantErr     bool
	}{
		{
			name:        "Grado 0 (Mismo nodo)",
			nodoInicial: "A",
			n:           0,
			want:        []string{"A"},
			wantErr:     false,
		},
		{
			name:        "Grado 1 (Vecinos directos)",
			nodoInicial: "A",
			n:           1,
			want:        []string{"B", "D"},
			wantErr:     false,
		},
		{
			name:        "Grado 2 (Amigos de amigos)",
			nodoInicial: "A",
			n:           2,
			want:    []string{"C", "E"},
			wantErr: false,
		},
		{
			name:        "Grado inalcanzable",
			nodoInicial: "A",
			n:           5,
			want:        nil, // O slice vacío
			wantErr:     false,
		},
		{
			name:        "Nodo inexistente",
			nodoInicial: "Z",
			n:           1,
			want:        nil,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BFS(mockGraph, tt.nodoInicial, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("BFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Ordenamos los slices para comparar contenidos sin importar el orden
			sort.Strings(got)
			sort.Strings(tt.want)

			// Manejo de nil vs slice vacío para reflect.DeepEqual
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
