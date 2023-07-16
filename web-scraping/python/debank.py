from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import NoSuchElementException

def high_level_scrape(url):
    # Configure Chrome options to run headless (without a visible browser window)
    chrome_options = Options()
    chrome_options.add_argument("--headless")

    # Instantiate the Chrome driver
    driver = webdriver.Chrome(options=chrome_options)

    try:
        # Load the URL in the headless Chrome browser
        driver.get(url)

        # Wait for the page to load (you can adjust the wait time based on the website's performance)
        driver.implicitly_wait(10)  # Waits up to 10 seconds for elements to appear

        # Print the raw page content (optional)
        print(driver.page_source)

        # Find the specific data using the provided selector
        selector = "#root > div > div.DesktopFrame_mainContainer__2V8Re > div.container_mainSubContainer__39U6P > div.DesktopContainer_main__MG15v.container_pageCenterSubContainer__3encx"
        try:
            specific_data_element = driver.find_element_by_css_selector(selector)
            extracted_data = specific_data_element.text
            print("Extracted Data:")
            print(extracted_data)
        except NoSuchElementException:
            print("Selector did not match any elements. Please check the selector.")

    finally:
        # Close the browser window and terminate the WebDriver
        driver.quit()

if __name__ == "__main__":
    url = "https://debank.com/profile/0x1c45e086ed143aef83c1209521a2ff5369f39abc?chain=arb"
    high_level_scrape(url)
