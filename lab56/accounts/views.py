from django.contrib.auth import login, authenticate
from django.contrib.auth.forms import UserCreationForm
from django.shortcuts import render, redirect
from django.views.generic import View


class RegistrationView(View):
    http_method_names = ('get', 'post')

    def post(self, request, *args, **kwargs):
        form = UserCreationForm(request.POST)
        if form.is_valid():
            form.save()
            username = form.cleaned_data.get('username')
            raw_password = form.cleaned_data.get('password1')
            user = authenticate(username=username, password=raw_password)
            login(request, user)
            return redirect('home')
        return render(request, 'registration/registration.html', {'form': form})

    def get(self, request, *args, **kwargs):
        return render(request, 'registration/registration.html', {'form': UserCreationForm()})
