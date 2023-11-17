# Sample Go HTTP Api with FTP server & small svelte application for markup images

* All server-side stuff like ftp, database and server himself are launched as 3 docker containers

## Implemented features
* Basic jwt auth, hash & sanitized pwd
* Storing projects with their stuff like images/labels/annotations
* Export all annotations of selected image in CSV format

## Needed implement
* **Configure a client build adapter for deploying** 
* Added object detection app (native or third party resource)
* Separate units (backend with database/client and obj. detection app) & added kafka or rabbit mq as message broker between them

## How to use it
* Move to client folder and download all deps (***cd client && npm i && cd ..***)
* Run command: ***make docker-build*** (all deps are installed automatically like migrations, access rights, etc.)
* Run command: ***make run-client***

## Preview
https://github.com/skhe55/intro-ai/assets/72722478/ac668053-f4a2-4a28-a96a-c509f7a3fea6



