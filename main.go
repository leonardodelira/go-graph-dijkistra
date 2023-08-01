package main

import (
	"container/heap"
	"fmt"
)

func main() {
	graph := map[int][]Edge{
		0: {{1, 4}, {2, 2}},
		1: {{3, 5}},
		2: {{1, 1}, {3, 3}},
		3: {{4, 3}},
		4: {},
	}

	source := 0
	target := 3
	distance, path := dijkstra(graph, source, target)

	fmt.Printf("Distância de %d para %d: %d\n", source, target, distance)
	fmt.Printf("Caminho mais curto: %v\n", path)
}

func dijkstra(graph map[int][]Edge, source int, target int) (int, []int) {
	dist := make(map[int]int)
	prev := make(map[int]int)
	visited := make(map[int]bool)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for node := range graph {
		dist[node] = -1
	}
	dist[source] = 0
	heap.Push(&pq, &Node{ID: source, Distance: 0})

	for pq.Len() > 0 {
		//Extrai o nó com menor distancia
		current := heap.Pop(&pq).(*Node)

		if visited[current.ID] {
			continue
		}

		visited[current.ID] = true

		//Relaxamento das arestas, ou seja, vejo com quais outros nós meu nó atual está ligado.
		for _, edge := range graph[current.ID] {
			//Calcula a distancia até o próximo nó
			distance := current.Distance + edge.Weight

			//Se a nova distancia é menor que a distancia atual para o próximo nó, atualiza.
			if dist[edge.To] == -1 || distance < dist[edge.To] {
				//Atuliza a distancia para o próximo nó e também marca que o nó atual está ligado com o próximo nó
				dist[edge.To] = distance
				prev[edge.To] = current.ID

				//Adiciona o novo nó a fila de prioridade
				heap.Push(&pq, &Node{ID: edge.To, Distance: distance})
			}
		}
	}

	//Constrói o caminho do nó de origem para o nó de destino
	/*
		Ex do prev:
		(inicio, source) 1 -> 4 -> 5 -> 7 (fim, target)
		Fazemos uma lógica começando do node "target" até chegar no "source". Isso é possível porque no algoritmo acima nós sempre atualizamos
		qual era o node "prev" de cada elemento, considerando sempre o menor valor.
	*/
	path := []int{}
	node := target
	for node != source {
		path = append([]int{node}, path...)
		node = prev[node]
	}
	path = append([]int{source}, path...)

	return dist[target], path
}
