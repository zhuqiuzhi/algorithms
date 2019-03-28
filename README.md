# algorithms

[![Build Status](https://travis-ci.com/zhuqiuzhi/algorithms.svg?branch=master)](https://travis-ci.org/zhuqiuzhi/algorithms)
[![codecov](https://codecov.io/gh/zhuqiuzhi/algorithms/branch/master/graph/badge.svg)](https://codecov.io/gh/zhuqiuzhi/algorithms)
[![Go Report Card](https://goreportcard.com/badge/github.com/zhuqiuzhi/algorithms)](https://goreportcard.com/report/github.com/zhuqiuzhi/algorithms)
[![GoDoc](https://godoc.org/github.com/zhuqiuzhi/algorithms?status.svg)](https://godoc.org/github.com/zhuqiuzhi/algorithms)
[![Sourcegraph](https://sourcegraph.com/github.com/zhuqiuzhi/algorithms/-/badge.svg)](https://sourcegraph.com/github.com/zhuqiuzhi/algorithms)

algorithms is solutions to algorithmic puzzles in the programming language Go.

Table of Contents
=================

* [Strategy](#strategy)
* [Iteration](#iteration)
* [Recursion](#recursion)
* [Exhaustive Search or Brute Force](#exhaustive-search-or-brute-force)
* [Backtracking](#backtracking)
* [Heuristics](#heuristics)
* [Divide and Conquer](#divide-and-conquer)
* [Dynamic Programming](#dynamic-programming)
* [Branch and Bound](#branch-and-bound)
* [Using specialized data structure](#using-specialized-data-structure)

## Strategy 

[Computer science distilled](https://code.energy/computer-science-distilled/) 总结了以下算法设计策略

- Handle repetitive tasks through **iteration**,
- Iterate elegantly using **recursion**,
- Use **brute force** when you’re lazy but powerful,
- Test bad options and then **backtrack**,
- Save time with **heuristics** for a reasonable way out,
- **Divide and conquer** your toughest opponents,
- Identify old issues **dynamically** not to waste energy again,
- **Bound** your problem so the solution doesn’t escape.

## Iteration

The iterative strategy consists in using loops (e.g. for, while) to repeat a process until a condition is met. Each step in a loop is called an iteration. It’s great for running through an input and applying the same operations on every part of it.

Puzzle |  Solution | Playgroud | Demo 
------ |  -------- | --------- | ----
[Merge sea and fresh into a sorted list](https://code.energy/computer-science-distilled/) | [mergeSortedList.go](/mergeSortedList.go) | |

## Recursion 

We say there’s recursion when a function delegates work to clones of itself. A recursive algorithm will naturally come to mind for solving a problem defined in terms of itself.

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----
[Selection Sort](https://en.wikipedia.org/wiki/Selection_sort) | [selectionSort.go](/selectionSort.go) | |

## Exhaustive Search or Brute Force

An exhaustive search, or brute force, algorithm examines every possible alternative to find one particular solution.

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----
[Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem) | [knapsack.go](/knapsack.go) | | 

## Backtracking

Backtracking works best in problems where the solution is a se- quence of choices and making a choice restrains subsequent choices. It identifies as soon as possible the choices you’ve made cannot give you the solution you want, so you can sooner step back and try something else. Fail early, fail often.

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----
[Eight queens puzzle](https://en.wikipedia.org/wiki/Eight_queens_puzzle) | [queens.go](/queens.go) | [Leetcode](https://leetcode-cn.com/problems/n-queens/) | [The 8 Queens Problem](https://code.energy/8-queens-problem)

## Heuristics

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----

## Divide and Conquer

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----

## Dynamic Programming

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----
[Fibonacci number](https://en.wikipedia.org/wiki/Fibonacci_number) | [fibonacci.go](/fibonacci.go) | [Fibonacci Number](https://leetcode.com/problems/fibonacci-number/) |

## Branch and Bound

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----

## Using specialized data structure

Puzzle | Solution | Playgroud | Demo 
------ | -------- | --------- | ----
[第一个只出现一次的字符](https://book.douban.com/subject/6966465/) | [firstNotRepeatingChar.go](/firstNotRepeatingChar.go) | |
[Two Sum](https://leetcode-cn.com/problems/two-sum/) | [two-sum.go](/two-sum.go) | [Leetcode](https://leetcode-cn.com/problems/two-sum/) | 
