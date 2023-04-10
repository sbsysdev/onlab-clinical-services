# OnLab-Clinical Services

## Instructions for ***development***

- To **start** the **Docker** image app, execute the command: **`docker compose -f docker-compose.yml up -d`**

- Open the **local** link: **[localhost:8080](http://localhost:8080)**

- To **stop** the **Docker** image app, execute the command: **`docker compose -f docker-compose.yml down`**

> ### If there are any changes in the source code, run the commands to remove the old Docker image and create a new one when the current Docker image is stopped
>
> - To get the current **Docker image name**, execute the command: **`docker images`**
>
> - To **delete** the current **Dokcer image**, execute the commad: **`docker image rm <dokcer-image-name>`**
