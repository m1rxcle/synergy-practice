from django.db import models


# Create your models here.
class Django(models.Model):
    name = models.CharField(max_length=50, db_index=True)

    def __str__(self):
        return self.name
