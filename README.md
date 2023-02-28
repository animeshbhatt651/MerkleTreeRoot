# MerkleTreeRoot
Computing the merkle tree root corresponding to a set of transactions as fast and efficiently as possible.

#This implementation uses a recursive approach to build the merkle tree and splits the work of building the left and right subtrees into separate goroutines, which run concurrently. This can significantly speed up the calculation of the merkle tree root, especially for large numbers of transactions.

:point_right:To use this implementation, save it in a file called merkle.

:point_right:go and create a file called transactions.txt in the same directory with one hex-encoded transaction hash per line. Then, run the following command:
:point_right: go run MerkleTreeRoot.go :point_left:

This will show the output as merkle root hash in hexadecimal format.
![output_merkleTreeRoot](https://user-images.githubusercontent.com/75935128/221887210-706d2d9e-b48f-47cf-b149-1b507cc27198.jpg)
following hexadecimal code : 6533da470ce32a60326226601cd6ad8b12739fd5962453c1902cc3075b5b8938

Merkle tree also known as hash tree is a data structure used for data verification and synchronization. 
It is a tree data structure where each non-leaf node is a hash of it’s child nodes. All the leaf nodes are at the same depth and are as far left as possible. 
It maintains data integrity and uses hash functions for this purpose. 

Hash Functions: 
So before understanding how Merkle trees work, we need to understand how hash functions work. 
A hash function maps an input to a fixed output and this output is called hash. The output is unique for every input and this enables fingerprinting of data. So, huge amounts of data can be easily identified through their hash. 

![Merkle-Tree-1](https://user-images.githubusercontent.com/75935128/221894301-34dbddb5-7c05-4649-aae1-2f8e7334ed65.png)
