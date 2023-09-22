# dagger-dash-days

Dagger Pipeline Templates (`dpt`) is a cli tool that can be used local and in Git[Hub/Lab] pipelines alike.

## Goal

> Create a portable pipeline that could enable Developers to focus on what they're building. 
> 
> NOT on how to build/deploy/test/ATO/security-scan/all-the-things

We want to enable developers to leverage the Defense Unicorns ecosystem without becoming experts in it. 

## Quick Start

1. Copy github workflow to your project
```
cp dagger-dash-days/examples/github/main.yaml <YOUR-PROJECT>/.github/workflows/main.yaml
```

2. That's it... you're done. Just go write code.

### Okay, but really

All of the work is done in `dpt` the github workflow is just a simple wrapper to leverage it. The same would be true for a `gitlab-ci` or a `jenkinsfile`

If a project doesn't have a `zarf.yaml` file, `dpt` will let the developer know exactly what needs to happen... (in the future will create a `zarf.yaml` for them)

The Vision: dpt would have the BEST defaults for all the things. So that a developer can just write code and not worry about the rest. The pipeline stages would fail, but provide clear indicators and help to the developers on what needs to be done to fix it.

## What Pain are we solving?

```mermaid
graph TD
    style A fill:##2196F3,stroke:#1565C0,stroke-width:5px
    style B fill:##2196F3,stroke:#1976D2,stroke-width:5px
    style C fill:#1f8f1b,stroke:#1976D2,stroke-width:5px
    style D fill:#1f8f1b,stroke:#1976D2,stroke-width:5px
    style E fill:##a13310,stroke:#E64A19,stroke-width:5px
    style F fill:##a13310,stroke:#E64A19,stroke-width:5px
    style G fill:##a13310,stroke:#E64A19,stroke-width:5px
    style H fill:##a13310,stroke:#E64A19,stroke-width:5px
    style I fill:##a13310,stroke:#E64A19,stroke-width:5px
    style J fill:##a13310,stroke:#E64A19,stroke-width:5px
    style K fill:##a13310,stroke:#E64A19,stroke-width:5px
    style L fill:##a13310,stroke:#E64A19,stroke-width:5px
    style M fill:##a13310,stroke:#E64A19,stroke-width:5px
    style N fill:##a13310,stroke:#E64A19,stroke-width:5px
    style O fill:##a13310,stroke:#E64A19,stroke-width:5px
    style P fill:##a13310,stroke:#E64A19,stroke-width:5px
    
    A[Developer Goal: release my capability]
    B[init repo: src/, Dockerfile, Makefile]
    C[use dpt copy/paste .github/workflows/main.yaml]
    D[write codezzzz]

    E[docker lint stage]
    F[trufflehog: password scanning]
    G[build src]
    H[lint]
    I[unit tests]
    J[dependency check]
    K[build image/s]
    L[neuvector scan]
    M[build zarf package]
    N[publish zarf package]
    O[deploy staging]
    P[run integration tests]
    choose{suck life \n or \n not suck life}

    A --First--> B
    B --> choose
    choose --not suck life--> C
    choose --suck life--> E
    C --> D

    E --still sucks--> F
    F --still sucks--> G
    G --still sucks--> H
    H --still sucks--> I
    I --still sucks--> J
    J --still sucks--> K
    K --still sucks--> L
    L --still sucks--> M
    M --still sucks--> N
    N --still sucks--> O
    O --still sucks--> P
    P --> D
```


## Just try it

```
cd test/pod-info && go run ../../main.go package create
```

## TODOs

- [x] go library and cli
- [x] inst  zarf.yaml
- [ ] detect zarf.yaml ? init zarf.yaml : be happy
- [x] zarf package create
- [x] Do it in a github pipeline
  - [x] `cp dagger-dash-days/examples/github/main.yaml <pod-info>/.github/workflows/main.yaml`
- [ ] Do it in a gitlab pipeline
  - [ ] `cp dagger-dash-days/examples/gitlab/main.yaml <pod-info>/.gitlab-ci.yaml`
- [ ] expand use cases, maybe integrate buildpacks???
- [ ] Use Ironbank images
- [ ] Use Zarf registry (runtime)
- [ ] Publish Zarf Package
- [ ] Deploy Package (easy user configuration for target environment)
- [ ] Make the CLI better (better output, logging, help, etc....)