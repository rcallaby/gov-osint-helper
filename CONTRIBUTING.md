# Contributing to gov-osint-helper

Thank you for your interest in contributing to **gov-osint-helper**. This project is an open-source tool designed to query and aggregate results from multiple government databases. We welcome contributions that improve functionality, security, documentation, and usability.

Please review the following guidelines before submitting contributions.

---

## How to Contribute

### 1. Reporting Issues

* Use the [GitHub Issues](../../issues) tracker for bug reports, feature requests, and questions.
* Clearly describe the problem, steps to reproduce, and expected behavior.
* For security-related issues, **do not open a public issue**. Instead, please email the maintainers directly at:
  `security@gov-osint-helper.org`

### 2. Branching and Pull Requests

* Work from the latest `main` branch.
* Create a feature branch for your work:

  ```bash
  git checkout -b feature/your-feature-name
  ```
* Keep commits atomic and write clear commit messages.
* Rebase your branch on `main` before opening a pull request:

  ```bash
  git fetch origin
  git rebase origin/main
  ```
* Open a pull request (PR) against `main` with:

  * A descriptive title.
  * A clear summary of changes.
  * Reference to related issues, if applicable.

### 3. Code Style

* Follow [PEP 8](https://peps.python.org/pep-0008/) (for Python code).
* Use docstrings for public functions and modules.
* Ensure all code is type-annotated where possible.
* Avoid hardcoding values; prefer configuration files or environment variables.

### 4. Testing

* Add or update unit tests for all changes.
* Place tests in the `tests/` directory.
* Ensure tests pass before submitting a PR:

  ```bash
  pytest
  ```
* Contributions without adequate test coverage may not be accepted.

### 5. Documentation

* Update `README.md` or create/update documentation under the `docs/` folder when introducing new features.
* Use clear, concise language.

---

## Commit Message Guidelines

* Use present tense (“Add feature” not “Added feature”).
* Limit subject line to 72 characters.
* Reference issues/tickets in the footer when applicable.

Example:

```
Add database query abstraction layer

- Implemented modular query interface
- Supports additional government APIs
- Includes unit tests for validation

Closes #42
```

---

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this standard.

---

## Licensing

By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).



