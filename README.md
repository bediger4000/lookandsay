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

## Interview Analysis

This is a good interview problem.
It explains what it wants,
the candidate doesn't have to have any physical insight,
or do 3-D mental visualizations.
There's no single best algorithm to accomplish the task.
The candidate can design a somewhat tricky flow-of-control,
and has a choice of what type to use for sequence values.
The numerical overflow is a way for the interviewer to
gauge how a candidate responds to initial assumptions
that might be false or cause problems.
Just enough code results that the candidate can choose variable
names wisely.

This problem can even be extended by the interviewer,
or the candidate can suggest features:

* The ability to use a different number as the start of the sequence
* Why the elements of the sequence never get a digit numerically larger than 3
* Is it possible to generate the Nth term of the sequence without creating
the preceding N-1 terms?

I'm not sure this rates a "[Medium]" in light of other problems
the Daily Coding Problem has rated "[Easy]".

### Digits in elements of the sequence

As far as why the elements of the sequence never get a digit numerically larger than 3,
I think it goes like this:

The 5th value of the sequence is 111221.
That illustrates that saying some sequence values ends up as "N Ns, N Ms".
That's 3 Ns in a row in the next sequence value, the value you're currently saying.
The way you "say" the digits of the next term in the sequence, N and M are different,
so the only way to get four of a digit in a row in a term,
is for the preceding term to have 11 '1' digits,
followed by a single other digit.
But it's not possible to have 11 '1' digits in a row in the preceding term
because of the "N Ns, N Ms" convention.

### Finding the Nth term without finding the preceding terms

```
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
13      1321132132111213122112311311222113111221131221
14      11131221131211131231121113112221121321132132211331222113112211
15      311311222113111231131112132112311321322112111312211312111322212311322113212221
```

There's a clear pattern in the final 4 digits of even-numbered terms (2211)
vs the final 3 digits of the odd-numbered terms (221).
I can't work out any more of the pattern, if there is one.
