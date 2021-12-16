# AOC 2021

https://adventofcode.com/2021

## Day 01

adding basic lib as you always need something like this in AOC :)
Learning/Fun: start using https://github.com/spf13/cobra 

## Day 02

Easy. Need refactoring for cobra and stuff....

## Day 03

Remove Cobra.... learned enough and it´s "Kanonen auf Spatzen" in this context

## Day 04

baking cookies while programming is not very efficient...

## Day 05

wasted lots of time for stupid bug (didn't clear data for part2 -> double adding lines ....) :(

## Day 06

First solution was not scalable. After rethinking I found a trivial way to implement.

## Day 07

Solution Part 1 should be fine but right now the server is broken... waiting for input data.....
Server is up and running again: Solution was bad, because of bad reading by me.... (Return Fuel and NOT Position)

## Day 08

This was tough... not really programming but analysing combinations of 7-Segment-Displays.
Didn't really like it though

## Day 09

Easy with a "border" around the input and some recursion...

## Day 10

Took some time to get the best idea to follow. Stack was the right choice :)
There are lots of beautifying options but hey..... (maybe later)

## Day 11

Easy

## Day 12

Cheated.... had the right idea with data structure but didn't get the recursion right...
Had to do some working besides AOC2021 so I took the work from https://github.com/pvainio/adventofcode/blob/main/2021/go/d12/main.go 

## Day 13

That was fun and not too hard to implement...

## Day 14

Again a "does not scale" challenge.... have to rethink my structures.
After thinking: Don´t work on a string but on count of rune-pairs as you just need the count of runes at the end.

## Day 15

Puuuuh, dijkstra is the solution.
After extending the graph to the Part2 dimension my solution for Part1 is now wrong. I guess there is a path outside of the Part1 boundary which is now shorter than the correct answer.
Assumption was right, I introduced a multiply parameter and now the solution is correct.

## Day 16

Ouch.... misinterpreted the padding stuff (used it on packet level..) and lots of classic "off by one" problems.
Finally made part 1 and I couldn't care less for part 2 right now