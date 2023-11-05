import os
import time
import random
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from dotenv import load_dotenv

def high_level_scrape(url):
    chrome_options = Options()
    chrome_options.add_argument("--headless")

    driver = webdriver.Chrome(options=chrome_options)

    try:
        driver.get(url)

        mimic_human_behavior(driver)

        headers = wait_for_element(driver)
        print(headers)

    finally:
        driver.quit()

def mimic_human_behavior(driver):
    time.sleep(random.uniform(2, 5))
    
def wait_for_element(driver):
    wait = WebDriverWait(driver, 60)
    wait.until(EC.visibility_of_element_located((By.CSS_SELECTOR, '#Overview_defiItem__1e5s9 > div:nth-child(2)')))
    element = driver.find_element(By.CSS_SELECTOR, '#Overview_defiItem__1e5s9 > div:nth-child(2)')
    data = element.text

    return data



if __name__ == "__main__":
    load_dotenv()
    url = os.getenv("URL")
    url = url + "?chain=arb"

    if not url:
        print("URL not found in .env file.")
    else:
        high_level_scrape(url)
