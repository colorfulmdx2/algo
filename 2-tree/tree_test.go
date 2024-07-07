package tree

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkPostorderTraversal1(b *testing.B) {
	root := generateTree(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostorderTraversal1(root)
	}
}

func BenchmarkPostorderTraversal2(b *testing.B) {
	root := generateTree(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PostorderTraversal2(root)
	}
}

// Generate a random binary 2-tree with the specified number of nodes.
func generateTree(nodeCount int) *TreeNode {
	if nodeCount == 0 {
		return nil
	}

	root := &TreeNode{Val: 1}
	nodeCount--
	currentLevel := []*TreeNode{root}

	rand.Seed(time.Now().UnixNano())
	for nodeCount > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range currentLevel {
			if nodeCount > 0 {
				node.Left = &TreeNode{Val: rand.Int()}
				nextLevel = append(nextLevel, node.Left)
				nodeCount--
			}
			if nodeCount > 0 {
				node.Right = &TreeNode{Val: rand.Int()}
				nextLevel = append(nextLevel, node.Right)
				nodeCount--
			}
		}
		currentLevel = nextLevel
	}

	return root
}
