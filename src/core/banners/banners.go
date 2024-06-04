package banners

import (
    "time"
    "bcolors"
    "math/rand"
)

type Banners struct{}

func Graphics() {
    rand.Seed(time.Now().UnixNano())
    menu := rand.Intn(8) + 1

    switch menu {
    case 1:
        bcolors.Colors(`
                     _,._
                 __.'   _)
                <_,)'.-"a\
                  /' (    \
      _.-----..,-'   ('"--^
     //              |
    (|   ';      ,   |
      \   ;.----/  , /
       ) // /   | |\ \       John 3:16
       \ \\'\   | |/ /      Jesus Christ
        \ \\ \  | |\/  The Lamb that was slain
         '" '"   '"         for our sins.
` + bcolors.ENDC)

    case 2:
        bcolors.Colors(`
                 _      xxxx      _
                /_;-.__ / _\  _.-;_\
                   '-._ ' _/''.-'
                       '\   /'
                        |  /
                       /-.(
                       \_._\
                        \ \';
                         > |/       John 3:16
                        / //      Jesus Christ
                        |//   The Lamb that was slain
                        \(\       for our sins.
` + bcolors.ENDC)

    case 3:
        bcolors.Colors(`
                               .--,
                           ,.-( (o)\
                          /   .)/\ ')
                        .',./'/   )/
                    ()=///=))))==()
                      /
                      wake up, Christian
              Lord God Jesus Christ L❤️.VE'S you
                  follow the white Pigeon.
                    knock, knock, knock,
                          Man Of God.
` + bcolors.ENDC)

    case 4:
        bcolors.Colors(`
                                       ___
             _______                  /__/
            |.-----.|            ,---[___]*
            ||     ||           /    printer
            ||_____||    _____ /        ____
            |o_____*|   [o_+_+]--------[=i==]
             |     ________|           drive
             |  __|_
             '-/_==_\       Jesus Christ is
              /_____\\  The Lamb that was slain
                             for our sins.
                               John 3:16
` + bcolors.ENDC)

    case 5:
        bcolors.Colors(`
           __________   __________ __________
          |          |\|          |          |\
          |  *    *  |||  *  *  * |        * ||
          |  *    *  |||          |     *    ||
          |  *    *  |||  *  *  * |  *       ||
          |__________|||__________|__________||
          |          || '---------------------'
          |  *    *  ||       Jesus Christ
          |          ||  The Lamb that was slain
          |  *    *  ||       for our sins.
          |__________||         John 3:16
           '----------'
` + bcolors.ENDC)

    case 6:
        bcolors.Colors(`
                              ____
                       __,-~~/~   '---.
                     _/_,---(      ,   )_
                 ___/        <    /   )  \___
   - -----===;;;'===----------------------===';;;===----- -
                    \/  ~"~"~"~"~"~\~"~)~"/
                    (_ (   \  (     >    \)
                      \_( _ <         >_>'
     Jesus Christ is    ~ '-i' ::>|--"
 The Lamb that was slain    I;|.|.|
       for our sins.       <|i::|i|'.
        John 3:16         (' ^'"'-' ")
` + bcolors.ENDC)

    case 7:
        bcolors.Colors(`
                       .--.       .--.
                   _  '    \     /    '  _
                    '\.===. \.^./ .===./'
                           \/'"'\/
                        ,  |     |  ,
                       / '\|'-.-'|/' \
                      /    |  \  |    \
                   .-' ,-''|   ; |''-, '-.
                       |   |    \|   |
                       |   |    ;|   |
     Jesus Christ is   |   \    //   |
The Lamb that was slain|    '._//'   |
     for our sins.    .'             '.
       John 3:16   _,'                 ',_
` + bcolors.ENDC)

    case 8:
        bcolors.Colors(`
                   _,.---.---.---.--.._ 
               _.-' '--.'---.'---'-. _,'--.._
              /'--._ .'.     '.     ','-.'-._\
             ||   \  '.'---.__'__..-'. ,''-._/
        _  ,'\ '-._\   \    '.    '_.-'-._,''-.
     ,'   '-_ \/ '-.'--.\    _\_.-'\__.-'-.'-._'.
    (_.o> ,--. '._/'--.-',--'  \_.-'       \'-._ \
     '---'    '._ '---._/__,----'           '-. '-\
               /_, ,  _..-'                    '-._\
               \_, \/ ._(      Jesus Christ is
                \_, \/ ._\ The Lamb that was slain
                 '._,\/ ._\     for our sins.
                   '._// ./'-._
                     '-._-_-_.-'
` + bcolors.ENDC)
    }
}

func Banner() {
    Graphics()
}


