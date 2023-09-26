class ModelA(models.Model):
    pass

class ModelB(models.Model):
    a = ForeignKey(ModelA)

# Forward ForeignKey relationship
ModelB.objects.select_related('a').all()

# Reverse ForeignKey relationship
ModelA.objects.prefetch_related('modelb_set').all() 
