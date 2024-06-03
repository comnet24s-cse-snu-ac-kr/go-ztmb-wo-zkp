# ZTMB (Zero Tunneling MiddleBox) w/o zkp

- Logic implementation without ZKP written in go

## Quickstart

### Demo

- Build and run an example input:

```bash
make
./ztmb ./example/input.json
```

#### Output example (AES-256-GCM)

- Generated JSON file (name: `result.json`):

```json
{"key":["0","1","2","3","4","5","6","7","8","9","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24","25","26","27","28","29","30","31"],"nonce":["0","1","2","3","4","5","6","7","8","9","10","11"],"packet":["36","215","1","0","0","1","0","0","0","0","0","1","63","98","87","70","106","76","84","89","48","81","71","57","119","90","87","53","122","99","50","103","117","89","50","57","116","76","72","86","116","89","87","77","116","77","84","73","52","81","71","57","119","90","87","53","122","99","50","103","117","89","50","57","116","76","71","104","116","89","87","77","116","99","50","104","63","104","77","105","48","121","78","84","89","115","97","71","49","104","89","121","49","122","97","71","69","121","76","84","85","120","77","105","120","111","98","87","70","106","76","88","78","111","89","84","69","115","97","71","49","104","89","121","49","116","90","68","85","116","90","88","82","116","81","71","57","119","90","87","63","53","122","99","50","103","117","89","50","57","116","76","71","104","116","89","87","77","116","99","109","108","119","90","87","49","107","77","84","89","119","76","87","86","48","98","85","66","118","99","71","86","117","99","51","78","111","76","109","78","118","98","83","120","111","98","87","70","106","76","88","78","111","89","20","84","69","116","79","84","89","116","90","88","82","116","81","71","57","119","90","87","53","122","99","1","56","1","102","1","49","5","49","51","57","52","48","6","116","117","110","110","101","108","7","101","120","97","109","112","108","101","3","111","114","103","0","0","5","0","1","0","12","0","247","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0","0"],"ciphertext":["99","213","215","27","197","228","194","27","141","65","151","138","142","139","15","43","233","186","211","77","192","42","24","69","79","61","178","176","71","10","50","245","116","73","156","197","219","173","90","238","32","221","8","160","252","202","92","81","218","8","7","180","13","140","212","239","69","244","24","126","109","186","199","215","132","80","129","50","135","136","113","12","104","55","152","38","208","231","67","161","8","243","167","54","47","16","18","47","204","181","199","181","73","138","216","234","174","129","53","178","241","57","133","83","198","83","92","229","46","63","109","39","244","40","59","223","134","150","68","221","124","183","14","180","179","146","52","74","238","38","90","190","230","131","186","52","49","150","50","246","2","239","22","129","4","147","150","76","7","40","231","38","203","23","188","14","173","159","157","141","222","219","201","111","158","206","167","51","243","58","138","150","168","214","182","217","200","172","188","65","204","221","38","255","86","145","175","139","60","171","225","145","96","20","197","152","122","187","156","133","79","56","188","191","174","53","156","208","174","10","129","2","71","120","232","165","0","246","163","144","27","31","186","128","32","34","236","174","128","33","50","32","99","116","129","60","133","23","8","4","230","221","220","57","154","85","209","228","122","27","43","54","136","211","48","230","20","209","247","144","156","218","68","62","128","223","114","7","183","45","58","10","218","180","202","205","128","139","129","175","113","47","109","81","10","244","186","75","1","3","222","72","186","197","206","23","146","60","58","90","10","149","160","211","152","202","91","94","73","148","207","23","152","49","29","22","151","61","69","233","201","149","127","42","139","96","120","233","163","224","123","32","93","211","29","98","147","26","215","132","156","20","133","8","73","163","120","102","141","177","72","255","89","47","249","217","133","84","114","203","32","175","164","35","151","181","98","40","132","253","31","121","238","241","0","193","159","229","52","220","220","87","89","50","165","131","192","217","135","7","211","220","161","112","254","127","62","64","18","135","226","11","108","187","159","32","82","212","61","16","222","70","77","230","187","86","25","122","87","107","195","209","81","114","160","184","220","82","255","248","176","165","93","107","68","18","28","131","32","229","231","185","227","237","134","171","31","120","122","122","174","221","174","87","191","125","235","26","224","87","241","148","213","81","105","138","107","82","195","154","30","60","67","197","186","176","17","255","134","180","25","239","41","1","151","147","223","206","81","170","4","141","232","172","237","128","144","172","209","154","61","174","225","32","183","2","145","93","199","232","83","28"],"counter":["0","0","0","1"]}
```

- STDOUT:

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
Cipher (AES-256-GCM)
  Key:                    000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f
  Nonce:                  000102030405060708090a0b
  PreCounterBlockSuffix:  00000001
  Length:                 512
  Tag:                    bb876078ec94bd58e46b494b0a1a177a
  Hex:                    63d5d71bc5e4c21b8d41978a8e8b0f2be9bad34dc02a18454f3db2b0470a32f574499cc5dbad5aee20dd08a0fcca5c51da0807b40d8cd4ef45f4187e6dbac7d7845081328788710c68379826d0e743a108f3a7362f10122fccb5c7b5498ad8eaae8135b2f1398553c6535ce52e3f6d27f4283bdf869644dd7cb70eb4b392344aee265abee683ba34319632f602ef16810493964c0728e726cb17bc0ead9f9d8ddedbc96f9ecea733f33a8a96a8d6b6d9c8acbc41ccdd26ff5691af8b3cabe1916014c5987abb9c854f38bcbfae359cd0ae0a81024778e8a500f6a3901b1fba802022ecae802132206374813c85170804e6dddc399a55d1e47a1b2b3688d330e614d1f7909cda443e80df7207b72d3a0adab4cacd808b81af712f6d510af4ba4b0103de48bac5ce17923c3a5a0a95a0d398ca5b5e4994cf1798311d16973d45e9c9957f2a8b6078e9a3e07b205dd31d62931ad7849c14850849a378668db148ff592ff9d9855472cb20afa42397b5622884fd1f79eef100c19fe534dcdc575932a583c0d98707d3dca170fe7f3e401287e20b6cbb9f2052d43d10de464de6bb56197a576bc3d15172a0b8dc52fff8b0a55d6b44121c8320e5e7b9e3ed86ab1f787a7aaeddae57bf7deb1ae057f194d551698a6b52c39a1e3c43c5bab011ff86b419ef29019793dfce51aa048de8aced8090acd19a3daee120b702915dc7e8531c
```

- Input/Output JSON file examples are available in [/example](./example) dir

### Testing

```bash
make test
```

### Pre-built binaries

- Pre-built binaries are available in [/build](./build) dir
    - `build/ztmb-darwin-amd64`: x86_64 binary
    - `build/ztmb-darwin-arm64`: ARM64 binary
