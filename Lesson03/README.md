[Testing - Autumn 2025](https://github.com/arturomorarioja-kea/SD_Testing_E25/blob/main/README.md)

# Lesson 3 - 9 September

[-> 5 State transition testing + in-class exercise]: #

## Exercise solutions
- [Password field](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/04%20EP%20BV%20Password%20field.md)
- [E-shop](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/05%20EP%20BV%20E-shop.md)
- [Framing shop](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/06%20EP%20BV%20Framing%20shop.md)

## In-class exercises
  - Decision table testing
    - [Input form](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/07%20DT%20Input%20form.md)
      
[  - State Transition Diagrams]: #
[    - Checkout(https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson03/02%20ST%20Checkout.md)]: #

## Class takeaways
Check out the following slide decks on Itslearning:

- **Test Design Techniques - Black-box**, with especial attention to:
  - Decision tables
    
[  - State Transition Diagrams]: #
- **Introduction to Unit Testing**, specifically:
  - Test Doubles    
- **Unit Testing Best Practices and Anti-Patterns**
  - Best practices
    - Each test must verify only one behaviour
    - Test case selection should be comprehensive
    - Full regression testing should be run as often as possible
    - Code must be written so that it is testable (e.g., pure functions or methods)
  - Anti-patterns
    - Do never test private methods directly
    - Do never expose private state
    - Do not leak domain knowledge to the unit tests
    - Avoid code pollution

- **Unit Testing Approaches**, focusing on the following concepts
  - Private, shared and volatile dependencies
  - The Classical Approach to Unit Testing
    - Broad unit tests
    - Mocking only shared dependencies
    - Unit test isolation rather than code under test isolation
  - The London Approach to Unit Testing
    - Small unit tests
    - Everything is mocked
    - Code under test isolation (one unit test for each function/method)
  - Shall external dependencies be mocked (Khorikov) or not (Wassell)?

## Homework
- Reflect on all the above
- Check out the following code sample on unit test mocking:
  - Customer onboarding ([Pytest/Python](https://github.com/arturomorarioja/py_customer_onboarding_mock) | [Jest/JavaScript](https://github.com/arturomorarioja/js_customer_onboarding_mock) | [PHPUnit/PHP](https://github.com/arturomorarioja/php_customer_onboarding_mock))
- Practice the use of test doubles (mocks and stubs) in the unit testing framework(s) of your choice
- Solve the [employees exercise](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson03/01%20Employees.md)
  - Try to follow the specification to the letter
  - Testing dates might be problematic. Give it some thought
- Solve the following black-box test design exercises:
  - Decision table testing
    - [Driver's license](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/08%20EP%20BV%20DT%20Driver's%20license.md). It also involves equivalence partitions and boundary values
    - [Airline](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/09%20DT%20Airline.md)
    
[  - State Transition Diagrams]: #
[    - ATM(https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson03/03%20ST%20ATM.md)]: #
[    - Login(https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson03/04%20ST%20Login)]: #
