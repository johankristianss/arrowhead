# Introduction

## Create a car provider 
```console
mkdir carprovider;cd carprovider
```

```console
arrowhead systems register --name carprovider --address localhost --port 8880
```
```console
INFO[0000] Using openssl as certificate manager
Certificate request self-signature ok
subject=CN = carprovider.cloud1.ltu.arrowhead.eu
INFO[0000] System registered successfully, certificate stored in ./carprovider.p12 and ./carprovider.pub, config file stored in ./carprovider.env
```

The command will register a new system at the Arrowhead Service Registry core service, and generate a corresponding PKCS#12 certificate with CN *carprovider.cloud1.ltu.arrowhead.eu*.

```console
ls
```
```console
.rw------- 4.2k johan 15 Feb 09:04 -N  carprovider.p12
.rw-r--r--  451 johan 15 Feb 09:04 -N  carprovider.pub
```

Type the command below to list available systems:
```console
arrowhead systems ls
```
```console
╭─────┬─────────────────┬────────────────────┬──────╮
│ ID  │ SYSTEM NAME     │ ADDRESS            │ PORT │
├─────┼─────────────────┼────────────────────┼──────┤
│ 1   │ serviceregistry │ c1-serviceregistry │ 8443 │
│ 2   │ gateway         │ c1-gateway         │ 8453 │
│ 3   │ eventhandler    │ c1-eventhandler    │ 8455 │
│ 4   │ orchestrator    │ c1-orchestrator    │ 8441 │
│ 5   │ authorization   │ c1-authorization   │ 8445 │
│ 6   │ gatekeeper      │ c1-gatekeeper      │ 8449 │
│ 103 │ carprovider     │ localhost          │ 8880 │
╰─────┴─────────────────┴────────────────────┴──────╯
```

We can also filter to only list the car systems.
```console
arrowhead systems ls --filter car
```
```console
╭─────┬─────────────────┬───────────┬──────╮
│ ID  │ SYSTEM NAME     │ ADDRESS   │ PORT │
├─────┼─────────────────┼───────────┼──────┤
│ 103 │ carprovider     │ localhost │ 8880 │
╰─────┴─────────────────┴───────────┴──────╯
```

## Create a car consumer 
```console
cd ..;mkdir carconsumer;cd carconsumer
```
```console
arrowhead systems register --name carconsumer --address localhost --port 8881
```
```console
INFO[0000] Using openssl as certificate manager
INFO[0000] System registered successfully, PKCS#12 certificate stored in ./carconsumer.p12 and ./carconsumer.pub, config file stored in ./carconsumer.env
```

```console
arrowhead systems ls --filter car
```
```console
╭─────┬─────────────────┬───────────┬──────╮
│ ID  │ SYSTEM NAME     │ ADDRESS   │ PORT │
├─────┼─────────────────┼───────────┼──────┤
│ 103 │ carprovider     │ localhost │ 8880 │
│ 104 │ carconsumer     │ localhost │ 8881 │
╰─────┴─────────────────┴───────────┴──────╯
```

## Register services
First let's register a function to create cars.
```console
arrowhead services register --system carprovider --definition create-car --uri /carfactory -m POST
```
```console
INFO[0000] Service registered                            HTTPMethod=POST ServiceDefinition=create-car ServiceURI=/carfactory SystemName=carprovider
```

Also, register a function to fetch cars.
```console
arrowhead services register --system carprovider --definition get-car --uri /carfactory -m GET 
```
```console
INFO[0000] Service registered                            HTTPMethod=GET ServiceDefinition=get-car ServiceURI=/carfactory SystemName=carprovider
```

Lets list all registered services:
```console
arrowhead services ls --filter car
```
```console
╭────┬───────────────┬─────────────┬────────────────────┬───────────┬──────┬───────────────────╮
│ ID │ PROVIDER NAME │ URI         │ SERVICE DEFINITION │ ADDRESS   │ PORT │ METADATA          │
├────┼───────────────┼─────────────┼────────────────────┼───────────┼──────┼───────────────────┤
│ 84 │ carprovider   │ /carfactory │ create-car         │ localhost │ 8880 │ http-method: POST │
│ 85 │ carprovider   │ /carfactory │ get-car            │ localhost │ 8880 │ http-method: GET  │
╰────┴───────────────┴─────────────┴────────────────────┴───────────┴──────┴───────────────────╯
```

## Authorization
The next step is to add an anothorization rule allowing the *carconsumer* to acccess the *carprovider*.
```console
arrowhead auths add --consumer carconsumer --provider carprovider --service create-car
```
```console
NFO[0000] Authorization added                           AuthID=14
```

```console
arrowhead auths add --consumer carconsumer --provider carprovider --service get-car
```
```console
NFO[0000] Authorization added                           AuthID=15
```

The *carconsumer* can now access the *carprovider*. List list all authorization rules.

```console
arrowhead auths ls 
```
```console
╭────┬──────────────────────┬──────────────────────┬────────────────────┬──────────────────╮
│ ID │ CONSUMER SYSTEM NAME │ PROVIDER SYSTEM NAME │ SERVICE DEFINITION │ INTERFACES       │
├────┼──────────────────────┼──────────────────────┼────────────────────┼──────────────────┤
│ 14 │ carconsumer          │ carprovider          │ create-car         │ HTTP-SECURE-JSON │
│ 15 │ carconsumer          │ carprovider          │ get-car            │ HTTP-SECURE-JSON │
╰────┴──────────────────────┴──────────────────────┴────────────────────┴──────────────────╯
```

## Orchestrations
Let's try to find the carprovider. Note that the we must use the certificate keystore (carconsumer.p12) of the carconsumer to run the command below.

```console
arrowhead orchestrate --system carconsumer --address localhost --port 8881 --keystore ./carconsumer.p12 --password 123456 --service create-car --compact
```
```console
╭────────────────────────────────────╮
│ ORCHESTRATION RESULT               │
├──────────────────────┬─────────────┤
│ FIELD                │ VALUE       │
├──────────────────────┼─────────────┤
│ Provider Address     │ localhost   │
│ Provider Port        │ 8880        │
│ Service URI          │ /carfactory │
│ Service Definition   │ create-car  │
│ Provider System Name │ carprovider │
╰──────────────────────┴─────────────╯
╭──────────────────────────────────────────────────────────────────────────╮
│ AUTHORIZATION TOKENS                                                     │
├──────────────────┬───────────────────────────────────────────────────────┤
│ FIELD            │ VALUE                                                 │
├──────────────────┼───────────────────────────────────────────────────────┤
│ HTTP-SECURE-JSON │ eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2Q0JDLU... │
╰──────────────────┴───────────────────────────────────────────────────────╯
╭──────────────────────────────────────╮
│ INTERFACE                            │
├────────────────┬─────────────────────┤
│ FIELD          │ VALUE               │
├────────────────┼─────────────────────┤
│ Updated At     │ 2023-06-08 12:01:21 │
│ ID             │ 1                   │
│ Interface Name │ HTTP-SECURE-JSON    │
│ Created At     │ 2023-06-08 12:01:21 │
╰────────────────┴─────────────────────╯
```

## Provider implementation

