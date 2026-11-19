package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info reprezinta o pereche (id, valoare) retinuta in memoria unui nod.
type Info struct {
	ID    int
	Value int
}

// Memory este o multime de token-uri (reuniunea se face prin inserare).
type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	Crashed   bool
}

func IntFromLine(l string) []int {
	var r []int
	for _, f := range strings.Fields(l) {
		x, _ := strconv.Atoi(f)
		r = append(r, x)
	}
	return r
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	first := IntFromLine(sc.Text())
	n, s := first[0], first[1]
	sc.Scan()
	tokens := IntFromLine(sc.Text())
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		adj[i] = IntFromLine(sc.Text())
	}

	diam := diametru(adj)
	fmt.Println("Diametru:", diam)

	// M_i(0) = {(i, v_i)}
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{ID: i, Neighbors: adj[i], M: Memory{{i, tokens[i]}: {}}}
	}

	crashes := 0
	t := 0
	for {
		// Maybe alegem un nod care da crash
		crashing := -1
		if crashes < s && rand.Intn(2) == 0 {
			c := rand.Intn(n)
			if !nodes[c].Crashed {
				crashing = c
				crashes++
			}
		}

		// Fiecare nod viu (si cel care va da crash) trimite M_i(t) catre toti vecinii
		primite := make([][]Memory, n)
		for _, nd := range nodes {
			if nd.Crashed {
				continue
			}
			for _, v := range nd.Neighbors {
				primite[v] = append(primite[v], nd.M)
			}
		}
		if crashing != -1 {
			nodes[crashing].Crashed = true
		}

		//  M_i(t+1) = M_i(t) u U
		nou := make([]Memory, n)
		for i, nd := range nodes {
			nou[i] = Memory{}
			for k := range nd.M {
				nou[i][k] = struct{}{}
			}
			if nd.Crashed {
				continue
			}
			for _, m := range primite[i] {
				for k := range m {
					nou[i][k] = struct{}{}
				}
			}
		}
		for i := range nodes {
			nodes[i].M = nou[i]
		}

		// Conditia de oprire
		if t > (s+1)*diam {
			break
		}
		t++
	}

	fmt.Println("Iteratie oprire:", t)

	// Verificam ca toate nodurile vii au aceeasi memorie
	var ref Memory
	consens := true
	for _, nd := range nodes {
		if nd.Crashed {
			continue
		}
		if ref == nil {
			ref = nd.M // Primul nod viu
			lista := []Info{}
			for k := range ref {
				lista = append(lista, k)
			}
			sort.Slice(lista, func(a, b int) bool { return lista[a].ID < lista[b].ID })

			for _, x := range lista { // Afisare continutul nodului viu
				fmt.Printf("(%d, %d) ", x.ID, x.Value)
			}
			fmt.Println()
			continue
		}
		if len(nd.M) != len(ref) {
			consens = false
			break
		}
		for k := range ref {
			if _, ok := nd.M[k]; !ok {
				consens = false
				break
			}
		}
	}
	if consens {
		fmt.Println("Toate nodurile vi au aceeasi memorie!")
	} else {
		fmt.Println("Nodurile vi NU au aceeasi memorie!")
	}
}
