class JamesBond:

    def map_data(self, input_list_of_dict):
        data = []
        key_mapping = {
            "project_id": "id",
            "project_code": "code",
            "project_type": "type",
        }
        for data in input_list_of_dict:
            data.append(self.rename_keys(data, key_mapping))
        return {
            "result": data,
        }


    def rename_keys(self, input_dict, key_mapping):
        output_dict = {}
        
        for key, value in input_dict.items():
            if key in key_mapping:
                new_key = key_mapping[key]
                output_dict[new_key] = value
            else:
                output_dict[key] = value
                
        return output_dict
