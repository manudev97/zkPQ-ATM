

## GKR from Scratch for zkPQ-ATM
The GKR (Goldwasser-Kalai-Rothblum) protocol is a powerful tool in the realm of zero-knowledge proofs, designed to efficiently verify computations in arithmetic circuits. Its innovative approach allows a verifier with limited resources to check the correctness of complex computations without needing to run the entire circuit, making it ideal for systems where scalability and efficiency are critical.

### How does the GKR protocol work?
Instead of evaluating each gate and each layer of the circuit independently, the GKR protocol leverages the algebraic structure of the problem to recursively verify computations. It uses the sumcheck protocol to validate partial sums representing the value computed at each layer of the circuit. This recursive approach dramatically reduces the computational burden, as the verifier only needs to interact with a small portion of the circuit at each step, rather than committing to all intermediate results.

### Purpose of this document
In this document, we will explore how the GKR protocol can be applied to a specific circuit within the Expander framework. Through a practical example, we will describe step-by-step how the protocol works, how it integrates with the circuit, and how it can be implemented to improve the efficiency of zero-knowledge proofs.

## GKR Protocol

We will approach the GKR protocol from a practical perspective, starting by describing the type of computation we wish to prove. Although the GKR protocol is inherently interactive, it is important to note that it can be made non-interactive using the Fiat-Shamir transform, a technique widely used in zero-knowledge proof systems to eliminate the need for interaction between the prover and the verifier.

We will work on the M31 field, a finite field defined as $F_{2^{31} - 1}$. This field is especially efficient for implementations on 64-bit architectures, since its size (31 bits) allows to take full advantage of native CPU operations, resulting in fast and optimized computations. One of the key advantages of M31 is its fast modular arithmetic, which simplifies operations such as additions, multiplications, and reductions, which is crucial to ensure efficiency in complex arithmetic circuits.

M31 is not only efficient, but also ideal for ensuring the correctness of in-circuit calculations, as its algebraic structure allows for verifications to be performed accurately and without loss of performance. This makes it an excellent choice for systems requiring high performance and security, as is the case with zkPQ-ATM.

In addition to M31, Expander supports two other main fields [View Fields in Expander](https://github.com/PolyhedraZK/ExpanderCompilerCollection/tree/4989db69573cbbed8f57549b31abf7857146286a/ecgo/field), each with its own advantages and use cases: 
- BN254: This field is a standard in Ethereum and is optimal for pairing-based schemes. With a size of ~254 bits, BN254 is widely used in cryptographic applications that require blockchain compatibility and advanced operations such as BLS signatures.
- GF2: This binary field is ideal for boolean circuits, where operations are performed in bit arithmetic. GF2 is especially useful in applications requiring boolean logic verifications or calculations in binary systems.

The choice of M31 for this project is due to its unique balance between efficiency and ease of implementation, making it perfect for arithmetic circuits in zero-knowledge proof systems like the GKR protocol.


Arithmetic circuit in $F_{2^{31}-1}$ for the GKR protocol