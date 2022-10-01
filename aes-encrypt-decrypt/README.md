# AES messages and files encrypt and decrypt

Encrypt / Decrypt messages

Usage:
```bash
./main <enc | dec> msg <message> <secretKey>
```
Encrypt / Decrypt files

Usage:
```bash
./main <enc | dec> file <filename> <secretKey>
```

### Examples

##### Encrypt message
```bash
./main enc msg hello 123
```
Result: AAAAAAAAAAAAAAAAV2qe+jmHxzPsOYub7tuIdu8K5fbF

##### Decrypt message
```bash
./main enc msg AAAAAAAAAAAAAAAAV2qe+jmHxzPsOYub7tuIdu8K5fbF 123
```
Result: 123

##### Encrypt file
```bash
./main enc file test.png 123
```
##### Decrypt file
```bash
./main dec file enc_test.png 123
```