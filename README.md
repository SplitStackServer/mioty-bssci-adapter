
**WORK IN PROGRESS**

> DISCLAIMER: This project is not endorsed by or associated with any company or the [mioty<sup>&reg;</sup> alliance](https://mioty-alliance.com/). 

# mioty<sup>&reg;</sup> BSSCI Adapter

The mioty<sup>&reg;</sup> BSSCI Adapter is a service which converts mioty<sup>&reg;</sup> BSSCI protocols into a simplified format (JSON and Protobuf). mioty<sup>&reg;</sup> is a LPWAN technology that implements the TS-UNB protocol as specified in [ETSI TS 103 357-2 V2.1.1](https://www.etsi.org/deliver/etsi_ts/103300_103399/10335702/02.01.01_60/ts_10335702v020101p.pdf
).

It is inspired by the [ChirpStack](https://github.com/chirpstack/chirpstack) open-source LoRaWAN<sup>&reg;</sup> Network Server project and follows the structure of the [ChirpStack Gateway Bridge](https://github.com/chirpstack/chirpstack-gateway-bridge). Where possible this implementation reuses existing code and concepts.

This component is part of the [SplitStackServer](https://github.com/SplitStackServer/splitstack), a fork of ChirpStack, to provide an open-source mioty<sup>&reg;</sup> Network Server.

## Backends

The following backends are provided:

* [BSSCI V1.0.0](https://developers.mioty-alliance.com/wp-content/uploads/2025/01/BSSCI_specification_v1.0.0_rev1.pdf)
    * Variable MAC sub channel is mostly implemented except for vm.downlink. Variable MAC uplinks can contain 

## Integrations

The following integrations are provided:

* MQTT 3.1/3.11
    * Currently the topics are hardcoded as follows: 
        * root: /bssci/#
        * basestation: bssci/{{ .BsEui }}/#
            * state: bssci/{{ .BsEui }}/state
            * events: 
                * endpoint events: "bssci/{{ .BsEui }}/event/{{ .EventSource }}/{{ .EventType }}"
                    * {{ .EventSource }} is "ep"
                    * {{ .EventType }} is one of "otaa", "dl, "ul", "rx"
                * basestation events: "bssci/{{ .BsEui }}/event/{{ .EventSource }}/{{ .EventType }}"
                    * {{ .EventSource }} is "bs"
                    * {{ .EventType }} is one of "status", "con, "vm"
            * server commands: "bssci/{{ .BsEui }}/command/#"
            * server responses: "bssci/{{ .BsEui }}/response/#"


# Building 

This project uses `goreleaser` and `nfpm` to handle build and packaging tasks. `goreleaser` also builds a Docker image.

Helpful commands:


```bash
# install development requirements
make dev-requirements

# run the tests
make test

# compile
make build

# build for supported architectures using goreleaser
make dist

# compile snapshot for supported architectures using goreleaser
make snapshot
```

## License

mioty<sup>&reg;</sup> BSSCI Adapter is distributed under the MIT license. See 
[LICENSE](https://github.com/SplitStackServer/mioty-bssci-adapter/blob/main/LICENSE).
