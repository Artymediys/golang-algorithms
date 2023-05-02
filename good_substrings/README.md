# Task description

Given a string `s`, consisting of characters `a`, `b`, `c`, `d`.
A substring of `s` is a string that can be obtained by taking some consecutive characters from `s`.
For example, strings `bcd`, `abcdcdab`, `cdcdab` are substring of string `abcdcdab`, but `cc` and `cdcdcd` are not.
Call a string good if each of the characters `a`, `b`, `c`, `d` occurs at least once in it.
Find the length of the shortest good substring of the string `s` or determine that s has no good substring.
### Format of the input data:
The first line of input data contains a number `n` `(1 <= n <= 2*10^5)`, where `n` is the length of string `s`.
The second line of input data contains the string `s` itself, consisting of characters `a`, `b`, `c`, `d`.
### Output data format:
Output the length of the shortest good substring of the string `s`. If there are no good substrings, print `-1`.