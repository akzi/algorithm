package aho_corasick_automation

type (
	Node struct {
		output bool
		c      byte
		path   string
		child  [256]*Node
		fail   *Node
	}
	AhoCorasickAutomation struct {
		root    *Node
		output  []*Node
		size    int
		dictMap []string
	}
)

func NewAhoCorasickAutomation() *AhoCorasickAutomation {
	return &AhoCorasickAutomation{
		root: &Node{},
	}
}

func (ahoCorasickAutomation *AhoCorasickAutomation) Append(patten string) {
	node := ahoCorasickAutomation.root
	strLen := len(patten)
	if strLen == 0 {
		return
	}
	path := make([]byte, 0, len(patten))
	for i := 0; i < strLen; i++ {
		b := patten[i]
		path = append(path, b)
		if node.child[b] == nil {
			child := &Node{
				c:    b,
				path: string(path),
			}
			node.child[b] = child
		}
		node = node.child[b]
	}
	node.output = true
	ahoCorasickAutomation.output = append(ahoCorasickAutomation.output, node)
}

func (ahoCorasickAutomation *AhoCorasickAutomation) Build() {
	var queue []*Node
	ahoCorasickAutomation.root.fail = ahoCorasickAutomation.root
	for _, node := range ahoCorasickAutomation.root.child {
		if node == nil {
			continue
		}
		node.fail = ahoCorasickAutomation.root
		queue = append(queue, node)
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		for i := 0; i < 256; i++ {
			child := node.child[i]
			if child == nil {
				continue
			}
			fail := node.fail
			for fail != nil {
				if fail.child[i] != nil {
					child.fail = fail.child[i]
					break
				}
				if fail.fail == ahoCorasickAutomation.root {
					child.fail = ahoCorasickAutomation.root
					break
				}
				fail = fail.fail
			}
			queue = append(queue, child)
		}
	}
}

func (ahoCorasickAutomation *AhoCorasickAutomation) Search(str string) [][]int {
	var result [][]int
	node := ahoCorasickAutomation.root
	strLen := len(str)
	for i := 0; i < strLen; {
		b := str[i]
		child := node.child[b]
		if child == nil {
			if node.fail != ahoCorasickAutomation.root {
				node = node.fail
				continue
			}
			if node != ahoCorasickAutomation.root &&
				ahoCorasickAutomation.root.child[b] != nil {
				node = node.fail
				continue
			}
			i++
			node = node.fail
			continue
		}
		i++
		node = child
		if node.output {
			result = append(result, []int{i - len(node.path), len(node.path)})
			fail := node.fail
			for fail != ahoCorasickAutomation.root {
				if fail.output {
					result = append(result, []int{i - len(fail.path), len(fail.path)})
				}
				fail = fail.fail
			}
		}
	}
	return result
}
