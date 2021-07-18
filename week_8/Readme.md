# week 8



| 数据大小 | 操作 | p95(ms) | p99(ms) | avg(ms) |
| -------- | ---- | ------- | ------- | ------- |
| 10bytes  | get  | 0.279   | 0.399   | 0.179   |
| 10bytes  | set  | 0.271   | 0.399   | 0.186   |
| 20bytes  | get  | 0.303   | 0.527   | 0.195   |
| 20bytes  | set  | 0.279   | 0.423   | 0.185   |
| 50bytes  | get  | 0.295   | 0.399   | 0.184   |
| 50bytes  | set  | 0.279   | 0.343   | 0.199   |
| 100bytes | get  | 0.287   | 0.431   | 0.193   |
| 100bytes | set  | 0.271   | 0.343   | 0.185   |
| 200bytes | get  | 0.279   | 0.375   | 0.187   |
| 200bytes | set  | 0.279   | 0.351   | 0.186   |
| 1k       | get  | 0.303   | 0.399   | 0.195   |
| 1k       | set  | 0.279   | 0.375   | 0.190   |
| 5k       | get  | 0.287   | 0.439   | 0.203   |
| 5k       | set  | 0.287   | 0.351   | 0.179   |

value size 10 bytes, before usage: 905240 bytes, after usage 1676312 bytes , perusage: 77.107200 bytes  
value size 20 bytes, before usage: 905240 bytes, after usage 1756312 bytes , perusage: 85.107200 bytes  
value size 50 bytes, before usage: 905240 bytes, after usage 2076312 bytes , perusage: 117.107200 bytes  
value size 100 bytes, before usage: 905240 bytes, after usage 2636312 bytes , perusage: 173.107200 bytes  
value size 200 bytes, before usage: 905240 bytes, after usage 3756312 bytes , perusage: 285.107200 bytes  
value size 1024 bytes, before usage: 905240 bytes, after usage 14316312 bytes , perusage: 1341.107200 bytes  
value size 5120 bytes, before usage: 905240 bytes, after usage 62956312 bytes , perusage: 6205.107200 bytes  
