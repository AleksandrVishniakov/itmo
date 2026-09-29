cd lab0

echo "--- Task 1 ---"
chmod u+r clamperl
cd clamperl
wc -l $(ls) 2>/dev/null | sort -nr
cd ..
chmod u-r clamperl

echo -e "\n--- Task 2 ---"

# ls -Rltr | grep 'on' | head -n 3
find . -name "*on*" -ls 2>/dev/null | sort -k10 | head -n 3

echo -e "\n--- Task 3 ---"

cat clamperl/exeggcute clamperl/politoed clamperl/porygon cubone/grumpig cubone/silcoon cubone/masquerain | grep "Re"

echo -e "\n--- Task 4 ---"

# cat *e */*e */*/*e 2>/dev/null | sort
cat $(find . -name "*e" 2>/dev/null) 2>/dev/null | sort

echo -e "\n--- Task 5 ---"

find . -name "c*" -ls 2>&1 | sort -k10 -r | head -n 2

echo -e "\n--- Task 6 ---"

# cat f* */f* */*/f* 2>/tmp/e | sort -r | cat -n
cat $(find . -name "f*" 2>/tmp/e) 2>/tmp/e | sort -r | cat -n

