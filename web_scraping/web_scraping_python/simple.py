from bs4 import BeautifulSoup

import requests

target_url = "https://coinmarketcap.com/currencies/bitcoin/"
res = requests.get(target_url)
res.encoding = "utf-8"
print(res)
soup = BeautifulSoup(res.text, 'html.parser')
data_all = soup.prettify()
find_all = soup.find_all("div",{"class":"priceValue"})
print(find_all)

btc_value = []

for i in find_all:
    i = str(i).split('<div class="priceValue"><span>', 1)
    i = str(i).split('</span></div>', 1)
    btc_value = i
    break

output = btc_value[0].split("['', '", 1)
print('BTC price = ' + output[1])