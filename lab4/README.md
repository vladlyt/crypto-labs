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

`hashcat -a 0 -m 0 input-passwords-md5.csv  top-1000000-passwords.txt -o output-passwords-md5.csv`

Straight mode MD5 passwords takes 36 seconds to restore with known passwords
Recovered: 86678/175911 (49.27%) Digests


`hashcat -a 3 -m 0 input-passwords-md5.csv  -o output-passwords-md5-brute.csv`

Brute-force takes 16 minutes to restore 28800 passwords from 190000 and it'll take more than 10 hours
------

`hashcat -a 0 -m 110 input-passwords-sha1-semicolumn.csv  top-1000000-passwords.txt -o output-passwords-sha1.csv`

Straight mode SHA1 passwords takes 2 hours 10 minutes 55 seconds to restore with known passwords
Recovered: 100615/190000 (52.95%)

`hashcat -a 3 -m 110 input-passwords-sha1-semicolumn.csv -o output-passwords-sha1-brute.csv`

Brute-force SHA1 takes 41 minutes to recover 7 passwords from 190000 (it considered only 2 and 3 digit passwords). Total estimation time was so big.

------

`hashcat -a 0 -m 3200 input-passwords-bcrypt.csv top-1000000-passwords.txt -o output-passwords-bcrypt.csv`

BCRYPT will takes 8 years to decode 190000 hashes in straight mode.

------

Conclusion: The strongest hashing scheme turned out to be bcrypt. Estimation time takes more than 8 year even in dictionary mode.
The weakest hashing scheme turned out to be MD5, less weak – SHA1.

We used dictionary and brute-force attacks. Brute-force became more effective.


