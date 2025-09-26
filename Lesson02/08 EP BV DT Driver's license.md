### Driver's license
The operator of the driver's license test support system enters the following information into the system, for a candidate who is taking the exams for the first time:
- The number of points from the theoretical exam (integer number from 0 to 100)
- The number of errors made by the candidate during the practical exam (integer number 0 or greater)
The candidate must take both exams. A candidate is granted a driver's license if they meet the following two conditions: they scored at least 85 points on the theoretical test and made no more than two errors on the practical test. If a candidate fails one of the exams, they must repeat this exam. In addition, if the candidate fails both exams, they are required to take additional hours of driving lessons.

Use black-box analysis to identify a comprehensive series of test cases:
1. Create the corresponding decision table
2. Write test cases based on the decision table
3. Identify the corresponding equivalence partitions
4. Use 3-value boundary value analysis to identify further test cases
5. Identify edge cases
6. List all test case values
7. Implement in code a function that receives as parameters the number of points for the theory exam and the number of errors for the practical exam and that returns a data structure with four boolean properties: whether the driver's license is granted, whether the theory exam must be repeated, whether the practical exam must be repeated, and whether additional driving lessons must be taken. Write the corresponding unit tests based on the above analysis. Use the programming language and unit test framework of your choice

<sub>Adapted from Stapp, Lucjan, Roman, Adam, and Michaël Pilaeten (2024). _ISTQB Certified Tester Foundation Level: A Self-Study Guide Syllabus v4.0_. Springer.</sub>

#### Solution
1. Decision table

||Rule 1|Rule 2|Rule 3|Rule 4|
|-|:-:|:-:|:-:|:-:|
|**Conditions**|
|Points >= 85|T|T|F|F|
|Errors <= 2|T|F|T|F|
|**Actions**|
|Grant a driver's license|Y|N|N|N|
|Repeat theory exam|N|N|Y|Y|
|Repeat practical exam|N|Y|N|Y|
|Additional driving lessons|N|N|N|Y|
   
2. Decision table-based test cases

||Th ex points|Pr ex errors|Rule|
|-|--:|--:|-|
|TC1|90|1|Rule 1|
|TC2|90|5|Rule 2|
|TC3|50|1|Rule 3|
|TC4|50|5|Rule 4|

3. Equivalence partitions

||Partition types|Partitions|Test case values|
|-|-|--:|--:|
|Theory exam points|Valid|0-84|50|
||Valid|85-100|90|
||Invalid|101-MAX INTEGER|120|
|Practical exam errors|Valid|0-2|1|
||Valid|3-MAX INTEGER|5|

The test case values for valid partitions were already figured out after the decision table.

4. 3-value boundary value-based test cases

||Partitions|Boundary values|Test case values|
|-|-|--:|--:|
|Theory exam points|0-84|0|0 1|
||0-84|84|83 84 85|
||85-100|85|84 85 86|
||85-100|100|99 100 101|
|Practical exam errors|0-2|0|0 1|
||0-2|2|1 2 3|
||3-MAX INTEGER|3|2 3 4|

Unique test case values:
- Theory exam points: `0` `1` `83` `84` `85` `86` `99` `100` `101`
- Practical exam errors: `0` `1` `2` `3` `4`

5. Edge cases
- Theory exam points: `-1` `101` `MAX INTEGER` `MAX INTEGER + 1` `wrong data types, if applicable`
- Practical exam errors: `-1` `MAX INTEGER` `MAX INTEGER + 1` `wrong data types, if applicable`

6. Complete list of test case values
- Theory exam points: `-1` `0` `1` `50` `83` `84` `85` `86` `90` `99` `100` `101` `120` `MAX INTEGER` `MAX INTEGER + 1` `wrong data types, if applicable`
- Practical exam errors: `-1` `0` `1` `2` `3` `4` `5` `MAX INTEGER` `MAX INTEGER + 1` `wrong data types, if applicable`

7. Code solution

- [Pytest/Python](https://github.com/arturomorarioja/py_drivers_license_unit_tests)
- [Jest/JavaScript](https://github.com/arturomorarioja/js_drivers_license_unit_tests)
- [PHPUnit/PHP8](https://github.com/arturomorarioja/php_drivers_license_unit_tests)
