[![Circle CI](https://circleci.com/gh/malnick/specd/tree/master.svg?style=svg&circle-token=184482507ec4fafa04056d1ce9a0c11c97df8ff6)](https://circleci.com/gh/malnick/specd/tree/master)

***DISCLAIMER*** This project is WIP. Current foucs is on developing the HTTP API and reporting center for the tool. Future work will be implementing enforcing state for resources declared in the state.yaml.

# specD  
A lightweight server inspection daemon and HTTP API implemented in Go

## Overview
specd is a lightweight system spec reporting daemon implemented in go. specd’s philosophy is to provide basic primitives, a modern HTTP API for reportingg in a small micro service you can deploy on limited linux architectures. 

### Command Line Interface
```
NAME:
   specd - A lightweight configuration management utility

USAGE:
   specd [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
    run Start specd

GLOBAL OPTIONS:
   --verbose                    Verbose logging mode.
   --json-log                   JSON logging mode.
   --api-port "1015"            API port when running in API mode.
   --state, -s "./state.yaml"   Path to state.yaml [$specd_STATE_PATH]
   --help, -h                   show help
   --version, -v                print the version
```
#### Commands
##### run
###### report
Report the state of the resources provisioned in the state.yaml.
###### api
Start the HTTP API interface. 

#### Flags
##### -v --verbose
Verbose logging output. 

##### --json-logger
Log directly to JSON 


##### -ha --http-api
Expose a HTTP/S API for dynamic interaction and integration.

#### state.yaml
Declare complete system state in the state.yaml:

```yaml
---
# specd configuration
bootstrap: /home/centos/my_bootstrap.tar.gz
provider: redhat

# Host Properties
properties: 
  
  ipaddress: |
#!/bin/bash
ip addr eth0 
  
  custom_uname: |
#!/bin/bash
uname -a

# Template Dictionary
dictionary:
  arbitrary_variable: foobar

# Resource Declaration
package:
  docker:
    exists: true

files:
  name: /etc/systemd/system/docker.service.d/override.conf: 
    exists: true
    before: 
      - package::docker
    content: |
ExecPreStart=/bin/echo "My IP addr is {{ property::ipaddress }} and {{ dictionary::arbitrary_variable }}"
ExecStart=/usr/bin/docker daemon --storage-driver=overlay -H fd://

  /etc/yum.repos.d/docker.repo:
    exists: true
    mode: 0644
    after:
      - package::docker
    content: |
[dockerrepo]
name=Docker Repository
baseurl=https://yum.dockerproject.org/repo/main/centos/$releasever/
enabled=1
gpgcheck=1
gpgkey=https://yum.dockerproject.org/gpg

service:
  docker.service:
    exists: true
    after: 
      - package::docker
      - file::my_file
```

#### Resource Primatives
Resources primitives are inherited from goss. The primitives are slightly modified to accept `before` and `after` parameters which allow specd to construct a directed acyclic graph for enforcing state. Currently specd supports <n> resource primitives from goss: <supported primatvies list> 

### Configuration
#### Environment Variables

`specd_HOME`: Defaults . 
`specd_CONFIG`: Defaults ./state.yaml

### Reporting
#### Report Endpoints
Get the state of resources on the host, in JSON format, from the HTTP API or the CLI.

Given a state.yaml which declares:

```yaml
---
files:
  /tmp/foobar.txt:
    exists: true
    mode: 0644
```

Executing `specd run report` returns:

```json
[
        {
                "successful": true,
                "resource-id": "/tmp/foobar.txt",
                "resource-type": "File",
                "title": "/tmp/foobar.txt",
                "meta": null,
                "test-type": 0,
                "property": "exists",
                "err": null,
                "expected": [
                        "true"
                ],
                "found": [
                        "true"
                ],
                "human": "",
                "duration": 29270
        },
        {
                "successful": false,
                "resource-id": "/tmp/foobar.txt",
                "resource-type": "File",
                "title": "/tmp/foobar.txt",
                "meta": null,
                "test-type": 0,
                "property": "mode",
                "err": null,
                "expected": [
                        "420"
                ],
                "found": [
                        "\"0644\""
                ],
                "human": "Expected\n    \u003cstring\u003e: 0644\nto equal\n    \u003cint\u003e: 420",
                "duration": 44931
        }
]
```

`curl -XGET localhost:1015/specd/api/v1/report`:

```json
{
  ... same as `run report` ... 
}
```

### DAG Implementation
The directed graph is implemented with breadth first search by indegree on the resources declared in the state.yaml. This search checks for cyclic resource declaration and duplicate edges on insert. 

### HTTP API
#### API Endpoints
##### /specd/api/v1/
##### /specd/api/v1/run/
Reserved for mimicing the CLI experience as an HTTP API.

##### /specd/api/v1/run/report/
Methods: GET POST

GET report, must have state.yaml on box in expected location. 
POST report, accepts a POST with JSON of resources to return report data about. 

REQUEST: **GET** 

RETURNS:

```json
[
        {
                "successful": true,
                "resource-id": "/tmp/foobar.txt",
                "resource-type": "File",
                "title": "/tmp/foobar.txt",
                "meta": null,
                "test-type": 0,
                "property": "exists",
                "err": null,
                "expected": [
                        "true"
                ],
                "found": [
                        "true"
                ],
                "human": "",
                "duration": 29270
        },
        {
                "successful": false,
                "resource-id": "/tmp/foobar.txt",
                "resource-type": "File",
                "title": "/tmp/foobar.txt",
                "meta": null,
                "test-type": 0,
                "property": "mode",
                "err": null,
                "expected": [
                        "420"
                ],
                "found": [
                        "\"0644\""
                ],
                "human": "Expected\n    \u003cstring\u003e: 0644\nto equal\n    \u003cint\u003e: 420",
                "duration": 44931
        }
]
```

REQUEST: **POST**

```json
{
  "files": {
    "/tmp/foobar.txt": {
      "exists": true
    }
  }
}
```
RETURNS:

```json
[
        {
                "successful": true,
                "resource-id": "/tmp/foobar.txt",
                "resource-type": "File",
                "title": "/tmp/foobar.txt",
                "meta": null,
                "test-type": 0,
                "property": "exists",
                "err": null,
                "expected": [
                        "true"
                ],
                "found": [
                        "true"
                ],
                "human": "",
                "duration": 29270
        },
        {
                "successful": false,
                "resource-id": "/tmp/foobar.txt",
                "resource-type": "File",
                "title": "/tmp/foobar.txt",
                "meta": null,
                "test-type": 0,
                "property": "mode",
                "err": null,
                "expected": [
                        "420"
                ],
                "found": [
                        "\"0644\""
                ],
                "human": "Expected\n    \u003cstring\u003e: 0644\nto equal\n    \u003cint\u003e: 420",
                "duration": 44931
        }
]
```

##### /specd/api/v1/run/once/
Allows for overriding on the fly state.json

Request POST
```json
{
  state.json
}
```

Response RunReport{
```json
...
  "dag" {
    ...
  }
...
```
}
