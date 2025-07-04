### Project Awareness & Context

- **Always read `PLANNING.md`** (if available) at the start of a new conversation to understand the project's architecture, goals, style, and constraints.
- **Use consistent naming conventions, file structure, and architecture patterns** as described in `PLANNING.md`, or by following best practices and conventions for the technologies used if `PLANNING.md` is not available.

### Code Structure & Modularity

- **Never create a file longer than 500 lines of code.** If a file approaches this limit, refactor by splitting it into modules or helper files.
- **Organize code into clearly separated modules**, grouped by feature or responsibility. Always adhere to the single responsibility principle.
- **Always adhere to industry best practices**, especially paying close attention to the SOLID principles for code, and the FIRST principles for unit tests.

### Testing & Reliability

- **Always create unit tests for new features** (functions, methods, etc).
- **After updating any logic**, check whether existing unit tests need to be updated. If so, do it.
- **Always create at least the following tests:**
    - 1 test for expected use
    - 1 edge case
    - 1 failure case

### Style & Conventions

- **Use the technologies mentioned by the user**, and only move on to other technologies if completely necessary, while still prompting the user for confirmation.
- **Follow best practices and idiomatic code** for whatever technology you use.
- When it comes to inline comments, **avoid using them to talk to the user**. **Never use inline comments to talk about *WHAT* the code is doing**, but rather use them sparingly and only when completely necessary, and focus on the *WHY* instead**.

### Documentation & Explainability

- **Update `README.md`** when new features are added, dependencies change, or setup steps are modified.
- Document the code and it's constructs (modules, functions, classes, interfaces) with detailed but not overly descriptive information.

### Security & Error handling

- **Always make sure that the code you generate or update is secure** and follows best security practices and is protected against common security vulnerabilities/attacks.
- **Dont assume that because a feature is simple it does not need to be secure**.
- **Always manage errors as grecefully as possible**, and create custom errors and/or exceptions to better represent the possible issues.

### AI Behavior Rules
- **Never assume missing context. Ask questions if uncertain.**
- **Never hallucinate libraries or functions** – only use known, verified packages (use the context7 MCP toolset if available to search for information about them if needed).
- **Always confirm file paths and modules** exist before referencing them in code or tests. 