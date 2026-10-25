# Author = Vishniakov Aleksandr Sergeevich
# Group = P3110
# Date = 15.10.2025
# Variant = 2

import re

def task3(text: string):
    rule1 = r"^.{5,}$"
    rule2 = r"[0-9]"
    rule3 = r"[A-Z]"
    rule4 = r"[^0-9A-Za-z\s]"
    rule5 = r"[0-9]"
    rule6 = r"(?i)january|february|march|april|may|june|july|august|september|october|november|december"

    if len(re.findall(rule1, text)) == 0:
        return "Rule 1: Your password must be at least 5 characters."

    if len(re.findall(rule2, text)) == 0:
        return "Rule 2: Your password must include a digit."

    if len(re.findall(rule3, text)) == 0:
        return "Rule 3: Your password must include an uppercase letter."

    if len(re.findall(rule4, text)) == 0:
        return "Rule 4: Your password must include a special character."

    if sum(int(x) for x in re.findall(rule5, text)) != 25:
        return "Rule 5: The digits in your password must add up to 25."

    if len(re.findall(rule6, text)) == 0:
        return "Rule 6: Your password must include a month of the year."

    return "Your password is valid!"
