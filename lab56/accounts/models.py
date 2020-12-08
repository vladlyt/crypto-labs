from django.contrib.auth.models import AbstractUser
from encrypted_fields import fields

from django.contrib.auth.base_user import BaseUserManager


class UserManager(BaseUserManager):
    pass


class User(AbstractUser):
    phone = fields.EncryptedCharField(max_length=48, help_text="Your phone number", null=True, blank=True)
    address = fields.EncryptedCharField(max_length=200, help_text="Your address", null=True, blank=True)

    objects = UserManager()

    def __str__(self):
        return self.username
