# Huffman compression tool in Go

## Huffman encoding/decoding

The Huffman compression algorithm is used for lossless data compression.

Premise: text has an unequal distribution of characters. Therefore, we can assign the shortest "code" to the most occurring character, and the opposite for the least occurring.

Constraint: prefix-free codes should be selected for each character. A prefix-free encoding means that the bit string of a character does not overlap or cause ambiguity with the prefix of another character.

Example
```
Encoding:
    a: 1
    b: 01
    c: 10

encoding 101 is ambiguous, having two possible interpretations:
1. ab (1 => a, b => 01)
2. ca (10 => c, a => 1)
```

## Prefix-free Binary Tree Huffman encoder

1. Read the text and determine the frequency of each character occurring.
2. Build the binary tree from the frequencies.
3. Generate the prefix-code table from the tree.
4. Encode the text using the code table.
5. Encode the tree - we’ll need to include this in the output file so we can decode it.
6. Write the encoded tree and text to an output field
