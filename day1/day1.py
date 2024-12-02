def main() -> None:
    part1_solution = part1()
    print(f"Part 1: {part1_solution}")
    part2_solution = part2()
    print(f"Part 2: {part2_solution}")


def part1() -> int:
    lst1, lst2 = load_numbers("input.txt")
    lst1.sort()
    lst2.sort()

    total = 0
    for i in range(len(lst1)):
        diff = lst1[i] - lst2[i]
        total += abs(diff)
    return total


def part2() -> int:
    lst1, lst2 = load_numbers("input.txt")
    lst1.sort()
    lst2.sort()

    simularity_score = 0
    for num in lst1:
        count = lst2.count(num)
        simularity_score += num * count

    return simularity_score


def load_numbers(file_name: str) -> tuple[list[int], list[int]]:
    lst1 = []
    lst2 = []
    with open(file_name, mode="r") as f:
        for line in f.readlines():
            split_line = line.split()
            lst1.append(int(split_line[0]))
            lst2.append(int(split_line[1]))
    return (lst1, lst2)


if __name__ == "__main__":
    main()
