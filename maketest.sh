
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

echo \*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*
echo \* Setup script will create and register an identity and its subchain \ \ \ \ \ \ \*
echo \* Credits must be in EC2DKSYyRcNWf7RS963VFYgMExoHRYLHVeCfQ9PGPmNzwrcmgm2r \ \*
echo \*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*\*
echo - Creating Identity Chain - ChainID: 88598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"message":"0001555477654abd391809a5adea2d429fbfc5b9cbd89843d1bd0a3045a341a78d0214b0383340c404032fbd2221b95c42e413e8fd561667fe7bbca52bafd8a2fdcf3b29245e13afad7b45499ac5be7e79b0129dfc256671056620e68d33d4618e001bcd8e9f5b0b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da294f8cea562fe23fa54cd1fbab3308355e90320cb56f2dd21928f6c38a12c7bc5060ef57b2d094ac4bb9af50f817d61eea37409a227619497483386c62b1c8f30a"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"008888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d400a3000100000e4964656e7469747920436861696e00209c9c2c0cfb8f747ac1ae1e88e5186ced277ffabdd42a4f0c85a24f2b7d341a350020ea9ee62daf5a1500ceb7762bbb412dd0f38f4163022082d0fbdaee999a0068d300205d8b61022c64582e96e7f49998f9a2b8a746f169cc642c0aec695cfd60808c74002012429abdf57c788cd2e2aed51437a611c52800bc8153de289b4ba661f937ce4c0006313338353234"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Registering Identity Chain
curl -X POST --data '{"jsonrpc":"2.0","id":3,"params":{"entry":"0001555477654f22777bdb211df516857f693d277a2a1ad4949e6ad7af4c7cfa627bb24cd19dcf013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29a1fe621168cf339766f3607e638609cf59e218aca9b6d411f4f1d04e3ac5186277f050d5f004faadc7013cd475f93d0e69281d89c5c0456133a20a6bf123450c"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  
curl -X POST --data '{"jsonrpc":"2.0","id":4,"params":{"entry":"00e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85500a40001000018526567697374657220466163746f6d204964656e7469747900208888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d4002101cda3cbde71b6277c21ca3af6be5e8038a851898aa14b6b5a78ae28ff582135e200402424784e69452274f56616cb4d5ecaa7d7d41c68715c1b3b10153f5beae8d6fcd23e8f90e52d30f3d8ffc2ac2653374a70c8cebf5a69f32eca27eee0a9eb5c06"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   

echo - Creating Identity SubChain - ChainID: 883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c
curl -X POST --data '{"jsonrpc":"2.0","id":5,"params":{"message":"0001555477ff4188fe60fc3fa3ebc078b066e0ef5c2bfa27843d821fa50804036a509682f2190973836c368da33387897faf105986d8bee4c95c1cac1768760a420f4e505819c07e7c1661f90a8d64a25044ac7bf3c542a2049bc2dd2f8d62451d2d477c8eb1810b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da291c69e7409a90ea4467167f1f90f340700a6dc037f7fcda8613c2d24fdce169934bd90fd6e0f905bcecf4f4055c696d3be20b32db9d4887932132e732d078f50f"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  
curl -X POST --data '{"jsonrpc":"2.0","id":6,"params":{"entry":"00888888898fbc055f7e201b6e039bcfb859da557e10ae4d0a25b13f04accbc32700410001000011536572766572204d616e6167656d656e7400208888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d4000738313232333333"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Registering Identity SubChain
curl -X POST --data '{"jsonrpc":"2.0","id":7,"params":{"entry":"0001555477ff4268741b0c03163d7d900d869a36eff30b74c0cdd4d3ac4ca554fd5a869c25d8b6013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29e654faf2067a29d78d5c01a8239e4c5897d3c2131e17debda553f9a14f2bb16ff1bea3404e5ea707c70a60edb124f1431c9cc4de98b5e99aec2f546146416205"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  
curl -X POST --data '{"jsonrpc":"2.0","id":8,"params":{"entry":"008888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d400a6000100001a526567697374657220536572766572204d616e6167656d656e7400208888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d4002101cda3cbde71b6277c21ca3af6be5e8038a851898aa14b6b5a78ae28ff582135e2004081ccc8eaee4ecba4e0243103646ad6b992de5c1daaabdb9b160b8b87ff603c1c2541e8351e21aa3bf7f90f60ff72ceea603d0de8f71310b436591d774f3a680c"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  

echo - New Block Signing Key
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"entry":"000155547be7cf8b25d0c4074b01c988d703b3ce1d0e00c393930509a2616886fe98c31e8fd470013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29e3786b4496dc6b665d7b8482f72fe78ab3b347f1e8747b3a0f03f6193776de1798329dc3daee5bdfb4448ccadc463253f15d2c1d5971b4af80ca9d6b6cd85107"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00888888898fbc055f7e201b6e039bcfb859da557e10ae4d0a25b13f04accbc32700cb00010000154e657720426c6f636b205369676e696e67204b657900208888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d400207b1fa46a2178d5a19fcd200236314be0e47b20e44951f38c8e5ca5ba8b26d1f700060155547be7ce002101cda3cbde71b6277c21ca3af6be5e8038a851898aa14b6b5a78ae28ff582135e20040af2c63aae8aac7faedf91cbf3666ac16585a7e16ac6cfb9cfa7a022b99d5cfc90c835ac274c2d463205c39b33306d1827afce8148cdf99e119741dcf5488940a"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   

echo - New Bitcoin Key
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"entry":"000155547d224dedfd34c66725e03654b80a7822aa68b9b67d738c4ec4d303496434e8c18b3390013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29c8a419e3ba707818ef19e7db23aad63b1978b0b61f20196678291d0a487277a38d7d91a17dab719b06c726d8bca79d7d63ec02b901a52584df2d58a2e732ce08"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00888888898fbc055f7e201b6e039bcfb859da557e10ae4d0a25b13f04accbc32700bf000100000f4e657720426974636f696e204b657900208888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d40001000001000014c5b7fd920dce5f61934e792c7e6fcc829aff533d00060155547d224d002101cda3cbde71b6277c21ca3af6be5e8038a851898aa14b6b5a78ae28ff582135e20040ece6b7d2f92fbe5947bb05faef1fbf6bbb7d60f6b7c17c600795f9505c5e855c4cbeb707e39df696e5078183637d4c84e8c3d485c60ab977d45c6756861b5e06"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2 
echo

echo - New MHash
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"entry":"00015554808c1fce916b2cd7264a3f4608a7862c1a15eac82e3123242964c3e3ec0707c1b3e2ae013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29a3ab508dfa0ee53873ba121ff88f72a5bbefeff3a1253caf6793ac6b9f45836aaf684a548632a5746e0d33a8ac97ced29712af9bd24644ba9e2c8c04823ee30a"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00888888898fbc055f7e201b6e039bcfb859da557e10ae4d0a25b13f04accbc32700c900010000134e6577204d617472796f73686b61204861736800208888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d40020d4a7b3cdb84069eb7b299a5165678399f9af4ade2596f457224b30f526231e5b0006015554808c1e002101cda3cbde71b6277c21ca3af6be5e8038a851898aa14b6b5a78ae28ff582135e20040d09daca80fd1fbd3ea5660f858c99503088b4af33011cc8be330ab82670137cc71b16bcecd729eeb90e7f921e44c5780d79b59eb417046b38721879208de2101"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo  

echo  Identity Chain Created, info:
echo  Private Keys:
echo  Level 1: sk13iVjQEF2QNxgVf53x9AqxDFAZ52Y5bDUz47Z9qJLSup1A2WhxJ
echo  Level 2: sk22HwURovN6kz1kkW2t9svn7MXb4g8VZxj4D58SNRdhyMtvoh3Hh
echo  Level 3: sk34FdP7Xa5H6295SH4Ajebipe45LimFH8cKbYpwXJHMUBmkVFQWP
echo  Level 4: sk33ZM8yKNgtjrpBkDxPFQPmU3rxgsGxXqPJ9NZaE9GrLuEwungUE
echo  
echo  ChainIDs:
echo  Main Factom Identity List ChainID: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
echo  Root: 8888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d4
echo  Sub \: 888888898fbc055f7e201b6e039bcfb859da557e10ae4d0a25b13f04accbc327
echo
echo  Block Signing Key: e4f29f4d4abcf4510de3bd3424832210
echo  BtcKey: 3kmKYEeCcBHa1KMwSg55z8vm4HZe
echo  
echo  MHash Seed: 8888884ed5d23aa2086865bdf02f321d38d98c4717dfa1a78c030d12241d53d4
echo  MHash: d4a7b3cdb84069eb7b299a5165678399f9af4ade2596f457224b30f526231e5b

