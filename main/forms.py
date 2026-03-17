from .models import Django
from django.forms import ModelForm, TextInput
from django import forms


class User_Form(ModelForm):
    class Meta:
        model = Django
        fields = ["name"]
        widgets = {
            "name": TextInput(
                attrs={
                    "id": "input",
                    "class": "form-input",
                    "placeholder": "Ваше имя",
                    "type": "text",
                    "required": "required",
                },
            ),
        }
        error_messages = {
            "name": {"required": "Поле не должно быть пустым!"},
        }

    def clean_name(self):
        name = self.cleaned_data.get("name", "")

        if not name.strip():
            raise forms.ValidationError("Поле не должно быть пустым!")

        special_chars = "!@#$%^&*()_+-=[]{}|;:,.<>/?`~\\"

        if any(char in name for char in special_chars):
            raise forms.ValidationError("Имя не должно содержать специальные символы!")

        return name
