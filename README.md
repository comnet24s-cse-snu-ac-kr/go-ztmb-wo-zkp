# ZTMB (Zero Tunneling MiddleBox) w/o zkp

- Logic implementation without ZKP written in go

## Quickstart

### Demo

- Build and run an example input:

```bash
make
./ztmb ./example/input.json
```

- Generated JSON file (name: `result.json`):

```json
{"key":["0","1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26","27","28","29","30","31"],"nonce":["0","1","2","3","4","5","6","7","8","9","10","11"],"packet":["36","215","1","0","0","1","0","0","0","0","0","1","63","98","87","70","106","76","84","89","48","81","71","57","119","90","87","53","122","99","50","103","117","89","50","57","116","76","72","86","116","89","87","77","116","77","84","73","52","81","71","57","119","90","87","53","122","99","50","103","117","89","50","57","116","76","71","104","116","89","87","77","116","99","50","104","63","104","77","105","48","121","78","84","89","115","97","71","49","104","89","121","49","122","97","71","69","121","76","84","85","120","77","105","120","111","98","87","70","106","76","88","78","111","89","84","69","115","97","71","49","104","89","121","49","116","90","68","85","116","90","88","82","116","81","71","57","119","90","87","63","53","122","99","50","103","117","89","50","57","116","76","71","104","116","89","87","77","116","99","109","108","119","90","87","49","107","77","84","89","119","76","87","86","48","98","85","66","118","99","71","86","117","99","51","78","111","76","109","78","118","98","83","120","111","98","87","70","106","76","88","78","111","89","20","84","69","116","79","84","89","116","90","88","82","116","81","71","57","119","90","87","53","122","99","1","56","1","102","1","49","5","49","51","57","52","48","6","116","117","110","110","101","108","7","101","120","97","109","112","108","101","3","111","114","103","0","0","5","0","1","0","12","0","247","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0"],"ciphertext":["173","44","9","0","41","22","165","64","183","131","63","242","167","127","121","37","163","28","230","158","97","37","234","128","145","205","120","240","47","163","148","123","153","217","30","202","146","114","249","238","99","75","1","149","17","217","251","74","74","213","206","77","170","114","207","174","104","219","235","96","229","22","172","239","13","20","251","133","177","77","123","132","176","117","127","244","39","20","145","184","234","167","41","102","54","25","177","121","63","31","169","119","233","7","99","216","40","167","0","167","162","161","106","189","233","20","145","224","196","219","251","216","167","32","15","76","236","14","33","191","126","37","109","166","72","75","229","2","244","18","153","197","251","103","174","11","112","100","5","45","19","9","49","121","134","22","242","11","57","86","184","50","117","169","243","72","243","86","212","90","114","245","214","146","202","117","255","171","218","39","49","124","16","231","223","98","139","54","208","41","80","45","47","86","235","81","129","221","42","88","194","114","170","31","16","237","20","155","69","83","85","38","74","87","217","50","40","89","45","172","188","72","28","119","253","3","103","78","21","242","161","77","73","135","15","66","50","177","232","177","52","229","235","201","49","114","155","30","224","217","190","17","206","180","108","30","242","223","211","240","39","158","230","178","115","69","131","145","7","137","104","68","186","148","108","218","159","71","106","241","222","238","245","2","164","214","22","201","124","169","27","102","125","108","21","98","185","86","39","47","89","24","247","61","240","85","0","119","94","213","148","8","90","3","57","119","171","214","173","198","126","142","23","188","26","195","34","99","110","169","2","94","142","5","25","92","145","158","239","236","125","142","23","196","146","186","110","172","136","88","160","104","125","212","96","79","67","128","7","219","69","56","104","104","12","59","121","125","66","102","209","81","229","155","57","168","23","2","168","0","181","123","58","233","135","128","229","87","86","25","60","44","22","146","247","103","27","132","227","68","52","111","1","94","71","31","69","182","12","97","184","52","234","211","153","88","184","201","147","84","237","225","150","91","55","105","128","214","208","111","199","66","179","47","8","118","190","107","110","196","189","106","138","181","30","198","14","225","60","4","151","38","90","231","220","81","188","22","94","11","48","220","152","28","173","130","48","149","72","88","10","29","245","6","112","137","64","239","159","11","21","180","121","246","206","246","240","98","68","10","159","87","15","172","94","218","1","87","65","208","255","5","116","215","129","18","231","93","233","87","94","189","85","163","64","92","35","58","30","74","177","166","91","91","226","212","200","37","175","34","213","190","182","190","30","85","6","244"],"counter":["0","0","0","1"]}
```

- Example STDOUT:

```
Header
  ID:        0x24d7
  Flags:     1 0
  QDCOUNT:   0x0001
  ANCOUNT:   0x0000
  NSCOUNT:   0x0000
  ARCOUNT:   0x0001
Question #0
Question
  QNMAE:     bwFjlTy0QG9wZW5Zc2GuY29tlHvTywMtMti4Qg9WZw5zc2guY29tlGHTYWmtc2H.hMI0ynTYsAG1HYy1zageylTUxMIXObWfJLXNOYteSaG1HYy1tZDUtZxRTqg9wzW.5ZC2guy29tlGHTyWMtcMLWzw1kMTYWlwV0bUbvcGvUC3NoLMnVbsXObwfJLxnoY.tETOtytzxRTqg9WZW5zc.8.f.1.13940.tUNNeL.eXAmpLE.Org.
  QTYPE:     0x0005
  QCLASS:    0x0001
Additional Rerouces Record #0
RR OPT
  OPTCODE:  0x000c
  OPTLEN:   0x00f7
  PADDING:  0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
Cipher (Chacha20Poly1305)
  Key:                    000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f
  Nonce:                  000102030405060708090a0b
  PreCounterBlockSuffix:  00000001
  Hex:                    ad2c09002916a540b7833ff2a77f7925a31ce69e6125ea8091cd78f02fa3947b99d91eca9272f9ee634b019511d9fb4a4ad5ce4daa72cfae68dbeb60e516acef0d14fb85b14d7b84b0757ff4271491b8eaa729663619b1793f1fa977e90763d828a700a7a2a16abde91491e0c4dbfbd8a7200f4cec0e21bf7e256da6484be502f41299c5fb67ae0b7064052d130931798616f20b3956b83275a9f348f356d45a72f5d692ca75ffabda27317c10e7df628b36d029502d2f56eb5181dd2a58c272aa1f10ed149b455355264a57d93228592dacbc481c77fd03674e15f2a14d49870f4232b1e8b134e5ebc931729b1ee0d9be11ceb46c1ef2dfd3f0279ee6b27345839107896844ba946cda9f476af1deeef502a4d616c97ca91b667d6c1562b956272f5918f73df05500775ed594085a033977abd6adc67e8e17bc1ac322636ea9025e8e05195c919eefec7d8e17c492ba6eac8858a0687dd4604f438007db453868680c3b797d4266d151e59b39a81702a800b57b3ae98780e55756193c2c1692f7671b84e344346f015e471f45b60c61b834ead39958b8c99354ede1965b376980d6d06fc742b32f0876be6b6ec4bd6a8ab51ec60ee13c0497265ae7dc51bc165e0b30dc981cad82309548580a1df506708940ef9f0b15b479f6cef6f062440a9f570fac5eda015741d0ff0574d78112e75de9575ebd55a3405c233a1e4ab1a65b5be2d4c825af22d5beb6be1e5506f4
  Length:                 528
```

- Input/Output JSON file examples are available in [/example](./example) dir

### Pre-built binaries

- Pre-built binaries are available in [/build](./build) dir
    - `build/ztmb-darwin-amd64`: x86_64 binary
    - `build/ztmb-darwin-arm64`: ARM64 binary
