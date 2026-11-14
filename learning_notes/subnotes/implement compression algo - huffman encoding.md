# Huffman Encoding — C++ Implementation Roadmap

Build a **lossless, byte-oriented compressor and decompressor** in stages. Implement and test each piece before connecting the whole pipeline.

Start with an in-memory example, then add packed bits and a small custom file format. Huffman coding is the algorithm; the file format is a separate design choice.

## Read before starting

Read the first two references before coding. Keep the others handy for the relevant phases—you don't need to master everything up front.

1. **[Stanford CS106B — Huffman Coding](https://web.stanford.edu/class/archive/cs/cs106b/cs106b.1258/lectures/23-huffman/)**  
   Start here for worked examples of prefix-free codes, encoding trees, decoding, and the greedy tree-building algorithm. Focus on “Variable-Length Encoding” through “Flattening the Encoding Tree.” After reading, build the `BANANA` tree on paper.
2. **[cppreference — `std::priority_queue`](https://en.cppreference.com/w/cpp/container/priority_queue.html)**  
   Read the introduction, `Compare` parameter, and min-heap examples. Understand why the default queue returns the largest element and how to reverse that ordering. Useful for phases 4–5.
3. **[cppreference — `std::unique_ptr`](https://en.cppreference.com/w/cpp/memory/unique_ptr.html)**  
   Review ownership, automatic cleanup, and `.get()`. Decide who owns the tree nodes before allocating them. Useful from phase 3 onward.
4. **[LearnCpp — Bit manipulation with bitwise operators and bit masks](https://www.learncpp.com/cpp-tutorial/bit-manipulation-with-bitwise-operators-and-bit-masks/)**  
   Practice setting and testing bits with masks and shifts. This is preparation for packing codes into bytes in phase 9, not another compression algorithm to implement.
5. **[cppreference — `std::ifstream`](https://en.cppreference.com/w/cpp/io/basic_ifstream.html)**  
   Review binary mode, unformatted input (`get`/`read`), and stream error states. Avoid formatted extraction (`>>`) for arbitrary file bytes. Useful for phases 9–10.

**Ready-to-start checklist:** You can explain a prefix-free code, trace a recursive tree traversal, retrieve the smallest item from a heap, and distinguish a bit from a byte.

## The complete pipeline

```text
Input bytes
    → Count byte frequencies
    → Build a Huffman tree
    → Generate a byte-to-code lookup table
    → Encode and pack bits into bytes
    → Write header + compressed payload

Compressed file
    → Read header and rebuild the same tree
    → Decode payload bits
    → Recover the original bytes
```

## Phase 1 — Read the input

**Goal:** Read and store arbitrary bytes.

Start with a small example:

```cpp
std::string input = "BANANA";
```

Eventually, support all byte values from `0` to `255`, including null bytes. This project assumes 8-bit bytes. A `std::string` can hold embedded nulls, but don't treat its contents as a null-terminated C string. A `std::vector<unsigned char>` is another option.

**Checkpoint:** Verify the input length and inspect its byte values. File I/O can wait until the in-memory round trip works.

## Phase 2 — Count byte frequencies

**Goal:** Count how often each byte appears.

For `BANANA`:

| Byte | Frequency |
| --- | ---: |
| `A` | 3 |
| `B` | 1 |
| `N` | 2 |

Use a fixed-size table and wide counters:

```cpp
// Requires <array> and <cstdint>.
std::array<std::uint64_t, 256> frequencies{};

for (unsigned char c : input) {
    ++frequencies[c];
}
```

Index with `unsigned char`, not potentially negative plain `char` values.

**Checkpoint:** Print nonzero entries. Their sum must equal the input size.

## Phase 3 — Represent the Huffman tree

**Goal:** Define leaf nodes and internal nodes, with clear memory ownership.

```cpp
// Requires <cstddef> and <cstdint>.
struct Node {
    unsigned char character = 0; // Meaningful only for leaves.
    std::uint64_t frequency = 0;
    std::size_t order = 0;       // Unique creation ID for tie-breaking.
    Node* left = nullptr;
    Node* right = nullptr;
};
```

- A leaf has no children and represents one byte value.
- An internal node has two children and stores their combined frequency.
- Don't use `character == 0` to identify internal nodes: zero is a valid input byte.

One ownership strategy is a `std::vector<std::unique_ptr<Node>>` that owns every allocated node. The children and priority queue can hold non-owning `Node*` pointers. Keep the owning container alive while using the tree; don't delete those pointers separately.

**Checkpoint:** Create two leaves and a parent. Verify leaf detection and automatic cleanup.

## Phase 4 — Build a min-priority queue

**Goal:** Always retrieve the two nodes with the smallest frequencies.

```cpp
// Requires <queue>, <vector>, and your Compare definition.
std::priority_queue<Node*, std::vector<Node*>, Compare> pq;
```

Implement `Compare` so smaller frequencies have priority. With `std::priority_queue`, this means comparing frequencies using `>`, not `<`.

**Define deterministic ties now:**

1. Create leaves in ascending byte-value order, skipping zero-frequency bytes.
2. Give each new node a unique, increasing `order` value, including parents.
3. Compare by frequency first, then by `order`; smaller values win.
4. When merging, make the first removed node the left child.

If the file stores only frequencies, the decoder must follow these same rules. Equal frequencies do not uniquely determine a tree.

**Checkpoint:** For `BANANA`, the initial removal order is `B:1`, `N:2`, `A:3`. Also test equal-frequency nodes.

## Phase 5 — Construct the Huffman tree

**Goal:** Repeatedly merge the two least-frequent nodes.

While the queue contains more than one node:

1. Remove the two nodes with the lowest frequencies.
2. Create a parent whose frequency is their sum.
3. Attach the nodes as its left and right children.
4. Insert the parent into the queue.

The remaining node is the root. For `BANANA`, the rules above produce:

```text
B:1 + N:2 → parent:3
A:3 + parent:3 → root:6

        (6)
       /   \
     0/     \1
    A:3     (3)
           /   \
         0/     \1
        B:1     N:2
```

**Checkpoint:** The root frequency equals the input size, and each distinct input byte appears in exactly one leaf. Empty input has no root; one distinct byte produces a single leaf.

## Phase 6 — Generate binary codes

**Goal:** Map each byte to a prefix-free code: no code is a prefix of another.

Traverse from the root:

- Going left appends `0`.
- Going right appends `1`.
- Reaching a leaf saves the current path in the lookup table.

Traversal sketch:

```cpp
void generate_codes(const Node* node, std::string code) {
    if (!node) return;

    if (!node->left && !node->right) {
        // TODO: Save character -> code in your lookup table.
        // For a single-leaf tree, assign "0" instead of an empty code.
        return;
    }

    generate_codes(node->left, code + "0");
    generate_codes(node->right, code + "1");
}
```

For example, use `std::array<std::string, 256>` as the lookup table.

**Checkpoint:** Every input byte has a nonempty code. The tree above gives `A = 0`, `B = 10`, and `N = 11`. Other tree conventions can produce different, equally optimal codes.

## Phase 7 — Encode the input

**Goal:** Replace every byte with its code.

For each input byte, look up its code and append it to a temporary string of `'0'` and `'1'` characters.

Using the codes from phase 6:

```text
Input:   B  A  N  A  N  A
Codes:  10  0 11  0 11  0
Encoded: 100110110
```

This is a debugging representation, **not compressed binary storage yet**. Nine `'0'`/`'1'` characters occupy nine bytes on the assumed platform; nine packed bits need two bytes.

**Checkpoint:** The encoded bit count equals the sum of `frequency[byte] × code_length[byte]`. For `BANANA`, that is `3×1 + 1×2 + 2×2 = 9` bits.

## Phase 8 — Decode in memory

**Goal:** Reconstruct the original bytes before adding a file format.

For a tree with at least two leaves:

1. Start at the root.
2. Read `0` to move left or `1` to move right.
3. At a leaf, output its byte and return to the root.
4. Continue until the meaningful encoded bits are consumed.

Reject invalid bit characters and input that ends halfway through a code. Handle a single-leaf tree separately: with the convention above, each `0` represents one copy of its byte.

```text
Original: BANANA
Encoded:  100110110
Decoded:  BANANA
```

**Checkpoint:** `decode(encode(input)) == input`, including empty input and repeated copies of one byte. Don't wait until optimization to test these cases.

## Phase 9 — Pack bits and define the file format

**Goal:** Write a self-contained binary file that a fresh decoder can read.

First, implement a bit writer and bit reader. Pick a bit order—for example, fill each byte from its most-significant bit to its least-significant bit. Zero-pad the unused low bits of the final byte.

```text
Meaningful bits: 100110110
Stored bytes:   10011011 00000000
                          ^ seven padding bits after the first bit
```

Bit packing is necessary for actual compression; it isn't merely an optional optimization.

For a first file format, store a fixed frequency table:

| Field | Purpose |
| --- | --- |
| Magic bytes | Identify your format, e.g. `HUF1` |
| Version | Select the decoding rules |
| Original byte count | Know how many bytes to reconstruct |
| 256 frequency counters | Rebuild the exact tree using the specified tie rules |
| Packed payload | Store the encoded bits |

Document field widths, integer byte order, tree-building rules, bit order, and padding. Serialize fields explicitly—don't dump a C++ struct or pointers to disk.

Use the original byte count to stop decoding, rather than interpreting padding as more symbols. For this frequency-table format, you can also derive the meaningful bit count from the frequencies and code lengths. Validate the frequency sum, payload length, and padding. Empty input has all-zero frequencies and no payload; single-byte inputs follow the one-bit code convention.

**Checkpoint:** Save a file, discard the encoder's in-memory tree, and decode using only the file contents.

## Phase 10 — Test files, then optimize

Use binary-mode streams and check open, read, and write failures. Compare decompressed files byte-for-byte with their originals.

### Required tests

- [ ] Empty input.
- [ ] One byte and many repetitions of the same byte.
- [ ] `BANANA` and other small known examples.
- [ ] Equal-frequency symbols, to verify deterministic ties.
- [ ] All 256 byte values, including `0x00` and `0xFF`.
- [ ] Whitespace, newlines, and arbitrary binary files.
- [ ] Payloads ending exactly on a byte boundary and with partial final bytes.
- [ ] Random round-trip tests over many lengths.
- [ ] Truncated files, invalid headers, inconsistent sizes, and overflowing frequency sums.

### Improvements after correctness

- Measure **total file size**, including the header and padding—not just payload bits. A full table of 256 64-bit counters alone costs 2,048 bytes, so small files can grow substantially.
- Consider storing only nonzero frequencies, a serialized tree, or canonical code lengths to reduce header overhead.
- Replace the temporary bit string with direct writes to a bit writer.
- Add large-file processing: static Huffman coding generally needs two passes, one to count frequencies and one to encode. Non-seekable input needs buffering or a block-based design.
- Consider a checksum for corruption detection; structural validation alone cannot detect every altered payload.

**First milestone:** An in-memory round trip for `BANANA`, empty input, and a repeated single byte. Only then move on to the binary file format.

## NEXT step :  [[implementing DEFLATE (what git uses)]]