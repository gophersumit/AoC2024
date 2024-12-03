def is_safe(report):
    increasing = all(0 < report[i + 1] - report[i] <= 3 for i in range(len(report) - 1))
    decreasing = all(0 < report[i] - report[i + 1] <= 3 for i in range(len(report) - 1))
    return increasing or decreasing


def is_tolerable_safe(report):
    n = len(report)
    for i in range(n):
        modified = report[:i] + report[i + 1:]
        if is_safe(modified):
            return True

    return False
with open("sample.txt") as f:
    safeCount = 0
    tolerableSafe = 0
    for line in f:
        scores = list(map(int, line.split()))
        if is_safe(scores):
            safeCount += 1
            tolerableSafe+= 1
        elif is_tolerable_safe(scores):
            tolerableSafe+= 1


    print(safeCount)
    print(tolerableSafe)


with open("input.txt") as f:
    safeCount = 0
    tolerableSafe = 0
    for line in f:
        scores = list(map(int, line.split()))
        if is_safe(scores):
            safeCount += 1
            tolerableSafe+= 1
        elif is_tolerable_safe(scores):
            tolerableSafe+= 1


    print(safeCount)
    print(tolerableSafe)

