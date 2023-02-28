# MerkleTreeRoot
Computing the merkle tree root corresponding to a set of transactions as fast and efficiently as possible.

#This implementation uses a recursive approach to build the merkle tree and splits the work of building the left and right subtrees into separate goroutines, which run concurrently. This can significantly speed up the calculation of the merkle tree root, especially for large numbers of transactions.

:point_right:To use this implementation, save it in a file called merkle.

:point_right:go and create a file called transactions.txt in the same directory with one hex-encoded transaction hash per line. Then, run the following command:
:point_right: go run MerkleTreeRoot.go :point_left:

![Alt text](/animeshbhatt651/MerkleTreeRoot/output_merkleTreeRoot.jpg
 "Output as merkle root hash in hexadecimal format")
