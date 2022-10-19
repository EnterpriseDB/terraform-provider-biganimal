# Contributing

You are encouraged to contribute to the project via GitHub pull requests. This document describes guidelines to support you in submitting your contributions.

## Intro

We welcome contributions to code, open defects, improvements and  documentation. All members of the community, including contributors, are expected to uphold the [Code of Conduct](./CODE_OF_CONDUCT.md).

## Issues

[GitHub issues](https://github.com/istio/istio/issues/new/choose) can be used to report suspected defects or submit feature requests.

When reporting a misbehaviour please make sure to provide:

-  the version of the Provider you were using (e.g. version number,
  or git commit);

- the Operating system version you are using;

- the expected VS obtained result;

- the minimal steps needed to reproduce the issue 


## Proposing a Feature

We recommend to go through the following steps when wishing to propose a feature.

- Discuss your idea in Github discussions
- Once there is general agreement that the feature is useful, create a GitHub issue to follow up. The issue should cover why the feature is needed, scope, use cases and any other considerations/limitattions
- If the scope requires, you'll be asked to submit a design, along with technical and implementation details
- Once the major technical issues are resolved and agreed upon, post a note to summarise design decision and the general execution plan
- Submit a PR with your proposed code change. Please make sure to cover documentation for your feature, including usage examples when possible

Please favour small PRs instead over giant scary PRs. We therfore suggest splitting large features into a set of progressive PRs that build on top of one another.

If you would like to skip the process of submitting an issue and instead would prefer to just submit a pull request with your desired code changes then that's fine. But keep in mind that there is no guarantee of it being accepted and so it is usually best to get agreement on the idea/design before time is spent coding it. However, sometimes seeing the exact code change can help focus discussions, so the choice is up to you.

## How to make a Pull request

For existing issues simply respond to the issue and express interest in working on it. This communication is important in order to  prevent duplicated efforts. If your proposed fix or feature is not already covered by an issue you may consider opening one first.

To submit your proposed change:

- clone the affected repository,
- create a new branch for your changes called "dev/ISSUE_ID",
- submit a pull request (PR)

## PR Guidelines

A PR is more likely to be accepted if it has:

- a well described PR body (TODO link)
- a good commit message. We use [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/). (TODO link)
- code that follows the conventions in old code
- code that respects [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- good enough unit and regression testing (see below)
- all checks green for static code analysis and automatic tests


## Expected Level of Testing

Your change should include unit tests and regression tests to protect your feature or other features from breaking. Please make sure

- that tests follow the conventions of existing ones;
- for new features, to cover as many inputs/configurations as possible.


## Certificate of Origin

By contributing to this project you agree to the [Developer Certificate of Origin (DCO)](https://developercertificate.org/). This document was created by the Linux Kernel community and is a simple statement that you, as a contributor, have the legal right to make the contribution. By signing off your commits, you guarantee that you wrote the patch or otherwise have the right to contribute the material by the rules of the DCO.

A signed commit includes the following line:
```
Signed-off-by: Jane Doe <jane.doe@example.com>
```
Please use your real name (sorry, no pseudonyms or anonymous contributions).
If you set your `user.name` and `user.email` in your Git Config, you can sign your
commit automatically with `git commit -s`.