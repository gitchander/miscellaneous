package mscore

import (
	"fmt"

	"mvstick/utils"
)

//------------------------------------------------------------------------------

type stepLayer struct {
	nodes []*layerNode
}

type layerNode struct {
	prevIndex int // prev layer node index
	move      Move
	poses     []int
	solved    bool
}

//------------------------------------------------------------------------------

type posesHasher struct {
	g    *Grid
	data []byte
	bsh  *utils.BytesHasher
}

func newPosesHasher(g *Grid) *posesHasher {
	return &posesHasher{
		g:    g,
		data: make([]byte, encodePosesLen(g.bs)),
		bsh:  utils.NewBytesHasher(),
	}
}

func (p *posesHasher) AddIfNotExist() bool {
	n := encodePoses(p.data, p.g.bs)
	return p.bsh.AddIfNotExist(p.data[:n])
}

//------------------------------------------------------------------------------

func Solve(g *Grid) {

	startNode := &layerNode{
		poses:  gridGetPoses(g),
		solved: g.Solved(),
	}

	ph := newPosesHasher(g)
	ph.AddIfNotExist()

	layers := []*stepLayer{
		&stepLayer{
			nodes: []*layerNode{
				startNode,
			},
		},
	}

	const maxLevel = 100

	for level := 0; level < maxLevel; level++ {

		lastLayer := layers[len(layers)-1]

		newLayer := calcNewLayer(g, ph, lastLayer)
		if newLayer == nil {
			break
		}

		layers = append(layers, newLayer)

		var countSolved int
		for nodeIndex, node := range newLayer.nodes {
			if node.solved {
				countSolved++

				//---------------------------------------------------------
				var moves []Move // solution moves
				k := nodeIndex

				for ls := layers; len(ls) > 1; ls = ls[:len(ls)-1] {
					n := ls[len(ls)-1].nodes[k]
					moves = append(moves, n.move)
					k = n.prevIndex
				}
				reverseSlice(moves)
				fmt.Println("moves:", moves)
				fmt.Println("moves count:", len(moves))
				//---------------------------------------------------------

			}
		}
		if countSolved > 0 {
			break
		}
	}

	// if true {
	// 	fmt.Println("count hashes:", len(ph.hashes))
	// 	count := 0
	// 	for _, layer := range layers {
	// 		count += len(layer.nodes)
	// 		fmt.Println(">:", len(layer.nodes))
	// 	}
	// 	fmt.Println("count:", count)
	// }

	gridSetPoses(g, startNode.poses)
}

func calcNewLayer(g *Grid, ph *posesHasher, lastLayer *stepLayer) *stepLayer {

	newLayer := new(stepLayer)

	for nodeIndex, node := range lastLayer.nodes {

		gridSetPoses(g, node.poses)

		// walk all bricks
		for brickIndex := range g.bs {

			// walk brick dec
			for step := -1; ; step-- {
				m := MakeMove(brickIndex, step)
				ok := calcNewNode(g, ph, nodeIndex, m, newLayer)
				if !ok {
					break
				}
			}

			// walk brick inc
			for step := +1; ; step++ {
				m := MakeMove(brickIndex, step)
				ok := calcNewNode(g, ph, nodeIndex, m, newLayer)
				if !ok {
					break
				}
			}
		}
	}

	if len(newLayer.nodes) == 0 {
		return nil
	}

	return newLayer
}

func calcNewNode(g *Grid, ph *posesHasher, prevIndex int, m Move, newLayer *stepLayer) bool {

	err := g.MoveBrick(m)
	if err != nil {
		return false
	}

	if ph.AddIfNotExist() {
		newNode := &layerNode{
			prevIndex: prevIndex,
			move:      m,
			poses:     gridGetPoses(g),
			solved:    g.Solved(),
		}
		newLayer.nodes = append(newLayer.nodes, newNode)
	}

	g.MoveBrick(m.Reverse()) // move back

	return true
}
