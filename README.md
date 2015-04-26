# Lock
Lock algorithm in go

`filter/` Peterson's Filter Lock

`flag/` Flag Lock

`measure/` measure overtaken threads, when using filter lock

The algorithm chosen are more of a theoretical nature. 
Real-World implementations of lock algorithm uses special instruction of the CPU
like test-and-set or compare-and-swap to gain more efficiency.
