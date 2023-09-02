from datetime import datetime

class Utils:
        
    def get_formatted_datetime(self):
            current_datetime = datetime.now()
            formatted_datetime = current_datetime.strftime("%d-%b-%Y %H:%M")

            return formatted_datetime


    # input data: list of dict, search_keyword: string
    def custom_search_filter(self, input_data, search_keyword):
        result = []
        search_columns = ["a", "b"]
        
        for project in input_data:
            for key, value in project.items():
                if key in search_columns:
                    if search_keyword.lower() in str(value).lower():
                        result.append(project)
                        break

        return result
