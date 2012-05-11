package main

import (
        "fmt"
        "encoding/binary"
        "os"
        "io/ioutil"
)

const R int = 256   // ASCII ahphabet
var (
        root *Node
        size int
        compressedSize int
)

type Node struct {
        ch byte
        freq int
        left, right *Node
}

func (n *Node) isLeaf() bool {
        return n.left == nil && n.right == nil
}

func (n *Node) CompareTo(that Node) int {
        return n.freq - that.freq
}

func binaryRead(filename string) ([]byte, error) {
        b := make([]byte, compressedSize)
        f, _ := os.Open(filename)
        err := binary.Read(f, binary.LittleEndian, &b)
        if err != nil {
                fmt.Println("binary.Read failed:", err)
                return nil, err
        }
        return b, nil
}

func binaryWrite(bs []byte, filename string) (int, error) {
        f, _ := os.Create(filename)
        defer f.Close()
        r := make([]byte, 0)
        count := 0
        var temp byte = 0
        for _, b := range bs {
                if b == '1' {
                        temp |= 0x01
                } else {
                        temp &= 0xFE
                }
                count += 1
                if count == 8 {
                        r = append(r, temp)
                        temp = 0
                        count = 0
                } else {
                        temp <<= 1
                }
        }
        if count < 8 {  // Fill 0 to the remain bit and write it as byte
                temp <<= uint(8 - count)
                r = append(r, temp)
        }
        err := binary.Write(f, binary.LittleEndian, r)
        if err != nil {
                fmt.Println("binary.Write failed:", err)
                return 0, err
        }
        return len(r), nil
}

func buildCode(st []string, root *Node, s string) {
        if root.isLeaf() {
                st[root.ch] = s
                return
        }
        buildCode(st, root.left, s + "0")
        buildCode(st, root.right, s + "1")
}

func buildTrie(freq []int) *Node {
        pq := NewMinPQ(len(freq))
        var i int
        for i = 0; i < R; i++ {
                if freq[i] > 0 {
                        pq.Insert(Node{byte(i) , freq[i], nil, nil})
                }
        }

        for pq.Size() > 1 {
                x := pq.Delmin()
                y := pq.Delmin()
                parent := Node{0, x.freq + y.freq, &x, &y}
                pq.Insert(parent)
        }
        n := pq.Delmin()
        return &n
}

// Wrap function for concurrent binary file reading
func readBinary(filename string, c chan bool) error {
        in, err := binaryRead(filename)
        if err != nil {
                return err
        }
        total := 0
        for i := 0; i < len(in); i++ {
                b := in[i]
                //fmt.Printf("b is 0x%.2x\n", b)
                count := 0
                for count < 8 {
                        var r bool
                        if b & 0x80 == 0x80 {
                                r = true
                        } else {
                                r = false
                        }
                        c <- r
                        total++
                        b <<= 1
                        count++
                }
        }
        println("Send", total)
        return nil
}

// Using huffman code to compress a file and save it to file.zz
func Encode(filename string) error {
        outFilename := filename + ".zz"
        out := make([]byte, 0)
        // Read file into byte array
        filecontent, err := ioutil.ReadFile(filename)
        if err != nil {
                return err
        }
        size = len(filecontent)
        fmt.Printf("Number of bytes I read: %d\n", size)
        // Tabulate frequency counts
        freq := make([]int, R)
        for _, b := range filecontent {
                freq[b]++
        }
        // Build Huffman code trie
        root = buildTrie(freq)
        // Build code table
        st := make([]string, R)
        buildCode(st, root, "")
        // Huffman code to encode input
        for i:=0; i<len(filecontent); i++ {
                code := st[filecontent[i]]
                out = append(out, []byte(code)...)
        }
        compressedSize, err = binaryWrite(out, outFilename)
        fmt.Printf("Number of bytes I write: %d\n", compressedSize)
        fmt.Printf("Compression ratio: %f\n", float32(compressedSize)/float32(size))
        return err
}

// Using huffman code to decompress a file.zz
// Huffman tree is shared in this example, in real app, Huffman tree need
// to be saved in the beginning of compressed file.
func Decode(filename string) ([]byte, error) {
        result := make([]byte, 0)
        c := make(chan bool)
        go readBinary(filename, c)
        fmt.Println("\nDecoded content:")
        total := 0
        for i:=0; i<size; i++ {
                x := root
                for !x.isLeaf() {
                        isRight := <-c
                        total++
                        if isRight {
                                x = x.right
                        } else {
                                x = x.left
                        }
                }
                result = append(result, x.ch)
        }
        fmt.Println(string(result))
        return result, nil
}

func main() {
        Encode("./test")
        Decode("./test.zz")
}
