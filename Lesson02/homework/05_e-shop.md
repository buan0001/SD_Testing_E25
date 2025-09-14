### E-shop
You are testing the payment functionality of an e-shop. The system receives a positive amount of purchases in kroner with an accuracy of 1 øre. Based on this value, a discount is calculated according to the following rules:

|Amount|Discount|
|-|-:|
|Up to 300 kr|0%|
|Over 300 kr, up to 800 kr|5%|
|Over 800 kr|10%|

Use black-box analysis to identify a comprehensive series of test cases:
1. Identify the corresponding equivalence partitions and propose test cases based on them
- Partitions: MIN - (-1), 0, 0.01 - 300, 300.01 - 800, 800.01 - MAX
- Values: -20, 0, 150, 600, 1500
2. Use 3-value boundary value analysis to identify further test cases
- Boundary values: -2, -1, 0.01, 1, 2, 299, 300, 301, 799, 800, 801
3. Write down the full resulting list of test cases
- Combined list: -20, -2, -1, 0, 0.01, 1, 2, 150, 299, 300, 301, 600, 799, 800, 801, 1500
4. Implement the discount calculation function in code and write the corresponding unit tests in the language and unit test framework of your choice

<sub>Adapted from Stapp, Lucjan, Roman, Adam, and Michaël Pilaeten (2024). _ISTQB Certified Tester Foundation Level: A Self-Study Guide Syllabus v4.0_. Springer.</sub>