use expander_compiler::frontend::*;
use internal::Serde;
use num_bigint::BigInt;
use num_traits::{Num, ToPrimitive, FromPrimitive};
use serde::Deserialize;
use std::fs::File;
use std::io::{BufReader, BufWriter, Write};
use std::ops::BitAnd;
use circuit_std_rs::poseidon_m31::*;

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
fn _hex_to_field_bn254(hex: &str) -> [u64; 4] {
    let bigint = BigInt::from_str_radix(hex.trim_start_matches("0x"), 16).unwrap();
    let mut limbs = [0u64; 4];
    let mask = BigInt::from(0xFFFFFFFFFFFFFFFFu64);

    for i in 0..4 {
        limbs[i] = (&bigint >> (64 * i)).bitand(&mask).to_u64().unwrap_or(0);
    }
    limbs
}

fn hex_to_field_m31(hex: &str) -> u32 {
    let bigint = BigInt::from_str_radix(hex.trim_start_matches("0x"), 16)
        .unwrap_or(BigInt::from(0));
    
    // Reducimos el valor a 31 bits (M31)
    let field_value = (bigint % BigInt::from(1 << 31))
        .to_u32()
        .unwrap_or(0);

    field_value
}

fn m31_outputs_to_hex(outputs: &[M31]) -> String {
    let mut bigint = BigInt::from(0);

    // Construimos el número grande concatenando los valores de M31 en bloques de 31 bits
    for (i, output) in outputs.iter().enumerate() {
        let part = BigInt::from_u32(output.v).unwrap_or(BigInt::from(0));
        bigint += &part << (31 * i);
    }

    format!("0x{:X}", bigint) // Convertimos a hexadecimal
}

fn hex_to_m31_outputs(hex: &str, num_outputs: usize) -> Vec<M31> {
    let bigint = BigInt::from_str_radix(hex.trim_start_matches("0x"), 16)
        .unwrap_or(BigInt::from(0));

    let mut outputs = Vec::new();
    let modulus = BigInt::from(1 << 31);

    for i in 0..num_outputs {
        let value = (&bigint >> (31 * i)) % &modulus;
        outputs.push(M31 {
            v: value.to_u32().unwrap_or(0),
        });
    }

    outputs
}

// Read the JSON file and deserialize it
fn read_input_from_file(filename: &str) -> InputData {
    let file = File::open(filename).expect("Could not open input file");
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).expect("Error reading input file")
}

#[test]
fn test_hex_to_m31_outputs() {
    let hex_hash = "0xABC96F9A67F491D4596C6DA1F7B4266A25F4F215ABA36E453BBBB3329B3D2A292F4D768BDE4BFE8DEB03593B55B98FA5F79657BF8BAB280634BABCDCD3E4";
    let outputs = hex_to_m31_outputs(hex_hash, 16);

    // Imprimir el resultado
    for output in &outputs {
        println!("M31 {{ v: {} }},", output.v);
    }
}

#[test]
fn test_m31_outputs_to_hex() {
    let outputs = vec![
        M31 { v: 1021105124 },
        M31 { v: 1342990709 },
        M31 { v: 1593716396 },
        M31 { v: 2100280498 },
        M31 { v: 330652568 },
        M31 { v: 1371365483 },
        M31 { v: 586650367 },
        M31 { v: 345482939 },
        M31 { v: 849034538 },
        M31 { v: 175601510 },
        M31 { v: 1454280121 },
        M31 { v: 1362077584 },
        M31 { v: 528171622 },
        M31 { v: 187534772 },
        M31 { v: 436020341 },
        M31 { v: 1441052621 },
    ];

    let hex_hash = m31_outputs_to_hex(&outputs);
    println!("Hexadecimal Hash: {}", hex_hash);
}



fn main() {
    let input_data = read_input_from_file("input.json");
    let compile_result = compile(&WithdrawCircuit::default()).unwrap();

    // Create the assignment dynamically with the JSON values
    let assignment = WithdrawCircuit::<M31> {
        // Public variables
        root: M31::from(hex_to_field_m31(&input_data.root)),
        nullifier_hash: [
            M31 { v: 1021105124 },
            M31 { v: 1342990709 },
            M31 { v: 1593716396 },
            M31 { v: 2100280498 },
            M31 { v: 330652568 },
            M31 { v: 1371365483 },
            M31 { v: 586650367 },
            M31 { v: 345482939 },
            M31 { v: 849034538 },
            M31 { v: 175601510 },
            M31 { v: 1454280121 },
            M31 { v: 1362077584 },
            M31 { v: 528171622 },
            M31 { v: 187534772 },
            M31 { v: 436020341 },
            M31 { v: 1441052621 },
        ],
        // Private variables
        nullifier: [M31::from(114514); 8],
        secret: M31::from(hex_to_field_m31(&input_data.secret)),
        path_elements: [
            M31::from(hex_to_field_m31(&input_data.path_elements[0])),
            M31::from(hex_to_field_m31(&input_data.path_elements[1])),
        ],
        path_indices: [
            M31::from(input_data.path_indices[0]),
            M31::from(input_data.path_indices[1]),
        ]
    };

    println!("{:?}",M31::from(hex_to_field_m31(&input_data.root)));
    let witness = compile_result
        .witness_solver
        .solve_witness(&assignment)
        .unwrap();
    let output = compile_result.layered_circuit.run(&witness);
    assert_eq!(output, vec![true]);

    // Save the generated files
    let file = File::create("circuit.txt").unwrap();
    let writer = BufWriter::new(file);
    compile_result
        .layered_circuit
        .serialize_into(writer)
        .unwrap();

    let file = File::create("witness.txt").unwrap();
    let writer = BufWriter::new(file);
    witness.serialize_into(writer).unwrap();

    let file = File::create("witness_solver.txt").unwrap();
    let writer = BufWriter::new(file);
    compile_result
        .witness_solver
        .serialize_into(writer)
        .unwrap();

    // Generate and verify the test
    let mut expander_circuit = compile_result
        .layered_circuit
        .export_to_expander::<<M31Config as Config>::DefaultGKRFieldConfig>()
        .flatten();

    expander_circuit.identify_rnd_coefs();
    expander_circuit.identify_structure_info();

    let config = expander_config::Config::<<M31Config as Config>::DefaultGKRConfig>::new(
        expander_config::GKRScheme::Vanilla,
        mpi_config::MPIConfig::new(),
    );

    let (simd_input, simd_public_input) =
        witness.to_simd::<<M31Config as Config>::DefaultSimdField>();
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
    let verification_result =
        gkr::executor::verify(&mut expander_circuit, &config, &proof, &claimed_v);

    if verification_result {
        println!("✅ Proof verification successful!");
    } else {
        println!("❌ Proof verification failed!");
    }

    assert!(verification_result, "Proof verification failed!");
}

declare_circuit!(WithdrawCircuit {
    root: PublicVariable,
    nullifier_hash: [PublicVariable; 16],
    nullifier: [Variable; 8],
    secret: Variable,
    path_elements: [Variable; LEVELS],
    path_indices: [Variable; LEVELS]
});

impl Define<M31Config> for WithdrawCircuit<Variable> {
    fn define(&self, builder: &mut API<M31Config>) {
        // Ensure root is equal to a specific value
        let expected_root = M31::from(hex_to_field_m31(
            "0x11aecbbfb437e5677960ee0a9cf0e43975214b8c0cfe0327d2f618e73c05c5ea",
        ));
        builder.assert_is_equal(self.root, expected_root);

        // Restriction on secret
        let expected_secret = M31::from(hex_to_field_m31(
            "0x07ecd6ff9737eb11d19e77a84bf2d9269a72569fd656ee7773ea4a5f3cbc321d",
        ));
        builder.assert_is_equal(self.secret, expected_secret);

        // Restriction on path_elements
        let expected_path_elements = [
            M31::from(hex_to_field_m31(
                "0x1d83cc0d4195d4cc2315f3741630e89c22600c66c88684ed2b6e579472700dbb",
            )),
            M31::from(hex_to_field_m31(
                "0x22ad4ea9d906223178e5e07ce96027769ee28e66bcfc00237e69c08845cd3972",
            )),
        ];

        for i in 0..LEVELS {
            builder.assert_is_equal(self.path_elements[i], expected_path_elements[i]);
        }

        // Restriction on path_indices
        let expected_path_indices = [M31::from(1), M31::from(0)];

        for i in 0..LEVELS {
            builder.assert_is_equal(self.path_indices[i], expected_path_indices[i]);
        }

        let params = PoseidonM31Params::new(
            builder,
            POSEIDON_M31X16_RATE,
            16,
            POSEIDON_M31X16_FULL_ROUNDS,
            POSEIDON_M31X16_PARTIAL_ROUNDS,
        );
        let res: Vec<Variable> = params.hash_to_state(builder, &self.nullifier);
        dbg!(builder.display("VARIABLE", res[1]));
        (0..params.width).for_each(|i| {
            builder.assert_is_equal(&res[i], self.nullifier_hash[i]);
            dbg!(builder.constant_value(res[1]));
        });
    }
}
