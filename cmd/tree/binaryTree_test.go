package tree

import (
	"reflect"
	"testing"
)

func TestNewNodeTree(t *testing.T) {
	value := 3
	rootNode := NewNodeTree(nil, nil, value)

	got := rootNode.GetValue()
	if got != value {
		t.Errorf("got %d, wanted %d", got, value)
	}
}

func TestInsertNewNodeToTheLeft(t *testing.T) {
	rootNode := NewNodeTree(nil, nil, 3)

	newValue := 1
	rootNode.InsertNewNode(newValue)
	got := (*rootNode.leftNode).GetValue()
	if got != newValue {
		t.Errorf("got %d, wanted %d", got, newValue)
	}
}

func TestInsertNewNodeToTheRight(t *testing.T) {
	rootNode := NewNodeTree(nil, nil, 3)

	newValue := 5
	rootNode.InsertNewNode(newValue)
	got := (*rootNode.rightNode).GetValue()
	if got != newValue {
		t.Errorf("got %d, wanted %d", got, newValue)
	}
}

type insertNodesTestCase[T Node] struct {
	input []T
	want  *BinaryTree[T]
}

func getInsertStringNodesTestCases() []insertNodesTestCase[string] {
	valueListTestCase1 := []string{"Noah", "Oliver", "Eli", "James", "Sophie", "Will", "Omar", "Rachel"}
	caseNodeJames := NewNodeTree(nil, nil, "James")
	caseNodeEli := NewNodeTree(nil, caseNodeJames, "Eli")

	caseNodeRachel := NewNodeTree(nil, nil, "Rachel")
	caseNodeOmar := NewNodeTree(nil, caseNodeRachel, "Omar")
	caseNodeWill := NewNodeTree(nil, nil, "Will")
	caseNodeSophie := NewNodeTree(caseNodeOmar, caseNodeWill, "Sophie")
	caseNodeOliver := NewNodeTree(nil, caseNodeSophie, "Oliver")
	treeTestCase1 := NewNodeTree(caseNodeEli, caseNodeOliver, "Noah")

	insertNodeTestCase1 := insertNodesTestCase[string]{input: valueListTestCase1, want: treeTestCase1}

	insertNodeTests := []insertNodesTestCase[string]{
		insertNodeTestCase1,
	}
	return insertNodeTests
}

func TestInsertStringNodes(t *testing.T) {
	insertNodeTests := getInsertStringNodesTestCases()
	for _, testCase := range insertNodeTests {
		tree_ := NewNodeTree(nil, nil, "")
		for _, value := range testCase.input {
			tree_.InsertNewNode(value)
		}
		if reflect.DeepEqual(testCase.want, tree_) == false {
			t.Fail()
		}
		tree_ = nil
	}
}

func getInsertIntegerNodesTestCases() []insertNodesTestCase[int] {
	valueListTestCase1 := []int{25, 22, 40, 30, 45, 27, 20, 21, 48}

	case1Node21 := NewNodeTree(nil, nil, 21)
	case1Node20 := NewNodeTree(nil, case1Node21, 20)
	case1Node22 := NewNodeTree(case1Node20, nil, 22)
	case1Node27 := NewNodeTree(nil, nil, 27)
	case1Node30 := NewNodeTree(case1Node27, nil, 30)
	case1Node48 := NewNodeTree(nil, nil, 48)
	case1Node45 := NewNodeTree(nil, case1Node48, 45)
	case1node40 := NewNodeTree(case1Node30, case1Node45, 40)
	treeTestCase1 := NewNodeTree(case1Node22, case1node40, 25)

	insertNodeTestCase1 := insertNodesTestCase[int]{input: valueListTestCase1, want: treeTestCase1}

	// test case 2
	valueListTestCase2 := []int{40, 25, 20, 30, 45, 27, 22, 21, 48}
	case2Node21 := NewNodeTree(nil, nil, 21)
	case2Node22 := NewNodeTree(case2Node21, nil, 22)
	case2Node20 := NewNodeTree(nil, case2Node22, 20)
	case2Node27 := NewNodeTree(nil, nil, 27)
	case2Node30 := NewNodeTree(case2Node27, nil, 30)
	case2Node25 := NewNodeTree(case2Node20, case2Node30, 25)
	case2Node48 := NewNodeTree(nil, nil, 48)
	case2Node45 := NewNodeTree(nil, case2Node48, 45)
	treeTestCase2 := NewNodeTree(case2Node25, case2Node45, 40)

	insertNodeTestCase2 := insertNodesTestCase[int]{valueListTestCase2, treeTestCase2}

	// final list
	insertNodeTests := []insertNodesTestCase[int]{
		insertNodeTestCase1,
		insertNodeTestCase2,
	}
	return insertNodeTests
}

func TestInsertIntegerNodes(t *testing.T) {
	insertNodeTests := getInsertIntegerNodesTestCases()
	for _, testCase := range insertNodeTests {
		tree_ := NewNodeTree(nil, nil, 0)
		for _, value := range testCase.input {
			tree_.InsertNewNode(value)
		}
		if reflect.DeepEqual(testCase.want, tree_) == false {
			t.Fail()
		}
		tree_ = nil
	}
}