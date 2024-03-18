#! /usr/bin/env python3

"""
vector2d.py: a simplistic class demonstrating some special methods

It is simplistic for didactic reasons. It lacks proper error handling,
especially in the ``__add__`` and ``__mul__`` methods.

This example is greatly expanded later in the book.

Addition::

    >>> v1 = Vector(2, 4)
    >>> v2 = Vector(2, 1)
    >>> v1 + v2
    Vector(4, 5)

Absolute value::

    >>> v = Vector(3, 4)
    >>> abs(v)
    5.0

Scalar multiplication::

    >>> v * 3
    Vector(9, 12)
    >>> abs(v * 3)
    15.0

"""
import math
import os
import sys

my_path = sys.argv[0]
me = os.path.basename(my_path)
dirname = os.path.dirname(my_path)


class Vector:

    def __init__(self, x=0, y=0):
        self.x = x
        self.y = y

    def __repr__(self):
        return f'Vector({self.x!r}, {self.y!r})'

    def __abs__(self):
        return math.hypot(self.x, self.y)

    def __bool__(self):
        return bool(abs(self))

    def __add__(self, other):
        x = self.x + other.x
        y = self.y + other.y
        return Vector(x, y)

    def __mul__(self, scalar):
        return Vector(self.x * scalar, self.y * scalar)


def demo():
    print(' = demo =')
    #
    v1 = Vector(2, 4)
    v2 = Vector(2, 1)
    print(f'  {v1=}  {v2=}')
    print(f'  {str(v1)=}  {repr(v2)=}')
    v1_plus_v2 = v1 + v2
    print(f'  {v1_plus_v2=}    ( v1_plus_v2 = v1 + v2 ) ')
    print()
    #
    v1_mul_2 = v1 * 2
    print(f'  {v1_mul_2=}    ( v1_mul_2 = v1 * 2  {{ int(2), not v2 }} ) ')
    print()
    #
    print(f'  {bool(v2)=} ')
    print(f'  {bool(Vector(0, 0))=} ')
    print(f'  {bool(Vector(0, 1))=} ')
    print()
    #
    another_vector = Vector(3, 4)
    print(f'  {another_vector=}  {abs(another_vector)=}')
    print(f'  {another_vector * 3=}  {abs(another_vector * 3)=}')
    print(f'  {bool(another_vector)=} ')
    print()
    #


def main():
    print(f'{me=}   ( {dirname=} )')
    print()
    #
    demo()


if __name__ == '__main__':
    main()
