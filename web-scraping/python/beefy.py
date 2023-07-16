import os
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from dotenv import load_dotenv
from bs4 import BeautifulSoup
from time import sleep

def high_level_scrape(url):
    chrome_options = Options()
    chrome_options.add_argument("--headless")

    driver = webdriver.Chrome(options=chrome_options)

    try:
        driver.get(url)
        sleep(5)
        
        page_source = driver.page_source

        soup = BeautifulSoup(page_source, "html.parser")

        body = soup.find('body')
        if body:
            print(body.get_text())

    finally:
        driver.quit()

if __name__ == "__main__":
    load_dotenv()
    url = os.getenv("BEEFY_URL")

    if not url:
        print("URL not found in .env file.")
    else:
        high_level_scrape(url)
