// SPDX-License-Identifier: MIT

/* zkPQ-ATM: Private Transactions using Expander */

pragma solidity >=0.8.0 <0.9.0;
import "./MerkleTree.sol";
import "./ZKPQATM_Token.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";  // Import token erc20 interface

struct Proof {
    uint _proof
}

interface IVerifier {
  function verifyExpander(
		uint _proof,
		uint256[2] memory _input
	) external returns (bool);
}

contract zkPQATM is MerkleTree {
	uint256 public immutable denomination;
	IVerifier public immutable verifier;
	uint256 public balanceTotalATM;
	ZKPQATM_Token public zkatmToken;  // Token zkatm reference
	mapping (address _account => uint256 _balance) public balance;
	mapping (bytes32 => bool) public nullifierHashes;
	mapping (address _account => bool) public initialized;
	// Store commitments to avoid accidental deposits with the same commitment.
	mapping(bytes32 => bool) public commitments; 						 
	string public greeting = "Starting to build my ZKPQATM proyect!!!";
	// _verifier -> The direction of the verifier contract for this snark
 	// _hashher -> The Hash Poseidon Contract Directorate
 	// _denomination -> amount to be transferred for each deposit
 	// _merkleTreeHeight -> The height of the Merkle Tree of deposits
 	// Hasher Address (Poseidon 8 Args): -
 	// VERIFIER ADDRESS: -
	
	constructor(
    	IVerifier _verifier,
		address _hasher,
		uint256 _denomination,
    	uint32 _merkleTreeHeight,
		ZKPQATM_Token _zkatmToken  // Add this parameter for Token Zkatm
  	)MerkleTree(_merkleTreeHeight, _hasher){
    	verifier = _verifier;
		hasher = Hasher(_hasher);
		denomination = _denomination;
		zkatmToken = _zkatmToken;  // Initializes the reference of Token Zkatm
  	}

	function initializeBalance() external {
        require(!initialized[msg.sender], "The balance has already been initialized");
		// Transfer 50 ZKATM tokens to the user who calls the function
        uint256 tokenAmount = 50 * (10 ** zkatmToken.decimals());  // Adjust according to the decimals of the token
        zkatmToken.mint(msg.sender, tokenAmount);
		initialized[msg.sender] = true;
		balance[msg.sender] = zkatmToken.balanceOf(msg.sender);
    }

    event Deposit (address from, uint32 leafIndex, uint256 value, uint256 timestamp);
    event Transfer (address from, address to, uint256 value, uint256 timestamp);
    event Withdraw (address to, bytes32 _nullifierHash, uint256 value);
	
    error InsufficientBalance();

	function generateCommitment() public view returns (
		bytes32 nullifier, 
		bytes32 secret,
	 	bytes32 commitment, 
		bytes32 nullifierHash
		) {
		// You can adjust how you generate the nullifier
        nullifier = bytes32(uint256(keccak256(abi.encodePacked(block.timestamp, msg.sender))) % MerkleTree.FIELD_SIZE); 
		// You can adjust how you generate the secret
        secret = bytes32(uint256(keccak256(abi.encodePacked(nullifier, block.timestamp, msg.sender))) % MerkleTree.FIELD_SIZE);  
		commitment = hasher.poseidon([nullifier, secret]);
        nullifierHash = hasher.poseidon([nullifier, nullifier]);
    }

	// Commitment The Note Commitment, which is posidonhash (Nullifier + Secret)
    function deposit(bytes32 commitment, uint256 _amount) public  {
		require(!commitments[commitment], "The commitment has already been inserted");
		if(zkatmToken.balanceOf(msg.sender) < _amount) revert InsufficientBalance();
		require(_amount == (denomination * (10 ** zkatmToken.decimals())), "The amount deposited is not admitted");
		
		// Transfer the tokens from the user who calls the function to this contract
    	bool transferSuccess = zkatmToken.transferFrom(msg.sender, address(this), _amount);
    	require(transferSuccess, "Tokens transfer has failed");

		balance[msg.sender] = zkatmToken.balanceOf(msg.sender);
		uint32 insertedIndex = _insert(commitment);
		commitments[commitment] = true;
		balanceTotalATM += _amount;
		emit Deposit (msg.sender, insertedIndex, balance[msg.sender], block.timestamp);
	}

	// _proof -> proof Expander - ECC.
 	// _input -> an array of the circuit with the following public entries:
 	// - Merkle root of all deposits in the contract
 	// - Nullifer Hash of unique deposit to avoid double expenses
    function withdraw(uint256 _amount, Proof calldata _proof, bytes32 _root, bytes32 _nullifierHash) public {
		require(_amount == denomination, "The amount to be removed is not admitted");
		require(!nullifierHashes[_nullifierHash], "The note was already spent");
		require(isKnownRoot(_root), "I can't find Merkle's root"); // Be sure to use one of the recent last 3
		require(verifier.verifyExpander(_proof, [uint256(_root), uint256(_nullifierHash)]), "Proof to withdraw invalidates");
		nullifierHashes[_nullifierHash] = true;

		// Transfer the tokens from the contract to the user who calls the function
    	bool transferSuccess = zkatmToken.transfer(msg.sender, _amount);
    	require(transferSuccess, "Tokens transfer has failed");
		
		balance[msg.sender] = zkatmToken.balanceOf(msg.sender);
		balanceTotalATM -= _amount;
		emit Withdraw (msg.sender, _nullifierHash, _amount);
	}

    function transfer(address _receiver, uint256 _amount, Proof calldata _proof, bytes32 _root, bytes32 _nullifierHash) public {
		require(_amount == denomination, "The amount to be removed is not admitted");
		require(!nullifierHashes[_nullifierHash], "The note was already spent");
		require(isKnownRoot(_root), "I can't find Merkle's root"); // Be sure to use one of the recent last 3
		require(verifier.verifyExpander(_proof, [uint256(_root), uint256(_nullifierHash)]), "Proof to withdraw invalidates");
		nullifierHashes[_nullifierHash] = true;

		// Transfer the tokens from the contract to the user who calls the function
    	bool transferSuccess = zkatmToken.transfer(_receiver, _amount);
    	require(transferSuccess, "Tokens transfer has failed");
		balance[_receiver] = zkatmToken.balanceOf(_receiver);
		balanceTotalATM -= _amount;
		emit Transfer (msg.sender, _receiver, _amount, block.timestamp);
	}
	 function isSpent(bytes32 _nullifierHash) public view returns (bool) {
        return nullifierHashes[_nullifierHash];
    } 
}


// Next steps: 
// - Implement Expander proof generation and verification
// - Create frontend integration with Wagmi & Next.js
// - Develop testing and deployment scripts
// - Add Verifier Expander 