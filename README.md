# abchain

Small go binary with a single `/continue` HTTP endpoint chaining an alphabet
string one letter at a time. The goal is to generate some bogus traffic
floating around the cluster.

It starts with `/alphabet?continue=a` and ends with `abc...xyz`. Service calls
out to similar peers to continue each letter.


`ABC_MIN_WAIT` and `ABC_MAX_WAIT` vars can be used to tweak how often each pod
schedules a new alphabet chain. More replicas can also be used to saturate the
cluster.

```
❯ kubectl exec -n kube-system -t ds/cilium -- hubble observe -f -t l7 --dict
  TIMESTAMP: Jul 24 18:08:24.419
     SOURCE: default/abchain-86985f9c9c-84vqs:48382
DESTINATION: default/abchain-86985f9c9c-4x9rq:3770
       TYPE: http-request
    VERDICT: FORWARDED
    SUMMARY: HTTP/1.1 GET http://port-abc:3770/continue?alphabet=abcdefghijklmnopqrstuvwz
------------
  TIMESTAMP: Jul 24 18:08:24.435
     SOURCE: default/abchain-86985f9c9c-4x9rq:40088
DESTINATION: default/abchain-86985f9c9c-fgxs9:3770
       TYPE: http-request
    VERDICT: FORWARDED
    SUMMARY: HTTP/1.1 GET http://port-abc:3770/continue?alphabet=abcdefghijklmnopqrstuvwzy
------------
  TIMESTAMP: Jul 24 18:08:24.464
     SOURCE: default/abchain-86985f9c9c-fgxs9:51794
DESTINATION: default/abchain-86985f9c9c-84vqs:3770
       TYPE: http-request
    VERDICT: FORWARDED
    SUMMARY: HTTP/1.1 GET http://port-abc:3770/continue?alphabet=abcdefghijklmnopqrstuvwzyz
```
