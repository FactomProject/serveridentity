
echo ">factom-cli importaddress zeros Es2Rf7iM6PdsqfYCo3D1tnAR65SkLENyWJG1deUzpRMQmbh9F3eG"
factom-cli importaddress zeros Es2Rf7iM6PdsqfYCo3D1tnAR65SkLENyWJG1deUzpRMQmbh9F3eG

echo ">factom-cli importaddress sand Fs3E9gV6DXsYzf7Fqx1fVBQPQXV695eP3k5XbmHEZVRLkMdD9qCK"
factom-cli importaddress sand Fs3E9gV6DXsYzf7Fqx1fVBQPQXV695eP3k5XbmHEZVRLkMdD9qCK

echo ">factom-cli newtransaction trans1"
factom-cli newtransaction trans1

echo ">factom-cli addinput trans1 sand 10"
factom-cli addinput trans1 sand 10

echo ">factom-cli addecoutput trans1 zeros 10"
factom-cli addecoutput trans1 zeros 10

echo ">factom-cli addfee trans1 sand"
factom-cli addfee trans1 sand

echo ">factom-cli sign trans1"
factom-cli sign trans1

echo ">factom-cli transactions"
factom-cli transactions

echo ">factom-cli submit trans1"
factom-cli submit trans1

sleep 1

echo Factom Identity Chain List Create
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"message":"0001553ba74d8faa6ac2d4961882f42a345c7615f4133dde8e6d6e7c1b6b40ae4ff6ee52c393d024cbe2e7f360baad36a66b4f063f1f1b9f57f25deb35aad8fba8905cf2893eec1be40ce17636636117d9469de0f027cd74754e0e1871d249dfefac958d0f91de0b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da299999aa8cfd722db62c61e53c7dbf9fa4de1a64b9891844f1d53b78a4cea3294fb6b88e5b53e5f132e32e1b1176335ead8ed351787457b9219f7743cc51b42803"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85500004d61696e204964656e74697479204c697374"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  
echo \*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*
echo \* Setup script will create and register an identity and its subchain \ \ \ \ \ \ \*
echo \* Credits must be in EC2DKSYyRcNWf7RS963VFYgMExoHRYLHVeCfQ9PGPmNzwrcmgm2r \ \*
echo \*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*
echo - Creating Identity Chain - ChainID: 888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b0
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"message":"0001555ed13962607cc2e851bdda1978af755711c7cc3ae1f1db47596a7b5b8e2176c44be02935cce5233911330e6662ef3ed6c363369075b02fc0e66e97804af6602efa2b870e7305f37fc0772f46bf7a7b65574cc4e5a9f98ccdf5c9a354aea48556c6d8a3b30b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29177ad364e7e7e5f6f73111ed72648ff3f7458b403a9676598ac686cf8ac457ddd043519020fed558dccfa0a46d8fa0382eb26fb908fbba631167f1a0e8803709"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b000a5000100000e4964656e7469747920436861696e0020febde62ee7aa79aef4f32295d9bae9cd3e59bf775e7a765b9ac7175f9f0290d2002079d8cd6f29799f317bf8c9d15c526ff58178c3a756740cb43f59ea7cd8adec170020357aab63e331bf799859e1462c3d6479cdeacd93eb8b52a9b2deb2ce285c9994002051df0be0c47f3a6bc4382d5a2bb4cbdbbe88b77db6bfe02c0933ba937142952b00083130373333393331"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Registering Identity Chain
curl -X POST --data '{"jsonrpc":"2.0","id":3,"params":{"entry":"0001555ed139636f4fa354a90f01910b7cf998a6f0d5839a1b911f7138830fe2defe8e76d16a70013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da291d09d8515a71788fd50d08ec0e0b9d0b90afa3a552b9d48a16051749cefaa333e61c788ed7b2c00a1ca11e965f16b0af314c0509bc622eb3255ae2d9c31db40c"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":4,"params":{"entry":"00e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85500a40001000018526567697374657220466163746f6d204964656e746974790020888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b0002101cf80ab51d583b2a32360fc005ee18c52ea5d4205a175ef286f1cd5615cdc47c30040626400beb77b056ef0f6952487944b28b8a80f00da03e8803302c7797eec267bbb3f314ff01b570480fa6d46372b650948a20d2bf3196255d9a3fb47487ba20f"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Creating Server Management SubChain - ChainID: 888888be4b6ae38c4ef1e4bcaf8859cf1636ed265c7a790f2e6710b304b97acf
curl -X POST --data '{"jsonrpc":"2.0","id":5,"params":{"message":"0001555ed1aa4e24b00d17ee807dfbaf8a38ed2d24d4396ca9c6d4c63d21c722fef29cb2d6dfe20798ee84f8106b81386be3e16ba0f3fb7ad6df73c9f7d6e52b3fd6eb9b4d8638af1421cc9655857517b6c70f9d6955a9a9dde8cad83d4947a406c133b93f3d640b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29878df5e2b04464fec3fa4d48e851cf91e3c0c2303ab4ba2f2e9885d41e33007901e90962b90f2a00f7ea52f2ca4ca9b29307a1c5de8f0d4ce0f3fa476050f101"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":6,"params":{"entry":"00888888be4b6ae38c4ef1e4bcaf8859cf1636ed265c7a790f2e6710b304b97acf00420001000011536572766572204d616e6167656d656e740020888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b000083130353839363532"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Registering Server Management SubChain
curl -X POST --data '{"jsonrpc":"2.0","id":7,"params":{"entry":"0001555ed1aa4f1c8236d64141e4123aa3a190cf8f8c95c2b5db9512c1688f83b6e50b54dbc7bd013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29d95b94b81d0837427c9a71c47868bbadf51e10c1a3dde803fa244b9310e8b298a0078d76d28438433a585ed50d86bb20a6faf9747e27688497c921691fae0b08"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":8,"params":{"entry":"00888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b000a6000100001a526567697374657220536572766572204d616e6167656d656e740020888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b0002101cf80ab51d583b2a32360fc005ee18c52ea5d4205a175ef286f1cd5615cdc47c30040048ad6a4e8b710eb8b9b4d88d668409ece78f0d05e2790177c52d9312c506009d74c585522ed5088a5b2af30d5b3d5b82a72d326956756bbef4956bf8c1bce09"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
sleep 1
echo - New Bitcoin Key
curl -X POST --data '{"jsonrpc":"2.0","id":9,"params":{"entry":"0001555ed1aa506350b28f88f02a108344628a9e286f202cc57ac39a8e853408dfed369730663f013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da2999b3b98d058c6bc89e99752b34ca8f2a7227ce3e81df0ca0a68ae432dcfdca4c3940a4bd3b71e71dad61c9fa6eda8418f9ed34fb8d5469e0e31e24f01a337a0b"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":10,"params":{"entry":"00888888be4b6ae38c4ef1e4bcaf8859cf1636ed265c7a790f2e6710b304b97acf00c1000100000f4e657720426974636f696e204b65790020888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b0000100000100001483be7c11ced9c74d696676557d8e3b225a47565e00080000000057640cb6002101cf80ab51d583b2a32360fc005ee18c52ea5d4205a175ef286f1cd5615cdc47c300409e6d4d57f7e7d085649ed0e5f17c30ffa36525655e116fcc8069a17c33aac8fe80655315602bfcb54f8906445d44a4584ab02b3470cb53572fc69729e0e6bb0a"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - New Block Signing Key
curl -X POST --data '{"jsonrpc":"2.0","id":11,"params":{"entry":"0001555ed1aa50f31ea1fb3a00543020886de26f499ae6cf8b88e7c509fd954092a901d9985892013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da2934ac2b53ba614fc1f73d92cb6dcd24a87fa7818cf270c6aa96c0687aa296c4292ac32a294496a074aeedb4ba25d2de9f8238439ca82b8712b4f81a90a6c11300"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":12,"params":{"entry":"00888888be4b6ae38c4ef1e4bcaf8859cf1636ed265c7a790f2e6710b304b97acf00cd00010000154e657720426c6f636b205369676e696e67204b65790020888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b000206733bf558d8497a4d993ba44eeec26b0ccfd6d4f0875f4e687c7a80fd93b378900080000000057640cb6002101cf80ab51d583b2a32360fc005ee18c52ea5d4205a175ef286f1cd5615cdc47c300402306274b42c69122bee9577b0f1b38c53df0fa763e388a35333497706a8bd8c3454b2fe88d540c4cffe6d94b37cf734d9c8a4b9933a198baa5f49708933f760a"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - New Matryoshka Hash
curl -X POST --data '{"jsonrpc":"2.0","id":13,"params":{"entry":"0001555ed1aa869690ba5e5d6a20a80d48de3b56e032889295d333f4a98759a505d7b079bf019e013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da2988c001651fd4ea6dfaae8b0b1a28d80e0d8ee8eb254d270f3e1181e1f43820fe958bd5f9fa2abdc30dbe0addc75fb48ce2e453ba886a5f6ac5d0f26b0055bc06"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":14,"params":{"entry":"00888888be4b6ae38c4ef1e4bcaf8859cf1636ed265c7a790f2e6710b304b97acf00cb00010000134e6577204d617472796f73686b6120486173680020888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b00020aabf83809fb595eb8702d67f941198ab5a54849b50b9794a13cdd7e93da32cb500080000000057640cb6002101cf80ab51d583b2a32360fc005ee18c52ea5d4205a175ef286f1cd5615cdc47c300406493a4c0de2186eaaf13371013b7a92c79beaf53ff11ba471e390eff583601fe02b6d09e924b17cdf6b863ae65e4155802fa49fbbf514ac72317f1ee8c15620a"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   

echo  Identity Chain Created, info:
echo  Private Keys:
echo  Level 1: sk13Dc8BtXSzJYSPypTk1YHUWo5Jrt93Sj8kv2TYUBUMKxVMFgXVD
echo  Level 2: sk23tQZPwAG8pyLELWmgYSpUpMyRPYWFmsTFNcYJAaZQHq3MoUJQ1
echo  Level 3: sk32mtDfwgqAmTyVokjQVfNqsxhv2vBrRAsd3712RD2FKUXkR3shW
echo  Level 4: sk449f5Tyu8a5ogX1cy5dJL5oSeVTqubth8ue6on8gKySmrosNxLE
echo  
echo  ChainIDs:
echo  Main Factom Identity List ChainID: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
echo  Root: 888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b0
echo  Sub \: 888888be4b6ae38c4ef1e4bcaf8859cf1636ed265c7a790f2e6710b304b97acf
echo
echo  Block Signing Key: 7c6f3020011231fe31ad3ed3d78f4112a155458799ff32bc94f249d99654454a6733bf558d8497a4d993ba44eeec26b0ccfd6d4f0875f4e687c7a80fd93b3789
echo  BtcKey: 1D1biEdmKwq6CVkFPsDkYKry8Ng1opJwM3
echo  
echo  MHash Seed: 888888742029e8901501d7b163588c9f57916dc5ed283a9473b7dcbec2e5f9b0
echo  MHash: aabf83809fb595eb8702d67f941198ab5a54849b50b9794a13cdd7e93da32cb5

