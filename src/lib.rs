use std::collections::HashMap;

pub fn baconian_cipher(sc0: char, sc1: char, secret_text: &str) -> String {
    let mut lookup_table: HashMap<String, char> = HashMap::new();

    // convert secret text to binary
    let mut binary_secret_text = String::new();
    for letter in secret_text.chars() {
        if letter == sc0 {
            binary_secret_text.push('0');
        } else if letter == sc1 {
            binary_secret_text.push('1');
        }
    }

    // create lookup
    for (index, letter) in ('A'..='Z').enumerate() {
        let binary_string = format!("{:05b}", index);
        lookup_table.insert(binary_string, letter);
    }

    // construct the string
    let mut decoded = String::new();
    for index in (0..binary_secret_text.len()).step_by(5) {
        let block = &binary_secret_text[index..index + 5];
        if let Some(letter) = lookup_table.get(block) {
            decoded.push(*letter);
        }
    }

    decoded
}

