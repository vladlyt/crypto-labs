from django.contrib.auth import login, authenticate, get_user_model
from django.shortcuts import render, redirect
from django.urls import reverse
from django.views.generic import View, FormView, UpdateView

from accounts.forms import SignUpForm, UpdateUserDataForm

User = get_user_model()


class RegistrationView(View):
    http_method_names = ('get', 'post')

    def post(self, request, *args, **kwargs):
        form = SignUpForm(request.POST)
        if form.is_valid():
            form.save()
            username = form.cleaned_data.get('username')
            raw_password = form.cleaned_data.get('password1')
            user = authenticate(username=username, password=raw_password)
            login(request, user)
            return redirect('home')
        return render(request, 'registration/registration.html', {'form': form})

    def get(self, request, *args, **kwargs):
        return render(request, 'registration/registration.html', {'form': SignUpForm()})


class UpdateUserDataView(UpdateView):
    template_name = 'user_data.html'
    form_class = UpdateUserDataForm
    queryset = User.objects.all()

    def get_success_url(self):
        return reverse('home')
