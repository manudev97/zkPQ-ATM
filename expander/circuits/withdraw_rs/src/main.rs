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
    let file = File::open(filename).expect("No se pudo abrir el archivo de entrada");
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).expect("Error al leer JSON")
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
         builder.assert_is_equal(self.path_indices[0], 1);
         builder.assert_is_equal(self.path_indices[1], 0);
    }
}
