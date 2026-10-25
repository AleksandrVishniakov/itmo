import re

def task3(text):
    return re.sub(r"(?<![0-9]\.)(?<![0-9])[0-9]+(?!\.[0-9]+)", lambda m: str(int(m.group())**3*5-13), text)
    

print(task3("aksk34.aksk aksa 8+ 9 == 34-5.012 hello 8"))