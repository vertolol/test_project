# from fetch_products_data.spiders.fetch_data import FetchProductsDataSpider
# from scrapy.crawler import CrawlerProcess
# from scrapy.utils.project import get_project_settings
#
# import argparse
#
#
# process = CrawlerProcess(
#     get_project_settings()
# )


def link_preparation(raw_link: str) -> str:
    page = "?p={}&"
    result = list(raw_link.split("/"))
    result[6] = page + result[6][1:]
    return '/'.join(result)


# def start_spider(url: str) -> None:
#     url = link_preparation(url)
#     process.crawl(FetchProductsDataSpider, page_link_template=url)
#
#
# parser = argparse.ArgumentParser()
# parser.add_argument("--link", help="This is the 'link' variable")
#
# args = parser.parse_args()


# if args.link:
#     start_spider(args.link)
# else:
#     start_spider("https://www.dns-shop.ru/catalog/4f13d297614d7fd7/nabory-elektroinstrumentov/")
#
# process.start()

# scrapy crawl crawl_dns -a url=

l = "https://www.dns-shop.ru/catalog/8a9ddfba20724e77/ssd-nakopiteli/?order=1&groupBy=none&stock=2&price=8001-13000"


def url_to_template(raw_url):
    page = "?p={}&"
    result = list(raw_url.split("/"))
    result[6] = page + result[6][1:]
    return '/'.join(result)

print(url_to_template(l).format(1))