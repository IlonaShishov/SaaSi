# deployer
A Golang CLI tool to deploy configurations extracted from a live OpenShit/Kubernetes environment from the [exporter](../exporter/README.md) tool.

## Dependencies
* Use [kustomize](konveyor.io/tools/crane/) to export the original configuration and remove cluster specific settings
  (e.g. IP addresses, status, ...)

## Installer configuration
The `deployer` runs using a configuration to specify how to install the packaged configuration:
* For each namespace we can define a target namespace name, reuse the original one or apply a common strategy to transform
the initial name into the new one
* For all the mandatory and optional parameters defined in `ConfigMaps`, we can define the desired value
* For all the parameters defined as 1`Secret` keys, we must provide the desired value

```yaml
application:
  # This creates an installer package named APP
  name: APP
  # Optional, defines the automatic mapping for the target namespaces
  # %s will be replaced with the original namespace name
  namespaceMappingFormat: "%s-prod"
  namespaces:
  - name: NS1
    # If missing, the namespaceMappingFormat or the original name are applied 
    target: NEW-NS1
    params:
     - configMap: MAP-1
        params:
        - name: PARAM-1
          value: VALUE-1
        - name: "..."
          value: "..."
     - configMap: MAP-2
         params:
        - name: PARAM-1
          value: VALUE-1
       - name: "..."
         value: "..."
    secrets:
      - secret: SECRET-1
          params:
          - name: PARAM-1
            value: VALUE-1
          - name: "..."
            value: "..."
      - secret: SECRET-2
          params:
          - name: PARAM-1
            value: VALUE-1
          - name: "..."
            value: "..."
```

## Running the installer
Prerequisites:
* Install `oc` and `kustomize` CLI tools
* Install `go` CLI tool (at least version 1.19)
* `oc` login to the source OpenShift cluster

Run this command to install the package extracted in `./installer/MYAPP` from the given configuration `myapp.yaml`:
```bash
go run main.go myapp.yaml ./installer/MYAPP
```

**TODO**:
* Manage mandatory params
  * If __DEFAULT__ stop installation

## Open points
* Which user permissions are needed to install
