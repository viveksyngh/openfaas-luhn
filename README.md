# openfaas-luhn
Luhn Algorithm as open faas function (https://en.wikipedia.org/wiki/Luhn_algorithm)

### Usage

Deploy:

```bash
$ faas-cli deploy -f ./luhn.yml --gateway=http://<GATEWAY-IP> 
```

Invoke: 
```bash
$ echo -n '4539 1488 0343 6467' | faas-cli invoke luhn --gateway=<GATEWAY-URL>
```

### Example
![True](https://github.com/viveksyngh/openfaas-luhn/blob/master/screens/luhn-true.png)

![False](https://github.com/viveksyngh/openfaas-luhn/blob/master/screens/luhn-false.png)
