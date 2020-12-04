# Part 1

## What was used: 

`go1.15.1` used in this lab. 

For pseudo-random values was used default `math/rand` lib.

Totally created 200000 passwords in `generated-sha1.csv`, `generated-md5.csv`, `generated-bcrypt.csv`

### Top common passwords

Top common passwords got from https://github.com/danielmiessler/SecLists, they are stored in 
`top-100-passwords.txt` and `top-1000000-passwords.txt`

### How top 100 passwords was generated

Passwords is randomly selected from the `top-100-passwords.txt`

5% of all passwords 

### How top 1000000 passwords was generated

Passwords is randomly selected from the `top-1000000-passwords.txt`

80% of all passwords

### How random passwords was generated

Passwords is randomly generated from charset: `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%&`
with random length from 5 to 14

5% of all passwords

### How passwords with rules are generated

Passwords is randomly selected from the `top-1000000-passwords.txt`
and randomly applied one of the following rules:

`ToUpperCase`

`ToLowerCase`

`Capitalize`

`Reverse`

`AddNumbersToStart` (count of numbers from 1 up to 5 (included))

`AddNumbersToEnd` (count of numbers from 1 up to 5 (included))

10% of all passwords

# Part 2

Got passwords from https://github.com/o-rumiantsev/Cryptography/tree/master/lab4/part-1

Renamed them into `input-passwords-bcrypt.csv`, `input-passwords-md5.csv`, `input-passwords-sha1.csv`
