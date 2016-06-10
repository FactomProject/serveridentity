
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
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"message":"0001553b8dd6a0fcad13f97dd3efdc89d5eb5e67e254f4ff5f98f03d8c5e7bb86ba5dc730faf37c3f5294b202a42711f57e69a809e90f689cef6d568223b1e2963e500af23b1ef68804348f4804b1d422947a2b415cbd5f352767d931acaead68ab6a21135b6680b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29f0d4275600c294113245c4a393751bb8045ca05a7d874f2992da40b58a32734e5ff367d74e87a3ef1c6c962c8984ab2e717658cc729b0f5cb484129338862005"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"0088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f009f000100000e4964656e7469747920436861696e00204d9a880e8a437c97176f897b5b7a11c2f2abe2b16caf99d1412cd760f6cd0cab00208959030a14476c3b966eab52a94701ddcb82679f0452ee0470a88a219f381cec0020c0791b0a54501a2a6ae737668d36a0dddb5a1677be371aaee2567275017e0a40002012a0ae3b2a138c401d66e2fa67e592bbf57350da4e8aaedf53237f1e8090e25a00023239"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Registering Identity Chain
curl -X POST --data '{"jsonrpc":"2.0","id":3,"params":{"entry":"0001553b8dd6a13436ba7abebe32654dfe95980e56dbb21741adb78f3862f9172f5e0a83c896dc013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29f995e3381044404d9ac267effa24115e6c6e4b86f967e7afa432c09681d747f3d097a188db5dd9a86e4317e755a72bd1722e325b1e3442b2eb2793f862cb7d0a"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":4,"params":{"entry":"00e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85500a40001000018526567697374657220466163746f6d204964656e74697479002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f0021017e5c9951acb6c286f2496e162381537685fd58300c97c4422211441694ed939e004022f00670687e86495139c1bca3f62923aa8aedec257b39d10b0d0d018294eb991c93277c5f879794dfe9d307c9ef442642f455add282a1226acdae6555be7c03"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Creating Identity SubChain - ChainID: 883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c
curl -X POST --data '{"jsonrpc":"2.0","id":5,"params":{"message":"0001553b8dd6a2401109bbc183ff89236bebe6fdea8c270f55510a66d8a87c792a1457a8c9d0611c5b1ab6b0d3e57e12adac39ebd41abcf632caf1bce93f88c631190a09f5f1c548cd34781acfd5e9346c5233b39d7bbbfa175a64f6cf5e8a243f5201bb4a099c0b3b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da299d43fe453de70899a304afdcfaa918f5a8ba6d806d13ab2f8329ee8ae2a01c1301903dfd625e243db3abf7f3d57e8a7a684706f94c4b4402d3b6f4df918e2006"},"method":"commit-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":6,"params":{"entry":"00883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c003c0001000011536572766572204d616e6167656d656e74002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f00023730"},"method":"reveal-chain"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo - Registering Identity SubChain
curl -X POST --data '{"jsonrpc":"2.0","id":7,"params":{"entry":"0001553b8dd6a3434510c46577590690482b53f185aaa380c9049c437e4d6fc4cb449218a7d28b013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da293458f5c19d1a6eb112527d2270248d27de731b5ff52c9b22c769013dbfc2caed1e596105b671d6ced82eabc6c5d57a8e00ad4ed9f979541434ffd14156763903"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":8,"params":{"entry":"0088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f00a6000100001a526567697374657220536572766572204d616e6167656d656e74002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f0021017e5c9951acb6c286f2496e162381537685fd58300c97c4422211441694ed939e0040cbfb07a03b382475950e7d1031fcd5630b07255c01b351e44fa6ba92aadfe5de8da517e1254155e3e9b00a40ab8549fc0a429d214e07429c865665615fd5c404"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo   

echo - New Bitcoin Key
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"entry":"0001553b8f913808bdb02762bc55da9d8b1078343b3df6f99923b3d8fecab91071b11fd047820a013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da292c55e8f886fa6a60d59963118bc5bef0cd7f820fea9943185e3e78b2d44bd96482873cb9a1f4cf5c05486fb55519caf8cef84866f5cff1f7c72e5b78d8cf3303"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c00cb00010000154e657720426c6f636b205369676e696e67204b6579002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f00205a51c3eff4768513b6eeb3e10eeafb7a6832ba7e8dfb009e0aeabc05d584087c000601553b8f913700210194e72f2ec74ca123162369cc99d77168fefbc56963dd408757530a9b00ea89e70040cf9daeb576d40947e2002126c156266dc2aa4ba6fa94383d6831064fab4b65df821d6dd495aa2acda40feacf8c1a57264f61d88897769d1cb24b98725f473107"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo   

echo - New Bitcoin Key
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"entry":"0001553b90eb0bea2b5df18c7c04fcac6fdcbfcabebbcd7fc7552e2562a48c37c4c4647fd3827b013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da2911d68f947788f7aa77983af6de083492f5379e2dc1d69af5b4ed38ab448f12a094b6f3121eadd450cc49b9be9018ad7bfd890db8e9527fe1d8ea295fd4d59b03"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c00bf000100000f4e657720426974636f696e204b6579002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f0001000001000014c5b7fd920dce5f61934e792c7e6fcc829aff533d000601553b90eb0a00210194e72f2ec74ca123162369cc99d77168fefbc56963dd408757530a9b00ea89e700402d92ddd5f23817fdb194fe52785eeb1f9a2eb01dc402a462aa4519acae58fb51ff18c049f7d3a4373ff9ee9361b934516cd9c62d461b86cb234e1d5afec38e05"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo

echo - New Bitcoin Key
curl -X POST --data '{"jsonrpc":"2.0","id":1,"params":{"entry":"0001553b91a66cd8d9b51cd7ca3389da90d5f63d687b031d9de571a81fa875bd8b3efa8d965205013b6a27bcceb6a42d62a3a8d02a6f0d73653215771de243a63ac048a18b59da29694a0dc4c5439842e7c473f5285c9baf5d25543ed07c5329375dc251ea08aec3e73874d1410fdb65bef315696a6447af11ac8267b7f5e931f404646467138c0c"},"method":"commit-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
curl -X POST --data '{"jsonrpc":"2.0","id":2,"params":{"entry":"00883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c00c900010000134e6577204d617472796f73686b612048617368002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f002088598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f000601553b91a66b00210194e72f2ec74ca123162369cc99d77168fefbc56963dd408757530a9b00ea89e70040658f06a18dc564e69ac91c3d6525e331ae517d43a38b817db3d539c575832477a43ce417c78c78aaa67ad55ffec5bb2e793cdd070b75bd157c8c51be844e4d04"},"method":"reveal-entry"}' -H 'content-type:text/plain;' http://localhost:8088/v2
echo   
echo

echo  Identity Chain Created, info:
echo  Private Keys:
echo  Level 1: sk13MMUdir8DnMzpC836252fCv2zWbVaJcSXh8MCz7XzECd1SB2zg
echo  Level 2: sk23Utfzt1sX83ga2VBWgA9m24D6fJbYFJAznysa3CUPvoJ7mWZeo
echo  Level 3: sk34FdP7Xa5H6295SH4Ajebipe45LimFH8cKbYpwXJHMUBmkVFQWP
echo  Level 4: sk4431HNNAEVEzabTPAmYy1Xkvxv3ggzyS3kwYzjPoxfheMgKYXfp
echo  
echo  ChainIDs:
echo  Main Factom Identity List ChainID: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
echo  Root: 88598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f
echo  Sub \: 883111484c5f0f7213c7c1fdd32117e68b492b6dee6cb36f4f5a88c2528f646c
echo
echo  Block Signing Key: 670e0d158c1f8c171222a18742d63ed6
echo  BtcKey: 3kmKYEeCcBHa1KMwSg55z8vm4HZe
echo  
echo  MHash Seed: 88598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f
echo  MHash: 88598ef076ce7deedb93756634cede03fcb2abfa168822bc94a93b80a2a7b43f

