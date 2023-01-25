# ❎ Semangit
[![Master Workflow](https://github.com/emranprojects/semangit/actions/workflows/master.yml/badge.svg)](https://github.com/emranprojects/semangit/actions/workflows/master.yml)
![Test Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/emranbm/03d07927044bdfe73aee59e6193dd8d5/raw/badge-coverage-semangit.json)  
A simple tool to force version update in CI.  
We've all experienced commits forgetting to update the corresponding version on the codebase; like changing a helm chart template without increasing the `version` in `Chart.yaml`. *Semangit* helps to prevent such mistakes by checking whether a change needs a corresponding version increase.

# 📖 How to use
The _Semangit_ project can be used in various ways. The possible ways are listed below.

## 📃 Compiled binary
The compiled version of the project can be found in [releases page](https://github.com/divar-ir/semangit/releases), you can simply download the artifact and run it as a command:
```bash
./semangit run [OPTIONS]
```

## 🐳 Official image
_Semangit_ also comes with an official docker image. The image can be found in [here](https://hub.docker.com/r/divaar/semangit). A sample usage is as follows:
```dockerfile
FROM divaar/semangit:0.1.5

ENTRYPOINT   semangit run [OPTIONS]
```

# 📦 Plugins
_Semangit_ is founded based on plugins and each of them can have some options to be set, to see the details of all options run `semangit --help` command. Below are some plugins which are currently supported in _Semangit_ project. Feel free to [contribute](#-contribution) and expand the number of version checkers.

## ❎ Helm chart version checker
This plugin will force the Helm chart version update on CI if either the package files or the `values.yaml` file is changed.

### 📖 Usage
A sample usage of this plugin in CI template can be as bellow:
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

# 💡 Contribution
Fortunately the plugins of _Semangit_ project are written in a way that you don't have to get in the details of the project. Simply, just create a file in [versionanalyzers](https://github.com/divar-ir/semangit/tree/master/internal/models/versionanalyzers) directory and implement the [VersionAnalyzer](https://github.com/divar-ir/semangit/blob/9443ec422e425166de83269289c4a3ec3b22cd52/internal/models/version_analyzer.go#L3) interface. _Semangit_ will do the rest and register your plugin and after that, the new plugin can be used through the executing options:

```bash
semangit run 
--repo-dir .
--old-rev master
--new-rev HEAD
--log-level debug
--version-analyzer-name ${YOUR_DESIRED_PLUGIN_NAME} 
--{EXTRA-ARG-NAME} {EXTRA-ARG-VALUE}
```