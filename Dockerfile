FROM golang:1.24.3-alpine AS development

ENV PROJECT_PATH=/mioty-bssci-adapter
ENV PATH=$PATH:$PROJECT_PATH/build
ENV CGO_ENABLED=0
ENV GO_EXTRA_BUILD_ARGS="-a -installsuffix cgo"

RUN apk add --no-cache ca-certificates make git bash

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN make dev-requirements
RUN make

FROM alpine:3.21.3 AS production

RUN apk --no-cache add ca-certificates
COPY --from=development /mioty-bssci-adapter/build/mioty-bssci-adapter /usr/bin/mioty-bssci-adapter
USER nobody:nogroup
ENTRYPOINT ["/usr/bin/mioty-bssci-adapter"]
