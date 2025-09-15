### Driver's license

The operator of the driver's license test support system enters the following information into the system, for a candidate who is taking the exams for the first time:

- The number of points from the theoretical exam (integer number from 0 to 100)
- The number of errors made by the candidate during the practical exam (integer number 0 or greater)
  The candidate must take both exams. A candidate is granted a driver's license if they meet the following two conditions: they scored at least 85 points on the theoretical test and made no more than two errors on the practical test. If a candidate fails one of the exams, they must repeat this exam. In addition, if the candidate fails both exams, they are required to take additional hours of driving lessons.

Use black-box analysis to identify a comprehensive series of test cases:

1. Create the corresponding decision table

|                         | R1  | R2  | R3  | R4  |
| ----------------------- | --- | --- | --- | --- |
| **Conditions**          |     |     |     |     |
| At least 85 points      | F   | T   | T   | F   |
| No more than 2 mistakes | T   | F   | T   | F   |
| **Actions**             |     |     |     |     |
| Pass                    | N   | N   | Y   | N   |
| Redo practical          | N   | Y   | N   | Y   |
| Redo theoretical        | Y   | N   | N   | Y   |
| Must take more lessons  | N   | N   | N   | Y   |

2. Write test cases based on the decision table
3. Identify the corresponding equivalence partitions

- Points: 0-84, 85-100, 100-INF
- Mistakes: 0-2, 3-INF

4. Use 3-value boundary value analysis to identify further test cases
- Points: 0, 1, 83, 84, 85, 86, 99, 100,
- Mistakes: 0, 1, 2, 3, 4
5. Identify edge cases
- Points: (above max) 101, 102, (negative) -1, -2
- Mistakes: (negative) -2, -1 
6. List all test case values
- Points: -2, -1, 0, 1, 83, 84, 85, 86, 99, 100, 101, 102
- Mistakes: -2, -1, 0, 1, 2, 3, 4
7. Implement in code a function that receives as parameters the number of points for the theory exam and the number of errors for the practical exam and that returns a data structure with four boolean properties: whether the driver's license is granted, whether the theory exam must be repeated, whether the practical exam must be repeated, and whether additional driving lessons must be taken. Write the corresponding unit tests based on the above analysis. Use the programming language and unit test framework of your choice

<sub>Adapted from Stapp, Lucjan, Roman, Adam, and Michaël Pilaeten (2024). _ISTQB Certified Tester Foundation Level: A Self-Study Guide Syllabus v4.0_. Springer.</sub>
