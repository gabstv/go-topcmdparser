# WIP

#### linux
top -b -n 1 -d 00.10

#### darwin
top -l 1 -s 1



```
Processes: 280 total, 2 running, 14 stuck, 264 sleeping, 1585 threads 
2015/10/11 22:05:58
Load Avg: 1.48, 1.46, 1.49 
CPU usage: 4.20% user, 13.44% sys, 82.35% idle 
SharedLibs: 11M resident, 15M data, 0B linkedit.
MemRegions: 83646 total, 3249M resident, 62M private, 980M shared.
PhysMem: 8070M used (1427M wired), 116M unused.
VM: 1946G vsize, 1066M framework vsize, 4864(0) swapins, 6656(0) swapouts.
Networks: packets: 4300449/3935M in, 4169449/911M out.
Disks: 2433038/52G read, 902638/24G written.

PID    COMMAND          %CPU TIME     #TH   #WQ #PORTS MEM    PURG   CMPRS  PGRP  PPID  STATE    BOOSTS     %CPU_ME %CPU_OTHRS UID FAULTS    COW       MSGSENT   MSGRECV   SYSBSD     SYSMACH    CSW        PAGEINS IDLEW    POWER USER            #MREGS RPRVT VPRVT VSIZE KPRVT KSHRD
97124  Google Chrome He 0.0  00:06.26 12    0   97+    37M+   0B     73M+   413   413   sleeping *0[26+]    0.00000 0.00000    501 110496+   1470+     19614+    4312+     176097+    464369+    75979+     46319+  3064     0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
94896  MagicPrefsPlugin 0.0  00:00.68 5     0   406+   21M+   996K+  31M+   94896 1     sleeping *0[169+]   0.00000 0.00000    501 18513+    356+      5800+     2144+     12492+     9838+      6337+      200+    235      0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
94893  goboots          0.0  00:37.53 8     0   26+    2668K+ 0B     1260K+ 94893 6137  sleeping *0[1+]     0.00000 0.00000    501 5755+     711+      36+       14+       4429501+   542556+    4699303+   370+    1308891  0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
94838  pia_openvpn      0.0  00:34.37 1     0   8+     728K+  0B     108K+  94838 1     sleeping *0[1+]     0.00000 0.00000    0   885+      142+      28+       11+       2554781+   31+        680306+    0       685      0.0   root            N/A    N/A   N/A   N/A   N/A   N/A  
92041  Google Chrome He 0.0  00:02.10 11    0   96+    9784K+ 0B     25M+   413   413   sleeping *0[24+]    0.00000 0.00000    501 43874+    1443+     7493+     1122+     76953+     308426+    27718+     23006+  652      0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
77948  Google Chrome He 0.0  00:10.55 11    0   99+    58M+   0B     22M+   413   413   sleeping *0[30+]    0.00000 0.00000    501 316229+   1460+     11920+    3954+     235520+    389193+    77430+     242988+ 1522     0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
77771  Google Chrome He 0.0  00:23.30 12    0   97+    48M+   0B     25M+   413   413   sleeping *0[26+]    0.00000 0.00000    501 170402+   1482+     84814+    25727+    516770+    1508292+   235421+    82157+  15713    0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
77058  Google Chrome He 0.0  00:10.31 13    0   98+    173M+  0B     13M+   413   413   sleeping *0[26+]    0.00000 0.00000    501 252100+   1470+     19245+    3238+     224386+    486673+    79687+     92086+  1661     0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
76766  WorkflowServiceR 0.0  00:00.10 3     0   130+   700K+  0B     5128K+ 76766 1     sleeping *0[10+]    0.00000 0.00000    501 5207+     265+      1933+     786+      3244+      2715+      1197+      303+    12       0.0   gabs            N/A    N/A   N/A   N/A   N/A   N/A  
```