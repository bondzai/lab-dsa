import requests


url = 'https://notify-api.line.me/api/notify'
token = 'HSjKlkm6ylCcZoHgtaAbk37EZqOdW52lrXuIw1CFqNb'
token_cj = '7LDFKvhJLfzlAbDry9125QaDYTCZCGHLMcA5H5yJMcE'
headers = {'content-type':'application/x-www-form-urlencoded','Authorization':'Bearer '+ token}

message = 'Hello from Python task scheduler.(executed every 10 minutes)'
sticker_package = 11537
sticker_id = 52002735

r = requests.post(url, headers=headers, data = {
    'message':message,
    'stickerPackageId':sticker_package,
    'stickerId':sticker_id,
    })

print (r.text)