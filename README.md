# Puppet

Puppet is a service just for being controlled.

## Usage

you can run a puppet by docker:

```
docker run --name puppet pigeonligh/puppet
```

and then you can control it:

```
docker exec puppet ctr sleep 10s
# or
docker exec puppet ctr exit 1
```

## Why puppet?

Puppet can be used to simulate task and change status of task. For example, you can run puppet in Kubernetes Job and observe the actions of Job Controller.
