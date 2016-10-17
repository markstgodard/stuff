# Stuff :movie_camera:
Stuff is a **fake store** is a microservices-based demo application that illustrates the use of [CF Container Networking](https://github.com/cloudfoundry-incubator/netman-release).

The app uses [Amalgam8](http://amalagam8.io) for [sidecar](https://www.amalgam8.io/docs/sidecar) and Service Discovery / Routing [control plane](https://www.amalgam8.io/docs/control-plane).

This project is the store front app that uses the following microservices:
- [Products](http://github.com/markstgodard/stuff-products)
- [Reviews](http://github.com/markstgodard/stuff-reviews)

![alt text](https://raw.githubusercontent.com/markstgodard/stuff/master/shop.png "Shop Stuff")

About this demo app:
- has a single external route for store front app
- all other microservices communication and policy enforcement will be on the internal overlay [CF Container Network](https://github.com/cloudfoundry-incubator/netman-release)
- CF apps are deployed as [Docker](https://docker.com) containers and leverage a [sidecar](https://www.amalgam8.io/docs/sidecar) inside the container that handles service discovery, health checks and routing requests to other microservices
- amalgam8 control plane (service registry / controller) is not deployed with auth enabled, please see [a8 docs](https://www.amalgam8.io/docs/sidecar/sidecar-configuration-options) for more info.


## Prerequisites
- [CF](https://github.com/cloudfoundry/cf-release) deployment
- [Diego](https://github.com/cloudfoundry/diego-release) deployment
  - [Docker support](https://github.com/cloudfoundry/diego-design-notes/blob/master/docker-support.md) enabled
  - [Netman support](https://github.com/cloudfoundry-incubator/netman-release) enabled
- [cf cli](http://docs.cloudfoundry.org/cf-cli)
  - [netman network policy ](https://github.com/cloudfoundry-incubator/netman-release/releases) cf cli plugin installed
- [jq](https://stedolan.github.io/jq/)

## Configuration
The scripts in this example uses this [config file](./scripts/cf.cfg) to configure CF domains, app names, etc.
If you wish to use the scripts to deploy the demo apps, please change the values to match your target environment.
The defaults assume [bosh-lite](https://github.com/cloudfoundry/bosh-lite) and that you already are targeting a org and space.

## Deployment

### Deploy Amalgam8 Controller and Registry
```sh
./scripts/deploy-a8.sh
```

You should now see the amalgam8 controller and registry apps running:
```sh
$ cf apps
Getting apps in org demo / space demo as admin...
OK

name               requested state   instances   memory   disk   urls
stuff-controller   started           1/1         256M     1G     stuff-controller.bosh-lite.com
stuff-registry     started           1/1         256M     1G     stuff-registry.bosh-lite.com
```

### Deploy Store app

Run this script:
```sh
./scripts/deploy.sh
```

To deploy the [Products](https://github.com/markstgodard/stuff-products) and [Reviews](https://github.com/markstgodard/stuff-reviews) app, please checkout those projects and run the same `./scripts/deploy.sh` script.


### Check Apps
After deploying you should see the controller, registry and store apps with routes. The products and review apps will not have external routes.
```sh
$ cf apps
Getting apps in org demo / space demo as mark...
OK

name               requested state   instances   memory   disk   urls
stuff-controller   started           1/1         256M     1G     stuff-controller.bosh-lite.com
stuff-registry     started           1/1         256M     1G     stuff-registry.bosh-lite.com
stuff              started           1/1         256M     1G     stuff.bosh-lite.com
stuff-products     started           1/1         256M     1G
stuff-reviews      started           1/1         256M     1G
```

### Check Network Policy
The above scripts/commands will create the appropriate network policy to ensure that only appropriate apps can talk to each other.
```sh
$ cf access-list
Listing policies as admin...
OK

Source          Destination     Protocol        Port
stuff           stuff-products  tcp             8080
stuff-products  stuff-reviews   tcp             8080
```

### Try it
Go to store app, for example: [http://stuff.bosh-lite.com](http://stuff.bosh-lite.com)

## Now what?
For more information on CF Container Networking:
- [CF Container Networking release](https://gitcom.com/cloudfoundry-incubator/netman-release)
- Chat with us at the `#container-networking` channel on [CloudFoundry Slack](http://slack.cloudfoundry.org/)


## Disclaimer
`Don't try and actually buy stuff.`
