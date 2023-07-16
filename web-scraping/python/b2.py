import requests
from bs4 import BeautifulSoup

def scrape_hyperlinks(url):
    response = requests.get(url)
    
    if response.status_code == 200:
        html_content = response.text
    else:
        print("Failed to fetch the content. Status code:", response.status_code)
        return []

    soup = BeautifulSoup(html_content, 'html.parser')
    hyperlinks = []

    for link in soup.find_all('a'):
        hyperlink = link.get('href')
        if hyperlink:
            hyperlinks.append(hyperlink)

    return hyperlinks

if __name__ == "__main__":
    target_url = "https://app.beefy.com/dashboard/0x1c45e086ed143aef83c1209521a2ff5369f39abc"
    scraped_hyperlinks = scrape_hyperlinks(target_url)
    
    if scraped_hyperlinks:
        print("Scraped Hyperlinks:")
        for hyperlink in scraped_hyperlinks:
            print(hyperlink)
    else:
        print("No hyperlinks found.")
