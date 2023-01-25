# 🩺 Semangit
[![Master Workflow](https://github.com/emranprojects/semangit/actions/workflows/master.yml/badge.svg)](https://github.com/emranprojects/semangit/actions/workflows/master.yml)
![Test Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/emranbm/03d07927044bdfe73aee59e6193dd8d5/raw/badge-coverage-semangit.json)  
A simple tool to force version update in CI.  
We've all experienced commits forgetting to update the corresponding version on the codebase; like changing a helm chart template without increasing the `version` in `Chart.yaml`. *Semangit* helps to prevent such mistakes by checking whether a change needs a corresponding version increase.

# 📖 Quick start guide
By now, the _Semangit_ project only support _helm_ version checker. To use it on your CI, use one of the following templates:

## Gitlab
Sample usage of the _helm_ version checker in Gitlab CI is shown below:
```yaml
check-helm-version:
  stage: lint
  image: divaar/semangit:latest
  variables:
    GIT_DEPTH: 0
    GIT_STRATEGY: clone
    SEMANGIT_REPODIR: .
    SEMANGIT_OLDREVISION: ${CI_DEFAULT_BRANCH}
    SEMANGIT_NEWREVISION: ${CI_COMMIT_REF_NAME}
    SEMANGIT_HELMROOTDIR: .
    SEMANGIT_LOGLEVEL: info # Options: trace, debug, info, warn, error, fatal, panic
  script:
    - git branch ${SEMANGIT_OLDREVISION} origin/${SEMANGIT_OLDREVISION}
    - git branch ${SEMANGIT_NEWREVISION} origin/${SEMANGIT_NEWREVISION}
    - semangit run
  only:
    refs:
      - merge_requests # Override this section if you don't want this behaviour
  before_script: []
```
This plugin will fail the CI job if either the template files or the `values.yaml` file is changed.

## 📃 Compiled binary
The compiled version of the project can be found in [releases page](https://github.com/divar-ir/semangit/releases), you can simply download the artifact and run it as a command:
```bash
./semangit run [flags]
```

## 🐳 Official image
_Semangit_ also comes with an official docker image. The image can be found in [here](https://hub.docker.com/r/divaar/semangit). A sample usage is as follows:
```bash
docker run --rm divaar/semangit:0.1.5 semangit run [flags]
```

# 💡 Contribution
Fortunately, the plugins of _Semangit_ project are written in a way that you don't have to get into the details of the project. Simply, just create a file in [versionanalyzers](https://github.com/divar-ir/semangit/tree/master/internal/models/versionanalyzers) directory and implement the [VersionAnalyzer](https://github.com/divar-ir/semangit/blob/master/internal/models/version_analyzer.go) interface. _Semangit_ will do the rest and register your plugin and after that, the new plugin can be used through the executing options:

```bash
semangit run 
--repo-dir .
--old-rev master
--new-rev HEAD
--log-level debug
--version-analyzer-name ${YOUR_DESIRED_PLUGIN_NAME} 
--{PLUGIN_NAME}-{EXTRA-ARG-NAME} {EXTRA-ARG-VALUE}
```