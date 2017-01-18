# Parrot
[![Docs](https://img.shields.io/badge/docs-latest-blue.svg)](https://anthonynsimon.gitbooks.io/parrot/content/)
[![MIT License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://github.com/anthonynsimon/parrot/blob/master/LICENSE)

Self-hosted Localization Management Platform built with Go and Angular 2.  

[Documentation] (https://anthonynsimon.gitbooks.io/parrot/content/)  
[Website] (http://anthonynsimon.com/parrot.github.io)

> NOTICE: the project is still in heavy development and it is NOT recommended for use in production until a version 1.0 is reached.

<img src="http://anthonynsimon.com/parrot.github.io/images/parrot-screenshot-001.png" style="width: 720px;"/>

## Features

- Built-in UI (web app) ready to deploy.
- REST API to easily extend or integrate Parrot into your pipeline.
- Export to various formats: keyvaluejson, `po`, `strings`, `properties`, `xmlproperties`, `android`, `php`, `xlsx`, `yaml` and `csv`.
- Easily rename project strings, Parrot takes care of keeping locales in sync.
- Manage your project's team, assign collaborators and their roles.
- Control API Client access for your projects.
- Easy install/deploy using Docker.

## Try it out

The easiest way to get started is using `docker` and `docker-compose`. Simply clone this repo, navigate to the root of it and start the services:

```
$ git clone https://github.com/anthonynsimon/parrot.git
$ cd parrot
$ /bin/bash scripts/buildweb.sh
$ sudo /bin/bash scripts/start.sh
```

> Please note that to build the web app, `npm` is required.

This will build the web app and launch 3 containers: a Postgres **database**, the Parrot **API server** and Nginx as the **reverse proxy and static file server**.

Navigate to https://localhost/api/v1/ping and you should be able to see if the API is up (your browser will complain about an unknown certificate, see the HTTPS notice below for more info).

And to view the web app simply navigate to https://localhost, it should open the login page of the web app.

### Important note on HTTPS

For convinience, self-signed SSL certificates are provided for the reverse-proxy (nginx). Do **NOT** use them for anything other than development,
use your own certificates instead. We recommended automating the generation and renewal of the certificates via Let's Encrypt.
The `/etc/nginx/certs` and `/etc/nginx/vhost.d` volumes on the nginx container has been made available for this purpose.

## License
This project is licensed under the [MIT](https://github.com/anthonynsimon/parrot/blob/master/LICENSE) license.

## Issues
The recommended medium to report and track issues is by opening one on [Github](https://github.com/anthonynsimon/parrot).

## Contributing
Want to hack on the project? Any kind of contribution is welcome!
Simply follow the next steps:

- Fork the project.
- Create a new branch.
- Make your changes and write tests when practical.
- Commit your changes to the new branch.
- Send a pull request, it will be reviewed shortly.

In case you want to add a feature, please create a new issue and briefly explain what the feature would consist of. For bugs or requests, before creating an issue please check if one has already been created for it.
