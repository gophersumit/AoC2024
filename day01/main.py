def part1(file):
    sample = open(file, 'r')
    first, second = [], []
    for line in sample:
        words = line.split()
        f, s = words[0], words[1]
        first.append(int(f))
        second.append(int(s))

    first.sort()
    second.sort()

    distance = []
    for a, b in zip(first, second):
        d = abs(a - b)
        distance.append(d)
    print('distance for file:',file)
    print(sum(distance))

def part2(file):
    sample = open(file, 'r')
    first =  []
    second = {}
    similarity = []
    for line in sample:
        words = line.split()
        f, s = int(words[0]), int(words[1])
        first.append(f)
        if s not in second:
            second.setdefault(s, 1)
        else:
            second[s] += 1

    for a in first:
        if a in second:
            s = a* second[a]
            similarity.append(s)

    print('similarity for file:',file)
    print(sum(similarity))


part1('sample.txt')
part1('input.txt')


part2('sample.txt')
part2('input.txt')