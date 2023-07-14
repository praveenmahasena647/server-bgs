# server-bgs

This is the server part of a service that can change your background. 
In this repo the server sends a requst to APis and gets back some random pics. And by a simple TCP connection the pics would be 
passed to the [client side](https://github.com/praveenmahasena647/client-bgs).
Also this server is bind to my [vim-config](https://github.com/praveenmahasena647/nvim) by a simple key bind you could trigger this app.

## for dev purposes
```sh
https://github.com/praveenmahasena647/server-bgs.git && cd server.bgs
make run .
```
