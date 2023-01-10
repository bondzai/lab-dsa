from bs4 import BeautifulSoup
import requests

url = "https://stackpython.co/courses"
res = requests.get(url)
res.encoding = "utf-8"
print(res)
soup = BeautifulSoup(res.text, 'html.parser')
print(soup.prettify())