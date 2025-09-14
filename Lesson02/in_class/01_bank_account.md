Bank account

Suppose you have a bank account that offers variable interest rates:

    0.5 per cent for the first 10000 kr credit
    1 per cent for the next 10000 kr
    1.5 per cent for the rest

If you wanted to check that the bank was handling your account correctly:

    What input partitions might you use?
    What test case values could be inferred from said partitions?

______________________________________________________________________________________

What input partitions might you use?
- (-Inf) - (-1)
- 0 - 10000
- 10001 - 20000
- 20000 - Inf


What test case values could be inferred from said partitions?
- MIN_FLOAT64
- (-1)
- 0
- 5000
- 10000
- 10001
- 15000
- 20000
- 20001
- MAX_FLOAT64