package setups

import (
    "fmt"
    "os"
    "menus"
    "utils"
    "bcolors"
    "banners"
    "scriptures"
    "subprocess"
)

var (
    userInput string
)

func installFoundationTools(commands []string) {
    for _, cmd := range commands {
        subprocess.Popen(cmd)
        fmt.Println()
    }
}

func installGithubTools() {
    githubCommands := []string{
        `cd /root/.africana/; git clone https://github.com/r0jahsm0ntar1/africana-base.git --depth 1`,
        `pip3 install --upgrade setuptools`,
        `pip3 install -r /root/.africana/africana-base/requirements.txt`,
        `go install github.com/j3ssie/osmedeus@latest`,
        `go install github.com/hahwul/dalfox/v2@latest`,
        `bash <(curl -fsSL https://raw.githubusercontent.com/osmedeus/osmedeus-base/master/install.sh)`,
    }

    for _, cmd := range githubCommands {
        subprocess.Popen(cmd)
        fmt.Println()
    }
}

func promptForUpdate() {
    for {
        fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.DARKCYAN + "🧬Africana already installed. " + bcolors.YELLOW + "Update it? " + bcolors.RED + "(y/n)" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Installer" + bcolors.BLUE + ")" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "\n╰─🩺" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        fmt.Scan(&userInput)
        switch userInput {
        case "y", "Y", "yes", "Yes", "YES":
            subprocess.Popen(`cd /root/.africana/africana-base; git pull .`)
            subprocess.Popen(`git clone https://github.com/r0jahsm0ntar1/africana-framework --depth 1`)
            subprocess.Popen(`cd africana-framework; go build -v -x ./africana.go; mv africana /usr/local/bin; rm -rf ../africana-framework`)
            fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Updated. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
            return
        case "n", "N", "no", "No", "NO":
            utils.ClearScreen()
            banners.Banner()
            scriptures.Verse()
            return
        default:
            fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
        }
    }
}

func KaliSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)\n" + bcolors.ENDC)
        foundationCommands := []string{
            `apt-get update -y`,
            `apt-get install zsh git curl wget -y`,
            `mkdir -p /etc/apt/trusted.gpg.d/`,
            `curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > /etc/apt/trusted.gpg.d/playit.gpg`,
            `curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o /etc/apt/sources.list.d/playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
            `apt-get install -y tor squid privoxy iptables tmux openssh-client libpcap-dev npm openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux netexec aha set playit libssl-dev gcc hydra wine32:i386`,
            `winecfg /v win11`,
        }
        installFoundationTools(foundationCommands)
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing Github third party tools" + bcolors.ENDC)
        installGithubTools()
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        promptForUpdate()
    }
}

func UbuntuSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)\n" + bcolors.ENDC)
        foundationCommands := []string{
            `apt-get update -y`,
            `apt-get install zsh git curl wget -y`,
            `wget "https://archive.kali.org/archive-key.asc"; apt-key add ./archive-key.asc; rm -rf ./archive-key.asc`,
            `echo -n "Package: *" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin: release a=kali-rolling" >> /etc/apt/preferences.d/kali.pref; echo -n "Pin-Priority: 50" >> /etc/apt/preferences.d/kali.pref`,
            `mkdir -p /etc/apt/trusted.gpg.d/`,
            `curl -vSL https://playit-cloud.github.io/ppa/key.gpg | gpg --dearmor > /etc/apt/trusted.gpg.d/playit.gpg`,
            `curl -vSL https://playit-cloud.github.io/ppa/playit-cloud.list -o /etc/apt/sources.list.d/playit-cloud.list`,
            `dpkg --add-architecture i386`,
            `apt-get update -y`,
            `apt-get install -y tor squid privoxy iptables tmux openssh-client libpcap-dev npm openssh-server ftp ncat rlwrap powershell golang-go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser util-linux netexec aha set playit libssl-dev gcc hydra wine32:i386`,
            `winecfg /v win11`,
        }
        installFoundationTools(foundationCommands)
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing Github third party tools" + bcolors.ENDC)
        installGithubTools()
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        promptForUpdate()
    }
}

func ArchSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)" + bcolors.ENDC)
        foundationCommands := []string{
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm zsh git curl wget go`,
        }
        installFoundationTools(foundationCommands)
        BlackArchSetups()
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing Github third party tools" + bcolors.ENDC)
        installGithubTools()
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        promptForUpdate()
    }
}

func BlackArchSetups() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing foundation tools " + bcolors.BLUE + "Eg." + bcolors.YELLOW + "(curl, wget, Go)\n" + bcolors.ENDC)
        foundationCommands := []string{
            `curl -O https://blackarch.org/strap.sh`,
            `chmod +x strap.sh`,
            `./strap.sh`,
            `pacman -Syu --noconfirm`,
            `pacman -S --noconfirm blackarch`,
            `pacman -S --noconfirm base-devel tor squid privoxy iptables tmux openssh-client libpcap-dev npm openssh-server ftp ncat rlwrap go docker.io python3 python3-pip build-essential libssl-dev libffi-dev python3-dev python3-venv python3-pycurl python3-geoip python3-whois python3-requests python3-scapy libgeoip1t64 libgeoip-dev privoxy dnsmasq gophish wifipumpkin3 wifite airgeddon nuclei nikto nmap smbmap dnsrecon metasploit-framework dnsrecon feroxbuster dirsearch uniscan sqlmap commix dnsenum sslscan whatweb wafw00f wordlists wapiti xsser powershell-empire util-linux netexec aha set playit libssl-dev gcc hydra wine32:i386`,
            `winecfg /v win11`,
        }
        installFoundationTools(foundationCommands)
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Installing Github third party tools" + bcolors.ENDC)
        installGithubTools()
        fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana Fully Installed. " + bcolors.YELLOW + "Safe Hacking!" + bcolors.ENDC)
    } else {
        promptForUpdate()
    }
}

func WindowsSetups() {
    // Placeholder for Windows setup logic
}

func Uninstall() {
    filePath := "/root/.africana/africana-base/"
    if _, err := os.Stat(filePath); !os.IsNotExist(err) {
        for {
            fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.DARKCYAN + "🧬Are you sure you want to uninstall Africana? " + bcolors.RED + "(y/n)" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╭─(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.DARKGREY + bcolors.ITALIC + "Uninstaller" + bcolors.BLUE + ")" + bcolors.ENDC)
            fmt.Printf(bcolors.BLUE + "\n╰─🩺" + bcolors.GREEN + "❯ " + bcolors.ENDC)
            fmt.Scan(&userInput)
            switch userInput {
            case "y", "Y", "yes", "Yes", "YES":
                subprocess.Popen(`rm -rf /root/.africana/; rm -rf /usr/local/bin/africana`)
                fmt.Println(bcolors.ENDC + `¯\_(ツ)_/¯` + bcolors.DARKCYAN + "🧬Africana successfully uninstalled. " + bcolors.YELLOW + "Goodbye!" + bcolors.ENDC)
                return
            case "n", "N", "no", "No", "NO":
                utils.ClearScreen()
                banners.Banner()
                scriptures.Verse()
                menus.MenuOne()
                return
            default:
                fmt.Println(bcolors.GREEN + "           (" + bcolors.DARKCYAN + "Choices are [" + bcolors.ENDC + "y|n|Y|N|yes|no|YES|NO" + bcolors.DARKCYAN + "]" + bcolors.GREEN + ")" + bcolors.ENDC)
            }
        }
    } else {
        fmt.Println(bcolors.ENDC + "(҂`_´) " + bcolors.RED + "Africana is not installed. " + bcolors.ENDC)
    }
}
