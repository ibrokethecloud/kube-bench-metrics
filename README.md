## kube-bench-metrics

kube-bench-metrics is a wrapper to execute aqusecurity/kube-bench.

It parses the results of the scan and exposes the results as prometheus metrics.

These can now be scraped by prometheus with subsequent alerting via Alertmanager.

The helm chart available in the **charts** directory can be used to quickly get started.

The `versionOverride` variable in values.yaml needs to be updated to specify the correct
version of checks to be run from **cfg** directory.

The current versions in **cfg** are a copy of ```github.com/rancher/security-scan``` and are customised for RKE environments.

These can be overriden. I will add support for all upstream kube-bench versions as well.

To Do:
* Build automation
* Prometheus integration in helm chart
