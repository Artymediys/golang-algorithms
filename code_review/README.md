# Task description
Some company has `n` junior developers and `m` senior developers.
In order for the code written by a junior developer to get into the prod,
it must be checked by at least `k` senior developers. 
The senior developer spends `1` minute to check one program.
Each of the `n` junior developers has written a program which they want to send to the prod.
Unfortunately, the checking system is not distributed yet,
so two senior developers cannot check the same job at the same time.
Also, a senior-developer wants to do a really good code-review,
so he does not check two programs at the same time and is not distracted by other things during the check.
Determine how much time in minutes it will take all the senior developers
to review all the programs written by the junior developers.
### Input format:
A single input line contains the numbers `n`, `m`, `k` `(1 <= k <= m <= n <= 10^4)`,
where `n` is the number of junior developers, `m` is the number of senior developers,
`k` is the number of checks required per job.
### Output data format:
Output one number - the total time, in minutes, that the senior developers will spend checking.