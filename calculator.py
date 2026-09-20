"""A simple calculator."""


def add(a, b):
    """Add together two positive integers.

    Raises ValueError if either argument is not a positive integer.
    """
    for value in (a, b):
        if not isinstance(value, int) or isinstance(value, bool):
            raise ValueError("arguments must be integers")
        if value <= 0:
            raise ValueError("arguments must be positive integers")
    return a + b
