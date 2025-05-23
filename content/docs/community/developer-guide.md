---
title: "Developer Guide"
---

### Technology Choices

This project favors popular, well-supported tools to keep onboarding friction low.

* **GitHub** – central hub for collaboration, issues, and code reviews.
* **Docker** – build and run the entire stack with a single dependency. You can also build your own images if you prefer not to use those published on Docker Hub.
* **Vue 3 + TypeScript** – straightforward, type-safe frontend development.
* **Cypress + TypeScript** – acceptance testing with an intuitive UI.
* **Go** – language for backend and **ci-runner**, a CLI that can build, test, and deploy with fine-grained control.
* **GitHub Actions** – every push triggers the **ci-runner** pipeline, providing fast feedback and enforcing quality gates.

### Installation Requirements

There is a script that can install all the required development dependencies for the project on Ubuntu, but may also work for other versions or Debian. The user executing this script must have 'sudo' rights. The instructions for execution are:

```bash
git clone https://github.com/ocelot-cloud/cloud
cd ocelot-cloud/scripts
bash install-project-dependencies.sh
# Reboot computer
cd ocelot-cloud/components/ci-runner
go build && ./ci-runner test all -i
```

### Extend 'hosts' File

Add this to the `/etc/hosts` file:

```
127.0.0.1       ocelot-cloud.localhost
127.0.0.1       sampleapp.localhost
```

This is necessary for the component tests, the HTTP client in Go tries to look up domains like `*.localhost` in a DNS server instead of sending the request directly to 127.0.0.1.

### ci-runner

The ci-runner is designed to make the testing process as easy as possible. Here is a quick start:

```bash
cd src/ci-runner
go build && ./ci-runner -h
go build && ./ci-runner test -h
```

There are a lot of options. Feel free to experiment with them to get familiar.

### Development Mode

If you want to run component tests against the backend REST API, run the backend with mocked dependencies:

```bash
cd src/backend && bash run-native-mocked.sh
```

If you want to test the frontend via cypress, run these two scripts in addition to the backend:

```bash
cd src/frontend && bash run-native.sh
cd src/cypress && bash run-native.sh
```

You can now visit the Vue.js GUI at `http://localhost:8081/` and use the Cypress GUI to run acceptance tests if needed.

## Development Policies

### Lean Technology Stack

Stick to the languages Go and Typescript and the tools Vue.js, Cypress and Docker for all development to keep knowledge requirements for this project as low as possible. New technologies should only be introduced when there is a compelling need. Implement DevOps related code in the ci-runner tool which is written in Go and avoid bash scripts.

### Rigorous Testing Strategy

Only frequently and extensively tested features are stable features. And only stable features are allowed to be merged. Every feature you add should be checked by passing tests. All tests must be included in the "test all" ci-runner job that is frequently executed by GitHub actions.

What should be tested?

* 'happy path', if the execution of the production code flow works as expected
* 'unhappy path' tests, when unexpected things or errors happen during execution. The production code must be robust, i.e. able to handle errors correctly. For example: "can't create user as it already exists", "app does not exist", etc.

Depending on the situation, consider using one or more of the following test types:

* Unit tests (testing a single unit) or integration test (wiring several units together and testing by interacting with a high-level API).
* Component tests, i.e. running the backend and sending http requests to it for testing
* Acceptance tests with Cypress for checking GUI features