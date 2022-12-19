# Semangit
[![Master Workflow](https://github.com/emranprojects/semangit/actions/workflows/master.yml/badge.svg)](https://github.com/emranprojects/semangit/actions/workflows/master.yml)
![Test Coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/emranbm/03d07927044bdfe73aee59e6193dd8d5/raw/badge-coverage-semangit.json)  
A simple tool to force version update in CI.  
We've all experienced commits forgetting to update the corresponding version on the codebase; like changing a helm chart template without increasing the `version` in `Chart.yaml`. *Semangit* helps to prevent such mistakes by checking whether a change needs a corresponding version increase.

# Use cases
- Force Helm package version update on CI, if the package files are changed.
