package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	txs := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	var leaves [][]byte
	for _, tx := range txs {
		leaves = append(leaves, hash(tx))
	}
	root := buildTreeConcurrent(leaves)
	fmt.Println(hex.EncodeToString(root))
}
func hash(data []byte) []byte {
	h := sha256.Sum256(data)
	return h[:]
}
func buildTree(leaves [][]byte) []byte {
	if len(leaves) == 1 {
		return leaves[0]
	}
	if len(leaves)%2 != 0 {
		leaves = append(leaves, leaves[len(leaves)-1])
	}
	var level [][]byte
	for i := 0; i < len(leaves); i += 2 {
		hash := hash(append(leaves[i], leaves[i+1]...))
		level = append(level, hash)
	}
	return buildTree(level)
}
func buildTreeConcurrent(leaves [][]byte) []byte {
	if len(leaves) == 1 {
		return leaves[0]
	}
	if len(leaves)%2 != 0 {
		leaves = append(leaves, leaves[len(leaves)-1])
	}
	half := len(leaves) / 2
	left := leaves[:half]
	right := leaves[half:]
	var wg sync.WaitGroup
	wg.Add(2)
	var leftHash, rightHash []byte
	go func() {
		leftHash = buildTreeConcurrent(left)
		wg.Done()
	}()
	go func() {
		rightHash = buildTreeConcurrent(right)
		wg.Done()
	}()
	wg.Wait()
	hash := hash(append(leftHash, rightHash...))

	return hash
}
