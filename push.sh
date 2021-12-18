#/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
if [[ $(uname -m)!="armv7" ]]; then echo -e "\n${RED}ERROR: build need to be done on 'ARMV7' plattform${NC}\n"; exit 1; fi
(./build.sh && docker push www.greisslomat.at:5000/gm65:1.0 && echo -e "\n${GREEN}push successfull${NC}\n") || echo -e "\n${RED}push failed${NC}\n"
