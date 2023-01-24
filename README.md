# series

common `int` number series: flipper, linear, square, random, fibonacci...

```
$ cat main.go
package main

import (
        "fmt"

        "github.com/btwiuse/series/fibonacci"
        "github.com/btwiuse/series/flipper"
        "github.com/btwiuse/series/linear"
        "github.com/btwiuse/series/random"
        "github.com/btwiuse/series/square"
)

func main() {
        for i := 0; i < 100; i++ {
                fmt.Println(
                        flipper.Value(), "\t",
                        linear.Value(), "\t",
                        square.Value(), "\t",
                        random.Value(), "\t",
                        fibonacci.Value(), "\t",
                )
        }
}
```

```
$ go run main.go
0        0       0       5577006791947779410     0 
1        1       1       8674665223082153551     1 
0        2       4       6129484611666145821     1 
1        3       9       4037200794235010051     2 
0        4       16      3916589616287113937     3 
1        5       25      6334824724549167320     5 
0        6       36      605394647632969758      8 
1        7       49      1443635317331776148     13 
0        8       64      894385949183117216      21 
1        9       81      2775422040480279449     34 
0        10      100     4751997750760398084     55 
1        11      121     7504504064263669287     89 
0        12      144     1976235410884491574     144 
1        13      169     3510942875414458836     233 
0        14      196     2933568871211445515     377 
1        15      225     4324745483838182873     610 
0        16      256     2610529275472644968     987 
1        17      289     2703387474910584091     1597 
0        18      324     6263450610539110790     2584 
1        19      361     2015796113853353331     4181 
0        20      400     1874068156324778273     6765 
1        21      441     3328451335138149956     10946 
0        22      484     5263531936693774911     17711 
1        23      529     7955079406183515637     28657 
0        24      576     2703501726821866378     46368 
1        25      625     2740103009342231109     75025 
0        26      676     6941261091797652072     121393 
1        27      729     1905388747193831650     196418 
0        28      784     7981306761429961588     317811 
1        29      841     6426100070888298971     514229 
0        30      900     4831389563158288344     832040 
1        31      961     261049867304784443      1346269 
0        32      1024    1460320609597786623     2178309 
1        33      1089    5600924393587988459     3524578 
0        34      1156    8995016276575641803     5702887 
1        35      1225    732830328053361739      9227465 
0        36      1296    5486140987150761883     14930352 
1        37      1369    545291762129038907      24157817 
0        38      1444    6382800227808658932     39088169 
1        39      1521    2781055864473387780     63245986 
0        40      1600    1598098976185383115     102334155 
1        41      1681    4990765271833742716     165580141 
0        42      1764    5018949295715050020     267914296 
1        43      1849    2568779411109623071     433494437 
0        44      1936    3902890183311134652     701408733 
1        45      2025    4893789450120281907     1134903170 
0        46      2116    2338498362660772719     1836311903 
1        47      2209    2601737961087659062     2971215073 
0        48      2304    7273596521315663110     4807526976 
1        49      2401    3337066551442961397     7778742049 
0        50      2500    8121576815539813105     12586269025 
1        51      2601    2740376916591569721     20365011074 
0        52      2704    8249030965139585917     32951280099 
1        53      2809    898860202204764712      53316291173 
0        54      2916    9010467728050264449     86267571272 
1        55      3025    685213522303989579      139583862445 
0        56      3136    2050257992909156333     225851433717 
1        57      3249    6281838661429879825     365435296162 
0        58      3364    2227583514184312746     591286729879 
1        59      3481    2873287401706343734     956722026041 
0        60      3600    8603989663476771718     1548008755920 
1        61      3721    6842348953158377901     2504730781961 
0        62      3844    7388428680384065704     4052739537881 
1        63      3969    6735196588112087610     6557470319842 
0        64      4096    1687184559264975024     10610209857723 
1        65      4225    3950896730125624717     17167680177565 
0        66      4356    8273290538659802269     27777890035288 
1        67      4489    6296367092202729479     44945570212853 
0        68      4624    9029029644282286269     72723460248141 
1        69      4761    8505906760983331750     117669030460994 
0        70      4900    837825985403119657      190392490709135 
1        71      5041    4548432111829895923     308061521170129 
0        72      5184    8549944162621642512     498454011879264 
1        73      5329    8807817071862113702     806515533049393 
0        74      5476    3209308858241334655     1304969544928657 
1        75      5625    6371863560482907257     2111485077978050 
0        76      5776    6556961545928831643     3416454622906707 
1        77      5929    5199948958991797301     5527939700884757 
0        78      6084    5990482929064819019     8944394323791464 
1        79      6241    5089134323978233018     14472334024676221 
0        80      6400    6971241403795498694     23416728348467685 
1        81      6561    3724427934598140041     37889062373143906 
0        82      6724    1205043859388862788     61305790721611591 
1        83      6889    9093919513921919021     99194853094755497 
0        84      7056    8267293389953062911     160500643816367088 
1        85      7225    2970700287221458280     259695496911122585 
0        86      7396    6651414131918424343     420196140727489673 
1        87      7569    5944830206637008055     679891637638612258 
0        88      7744    788787457839692041      1100087778366101931 
1        89      7921    6175742077372812453     1779979416004714189 
0        90      8100    5743654948930018631     2880067194370816120 
1        91      8281    3409814636252858217     4660046610375530309 
0        92      8464    2184302455902443631     7540113804746346429 
1        93      8649    4937104021912138218     -6246583658587674878 
0        94      8836    1727040455672546632     1293530146158671551 
1        95      9025    2202916659517317514     -4953053512429003327 
0        96      9216    5793183108815074904     -3659523366270331776 
1        97      9409    1169089424364679180     -8612576878699335103 
0        98      9604    2594813965004488500     6174643828739884737 
1        99      9801    3784560248718450071     -2437933049959450366
```