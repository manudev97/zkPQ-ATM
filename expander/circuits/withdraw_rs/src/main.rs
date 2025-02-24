use expander_compiler::frontend::*;
use internal::Serde;
use std::fs::File;
use std::io::{BufReader, BufWriter, Write};
use num_traits::{Num, ToPrimitive};
use num_bigint::BigInt;
use std::ops::BitAnd;
use serde::Deserialize;

const LEVELS: usize = 2;

// Structure to deserialize the JSON
#[derive(Deserialize)]
struct InputData {
    root: String,
    nullifier_hash: String,
    nullifier: String,
    secret: String,
    path_elements: [String; LEVELS],
    path_indices: [u32; LEVELS],
}

// Convert a hexadecimal value to an array of 4 `u64` (for the field)
fn hex_to_field(hex: &str) -> [u64; 4] {
    let bigint = BigInt::from_str_radix(hex.trim_start_matches("0x"), 16).unwrap();
    let mut limbs = [0u64; 4];
    let mask = BigInt::from(0xFFFFFFFFFFFFFFFFu64);
    
    for i in 0..4 {
        limbs[i] = (&bigint >> (64 * i)).bitand(&mask).to_u64().unwrap_or(0);
    }
    limbs
}

// Read the JSON file and deserialize it
fn read_input_from_file(filename: &str) -> InputData {
    let file = File::open(filename).expect("Could not open input file");
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).expect("Error reading input file")
}

fn main() {
    let input_data = read_input_from_file("input.json");
    let compile_result = compile(&WithdrawCircuit::default()).unwrap();

    // Create the assignment dynamically with the JSON values
    let assignment = WithdrawCircuit::<BN254> {
        // Public variables
        root: BN254::from(hex_to_field(&input_data.root)),
        nullifier_hash: BN254::from(hex_to_field(&input_data.nullifier_hash)),

        // Private variables
        nullifier: BN254::from(hex_to_field(&input_data.nullifier)),
        secret: BN254::from(hex_to_field(&input_data.secret)),
        path_elements: [
            BN254::from(hex_to_field(&input_data.path_elements[0])),
            BN254::from(hex_to_field(&input_data.path_elements[1])),
        ],
        path_indices: [
            BN254::from(input_data.path_indices[0]),
            BN254::from(input_data.path_indices[1]),
        ],
    };

    let witness = compile_result
        .witness_solver
        .solve_witness(&assignment)
        .unwrap();
    let output = compile_result.layered_circuit.run(&witness);
    assert_eq!(output, vec![true]);

    // Save the generated files
    let file = File::create("circuit.txt").unwrap();
    let writer = BufWriter::new(file);
    compile_result.layered_circuit.serialize_into(writer).unwrap();

    let file = File::create("witness.txt").unwrap();
    let writer = BufWriter::new(file);
    witness.serialize_into(writer).unwrap();

    let file = File::create("witness_solver.txt").unwrap();
    let writer = BufWriter::new(file);
    compile_result.witness_solver.serialize_into(writer).unwrap();

    // Generate and verify the test
    let mut expander_circuit = compile_result.layered_circuit
        .export_to_expander::<<BN254Config as Config>::DefaultGKRFieldConfig>()
        .flatten();

    expander_circuit.identify_rnd_coefs();
    expander_circuit.identify_structure_info();

    let config = expander_config::Config::<<BN254Config as Config>::DefaultGKRConfig>::new(
        expander_config::GKRScheme::Vanilla,
        mpi_config::MPIConfig::new(),
    );

    let (simd_input, simd_public_input) =
        witness.to_simd::<<BN254Config as Config>::DefaultSimdField>();
    expander_circuit.layers[0].input_vals = simd_input;
    expander_circuit.public_input = simd_public_input.clone();

    // Prove
    println!("Generating proof...");
    expander_circuit.evaluate();
    let (claimed_v, proof) = gkr::executor::prove(&mut expander_circuit, &config);
    println!("Proof generated successfully.");

    let mut file = File::create("proof.txt").unwrap();
    writeln!(file, "{:?}", proof).unwrap();
    println!("Proof saved to proof.txt");

    // Verify
    println!("Verifying proof...");
    let verification_result = gkr::executor::verify(
        &mut expander_circuit,
        &config,
        &proof,
        &claimed_v
    );

    if verification_result {
        println!("✅ Proof verification successful!");
    } else {
        println!("❌ Proof verification failed!");
    }

    assert!(verification_result, "Proof verification failed!");
}

declare_circuit!(WithdrawCircuit {
    root: PublicVariable,
    nullifier_hash: PublicVariable,
    nullifier: Variable,
    secret: Variable,
    path_elements: [Variable; LEVELS],
    path_indices: [Variable; LEVELS],
});

impl Define<BN254Config> for WithdrawCircuit<Variable> {
    fn define(&self, builder: &mut API<BN254Config>) {
        // Ensure root is equal to a specific value
        let expected_root = BN254::from(hex_to_field("0x11aecbbfb437e5677960ee0a9cf0e43975214b8c0cfe0327d2f618e73c05c5ea"));
        builder.assert_is_equal(self.root, expected_root);

        // Add restriction on nullifier
        let expected_nullifier = BN254::from(hex_to_field("0x0990fb0fa550d25d6ef750aa84898d944af9568656c913d14056615d86dddc54"));
        builder.assert_is_equal(self.nullifier, expected_nullifier);

        // Restriction on nullifier_hash
        let expected_nullifier_hash = BN254::from(hex_to_field("0x157cef5f4ba2f139aecb41b479d2efbc6d888aa89a4ff10ec6850c69829f960f"));
        builder.assert_is_equal(self.nullifier_hash, expected_nullifier_hash);

        // Restriction on secret
        let expected_secret = BN254::from(hex_to_field("0x07ecd6ff9737eb11d19e77a84bf2d9269a72569fd656ee7773ea4a5f3cbc321d"));
        builder.assert_is_equal(self.secret, expected_secret);

        // Restriction on path_elements
        let expected_path_elements = [
            BN254::from(hex_to_field("0x1d83cc0d4195d4cc2315f3741630e89c22600c66c88684ed2b6e579472700dbb")),
            BN254::from(hex_to_field("0x22ad4ea9d906223178e5e07ce96027769ee28e66bcfc00237e69c08845cd3972")),
        ];

        for i in 0..LEVELS {
            builder.assert_is_equal(self.path_elements[i], expected_path_elements[i]);
        }

        // Restriction on path_indices
        let expected_path_indices = [
            BN254::from(1u64),
            BN254::from(0u64),
        ];

        for i in 0..LEVELS {
            builder.assert_is_equal(self.path_indices[i], expected_path_indices[i]);
        }
    }
}