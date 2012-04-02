package main

import (
        "fmt"
)

const (
        ALPHABET_SIZE = 26
)

// trie node
type trieNode struct {
        value int
        children []*trieNode
        isLeaf bool
}

// trie ADT
type trie struct {
        root *trieNode
        count int
}

// Initializes trie (root is dummy node)
func newTrie() *trie {
        c := make([]*trieNode, ALPHABET_SIZE)
        root := &trieNode{0, c, false}
        return &trie{root, 0}
}

// If not present, inserts key into trie
// If the key is prefix of trie node, marks it as leaf node
func (t *trie) insert(key string) {
        if !t.search(key) {
                nodePointer := t.root
                t.count++
                for i := 0; i < len(key); i++ {
                        index := key[i] - 'a'
                        if nodePointer.children[index] == nil {
                                c := make([]*trieNode, ALPHABET_SIZE)
                                node := &trieNode{1, c, false}
                                nodePointer.children[index] = node
                        } else {
                                nodePointer.children[index].value++
                        }
                        nodePointer = nodePointer.children[index]
                }
                nodePointer.isLeaf = true
        }
}

// Returns non zero, if key presents in trie
func (t *trie) search(key string) bool {
        nodePointer := t.root
        for i := 0; i < len(key); i++ {
                index := key[i] - 'a'
                if nodePointer.children[index] == nil {
                        return false
                }
                nodePointer = nodePointer.children[index]
        }

        // Find the leaf
        if nodePointer != nil {
                return nodePointer.isLeaf
        }
        return false
}

// Recursively delete the node
func deleteHelper(node *trieNode, key string, level int) {
        index := key[level] - 'a'
        if level != len(key) - 1 { // Not leaf node
            deleteHelper(node.children[index], key, level+1)
        } else {
                node.children[index].isLeaf = false
        }

        node.children[index].value--
        if node.children[index].value == 0 {
                node.children[index] = nil
        }
}

func (t *trie) delete(key string) {
        t.count--
        if t.search(key) {
                level := 0
                deleteHelper(t.root, key, level)
        }
}

func main() {
        // Construct trie
        t := newTrie()

        t.insert("the")
        t.insert("a")
        t.insert("there")
        t.insert("answer")
        t.insert("any")
        t.insert("by")
        t.insert("bye")
        t.insert("their")
        t.insert("thesis")

        t.delete("there")
        t.delete("the")

        fmt.Printf("%s --- %v\n", "the", t.search("the") )
        fmt.Printf("%s --- %v\n", "these", t.search("these") )
        fmt.Printf("%s --- %v\n", "there", t.search("there") )
        fmt.Printf("%s --- %v\n", "thesis", t.search("thesis") )
        fmt.Printf("%s --- %v\n", "thea", t.search("thea") )
        fmt.Printf("%s --- %v\n", "by", t.search("by") )
}
