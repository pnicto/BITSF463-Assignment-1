use bitsf463_assignment_1::*;

#[test]
fn test_baconian_cipher() {
    let decoded = baconian_cipher(
        'C',
        'G',
        "CCCGCGCCCGGGCCCCGGGGGCCGGCGGGCCCGGCGCCCGCCCCCCGGGGCCGGGGGCCC",
    );
    assert_eq!(decoded, "CRYPTOGRAPHY");
}
