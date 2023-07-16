import os
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from dotenv import load_dotenv

def high_level_scrape(url):
    # Configure Chrome options to run headless (without a visible browser window)
    chrome_options = Options()
    # chrome_options.add_argument("--headless")
    chrome_options.add_argument("--disable-gpu")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_options.add_argument("--disable-setuid-sandbox")
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-dev-shm-usage")

    # Instantiate the Chrome driver
    driver = webdriver.Chrome(options=chrome_options)

    try:
        # Load the URL in the headless Chrome browser
        driver.get(url)

        # Wait for the specific element to load
        WebDriverWait(driver, 10).until(EC.presence_of_element_located((By.CSS_SELECTOR, '#Overview_defiItem__1e5s9 > div:nth-child(2)')))

        # Get the element's data
        element = driver.find_element(By.CSS_SELECTOR, '#Overview_defiItem__1e5s9 > div:nth-child(2)')
        print(element.text)  # print the content of the element

    finally:
        driver.quit()

if __name__ == "__main__":
    load_dotenv()
    url = os.getenv("URL")
    url = url + "?chain=arb"

    if not url:
        print("URL not found in .env file.")
    else:
        high_level_scrape(url)
