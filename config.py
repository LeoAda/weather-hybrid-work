import json
CONFIG_FILE = "config.json"
class Config:
    def __init__(self):
        try:
            self.data = self.get_file()
        except:
            self.init_config()
            self.data = self.get_file()

    def set(self, key, value):
        self.data[key] = value
        self.set_file(self.data)

    def set_file(self,data):
        with open(CONFIG_FILE, 'w') as outfile:
            json.dump(data, outfile)
    
    def get(self, key):
        return self.data[key]

    def get_file(self):
        with open(CONFIG_FILE) as json_data_file:
            data = json.load(json_data_file)
        return data
    
    def init_config(self):
        data = {
        # Lille coordinates
        "localisation": [50.6333, 3.0667]
        }
        self.set_file(data)
        