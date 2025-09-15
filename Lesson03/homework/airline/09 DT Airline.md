### Airline discount policy

An airline offers only flights to India and Asia. Under special conditions, a discount is offered on the normal airfare:

- Passengers older than 18 with destinations in India are offered a discount of 20%, as long as the departure is not on a Monday or Friday
- For destinations outside of India, passengers are offered a discount of 25%, if the departure is not on a Monday or Friday
- Passengers who stay at least 6 days at their destination receive an additional discount of 10%
- Passengers older than 2 but younger than 18 years are offered a discount of 40% for all destinations
- Children 2 and under travel for free

Apply black-box test design:

1. Represent this information in a decision table.

|                           | R1  | R2  | R3  | R4  | R5  | R6  | R7  | R8  | R9  | R10 | R11 |
| ------------------------- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **Conditions**            |     |     |     |     |     |     |     |     |     |     |     |
| Destination outside India | -   | -   | T   | T   | F   | F   | -   | F   | T   | T   | -   |
| Monday or Friday          | T   | T   | F   | F   | F   | F   | T   | F   | F   | F   | -   |
| Minimum 6 day stay        | T   | F   | T   | F   | T   | F   | T   | F   | T   | F   | -   |
| Older than 18             | T   | T   | T   | T   | T   | T   | N/A | N/A | N/A | N/A | N/A |
| Age between 2 and 18      | N/A | N/A | N/A | N/A | N/A | N/A | T   | T   | T   | T   | N/A |
| Age under 2               | N/A | N/A | N/A | N/A | N/A | N/A | N/A | N/A | N/A | N/A | T   |
| **Actions**               |     |     |     |     |     |     |     |     |     |     |     |
| Discount %                | 10  | 0   | 35  | 25  | 30  | 20  | 50  | 40  | 75  | 65  | 100 |

2. Write the corresponding unit tests (one test case per business rule) using the programming language and unit test framework of your choice

<sub>Adapted from FlexRule, ["Preparing a decision table"](https://resource.flexrule.com/knowledge-base/preparing-a-decision-table/)</sub>
