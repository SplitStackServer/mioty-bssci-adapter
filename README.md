
**WORK IN PROGRESS**

> DISCLAIMER: This project is not endorsed by or associated with any company or the mioty<sup>&reg;</sup> alliance. 

# mioty<sup>&reg;</sup> BSSCI Adapter

The mioty<sup>&reg;</sup> BSSCI Adapter is a service which converts mioty<sup>&reg;</sup> BSSCI protocols into a simplified format (JSON and Protobuf). It is inspired by the [ChirpStack](https://github.com/chirpstack/chirpstack) open-source LoRaWAN<sup>&reg;</sup> Network Server project.

## Backends

The following backends are provided:

* [BSSCI V1.0.0](https://developers.mioty-alliance.com/wp-content/uploads/2025/01/BSSCI_specification_v1.0.0_rev1.pdf)
    * Variable MAC sub channel is mostly implemented except for vm.downlink

## Integrations

The following integrations are provided:

* MQTT 3.1/3.11
    * Topic layout: "basestation/{{ .BsEui }}/#"
        * /state
        * /command
        * /response
        * /event
            * /bs/{{ .EventType }}
            * /ep/{{ .EventType }}





## License

mioty<sup>&reg;</sup> BSSCI Adapter is distributed under the MIT license. See 
[LICENSE](https://github.com/ipaid2win/mioty-bssci-adapter/blob/main/LICENSE).
