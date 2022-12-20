package problem

import (
	"fmt"
	"sort"
)

//@todo to rewrite this is a stolen code i don't understand how it works

type Person struct {
	h int // height
	k int // number of people in front >= height
}

func (p *Person) String() string {
	return fmt.Sprintf("h: %d, k: %d\n", p.h, p.k)
}

type People []*Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].h == p[j].h {
		return p[i].k < p[j].k
	}

	return p[i].h >= p[j].h // Desc order
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func QueueReconstructionByHeight(input [][]int) [][]int {
	var p People
	for _, person := range input {
		p = append(p, &Person{
			h: person[0],
			k: person[1],
		})
	}
	sort.Sort(p)

	for i := 0; i < len(input); i++ {
		current := p[i]

		for j := current.k; j <= i; j++ {
			temp := p[j]
			p[j] = current
			current = temp
		}
	}

	var res [][]int

	for _, person := range p {
		res = append(res, []int{person.h, person.k})
	}

	return res
}
