# Author = Vishniakov Aleksandr Sergeevich
# Group = P3110
# Date = 15.10.2025
# Variant = 0

import re

def task2(text: string):
    p = r"^[0-9A-Za-z_.]+@([A-Za-z]+(?:\.[A-Za-z]+)+)$"
    a = re.findall(p, text)
    return a[0] if len(a)==1 else "Fail!"

