Flowers

A mail-order company selling flower seeds charges 30 kr for postage and packing on all orders up to and including 150 kr value and 40 kr for orders above 150 kr value and up to and including 300 kr value. For orders above 300 kr value there is no charge for postage and packing.

If you were using equivalence partitioning to prepare test cases for the postage and packing charges:

    What valid input partitions might you use?
    What test case values could be inferred from said partitions?
    With the information provided, would there be any invalid partitions?

_____________________________________________________________________________________

What valid input partitions might you use? (kr)
- 0 - 150
- 151 - 300
- 300 - Inf

What test case values could be inferred from said partitions?
- (-1)
- 0
- 75.123
- 150.0
- 151
- 300
- MAX_FLOAT

With the information provided, would there be any invalid partitions?
- MIN_FLOAT < 0