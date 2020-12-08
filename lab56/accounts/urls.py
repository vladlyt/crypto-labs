from django.urls import path
from . import views

urlpatterns = [
    path('signup/', views.RegistrationView.as_view(), name='signup'),
    path('user-data/<int:pk>/', views.UpdateUserDataView.as_view(), name='update-user-data'),
]
