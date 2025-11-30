import pdfplumber
import pandas as pd
import pymysql
from pymysql.cursors import DictCursor
from datetime import datetime

# -------------------------- 1. 配置参数 --------------------------
# 处罚PDF文件本地路径（替换为处罚PDF实际名称）
PDF_PATH = "1759654395264_a6eba0e59a.pdf"
# MySQL数据库连接参数
DB_CONFIG = {
    "host": "127.0.0.1",
    "port": 3306,
    "user": "ranking",
    "password": "ranking",
    "database": "ranking",
    "charset": "utf8mb4"
}
# 固定违规原因
VIOLATION_REASON = "史诗团本"
violation_time = datetime.strptime("2025-11-13", "%Y-%m-%d")

# -------------------------- 2. 读取PDF并提取数据 --------------------------
def extract_pdf_data(pdf_path):
    """提取PDF中的角色昵称、服务器名、违规次数"""
    all_text = ""
    # 读取PDF所有页面文本
    with pdfplumber.open(pdf_path) as pdf:
        for page in pdf.pages:
            page_text = page.extract_text()
            if page_text:
                all_text += page_text + "\n"  # 保留换行符用于分行

    # 按行解析数据
    lines = all_text.split("\n")
    data_list = []
    header_found = False  # 标记是否找到表头（角色昵称、服务器名、违规次数）

    for line in lines:
        line = line.strip()
        if not line:
            continue  # 跳过空行

        # 找到表头后开始处理数据行
        if "角色名" in line and "服务器名" in line:
            header_found = True
            continue
        if not header_found:
            continue

        # 定位违规次数（仅两种取值：首次违规/多次违规）
        if "首次违规" in line:
            violation_count = "首次违规"
            content = line.replace("首次违规", "").strip()
        elif "多次违规" in line:
            violation_count = "多次违规"
            content = line.replace("多次违规", "").strip()
        else:
            continue  # 跳过无明确违规次数的行

        # violation_count = "多次违规"
        # # 移除文本中的违规标记
        # content = line.replace("首次违规", "").replace("多次违规", "").strip()

        # 分割角色昵称和服务器名（从后往前找第一个空格，避免昵称含空格）
        if " " in content:
            space_index = content.rfind(" ")
            player_name = content[:space_index].strip()
            server_name = content[space_index:].strip()

            # 过滤无效数据（如昵称/服务器名为空）
            if player_name and server_name:
                data_list.append({
                    "player_name": player_name,
                    "server_name": server_name,
                    "violation_reason": VIOLATION_REASON,
                    "violation_count": violation_count,
                    "violation_time": violation_time  # 新增：添加固定处罚时间
                })

    return data_list


# -------------------------- 3. 插入MySQL数据库 --------------------------
def insert_into_mysql(data_list, db_config):
    """批量插入数据到player_violations表"""
    if not data_list:
        print("无有效数据可插入！")
        return

    # SQL插入语句（使用参数化查询避免SQL注入）
    insert_sql = """
                 INSERT INTO player_violations (player_name, server_name, violation_reason, violation_count, violation_time)
                 VALUES (%s, %s, %s, %s, %s) \
                 """

    # 准备插入参数（按SQL顺序整理数据）
    insert_params = [
        (
            item["player_name"],
            item["server_name"],
            item["violation_reason"],
            item["violation_count"],
            item["violation_time"]
        ) for item in data_list
    ]

    conn = None
    try:
        # 建立数据库连接
        conn = pymysql.connect(**db_config)
        cursor = conn.cursor()

        # 批量执行插入（效率高于单条插入）
        cursor.executemany(insert_sql, insert_params)
        conn.commit()  # 提交事务

        print(f"数据插入成功！共插入 {cursor.rowcount} 条记录")

    except pymysql.MySQLError as e:
        if conn:
            conn.rollback()  # 插入失败回滚事务
        print(f"数据库插入失败：{e}")

    finally:
        # 关闭连接
        if conn:
            cursor.close()
            conn.close()
            print("数据库连接已关闭")


# -------------------------- 4. 主执行逻辑 --------------------------
if __name__ == "__main__":
    # 步骤1：提取PDF数据
    print("开始读取PDF数据...")
    violation_data = extract_pdf_data(PDF_PATH)
    print(f"PDF数据提取完成，共获取 {len(violation_data)} 条有效记录")

    # 步骤2：插入数据库
    if violation_data:
        print("\n开始插入数据库...")
        insert_into_mysql(violation_data, DB_CONFIG)
    else:
        print("\n无数据可插入数据库")