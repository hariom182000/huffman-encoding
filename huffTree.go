package main

import "container/heap"

type Node struct {
    char   rune
	weight int
    left    *Node
    right  *Node
	index    int
}

type MinHeap []*Node



func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight 
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MinHeap) Push(x interface{}) {
	n := len(*h)
	node := x.(*Node)
	node.index = n
	*h = append(*h, node)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	node := old[n-1]
	*h = old[0 : n-1]
	return node
}

func newItem(  char   rune, weight int) *Node {
	return &Node{
		char:    char,
		weight: weight,
	}
}

func CreateTree(frequencyMap map[rune]int) *Node {

	h := &MinHeap{}
	heap.Init(h)
	for char, freq := range frequencyMap {
		heap.Push(h, newItem(char, freq))
	}
	for h.Len()>1 {
		left:=heap.Pop(h).(*Node);
		right:=heap.Pop(h).(*Node);
		merged := &Node{
			weight: left.weight + right.weight,
			left:   left,
			right:  right,
		}
		heap.Push(h,merged)
	}
	return heap.Pop(h).(*Node)
}


func GetPrefixTable(tree *Node ,prefixTable map[rune]string, path string ) {

if(tree.left==nil && tree.right==nil){
	prefixTable[tree.char]=path;
	return;
}
GetPrefixTable(tree.left,prefixTable,path+"0");
GetPrefixTable(tree.right,prefixTable,path+"1");

}