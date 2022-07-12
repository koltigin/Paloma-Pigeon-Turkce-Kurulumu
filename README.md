# Paloma Testnet-6 Pigeon Yazılımı Türkçe Kurulum Rehberi

![Logo!](assets/paloma.png)

# Pigeon

> Paloma doğrulayıcılarının mesajları herhangi bir blok zincirine iletmeleri için bir Golang zincirler arası mesaj aktarma sistemidir.

Paloma, herhangi bir blok zincirinde birden fazla akıllı sözleşmenin eşzamanlı kontrolünü isteyen Crosschain yazılım mühendisleri için,
herhangi bir veri kaynağıyla ölçeklenebilir, zincirler arası, akıllı sözleşme yürütmesine olanak tanıyan, merkezi olmayan ve fikir birliğine dayalı mesaj teslimi, hızlı durum farkındalığı, düşük maliyetli durum hesaplaması ve güçlü doğrulama sistemidir.


## Pigeon Yükleme

```shell
wget -O - https://github.com/palomachain/pigeon/releases/download/v0.2.5-alpha/pigeon_0.2.5-alpha_Linux_x86_64v3.tar.gz | \
tar -C /usr/local/bin -xvzf - pigeon
chmod +x /usr/local/bin/pigeon
mkdir ~/.pigeon
```

## EVM (ETH) Cüzdan

### Yeni Cüzdan Oluşturma
* Şifrenizi ve mnemoniclerinizi unutmayınız, kaydediniz.
```
pigeon evm keys generate-new ~/.pigeon/keys/evm/eth-main
```

### Mevcut Cüzdanı İçeri Aktarma
* Cüzdanınıza ait menemonicleriniz var ise bu kod ile cüzdanınızı içeri aktarınız.
* Ben Paloma formunda verdiğim ETH adresimi içeri aktaracağım.

```
pigeon evm keys import ~/.pigeon/keys/evm/eth-main
```

### Yapılandırma Kurulumu
* Burada Paloma cüzdanımızı içeri aktarıyoruz.
* VALIDATOR_ADINIZ kısmına Paloma validator adınızı yazınız.
```shell
palomad keys add VALIDATOR_ADINIZ --recover`
```
Yukarıdaki kodu girdiğinizde çöyle bir çıktı alacaksınız `override the existing name VALIDATOR_ADINIZ [y/N]:` buna yes yani y diyerek devam ediyoruz. Ardından sizden > `Enter your bip39 mnemonic` cüzdanınıza ait menemonicleri isteyecek onları yazıp işleme devam ediyoruz.

VALIDATOR env değişkenini ayarlama
```shell
export VALIDATOR="$(palomad keys list --list-names | head -n1)"
```
* Buradan sonraki aşamaya geçmeden önce [Alchemy](https://alchemy.com/?r=zc3NjI5NzM1NzMxN)'den bir hesap oluşturup ETH Mainnet App oluşturuyoruz. Burada 'View Key' bölümünden 'https' ile başlayan linkimizi alıyoruz.


* Yapılandırma dosyasını oluşturma `~/.pigeon/config.yaml`
Burada belirtilen dizine config.yaml dosyası oluşturuyoruz. Aşağıdaki kodu kendimize göre düzenleyerek tek seferde çalıştırıyoruz.
	* PALOMA_KEYRING_PASS ve ETH_PASS değiştirmeden bırakıyoruz.
	* VALIDATOR_ADINIZ bu bölüme validator adımızı yazıyoruz.
	* ETH_KEYINIZ bölümüne cüzdan oluşturduğunuzda sonda rakamlardan oluşan kodu ekliyoruz. Bu kod aslında cüzdan adresinizin 0x bölümü olamadanki hali.
	

```yaml
cat <<EOT >~/.pigeon/config.yaml

loop-timeout: 5s

paloma:
  chain-id: paloma-testnet-6
  call-timeout: 20s
  keyring-dir: ~/.paloma
  keyring-pass-env-name: PALOMA_KEYRING_PASS
  keyring-type: test
  signing-key: VALIDATOR_ADINIZ
  base-rpc-url: http://localhost:26657
  gas-adjustment: 1.5
  gas-prices: 0.001ugrain
  account-prefix: paloma

evm:
  eth-main:
    chain-id: 1
    base-rpc-url: ALCHEMY_LINKINIZ
    keyring-pass-env-name: ETH_PASS
    signing-key: ETH_KEYINIZ
    keyring-dir: ~/.pigeon/keys/evm/eth-main
EOT
```

### Pigeon'u Başlatma

İlk güvercinimiz birkaç anahtara ihtiyacı vardır, bunun için 'env.sh' doyası oluşturuyoruz:

```shell
cat <<EOT >~/.pigeon/env.sh
PALOMA_KEYRING_PASS=PALOMA_CUZDAN_SIFRENIZ
ETH_RPC_URL=ALCHEMY_LINKINIZ
ETH_PASSWORD=ETH_CUZDAN_SIFRENIZ
ETH_SIGNING_KEY=ETH_HEX_KEYINIZ
VALIDATOR=VALIDATOR_ADINIZ
EOT
```

* Sonra aşağıdaki kodla güvercini çalıştırıyoruz:

```shell
source ~/.pigeon/env.sh
pigeon start
```

#### Servis Dosyası Oluşturma


```shell
cat <<EOT >/etc/systemd/system/pigeond.service
[Unit]
Description=Pigeon daemon
After=network-online.target
ConditionPathExists=/usr/local/bin/pigeon

[Service]
Type=simple
Restart=always
RestartSec=5
User=${USER}
WorkingDirectory=~
EnvironmentFile=${HOME}/.pigeon/env.sh
ExecStart=/usr/local/bin/pigeon start
ExecReload=

[Install]
WantedBy=multi-user.target
EOT
```

Son olarak güvercini başlatıyoruz.

```shell
service pigeond start
```

#### Servis Durumunu Görme:
```
service pigeond status
```
#### Logları İzleme:
```
journalctl -u pigeond.service -f -n 100
```

### Pigeon Değişkenlerinin Tanımları ve Açıklamaları
* Bu bölümü daha sonra Türkçeleştireceğim.
  - for paloma key:
	- keyring-dir
      - right now it's not really super important where this points. The important things for the future is that pigeon needs to send transactions to Paloma using it's validator (operator) key!
	  - it's best to leave it as is
	- keyring-pass-env-name
	  - this one is super important!
	  - it is the name of the ENV variable where password to unlock the keyring is stored!
	  - you are not writing password here!! You are writing the ENV variable's name where the password is stored.
	  - you should obviously use a bit more advanced method than shown here, but here is the example:
	    - if the `keyring-pass-env-name` is set to `MY_SUPER_SECRET_PASS` then you should provide ENV variable `MY_SUPER_SECRET_PASS` and store the password there
	    - e.g. `MY_SUPER_SECRET_PASS=abcd pigeon start`
	- keyring-type
	  - it should be the same as it's defined for paloma's client. Look under the ~/.paloma/config/client.toml
	- signing-key
	  - right now it's again not important which key we are using. It can be any key that has enough balance to submit TXs to Paloma. It's best to use the same key that's set up for the validator.
	- gas-adustment:
	  - gas multiplier. The pigeon will estimate the gas to run a TX and then it will multiply it with gas-adjustment (if it's a positive number)
 - for evm -> eth-main:
	- keyring-pass-env-name: as as above for paloma.
	- signing-key
	  - address of the key from the keyring used to sign and send TXs to EVM network (one that you got when running `pigeon evm keys generate-new` from the install section)
	- keyring-dir:
	  - a directory where keys to communicate with the EVM network is stored

