from django.shortcuts import render, get_object_or_404
from django.http import HttpResponse


# Create your views here.


def index(request):
    name = "Mira"
    return render(request, "main/index.html")


def get_user(request):
    return render(request, "main/get-user.html")
