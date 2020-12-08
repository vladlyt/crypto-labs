## Lab 5

For this lab we used Django framework on Python 3.8.1 version.
This is very powerful framework with a lot of security mechanisms in it (like XSS protection, CSRF protection, SQL injection protection and so on)

For password hasher we using bcrypt algorithm from `django.contrib.auth.hashers.BCryptSHA256PasswordHasher`. 
It is very popular and secure algorithm for storing passwords.

Also added we've added some restrictions on customers passwords such as:

    Password can’t be too similar to your other personal information.
    Password must contain at least 8 characters.
    Password can’t be a commonly used password.
    Password can’t be entirely numeric.
    

## Lab 6

For database encryption we used encrypted_fields library. 
This is lightweight lib, that provide us to store AEAD-encrypted data in database for fields, that we want to be encrypted.
It is using AES256 to encrypt data and also using `FIELD_ENCRYPTION_KEYS` variable in settings, to decode/encode data using first key from this variable.
In our case this is just data from `.encryption_keys` file, but in real project this should be 
KMS provided by any major cloud provider. Also this keys can be rotated.

Information can be stolen if someone will get access to the database, also will know the `.encryption_keys` file located on the server
or will get access to the KMS (which is much harder). Or if he will be able to exec some malicious code onto the running production server, 
because we are storing `FIELD_ENCRYPTION_KEYS` in global settings. But it is very hard to make in real life =)

