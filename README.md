# dagger-dash-days
Dagger repo for dash days

_Goal's_
- We want capability providers/developers to focus on what they're building - not on how to build/deploy/test/ATO/security-scan/all-the-things it
- We want capability providers/developers to focus on what they're building AND make it easy to leverage Defense Unicorns ecosystem
  - don't have a zarf.yaml - no problem
  - don't have a Dockerfile - no problem
  - only have src code - no problem
  - ...
- We (Unicorns) want to deploy anything, anywhere. This can only be accomplished with some opinionation. This project is an effort to standarize and opinionate the BEST way to get source code to prod in/for highly regulated environments.


## TODOs

TODO:
- [x] go library and cli
- [x] inst  zarf.yaml
- [ ] detect zarf.yaml ? init zarf.yaml : be happy
- [ ] zarf package create
- [ ] Do it in a github pipeline
  - [ ] `cp dagger-dash-days/examples/github/main.yaml <pod-info>/.github/workflows/main.yaml`
- [ ] Do it in a gitlab pipeline
  - [ ] `cp dagger-dash-days/examples/gitlab/main.yaml <pod-info>/.gitlab-ci.yaml`

TODO - Extra Credit and General Questions:
- [ ] Use Ironbank images
- [ ] can we define GH workflow and make examples/github/main.yaml stupid simple?
  - [ ] ? https://docs.github.com/en/actions/using-workflows/creating-starter-workflows-for-your-organization
    ```yaml
    name: Example GitHub Workflow to leverage Dagger Pipeline Templates

    runs:
        steps:
            - name: all-things
            uses: defenseunicorns/dagger-pipeline-templates@main
    ```

## What Pain are we solving?

### Capability Provider creates a new project

Goal: Release my capability
FROM THIS....
```
1. init repo (src/, Dockerfile, Makefile, etc...)
2. pipelines
   1. docker lint stage
   2. trufflehog (password scanning)
   3. build src
   4. lint
   5. unit tests
   6. dependecny check
   7. build image/s
   8. neuvector scan
   9. build zarf package
   10. publish zarf package
   11. deploy staging
   12. run integration tests
   13. release
```

TO THIS...
```
#github
1. init repo (src/, Dockerfile, Makefile)
2. goto dpt/examples/github/main.yaml
   1. copy past to .github/workflows/main.yaml
---
#gitlab
1. init repo (src/, Dockerfile, Makefile)
2. goto dpt/examples/gitlab/main.yaml
   1. copy paste to .gitlab-ci.yaml
```

OR MAYBE EVEN THIS...
```
Given Software Factory
1. Interact with a UI to create a new project
```


