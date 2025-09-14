### Framing shop

The system calculates the price of picture framing based on the given parameters: width and height of the picture (in centimeters). The valid width of the picture is between 30 and 100 cm inclusive. The valid height of the picture is between 30 and 60 cm inclusive. The system calculates the area of the image as the product of width and height. If the surface area exceeds 1600 cm<sup>2</sup>, the framing price is 3500 kr. Otherwise, the framing price is 3000 kr.

Use black-box analysis to identify a comprehensive series of test cases:

1. Identify the corresponding equivalence partitions and propose test cases based on them

- Width:
  - Part: MIN - (-1), 0, 1 - 29, 30 - 100, 101 - MAX
  - Test values: -50, 0, 15, 50, 150
- Height:
  - Partitions: MIN - (-1), 0, 1 - 29, 30 - 60, 61 - MAX
  - Test values: Same as above

2. Use 3-value boundary value analysis to identify further test cases

- Shared: -2, -1, 1, 2, 28, 29, 30, 31
  - Height: 59, 60, 61, 62
  - Width: 99, 100, 101, 102

3. Write down the full resulting list of test case values

- Shared: -50, -2, -1, 0, 1, 2, 15, 28, 29, 30, 31, 40x40, 50, 150
  - plus above for boundary height/width values

4. Implement the discount calculation function in code and write the corresponding unit tests in the language and unit test framework of your choice

<sub>Adapted from Stapp, Lucjan, Roman, Adam, and Michaël Pilaeten (2024). _ISTQB Certified Tester Foundation Level: A Self-Study Guide Syllabus v4.0_. Springer.</sub>
