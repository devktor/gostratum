## gostratum
Stratum (Electrum) client implementation in GOLANG. 
###documentation
[API Reference](doc/api.md)<br />
[Bitcoin Address Watch example](examples/address_watch.go)<br />
###status
Package is not tested yet, is not recommended to use it in production at this point
###installation
go get https://github.com/devktor/gostratum
###concurency
It is safe to use Client object in multiple goroutines.
###protocol compliance
Implements protocol specifications according to https://docs.google.com/document/d/17zHy1SUlhgtCMbypO8cHgpWH73V5iUQKk_0rWvMqSNs/edit?hl=en_US


