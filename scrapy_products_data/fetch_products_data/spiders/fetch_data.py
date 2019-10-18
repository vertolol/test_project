import scrapy


class FetchProductsDataSpider(scrapy.Spider):
    name = "crawl_dns"

    base_url = "https://www.dns-shop.ru"
    start_urls = []

    def __init__(self, url: str):
        super().__init__(self)
        self.page_number = 1
        self.raw_url = url
        print(f"----------------{url}----------------")
        self.page_url_template = self.url_to_template()
        self.page_url = self.page_url_template.format(self.page_number)

        self.start_urls.append(self.page_url)

    def url_to_template(self):
        page = "?p={}&"
        result = list(self.raw_url.split("/"))
        result[6] = page + result[6][1:]
        return '/'.join(result)

    def parse(self, response):
        product_links_xpath = r"//div[@class='product-info__title']/div/a"

        for link in response.xpath(product_links_xpath):
            link = link.attrib['href']
            yield response.follow(link, callback=self.parse_data)

        next_page_xpath = "//div[2]/div/ul/li/a[@class='pagination-widget__page-link pagination-widget__page-link_next ']"
        next_page = response.xpath(next_page_xpath).get()

        if next_page is not None:
            self.page_number += 1
            next_page_link = self.page_url_template.format(self.page_number)
            yield response.follow(next_page_link, callback=self.parse)

    def parse_data(self, response):
        name = response.xpath("//body/div/div/h1/text()").get(),
        price = response.xpath("//div[2]/div/div[3]/div/div/div/div/div/div/span/text()").get(),
        description = response.xpath("//div[4]/div/div/div/div/p/text()").getall()

        yield {
            "name": name,
            "price": price,
            "description": description,
        }
