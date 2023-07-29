import os
import requests
from bs4 import BeautifulSoup
from urllib.parse import urljoin
import re
import argparse
import time



# 发送HTTP请求并下载文件
def download_file(url, path):
    response = requests.get(url, stream=True)
    response.raise_for_status()
    with open(path, 'wb') as file:
        for chunk in response.iter_content(chunk_size=8192):
            if chunk:
                file.write(chunk)

# 解析网页内容，提取文件链接
def parse_page(url):
    response = requests.get(url)
    response.raise_for_status()
    soup = BeautifulSoup(response.text, 'html.parser')
    files = []
    directories = []
    for link in soup.find_all('a'):
        href = link.get('href')
        if href.endswith('/'):
            directories.append(href)
        else:
            files.append(href)
    return directories, files

# 递归下载文件
def download_recursive(url, path, days):
    directories, files = parse_page(url)

    # 获取已下载的文件
    downloaded_files = os.listdir(path)

    # 去除不需要的文件
    pattern = r'^.*\.[a-zA-Z0-9]+$'
    valid_files = []
    for file in files:
        if re.match(pattern, file) and file not in downloaded_files:
            valid_files.append(file)
    
    # 仅保留最近的days天的文件
    target_day = time.strftime("%y%m%d", time.localtime(time.time() - 86400 * days))

    valid_files.sort()
    files = []
    for file in valid_files:
        if file.split('.')[1][2:] >= target_day:
            files.append(file)

    # 下载当前目录的文件
    for file in files:
        file_url = urljoin(url, file)
        file_temp_path = os.path.join(path, "temp")
        file_path = os.path.join(file_temp_path ,file)
        print("Downloading:", file_url)
        download_file(file_url, file_path)
        os.rename(file_path, os.path.join(path, file))
    
    # 去除父目录，父目录名以/开头以/结尾
    pattern = r'^/.*/$'
    valid_directories = []
    for directory in directories:
        if not re.match(pattern, directory):
            valid_directories.append(directory)
    directories = valid_directories

    # 递归下载子目录中的文件
    for directory in directories:
        directory_url = urljoin(url, directory)
        directory_path = os.path.join(path, directory)
        os.makedirs(directory_path, exist_ok=True)
        print("Entering:", directory_url)
        download_recursive(directory_url, directory_path)

# 开始下载
if __name__ == "__main__":
    # 设置要下载的网站URL和目标文件夹

    # 解析参数
    parser = argparse.ArgumentParser()
    parser.add_argument("url", help="URL of the website to download")
    parser.add_argument("path", help="Path of the folder to save files")
    parser.add_argument("days", help="Number of days to keep", type=int)
    args = parser.parse_args()
    base_url = args.url
    target_folder = args.path
    days = args.days

    # python ./script/download.py https://archive.routeviews.org/bgpdata/2023.07/RIBS/ D:/code/platform/BGP/data/unprocessed 10
    # os.makedirs(target_folder+"/temp", exist_ok=True)

    # 删除temp文件夹，重新创建
    if not os.path.exists(target_folder+"/temp"):
        os.makedirs(target_folder+"/temp")

    download_recursive(base_url, target_folder, days)