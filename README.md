Server Identity Management
========
An application that creates and manages a Factom Server's Identity. This application will not store or save any private keys once the application has terminated, so it is up to the user to copy down the keys displayed and store them appropriately.

## Creating a new Server Identity
To create a new server identity is very simple, after compiling the source code run:
```
serveridentity start fresh
```

1. A lot of information will be outputted, **copy down** the private keys, identity keys, and the entry credit address private key.
  * You can pipe the output to a file to make it easier to copy down, but make sure to delete all trace of the file as these keys should be stored securely
    * ```serveridentity start fresh > out.txt```
2. Press 'ENTER' when you copied down your keys, the next step will take some time. Wait for it to complete
3. Various curl commands will be outputted. Copy down these commands, as they are needed to be run in the order they appear to successfully create an identity chain and subchain.
  * The ChainIDs of the identity chain and subchain will be located here
  * A file called 'startidentity.sh' will be created and can be run instead of doing the curl commands manually
  * The commands or 'startidentity.sh' must be done on an online machine with factomd running and funds must be available in the entry credit address outputted
4. Now the identity is ready to be uploaded to the Factom network. Start up factomd and fctwallet
5. Move at least 24 entry credits to the addressed specified in the output of step 1. This address will pay for the entries
6. Manually enter the curl commands in order, or run ```sh startidentity.sh```
7. If all success messages are returned, your server identity has been successfully created and registered.

## Create a new Block Signing Key
```
serveridentity newkey block
 ```
* Follow Prompts

## Create a new Matryoshka Hash
```
serveridentity mhash
 ```
* Follow Prompts

## Create a new Bitcoin Key
```
serveridentity newkey btc
 ```
 * Follow Prompts


 ## Get Pubkey from Private
 ```
 serveridentity get pubkey PRIVATEKEY
  ```
* Follow Prompts
* Example: ``` serveridentity get pubkey sk11pz4AG9XgB1eNVkbppYAWsgyg7sftDXqBASsagKJqvVRKYodCU ```

## Get Identity Key from Private
```
serveridentity get idkey PRIVATEKEY
 ```
* Follow Prompts
* Example: ``` serveridentity get pubkey sk11pz4AG9XgB1eNVkbppYAWsgyg7sftDXqBASsagKJqvVRKYodCU ```
