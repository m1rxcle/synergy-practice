from django.shortcuts import render, redirect
from django.http import HttpResponse
from .models import Django
from .forms import User_Form
from urllib.parse import urlencode
from django.urls import reverse

# Create your views here.


def index(request):

    if request.method == "POST":
        form = User_Form(request.POST)

        name = request.POST.get("name", "")

        if form.is_valid():
            user, created = Django.objects.get_or_create(name=name)
            query_string = urlencode({"name": user.name})
            url = reverse("main:get_user")

            return redirect(f"{url}?{query_string}")
    else:
        form = User_Form()

    return render(request, "main/index.html", {"form": form})


def get_user(request):

    name_from_url = request.GET.get("name")

    user = Django.objects.filter(name=name_from_url).first()

    return render(request, "main/get-user.html", {"name": user})
