### Airline discount policy
An airline offers only flights to India and Asia. Under special conditions, a discount is offered on the normal airfare:
- Passengers older than 18 with destinations in India are offered a discount of 20%, as long as the departure is not on a Monday or Friday
- For destinations outside of India, passengers are offered a discount of 25%, if the departure is not on a Monday or Friday
- Passengers who stay at least 6 days at their destination receive an additional discount of 10%
- Passengers older than 2 but younger than 18 years are offered a discount of 40% for all destinations
- Children 2 and under travel for free

Apply black-box test design:
1. Represent this information in a decision table.
2. Write the corresponding unit tests (one test case per business rule) using the programming language and unit test framework of your choice

<sub>Adapted from FlexRule, ["Preparing a decision table"](https://resource.flexrule.com/knowledge-base/preparing-a-decision-table/)</sub>

#### Solution

1. Decision table
   
||R1|R2|R3|R4|R5|R6|R7|R8|R9|R10|R11|R12|R13|R14|R15|R16|R17|R18|R19|R20|R21|R22|R23|R24|
|-|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|**Conditions**|||||||||||||||||||||||||
|Destination|I|I|I|I|I|I|I|I|I|I|I|I|A|A|A|A|A|A|A|A|A|A|A|A|
|Passenger age|<=2|<=2|<=2|<=2|3-18|3-18|3-18|3-18|>18|>18|>18|>18|<=2|<=2|<=2|<=2|3-18|3-18|3-18|3-18|>18|>18|>18|>18|
|Depart Mon/Fri|T|T|F|F|T|T|F|F|T|T|F|F|T|T|F|F|T|T|F|F|T|T|F|F|
|Stay>=6 days|T|F|T|F|T|F|T|F|T|F|T|F|T|F|T|F|T|F|T|F|T|F|T|F|
|**Actions**|||||||||||||||||||||||||
|No discount|N|N|N|N|N|N|N|N|N|Y|N|N|N|N|N|N|N|N|N|N|N|Y|N|N|
|10% discount|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|Y|N|
|20% discount|N|N|N|N|N|N|N|N|N|N|Y|Y|N|N|N|N|N|N|N|N|N|N|N|N|
|25% discount|N|N|N|N|N|N|N|N|N|N|N|N|N|N|Y|Y|N|N|Y|Y|N|N|Y|Y|
|40% discount|N|N|N|N|Y|Y|Y|Y|N|N|N|N|N|N|N|N|Y|Y|Y|Y|N|N|N|N|
|Travel free|Y|Y|Y|Y|N|N|N|N|N|N|N|N|Y|Y|Y|Y|N|N|N|N|N|N|N|N|

Reduced solution:

||R1|R2|R3|R4|R5|R6|R7|R8|R9|R10|R11|R12|R13|
|-|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|**Conditions**||||||||||||||
|Destination|-|I|I|-|-|I|I|A|A|A|A|A|A|
|Passenger age|<=2|3-18|3-18|>18|>18|>18|>18|3-18|3-18|3-18|3-18|>18|>18|
|Depart Mon/Fri|-|-|-|T|T|F|F|T|T|F|F|F|F|
|Stay>=6 days|-|T|F|F|T|T|F|T|F|T|F|T|F|
|**Actions**||||||||||||||
|No discount|N|N|N|Y|N|N|N|N|N|N|N|N|N|
|10% discount|N|Y|N|N|Y|Y|N|Y|N|Y|N|Y|N|
|20% discount|N|N|N|N|N|Y|Y|N|N|N|N|N|N|
|25% discount|N|N|N|N|N|N|N|N|N|Y|Y|Y|Y|
|40% discount|N|Y|Y|N|N|N|N|Y|Y|Y|Y|N|N|
|Travel free|Y|N|N|N|N|N|N|N|N|N|N|N|N|

2. Code solution

- [Pytest/Python](https://github.com/arturomorarioja/py_airline_unit_tests)
- [Jest/JavaScript](https://github.com/arturomorarioja/js_airline_unit_tests)
- [PHPUnit/PHP8](https://github.com/arturomorarioja/php_airline_unit_tests)
