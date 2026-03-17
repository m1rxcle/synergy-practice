from django.contrib import admin
from .models import Django


# Register your models here.
@admin.register(Django)
class UserAdmin(admin.ModelAdmin):
    list_display = [
        "name",
    ]
