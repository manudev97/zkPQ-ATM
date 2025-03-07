use circuit_std_rs::poseidon_m31::*;
use expander_compiler::frontend::*;
// use circuit_std_rs::utils::*;
use expander_compiler::declare_circuit;
use internal::Serde;

use num_bigint::BigInt;
use num_traits::{FromPrimitive, Num, ToPrimitive};
use serde::Deserialize;
use std::fs::File;
use std::io::{BufReader, BufWriter, Write};
use std::ops::BitAnd;

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
    let bigint =
        BigInt::from_str_radix(hex.trim_start_matches("0x"), 16).unwrap_or(BigInt::from(0));

    // We reduce the value to 31 bits (M31)
    let field_value = (bigint % BigInt::from(1 << 31)).to_u32().unwrap_or(0);

    field_value
}

fn _m31_outputs_to_hex(outputs: &[M31]) -> String {
    let mut bigint = BigInt::from(0);

    // We build the large number by concatenating the values ​​of M31 into blocks of 31 bits
    for (i, output) in outputs.iter().enumerate() {
        let part = BigInt::from_u32(output.v).unwrap_or(BigInt::from(0));
        bigint += &part << (31 * i);
    }

    format!("0x{:x}", bigint)
}

fn hex_to_m31_outputs(hex: &str, num_outputs: usize) -> [M31; 16] {
    let bigint =
        BigInt::from_str_radix(hex.trim_start_matches("0x"), 16).unwrap_or(BigInt::from(0));

    let modulus = BigInt::from(1 << 31);
    let mut outputs = [M31 { v: 0 }; 16];

    for i in 0..num_outputs.min(16) {
        let value = (&bigint >> (31 * i)) % &modulus;
        outputs[i] = M31 {
            v: value.to_u32().unwrap_or(0),
        };
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
    let hex_hash = "1bf52";
    let outputs = hex_to_m31_outputs(hex_hash, 16);

    // Imprimir el resultado
    for output in &outputs {
        println!("M31 {{ v: {} }},", output.v);
    }
}

// #[test]
// fn test_m31_outputs_to_hex() {
//     let outputs = vec![
//             M31 { v: 1021105124 },
//             M31 { v: 1342990709 },
//             M31 { v: 1593716396 },
//             M31 { v: 2100280498 },
//             M31 { v: 330652568 },
//             M31 { v: 1371365483 },
//             M31 { v: 586650367 },
//             M31 { v: 345482939 },
//             M31 { v: 849034538 },
//             M31 { v: 175601510 },
//             M31 { v: 1454280121 },
//             M31 { v: 1362077584 },
//             M31 { v: 528171622 },
//             M31 { v: 187534772 },
//             M31 { v: 436020341 },
//             M31 { v: 1441052621 },
//         ];

//     let hex_hash = m31_outputs_to_hex(&outputs);
//     println!("Hexadecimal Hash: {}", hex_hash);
// }

fn main() {
    use extra::debug_eval;
    let input_data = read_input_from_file("input-m31.json");
    let compile_result = compile(&WithdrawCircuit::default(), CompileOptions::default()).unwrap();

    let mut hint_registry = HintRegistry::<M31>::new();

    // Create the assignment dynamically with the JSON values
    let assignment = WithdrawCircuit::<M31> {
        // Public variables
        root: M31::from(hex_to_field_m31(&input_data.root)),
        nullifier_hash: hex_to_m31_outputs(&input_data.nullifier_hash, 16),
        // Private variables
        nullifier: M31::from(hex_to_field_m31(&input_data.nullifier)),
        secret: M31::from(hex_to_field_m31(&input_data.secret)),
        path_elements: [
            M31::from(hex_to_field_m31(&input_data.path_elements[0])),
            M31::from(hex_to_field_m31(&input_data.path_elements[1])),
        ],
        path_indices: [
            M31::from(input_data.path_indices[0]),
            M31::from(input_data.path_indices[1]),
        ],
    };

    //println!("{:?}",M31::from(hex_to_field_m31(&input_data.root)));

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

    debug_eval(&WithdrawCircuit::default(), &assignment, hint_registry);

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
    nullifier: Variable,
    secret: Variable,
    path_elements: [Variable; LEVELS],
    path_indices: [Variable; LEVELS]
});

fn dual_mux<C: Config, B: RootAPI<C>>(
    api: &mut B,
    in0: Variable,
    in1: Variable,
    s: Variable,
) -> (Variable, Variable) {
    api.assert_is_bool(s);

    let in1_minus_in0 = api.sub(in1, in0);
    let in0_minus_in1 = api.sub(in0, in1);
    let mul1 = api.mul(in1_minus_in0, s);
    let mul2 = api.mul(in0_minus_in1, s);

    // Calculate out[0] = (in1 - in0) * s + in0
    let out0 = api.add(mul1, in0);

    // Calculate out[1] = (in0 - in1) * s + in1
    let out1 = api.add(mul2, in1);

    (out0, out1)
}

fn merkle_tree_checker<C: Config, B: RootAPI<C>>(
    api: &mut B,
    leaf: Vec<Variable>,        // leaf es un Vec<Variable>
    root: Variable,             // root es un solo Variable
    path_elements: &[Variable], // path_elements es un slice de Variable
    path_indices: &[Variable],  // path_indices es un slice de Variable
) {
    let levels = path_elements.len();
    let mut current_hash = leaf; // current_hash es un Vec<Variable>

    for i in 0..levels {
        // Para cada nivel del árbol, aplicamos el dual_mux a cada elemento del Vec
        let mut new_current_hash = Vec::with_capacity(current_hash.len()); // Nuevo Vec para el hash actualizado

        for j in 0..current_hash.len() {
            // Aplicamos dual_mux a cada elemento del Vec
            let (left, right) = dual_mux(api, current_hash[j], path_elements[i], path_indices[i]);

            // Hasher Poseidon para combinar left y right
            let params = PoseidonM31Params::new(
                api,
                POSEIDON_M31X16_RATE,
                16,
                POSEIDON_M31X16_FULL_ROUNDS,
                POSEIDON_M31X16_PARTIAL_ROUNDS,
            );

            // Aplicamos el hash a left y right
            let hash_result = params.hash_to_state(api, &[left, right]);
            new_current_hash.push(hash_result[0]); // Asumimos que el hash_result es un array de 1 elemento
        }

        // Actualizamos current_hash para el siguiente nivel
        current_hash = new_current_hash;
    }

    // Verificar que la raíz calculada coincide con la raíz proporcionada
    // Aquí asumimos que root es un solo Variable que representa el hash final
    api.assert_is_equal(current_hash[0], root);
}

impl Define<M31Config> for WithdrawCircuit<Variable> {
    fn define<Builder: RootAPI<M31Config>>(&self, api: &mut Builder) {
        // Ensure root is equal to a specific value
        let expected_root = M31::from(hex_to_field_m31(
            "0x11aecbbfb437e5677960ee0a9cf0e43975214b8c0cfe0327d2f618e73c05c5ea",
        ));
        api.assert_is_equal(self.root, expected_root);

        dbg!(api.constant_value(self.root));
        dbg!(api.display("Root", self.root));

        println!("{:?}", api.constant_value(11212));
        println!("{:?}", self.root);
        let expected_nullifier = M31::from(hex_to_field_m31("0x1bf52"));
        api.assert_is_equal(self.nullifier, expected_nullifier);

        // Restriction on secret
        let expected_secret = M31::from(hex_to_field_m31("0x07ecd"));
        api.assert_is_equal(self.secret, expected_secret);

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
            api.assert_is_equal(self.path_elements[i], expected_path_elements[i]);
        }

        // Restriction on path_indices
        let expected_path_indices = [M31::from(1), M31::from(0)];

        for i in 0..LEVELS {
            api.assert_is_equal(self.path_indices[i], expected_path_indices[i]);
        }

        let params = PoseidonM31Params::new(
            api,
            POSEIDON_M31X16_RATE,
            16,
            POSEIDON_M31X16_FULL_ROUNDS,
            POSEIDON_M31X16_PARTIAL_ROUNDS,
        );

        // Restriction on nullifier_hash
        let res: Vec<Variable> = params.hash_to_state(api, &[self.nullifier; 8]);
        //dbg!(builder.display("VARIABLE", res[1]));

        (0..params.width).for_each(|i| {
            api.assert_is_equal(&res[i], self.nullifier_hash[i]);
            //dbg!(builder.constant_value(res[1]));
        });

        // Calculate the commitment hash (commitment hasher)
        let commitment_hash: Vec<Variable> =
            params.hash_to_state(api, &[self.nullifier, self.secret]);

        for (i, hash_var) in commitment_hash.iter().enumerate() {
            if let Some(value) = api.constant_value(*hash_var) {
                println!("commitment_hash[{}] = {:?}", i, value);
            } else {
                println!("commitment_hash[{}] is not a constant", i);
            }
        }

        dbg!(commitment_hash);

        // for (i, hash_var) in commitment_hash.iter().enumerate() {
        //     builder.display(&format!("commitment_hash[{}]", i), *hash_var);
        // }

        // Llamar al subcircuito MerkleTreeChecker
        // let merkle_root = builder.memorized_simple_call(
        //     move |api: &mut API<M31Config>, inputs: &Vec<Variable>| {
        //         // Extraer los valores de inputs
        //         let leaf = commitment_hash.clone(); // leaf es un Vec<Variable>
        //         let root = inputs[0]; // root es un solo Variable
        //         let path_elements = &inputs[1..1 + LEVELS]; // path_elements es un slice de Variable
        //         let path_indices = &inputs[1 + LEVELS..1 + 2 * LEVELS]; // path_indices es un slice de Variable

        //         // Llamar a merkle_tree_checker
        //         merkle_tree_checker(api, leaf, root, path_elements, path_indices);

        //         // Devolver la raíz como salida
        //         vec![root]
        //     },
        //     &[
        //         // Pasar root
        //         self.root,
        //         // Pasar path_elements
        //         self.path_elements[0],
        //         self.path_elements[1],
        //         // Pasar path_indices
        //         self.path_indices[0],
        //         self.path_indices[1],
        //     ],
        // );

        // // Verificar que la raíz calculada coincide con la raíz proporcionada
        // builder.assert_is_equal(merkle_root[0], self.root);
    }
}
