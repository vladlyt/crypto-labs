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
    

