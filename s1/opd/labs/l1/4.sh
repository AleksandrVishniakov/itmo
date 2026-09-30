cd lab0

echo "--- Task 1 ---"
chmod u+r clamperl
cd clamperl
wc -l $(ls) 2>/dev/null | sort -nr
cd ..
chmod u-r clamperl

echo -e "\n--- Task 2 ---"

ls -Rltr | grep 'on' | grep -v "^./" | head -n 3

echo -e "\n--- Task 3 ---"

cat clamperl/exeggcute clamperl/politoed clamperl/porygon cubone/grumpig cubone/silcoon cubone/masquerain | grep "Re"

echo -e "\n--- Task 4 ---"

cat *e */*e */*/*e 2>/dev/null | sort

echo -e "\n--- Task 5 ---"

ls -ldt $(echo c* */c* */*/c*) 2>&1 | grep -v "^ls:\ \*"

echo -e "\n--- Task 6 ---"

cat f* */f* */*/f* 2>/tmp/e | sort -r | cat -n

