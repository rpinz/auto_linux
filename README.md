                       #                                             ***********
                      ###                                        **************
                     #####                                     ******** ******
                    #########                                  *****  *******
                  ##########                                  ***   ********
               ########                                       *   ********
            ##########      ###                                 ******
         ############    ########                                ########
        ############    ##########    #             ##########       ########
      ##############    ###########   ##         ################        ######
     ###############    ##########    ####      ######         ###          #####
     ################    ########    ######    #####             ##          #####
    ###############                 ########  #####                           ####
    ###########                   ##########  #####                           #####
    ##########                  ############  #####                           #####
    ########       ########      ###########  #####                           #####
    #######      ############     ##########  #####                          ######
     ######     ##############     ########   #####                          ######
      #####     ##############     #######     #######                      ######
       ####     ##############     ######       ########                  #######
        ###      ############     ######           ###########################
          ##      ##########      ####               #######################
            #                    ###                    ##################
                               ##                            #######
# auto_linux

Installers for Drupal and OrangeHRM.

#### Make the installers:
``` bash
make
```

#### Install Drupal:
``` bash
./drupal_auto_linux:
  --dbuser="drupal"
  --dbname="drupal"
  --dbhost="localhost"
  --dbport="3306"
  --admin="admin"
  --site_mail="noone@nowhere.no"
  --site_name="name"
  --account_pass="ubuntu"
```

#### Install OrangeHRM:
``` bash
./ohrm_auto_linux:
  --dbuser="drupal"
  --dbname="drupal"
  --dbhost="localhost"
  --dbport="3306"
  --admin="admin"
  --account_pass="ubuntu"
```
