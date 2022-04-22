# USTOJ - Master

This is for the course project of CSIT6000O Advanced Cloud Computing, which includes:
[ustoj_deployment](https://github.com/KMFtcy/ustoj_deployment),
[ustoj-master](https://github.com/Dexter1636/ustoj-master), and
[ustoj_front](https://github.com/1023198294/ustoj_front).

## Quick Start

1. Clone the project.

    ```
    git clone https://github.com/Dexter1636/ustoj-master.git
    ```

2. Download modules.

    ```
    cd ustoj-master
    go mod download
    ```

3. Add application config file and test config file.

   Write the following code to your `ustoj-master/config/application.yaml`:

    ```yaml
    server:
      port: 8080
    
    datasource:
      driverName: mysql
      host: <hostname>
      port: <port>
      database: <database_name>
      username: <username>
      password: <password>
      charset: utf8
   
    redis:
      host: <hostname>
      port: <port>
      db: <db>
      user: <username>
      password: <password>
    
    logger:
      level: info
    ```

   And the same for `ustoj-master/config/test.yaml`, which is used for testing.

4. Run.
    ```
    cd api-server
    go build -o build/
    ./build/api-server <config-file-path>
    
    cd scheduler
    go build -o build/
    ./build/scheduler <config-file-path>
    ```

## Build Docker image

1. For api server, run the following command under root directory:
```bash
docker build -f Dockerfile-apiserver . -t ustoj/api-server
```
2. For scheduler, run the following command under root directory:
```bash
docker build -f Dockerfile-scheduler . -t ustoj/scheduler
```

## Note

Do NOT track `application.yaml` and `test.yaml` since they contain sensitive data.

# Documents

- [Database](doc/database.md)
