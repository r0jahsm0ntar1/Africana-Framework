package main

import (
    "os"
    "fmt"
    "bufio"
    "os/exec"
    "strings"
    "bcolors"
    "internals"
)

func printBanner() {
    fmt.Println(bcolors.BLUE + `                         _,._                         ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                     __.'   _)                        ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                    <_,)'.-'a\                        ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                      /' (    \                       ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `          _.-----..,-'   (''--^                       ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `         //              |                            ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `        (|   ';      ,   |                            ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `          \   ;.----/  ,/                             ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `           ) // /   | |\ \                            ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `           \ \\'\   | |/ /     Jesus Christ           ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `            \ \\ \  | |\/  The Lamb that was slain    ` + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `             '" '"  '"'        for our sins.          ` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `         __                 _____ _____     _     _   ` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `      __|  |___ ___ _ _ ___|     |  |  |___|_|___| |_ ` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `     |  |  | -_|_ -| | |_ -|   --|     |  _| |_ -|  _|` + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + `     |_____|___|___|___|___|_____|__|__|_| |_|___|_|  ` + bcolors.ENDC)
    fmt.Println("                                                                                    ")
}

func exitBanner() {
    fmt.Println(`
                   _,.---.---.---.--.._ 
               _.-' '--.'---.'---'-. _,'--.._
              /'--._ .'.     '.     ','-.'-._\
             ||   \  '.'---.__'__..-'. ,''-._/
        _  ,'\ '-._\   \    '.    '_.-'-._,''-.
     ,'   '-_ \/ '-.'--.\    _\_.-'\__.-'-.'-._'.
    (_.o> ,--. '._/'--.-',--'  \_.-'       \'-._ \
     '---'    '._ '---._/__,----'           '-. '-\
               /_, ,  _..-'                    '-._\
               \_, \/ ._(
                \_, \/ ._\
                 '._,\/ ._\
                   '._// ./'-._
                     '-._-_-_.-'
    `)
    fmt.Println(bcolors.GREEN + "\n                  wake up, Christian             " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "           Lord God Jesus Christ L❤️.VE'S you      " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "                follow the white Pigeon.           " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "                   knock, knock, knock,            " + bcolors.ENDC)
    fmt.Println(bcolors.GREEN + "                       Man Of God.               \n" + bcolors.ENDC)
}

func menu_main() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ 🍄" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Install or update africana-framework.." + bcolors.DARKCYAN + "(Start here )" + bcolors.BLUE + "🐞] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. System Security Configuration........." + bcolors.DARKCYAN + "(Setup tor &)" + bcolors.BLUE + "🐈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Local Network Attack Vectors.........." + bcolors.DARKCYAN + "(Mitm, sniff)" + bcolors.BLUE + "🐹] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Generate Undetectable Backdoors......." + bcolors.DARKCYAN + "(C2 & shells)" + bcolors.BLUE + "🐭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. WiFi Attack Vectors..................." + bcolors.DARKCYAN + "(Wifite, air)" + bcolors.BLUE + "🦝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Crack Hash, Pcap & Brute Passwords...." + bcolors.DARKCYAN + "(Hashcat, jo)" + bcolors.BLUE + "🐙] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Social-Engineering Attacks............" + bcolors.DARKCYAN + "(Gophish, gi)" + bcolors.BLUE + "🧪] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Website Attack Vectors................" + bcolors.DARKCYAN + "(Osmedeus, j)" + bcolors.BLUE + "🌍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Help, Credits, Tricks and About......." + bcolors.DARKCYAN + "(🕊️  " + bcolors.RED + "︻╦╤─" + bcolors.GREEN + "JC❤️sU" + bcolors.DARKCYAN + ")" + bcolors.BLUE + "] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit africana-framework..............." + bcolors.DARKCYAN + "(Try option " + bcolors.YELLOW + "99" + bcolors.DARKCYAN + ")" + bcolors.BLUE + "] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + `                                              ¯\_(ツ)_/¯                                                 ` + bcolors.ENDC)
}

func menu_1() {
    println(bcolors.BLUE + "\n        ~>[ 🍄" + bcolors.ENDC + bcolors.BOLD + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    println(bcolors.BLUE + "[ 1. Kali-Linux..........................." + bcolors.DARKCYAN + "(It is Stable)" + bcolors.BLUE + "🍒] " + bcolors.ENDC)
    println(bcolors.BLUE + "[ 2. Ubuntu-Linux.......................................🥝] " + bcolors.ENDC)
    println(bcolors.BLUE + "[ 3. Arch-Linux/ Black-Arch-Linux/ Manjaro-Linux..........] " + bcolors.ENDC)
    println(bcolors.BLUE + "[ 4. Uninstall-Africana...................................] " + bcolors.ENDC)
    println(bcolors.BLUE + "[ 0. Exit.................................................] \n" + bcolors.ENDC)
}

func menu_2() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ 🎭" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Install & Setup " + bcolors.RED + "Tor" + bcolors.BLUE + "......................" + bcolors.DARKCYAN + "(start here)" + bcolors.BLUE + "] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Start anonymizing through tor......................🧅]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Stop tor & restore all iptables....................🐝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Chains....." + bcolors.GREEN + "local 🐌 " + bcolors.GREEN + "squid 🐙 " + bcolors.GREEN + "privoxy 🎭 " + bcolors.GREEN + "tor 🧄 " + bcolors.GREEN + "WEB 🌍" + bcolors.BLUE + "] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Check if using " + bcolors.GREEN + "tor 🧄 " + bcolors.BLUE + ".............................🦨] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit and go back fo main menu........................] \n" + bcolors.ENDC)
}

func menu_3() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Discover Targets..................................🐹] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Port discovery on the target......................🐾] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Vuln' Scann the Target............................🦉] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. S.M.B enumration on the target....................🪳] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Exploit S.M.B on target...........................🐞] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Sniff Internal Packets............................🪰] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Lunch Responder with IPv6 Support.................🍈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Start Beefxss & Bettercap For (M.I.B).............🐝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Get Reverse shell through snifing (XSS Injection).🍇] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go to main menu..............................] \n" + bcolors.ENDC)
}

func menu_4() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Blackjack......................" + bcolors.DARKCYAN + "(All Distro...R.A.T)" + bcolors.BLUE + "🐚] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. ShellzGen......................" + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🦐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. PowerJoker....................." + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🍐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. MeterPeter....................." + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Havoc C2......................." + bcolors.DARKCYAN + "(Windows Rev Shells)" + bcolors.BLUE + "🍵] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Teardroid......................" + bcolors.DARKCYAN + "(Android 4 > 13 Rat)" + bcolors.BLUE + "🥙] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. AndroidRAT....................." + bcolors.DARKCYAN + "(Android 4 > 10 Rat)" + bcolors.BLUE + "🐭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Chameleon......................" + bcolors.DARKCYAN + "(Obfuscation .Ps1 S)" + bcolors.BLUE + "🍤] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Gh0x0st........................" + bcolors.DARKCYAN + "(Obfuscation .Ps1 S)" + bcolors.BLUE + "🍈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)
}

func menu_5() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ 📶" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. ......................Wifite............" + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🎯] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. ....................Bettercap..........." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. ...................Wifipumpkin3" + bcolors.DARKCYAN + "(Automated Phishing)" + bcolors.BLUE + "🍍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. ....................Airgeddon.............." + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. ...................wifiPumpkin3............" + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🐚] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. ..............Exit & Go To Main Menu.................] \n" + bcolors.ENDC)
}

func menu_6() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Online......................." + bcolors.DARKCYAN + "(automated_Bruteforce)" + bcolors.BLUE + "🍋] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Offline................." + bcolors.DARKCYAN + "(automated/ maunal/ Hashes)" + bcolors.BLUE + "🥭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)
}

func menu_6_1() {
    fmt.Println(bcolors.BLUE + "\n      ~>[ 🔐" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. ...................Bruteforce " + bcolors.DARKCYAN + "SSH" + bcolors.BLUE + "..................🫑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. ...................Bruteforce " + bcolors.DARKCYAN + "FTP" + bcolors.BLUE + "..................🔑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. ...................Bruteforce " + bcolors.DARKCYAN + "SMB" + bcolors.BLUE + "..................🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. ...................Bruteforce " + bcolors.DARKCYAN + "RDP" + bcolors.BLUE + "..................🍒] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. ...................Bruteforce " + bcolors.DARKCYAN + "LDAP" + bcolors.BLUE + ".................🧉] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. ...................Bruteforce " + bcolors.DARKCYAN + "SMTP" + bcolors.BLUE + ".................🌈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. ...................Bruteforce " + bcolors.DARKCYAN + "Telnet" + bcolors.BLUE + "...............🐚] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. ...................Bruteforce " + bcolors.DARKCYAN + "HTTP/S" + bcolors.BLUE + "...............🪀] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. ...................Bruteforce " + bcolors.DARKCYAN + "All/SS" + bcolors.BLUE + "...............🩴] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. ..............Exit & Go To Main Menu.................] \n" + bcolors.ENDC)
}

func menu_6_2() {
    fmt.Println(bcolors.BLUE + "\n      ~>[ 🔐" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. ....................Aircrack_ng........." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🫑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. .......................John............." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🥭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. ......................To Add.......................🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. ....................Hash-Buster....................🫑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. ..............Exit & Go To Main Menu.................] \n" + bcolors.ENDC)
}

func menu_7() {
    fmt.Println(bcolors.BLUE + "\n      ~>[ 🍄" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. ..............Gophish.(Browser Gui)." + bcolors.DARKCYAN + "(All Templetes)" + bcolors.BLUE + "🍵] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. ...............Good Ginx (Advanced)." + bcolors.DARKCYAN + "(OTP Bypass...)" + bcolors.BLUE + "🍹] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. ...................AdvPhishing......" + bcolors.DARKCYAN + "(OTP Bypass...)" + bcolors.BLUE + "🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. ....................Setoolkit......." + bcolors.DARKCYAN + "(Web Cloning..)" + bcolors.BLUE + "🧉] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. ....................Anonphisher....." + bcolors.DARKCYAN + "(OTP Bypass...)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. ....................Cyberphish......" + bcolors.DARKCYAN + "(phish temps..)" + bcolors.BLUE + "🫑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. ......................To Add.........................] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. ...............Exit & Go To Main Menu................] \n" + bcolors.ENDC)
}

func menu_8() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ 🦟" + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Start Passive Web recon & Subdomain Enumration.....🌍] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Gather e-mails & subdomain namesfrom public sources🪰] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Start Bruteforcing Host's Root Files...............🚀] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Start SQL, XSS & SSRF Detection & Eploitation......🐌] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Launch OWASP Nettacker project MainMenu............🦣] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Jok3r v3.5 Insane Mult Reconing Engine Main Menu...👊] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Osmedeus Next Generation Workflow Engine Main Menu.🍈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Ufonet Next Generation DDOS Tool Main Menu.........🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Launch Heavy Automation Attacks On The Host........🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit Web Scanner & Go To Main Menu...................] \n" + bcolors.ENDC)
}

func menu_8_4() {
    fmt.Println(bcolors.BLUE + "\n      ~>[ " + bcolors.ENDC + bcolors.UNDERL + "💉Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. ............Sql injection with sqlmap..." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🐞] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. .......Xss injection with Dalfox, Xsser." + bcolors.DARKCYAN + "(Automated)" + bcolors.BLUE + "🪰] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. ............Sql injection with Sqlmap......" + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🐛] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. ............Xss injection with Xsser......." + bcolors.DARKCYAN + "(Manual)" + bcolors.BLUE + "🐌] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. .............Exit & Go To Main Menu..................] \n" + bcolors.ENDC)
}

func menu_8_5() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Start Nettacker port & web content discovery.......📡] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Recon, find and scan subdomains....................🐾] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Launch admin_scan to find admin panel..............🦨] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Start insane information gathering on host.........🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks...........🧭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Run CVE scans on the target host...................🍹] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Search for (critical vulns & easy to exploit)......🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Automate all modules & security checks on target...🥑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. View scanned Nettacker report list.................🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)
}

func menu_8_6() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Install all jok3r tools..." + bcolors.DARKCYAN + "(Pleas start here if not)" + bcolors.BLUE + "🃏] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Update all the tools in the toolbox................🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Show all the tools in the toolbox..................🍒] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Show supported products for all services...........🍵] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Run only recon & vulnscansecurity checks...........🧭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Run all & intense security checks against an URL...🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Search for (critical vulns & easy to exploit)......🌈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. View the full resultsfrom the security checks......🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Clean database & delete results....................🥑] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)
}

func menu_8_7() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Update Osmedeus & Run diagnostics to check config..🍄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Start a simple scan with other flow................🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Launch vuln and directory scan on domains..........🧭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Scan list of targets (Full path of target needed)..🥐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Cloud - Run scan in Distributed Cloud mode.........🌏] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Performing Full web vuln & secret scan on host.....🏹] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Update the vulnerability database before attacking.📜] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Start web UI server................................🪐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. View scanned osmedeus report list..................🌈] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)
}

func menu_eight_8() {
    fmt.Println(bcolors.BLUE + "\n        ~>[ " + bcolors.ENDC + bcolors.UNDERL + "Select a number from the table below" + bcolors.ENDC + bcolors.BLUE + " ]<~\n" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 1. Download list of " + bcolors.RED + "Bots " + bcolors.BLUE + "from" + bcolors.YELLOW + " Community " + bcolors.BLUE + "server.......🍄.]" + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 2. Test if bots are alive............................📡.] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 3. Launch palantir..................." + bcolors.DARKCYAN + "(Palantir 3.14..)" + bcolors.BLUE + "🧭] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 4. Launch socking_waves.............." + bcolors.DARKCYAN + "(Knockout!......)" + bcolors.BLUE + "🦠] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 5. Launch xcom-1....................." + bcolors.DARKCYAN + "(Only DDoS......)" + bcolors.BLUE + "🥐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 6. Launch xcom-2....................." + bcolors.DARKCYAN + "(Only DoS.......)" + bcolors.BLUE + "🍵] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 7. Launch ufonet-gui................." + bcolors.DARKCYAN + "(Gui on browser.)" + bcolors.BLUE + "🧄] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 8. Start Grider.............." + bcolors.DARKCYAN + "........(Grider.........)" + bcolors.BLUE + "🪐] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 9. Launch Armageddon!........" + bcolors.DARKCYAN + "........(Launch ALL!!...)" + bcolors.BLUE + "🥝] " + bcolors.ENDC)
    fmt.Println(bcolors.BLUE + "[ 0. Exit & Go To Main Menu...............................] \n" + bcolors.ENDC)
}

func screenClear() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
            fmt.Println("Error running creenClear:", err)
    }
}

func main() {
    screenClear()
    scanner := bufio.NewScanner(os.Stdin)
    for {
        printBanner(); menu_main()
        fmt.Print(bcolors.GREEN + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework" + bcolors.GREEN + ")# " + bcolors.ENDC)
        scanner.Scan()
        userInput := scanner.Text()

        switch strings.ToLower(userInput) {
        case "0":
            screenClear(); exitBanner()
            return
        case "1":
            screenClear(); menu_1(); internals.RunNmap()
        case "2":
            screenClear(); menu_2(); internals.RunNmap()
        case "3":
            screenClear(); menu_3()
        case "4":
            screenClear(); menu_4()
        case "5":
            screenClear(); menu_5()
        case "6":
            screenClear(); menu_6()
        case "7":
            screenClear(); menu_7()
        case "8":
            screenClear(); menu_8()
        case "9":
            screenClear(); menu_8()
        default:
            fmt.Println(bcolors.BLUE + "\n~>[ " + bcolors.RED + "Poor choice of selection. Please select ~> " + bcolors.DARKCYAN + "from 0 to 9 " + bcolors.BLUE + "]<~" + bcolors.ENDC)
        }
    }
}

