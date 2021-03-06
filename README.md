Server Identity Management
========
An application that creates and manages a Factom Server's Identity. This application will not store or save any private keys once the application has terminated, so it is up to the user to copy down the keys displayed and store them appropriately.

Here is a list of the entries that are created by this program and which chains they will be entered into. Each column represents a factom chain. Italicized names are the first entry in a chain, while unbolded are new entries.

|Factom Identity List<br />ChainID:[88888800...]|Identity Chain <br />ChainID:[888888...]|Server Management SubChain<br />ChainID:[888888...]
|:---:|:---:|:---:|
|"Register Factom Identity"<br />sk1 key sign|<i>"Identity Chain"</i>|<i>"Server Management"</i>
||"Register Server Management"<br />sk1 key sign|"New Block Signing Key" <br />sk1 key sign
|||"New Bitcoin Key"<br />sk1 key sign|
|||"New Matryoshka Hash"<br />sk1 key sign
## Compiling
* /identity/varsconfig.go contains various hardcode options. Until release, refer to these when creating an identity.


Using Server Management Tool
========

There are two ways to create an Identity. You can choose the manual way, and choose each key, or do the quick automatic way and allow keys to be chosen for you randomly.

### Identity Check Tool
You can also check to see if an identity is "full" and able to be promoted. It will return true/false to indicate full/not full
```
serveridentity check CHAINID -host={HOST} -v={true/false}
```
* host = Location of factomd api
  * default: localhost:8088
* v = printout keys if all found
  * default: false
    * Only prints true/false as return

# Automatic (Quick) Method -- Recommended for new Identity
## Requirements

* signwith25519
  * Go Install 'signwith25519' in the subdirectory 'serveridentity/signwith25519/'
* factom-cli
* factomd
* factom-walletd

## Generating a new Identity

All Keys will be generated randomly and printed out to stdout
```
serveridentity full elements {ENTRY_CREDIT_PRIVATE_KEY} -n={OPTIONAL_FILENAME}
```

Several files will be produced:
* Scripts to add the Identity to the blockchain
  * Script utilizes factom-cli
  * Name of script is by default 'fullidentity.sh' or ''{OPTIONAL_FILENAME}.sh' if provided
* Config file needed for the server
  * Place in ~/factom/m2 and rename to factomd.conf
  * Name of config is by default 'fullidentity.conf' or ''{OPTIONAL_FILENAME}.conf' if provided

### Building Identity and Launching Factomd with new Identity
1. Copy down keys in stdout
* Start factomd and wait for it to sync
* Start factom-walletd and use cli/gui to import the private entry credit address
  * Load with entry credits (>= 26)
* Run 'fullidentity.sh' (or given name)
* Place config in required spot and rename
* Restart factomd with config file




# Manual Way

## Creating a new Server Identity
To create a new server identity is very simple, after compiling the source code run:
```
serveridentity start fresh
```

1. A lot of information will be outputted, **copy down** the private keys and the entry credit address private key. The identity keys and public key can be found using the private key so are not needed to be copied down.
  * You can pipe the output to a file to make it easier to copy down, but make sure to delete all trace of the file as these keys should be stored securely
    * ```serveridentity start fresh > out.txt```
    * type ``` y ``` press enter, then press ```ENTER``` again and wait. That will guide you through the prompts and start the generating
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
serveridentity newkey block [-s]
 ```
* Follow Prompts, -s flag will generate a script to run the curl commands
* Prompts
  1. Identity ChainID: Root identity chain ID
  2. Subchain ID: Sub identity chain ID
  3. Private Key: Level 1 (sk1) private Key
  4. Entry Credit Key: Can choose 'any' or import an entry credit address. Entry credits will be taken from this address to pay the cost of the entry

## Create a new Bitcoin Key
```
serveridentity newkey btc [-s]
 ```

 * Follow Prompts, -s flag will generate a script to run the curl commands
 * Prompts
   1. Identity ChainID: Root identity chain ID
   2. Subchain ID: Sub identity chain ID
   3. Bitcoin Key Level: Integer 0-3, start at 0 for first Key
   4. Bitcoin Address Type: 0 for PKPKH or 1 for P2SH
   5. Bitcoin address P2PKH or P2SH Type
   6. Private Key: Level 1 (sk1) private Key
   7. Entry Credit Key: Can choose 'any' or import an entry credit address. Entry credits will be taken from this address to pay the cost of the entry


## Create a new Matryoshka Hash
```
serveridentity mhash [-s]
```

   * Follow Prompts, , -s flag will generate a script to run the curl commands
     1. Identity ChainID: Root identity chain ID
     2. Subchain ID: Sub identity chain ID
     3. Seed: Hex string of max length 64
     4. Private Key: Level 1 (sk1) private Key
     5. Entry Credit Key: Can choose 'any' or import an entry credit address. Entry credits will be taken from this address to pay the cost of the entry


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

* Example: ``` serveridentity get idkey sk11pz4AG9XgB1eNVkbppYAWsgyg7sftDXqBASsagKJqvVRKYodCU ```

## Get Private Key Hex from Private Sk# Key
```
serveridentity get privkey PRIVATEKEY
```

* Example: ``` serveridentity get privkey sk11pz4AG9XgB1eNVkbppYAWsgyg7sftDXqBASsagKJqvVRKYodCU ```

## Get Bitcoin Hash160 from Public key
```
serveridentity get btckey BTCKEY
```

* Example: ``` serveridentity get btckey 1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2 ```
