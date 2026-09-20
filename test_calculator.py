import unittest

from calculator import add_positive_integers


class AddPositiveIntegersTest(unittest.TestCase):
    def test_adds_two_positive_integers(self):
        self.assertEqual(add_positive_integers(2, 3), 5)

    def test_adds_multiple_positive_integers(self):
        self.assertEqual(add_positive_integers(1, 2, 3, 4), 10)

    def test_adds_large_positive_integers(self):
        self.assertEqual(add_positive_integers(1_000_000, 2_500_000), 3_500_000)

    def test_rejects_zero(self):
        with self.assertRaisesRegex(ValueError, "positive"):
            add_positive_integers(1, 0)

    def test_rejects_non_integer_values(self):
        with self.assertRaisesRegex(TypeError, "integers"):
            add_positive_integers(1, "2")


if __name__ == "__main__":
    unittest.main()
