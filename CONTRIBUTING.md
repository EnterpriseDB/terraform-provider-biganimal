# Contributing

- [Intro](#intro)
  - [Report an Issue](#issues)
  - [Find an Issue](#find-an-issue)
  - [Proposing a Feature](#proposing-a-feature)
  - [Developer Environment Setup](#developer-environment-setup)
  - [How to Make a Pull Request](#how-to-make-a-pull-request)
  - [Commit and PR Guidelines](#commit-and-pr-guidelines)
  - [Expected Level of Testing](#expected-level-of-testing)
  - [Certificate of Origin](#certificate-of-origin)

  

You are encouraged to contribute to the project by forking the project and submitting [GitHub pull requests](https://github.com/EnterpriseDB/terraform-provider-biganimal/pulls). This document describes guidelines to support you in submitting your contributions.

## Intro

We welcome contributions to code, open defects, improvements and  documentation. All members of the community, including contributors, are expected to uphold the [Code of Conduct](./CODE_OF_CONDUCT.md).

## Report an Issue

[GitHub issues](https://github.com/istio/istio/issues/new/choose) can be used to report suspected defects or submit feature requests.

When reporting a misbehaviour please make sure to provide:

-  the version of the Provider you were using (e.g. version number,
  or git commit);
- the Operating system version you are using;
- the expected VS obtained result;
- the minimal steps needed to reproduce the issue

## Find an Issue

We have good first issues for new contributors and help wanted issues suitable
for any contributor. [good first issue](TODO) has extra information to
help you make your first contribution. [help wanted](TODO) are issues
suitable for someone who isn't a core maintainer and is good to move onto after
your first pull request.

Sometimes there won’t be any issues with these labels. That’s ok! There is
likely still something for you to work on. If you want to contribute but you
don’t know where to start or can't find a suitable issue, you can ⚠️ **explain how people can ask for an issue to work on**.

Once you see an issue that you'd like to work on, please post a comment saying
that you want to work on it. Something like "I want to work on this" is fine.

## Proposing a Feature

We recommend to go through the following steps when wishing to propose a feature.

- Discuss your idea in Github discussions
- Once there is general agreement that the feature is useful, create a GitHub issue to follow up. The issue should cover why the feature is needed, scope, use cases and any other considerations/limitattions
- If the scope requires, you'll be asked to submit a design, along with technical and implementation details
- Once the major technical issues are resolved and agreed upon, post a note to summarise design decision and the general execution plan
- Submit a PR with your proposed code change. Please make sure to cover documentation for your feature, including usage examples when possible

Please favour small PRs instead over giant scary PRs. We therefore suggest splitting large features into a set of progressive PRs that build on top of one another.

If you want to skip submitting an issue and instead prefer to just send a pull request with your proposed code change, we'll still consider it. However, please be aware that it may not be accepted: it is recommended to get agreement on the idea/design before spending time coding it. We know that, sometimes, sharing directly the code change can help focus discussions, so we leave you the option.

## Developer Environment Setup

(TODO) Explain how to:

- Get the source code
- Retrieve any dependencies
- Build the source code
- Run the project locally
- Test the source code, unit and “integration” or “end-to-end”
- Generate and preview the documentation locally

## How to make a Pull Request

For existing issues simply respond to the issue and express interest in working on it. This communication is important in order to  prevent duplicated efforts. If your proposed fix or feature is not already covered by an issue you may consider opening one first.

To submit your proposed change:

- clone the affected repository,
- create a new branch for your changes called "dev/ISSUE_ID",
- commit your code,
- submit a pull request (PR)

## Commit and PR Guidelines

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