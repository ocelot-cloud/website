---
title: "Contribution Guide"
tags: ["contributing", "community"]
---

### Report bugs

If you find a bug, please open an issue on GitHub with a clear title and detailed description. Include the steps to reproduce the bug, the expected result, and the actual result. If possible, include screenshots or logs. Tag the issue as "bug".

### Feature proposals

If you have an idea for improving the project, you can submit your suggestion as a GitHub issue. Describe the feature, its benefits, and if possible, how it could be implemented. Label the issue as a "feature proposal".

### Create a pull request

To contribute code, fork the repository, create a new branch, make your changes, and then submit a pull request (PR). Branches should have a name of the following form `<issueNumber>-<descriptiveTitle>`. Make sure your PR has a clear title, a brief description of the changes, and any relevant issue numbers. All code contributions should pass existing tests and include new tests for added features.

## Developer Guide

### Technological Decisions

This project is designed to make development as easy and accessible as possible. Notable decisions in this regard are:

- The **ci-runner** component is a tool written in Go that is capable of performing many CI operations such as build, test, deploy, and specifically provides very fine-grained access to the test suite according to current your development needs.
- **Easy Tools**: Vue.js and Cypress are intuitive tools for web development and writing acceptance tests. Both technologies are configured to use Typescript mainly because of its type safety. The backend and ci-runner language is Go, which is easy to learn and, pleasant to develop, and its executables have low resource requirements.
- The **CI pipelines**, or "Actions" in GitHub parlance, execute the test suite via ci-runner for new commits pushed to GitHub. This enforces frequent quality assurance and early feedback for developers.

**Installation Requirements**

- Go
- Node.js
- Cypress Dependencies
- Docker and Docker Compose v2 (not v1 "docker-compose")
- Add these lines in the `/etc/hosts` file:

```
127.0.0.1       sampleapp.localhost
127.0.0.1       ocelot-cloud.localhost
# add further lines as needed during development
```

You may need to add more addresses as tests are executed against them in the test suite. Browsers seem to automatically map `.localhost` addresses to IP 127.0.0.1, but the applications we build do not. Hence the customization of the `hosts` file.

### Automatic Development Preparation

There is a script that can install all the required development dependencies for the project on Debian but may also work for other versions or Ubuntu. The instructions for execution are:

```bash
git clone https://github.com/ocelot-cloud/ocelot-cloud
cd ocelot-cloud/scripts
sudo bash install.sh
# Reboot computer
cd ocelot-cloud/src/ci-runner
go build && ./ci-runner test all
```

If you work on another system, the script may still be helpful for guiding you through installing the dependencies.

### ci-runner

The ci-runner is designed to make the testing process as easy as possible. Here is a quick start that prints the available options:

```bash
cd src/ci-runner
go build && ./ci-runner --help
```

There are a lot of options. Feel free to experiment with them to get familiar.

### Development Mode

To facilitate the development of interactions between the frontend and the backend, the 'native' mode has been introduced. This is a special setup in which the frontend and backend run separately, but natively on the system, allowing Vue.js to run its own process that enables hot loading of source code changes, reducing development time. Open a console and run one of these commands to start the setup in development mode:

```bash
cd src/cloud/backend           && bash run-native-mocked.sh
cd src/cloud/frontend          && bash run-native.sh
cd src/cloud/acceptance-tests  && bash run-native.sh
```

You can now visit the Vue.js GUI at `http://localhost:8081/` and use the Cypress GUI to run acceptance tests if needed.

## Development Policies

### Lean Technology Stack

Stick to the languages Go and Typescript and the tools Vue.js, Cypress and Docker for all development to keep knowledge requirements for this project as low as possible. New technologies should only be introduced when there is a compelling need. Implement DevOps related code in the ci-runner tool which is written in Go. New languages or bash scripts are not allowed.

### Testing Strategy

Only frequently and extensively tested features are stable features. And in this project, only stable features may be merged. Every feature you add should be checked by passing tests. All tests must be included in the ci-runner in a job that is frequently executed by GitHub Actions.

What should be tested?

- 'happy path', if the execution of the production code flow works as expected
- 'unhappy path' tests, where invalid input, unexpected things, or errors occur during execution. Certain entities to be created already exist, or the operation requires entities that have not yet exited, etc. The code must be able to handle such errors correctly to be robust in production.
- border cases such as nil/null, zero, negative numbers, empty strings, empty collections, etc.
- Security and its testing must be built in from the beginning, where necessary. For example, a new REST API endpoint requires authentication, input validation etc.

Depending on the situation, one or more of the following types of tests may be necessary:

- Unit tests (testing a single unit) or core tests (wiring several units together and testing by interacting with a high-level API).
- Component tests, i.e. running the backend and interacting with it via REST API.
- Acceptance tests must be written using Cypress for all GUI features. Acceptance tests are run only in native mode and against a mocked backend to avoid error-prone waits.

### Coding Practices

- The extensive test suite makes it unlikely that code changes will introduce new bugs. Therefore, any kind of code improvement is encouraged. If you accidentally break something, the tests will tell you, so you can fix it right away. Feel free to add any tests that may be missing.
- Contributions must be refactored until they are clean code, i.e.
  - readable: easy to understand
  - maintainable: easy to apply future changes, e.g. by avoiding duplications
- 3rd party libraries must be wrapped.
