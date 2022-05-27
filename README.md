# Poc Go-Mongo Palindromo

## Objetivo
La idea era hacer una aplicación que acceda a una base de datos de productos y que devuelva estos, permitiendo un filtro de texto, a través de una API json, opcionalmente puede incluir un cliente web que consuma la api para desplegar adecuadamente la información.

Además, cuando el filtro sea un texto palíndromo, se debe devolver el producto con un descuento del 50%.


## Uso
```sh
> docker-compose up
```

Este comando gatillará la descarga de la imagen de mongo, si es que no se encuentra en la máquina local, luego comenzará la construcción de la imagen de la aplicación.

Una vez que ambas imágenes se encuentran en el repositorio local, se iniciará la base de datos, lo que gatillará la ejecución de un script de inicio que creará la base de datos, usuarios y datos.

Una vez que la base de datos se encuentra operativa, comienza a ejecutarse la aplicación, la cual creará una conexión permanente a mongo, e iniciará un servidor a la escucha en el puerto 3000.

___Salida esperada del comando anterior___
```sh
Creating network "poc_2022_go_mongo_palindrome_default" with the default driver
Pulling persistence (mongo:3.6.8)...
3.6.8: Pulling from library/mongo
f17d81b4b692: Pull complete
731b1bd9aa76: Pull complete
503d5a66072c: Pull complete
60a7a8d9b4a0: Pull complete
aa55804e1748: Pull complete
f619aedb83bc: Pull complete
f15dff03b4af: Pull complete
8a32790de47e: Pull complete
c659c71afaf3: Pull complete
6301081f09ef: Pull complete
Digest: sha256:75a36c28b6ab0ff5129fcbf79ee62a58f094722a68251067ab113d6da4fb869b
Status: Downloaded newer image for mongo:3.6.8
Building app
Step 1/8 : FROM golang:1.17.3-buster
 ---> 848fd7d0fb9d
Step 2/8 : WORKDIR /go/src/ppamo/server
 ---> Using cache
 ---> 825ff497e82e
Step 3/8 : COPY ./ /go/src/ppamo/server
 ---> Using cache
 ---> d73696db1044
Step 4/8 : RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server .
 ---> Using cache
 ---> c57b8ed4d935
Step 5/8 : FROM scratch
 --->
Step 6/8 : COPY --from=0 /go/src/ppamo/server/bin/server /server
 ---> e58a0165cee7
Step 7/8 : COPY ./static /static
 ---> 8740690a7665
Step 8/8 : CMD ["/server"]
 ---> Running in 7b0c58df2e7a
Removing intermediate container 7b0c58df2e7a
 ---> b9f40492680c
Successfully built b9f40492680c
Successfully tagged app:0.1.0
WARNING: Image for service app was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating poc_db ... done
Creating poc_server ... done
Attaching to poc_db, poc_server
poc_db         | about to fork child process, waiting until server is ready for connections.
poc_db         | forked process: 26
poc_db         | 2022-05-27T09:38:55.189+0000 I CONTROL  [main] ***** SERVER RESTARTED *****
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] MongoDB starting : pid=26 port=27017 dbpath=/data/db 64-bit host=eda3b0bed89d
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] db version v3.6.8
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] git version: 6bc9ed599c3fa164703346a22bad17e33fa913e4
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] OpenSSL version: OpenSSL 1.1.0f  25 May 2017
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] allocator: tcmalloc
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] modules: none
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] build environment:
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten]     distmod: debian92
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten]     distarch: x86_64
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten]     target_arch: x86_64
poc_db         | 2022-05-27T09:38:55.200+0000 I CONTROL  [initandlisten] options: { net: { bindIp: "127.0.0.1", port: 27017, ssl: { mode: "disabled" } }, processManagement: { fork: true, pidFilePath: "/tmp/dock
er-entrypoint-temp-mongod.pid" }, systemLog: { destination: "file", logAppend: true, path: "/proc/1/fd/1" } }
poc_db         | 2022-05-27T09:38:55.200+0000 I STORAGE  [initandlisten]
poc_db         | 2022-05-27T09:38:55.200+0000 I STORAGE  [initandlisten] ** WARNING: Using the XFS filesystem is strongly recommended with the WiredTiger storage engine
poc_db         | 2022-05-27T09:38:55.200+0000 I STORAGE  [initandlisten] **          See http://dochub.mongodb.org/core/prodnotes-filesystem
poc_db         | 2022-05-27T09:38:55.200+0000 I STORAGE  [initandlisten] wiredtiger_open config: create,cache_size=7444M,session_max=20000,eviction=(threads_min=4,threads_max=4),config_base=false,statistics=(fa
st),cache_cursors=false,compatibility=(release="3.0",require_max="3.0"),log=(enabled=true,archive=true,path=journal,compressor=snappy),file_manager=(close_idle_time=100000),statistics_log=(wait=0),verbose=(reco
very_progress),
poc_db         | 2022-05-27T09:38:55.697+0000 I STORAGE  [initandlisten] WiredTiger message [1653629935:697960][26:0x7f742c9fd580], txn-recover: Set global recovery timestamp: 0
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten]
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten] ** WARNING: Access control is not enabled for the database.
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten] **          Read and write access to data and configuration is unrestricted.
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten]
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten]
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten] ** WARNING: /sys/kernel/mm/transparent_hugepage/enabled is 'always'.
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten] **        We suggest setting it to 'never'
poc_db         | 2022-05-27T09:38:55.732+0000 I CONTROL  [initandlisten]
poc_db         | 2022-05-27T09:38:55.732+0000 I STORAGE  [initandlisten] createCollection: admin.system.version with provided UUID: 8b178fa9-c5c3-4c0e-a0c3-a4e1adfc92d8
poc_db         | 2022-05-27T09:38:55.744+0000 I COMMAND  [initandlisten] setting featureCompatibilityVersion to 3.6
poc_server     | 2022/05/27 09:38:55 > Creating connection to mongo using mongodb://app:654321@persistence:27017/products
poc_db         | 2022-05-27T09:38:55.748+0000 I STORAGE  [initandlisten] createCollection: local.startup_log with generated UUID: f74d6ccf-555b-466e-9218-7b79881b888c
poc_db         | 2022-05-27T09:38:55.757+0000 I FTDC     [initandlisten] Initializing full-time diagnostic data capture with directory '/data/db/diagnostic.data'
poc_db         | 2022-05-27T09:38:55.757+0000 I STORAGE  [LogicalSessionCacheRefresh] createCollection: config.system.sessions with generated UUID: 54af787a-330d-4b9d-9d4b-ded8e547ae3a
poc_db         | 2022-05-27T09:38:55.757+0000 I NETWORK  [initandlisten] waiting for connections on port 27017
poc_db         | child process started successfully, parent exiting
poc_db         | 2022-05-27T09:38:55.771+0000 I INDEX    [LogicalSessionCacheRefresh] build index on: config.system.sessions properties: { v: 2, key: { lastUse: 1 }, name: "lsidTTLIndex", ns: "config.system.ses
sions", expireAfterSeconds: 1800 }
poc_db         | 2022-05-27T09:38:55.771+0000 I INDEX    [LogicalSessionCacheRefresh]    building index using bulk method; build may temporarily use up to 500 megabytes of RAM
poc_db         | 2022-05-27T09:38:55.772+0000 I INDEX    [LogicalSessionCacheRefresh] build index done.  scanned 0 total records. 0 secs
poc_db         | 2022-05-27T09:38:55.804+0000 I NETWORK  [listener] connection accepted from 127.0.0.1:50320 #1 (1 connection now open)
poc_db         | 2022-05-27T09:38:55.805+0000 I NETWORK  [conn1] received client metadata from 127.0.0.1:50320 conn1: { application: { name: "MongoDB Shell" }, driver: { name: "MongoDB Internal Client", version
: "3.6.8" }, os: { type: "Linux", name: "PRETTY_NAME="Debian GNU/Linux 9 (stretch)"", architecture: "x86_64", version: "Kernel 4.19.0-18-amd64" } }
poc_db         | 2022-05-27T09:38:55.807+0000 I NETWORK  [conn1] end connection 127.0.0.1:50320 (0 connections now open)
poc_db         | 2022-05-27T09:38:55.854+0000 I NETWORK  [listener] connection accepted from 127.0.0.1:50322 #2 (1 connection now open)
poc_db         | 2022-05-27T09:38:55.855+0000 I NETWORK  [conn2] received client metadata from 127.0.0.1:50322 conn2: { application: { name: "MongoDB Shell" }, driver: { name: "MongoDB Internal Client", version
: "3.6.8" }, os: { type: "Linux", name: "PRETTY_NAME="Debian GNU/Linux 9 (stretch)"", architecture: "x86_64", version: "Kernel 4.19.0-18-amd64" } }
poc_db         | 2022-05-27T09:38:55.866+0000 I STORAGE  [conn2] createCollection: admin.system.users with generated UUID: 9cf502fa-d482-4e3c-a792-156eb78d4400
poc_db         | Successfully added user: {
poc_db         |        "user" : "admin",
poc_db         |        "roles" : [
poc_db         |                {
poc_db         |                        "role" : "root",
poc_db         |                        "db" : "admin"
poc_db         |                }
poc_db         |        ]
poc_db         | }
poc_db         | 2022-05-27T09:38:55.880+0000 E -        [main] Error saving history file: FileOpenFailed: Unable to open() file /home/mongodb/.dbshell: No such file or directory
poc_db         | 2022-05-27T09:38:55.881+0000 I NETWORK  [conn2] end connection 127.0.0.1:50322 (0 connections now open)
poc_db         |
poc_db         | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/01-products.json
poc_db         |
poc_db         | /usr/local/bin/docker-entrypoint.sh: running /docker-entrypoint-initdb.d/init.sh
poc_db         | - Creating user:
poc_db         | MongoDB shell version v3.6.8
poc_db         | connecting to: mongodb://127.0.0.1:27017/products
poc_db         | 2022-05-27T09:38:55.922+0000 I NETWORK  [listener] connection accepted from 127.0.0.1:50324 #3 (1 connection now open)
poc_db         | 2022-05-27T09:38:55.922+0000 I NETWORK  [conn3] received client metadata from 127.0.0.1:50324 conn3: { application: { name: "MongoDB Shell" }, driver: { name: "MongoDB Internal Client", version
: "3.6.8" }, os: { type: "Linux", name: "PRETTY_NAME="Debian GNU/Linux 9 (stretch)"", architecture: "x86_64", version: "Kernel 4.19.0-18-amd64" } }
poc_db         | MongoDB server version: 3.6.8
poc_db         | 2022-05-27T09:38:55.934+0000 I ACCESS   [conn3] Successfully authenticated as principal admin on admin
poc_db         | Successfully added user: { "user" : "app", "roles" : [ "readWrite" ] }
poc_db         | bye
poc_db         | 2022-05-27T09:38:55.948+0000 E -        [main] Error saving history file: FileOpenFailed: Unable to open() file /home/mongodb/.dbshell: No such file or directory
poc_db         | 2022-05-27T09:38:55.949+0000 I NETWORK  [conn3] end connection 127.0.0.1:50324 (0 connections now open)
poc_db         | - Importing data:
poc_db         | 2022-05-27T09:38:55.959+0000 I NETWORK  [listener] connection accepted from 127.0.0.1:50326 #4 (1 connection now open)
poc_db         | 2022-05-27T09:38:55.968+0000 I ACCESS   [conn4] Successfully authenticated as principal admin on admin
poc_db         | 2022-05-27T09:38:55.968+0000   connected to: localhost
poc_db         | 2022-05-27T09:38:55.968+0000 I STORAGE  [conn4] createCollection: products.docs with generated UUID: d32af382-40da-4ad4-8670-48d176083ddc
poc_db         | 2022-05-27T09:38:56.240+0000   imported 3000 documents
poc_db         | 2022-05-27T09:38:56.241+0000 I NETWORK  [conn4] end connection 127.0.0.1:50326 (0 connections now open)
poc_db         |
poc_db         | 2022-05-27T09:38:56.249+0000 I CONTROL  [main] ***** SERVER RESTARTED *****
poc_db         | killing process with pid: 26
poc_db         | 2022-05-27T09:38:56.253+0000 I CONTROL  [signalProcessingThread] got signal 15 (Terminated), will terminate after current cmd ends
poc_db         | 2022-05-27T09:38:56.253+0000 I NETWORK  [signalProcessingThread] shutdown: going to close listening sockets...
poc_db         | 2022-05-27T09:38:56.253+0000 I NETWORK  [signalProcessingThread] removing socket file: /tmp/mongodb-27017.sock
poc_db         | 2022-05-27T09:38:56.254+0000 I FTDC     [signalProcessingThread] Shutting down full-time diagnostic data capture
poc_db         | 2022-05-27T09:38:56.256+0000 I STORAGE  [signalProcessingThread] WiredTigerKVEngine shutting down
poc_db         | 2022-05-27T09:38:56.396+0000 I STORAGE  [signalProcessingThread] shutdown: removing fs lock...
poc_db         | 2022-05-27T09:38:56.397+0000 I CONTROL  [signalProcessingThread] now exiting
poc_db         | 2022-05-27T09:38:56.397+0000 I CONTROL  [signalProcessingThread] shutting down with code:0
poc_db         |
poc_db         | MongoDB init process complete; ready for start up.
poc_db         |
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] MongoDB starting : pid=1 port=27017 dbpath=/data/db 64-bit host=eda3b0bed89d
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] db version v3.6.8
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] git version: 6bc9ed599c3fa164703346a22bad17e33fa913e4
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] OpenSSL version: OpenSSL 1.1.0f  25 May 2017
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] allocator: tcmalloc
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] modules: none
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] build environment:
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten]     distmod: debian92
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten]     distarch: x86_64
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten]     target_arch: x86_64
poc_db         | 2022-05-27T09:38:57.326+0000 I CONTROL  [initandlisten] options: { net: { bindIpAll: true }, security: { authorization: "enabled" } }
poc_db         | 2022-05-27T09:38:57.326+0000 I -        [initandlisten] Detected data files in /data/db created by the 'wiredTiger' storage engine, so setting the active storage engine to 'wiredTiger'.
poc_db         | 2022-05-27T09:38:57.326+0000 I STORAGE  [initandlisten]
poc_db         | 2022-05-27T09:38:57.326+0000 I STORAGE  [initandlisten] ** WARNING: Using the XFS filesystem is strongly recommended with the WiredTiger storage engine
poc_db         | 2022-05-27T09:38:57.326+0000 I STORAGE  [initandlisten] **          See http://dochub.mongodb.org/core/prodnotes-filesystem
poc_db         | 2022-05-27T09:38:57.327+0000 I STORAGE  [initandlisten] wiredtiger_open config: create,cache_size=7444M,session_max=20000,eviction=(threads_min=4,threads_max=4),config_base=false,statistics=(fa
st),cache_cursors=false,compatibility=(release="3.0",require_max="3.0"),log=(enabled=true,archive=true,path=journal,compressor=snappy),file_manager=(close_idle_time=100000),statistics_log=(wait=0),verbose=(reco
very_progress),
poc_db         | 2022-05-27T09:38:57.906+0000 I STORAGE  [initandlisten] WiredTiger message [1653629937:906606][1:0x7f62246dc580], txn-recover: Main recovery loop: starting at 1/798080
poc_db         | 2022-05-27T09:38:57.995+0000 I STORAGE  [initandlisten] WiredTiger message [1653629937:995368][1:0x7f62246dc580], txn-recover: Recovering log 1 through 2
poc_db         | 2022-05-27T09:38:58.053+0000 I STORAGE  [initandlisten] WiredTiger message [1653629938:53296][1:0x7f62246dc580], txn-recover: Recovering log 2 through 2
poc_db         | 2022-05-27T09:38:58.098+0000 I STORAGE  [initandlisten] WiredTiger message [1653629938:98261][1:0x7f62246dc580], txn-recover: Set global recovery timestamp: 0
poc_db         | 2022-05-27T09:38:58.116+0000 I CONTROL  [initandlisten]
poc_db         | 2022-05-27T09:38:58.116+0000 I CONTROL  [initandlisten] ** WARNING: /sys/kernel/mm/transparent_hugepage/enabled is 'always'.
poc_db         | 2022-05-27T09:38:58.116+0000 I CONTROL  [initandlisten] **        We suggest setting it to 'never'
poc_db         | 2022-05-27T09:38:58.116+0000 I CONTROL  [initandlisten]
poc_db         | 2022-05-27T09:38:58.126+0000 I FTDC     [initandlisten] Initializing full-time diagnostic data capture with directory '/data/db/diagnostic.data'
poc_db         | 2022-05-27T09:38:58.127+0000 I NETWORK  [initandlisten] waiting for connections on port 27017
poc_db         | 2022-05-27T09:38:58.240+0000 I NETWORK  [listener] connection accepted from 172.27.0.3:44730 #1 (1 connection now open)
poc_db         | 2022-05-27T09:38:58.241+0000 I NETWORK  [conn1] received client metadata from 172.27.0.3:44730 conn1: { driver: { name: "mongo-go-driver", version: "v1.9.1" }, os: { type: "linux", architecture
: "amd64" }, platform: "go1.17.3" }
poc_db         | 2022-05-27T09:38:58.246+0000 I NETWORK  [listener] connection accepted from 172.27.0.3:44732 #2 (2 connections now open)
poc_db         | 2022-05-27T09:38:58.248+0000 I NETWORK  [conn2] received client metadata from 172.27.0.3:44732 conn2: { driver: { name: "mongo-go-driver", version: "v1.9.1" }, os: { type: "linux", architecture
: "amd64" }, platform: "go1.17.3" }
poc_db         | 2022-05-27T09:38:58.284+0000 I ACCESS   [conn2] Successfully authenticated as principal app on products
poc_server     | 2022/05/27 09:38:58 > Starting server
poc_server     | 2022/05/27 09:38:58 > Mux router listening at 0.0.0.0:3000
```

&nbsp;

## Aplicación

La aplicación escucha peticiones HTTP por el puerto 3000, y esta preparada para servir los archivos estáticos que se encuentren en la carpeta */static* del repositorio.   Además cuenta con rutas estáticas que devuelven información relacionada a los productos:
- **/product/{id}**: devuelve un producto en particular, dado un id de producto.
- **/products**: devuelve el listado completo de productos en la base de datos.   Esta ruta también acepta un filtro de busqueda, a través de la variable **q**, que contiene el texto que se utilizará para filtrar los datos.

&nbsp;

#### Rutas

##### - /
Esta ruta devuelve el archivo estático por defecto, en éste caso el contenido del archivo [static/index.html](/static/index.html) del repositorio

```sh
> curl http://localhost:3000/
Hello World!! Hello World!! Hello World!!
```

##### - /product/{id}
Esta ruta devuelve un producto en particular, identificado por su id
```sh
> curl http://localhost:3000/product/2198 | jq
{
  "id": 2198,
  "brand": "dsaasd",
  "description": "frky givgxzqk",
  "image": "www.lider.cl/catalogo/images/homeIcon.svg",
  "price": 650078,
  "discount": 0
}
```

En caso que el identificador no se encuentre en la base de datos devuelve un error genérico:
```sh
> curl http://localhost:3000/product/21981 | jq
{
  "message": "Error getting the product"
}
```

##### - /products
La ruta */products* es un endpoint de la API que devuelve la lista completa de productos en la base de datos:
```sh
> curl http://localhost:3000/products | jq
[
  {
    "id": 1,
    "brand": "ooy eqrceli",
    "description": "rlñlw brhrka",
    "image": "www.lider.cl/catalogo/images/whiteLineIcon.svg",
    "price": 498724,
    "discount": 0
  },
  {
    "id": 2,
    "brand": "dsaasd",
    "description": "zlrwax bñyrh",
    "image": "www.lider.cl/catalogo/images/babyIcon.svg",
    "price": 130173,
    "discount": 0
  },
  .
  .
  .
```

Los datos que devuelve la ruta pueden ser filtrados por texto en la marca o descripción, usando la variable *q*:
```sh
> curl http://localhost:3000/products?q=saa | jq
[
  {
    "id": 2,
    "brand": "dsaasd",
    "description": "zlrwax bñyrh",
    "image": "www.lider.cl/catalogo/images/babyIcon.svg",
    "price": 130173,
    "discount": 0
  },
  {
    "id": 238,
    "brand": "dsaasd",
    "description": "xhabvp ciowh",
    "image": "www.lider.cl/catalogo/images/babyIcon.svg",
    "price": 272633,
    "discount": 0
  },
  .
  .
  .
```

Y cuando el texto de busqueda es palíndromo, se aplica un descuento del 50%:
```sh
> curl http://localhost:3000/products?q=saas | jq
[
  {
    "id": 2,
    "brand": "dsaasd",
    "description": "zlrwax bñyrh",
    "image": "www.lider.cl/catalogo/images/babyIcon.svg",
    "price": 130173,
    "discount": 50
  },
  {
    "id": 238,
    "brand": "dsaasd",
    "description": "xhabvp ciowh",
    "image": "www.lider.cl/catalogo/images/babyIcon.svg",
    "price": 272633,
    "discount": 50
  },
  .
  .
  .
```
