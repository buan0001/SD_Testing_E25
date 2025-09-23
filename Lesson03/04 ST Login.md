### Login
Draw a State Transition Diagram for an application login:

- A correct attempt logs the user in
- 3 failed attempts block the account

#### Solution

<img width="659" height="457" alt="image" src="https://github.com/user-attachments/assets/4c33cd23-c9fa-4bf2-bd86-637f508890c6" />

Alternative implementations:

- Showing an "account blocked" state
- Creating one transition for "1st pwd wrong" and another one for "2nd pwd wrong" (could be unmanageable if the number of attempts is big)
