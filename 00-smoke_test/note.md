In order to run this, you have two options:

1. For local testing, you can run it with: 
  ```fish
  env IP="127.0.0.1" PORT="8000" go run main.go    
  ```

2. For actual testing (i.e., accessing it outside the local network), use:
  ```fish
  evn IP="0.0.0.0" PORT="8000" go run main.go    
  ```

"127.0.0.1" is an address assigned to the loopback (local-only) interface,
hence, if you try to test it, it will not be visible to the testing service.
