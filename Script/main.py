import requests
from bs4 import BeautifulSoup
import pymysql
from datetime import datetime

# 1. 抓取网页表格
url = "https://wow.blizzard.cn/news/20250821/40565_1255278.html"
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
}
response = requests.get(url, headers=headers)
response.encoding = 'utf-8'

soup = BeautifulSoup(response.text, "html.parser")
table = soup.select_one(".detail table")

data = []
if table:
    rows = table.find_all("tr")
    for row in rows:
        cols = row.find_all(["td", "th"])
        cols_text = [col.get_text(strip=True) for col in cols]
        data.append(cols_text)
else:
    print("未找到 table")
    exit()

# 2. 连接 MySQL 数据库
conn = pymysql.connect(
    host="localhost",
    port=3306,
    user="root",
    password="",
    database="ranking",  # 替换为你的数据库名
    charset="utf8mb4"
)
cursor = conn.cursor()

# 3. 插入数据
violation_time = datetime.strptime("2025-08-21", "%Y-%m-%d")  # 固定日期
for row in data[1:]:  # 跳过表头
    player_name = row[0] if len(row) > 0 else ''
    server_name = row[1] if len(row) > 1 else ''
    violation_reason = row[2] if len(row) > 2 else ''
    violation_count = row[3] if len(row) > 3 else ''

    sql = """
        INSERT INTO player_violations
        (player_name, server_name, violation_reason, violation_count, violation_time)
        VALUES (%s, %s, %s, %s, %s)
    """
    cursor.execute(sql, (player_name, server_name, violation_reason, violation_count, violation_time))

conn.commit()
cursor.close()
conn.close()

print("数据插入完成！")
