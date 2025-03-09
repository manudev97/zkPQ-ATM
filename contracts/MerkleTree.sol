// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface Hasher {
  function poseidon(bytes32[2] calldata leftRight) external pure returns (bytes32);
}

contract MerkleTree {
    uint256 public constant FIELD_SIZE =
        21888242871839275222246405745257275088548364400416034343698204186575808495617; // Order (cardinality) of the curve
    uint256 public constant ZERO_VALUE =
        13187267684982673088682801604133345722716467932170606900401783776316647633885; // = keccak256("zkPQATM") % FIELD_SIZE

    Hasher public hasher;
    uint32 public immutable levels;

    // The following variables are made public to facilitate tests and
    // I know that they should not access them by normal code
    bytes32[] public filledSubtrees; // Arrangement with roots of the left sub -leaf (of size _levels)
    bytes32[] public zeros; // Arrangement with roots from the right to the right of a sheet (_levels size)
    uint32 public currentRootIndex = 0;
    uint32 public nextIndex = 0;
    uint32 public constant ROOT_HISTORY_SIZE = 3;  // It has to be less than 2 ** Levels (each new new root sheet)
    bytes32[ROOT_HISTORY_SIZE] public roots;

    constructor(uint32 _levels, address _hasher) {
        require(_levels > 0, "_levels debe ser mayor que cero");
        require(_levels < 6, "_levels debe ser menor que 6");
        levels = _levels;
        hasher = Hasher(_hasher);
    
        bytes32 currentZero = bytes32(ZERO_VALUE);
        zeros.push(currentZero);
        filledSubtrees.push(currentZero);

        // Obtaining the first root for all empty leaves
        for (uint32 i = 1; i < _levels; i++) {
            currentZero = hashLeftRight(currentZero, currentZero);
            zeros.push(currentZero); 
            filledSubtrees.push(currentZero);
        }

        roots[0] = hashLeftRight(currentZero, currentZero); 
  }

    // Having 2 leaves of the tree, returns Poseidon (Leftright)
    function hashLeftRight(bytes32 _left, bytes32 _right) public view returns (bytes32) {
        require(uint256(_left) < FIELD_SIZE, "_left debe estar dentro del campo" );
        require(uint256(_right) < FIELD_SIZE, "_right debe estar dentro del campo" );
        bytes32[2] memory leftright = [_left, _right];
        return hasher.poseidon(leftright);
    }

    function _insert(bytes32 _leaf) internal returns  (uint32 index) {
        uint32 currentIndex = nextIndex;
        require(currentIndex != uint32(2)**levels, "Arbol Merkle lleno. No se pueden agregar hojas.");
        nextIndex += 1;
        bytes32 currentLevelHash = _leaf;
        bytes32 left;
        bytes32 right;

        for (uint32 i = 0; i < levels; i++) {
            if (currentIndex % 2 == 0) {
                left = currentLevelHash;
                right = zeros[i];
                
                filledSubtrees[i] = currentLevelHash;
            } else {
                left = filledSubtrees[i];
                right = currentLevelHash;
            }

            currentLevelHash = hashLeftRight(left, right);

            currentIndex /= 2;
        }

        currentRootIndex = (currentRootIndex + 1) % ROOT_HISTORY_SIZE;
        roots[currentRootIndex] = currentLevelHash;
        return nextIndex - 1;
    }

    // If the root is present in the root history
    function isKnownRoot(bytes32 _root) public view returns (bool) {
        if (_root == 0) return false;

        uint32 i = currentRootIndex;
        do {
            if (_root == roots[i]) return true;
            if (i == 0) i = ROOT_HISTORY_SIZE;
            i--;
        } while (i != currentRootIndex);
        return false;
    }

    function getLastRoot() public view returns (bytes32) {
        return roots[currentRootIndex];
    }
}
