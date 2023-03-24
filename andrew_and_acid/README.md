# Task description

Andrew works in a secret chemical laboratory, which produces a dangerous acid with amazing properties.
Andrew has `n` infinitely large tanks arranged in one row.
Initially, each tank contains a certain amount of acid.
Andrei's superiors require that all the tanks contain the same amount of acid.
Unfortunately, the pouring machine is imperfect.
In one operation it can pour one liter of acid into each of the first `k (1 <= k <= n)` reservoirs.
Note that for different operations `k` may be different. Since the acid is very expensive,
Andrew is not allowed to pour the acid out of the tanks.
Andrew asks you to find out if it is possible to equalize the volumes of acid in the tanks
and if it is possible, calculate the minimum number of operations that would be needed to achieve this.

Input format
The first line contains the number `n (1 <= n <= 100 000)` - the number of acid tanks.
The second line contains `n` integers `a_i (1 <= a_i <= 10^9)`,
where `a_i` means the original volume of acid in the `i-th` tank in liters.

Output format
If the volumes of acid in the tanks can be equalized, output the minimum number of operations necessary for this.
If this is not possible, output "-1".
