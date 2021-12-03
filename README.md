# Daily Coding Problem: Problem #1056 [Medium]
## The "Look and Say" Sequence

This problem was asked by Epic.

The "look and say" sequence is defined as follows:
beginning with the term 1,
each subsequent term visually describes the digits appearing in the previous term.
The first few terms are as follows:

```
1
11
21
1211
111221
```

As an example, the fourth term is 1211, since the third term consists of one 2 and one 1.

Given an integer `N`, print the `Nth` term of this sequence.

## Analysis

I made an [erroneous program](las.go) that looks through the previous sequence value first,
then "says" digits and counts in digit-numerical order.
This sequence converges on a value:

```
1       1
2       11
3       21
4       1112
5       3112
6       211213
7       312213
8       212223
9       114213
10      31121314
11      41122314
12      31221324
13      21322314
14      21322314
```

I found this mistaken version amusing: it has fixed point(s).

If you keep the sequence values as 64-bit int,
the algorithm overflows on the 10th value in the sequence:

```sh
$ go build las2.go
$ ./las2 10
1       1
2       11
3       21
4       1211
5       111221
6       312211
7       13112221
8       1113213211
9       31131211131221
2021/12/02 17:40:57 strconv.Atoi: parsing "13211311123113112211": value out of range
```

The exact error comes from [my code](las2.go) using a string-conversion,
but any code keeping the sequence as an integer will overflow very rapidly.

I wrote a [third version](las3.go) that works with string representations
of sequence values, never converting to an integer internally.

```sh
$ go build las3.go
$  ./las3 12
1       1
2       11
3       21
4       1211
5       111221
6       312211
7       13112221
8       1113213211
9       31131211131221
10      13211311123113112211
11      11131221133112132113212221
12      3113112221232112111312211312113211
```

One follow-on here is to add the ability of changing the starting string.

## Interview Analysis

This is a good interview problem.
It explains what it wants,
the candidate doesn't have to have any physical insight,
or do 3-D visualizations.
There's no single best algorithm to accomplish the task.
The candidate can design a somewhat tricky flow-of-control,
and has a choice of how to keep sequence values.
The numerical overflow is a way for the interviewer to
gauge how a candidate responds to initial assumptions
that might be false or cause problems.
Just enough code results that the candidate can choose variable
names wisely.

I'm not sure this rates a "[Medium]" in light of other problems
the Daily Coding Problem has rated "[Easy]".
