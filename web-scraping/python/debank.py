import os
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from dotenv import load_dotenv

def high_level_scrape(url):
    chrome_options = Options()
    chrome_options.add_argument("--headless")

    # Instantiate the Chrome driver
    driver = webdriver.Chrome(options=chrome_options)

    try:
        # Load the URL in the headless Chrome browser
        driver.get(url)

        # Wait for the headers to be present
        headers = wait_for_headers(driver)
        print(headers)

    finally:
        driver.quit()

def wait_for_headers(driver):
    # Define a wait condition for header elements to be present
    wait = WebDriverWait(driver, 10)
    wait.until(EC.presence_of_all_elements_located((By.XPATH, "//h1 | //h2 | //h3 | //h4 | //h5 | //h6")))

    # Find all header elements using XPath
    headers = driver.find_elements(By.XPATH, "//h1 | //h2 | //h3 | //h4 | //h5 | //h6")

    # Extract the text from the headers
    header_texts = [header.text for header in headers]

    return header_texts

if __name__ == "__main__":
    load_dotenv()
    url = os.getenv("URL")
    url = url + "?chain=arb"

    if not url:
        print("URL not found in .env file.")
    else:
        high_level_scrape(url)
