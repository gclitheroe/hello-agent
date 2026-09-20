"""Tests for the calculator."""

import unittest

from calculator import add


class TestAdd(unittest.TestCase):
    def test_add_small_positive_integers(self):
        self.assertEqual(add(1, 2), 3)

    def test_add_larger_positive_integers(self):
        self.assertEqual(add(100, 250), 350)

    def test_add_is_commutative(self):
        self.assertEqual(add(7, 13), add(13, 7))

    def test_add_rejects_non_positive_integers(self):
        with self.assertRaises(ValueError):
            add(0, 5)

    def test_add_rejects_non_integers(self):
        with self.assertRaises(ValueError):
            add(1.5, 2)


if __name__ == "__main__":
    unittest.main()
