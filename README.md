# openfaas-luhn
Luhn Algorithm as open faas function (https://en.wikipedia.org/wiki/Luhn_algorithm).

Luhn Algorithm is a simple checksum formula used to validate a variety of identification numbers, such as credit card numbers, IMEI numbers, National Provider Identifier numbers in the United States, Canadian Social Insurance Numbers, Israel ID Numbers and Greek Social Security Numbers (ΑΜΚΑ).


### Usage

Deploy:

```bash
$ faas-cli deploy -f ./luhn.yml --gateway=http://<GATEWAY-IP> 
```

Invoke: 
```bash
$ echo -n '4539 1488 0343 6467' | faas-cli invoke luhn --gateway=<GATEWAY-IP>
```

### Example
![True](https://github.com/viveksyngh/openfaas-luhn/blob/master/screens/luhn-true.png)

![False](https://github.com/viveksyngh/openfaas-luhn/blob/master/screens/luhn-false.png)
