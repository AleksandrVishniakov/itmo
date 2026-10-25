# Author = Vishniakov Aleksandr Sergeevich
# Group = P3110
# Date = 15.10.2025
# Variant = 3

import re

def task1(text: string):
    p = r"\b([А-ЯЁ][a-яё]+(?:\-[A-ЯЁ][а-яё]+)?)\s+(?:[А-ЯЁ]\.)(?:-[А-ЯЁ]\.)*(?:[А-ЯЁ]\.)?(?:[^А-ЯЁ]|$)"
    a = re.findall(p, text)
    return sorted(a)

