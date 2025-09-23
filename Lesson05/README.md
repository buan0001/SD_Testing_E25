[Testing - Autumn 2025](https://github.com/arturomorarioja-kea/SD_Testing_E25/blob/main/README.md)

# Lesson 5 - 23 September

## Exercise solutions
- [Coverage](https://github.com/arturomorarioja-ek/SD_Testing_E25/edit/main/Lesson04/01%20Coverage.md) (Number and Employees)
- Decision Testing
  - [Airline](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson02/09%20DT%20Airline.md)

- State Transition Testing
  - [ATM](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson03/03%20ST%20ATM.md)
  - [Login](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson03/04%20ST%20Login)

## Class takeaways
Check out the following slide decks on Itslearning:
- **Integration Testing**, with especial attention to
  - Advantages: protection against regressions, resistance to refactoring
  - Disadvantages: slow, difficult to maintain
- **API Testing**. Focus on:
  - How do API calls usually fail?
  - What to test for?
  - Postman (although you can use Insomnia, ThunderClient or any other similar platform)
- **Database Testing**
- **Continuous Testing**. Notice:
  - The difference between CI, CT, CD and the other CD

## Homework

- Integration testing
  - Solve the [measure converter exercise](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson05/01%20Measure%20converter.md), where you will:
    - Apply your black-box and white-box test design knowledge
    - Decide what to mock and what not to mock (and, most importantly, why)
- API testing
  - Check out the [library API v3](https://github.com/arturomorarioja/py_library_api_v3) Postman tests
    - [Collection](https://github.com/arturomorarioja/py_library_api_v3/blob/main/postman/Library%20API%20v3.postman_collection.json)
    - [Environment](https://github.com/arturomorarioja/py_library_api_v3/blob/main/postman/Library%20API%20v3.postman_environment.json)
  - Practice API testing in existing APIs of yours:
    - Create collections to group requests to the same API and environments to define variables
    - Write tests under "Scripts". You can use snippets and the built-in AI tool
    - Remember to write positive and negative tests
    - Sort your tests so that you can run them in a row
  - Solve the [customers API exercise](https://github.com/arturomorarioja-ek/SD_Testing_E25/blob/main/Lesson05/02%20Customers%20API.md)
- Continuous Testing
  - Try one or several CI/CD applications on applications of yours
    - Create continuous integration jobs and pipelines
    - Run tests in the pipeline (unit tests, integration tests, linting, static code analysis tools)
