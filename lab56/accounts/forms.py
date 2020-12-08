from django.contrib.auth import get_user_model
from django.contrib.auth.forms import UserCreationForm, UsernameField
from django.forms import ModelForm

User = get_user_model()


class SignUpForm(UserCreationForm):
    class Meta:
        model = User
        fields = ("username",)
        field_classes = {'username': UsernameField}


class UpdateUserDataForm(ModelForm):
    class Meta:
        model = User
        fields = ('phone', 'address')

    def save(self, commit=True):
        print("LOL", self.data)
        super().save(commit)
