FROM registry.cn-shanghai.aliyuncs.com/pigeonligh/build:golang-1.16-alpine AS build

WORKDIR /build

ADD go.* ./

RUN go mod download

ADD cmd ./cmd
ADD common ./common
ADD Makefile ./

RUN make

FROM registry.cn-shanghai.aliyuncs.com/pigeonligh/runtime:alpine-3.14

COPY --from=build /build/_output/bin/puppet /usr/bin/
COPY --from=build /build/_output/bin/control /usr/bin/

RUN ln -s /usr/bin/control /usr/bin/ctr

ENTRYPOINT ["/usr/bin/puppet"]
