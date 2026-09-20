"""A small calculator for positive integers."""


def add_positive_integers(*numbers: int) -> int:
    """Return the sum of one or more positive integers."""
    if not numbers:
        raise ValueError("at least one positive integer is required")

    for number in numbers:
        if not isinstance(number, int) or isinstance(number, bool):
            raise TypeError("all values must be integers")
        if number <= 0:
            raise ValueError("all integers must be positive")

    return sum(numbers)
