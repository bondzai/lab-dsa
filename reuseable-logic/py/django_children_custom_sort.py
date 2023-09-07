# Many to many field models

class DataGroupSerializer(serializers.ModelSerializer):
    class Meta:
        model = DataGroup
        fields = "__all__"

    def to_representation(self, instance):
        data = super().to_representation(instance)
        
        sorted_data = sorted(
            instance.field_name.all(),
            key=lambda target_field: target_field.find_field
        )

        data["field_name"] = ChildrenSerializer(
            sorted_data, many=True
        ).data
        
        return data
