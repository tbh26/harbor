#! /usr/bin/env python3

import collections
import os
import sys

my_path = sys.argv[0]
me = os.path.basename(my_path)
dirname = os.path.dirname(my_path)

Card = collections.namedtuple('Card', ['rank', 'suit'])


class FrenchDeck:
    ranks = [str(n) for n in range(2, 11)] + list('JQKA')
    suits = 'spades diamonds clubs hearts'.split()

    def __init__(self):
        self._cards = [Card(rank, suit) for suit in self.suits for rank in self.ranks]

    def __len__(self):
        return len(self._cards)

    def __getitem__(self, position):
        return self._cards[position]


def first_demo():
    print(' = first_demo =')
    print()
    #
    some_card = Card(rank='Q', suit='spades')
    other_card = Card(rank='J', suit='clubs')
    print(f'  {some_card=}')
    print(f'  {other_card=}')
    print(f'  {Card(rank='10', suit='hearts')}')
    print()


def next_demo():
    print(' = next_demo =')
    print()

    fd = FrenchDeck()
    print(f'  {fd[0]=}')
    print(f'  {fd[1]=}')
    print(f'  {fd[-1]=}')
    print(f'  {len(fd)=}')
    print()

    for n, card in enumerate(fd, 1):
        print(f'    card {n: >2} == {card} ')
    print()


def other_demo():
    print(' = other_demo =')
    print()

    suit_values = dict(spades=3, hearts=2, diamonds=1, clubs=0)
    print(f'  {suit_values=}')

    def spades_high(card):
        rank_value = FrenchDeck.ranks.index(card.rank)
        return rank_value * len(suit_values) + suit_values[card.suit]

    some_cards = [Card(rank='3', suit='clubs'), Card(rank='J', suit='diamonds'), Card(rank='10', suit='hearts'), Card(rank='A', suit='spades')]
    for card in some_cards:
        print(f'  {card=}\t\t{spades_high(card)=}')
    print()

    print('  listing sorted(FrenchDeck(), key=spades_high) ==> ')
    fd = FrenchDeck()
    for card in sorted(fd, key=spades_high):
        print(f'   {card=}\t\t{spades_high(card)=}')
    print()


def main():
    print(f'{me=}   ( {dirname=} )')
    print()
    #
    first_demo()
    #
    next_demo()
    #
    other_demo()


if __name__ == '__main__':
    main()
