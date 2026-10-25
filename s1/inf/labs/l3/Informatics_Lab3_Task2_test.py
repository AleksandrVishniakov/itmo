# Author = Vishniakov Aleksandr Sergeevich
# Group = P3110
# Date = 15.10.2025
# Variant = 0

import pytest
from Informatics_Lab3_Task2 import task2


@pytest.mark.parametrize(
    ("text", "expected"),
    [
        ("students.spam@yandex.ru", "yandex.ru"),
        ("example@example.com", "example.com"),
        ("user@y4ndex.ru", "Fail!"),
        ("test@exapmle", "Fail!"),
        ("a@b@c.com", "Fail!"),
        ("name_SURNam3@mail.co.uk", "mail.co.uk"),
        ("students@spam.ru@yandex.ru", "Fail!"),
        ("students.spam@.ru", "Fail!"),
        ("students.spam@yandex.ru.com", "yandex.ru.com"),
        ("students.spam@yandex...ru", "Fail!"),
    ],
)
def test_task2(text, expected):
    got = task2(text)
    assert expected == got
