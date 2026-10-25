# Author = Vishniakov Aleksandr Sergeevich
# Group = P3110
# Date = 15.10.2025
# Variant = 3

import pytest
from Informatics_Lab3_Task3 import task3

@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("KL4D0VK4_P4VLU*xx!july23>13377331WxR", "Rule 5: The digits in your password must add up to 25."),
        ("Ab1*", "Rule 1: Your password must be at least 5 characters."),
        ("Abc*Xyz!!", "Rule 2: Your password must include a digit."),
        ("abc*1xyz", "Rule 3: Your password must include an uppercase letter."),
        ("abcxyzMay12", "Rule 4: Your password must include a special character."),
        ("Abc*May123", "Rule 5: The digits in your password must add up to 25."),
        ("Abc*XYZ25972", "Rule 6: Your password must include a month of the year."),
        ("Pass*July25972", "Your password is valid!"),
    ]
)

def test_task3(text, expected):
    got = task3(text)
    assert expected == got
